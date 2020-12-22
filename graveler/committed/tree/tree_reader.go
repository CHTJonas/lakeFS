package tree

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/treeverse/lakefs/graveler"
	"github.com/treeverse/lakefs/graveler/committed/sstable"
)

// treeRepo is a singleton containing the caches and common data. Will certainly be replaced once integrated into
// lakeFS startup
type treeRepo struct {
	treesMap   Cache
	partManger sstable.Manager
	pathBase   string
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}

type treeIterator struct {
	tree             []part
	currentIter      graveler.ValueIterator
	currentPartIndex int // index of current part
	err              error
	closed           bool
	trees            *treeRepo // referenc to tree repo needed in some actions on a tree
}

// InitTreeRepository creates the tree cache, and stores part Manager for operations of parts (currently implemented as sstables).
// should be called at process init.
// decisions on who calls it and how to get a treesRepository will be taken later
func InitTreesRepository(cache Cache, manager sstable.Manager) *treeRepo {
	return &treeRepo{
		//treesMap:   cache.NewCacheMap(CacheMapSize, CacheTrimSize, InitialWeight, AdditionalWeight, TrimFactor),
		treesMap:   cache,
		partManger: manager,
		pathBase:   "testdata/",
	}
}

func (trees *treeRepo) GetValue(treeID graveler.TreeID, key graveler.Key) (*graveler.ValueRecord, error) {
	tree, err := trees.GetTree(treeID)
	if err != nil {
		return nil, err
	}
	partIterator, _, err := trees.getSSTIteratorForKey(tree, key)
	if err != nil {
		return nil, err
	}
	defer partIterator.Close()
	partIterator.SeekGE(key)
	if partIterator.Next() {
		val := partIterator.Value()
		if bytes.Equal(val.Key, key) {
			return val, nil
		} else {
			return nil, ErrNotFound
		}
	}
	return nil, partIterator.Err()
}

func (trees *treeRepo) NewIteratorFromTreeID(treeID graveler.TreeID, start graveler.Key) (graveler.ValueIterator, error) {
	tree, err := trees.GetTree(treeID)
	if err != nil {
		return nil, err
	}
	return trees.newIterator(tree, start)
}

func (trees *treeRepo) NewIteratorFromTreeObject(treeSlice Tree, start graveler.Key) (graveler.ValueIterator, error) {
	return trees.newIterator(treeSlice, start)
}

func (trees *treeRepo) newIterator(tree Tree, from graveler.Key) (graveler.ValueIterator, error) {
	partIterator, partIndex, err := trees.getSSTIteratorForKey(tree, from)
	if err != nil {
		return nil, err
	}
	scanner := &treeIterator{
		tree:             tree.treeSlice,
		currentIter:      partIterator,
		currentPartIndex: partIndex,
		trees:            trees,
	}
	return scanner, nil
}

func (trees *treeRepo) getSSTIteratorForKey(tree Tree, key graveler.Key) (graveler.ValueIterator, int, error) {
	treeSlice := tree.treeSlice
	partIndex := findPartIndexForPath(treeSlice, key)
	if partIndex >= len(treeSlice) {
		return nil, 0, ErrPathBiggerThanMaxPath
	}
	partName := treeSlice[partIndex].PartName
	partIterator, err := trees.partManger.NewSSTableIterator(partName, key)
	if err != nil {
		return nil, 0, err
	}
	return partIterator, partIndex, nil
}

func (trees *treeRepo) GetTree(treeID graveler.TreeID) (Tree, error) {
	t, exists := trees.treesMap.Get(string(treeID))
	if exists {
		tree := t.(Tree)
		return tree, nil
	}
	fName := filepath.Join(trees.pathBase, string(treeID)+".json")
	jsonBytes, err := ioutil.ReadFile(fName)
	if err != nil {
		return Tree{}, err
	}
	treeSlice := make([]part, 0)
	jsonTreeSlice := make([]jsonTreePart, 0)

	err = json.Unmarshal(jsonBytes, &jsonTreeSlice)
	if err != nil {
		return Tree{}, err
	}
	for _, p := range jsonTreeSlice {
		treeSlice = append(treeSlice, part{
			PartName:        sstable.ID(p.PartName),
			MaxKey:          graveler.Key(p.MaxKey),
			MinKey:          graveler.Key(p.MinKey),
			NumberOfRecords: p.NumberOfRecords,
		})
	}
	tree := Tree{treeSlice: treeSlice}
	trees.treesMap.Set(string(treeID), tree)
	return tree, nil
}

func (t *treeIterator) SeekGE(start graveler.Key) {
	var err error
	partIndex := findPartIndexForPath(t.tree, start)
	if partIndex != t.currentPartIndex {
		t.currentPartIndex = partIndex
		t.currentIter.Close()
		t.currentIter, err = t.trees.partManger.NewSSTableIterator(t.tree[partIndex].PartName, start)
		if err != nil {
			t.err = err
			return
		}
	}
	t.currentIter.SeekGE(start)
}

func findPartIndexForPath(tree []part, path graveler.Key) int {
	n := len(tree)
	pos := sort.Search(n, func(i int) bool {
		return bytes.Compare(tree[i].MaxKey, path) >= 0
	})
	return pos
}

func (t *treeIterator) Next() bool {
	var err error
	if t.closed {
		return false
	}
	if t.currentIter.Next() {
		return true
	}
	t.err = t.currentIter.Err()
	t.currentIter.Close()
	// assert: if Next returned false and err == nil - reached end of part
	if t.err != nil {
		t.closed = true
		return false
	}
	// assert:  if Next returned false and err == nil - reached end of part
	if t.currentPartIndex >= len(t.tree)-1 {
		t.closed = true
		return false
	}
	t.currentPartIndex++
	requiredPartName := t.tree[t.currentPartIndex].PartName
	t.currentIter, err = t.trees.partManger.NewSSTableIterator(requiredPartName, nil)
	if err != nil {
		t.currentIter.Close()
		t.closed = true
		t.err = err
		return false
	}
	return t.currentIter.Next()
}

func (t *treeIterator) Err() error {
	if t.currentIter == nil {
		return ErrScannerIsNil
	}
	return t.currentIter.Err()
}

func (t *treeIterator) Value() *graveler.ValueRecord {
	if t.currentIter == nil || t.closed {
		return nil
	}
	return t.currentIter.Value()
}

func (t *treeIterator) Close() {
	if t.currentIter == nil {
		return
	}
	t.currentIter.Close()
	t.closed = true
}

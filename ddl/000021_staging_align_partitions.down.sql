BEGIN;
ALTER TABLE graveler_staging_kv_p0 RENAME TO kv_staging_p0;
ALTER TABLE graveler_staging_kv_p1 RENAME TO kv_staging_p1;
ALTER TABLE graveler_staging_kv_p2 RENAME TO kv_staging_p2;
ALTER TABLE graveler_staging_kv_p3 RENAME TO kv_staging_p3;
ALTER TABLE graveler_staging_kv_p4 RENAME TO kv_staging_p4;
ALTER TABLE graveler_staging_kv_p5 RENAME TO kv_staging_p5;
ALTER TABLE graveler_staging_kv_p6 RENAME TO kv_staging_p6;
ALTER TABLE graveler_staging_kv_p7 RENAME TO kv_staging_p7;
ALTER TABLE graveler_staging_kv_p8 RENAME TO kv_staging_p8;
ALTER TABLE graveler_staging_kv_p9 RENAME TO kv_staging_p9;
COMMIT;

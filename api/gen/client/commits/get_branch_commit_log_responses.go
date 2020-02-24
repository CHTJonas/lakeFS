// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/treeverse/lakefs/api/gen/models"
)

// GetBranchCommitLogReader is a Reader for the GetBranchCommitLog structure.
type GetBranchCommitLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBranchCommitLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBranchCommitLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBranchCommitLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBranchCommitLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetBranchCommitLogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBranchCommitLogOK creates a GetBranchCommitLogOK with default headers values
func NewGetBranchCommitLogOK() *GetBranchCommitLogOK {
	return &GetBranchCommitLogOK{}
}

/*GetBranchCommitLogOK handles this case with default header values.

commit log
*/
type GetBranchCommitLogOK struct {
	Payload *GetBranchCommitLogOKBody
}

func (o *GetBranchCommitLogOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches/{branchId}/commits][%d] getBranchCommitLogOK  %+v", 200, o.Payload)
}

func (o *GetBranchCommitLogOK) GetPayload() *GetBranchCommitLogOKBody {
	return o.Payload
}

func (o *GetBranchCommitLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetBranchCommitLogOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBranchCommitLogUnauthorized creates a GetBranchCommitLogUnauthorized with default headers values
func NewGetBranchCommitLogUnauthorized() *GetBranchCommitLogUnauthorized {
	return &GetBranchCommitLogUnauthorized{}
}

/*GetBranchCommitLogUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBranchCommitLogUnauthorized struct {
	Payload interface{}
}

func (o *GetBranchCommitLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches/{branchId}/commits][%d] getBranchCommitLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBranchCommitLogUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *GetBranchCommitLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBranchCommitLogNotFound creates a GetBranchCommitLogNotFound with default headers values
func NewGetBranchCommitLogNotFound() *GetBranchCommitLogNotFound {
	return &GetBranchCommitLogNotFound{}
}

/*GetBranchCommitLogNotFound handles this case with default header values.

branch not found
*/
type GetBranchCommitLogNotFound struct {
	Payload *models.Error
}

func (o *GetBranchCommitLogNotFound) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches/{branchId}/commits][%d] getBranchCommitLogNotFound  %+v", 404, o.Payload)
}

func (o *GetBranchCommitLogNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBranchCommitLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBranchCommitLogDefault creates a GetBranchCommitLogDefault with default headers values
func NewGetBranchCommitLogDefault(code int) *GetBranchCommitLogDefault {
	return &GetBranchCommitLogDefault{
		_statusCode: code,
	}
}

/*GetBranchCommitLogDefault handles this case with default header values.

generic error response
*/
type GetBranchCommitLogDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get branch commit log default response
func (o *GetBranchCommitLogDefault) Code() int {
	return o._statusCode
}

func (o *GetBranchCommitLogDefault) Error() string {
	return fmt.Sprintf("[GET /repositories/{repositoryId}/branches/{branchId}/commits][%d] getBranchCommitLog default  %+v", o._statusCode, o.Payload)
}

func (o *GetBranchCommitLogDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBranchCommitLogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetBranchCommitLogOKBody get branch commit log o k body
swagger:model GetBranchCommitLogOKBody
*/
type GetBranchCommitLogOKBody struct {

	// results
	Results []*models.Commit `json:"results"`
}

// Validate validates this get branch commit log o k body
func (o *GetBranchCommitLogOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBranchCommitLogOKBody) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getBranchCommitLogOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBranchCommitLogOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBranchCommitLogOKBody) UnmarshalBinary(b []byte) error {
	var res GetBranchCommitLogOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
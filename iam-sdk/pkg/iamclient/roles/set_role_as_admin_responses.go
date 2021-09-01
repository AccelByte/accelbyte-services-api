// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetRoleAsAdminReader is a Reader for the SetRoleAsAdmin structure.
type SetRoleAsAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoleAsAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetRoleAsAdminNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetRoleAsAdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetRoleAsAdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSetRoleAsAdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetRoleAsAdminNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/roles/{roleId}/admin returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSetRoleAsAdminNoContent creates a SetRoleAsAdminNoContent with default headers values
func NewSetRoleAsAdminNoContent() *SetRoleAsAdminNoContent {
	return &SetRoleAsAdminNoContent{}
}

/*SetRoleAsAdminNoContent handles this case with default header values.

  Operation succeeded
*/
type SetRoleAsAdminNoContent struct {
}

func (o *SetRoleAsAdminNoContent) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/admin][%d] setRoleAsAdminNoContent ", 204)
}

func (o *SetRoleAsAdminNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetRoleAsAdminBadRequest creates a SetRoleAsAdminBadRequest with default headers values
func NewSetRoleAsAdminBadRequest() *SetRoleAsAdminBadRequest {
	return &SetRoleAsAdminBadRequest{}
}

/*SetRoleAsAdminBadRequest handles this case with default header values.

  Invalid request
*/
type SetRoleAsAdminBadRequest struct {
}

func (o *SetRoleAsAdminBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/admin][%d] setRoleAsAdminBadRequest ", 400)
}

func (o *SetRoleAsAdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetRoleAsAdminUnauthorized creates a SetRoleAsAdminUnauthorized with default headers values
func NewSetRoleAsAdminUnauthorized() *SetRoleAsAdminUnauthorized {
	return &SetRoleAsAdminUnauthorized{}
}

/*SetRoleAsAdminUnauthorized handles this case with default header values.

  Unauthorized access
*/
type SetRoleAsAdminUnauthorized struct {
}

func (o *SetRoleAsAdminUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/admin][%d] setRoleAsAdminUnauthorized ", 401)
}

func (o *SetRoleAsAdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetRoleAsAdminForbidden creates a SetRoleAsAdminForbidden with default headers values
func NewSetRoleAsAdminForbidden() *SetRoleAsAdminForbidden {
	return &SetRoleAsAdminForbidden{}
}

/*SetRoleAsAdminForbidden handles this case with default header values.

  Forbidden
*/
type SetRoleAsAdminForbidden struct {
}

func (o *SetRoleAsAdminForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/admin][%d] setRoleAsAdminForbidden ", 403)
}

func (o *SetRoleAsAdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetRoleAsAdminNotFound creates a SetRoleAsAdminNotFound with default headers values
func NewSetRoleAsAdminNotFound() *SetRoleAsAdminNotFound {
	return &SetRoleAsAdminNotFound{}
}

/*SetRoleAsAdminNotFound handles this case with default header values.

  Data not found
*/
type SetRoleAsAdminNotFound struct {
}

func (o *SetRoleAsAdminNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/admin][%d] setRoleAsAdminNotFound ", 404)
}

func (o *SetRoleAsAdminNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
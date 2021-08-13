// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminDeleteUserRolesV3Reader is a Reader for the AdminDeleteUserRolesV3 structure.
type AdminDeleteUserRolesV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteUserRolesV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminDeleteUserRolesV3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminDeleteUserRolesV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeleteUserRolesV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminDeleteUserRolesV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminDeleteUserRolesV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/v3/admin/namespaces/{namespace}/users/{userId}/roles returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminDeleteUserRolesV3NoContent creates a AdminDeleteUserRolesV3NoContent with default headers values
func NewAdminDeleteUserRolesV3NoContent() *AdminDeleteUserRolesV3NoContent {
	return &AdminDeleteUserRolesV3NoContent{}
}

/*AdminDeleteUserRolesV3NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminDeleteUserRolesV3NoContent struct {
}

func (o *AdminDeleteUserRolesV3NoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminDeleteUserRolesV3NoContent ", 204)
}

func (o *AdminDeleteUserRolesV3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteUserRolesV3BadRequest creates a AdminDeleteUserRolesV3BadRequest with default headers values
func NewAdminDeleteUserRolesV3BadRequest() *AdminDeleteUserRolesV3BadRequest {
	return &AdminDeleteUserRolesV3BadRequest{}
}

/*AdminDeleteUserRolesV3BadRequest handles this case with default header values.

  Invalid request
*/
type AdminDeleteUserRolesV3BadRequest struct {
}

func (o *AdminDeleteUserRolesV3BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminDeleteUserRolesV3BadRequest ", 400)
}

func (o *AdminDeleteUserRolesV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteUserRolesV3Unauthorized creates a AdminDeleteUserRolesV3Unauthorized with default headers values
func NewAdminDeleteUserRolesV3Unauthorized() *AdminDeleteUserRolesV3Unauthorized {
	return &AdminDeleteUserRolesV3Unauthorized{}
}

/*AdminDeleteUserRolesV3Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminDeleteUserRolesV3Unauthorized struct {
}

func (o *AdminDeleteUserRolesV3Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminDeleteUserRolesV3Unauthorized ", 401)
}

func (o *AdminDeleteUserRolesV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteUserRolesV3Forbidden creates a AdminDeleteUserRolesV3Forbidden with default headers values
func NewAdminDeleteUserRolesV3Forbidden() *AdminDeleteUserRolesV3Forbidden {
	return &AdminDeleteUserRolesV3Forbidden{}
}

/*AdminDeleteUserRolesV3Forbidden handles this case with default header values.

  Forbidden
*/
type AdminDeleteUserRolesV3Forbidden struct {
}

func (o *AdminDeleteUserRolesV3Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminDeleteUserRolesV3Forbidden ", 403)
}

func (o *AdminDeleteUserRolesV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteUserRolesV3NotFound creates a AdminDeleteUserRolesV3NotFound with default headers values
func NewAdminDeleteUserRolesV3NotFound() *AdminDeleteUserRolesV3NotFound {
	return &AdminDeleteUserRolesV3NotFound{}
}

/*AdminDeleteUserRolesV3NotFound handles this case with default header values.

  Data not found
*/
type AdminDeleteUserRolesV3NotFound struct {
}

func (o *AdminDeleteUserRolesV3NotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminDeleteUserRolesV3NotFound ", 404)
}

func (o *AdminDeleteUserRolesV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

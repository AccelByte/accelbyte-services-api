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

// AdminResetPasswordV2Reader is a Reader for the AdminResetPasswordV2 structure.
type AdminResetPasswordV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminResetPasswordV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminResetPasswordV2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminResetPasswordV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminResetPasswordV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminResetPasswordV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminResetPasswordV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminResetPasswordV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminResetPasswordV2NoContent creates a AdminResetPasswordV2NoContent with default headers values
func NewAdminResetPasswordV2NoContent() *AdminResetPasswordV2NoContent {
	return &AdminResetPasswordV2NoContent{}
}

/*AdminResetPasswordV2NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminResetPasswordV2NoContent struct {
}

func (o *AdminResetPasswordV2NoContent) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password][%d] adminResetPasswordV2NoContent ", 204)
}

func (o *AdminResetPasswordV2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminResetPasswordV2BadRequest creates a AdminResetPasswordV2BadRequest with default headers values
func NewAdminResetPasswordV2BadRequest() *AdminResetPasswordV2BadRequest {
	return &AdminResetPasswordV2BadRequest{}
}

/*AdminResetPasswordV2BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr><tr><td>20002</td><td>validation error</td></tr><tr><td>10142</td><td>new password cannot be same with original</td></tr><tr><td>10143</td><td>password not match</td></tr></table>
*/
type AdminResetPasswordV2BadRequest struct {
}

func (o *AdminResetPasswordV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password][%d] adminResetPasswordV2BadRequest ", 400)
}

func (o *AdminResetPasswordV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminResetPasswordV2Unauthorized creates a AdminResetPasswordV2Unauthorized with default headers values
func NewAdminResetPasswordV2Unauthorized() *AdminResetPasswordV2Unauthorized {
	return &AdminResetPasswordV2Unauthorized{}
}

/*AdminResetPasswordV2Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminResetPasswordV2Unauthorized struct {
}

func (o *AdminResetPasswordV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password][%d] adminResetPasswordV2Unauthorized ", 401)
}

func (o *AdminResetPasswordV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminResetPasswordV2Forbidden creates a AdminResetPasswordV2Forbidden with default headers values
func NewAdminResetPasswordV2Forbidden() *AdminResetPasswordV2Forbidden {
	return &AdminResetPasswordV2Forbidden{}
}

/*AdminResetPasswordV2Forbidden handles this case with default header values.

  Forbidden
*/
type AdminResetPasswordV2Forbidden struct {
}

func (o *AdminResetPasswordV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password][%d] adminResetPasswordV2Forbidden ", 403)
}

func (o *AdminResetPasswordV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminResetPasswordV2NotFound creates a AdminResetPasswordV2NotFound with default headers values
func NewAdminResetPasswordV2NotFound() *AdminResetPasswordV2NotFound {
	return &AdminResetPasswordV2NotFound{}
}

/*AdminResetPasswordV2NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type AdminResetPasswordV2NotFound struct {
}

func (o *AdminResetPasswordV2NotFound) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password][%d] adminResetPasswordV2NotFound ", 404)
}

func (o *AdminResetPasswordV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminResetPasswordV2InternalServerError creates a AdminResetPasswordV2InternalServerError with default headers values
func NewAdminResetPasswordV2InternalServerError() *AdminResetPasswordV2InternalServerError {
	return &AdminResetPasswordV2InternalServerError{}
}

/*AdminResetPasswordV2InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type AdminResetPasswordV2InternalServerError struct {
}

func (o *AdminResetPasswordV2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/password][%d] adminResetPasswordV2InternalServerError ", 500)
}

func (o *AdminResetPasswordV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

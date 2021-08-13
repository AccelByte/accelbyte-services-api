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

// UserVerificationReader is a Reader for the UserVerification structure.
type UserVerificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserVerificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserVerificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserVerificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUserVerificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserVerificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUserVerificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUserVerificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/namespaces/{namespace}/users/{userId}/verification returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUserVerificationNoContent creates a UserVerificationNoContent with default headers values
func NewUserVerificationNoContent() *UserVerificationNoContent {
	return &UserVerificationNoContent{}
}

/*UserVerificationNoContent handles this case with default header values.

  Operation succeeded
*/
type UserVerificationNoContent struct {
}

func (o *UserVerificationNoContent) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verification][%d] userVerificationNoContent ", 204)
}

func (o *UserVerificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserVerificationBadRequest creates a UserVerificationBadRequest with default headers values
func NewUserVerificationBadRequest() *UserVerificationBadRequest {
	return &UserVerificationBadRequest{}
}

/*UserVerificationBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type UserVerificationBadRequest struct {
}

func (o *UserVerificationBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verification][%d] userVerificationBadRequest ", 400)
}

func (o *UserVerificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserVerificationUnauthorized creates a UserVerificationUnauthorized with default headers values
func NewUserVerificationUnauthorized() *UserVerificationUnauthorized {
	return &UserVerificationUnauthorized{}
}

/*UserVerificationUnauthorized handles this case with default header values.

  Unauthorized access
*/
type UserVerificationUnauthorized struct {
}

func (o *UserVerificationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verification][%d] userVerificationUnauthorized ", 401)
}

func (o *UserVerificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserVerificationForbidden creates a UserVerificationForbidden with default headers values
func NewUserVerificationForbidden() *UserVerificationForbidden {
	return &UserVerificationForbidden{}
}

/*UserVerificationForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10152</td><td>verification code not found</td></tr><tr><td>10137</td><td>code is expired</td></tr><tr><td>10136</td><td>code is either been used or not valid anymore</td></tr><tr><td>10138</td><td>code not match</td></tr><tr><td>10149</td><td>verification contact type doesn't match</td></tr><tr><td>10148</td><td>verification code context doesn't match the required context</td></tr><tr><td>10162</td><td>invalid verification</td></tr></table>
*/
type UserVerificationForbidden struct {
}

func (o *UserVerificationForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verification][%d] userVerificationForbidden ", 403)
}

func (o *UserVerificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserVerificationNotFound creates a UserVerificationNotFound with default headers values
func NewUserVerificationNotFound() *UserVerificationNotFound {
	return &UserVerificationNotFound{}
}

/*UserVerificationNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10139</td><td>platform account not found</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type UserVerificationNotFound struct {
}

func (o *UserVerificationNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verification][%d] userVerificationNotFound ", 404)
}

func (o *UserVerificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserVerificationInternalServerError creates a UserVerificationInternalServerError with default headers values
func NewUserVerificationInternalServerError() *UserVerificationInternalServerError {
	return &UserVerificationInternalServerError{}
}

/*UserVerificationInternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type UserVerificationInternalServerError struct {
}

func (o *UserVerificationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verification][%d] userVerificationInternalServerError ", 500)
}

func (o *UserVerificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

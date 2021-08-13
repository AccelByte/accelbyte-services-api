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

// SendVerificationCodeReader is a Reader for the SendVerificationCode structure.
type SendVerificationCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVerificationCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSendVerificationCodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendVerificationCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSendVerificationCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSendVerificationCodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSendVerificationCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewSendVerificationCodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewSendVerificationCodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSendVerificationCodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/namespaces/{namespace}/users/{userId}/verificationcode returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSendVerificationCodeNoContent creates a SendVerificationCodeNoContent with default headers values
func NewSendVerificationCodeNoContent() *SendVerificationCodeNoContent {
	return &SendVerificationCodeNoContent{}
}

/*SendVerificationCodeNoContent handles this case with default header values.

  Operation succeeded
*/
type SendVerificationCodeNoContent struct {
}

func (o *SendVerificationCodeNoContent) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeNoContent ", 204)
}

func (o *SendVerificationCodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeBadRequest creates a SendVerificationCodeBadRequest with default headers values
func NewSendVerificationCodeBadRequest() *SendVerificationCodeBadRequest {
	return &SendVerificationCodeBadRequest{}
}

/*SendVerificationCodeBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type SendVerificationCodeBadRequest struct {
}

func (o *SendVerificationCodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeBadRequest ", 400)
}

func (o *SendVerificationCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeUnauthorized creates a SendVerificationCodeUnauthorized with default headers values
func NewSendVerificationCodeUnauthorized() *SendVerificationCodeUnauthorized {
	return &SendVerificationCodeUnauthorized{}
}

/*SendVerificationCodeUnauthorized handles this case with default header values.

  Unauthorized access
*/
type SendVerificationCodeUnauthorized struct {
}

func (o *SendVerificationCodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeUnauthorized ", 401)
}

func (o *SendVerificationCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeForbidden creates a SendVerificationCodeForbidden with default headers values
func NewSendVerificationCodeForbidden() *SendVerificationCodeForbidden {
	return &SendVerificationCodeForbidden{}
}

/*SendVerificationCodeForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10146</td><td>userID not match</td></tr></table>
*/
type SendVerificationCodeForbidden struct {
}

func (o *SendVerificationCodeForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeForbidden ", 403)
}

func (o *SendVerificationCodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeNotFound creates a SendVerificationCodeNotFound with default headers values
func NewSendVerificationCodeNotFound() *SendVerificationCodeNotFound {
	return &SendVerificationCodeNotFound{}
}

/*SendVerificationCodeNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr><tr><td>10171</td><td>email address not found</td></tr><tr><td>10139</td><td>platform account not found</td></tr></table>
*/
type SendVerificationCodeNotFound struct {
}

func (o *SendVerificationCodeNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeNotFound ", 404)
}

func (o *SendVerificationCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeConflict creates a SendVerificationCodeConflict with default headers values
func NewSendVerificationCodeConflict() *SendVerificationCodeConflict {
	return &SendVerificationCodeConflict{}
}

/*SendVerificationCodeConflict handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10140</td><td>user verified</td></tr><tr><td>10133</td><td>email already used</td></tr></table>
*/
type SendVerificationCodeConflict struct {
}

func (o *SendVerificationCodeConflict) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeConflict ", 409)
}

func (o *SendVerificationCodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeTooManyRequests creates a SendVerificationCodeTooManyRequests with default headers values
func NewSendVerificationCodeTooManyRequests() *SendVerificationCodeTooManyRequests {
	return &SendVerificationCodeTooManyRequests{}
}

/*SendVerificationCodeTooManyRequests handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20007</td><td>too many requests</td></tr></table>
*/
type SendVerificationCodeTooManyRequests struct {
}

func (o *SendVerificationCodeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeTooManyRequests ", 429)
}

func (o *SendVerificationCodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendVerificationCodeInternalServerError creates a SendVerificationCodeInternalServerError with default headers values
func NewSendVerificationCodeInternalServerError() *SendVerificationCodeInternalServerError {
	return &SendVerificationCodeInternalServerError{}
}

/*SendVerificationCodeInternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type SendVerificationCodeInternalServerError struct {
}

func (o *SendVerificationCodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/verificationcode][%d] sendVerificationCodeInternalServerError ", 500)
}

func (o *SendVerificationCodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

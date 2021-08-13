// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// PublicUpdateUserV2Reader is a Reader for the PublicUpdateUserV2 structure.
type PublicUpdateUserV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicUpdateUserV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicUpdateUserV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicUpdateUserV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicUpdateUserV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicUpdateUserV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewPublicUpdateUserV2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPublicUpdateUserV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /iam/v2/public/namespaces/{namespace}/users/{userId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicUpdateUserV2OK creates a PublicUpdateUserV2OK with default headers values
func NewPublicUpdateUserV2OK() *PublicUpdateUserV2OK {
	return &PublicUpdateUserV2OK{}
}

/*PublicUpdateUserV2OK handles this case with default header values.

  OK
*/
type PublicUpdateUserV2OK struct {
	Payload []*iamclientmodels.ModelUserResponse
}

func (o *PublicUpdateUserV2OK) Error() string {
	return fmt.Sprintf("[PATCH /iam/v2/public/namespaces/{namespace}/users/{userId}][%d] publicUpdateUserV2OK  %+v", 200, o.Payload)
}

func (o *PublicUpdateUserV2OK) GetPayload() []*iamclientmodels.ModelUserResponse {
	return o.Payload
}

func (o *PublicUpdateUserV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicUpdateUserV2BadRequest creates a PublicUpdateUserV2BadRequest with default headers values
func NewPublicUpdateUserV2BadRequest() *PublicUpdateUserV2BadRequest {
	return &PublicUpdateUserV2BadRequest{}
}

/*PublicUpdateUserV2BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr><tr><td>10131</td><td>invalid date of birth</td></tr><tr><td>10155</td><td>country is not defined</td></tr><tr><td>10154</td><td>country not found</td></tr><tr><td>10130</td><td>user under age</td></tr><tr><td>10132</td><td>invalid email address</td></tr></table>
*/
type PublicUpdateUserV2BadRequest struct {
}

func (o *PublicUpdateUserV2BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iam/v2/public/namespaces/{namespace}/users/{userId}][%d] publicUpdateUserV2BadRequest ", 400)
}

func (o *PublicUpdateUserV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicUpdateUserV2Unauthorized creates a PublicUpdateUserV2Unauthorized with default headers values
func NewPublicUpdateUserV2Unauthorized() *PublicUpdateUserV2Unauthorized {
	return &PublicUpdateUserV2Unauthorized{}
}

/*PublicUpdateUserV2Unauthorized handles this case with default header values.

  Unauthorized access
*/
type PublicUpdateUserV2Unauthorized struct {
}

func (o *PublicUpdateUserV2Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /iam/v2/public/namespaces/{namespace}/users/{userId}][%d] publicUpdateUserV2Unauthorized ", 401)
}

func (o *PublicUpdateUserV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicUpdateUserV2NotFound creates a PublicUpdateUserV2NotFound with default headers values
func NewPublicUpdateUserV2NotFound() *PublicUpdateUserV2NotFound {
	return &PublicUpdateUserV2NotFound{}
}

/*PublicUpdateUserV2NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10139</td><td>platform account not found</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type PublicUpdateUserV2NotFound struct {
}

func (o *PublicUpdateUserV2NotFound) Error() string {
	return fmt.Sprintf("[PATCH /iam/v2/public/namespaces/{namespace}/users/{userId}][%d] publicUpdateUserV2NotFound ", 404)
}

func (o *PublicUpdateUserV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicUpdateUserV2Conflict creates a PublicUpdateUserV2Conflict with default headers values
func NewPublicUpdateUserV2Conflict() *PublicUpdateUserV2Conflict {
	return &PublicUpdateUserV2Conflict{}
}

/*PublicUpdateUserV2Conflict handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10133</td><td>email already used</td></tr></table>
*/
type PublicUpdateUserV2Conflict struct {
}

func (o *PublicUpdateUserV2Conflict) Error() string {
	return fmt.Sprintf("[PATCH /iam/v2/public/namespaces/{namespace}/users/{userId}][%d] publicUpdateUserV2Conflict ", 409)
}

func (o *PublicUpdateUserV2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicUpdateUserV2InternalServerError creates a PublicUpdateUserV2InternalServerError with default headers values
func NewPublicUpdateUserV2InternalServerError() *PublicUpdateUserV2InternalServerError {
	return &PublicUpdateUserV2InternalServerError{}
}

/*PublicUpdateUserV2InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type PublicUpdateUserV2InternalServerError struct {
}

func (o *PublicUpdateUserV2InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /iam/v2/public/namespaces/{namespace}/users/{userId}][%d] publicUpdateUserV2InternalServerError ", 500)
}

func (o *PublicUpdateUserV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

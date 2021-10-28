// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// UpdateMyZipCodeReader is a Reader for the UpdateMyZipCode structure.
type UpdateMyZipCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMyZipCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMyZipCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMyZipCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMyZipCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateMyZipCodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /v1/public/namespaces/{namespace}/users/me/profiles/zipCode returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateMyZipCodeOK creates a UpdateMyZipCodeOK with default headers values
func NewUpdateMyZipCodeOK() *UpdateMyZipCodeOK {
	return &UpdateMyZipCodeOK{}
}

/*UpdateMyZipCodeOK handles this case with default header values.

  Successful operation
*/
type UpdateMyZipCodeOK struct {
	Payload *basicclientmodels.UserZipCode
}

func (o *UpdateMyZipCodeOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/public/namespaces/{namespace}/users/me/profiles/zipCode][%d] updateMyZipCodeOK  %+v", 200, o.Payload)
}

func (o *UpdateMyZipCodeOK) GetPayload() *basicclientmodels.UserZipCode {
	return o.Payload
}

func (o *UpdateMyZipCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.UserZipCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMyZipCodeBadRequest creates a UpdateMyZipCodeBadRequest with default headers values
func NewUpdateMyZipCodeBadRequest() *UpdateMyZipCodeBadRequest {
	return &UpdateMyZipCodeBadRequest{}
}

/*UpdateMyZipCodeBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type UpdateMyZipCodeBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *UpdateMyZipCodeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/public/namespaces/{namespace}/users/me/profiles/zipCode][%d] updateMyZipCodeBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMyZipCodeBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *UpdateMyZipCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMyZipCodeUnauthorized creates a UpdateMyZipCodeUnauthorized with default headers values
func NewUpdateMyZipCodeUnauthorized() *UpdateMyZipCodeUnauthorized {
	return &UpdateMyZipCodeUnauthorized{}
}

/*UpdateMyZipCodeUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type UpdateMyZipCodeUnauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdateMyZipCodeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/public/namespaces/{namespace}/users/me/profiles/zipCode][%d] updateMyZipCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMyZipCodeUnauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateMyZipCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMyZipCodeForbidden creates a UpdateMyZipCodeForbidden with default headers values
func NewUpdateMyZipCodeForbidden() *UpdateMyZipCodeForbidden {
	return &UpdateMyZipCodeForbidden{}
}

/*UpdateMyZipCodeForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permission</td></tr></table>
*/
type UpdateMyZipCodeForbidden struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdateMyZipCodeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/public/namespaces/{namespace}/users/me/profiles/zipCode][%d] updateMyZipCodeForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMyZipCodeForbidden) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateMyZipCodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

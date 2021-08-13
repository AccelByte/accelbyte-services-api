// Code generated by go-swagger; DO NOT EDIT.

package clients

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

// CreateClientReader is a Reader for the CreateClient structure.
type CreateClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClientCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateClientConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/clients returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCreateClientCreated creates a CreateClientCreated with default headers values
func NewCreateClientCreated() *CreateClientCreated {
	return &CreateClientCreated{}
}

/*CreateClientCreated handles this case with default header values.

  Created
*/
type CreateClientCreated struct {
	Payload *iamclientmodels.ClientmodelClientCreationResponse
}

func (o *CreateClientCreated) Error() string {
	return fmt.Sprintf("[POST /iam/clients][%d] createClientCreated  %+v", 201, o.Payload)
}

func (o *CreateClientCreated) GetPayload() *iamclientmodels.ClientmodelClientCreationResponse {
	return o.Payload
}

func (o *CreateClientCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ClientmodelClientCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClientBadRequest creates a CreateClientBadRequest with default headers values
func NewCreateClientBadRequest() *CreateClientBadRequest {
	return &CreateClientBadRequest{}
}

/*CreateClientBadRequest handles this case with default header values.

  Invalid request
*/
type CreateClientBadRequest struct {
}

func (o *CreateClientBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/clients][%d] createClientBadRequest ", 400)
}

func (o *CreateClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClientUnauthorized creates a CreateClientUnauthorized with default headers values
func NewCreateClientUnauthorized() *CreateClientUnauthorized {
	return &CreateClientUnauthorized{}
}

/*CreateClientUnauthorized handles this case with default header values.

  Unauthorized access
*/
type CreateClientUnauthorized struct {
}

func (o *CreateClientUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/clients][%d] createClientUnauthorized ", 401)
}

func (o *CreateClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClientForbidden creates a CreateClientForbidden with default headers values
func NewCreateClientForbidden() *CreateClientForbidden {
	return &CreateClientForbidden{}
}

/*CreateClientForbidden handles this case with default header values.

  Forbidden
*/
type CreateClientForbidden struct {
}

func (o *CreateClientForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/clients][%d] createClientForbidden ", 403)
}

func (o *CreateClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClientConflict creates a CreateClientConflict with default headers values
func NewCreateClientConflict() *CreateClientConflict {
	return &CreateClientConflict{}
}

/*CreateClientConflict handles this case with default header values.

  Client exists
*/
type CreateClientConflict struct {
}

func (o *CreateClientConflict) Error() string {
	return fmt.Sprintf("[POST /iam/clients][%d] createClientConflict ", 409)
}

func (o *CreateClientConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

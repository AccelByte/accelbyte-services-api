// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// CreateStoreReader is a Reader for the CreateStore structure.
type CreateStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateStoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateStoreConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateStoreUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /admin/namespaces/{namespace}/stores returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCreateStoreCreated creates a CreateStoreCreated with default headers values
func NewCreateStoreCreated() *CreateStoreCreated {
	return &CreateStoreCreated{}
}

/*CreateStoreCreated handles this case with default header values.

  successful operation
*/
type CreateStoreCreated struct {
	Payload *platformclientmodels.StoreInfo
}

func (o *CreateStoreCreated) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/stores][%d] createStoreCreated  %+v", 201, o.Payload)
}

func (o *CreateStoreCreated) GetPayload() *platformclientmodels.StoreInfo {
	return o.Payload
}

func (o *CreateStoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.StoreInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStoreConflict creates a CreateStoreConflict with default headers values
func NewCreateStoreConflict() *CreateStoreConflict {
	return &CreateStoreConflict{}
}

/*CreateStoreConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30174</td><td>Draft store already exists in namespace [{namespace}]</td></tr></table>
*/
type CreateStoreConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *CreateStoreConflict) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/stores][%d] createStoreConflict  %+v", 409, o.Payload)
}

func (o *CreateStoreConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *CreateStoreConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStoreUnprocessableEntity creates a CreateStoreUnprocessableEntity with default headers values
func NewCreateStoreUnprocessableEntity() *CreateStoreUnprocessableEntity {
	return &CreateStoreUnprocessableEntity{}
}

/*CreateStoreUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type CreateStoreUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *CreateStoreUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/stores][%d] createStoreUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateStoreUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *CreateStoreUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

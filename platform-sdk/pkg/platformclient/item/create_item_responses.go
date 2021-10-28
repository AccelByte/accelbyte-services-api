// Code generated by go-swagger; DO NOT EDIT.

package item

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

// CreateItemReader is a Reader for the CreateItem structure.
type CreateItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateItemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateItemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /admin/namespaces/{namespace}/items returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCreateItemCreated creates a CreateItemCreated with default headers values
func NewCreateItemCreated() *CreateItemCreated {
	return &CreateItemCreated{}
}

/*CreateItemCreated handles this case with default header values.

  successful operation
*/
type CreateItemCreated struct {
	Payload *platformclientmodels.FullItemInfo
}

func (o *CreateItemCreated) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/items][%d] createItemCreated  %+v", 201, o.Payload)
}

func (o *CreateItemCreated) GetPayload() *platformclientmodels.FullItemInfo {
	return o.Payload
}

func (o *CreateItemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.FullItemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateItemBadRequest creates a CreateItemBadRequest with default headers values
func NewCreateItemBadRequest() *CreateItemBadRequest {
	return &CreateItemBadRequest{}
}

/*CreateItemBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30322</td><td>Bundle item [{itemId}] can't be bundled</td></tr><tr><td>30325</td><td>Code item [{itemId}] can't be bundled</td></tr><tr><td>30326</td><td>Subscription item [{itemId}] can't be bundled</td></tr><tr><td>30328</td><td>Season item [{itemId}] can't be bundled</td></tr><tr><td>30021</td><td>Default language [{language}] required</td></tr><tr><td>30321</td><td>Invalid item discount amount</td></tr><tr><td>30022</td><td>Default region [{region}] is required</td></tr><tr><td>30323</td><td>Target namespace is required</td></tr><tr><td>30327</td><td>Invalid item trial price</td></tr></table>
*/
type CreateItemBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *CreateItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/items][%d] createItemBadRequest  %+v", 400, o.Payload)
}

func (o *CreateItemBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *CreateItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateItemNotFound creates a CreateItemNotFound with default headers values
func NewCreateItemNotFound() *CreateItemNotFound {
	return &CreateItemNotFound{}
}

/*CreateItemNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30241</td><td>Category [{categoryPath}] does not exist in namespace [{namespace}]</td></tr><tr><td>36141</td><td>Currency [{currencyCode}] does not exist in namespace [{namespace}]</td></tr><tr><td>30141</td><td>Store [{storeId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type CreateItemNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *CreateItemNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/items][%d] createItemNotFound  %+v", 404, o.Payload)
}

func (o *CreateItemNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *CreateItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateItemConflict creates a CreateItemConflict with default headers values
func NewCreateItemConflict() *CreateItemConflict {
	return &CreateItemConflict{}
}

/*CreateItemConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30173</td><td>Published store can't modify content</td></tr><tr><td>30373</td><td>ItemType [{itemType}] is not allowed in namespace [{namespace}]</td></tr></table>
*/
type CreateItemConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *CreateItemConflict) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/items][%d] createItemConflict  %+v", 409, o.Payload)
}

func (o *CreateItemConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *CreateItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateItemUnprocessableEntity creates a CreateItemUnprocessableEntity with default headers values
func NewCreateItemUnprocessableEntity() *CreateItemUnprocessableEntity {
	return &CreateItemUnprocessableEntity{}
}

/*CreateItemUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type CreateItemUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *CreateItemUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/items][%d] createItemUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateItemUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *CreateItemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// UpdateItemReader is a Reader for the UpdateItem structure.
type UpdateItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewUpdateItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewUpdateItemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/items/{itemId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateItemOK creates a UpdateItemOK with default headers values
func NewUpdateItemOK() *UpdateItemOK {
	return &UpdateItemOK{}
}

/*UpdateItemOK handles this case with default header values.

  successful operation
*/
type UpdateItemOK struct {
	Payload *platformclientmodels.FullItemInfo
}

func (o *UpdateItemOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}][%d] updateItemOK  %+v", 200, o.Payload)
}

func (o *UpdateItemOK) GetPayload() *platformclientmodels.FullItemInfo {
	return o.Payload
}

func (o *UpdateItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.FullItemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateItemBadRequest creates a UpdateItemBadRequest with default headers values
func NewUpdateItemBadRequest() *UpdateItemBadRequest {
	return &UpdateItemBadRequest{}
}

/*UpdateItemBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30322</td><td>Bundle item [{itemId}] can't be bundled</td></tr><tr><td>30325</td><td>Code item [{itemId}] can't be bundled</td></tr><tr><td>30326</td><td>Subscription item [{itemId}] can't be bundled</td></tr><tr><td>30328</td><td>Season item [{itemId}] can't be bundled</td></tr><tr><td>30021</td><td>Default language [{language}] required</td></tr><tr><td>30321</td><td>Invalid item discount amount</td></tr><tr><td>30022</td><td>Default region [{region}] is required</td></tr><tr><td>30323</td><td>Target namespace is required</td></tr><tr><td>30327</td><td>Invalid item trial price</td></tr></table>
*/
type UpdateItemBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *UpdateItemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}][%d] updateItemBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateItemBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateItemNotFound creates a UpdateItemNotFound with default headers values
func NewUpdateItemNotFound() *UpdateItemNotFound {
	return &UpdateItemNotFound{}
}

/*UpdateItemNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30141</td><td>Store [{storeId}] does not exist in namespace [{namespace}]</td></tr><tr><td>30341</td><td>Item [{itemId}] does not exist in namespace [{namespace}]</td></tr><tr><td>30241</td><td>Category [{categoryPath}] does not exist in namespace [{namespace}]</td></tr><tr><td>36141</td><td>Currency [{currencyCode}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type UpdateItemNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *UpdateItemNotFound) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}][%d] updateItemNotFound  %+v", 404, o.Payload)
}

func (o *UpdateItemNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateItemConflict creates a UpdateItemConflict with default headers values
func NewUpdateItemConflict() *UpdateItemConflict {
	return &UpdateItemConflict{}
}

/*UpdateItemConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30173</td><td>Published store can't modify content</td></tr><tr><td>30371</td><td>Item maxCount not allow reduce</td></tr><tr><td>30372</td><td>ItemType is not updatable</td></tr></table>
*/
type UpdateItemConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *UpdateItemConflict) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}][%d] updateItemConflict  %+v", 409, o.Payload)
}

func (o *UpdateItemConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateItemUnprocessableEntity creates a UpdateItemUnprocessableEntity with default headers values
func NewUpdateItemUnprocessableEntity() *UpdateItemUnprocessableEntity {
	return &UpdateItemUnprocessableEntity{}
}

/*UpdateItemUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type UpdateItemUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *UpdateItemUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}][%d] updateItemUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateItemUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *UpdateItemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

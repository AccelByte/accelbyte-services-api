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

// EnableItemReader is a Reader for the EnableItem structure.
type EnableItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnableItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewEnableItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/items/{itemId}/enable returns an error %d: %s", response.Code(), string(data))
	}
}

// NewEnableItemOK creates a EnableItemOK with default headers values
func NewEnableItemOK() *EnableItemOK {
	return &EnableItemOK{}
}

/*EnableItemOK handles this case with default header values.

  successful operation
*/
type EnableItemOK struct {
	Payload *platformclientmodels.FullItemInfo
}

func (o *EnableItemOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}/enable][%d] enableItemOK  %+v", 200, o.Payload)
}

func (o *EnableItemOK) GetPayload() *platformclientmodels.FullItemInfo {
	return o.Payload
}

func (o *EnableItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.FullItemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableItemNotFound creates a EnableItemNotFound with default headers values
func NewEnableItemNotFound() *EnableItemNotFound {
	return &EnableItemNotFound{}
}

/*EnableItemNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30141</td><td>Store [{storeId}] does not exist in namespace [{namespace}]</td></tr><tr><td>30341</td><td>Item [{itemId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type EnableItemNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *EnableItemNotFound) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}/enable][%d] enableItemNotFound  %+v", 404, o.Payload)
}

func (o *EnableItemNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *EnableItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableItemConflict creates a EnableItemConflict with default headers values
func NewEnableItemConflict() *EnableItemConflict {
	return &EnableItemConflict{}
}

/*EnableItemConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30173</td><td>Published store can't modify content</td></tr></table>
*/
type EnableItemConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *EnableItemConflict) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/items/{itemId}/enable][%d] enableItemConflict  %+v", 409, o.Payload)
}

func (o *EnableItemConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *EnableItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

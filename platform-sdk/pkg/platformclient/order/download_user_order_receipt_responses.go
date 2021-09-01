// Code generated by go-swagger; DO NOT EDIT.

package order

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

// DownloadUserOrderReceiptReader is a Reader for the DownloadUserOrderReceipt structure.
type DownloadUserOrderReceiptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadUserOrderReceiptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 404:
		result := NewDownloadUserOrderReceiptNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewDownloadUserOrderReceiptConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDownloadUserOrderReceiptNotFound creates a DownloadUserOrderReceiptNotFound with default headers values
func NewDownloadUserOrderReceiptNotFound() *DownloadUserOrderReceiptNotFound {
	return &DownloadUserOrderReceiptNotFound{}
}

/*DownloadUserOrderReceiptNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>32141</td><td>Order [{orderNo}] does not exist</td></tr></table>
*/
type DownloadUserOrderReceiptNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DownloadUserOrderReceiptNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf][%d] downloadUserOrderReceiptNotFound  %+v", 404, o.Payload)
}

func (o *DownloadUserOrderReceiptNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DownloadUserOrderReceiptNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadUserOrderReceiptConflict creates a DownloadUserOrderReceiptConflict with default headers values
func NewDownloadUserOrderReceiptConflict() *DownloadUserOrderReceiptConflict {
	return &DownloadUserOrderReceiptConflict{}
}

/*DownloadUserOrderReceiptConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>32173</td><td>Receipt of order [{orderNo}] is not downloadable</td></tr></table>
*/
type DownloadUserOrderReceiptConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DownloadUserOrderReceiptConflict) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf][%d] downloadUserOrderReceiptConflict  %+v", 409, o.Payload)
}

func (o *DownloadUserOrderReceiptConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DownloadUserOrderReceiptConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
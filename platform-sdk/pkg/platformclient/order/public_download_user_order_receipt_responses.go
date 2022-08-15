// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// PublicDownloadUserOrderReceiptReader is a Reader for the PublicDownloadUserOrderReceipt structure.
type PublicDownloadUserOrderReceiptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicDownloadUserOrderReceiptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicDownloadUserOrderReceiptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicDownloadUserOrderReceiptNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewPublicDownloadUserOrderReceiptConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicDownloadUserOrderReceiptOK creates a PublicDownloadUserOrderReceiptOK with default headers values
func NewPublicDownloadUserOrderReceiptOK() *PublicDownloadUserOrderReceiptOK {
	return &PublicDownloadUserOrderReceiptOK{}
}

/*PublicDownloadUserOrderReceiptOK handles this case with default header values.

  Successful operation
*/
type PublicDownloadUserOrderReceiptOK struct {
}

func (o *PublicDownloadUserOrderReceiptOK) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf][%d] publicDownloadUserOrderReceiptOK ", 200)
}

func (o *PublicDownloadUserOrderReceiptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicDownloadUserOrderReceiptNotFound creates a PublicDownloadUserOrderReceiptNotFound with default headers values
func NewPublicDownloadUserOrderReceiptNotFound() *PublicDownloadUserOrderReceiptNotFound {
	return &PublicDownloadUserOrderReceiptNotFound{}
}

/*PublicDownloadUserOrderReceiptNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>32141</td><td>Order [{orderNo}] does not exist</td></tr></table>
*/
type PublicDownloadUserOrderReceiptNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *PublicDownloadUserOrderReceiptNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf][%d] publicDownloadUserOrderReceiptNotFound  %+v", 404, o.ToJSONString())
}

func (o *PublicDownloadUserOrderReceiptNotFound) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicDownloadUserOrderReceiptNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicDownloadUserOrderReceiptNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicDownloadUserOrderReceiptConflict creates a PublicDownloadUserOrderReceiptConflict with default headers values
func NewPublicDownloadUserOrderReceiptConflict() *PublicDownloadUserOrderReceiptConflict {
	return &PublicDownloadUserOrderReceiptConflict{}
}

/*PublicDownloadUserOrderReceiptConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>32173</td><td>Receipt of order [{orderNo}] is not downloadable</td></tr></table>
*/
type PublicDownloadUserOrderReceiptConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *PublicDownloadUserOrderReceiptConflict) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf][%d] publicDownloadUserOrderReceiptConflict  %+v", 409, o.ToJSONString())
}

func (o *PublicDownloadUserOrderReceiptConflict) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicDownloadUserOrderReceiptConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicDownloadUserOrderReceiptConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

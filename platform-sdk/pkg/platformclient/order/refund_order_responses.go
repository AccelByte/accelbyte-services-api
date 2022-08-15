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

// RefundOrderReader is a Reader for the RefundOrder structure.
type RefundOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefundOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefundOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRefundOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewRefundOrderConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRefundOrderUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/admin/namespaces/{namespace}/orders/{orderNo}/refund returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRefundOrderOK creates a RefundOrderOK with default headers values
func NewRefundOrderOK() *RefundOrderOK {
	return &RefundOrderOK{}
}

/*RefundOrderOK handles this case with default header values.

  successful operation
*/
type RefundOrderOK struct {
	Payload *platformclientmodels.OrderInfo
}

func (o *RefundOrderOK) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/orders/{orderNo}/refund][%d] refundOrderOK  %+v", 200, o.ToString())
}

func (o *RefundOrderOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *RefundOrderOK) GetPayload() *platformclientmodels.OrderInfo {
	return o.Payload
}

func (o *RefundOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.OrderInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefundOrderNotFound creates a RefundOrderNotFound with default headers values
func NewRefundOrderNotFound() *RefundOrderNotFound {
	return &RefundOrderNotFound{}
}

/*RefundOrderNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>32141</td><td>Order [{orderNo}] does not exist</td></tr></table>
*/
type RefundOrderNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *RefundOrderNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/orders/{orderNo}/refund][%d] refundOrderNotFound  %+v", 404, o.ToString())
}

func (o *RefundOrderNotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *RefundOrderNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *RefundOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefundOrderConflict creates a RefundOrderConflict with default headers values
func NewRefundOrderConflict() *RefundOrderConflict {
	return &RefundOrderConflict{}
}

/*RefundOrderConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>32171</td><td>Order [{orderNo}] is not refundable</td></tr></table>
*/
type RefundOrderConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *RefundOrderConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/orders/{orderNo}/refund][%d] refundOrderConflict  %+v", 409, o.ToString())
}

func (o *RefundOrderConflict) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *RefundOrderConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *RefundOrderConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefundOrderUnprocessableEntity creates a RefundOrderUnprocessableEntity with default headers values
func NewRefundOrderUnprocessableEntity() *RefundOrderUnprocessableEntity {
	return &RefundOrderUnprocessableEntity{}
}

/*RefundOrderUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type RefundOrderUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *RefundOrderUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/orders/{orderNo}/refund][%d] refundOrderUnprocessableEntity  %+v", 422, o.ToString())
}

func (o *RefundOrderUnprocessableEntity) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *RefundOrderUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *RefundOrderUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

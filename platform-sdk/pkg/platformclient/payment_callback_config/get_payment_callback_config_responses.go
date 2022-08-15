// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package payment_callback_config

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

// GetPaymentCallbackConfigReader is a Reader for the GetPaymentCallbackConfig structure.
type GetPaymentCallbackConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentCallbackConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentCallbackConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPaymentCallbackConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/payment/config/callback returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetPaymentCallbackConfigOK creates a GetPaymentCallbackConfigOK with default headers values
func NewGetPaymentCallbackConfigOK() *GetPaymentCallbackConfigOK {
	return &GetPaymentCallbackConfigOK{}
}

/*GetPaymentCallbackConfigOK handles this case with default header values.

  successful operation
*/
type GetPaymentCallbackConfigOK struct {
	Payload *platformclientmodels.PaymentCallbackConfigInfo
}

func (o *GetPaymentCallbackConfigOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/payment/config/callback][%d] getPaymentCallbackConfigOK  %+v", 200, o.ToString())
}

func (o *GetPaymentCallbackConfigOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetPaymentCallbackConfigOK) GetPayload() *platformclientmodels.PaymentCallbackConfigInfo {
	return o.Payload
}

func (o *GetPaymentCallbackConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.PaymentCallbackConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentCallbackConfigNotFound creates a GetPaymentCallbackConfigNotFound with default headers values
func NewGetPaymentCallbackConfigNotFound() *GetPaymentCallbackConfigNotFound {
	return &GetPaymentCallbackConfigNotFound{}
}

/*GetPaymentCallbackConfigNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>33243</td><td>Payment callback config for [{namespace}] does not exist</td></tr></table>
*/
type GetPaymentCallbackConfigNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GetPaymentCallbackConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/payment/config/callback][%d] getPaymentCallbackConfigNotFound  %+v", 404, o.ToString())
}

func (o *GetPaymentCallbackConfigNotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetPaymentCallbackConfigNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetPaymentCallbackConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

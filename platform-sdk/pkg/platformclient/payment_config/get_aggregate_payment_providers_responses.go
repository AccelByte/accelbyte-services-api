// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package payment_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAggregatePaymentProvidersReader is a Reader for the GetAggregatePaymentProviders structure.
type GetAggregatePaymentProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAggregatePaymentProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAggregatePaymentProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/payment/config/provider/aggregate returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetAggregatePaymentProvidersOK creates a GetAggregatePaymentProvidersOK with default headers values
func NewGetAggregatePaymentProvidersOK() *GetAggregatePaymentProvidersOK {
	return &GetAggregatePaymentProvidersOK{}
}

/*GetAggregatePaymentProvidersOK handles this case with default header values.

  successful operation
*/
type GetAggregatePaymentProvidersOK struct {
	Payload []string
}

func (o *GetAggregatePaymentProvidersOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/payment/config/provider/aggregate][%d] getAggregatePaymentProvidersOK  %+v", 200, o.Payload)
}

func (o *GetAggregatePaymentProvidersOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAggregatePaymentProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

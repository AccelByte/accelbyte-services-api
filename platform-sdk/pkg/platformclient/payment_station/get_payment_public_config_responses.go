// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package payment_station

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPaymentPublicConfigReader is a Reader for the GetPaymentPublicConfig structure.
type GetPaymentPublicConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentPublicConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentPublicConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/public/namespaces/{namespace}/payment/publicconfig returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetPaymentPublicConfigOK creates a GetPaymentPublicConfigOK with default headers values
func NewGetPaymentPublicConfigOK() *GetPaymentPublicConfigOK {
	return &GetPaymentPublicConfigOK{}
}

/*GetPaymentPublicConfigOK handles this case with default header values.

  successful operation
*/
type GetPaymentPublicConfigOK struct {
	Payload map[string]interface{}
}

func (o *GetPaymentPublicConfigOK) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/payment/publicconfig][%d] getPaymentPublicConfigOK  %+v", 200, o.Payload)
}

func (o *GetPaymentPublicConfigOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetPaymentPublicConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

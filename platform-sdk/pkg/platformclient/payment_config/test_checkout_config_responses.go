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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// TestCheckoutConfigReader is a Reader for the TestCheckoutConfig structure.
type TestCheckoutConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestCheckoutConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestCheckoutConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /platform/admin/payment/config/merchant/checkoutconfig/test returns an error %d: %s", response.Code(), string(data))
	}
}

// NewTestCheckoutConfigOK creates a TestCheckoutConfigOK with default headers values
func NewTestCheckoutConfigOK() *TestCheckoutConfigOK {
	return &TestCheckoutConfigOK{}
}

/*TestCheckoutConfigOK handles this case with default header values.

  successful operation
*/
type TestCheckoutConfigOK struct {
	Payload *platformclientmodels.TestResult
}

func (o *TestCheckoutConfigOK) Error() string {
	return fmt.Sprintf("[POST /platform/admin/payment/config/merchant/checkoutconfig/test][%d] testCheckoutConfigOK  %+v", 200, o.Payload)
}

func (o *TestCheckoutConfigOK) GetPayload() *platformclientmodels.TestResult {
	return o.Payload
}

func (o *TestCheckoutConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.TestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package payment_config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// TestXsollaConfigReader is a Reader for the TestXsollaConfig structure.
type TestXsollaConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestXsollaConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestXsollaConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /platform/admin/payment/config/merchant/xsollaconfig/test returns an error %d: %s", response.Code(), string(data))
	}
}

// NewTestXsollaConfigOK creates a TestXsollaConfigOK with default headers values
func NewTestXsollaConfigOK() *TestXsollaConfigOK {
	return &TestXsollaConfigOK{}
}

/*TestXsollaConfigOK handles this case with default header values.

  successful operation
*/
type TestXsollaConfigOK struct {
	Payload *platformclientmodels.TestResult
}

func (o *TestXsollaConfigOK) Error() string {
	return fmt.Sprintf("[POST /platform/admin/payment/config/merchant/xsollaconfig/test][%d] testXsollaConfigOK  %+v", 200, o.ToJSONString())
}

func (o *TestXsollaConfigOK) ToJSONString() string {
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

func (o *TestXsollaConfigOK) GetPayload() *platformclientmodels.TestResult {
	return o.Payload
}

func (o *TestXsollaConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(platformclientmodels.TestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

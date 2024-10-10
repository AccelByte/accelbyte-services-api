// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package public

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// GetDefaultProviderReader is a Reader for the GetDefaultProvider structure.
type GetDefaultProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefaultProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/public/provider/default returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetDefaultProviderOK creates a GetDefaultProviderOK with default headers values
func NewGetDefaultProviderOK() *GetDefaultProviderOK {
	return &GetDefaultProviderOK{}
}

/*GetDefaultProviderOK handles this case with default header values.

  Default provider got
*/
type GetDefaultProviderOK struct {
	Payload *dsmcclientmodels.ModelsDefaultProvider
}

func (o *GetDefaultProviderOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/public/provider/default][%d] getDefaultProviderOK  %+v", 200, o.ToJSONString())
}

func (o *GetDefaultProviderOK) ToJSONString() string {
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

func (o *GetDefaultProviderOK) GetPayload() *dsmcclientmodels.ModelsDefaultProvider {
	return o.Payload
}

func (o *GetDefaultProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(dsmcclientmodels.ModelsDefaultProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

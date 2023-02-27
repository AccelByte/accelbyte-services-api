// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package stat_cycle_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// GetStatCycles1Reader is a Reader for the GetStatCycles1 structure.
type GetStatCycles1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatCycles1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatCycles1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /social/v1/public/namespaces/{namespace}/statCycles returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetStatCycles1OK creates a GetStatCycles1OK with default headers values
func NewGetStatCycles1OK() *GetStatCycles1OK {
	return &GetStatCycles1OK{}
}

/*GetStatCycles1OK handles this case with default header values.

  successful operation
*/
type GetStatCycles1OK struct {
	Payload *socialclientmodels.StatCyclePagingSlicedResult
}

func (o *GetStatCycles1OK) Error() string {
	return fmt.Sprintf("[GET /social/v1/public/namespaces/{namespace}/statCycles][%d] getStatCycles1OK  %+v", 200, o.ToJSONString())
}

func (o *GetStatCycles1OK) ToJSONString() string {
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

func (o *GetStatCycles1OK) GetPayload() *socialclientmodels.StatCyclePagingSlicedResult {
	return o.Payload
}

func (o *GetStatCycles1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(socialclientmodels.StatCyclePagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

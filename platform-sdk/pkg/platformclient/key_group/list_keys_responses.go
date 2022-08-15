// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package key_group

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

// ListKeysReader is a Reader for the ListKeys structure.
type ListKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}/keys returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListKeysOK creates a ListKeysOK with default headers values
func NewListKeysOK() *ListKeysOK {
	return &ListKeysOK{}
}

/*ListKeysOK handles this case with default header values.

  successful operation
*/
type ListKeysOK struct {
	Payload *platformclientmodels.KeyPagingSliceResult
}

func (o *ListKeysOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}/keys][%d] listKeysOK  %+v", 200, o.ToString())
}

func (o *ListKeysOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ListKeysOK) GetPayload() *platformclientmodels.KeyPagingSliceResult {
	return o.Payload
}

func (o *ListKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.KeyPagingSliceResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

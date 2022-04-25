// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// QueryUserIAPOrdersReader is a Reader for the QueryUserIAPOrders structure.
type QueryUserIAPOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryUserIAPOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryUserIAPOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/users/{userId}/iap returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueryUserIAPOrdersOK creates a QueryUserIAPOrdersOK with default headers values
func NewQueryUserIAPOrdersOK() *QueryUserIAPOrdersOK {
	return &QueryUserIAPOrdersOK{}
}

/*QueryUserIAPOrdersOK handles this case with default header values.

  successful operation
*/
type QueryUserIAPOrdersOK struct {
	Payload *platformclientmodels.IAPOrderPagingSlicedResult
}

func (o *QueryUserIAPOrdersOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/users/{userId}/iap][%d] queryUserIAPOrdersOK  %+v", 200, o.Payload)
}

func (o *QueryUserIAPOrdersOK) GetPayload() *platformclientmodels.IAPOrderPagingSlicedResult {
	return o.Payload
}

func (o *QueryUserIAPOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.IAPOrderPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

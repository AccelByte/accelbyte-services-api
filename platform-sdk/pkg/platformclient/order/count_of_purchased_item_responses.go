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

// CountOfPurchasedItemReader is a Reader for the CountOfPurchasedItem structure.
type CountOfPurchasedItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountOfPurchasedItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountOfPurchasedItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/users/{userId}/orders/countOfItem returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCountOfPurchasedItemOK creates a CountOfPurchasedItemOK with default headers values
func NewCountOfPurchasedItemOK() *CountOfPurchasedItemOK {
	return &CountOfPurchasedItemOK{}
}

/*CountOfPurchasedItemOK handles this case with default header values.

  successful operation
*/
type CountOfPurchasedItemOK struct {
	Payload *platformclientmodels.PurchasedItemCount
}

func (o *CountOfPurchasedItemOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/users/{userId}/orders/countOfItem][%d] countOfPurchasedItemOK  %+v", 200, o.ToString())
}

func (o *CountOfPurchasedItemOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *CountOfPurchasedItemOK) GetPayload() *platformclientmodels.PurchasedItemCount {
	return o.Payload
}

func (o *CountOfPurchasedItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.PurchasedItemCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

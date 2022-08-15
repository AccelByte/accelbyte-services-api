// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// PublicQueryUserStatItemsReader is a Reader for the PublicQueryUserStatItems structure.
type PublicQueryUserStatItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicQueryUserStatItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicQueryUserStatItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /social/v1/public/namespaces/{namespace}/users/{userId}/statitems returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicQueryUserStatItemsOK creates a PublicQueryUserStatItemsOK with default headers values
func NewPublicQueryUserStatItemsOK() *PublicQueryUserStatItemsOK {
	return &PublicQueryUserStatItemsOK{}
}

/*PublicQueryUserStatItemsOK handles this case with default header values.

  successful operation
*/
type PublicQueryUserStatItemsOK struct {
	Payload *socialclientmodels.UserStatItemPagingSlicedResult
}

func (o *PublicQueryUserStatItemsOK) Error() string {
	return fmt.Sprintf("[GET /social/v1/public/namespaces/{namespace}/users/{userId}/statitems][%d] publicQueryUserStatItemsOK  %+v", 200, o.ToString())
}

func (o *PublicQueryUserStatItemsOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicQueryUserStatItemsOK) GetPayload() *socialclientmodels.UserStatItemPagingSlicedResult {
	return o.Payload
}

func (o *PublicQueryUserStatItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.UserStatItemPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

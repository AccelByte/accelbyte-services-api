// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package category

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

// PublicGetChildCategoriesReader is a Reader for the PublicGetChildCategories structure.
type PublicGetChildCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetChildCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetChildCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/public/namespaces/{namespace}/categories/{categoryPath}/children returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetChildCategoriesOK creates a PublicGetChildCategoriesOK with default headers values
func NewPublicGetChildCategoriesOK() *PublicGetChildCategoriesOK {
	return &PublicGetChildCategoriesOK{}
}

/*PublicGetChildCategoriesOK handles this case with default header values.

  successful operation
*/
type PublicGetChildCategoriesOK struct {
	Payload []*platformclientmodels.CategoryInfo
}

func (o *PublicGetChildCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/categories/{categoryPath}/children][%d] publicGetChildCategoriesOK  %+v", 200, o.ToString())
}

func (o *PublicGetChildCategoriesOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicGetChildCategoriesOK) GetPayload() []*platformclientmodels.CategoryInfo {
	return o.Payload
}

func (o *PublicGetChildCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

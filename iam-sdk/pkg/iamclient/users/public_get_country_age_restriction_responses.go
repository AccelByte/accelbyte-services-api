// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// PublicGetCountryAgeRestrictionReader is a Reader for the PublicGetCountryAgeRestriction structure.
type PublicGetCountryAgeRestrictionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetCountryAgeRestrictionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetCountryAgeRestrictionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicGetCountryAgeRestrictionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicGetCountryAgeRestrictionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v2/public/namespaces/{namespace}/countries/{countryCode}/agerestrictions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetCountryAgeRestrictionOK creates a PublicGetCountryAgeRestrictionOK with default headers values
func NewPublicGetCountryAgeRestrictionOK() *PublicGetCountryAgeRestrictionOK {
	return &PublicGetCountryAgeRestrictionOK{}
}

/*PublicGetCountryAgeRestrictionOK handles this case with default header values.

  OK
*/
type PublicGetCountryAgeRestrictionOK struct {
	Payload []*iamclientmodels.AccountcommonCountryAgeRestriction
}

func (o *PublicGetCountryAgeRestrictionOK) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/countries/{countryCode}/agerestrictions][%d] publicGetCountryAgeRestrictionOK  %+v", 200, o.ToJSONString())
}

func (o *PublicGetCountryAgeRestrictionOK) ToJSONString() string {
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

func (o *PublicGetCountryAgeRestrictionOK) GetPayload() []*iamclientmodels.AccountcommonCountryAgeRestriction {
	return o.Payload
}

func (o *PublicGetCountryAgeRestrictionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetCountryAgeRestrictionUnauthorized creates a PublicGetCountryAgeRestrictionUnauthorized with default headers values
func NewPublicGetCountryAgeRestrictionUnauthorized() *PublicGetCountryAgeRestrictionUnauthorized {
	return &PublicGetCountryAgeRestrictionUnauthorized{}
}

/*PublicGetCountryAgeRestrictionUnauthorized handles this case with default header values.

  Unauthorized access
*/
type PublicGetCountryAgeRestrictionUnauthorized struct {
}

func (o *PublicGetCountryAgeRestrictionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/countries/{countryCode}/agerestrictions][%d] publicGetCountryAgeRestrictionUnauthorized ", 401)
}

func (o *PublicGetCountryAgeRestrictionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicGetCountryAgeRestrictionNotFound creates a PublicGetCountryAgeRestrictionNotFound with default headers values
func NewPublicGetCountryAgeRestrictionNotFound() *PublicGetCountryAgeRestrictionNotFound {
	return &PublicGetCountryAgeRestrictionNotFound{}
}

/*PublicGetCountryAgeRestrictionNotFound handles this case with default header values.

  Data not found
*/
type PublicGetCountryAgeRestrictionNotFound struct {
}

func (o *PublicGetCountryAgeRestrictionNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/countries/{countryCode}/agerestrictions][%d] publicGetCountryAgeRestrictionNotFound ", 404)
}

func (o *PublicGetCountryAgeRestrictionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

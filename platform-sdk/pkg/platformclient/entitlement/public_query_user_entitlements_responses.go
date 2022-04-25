// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// PublicQueryUserEntitlementsReader is a Reader for the PublicQueryUserEntitlements structure.
type PublicQueryUserEntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicQueryUserEntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicQueryUserEntitlementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/public/namespaces/{namespace}/users/{userId}/entitlements returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicQueryUserEntitlementsOK creates a PublicQueryUserEntitlementsOK with default headers values
func NewPublicQueryUserEntitlementsOK() *PublicQueryUserEntitlementsOK {
	return &PublicQueryUserEntitlementsOK{}
}

/*PublicQueryUserEntitlementsOK handles this case with default header values.

  successful operation
*/
type PublicQueryUserEntitlementsOK struct {
	Payload *platformclientmodels.EntitlementPagingSlicedResult
}

func (o *PublicQueryUserEntitlementsOK) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/users/{userId}/entitlements][%d] publicQueryUserEntitlementsOK  %+v", 200, o.Payload)
}

func (o *PublicQueryUserEntitlementsOK) GetPayload() *platformclientmodels.EntitlementPagingSlicedResult {
	return o.Payload
}

func (o *PublicQueryUserEntitlementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.EntitlementPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

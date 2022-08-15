// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// GetUserAppEntitlementOwnershipByAppIDReader is a Reader for the GetUserAppEntitlementOwnershipByAppID structure.
type GetUserAppEntitlementOwnershipByAppIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAppEntitlementOwnershipByAppIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserAppEntitlementOwnershipByAppIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/byAppId returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserAppEntitlementOwnershipByAppIDOK creates a GetUserAppEntitlementOwnershipByAppIDOK with default headers values
func NewGetUserAppEntitlementOwnershipByAppIDOK() *GetUserAppEntitlementOwnershipByAppIDOK {
	return &GetUserAppEntitlementOwnershipByAppIDOK{}
}

/*GetUserAppEntitlementOwnershipByAppIDOK handles this case with default header values.

  successful operation
*/
type GetUserAppEntitlementOwnershipByAppIDOK struct {
	Payload *platformclientmodels.Ownership
}

func (o *GetUserAppEntitlementOwnershipByAppIDOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/byAppId][%d] getUserAppEntitlementOwnershipByAppIdOK  %+v", 200, o.ToJSONString())
}

func (o *GetUserAppEntitlementOwnershipByAppIDOK) ToJSONString() string {
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

func (o *GetUserAppEntitlementOwnershipByAppIDOK) GetPayload() *platformclientmodels.Ownership {
	return o.Payload
}

func (o *GetUserAppEntitlementOwnershipByAppIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.Ownership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

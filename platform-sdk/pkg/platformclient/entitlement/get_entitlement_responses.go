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

// GetEntitlementReader is a Reader for the GetEntitlement structure.
type GetEntitlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntitlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntitlementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEntitlementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/entitlements/{entitlementId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetEntitlementOK creates a GetEntitlementOK with default headers values
func NewGetEntitlementOK() *GetEntitlementOK {
	return &GetEntitlementOK{}
}

/*GetEntitlementOK handles this case with default header values.

  successful operation
*/
type GetEntitlementOK struct {
	Payload *platformclientmodels.EntitlementInfo
}

func (o *GetEntitlementOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/entitlements/{entitlementId}][%d] getEntitlementOK  %+v", 200, o.ToString())
}

func (o *GetEntitlementOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetEntitlementOK) GetPayload() *platformclientmodels.EntitlementInfo {
	return o.Payload
}

func (o *GetEntitlementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.EntitlementInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntitlementNotFound creates a GetEntitlementNotFound with default headers values
func NewGetEntitlementNotFound() *GetEntitlementNotFound {
	return &GetEntitlementNotFound{}
}

/*GetEntitlementNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>31141</td><td>Entitlement [{entitlementId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type GetEntitlementNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GetEntitlementNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/entitlements/{entitlementId}][%d] getEntitlementNotFound  %+v", 404, o.ToString())
}

func (o *GetEntitlementNotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetEntitlementNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetEntitlementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

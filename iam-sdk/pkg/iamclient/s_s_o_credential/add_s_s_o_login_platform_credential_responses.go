// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_credential

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

// AddSSOLoginPlatformCredentialReader is a Reader for the AddSSOLoginPlatformCredential structure.
type AddSSOLoginPlatformCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSSOLoginPlatformCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddSSOLoginPlatformCredentialCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSSOLoginPlatformCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddSSOLoginPlatformCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAddSSOLoginPlatformCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAddSSOLoginPlatformCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAddSSOLoginPlatformCredentialCreated creates a AddSSOLoginPlatformCredentialCreated with default headers values
func NewAddSSOLoginPlatformCredentialCreated() *AddSSOLoginPlatformCredentialCreated {
	return &AddSSOLoginPlatformCredentialCreated{}
}

/*AddSSOLoginPlatformCredentialCreated handles this case with default header values.

  SSO Credential Created
*/
type AddSSOLoginPlatformCredentialCreated struct {
	Payload *iamclientmodels.ModelSSOPlatformCredentialResponse
}

func (o *AddSSOLoginPlatformCredentialCreated) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] addSSOLoginPlatformCredentialCreated  %+v", 201, o.ToJSONString())
}

func (o *AddSSOLoginPlatformCredentialCreated) ToJSONString() string {
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

func (o *AddSSOLoginPlatformCredentialCreated) GetPayload() *iamclientmodels.ModelSSOPlatformCredentialResponse {
	return o.Payload
}

func (o *AddSSOLoginPlatformCredentialCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelSSOPlatformCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSSOLoginPlatformCredentialBadRequest creates a AddSSOLoginPlatformCredentialBadRequest with default headers values
func NewAddSSOLoginPlatformCredentialBadRequest() *AddSSOLoginPlatformCredentialBadRequest {
	return &AddSSOLoginPlatformCredentialBadRequest{}
}

/*AddSSOLoginPlatformCredentialBadRequest handles this case with default header values.

  Bad Request
*/
type AddSSOLoginPlatformCredentialBadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AddSSOLoginPlatformCredentialBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] addSSOLoginPlatformCredentialBadRequest  %+v", 400, o.ToJSONString())
}

func (o *AddSSOLoginPlatformCredentialBadRequest) ToJSONString() string {
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

func (o *AddSSOLoginPlatformCredentialBadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AddSSOLoginPlatformCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSSOLoginPlatformCredentialUnauthorized creates a AddSSOLoginPlatformCredentialUnauthorized with default headers values
func NewAddSSOLoginPlatformCredentialUnauthorized() *AddSSOLoginPlatformCredentialUnauthorized {
	return &AddSSOLoginPlatformCredentialUnauthorized{}
}

/*AddSSOLoginPlatformCredentialUnauthorized handles this case with default header values.

  Unauthorized
*/
type AddSSOLoginPlatformCredentialUnauthorized struct {
}

func (o *AddSSOLoginPlatformCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] addSSOLoginPlatformCredentialUnauthorized ", 401)
}

func (o *AddSSOLoginPlatformCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSSOLoginPlatformCredentialForbidden creates a AddSSOLoginPlatformCredentialForbidden with default headers values
func NewAddSSOLoginPlatformCredentialForbidden() *AddSSOLoginPlatformCredentialForbidden {
	return &AddSSOLoginPlatformCredentialForbidden{}
}

/*AddSSOLoginPlatformCredentialForbidden handles this case with default header values.

  Forbidden
*/
type AddSSOLoginPlatformCredentialForbidden struct {
}

func (o *AddSSOLoginPlatformCredentialForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] addSSOLoginPlatformCredentialForbidden ", 403)
}

func (o *AddSSOLoginPlatformCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSSOLoginPlatformCredentialInternalServerError creates a AddSSOLoginPlatformCredentialInternalServerError with default headers values
func NewAddSSOLoginPlatformCredentialInternalServerError() *AddSSOLoginPlatformCredentialInternalServerError {
	return &AddSSOLoginPlatformCredentialInternalServerError{}
}

/*AddSSOLoginPlatformCredentialInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AddSSOLoginPlatformCredentialInternalServerError struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AddSSOLoginPlatformCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] addSSOLoginPlatformCredentialInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *AddSSOLoginPlatformCredentialInternalServerError) ToJSONString() string {
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

func (o *AddSSOLoginPlatformCredentialInternalServerError) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AddSSOLoginPlatformCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// AdminGetUserMappingReader is a Reader for the AdminGetUserMapping structure.
type AdminGetUserMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetUserMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetUserMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminGetUserMappingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGetUserMappingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminGetUserMappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminGetUserMappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/admin/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminGetUserMappingOK creates a AdminGetUserMappingOK with default headers values
func NewAdminGetUserMappingOK() *AdminGetUserMappingOK {
	return &AdminGetUserMappingOK{}
}

/*AdminGetUserMappingOK handles this case with default header values.

  OK
*/
type AdminGetUserMappingOK struct {
	Payload *iamclientmodels.ModelGetUserMappingV3
}

func (o *AdminGetUserMappingOK) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] adminGetUserMappingOK  %+v", 200, o.ToJSONString())
}

func (o *AdminGetUserMappingOK) ToJSONString() string {
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

func (o *AdminGetUserMappingOK) GetPayload() *iamclientmodels.ModelGetUserMappingV3 {
	return o.Payload
}

func (o *AdminGetUserMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelGetUserMappingV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetUserMappingBadRequest creates a AdminGetUserMappingBadRequest with default headers values
func NewAdminGetUserMappingBadRequest() *AdminGetUserMappingBadRequest {
	return &AdminGetUserMappingBadRequest{}
}

/*AdminGetUserMappingBadRequest handles this case with default header values.

  Invalid request
*/
type AdminGetUserMappingBadRequest struct {
}

func (o *AdminGetUserMappingBadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] adminGetUserMappingBadRequest ", 400)
}

func (o *AdminGetUserMappingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminGetUserMappingUnauthorized creates a AdminGetUserMappingUnauthorized with default headers values
func NewAdminGetUserMappingUnauthorized() *AdminGetUserMappingUnauthorized {
	return &AdminGetUserMappingUnauthorized{}
}

/*AdminGetUserMappingUnauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminGetUserMappingUnauthorized struct {
}

func (o *AdminGetUserMappingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] adminGetUserMappingUnauthorized ", 401)
}

func (o *AdminGetUserMappingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminGetUserMappingForbidden creates a AdminGetUserMappingForbidden with default headers values
func NewAdminGetUserMappingForbidden() *AdminGetUserMappingForbidden {
	return &AdminGetUserMappingForbidden{}
}

/*AdminGetUserMappingForbidden handles this case with default header values.

  Forbidden
*/
type AdminGetUserMappingForbidden struct {
}

func (o *AdminGetUserMappingForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] adminGetUserMappingForbidden ", 403)
}

func (o *AdminGetUserMappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminGetUserMappingNotFound creates a AdminGetUserMappingNotFound with default headers values
func NewAdminGetUserMappingNotFound() *AdminGetUserMappingNotFound {
	return &AdminGetUserMappingNotFound{}
}

/*AdminGetUserMappingNotFound handles this case with default header values.

  Data not found
*/
type AdminGetUserMappingNotFound struct {
}

func (o *AdminGetUserMappingNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v3/admin/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] adminGetUserMappingNotFound ", 404)
}

func (o *AdminGetUserMappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

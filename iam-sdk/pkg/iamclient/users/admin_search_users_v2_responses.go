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

// AdminSearchUsersV2Reader is a Reader for the AdminSearchUsersV2 structure.
type AdminSearchUsersV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminSearchUsersV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminSearchUsersV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminSearchUsersV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminSearchUsersV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminSearchUsersV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v2/admin/namespaces/{namespace}/users returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminSearchUsersV2OK creates a AdminSearchUsersV2OK with default headers values
func NewAdminSearchUsersV2OK() *AdminSearchUsersV2OK {
	return &AdminSearchUsersV2OK{}
}

/*AdminSearchUsersV2OK handles this case with default header values.

  OK
*/
type AdminSearchUsersV2OK struct {
	Payload *iamclientmodels.ModelSearchUsersByPlatformIDResponse
}

func (o *AdminSearchUsersV2OK) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/users][%d] adminSearchUsersV2OK  %+v", 200, o.ToString())
}

func (o *AdminSearchUsersV2OK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *AdminSearchUsersV2OK) GetPayload() *iamclientmodels.ModelSearchUsersByPlatformIDResponse {
	return o.Payload
}

func (o *AdminSearchUsersV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelSearchUsersByPlatformIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSearchUsersV2BadRequest creates a AdminSearchUsersV2BadRequest with default headers values
func NewAdminSearchUsersV2BadRequest() *AdminSearchUsersV2BadRequest {
	return &AdminSearchUsersV2BadRequest{}
}

/*AdminSearchUsersV2BadRequest handles this case with default header values.

  Invalid request
*/
type AdminSearchUsersV2BadRequest struct {
}

func (o *AdminSearchUsersV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/users][%d] adminSearchUsersV2BadRequest ", 400)
}

func (o *AdminSearchUsersV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminSearchUsersV2Unauthorized creates a AdminSearchUsersV2Unauthorized with default headers values
func NewAdminSearchUsersV2Unauthorized() *AdminSearchUsersV2Unauthorized {
	return &AdminSearchUsersV2Unauthorized{}
}

/*AdminSearchUsersV2Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminSearchUsersV2Unauthorized struct {
}

func (o *AdminSearchUsersV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/users][%d] adminSearchUsersV2Unauthorized ", 401)
}

func (o *AdminSearchUsersV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminSearchUsersV2Forbidden creates a AdminSearchUsersV2Forbidden with default headers values
func NewAdminSearchUsersV2Forbidden() *AdminSearchUsersV2Forbidden {
	return &AdminSearchUsersV2Forbidden{}
}

/*AdminSearchUsersV2Forbidden handles this case with default header values.

  Forbidden
*/
type AdminSearchUsersV2Forbidden struct {
}

func (o *AdminSearchUsersV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/users][%d] adminSearchUsersV2Forbidden ", 403)
}

func (o *AdminSearchUsersV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

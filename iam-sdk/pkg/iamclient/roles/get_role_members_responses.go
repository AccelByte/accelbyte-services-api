// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// GetRoleMembersReader is a Reader for the GetRoleMembers structure.
type GetRoleMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoleMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRoleMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRoleMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoleMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/roles/{roleId}/members returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetRoleMembersOK creates a GetRoleMembersOK with default headers values
func NewGetRoleMembersOK() *GetRoleMembersOK {
	return &GetRoleMembersOK{}
}

/*GetRoleMembersOK handles this case with default header values.

  OK
*/
type GetRoleMembersOK struct {
	Payload *iamclientmodels.ModelRoleMembersResponse
}

func (o *GetRoleMembersOK) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/members][%d] getRoleMembersOK  %+v", 200, o.ToString())
}

func (o *GetRoleMembersOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetRoleMembersOK) GetPayload() *iamclientmodels.ModelRoleMembersResponse {
	return o.Payload
}

func (o *GetRoleMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelRoleMembersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleMembersBadRequest creates a GetRoleMembersBadRequest with default headers values
func NewGetRoleMembersBadRequest() *GetRoleMembersBadRequest {
	return &GetRoleMembersBadRequest{}
}

/*GetRoleMembersBadRequest handles this case with default header values.

  Invalid request
*/
type GetRoleMembersBadRequest struct {
}

func (o *GetRoleMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/members][%d] getRoleMembersBadRequest ", 400)
}

func (o *GetRoleMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleMembersUnauthorized creates a GetRoleMembersUnauthorized with default headers values
func NewGetRoleMembersUnauthorized() *GetRoleMembersUnauthorized {
	return &GetRoleMembersUnauthorized{}
}

/*GetRoleMembersUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetRoleMembersUnauthorized struct {
}

func (o *GetRoleMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/members][%d] getRoleMembersUnauthorized ", 401)
}

func (o *GetRoleMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleMembersForbidden creates a GetRoleMembersForbidden with default headers values
func NewGetRoleMembersForbidden() *GetRoleMembersForbidden {
	return &GetRoleMembersForbidden{}
}

/*GetRoleMembersForbidden handles this case with default header values.

  Forbidden
*/
type GetRoleMembersForbidden struct {
}

func (o *GetRoleMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/members][%d] getRoleMembersForbidden ", 403)
}

func (o *GetRoleMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleMembersNotFound creates a GetRoleMembersNotFound with default headers values
func NewGetRoleMembersNotFound() *GetRoleMembersNotFound {
	return &GetRoleMembersNotFound{}
}

/*GetRoleMembersNotFound handles this case with default header values.

  Data not found
*/
type GetRoleMembersNotFound struct {
}

func (o *GetRoleMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/members][%d] getRoleMembersNotFound ", 404)
}

func (o *GetRoleMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

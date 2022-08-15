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

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/roles/{roleId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/*GetRoleOK handles this case with default header values.

  OK
*/
type GetRoleOK struct {
	Payload *iamclientmodels.ModelRoleResponse
}

func (o *GetRoleOK) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}][%d] getRoleOK  %+v", 200, o.ToString())
}

func (o *GetRoleOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetRoleOK) GetPayload() *iamclientmodels.ModelRoleResponse {
	return o.Payload
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleUnauthorized creates a GetRoleUnauthorized with default headers values
func NewGetRoleUnauthorized() *GetRoleUnauthorized {
	return &GetRoleUnauthorized{}
}

/*GetRoleUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetRoleUnauthorized struct {
}

func (o *GetRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}][%d] getRoleUnauthorized ", 401)
}

func (o *GetRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleForbidden creates a GetRoleForbidden with default headers values
func NewGetRoleForbidden() *GetRoleForbidden {
	return &GetRoleForbidden{}
}

/*GetRoleForbidden handles this case with default header values.

  Forbidden
*/
type GetRoleForbidden struct {
}

func (o *GetRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}][%d] getRoleForbidden ", 403)
}

func (o *GetRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleNotFound creates a GetRoleNotFound with default headers values
func NewGetRoleNotFound() *GetRoleNotFound {
	return &GetRoleNotFound{}
}

/*GetRoleNotFound handles this case with default header values.

  Data not found
*/
type GetRoleNotFound struct {
}

func (o *GetRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}][%d] getRoleNotFound ", 404)
}

func (o *GetRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

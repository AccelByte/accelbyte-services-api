// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/group-sdk/pkg/groupclientmodels"
)

// GetSingleGroupPublicV1Reader is a Reader for the GetSingleGroupPublicV1 structure.
type GetSingleGroupPublicV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleGroupPublicV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSingleGroupPublicV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSingleGroupPublicV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSingleGroupPublicV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSingleGroupPublicV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSingleGroupPublicV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSingleGroupPublicV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /group/v1/public/namespaces/{namespace}/groups/{groupId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetSingleGroupPublicV1OK creates a GetSingleGroupPublicV1OK with default headers values
func NewGetSingleGroupPublicV1OK() *GetSingleGroupPublicV1OK {
	return &GetSingleGroupPublicV1OK{}
}

/*GetSingleGroupPublicV1OK handles this case with default header values.

  OK
*/
type GetSingleGroupPublicV1OK struct {
	Payload *groupclientmodels.ModelsGroupResponseV1
}

func (o *GetSingleGroupPublicV1OK) Error() string {
	return fmt.Sprintf("[GET /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] getSingleGroupPublicV1OK  %+v", 200, o.ToString())
}

func (o *GetSingleGroupPublicV1OK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetSingleGroupPublicV1OK) GetPayload() *groupclientmodels.ModelsGroupResponseV1 {
	return o.Payload
}

func (o *GetSingleGroupPublicV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ModelsGroupResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleGroupPublicV1BadRequest creates a GetSingleGroupPublicV1BadRequest with default headers values
func NewGetSingleGroupPublicV1BadRequest() *GetSingleGroupPublicV1BadRequest {
	return &GetSingleGroupPublicV1BadRequest{}
}

/*GetSingleGroupPublicV1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type GetSingleGroupPublicV1BadRequest struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *GetSingleGroupPublicV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] getSingleGroupPublicV1BadRequest  %+v", 400, o.ToString())
}

func (o *GetSingleGroupPublicV1BadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetSingleGroupPublicV1BadRequest) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetSingleGroupPublicV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleGroupPublicV1Unauthorized creates a GetSingleGroupPublicV1Unauthorized with default headers values
func NewGetSingleGroupPublicV1Unauthorized() *GetSingleGroupPublicV1Unauthorized {
	return &GetSingleGroupPublicV1Unauthorized{}
}

/*GetSingleGroupPublicV1Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type GetSingleGroupPublicV1Unauthorized struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *GetSingleGroupPublicV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] getSingleGroupPublicV1Unauthorized  %+v", 401, o.ToString())
}

func (o *GetSingleGroupPublicV1Unauthorized) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetSingleGroupPublicV1Unauthorized) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetSingleGroupPublicV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleGroupPublicV1Forbidden creates a GetSingleGroupPublicV1Forbidden with default headers values
func NewGetSingleGroupPublicV1Forbidden() *GetSingleGroupPublicV1Forbidden {
	return &GetSingleGroupPublicV1Forbidden{}
}

/*GetSingleGroupPublicV1Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr><tr><td>20022</td><td>token is not user token</td></tr></table>
*/
type GetSingleGroupPublicV1Forbidden struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *GetSingleGroupPublicV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] getSingleGroupPublicV1Forbidden  %+v", 403, o.ToString())
}

func (o *GetSingleGroupPublicV1Forbidden) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetSingleGroupPublicV1Forbidden) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetSingleGroupPublicV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleGroupPublicV1NotFound creates a GetSingleGroupPublicV1NotFound with default headers values
func NewGetSingleGroupPublicV1NotFound() *GetSingleGroupPublicV1NotFound {
	return &GetSingleGroupPublicV1NotFound{}
}

/*GetSingleGroupPublicV1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>73333</td><td>group not found</td></tr></table>
*/
type GetSingleGroupPublicV1NotFound struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *GetSingleGroupPublicV1NotFound) Error() string {
	return fmt.Sprintf("[GET /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] getSingleGroupPublicV1NotFound  %+v", 404, o.ToString())
}

func (o *GetSingleGroupPublicV1NotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetSingleGroupPublicV1NotFound) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetSingleGroupPublicV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleGroupPublicV1InternalServerError creates a GetSingleGroupPublicV1InternalServerError with default headers values
func NewGetSingleGroupPublicV1InternalServerError() *GetSingleGroupPublicV1InternalServerError {
	return &GetSingleGroupPublicV1InternalServerError{}
}

/*GetSingleGroupPublicV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetSingleGroupPublicV1InternalServerError struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *GetSingleGroupPublicV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] getSingleGroupPublicV1InternalServerError  %+v", 500, o.ToString())
}

func (o *GetSingleGroupPublicV1InternalServerError) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetSingleGroupPublicV1InternalServerError) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetSingleGroupPublicV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

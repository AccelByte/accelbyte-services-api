// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/group-sdk/pkg/groupclientmodels"
)

// DeleteGroupConfigurationV1Reader is a Reader for the DeleteGroupConfigurationV1 structure.
type DeleteGroupConfigurationV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupConfigurationV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGroupConfigurationV1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGroupConfigurationV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteGroupConfigurationV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteGroupConfigurationV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteGroupConfigurationV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteGroupConfigurationV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteGroupConfigurationV1NoContent creates a DeleteGroupConfigurationV1NoContent with default headers values
func NewDeleteGroupConfigurationV1NoContent() *DeleteGroupConfigurationV1NoContent {
	return &DeleteGroupConfigurationV1NoContent{}
}

/*DeleteGroupConfigurationV1NoContent handles this case with default header values.

  No Content
*/
type DeleteGroupConfigurationV1NoContent struct {
}

func (o *DeleteGroupConfigurationV1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode}][%d] deleteGroupConfigurationV1NoContent ", 204)
}

func (o *DeleteGroupConfigurationV1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupConfigurationV1BadRequest creates a DeleteGroupConfigurationV1BadRequest with default headers values
func NewDeleteGroupConfigurationV1BadRequest() *DeleteGroupConfigurationV1BadRequest {
	return &DeleteGroupConfigurationV1BadRequest{}
}

/*DeleteGroupConfigurationV1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type DeleteGroupConfigurationV1BadRequest struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *DeleteGroupConfigurationV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode}][%d] deleteGroupConfigurationV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteGroupConfigurationV1BadRequest) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteGroupConfigurationV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupConfigurationV1Unauthorized creates a DeleteGroupConfigurationV1Unauthorized with default headers values
func NewDeleteGroupConfigurationV1Unauthorized() *DeleteGroupConfigurationV1Unauthorized {
	return &DeleteGroupConfigurationV1Unauthorized{}
}

/*DeleteGroupConfigurationV1Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type DeleteGroupConfigurationV1Unauthorized struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *DeleteGroupConfigurationV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode}][%d] deleteGroupConfigurationV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGroupConfigurationV1Unauthorized) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteGroupConfigurationV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupConfigurationV1Forbidden creates a DeleteGroupConfigurationV1Forbidden with default headers values
func NewDeleteGroupConfigurationV1Forbidden() *DeleteGroupConfigurationV1Forbidden {
	return &DeleteGroupConfigurationV1Forbidden{}
}

/*DeleteGroupConfigurationV1Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr><tr><td>20022</td><td>token is not user token</td></tr></table>
*/
type DeleteGroupConfigurationV1Forbidden struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *DeleteGroupConfigurationV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode}][%d] deleteGroupConfigurationV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteGroupConfigurationV1Forbidden) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteGroupConfigurationV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupConfigurationV1NotFound creates a DeleteGroupConfigurationV1NotFound with default headers values
func NewDeleteGroupConfigurationV1NotFound() *DeleteGroupConfigurationV1NotFound {
	return &DeleteGroupConfigurationV1NotFound{}
}

/*DeleteGroupConfigurationV1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>73131</td><td>global configuration not found</td></tr></table>
*/
type DeleteGroupConfigurationV1NotFound struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *DeleteGroupConfigurationV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode}][%d] deleteGroupConfigurationV1NotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupConfigurationV1NotFound) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteGroupConfigurationV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupConfigurationV1InternalServerError creates a DeleteGroupConfigurationV1InternalServerError with default headers values
func NewDeleteGroupConfigurationV1InternalServerError() *DeleteGroupConfigurationV1InternalServerError {
	return &DeleteGroupConfigurationV1InternalServerError{}
}

/*DeleteGroupConfigurationV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteGroupConfigurationV1InternalServerError struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *DeleteGroupConfigurationV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /group/v1/admin/namespaces/{namespace}/configuration/{configurationCode}][%d] deleteGroupConfigurationV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGroupConfigurationV1InternalServerError) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteGroupConfigurationV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

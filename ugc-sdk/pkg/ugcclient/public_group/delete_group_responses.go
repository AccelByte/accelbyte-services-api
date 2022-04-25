// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package public_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// DeleteGroupReader is a Reader for the DeleteGroup structure.
type DeleteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/groups/{groupId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteGroupNoContent creates a DeleteGroupNoContent with default headers values
func NewDeleteGroupNoContent() *DeleteGroupNoContent {
	return &DeleteGroupNoContent{}
}

/*DeleteGroupNoContent handles this case with default header values.

  No Content
*/
type DeleteGroupNoContent struct {
}

func (o *DeleteGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/groups/{groupId}][%d] deleteGroupNoContent ", 204)
}

func (o *DeleteGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupUnauthorized creates a DeleteGroupUnauthorized with default headers values
func NewDeleteGroupUnauthorized() *DeleteGroupUnauthorized {
	return &DeleteGroupUnauthorized{}
}

/*DeleteGroupUnauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteGroupUnauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/groups/{groupId}][%d] deleteGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteGroupUnauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupNotFound creates a DeleteGroupNotFound with default headers values
func NewDeleteGroupNotFound() *DeleteGroupNotFound {
	return &DeleteGroupNotFound{}
}

/*DeleteGroupNotFound handles this case with default header values.

  Not Found
*/
type DeleteGroupNotFound struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/groups/{groupId}][%d] deleteGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteGroupNotFound) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupInternalServerError creates a DeleteGroupInternalServerError with default headers values
func NewDeleteGroupInternalServerError() *DeleteGroupInternalServerError {
	return &DeleteGroupInternalServerError{}
}

/*DeleteGroupInternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteGroupInternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/groups/{groupId}][%d] deleteGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteGroupInternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

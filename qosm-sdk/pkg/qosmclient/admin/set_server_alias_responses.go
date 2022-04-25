// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/qosm-sdk/pkg/qosmclientmodels"
)

// SetServerAliasReader is a Reader for the SetServerAlias structure.
type SetServerAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetServerAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetServerAliasNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetServerAliasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetServerAliasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSetServerAliasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /qosm/admin/servers/{region}/alias returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSetServerAliasNoContent creates a SetServerAliasNoContent with default headers values
func NewSetServerAliasNoContent() *SetServerAliasNoContent {
	return &SetServerAliasNoContent{}
}

/*SetServerAliasNoContent handles this case with default header values.

  record updated
*/
type SetServerAliasNoContent struct {
}

func (o *SetServerAliasNoContent) Error() string {
	return fmt.Sprintf("[POST /qosm/admin/servers/{region}/alias][%d] setServerAliasNoContent ", 204)
}

func (o *SetServerAliasNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetServerAliasBadRequest creates a SetServerAliasBadRequest with default headers values
func NewSetServerAliasBadRequest() *SetServerAliasBadRequest {
	return &SetServerAliasBadRequest{}
}

/*SetServerAliasBadRequest handles this case with default header values.

  malformed request
*/
type SetServerAliasBadRequest struct {
	Payload *qosmclientmodels.ResponseError
}

func (o *SetServerAliasBadRequest) Error() string {
	return fmt.Sprintf("[POST /qosm/admin/servers/{region}/alias][%d] setServerAliasBadRequest  %+v", 400, o.Payload)
}

func (o *SetServerAliasBadRequest) GetPayload() *qosmclientmodels.ResponseError {
	return o.Payload
}

func (o *SetServerAliasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(qosmclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetServerAliasNotFound creates a SetServerAliasNotFound with default headers values
func NewSetServerAliasNotFound() *SetServerAliasNotFound {
	return &SetServerAliasNotFound{}
}

/*SetServerAliasNotFound handles this case with default header values.

  server record not found
*/
type SetServerAliasNotFound struct {
	Payload *qosmclientmodels.ResponseError
}

func (o *SetServerAliasNotFound) Error() string {
	return fmt.Sprintf("[POST /qosm/admin/servers/{region}/alias][%d] setServerAliasNotFound  %+v", 404, o.Payload)
}

func (o *SetServerAliasNotFound) GetPayload() *qosmclientmodels.ResponseError {
	return o.Payload
}

func (o *SetServerAliasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(qosmclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetServerAliasInternalServerError creates a SetServerAliasInternalServerError with default headers values
func NewSetServerAliasInternalServerError() *SetServerAliasInternalServerError {
	return &SetServerAliasInternalServerError{}
}

/*SetServerAliasInternalServerError handles this case with default header values.

  Internal Server Error
*/
type SetServerAliasInternalServerError struct {
	Payload *qosmclientmodels.ResponseError
}

func (o *SetServerAliasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /qosm/admin/servers/{region}/alias][%d] setServerAliasInternalServerError  %+v", 500, o.Payload)
}

func (o *SetServerAliasInternalServerError) GetPayload() *qosmclientmodels.ResponseError {
	return o.Payload
}

func (o *SetServerAliasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(qosmclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

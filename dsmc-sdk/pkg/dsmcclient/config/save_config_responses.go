// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// SaveConfigReader is a Reader for the SaveConfig structure.
type SaveConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSaveConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSaveConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSaveConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSaveConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dsmcontroller/admin/configs returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSaveConfigNoContent creates a SaveConfigNoContent with default headers values
func NewSaveConfigNoContent() *SaveConfigNoContent {
	return &SaveConfigNoContent{}
}

/*SaveConfigNoContent handles this case with default header values.

  config added/updated
*/
type SaveConfigNoContent struct {
}

func (o *SaveConfigNoContent) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/configs][%d] saveConfigNoContent ", 204)
}

func (o *SaveConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveConfigBadRequest creates a SaveConfigBadRequest with default headers values
func NewSaveConfigBadRequest() *SaveConfigBadRequest {
	return &SaveConfigBadRequest{}
}

/*SaveConfigBadRequest handles this case with default header values.

  malformed request
*/
type SaveConfigBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *SaveConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/configs][%d] saveConfigBadRequest  %+v", 400, o.Payload)
}

func (o *SaveConfigBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *SaveConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveConfigUnauthorized creates a SaveConfigUnauthorized with default headers values
func NewSaveConfigUnauthorized() *SaveConfigUnauthorized {
	return &SaveConfigUnauthorized{}
}

/*SaveConfigUnauthorized handles this case with default header values.

  Unauthorized
*/
type SaveConfigUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *SaveConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/configs][%d] saveConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *SaveConfigUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *SaveConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveConfigInternalServerError creates a SaveConfigInternalServerError with default headers values
func NewSaveConfigInternalServerError() *SaveConfigInternalServerError {
	return &SaveConfigInternalServerError{}
}

/*SaveConfigInternalServerError handles this case with default header values.

  Internal Server Error
*/
type SaveConfigInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *SaveConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/configs][%d] saveConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *SaveConfigInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *SaveConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package image_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// CreateImagePatchReader is a Reader for the CreateImagePatch structure.
type CreateImagePatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImagePatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateImagePatchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateImagePatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateImagePatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateImagePatchConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCreateImagePatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dsmcontroller/admin/images/patches returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCreateImagePatchCreated creates a CreateImagePatchCreated with default headers values
func NewCreateImagePatchCreated() *CreateImagePatchCreated {
	return &CreateImagePatchCreated{}
}

/*CreateImagePatchCreated handles this case with default header values.

  image patch created
*/
type CreateImagePatchCreated struct {
}

func (o *CreateImagePatchCreated) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/images/patches][%d] createImagePatchCreated ", 201)
}

func (o *CreateImagePatchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateImagePatchBadRequest creates a CreateImagePatchBadRequest with default headers values
func NewCreateImagePatchBadRequest() *CreateImagePatchBadRequest {
	return &CreateImagePatchBadRequest{}
}

/*CreateImagePatchBadRequest handles this case with default header values.

  malformed request
*/
type CreateImagePatchBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateImagePatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/images/patches][%d] createImagePatchBadRequest  %+v", 400, o.ToString())
}

func (o *CreateImagePatchBadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *CreateImagePatchBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateImagePatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImagePatchUnauthorized creates a CreateImagePatchUnauthorized with default headers values
func NewCreateImagePatchUnauthorized() *CreateImagePatchUnauthorized {
	return &CreateImagePatchUnauthorized{}
}

/*CreateImagePatchUnauthorized handles this case with default header values.

  Unauthorized
*/
type CreateImagePatchUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateImagePatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/images/patches][%d] createImagePatchUnauthorized  %+v", 401, o.ToString())
}

func (o *CreateImagePatchUnauthorized) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *CreateImagePatchUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateImagePatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImagePatchConflict creates a CreateImagePatchConflict with default headers values
func NewCreateImagePatchConflict() *CreateImagePatchConflict {
	return &CreateImagePatchConflict{}
}

/*CreateImagePatchConflict handles this case with default header values.

  conflict. duplicate image patch version record
*/
type CreateImagePatchConflict struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateImagePatchConflict) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/images/patches][%d] createImagePatchConflict  %+v", 409, o.ToString())
}

func (o *CreateImagePatchConflict) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *CreateImagePatchConflict) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateImagePatchConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImagePatchInternalServerError creates a CreateImagePatchInternalServerError with default headers values
func NewCreateImagePatchInternalServerError() *CreateImagePatchInternalServerError {
	return &CreateImagePatchInternalServerError{}
}

/*CreateImagePatchInternalServerError handles this case with default header values.

  Internal Server Error
*/
type CreateImagePatchInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CreateImagePatchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/admin/images/patches][%d] createImagePatchInternalServerError  %+v", 500, o.ToString())
}

func (o *CreateImagePatchInternalServerError) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *CreateImagePatchInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CreateImagePatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

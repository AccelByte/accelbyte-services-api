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

// GetImagePatchDetailReader is a Reader for the GetImagePatchDetail structure.
type GetImagePatchDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagePatchDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImagePatchDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetImagePatchDetailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetImagePatchDetailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetImagePatchDetailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/admin/namespaces/{namespace}/images/versions/{version}/patches/{versionPatch} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetImagePatchDetailOK creates a GetImagePatchDetailOK with default headers values
func NewGetImagePatchDetailOK() *GetImagePatchDetailOK {
	return &GetImagePatchDetailOK{}
}

/*GetImagePatchDetailOK handles this case with default header values.

  ok
*/
type GetImagePatchDetailOK struct {
	Payload *dsmcclientmodels.ModelsGetImagePatchDetailResponse
}

func (o *GetImagePatchDetailOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/images/versions/{version}/patches/{versionPatch}][%d] getImagePatchDetailOK  %+v", 200, o.ToJSONString())
}

func (o *GetImagePatchDetailOK) ToJSONString() string {
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

func (o *GetImagePatchDetailOK) GetPayload() *dsmcclientmodels.ModelsGetImagePatchDetailResponse {
	return o.Payload
}

func (o *GetImagePatchDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsGetImagePatchDetailResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagePatchDetailUnauthorized creates a GetImagePatchDetailUnauthorized with default headers values
func NewGetImagePatchDetailUnauthorized() *GetImagePatchDetailUnauthorized {
	return &GetImagePatchDetailUnauthorized{}
}

/*GetImagePatchDetailUnauthorized handles this case with default header values.

  Unauthorized
*/
type GetImagePatchDetailUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetImagePatchDetailUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/images/versions/{version}/patches/{versionPatch}][%d] getImagePatchDetailUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *GetImagePatchDetailUnauthorized) ToJSONString() string {
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

func (o *GetImagePatchDetailUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetImagePatchDetailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagePatchDetailNotFound creates a GetImagePatchDetailNotFound with default headers values
func NewGetImagePatchDetailNotFound() *GetImagePatchDetailNotFound {
	return &GetImagePatchDetailNotFound{}
}

/*GetImagePatchDetailNotFound handles this case with default header values.

  image version not found
*/
type GetImagePatchDetailNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetImagePatchDetailNotFound) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/images/versions/{version}/patches/{versionPatch}][%d] getImagePatchDetailNotFound  %+v", 404, o.ToJSONString())
}

func (o *GetImagePatchDetailNotFound) ToJSONString() string {
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

func (o *GetImagePatchDetailNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetImagePatchDetailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagePatchDetailInternalServerError creates a GetImagePatchDetailInternalServerError with default headers values
func NewGetImagePatchDetailInternalServerError() *GetImagePatchDetailInternalServerError {
	return &GetImagePatchDetailInternalServerError{}
}

/*GetImagePatchDetailInternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetImagePatchDetailInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetImagePatchDetailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/images/versions/{version}/patches/{versionPatch}][%d] getImagePatchDetailInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *GetImagePatchDetailInternalServerError) ToJSONString() string {
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

func (o *GetImagePatchDetailInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetImagePatchDetailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

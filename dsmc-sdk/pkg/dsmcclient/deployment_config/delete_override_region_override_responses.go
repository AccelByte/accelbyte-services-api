// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package deployment_config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// DeleteOverrideRegionOverrideReader is a Reader for the DeleteOverrideRegionOverride structure.
type DeleteOverrideRegionOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOverrideRegionOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOverrideRegionOverrideOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOverrideRegionOverrideBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOverrideRegionOverrideUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteOverrideRegionOverrideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteOverrideRegionOverrideInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteOverrideRegionOverrideOK creates a DeleteOverrideRegionOverrideOK with default headers values
func NewDeleteOverrideRegionOverrideOK() *DeleteOverrideRegionOverrideOK {
	return &DeleteOverrideRegionOverrideOK{}
}

/*DeleteOverrideRegionOverrideOK handles this case with default header values.

  deployment region override deleted
*/
type DeleteOverrideRegionOverrideOK struct {
	Payload *dsmcclientmodels.ModelsDeploymentWithOverride
}

func (o *DeleteOverrideRegionOverrideOK) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] deleteOverrideRegionOverrideOK  %+v", 200, o.ToJSONString())
}

func (o *DeleteOverrideRegionOverrideOK) ToJSONString() string {
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

func (o *DeleteOverrideRegionOverrideOK) GetPayload() *dsmcclientmodels.ModelsDeploymentWithOverride {
	return o.Payload
}

func (o *DeleteOverrideRegionOverrideOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(dsmcclientmodels.ModelsDeploymentWithOverride)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOverrideRegionOverrideBadRequest creates a DeleteOverrideRegionOverrideBadRequest with default headers values
func NewDeleteOverrideRegionOverrideBadRequest() *DeleteOverrideRegionOverrideBadRequest {
	return &DeleteOverrideRegionOverrideBadRequest{}
}

/*DeleteOverrideRegionOverrideBadRequest handles this case with default header values.

  malformed request
*/
type DeleteOverrideRegionOverrideBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteOverrideRegionOverrideBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] deleteOverrideRegionOverrideBadRequest  %+v", 400, o.ToJSONString())
}

func (o *DeleteOverrideRegionOverrideBadRequest) ToJSONString() string {
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

func (o *DeleteOverrideRegionOverrideBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteOverrideRegionOverrideBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOverrideRegionOverrideUnauthorized creates a DeleteOverrideRegionOverrideUnauthorized with default headers values
func NewDeleteOverrideRegionOverrideUnauthorized() *DeleteOverrideRegionOverrideUnauthorized {
	return &DeleteOverrideRegionOverrideUnauthorized{}
}

/*DeleteOverrideRegionOverrideUnauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteOverrideRegionOverrideUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteOverrideRegionOverrideUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] deleteOverrideRegionOverrideUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *DeleteOverrideRegionOverrideUnauthorized) ToJSONString() string {
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

func (o *DeleteOverrideRegionOverrideUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteOverrideRegionOverrideUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOverrideRegionOverrideNotFound creates a DeleteOverrideRegionOverrideNotFound with default headers values
func NewDeleteOverrideRegionOverrideNotFound() *DeleteOverrideRegionOverrideNotFound {
	return &DeleteOverrideRegionOverrideNotFound{}
}

/*DeleteOverrideRegionOverrideNotFound handles this case with default header values.

  deployment  not found
*/
type DeleteOverrideRegionOverrideNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteOverrideRegionOverrideNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] deleteOverrideRegionOverrideNotFound  %+v", 404, o.ToJSONString())
}

func (o *DeleteOverrideRegionOverrideNotFound) ToJSONString() string {
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

func (o *DeleteOverrideRegionOverrideNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteOverrideRegionOverrideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOverrideRegionOverrideInternalServerError creates a DeleteOverrideRegionOverrideInternalServerError with default headers values
func NewDeleteOverrideRegionOverrideInternalServerError() *DeleteOverrideRegionOverrideInternalServerError {
	return &DeleteOverrideRegionOverrideInternalServerError{}
}

/*DeleteOverrideRegionOverrideInternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteOverrideRegionOverrideInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteOverrideRegionOverrideInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/versions/{version}/regions/{region}][%d] deleteOverrideRegionOverrideInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *DeleteOverrideRegionOverrideInternalServerError) ToJSONString() string {
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

func (o *DeleteOverrideRegionOverrideInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteOverrideRegionOverrideInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

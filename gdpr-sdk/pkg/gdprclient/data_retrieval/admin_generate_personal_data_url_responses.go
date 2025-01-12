// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package data_retrieval

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/gdpr-sdk/pkg/gdprclientmodels"
)

// AdminGeneratePersonalDataURLReader is a Reader for the AdminGeneratePersonalDataURL structure.
type AdminGeneratePersonalDataURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGeneratePersonalDataURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGeneratePersonalDataURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminGeneratePersonalDataURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGeneratePersonalDataURLUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminGeneratePersonalDataURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminGeneratePersonalDataURLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /gdpr/admin/namespaces/{namespace}/users/{userId}/requests/{requestDate}/generate returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminGeneratePersonalDataURLOK creates a AdminGeneratePersonalDataURLOK with default headers values
func NewAdminGeneratePersonalDataURLOK() *AdminGeneratePersonalDataURLOK {
	return &AdminGeneratePersonalDataURLOK{}
}

/*AdminGeneratePersonalDataURLOK handles this case with default header values.

  OK
*/
type AdminGeneratePersonalDataURLOK struct {
	Payload *gdprclientmodels.ModelsUserDataURL
}

func (o *AdminGeneratePersonalDataURLOK) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/requests/{requestDate}/generate][%d] adminGeneratePersonalDataUrlOK  %+v", 200, o.ToJSONString())
}

func (o *AdminGeneratePersonalDataURLOK) ToJSONString() string {
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

func (o *AdminGeneratePersonalDataURLOK) GetPayload() *gdprclientmodels.ModelsUserDataURL {
	return o.Payload
}

func (o *AdminGeneratePersonalDataURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(gdprclientmodels.ModelsUserDataURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGeneratePersonalDataURLBadRequest creates a AdminGeneratePersonalDataURLBadRequest with default headers values
func NewAdminGeneratePersonalDataURLBadRequest() *AdminGeneratePersonalDataURLBadRequest {
	return &AdminGeneratePersonalDataURLBadRequest{}
}

/*AdminGeneratePersonalDataURLBadRequest handles this case with default header values.

  Bad Request
*/
type AdminGeneratePersonalDataURLBadRequest struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminGeneratePersonalDataURLBadRequest) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/requests/{requestDate}/generate][%d] adminGeneratePersonalDataUrlBadRequest  %+v", 400, o.ToJSONString())
}

func (o *AdminGeneratePersonalDataURLBadRequest) ToJSONString() string {
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

func (o *AdminGeneratePersonalDataURLBadRequest) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminGeneratePersonalDataURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGeneratePersonalDataURLUnauthorized creates a AdminGeneratePersonalDataURLUnauthorized with default headers values
func NewAdminGeneratePersonalDataURLUnauthorized() *AdminGeneratePersonalDataURLUnauthorized {
	return &AdminGeneratePersonalDataURLUnauthorized{}
}

/*AdminGeneratePersonalDataURLUnauthorized handles this case with default header values.

  Unauthorized
*/
type AdminGeneratePersonalDataURLUnauthorized struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminGeneratePersonalDataURLUnauthorized) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/requests/{requestDate}/generate][%d] adminGeneratePersonalDataUrlUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *AdminGeneratePersonalDataURLUnauthorized) ToJSONString() string {
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

func (o *AdminGeneratePersonalDataURLUnauthorized) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminGeneratePersonalDataURLUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGeneratePersonalDataURLNotFound creates a AdminGeneratePersonalDataURLNotFound with default headers values
func NewAdminGeneratePersonalDataURLNotFound() *AdminGeneratePersonalDataURLNotFound {
	return &AdminGeneratePersonalDataURLNotFound{}
}

/*AdminGeneratePersonalDataURLNotFound handles this case with default header values.

  Not Found
*/
type AdminGeneratePersonalDataURLNotFound struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminGeneratePersonalDataURLNotFound) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/requests/{requestDate}/generate][%d] adminGeneratePersonalDataUrlNotFound  %+v", 404, o.ToJSONString())
}

func (o *AdminGeneratePersonalDataURLNotFound) ToJSONString() string {
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

func (o *AdminGeneratePersonalDataURLNotFound) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminGeneratePersonalDataURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGeneratePersonalDataURLInternalServerError creates a AdminGeneratePersonalDataURLInternalServerError with default headers values
func NewAdminGeneratePersonalDataURLInternalServerError() *AdminGeneratePersonalDataURLInternalServerError {
	return &AdminGeneratePersonalDataURLInternalServerError{}
}

/*AdminGeneratePersonalDataURLInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminGeneratePersonalDataURLInternalServerError struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminGeneratePersonalDataURLInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/requests/{requestDate}/generate][%d] adminGeneratePersonalDataUrlInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *AdminGeneratePersonalDataURLInternalServerError) ToJSONString() string {
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

func (o *AdminGeneratePersonalDataURLInternalServerError) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminGeneratePersonalDataURLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package session

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/sessionbrowser-sdk/pkg/sessionbrowserclientmodels"
)

// DeleteSessionLocalDSReader is a Reader for the DeleteSessionLocalDS structure.
type DeleteSessionLocalDSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSessionLocalDSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSessionLocalDSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSessionLocalDSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSessionLocalDSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteSessionLocalDSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /sessionbrowser/namespaces/{namespace}/gamesession/{sessionID}/localds returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteSessionLocalDSOK creates a DeleteSessionLocalDSOK with default headers values
func NewDeleteSessionLocalDSOK() *DeleteSessionLocalDSOK {
	return &DeleteSessionLocalDSOK{}
}

/*DeleteSessionLocalDSOK handles this case with default header values.

  session get
*/
type DeleteSessionLocalDSOK struct {
	Payload *sessionbrowserclientmodels.ModelsSessionResponse
}

func (o *DeleteSessionLocalDSOK) Error() string {
	return fmt.Sprintf("[DELETE /sessionbrowser/namespaces/{namespace}/gamesession/{sessionID}/localds][%d] deleteSessionLocalDSOK  %+v", 200, o.ToJSONString())
}

func (o *DeleteSessionLocalDSOK) ToJSONString() string {
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

func (o *DeleteSessionLocalDSOK) GetPayload() *sessionbrowserclientmodels.ModelsSessionResponse {
	return o.Payload
}

func (o *DeleteSessionLocalDSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(sessionbrowserclientmodels.ModelsSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSessionLocalDSBadRequest creates a DeleteSessionLocalDSBadRequest with default headers values
func NewDeleteSessionLocalDSBadRequest() *DeleteSessionLocalDSBadRequest {
	return &DeleteSessionLocalDSBadRequest{}
}

/*DeleteSessionLocalDSBadRequest handles this case with default header values.

  malformed request
*/
type DeleteSessionLocalDSBadRequest struct {
	Payload *sessionbrowserclientmodels.ResponseError
}

func (o *DeleteSessionLocalDSBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sessionbrowser/namespaces/{namespace}/gamesession/{sessionID}/localds][%d] deleteSessionLocalDSBadRequest  %+v", 400, o.ToJSONString())
}

func (o *DeleteSessionLocalDSBadRequest) ToJSONString() string {
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

func (o *DeleteSessionLocalDSBadRequest) GetPayload() *sessionbrowserclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteSessionLocalDSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(sessionbrowserclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSessionLocalDSNotFound creates a DeleteSessionLocalDSNotFound with default headers values
func NewDeleteSessionLocalDSNotFound() *DeleteSessionLocalDSNotFound {
	return &DeleteSessionLocalDSNotFound{}
}

/*DeleteSessionLocalDSNotFound handles this case with default header values.

  session not found
*/
type DeleteSessionLocalDSNotFound struct {
	Payload *sessionbrowserclientmodels.ResponseError
}

func (o *DeleteSessionLocalDSNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sessionbrowser/namespaces/{namespace}/gamesession/{sessionID}/localds][%d] deleteSessionLocalDSNotFound  %+v", 404, o.ToJSONString())
}

func (o *DeleteSessionLocalDSNotFound) ToJSONString() string {
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

func (o *DeleteSessionLocalDSNotFound) GetPayload() *sessionbrowserclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteSessionLocalDSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(sessionbrowserclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSessionLocalDSInternalServerError creates a DeleteSessionLocalDSInternalServerError with default headers values
func NewDeleteSessionLocalDSInternalServerError() *DeleteSessionLocalDSInternalServerError {
	return &DeleteSessionLocalDSInternalServerError{}
}

/*DeleteSessionLocalDSInternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteSessionLocalDSInternalServerError struct {
	Payload *sessionbrowserclientmodels.ResponseError
}

func (o *DeleteSessionLocalDSInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /sessionbrowser/namespaces/{namespace}/gamesession/{sessionID}/localds][%d] deleteSessionLocalDSInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *DeleteSessionLocalDSInternalServerError) ToJSONString() string {
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

func (o *DeleteSessionLocalDSInternalServerError) GetPayload() *sessionbrowserclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteSessionLocalDSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(sessionbrowserclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

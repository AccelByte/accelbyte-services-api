// Code generated by go-swagger; DO NOT EDIT.

package server

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

// ShutdownServerReader is a Reader for the ShutdownServer structure.
type ShutdownServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewShutdownServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShutdownServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewShutdownServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewShutdownServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewShutdownServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dsmcontroller/namespaces/{namespace}/servers/shutdown returns an error %d: %s", response.Code(), string(data))
	}
}

// NewShutdownServerNoContent creates a ShutdownServerNoContent with default headers values
func NewShutdownServerNoContent() *ShutdownServerNoContent {
	return &ShutdownServerNoContent{}
}

/*ShutdownServerNoContent handles this case with default header values.

  server removed
*/
type ShutdownServerNoContent struct {
}

func (o *ShutdownServerNoContent) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/shutdown][%d] shutdownServerNoContent ", 204)
}

func (o *ShutdownServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShutdownServerBadRequest creates a ShutdownServerBadRequest with default headers values
func NewShutdownServerBadRequest() *ShutdownServerBadRequest {
	return &ShutdownServerBadRequest{}
}

/*ShutdownServerBadRequest handles this case with default header values.

  malformed request
*/
type ShutdownServerBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ShutdownServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/shutdown][%d] shutdownServerBadRequest  %+v", 400, o.Payload)
}

func (o *ShutdownServerBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ShutdownServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownServerUnauthorized creates a ShutdownServerUnauthorized with default headers values
func NewShutdownServerUnauthorized() *ShutdownServerUnauthorized {
	return &ShutdownServerUnauthorized{}
}

/*ShutdownServerUnauthorized handles this case with default header values.

  Unauthorized
*/
type ShutdownServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ShutdownServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/shutdown][%d] shutdownServerUnauthorized  %+v", 401, o.Payload)
}

func (o *ShutdownServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ShutdownServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownServerNotFound creates a ShutdownServerNotFound with default headers values
func NewShutdownServerNotFound() *ShutdownServerNotFound {
	return &ShutdownServerNotFound{}
}

/*ShutdownServerNotFound handles this case with default header values.

  server not found
*/
type ShutdownServerNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ShutdownServerNotFound) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/shutdown][%d] shutdownServerNotFound  %+v", 404, o.Payload)
}

func (o *ShutdownServerNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ShutdownServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownServerInternalServerError creates a ShutdownServerInternalServerError with default headers values
func NewShutdownServerInternalServerError() *ShutdownServerInternalServerError {
	return &ShutdownServerInternalServerError{}
}

/*ShutdownServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type ShutdownServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ShutdownServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dsmcontroller/namespaces/{namespace}/servers/shutdown][%d] shutdownServerInternalServerError  %+v", 500, o.Payload)
}

func (o *ShutdownServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ShutdownServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

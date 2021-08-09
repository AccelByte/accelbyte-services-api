// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// DeleteServerReader is a Reader for the DeleteServer structure.
type DeleteServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServerNoContent creates a DeleteServerNoContent with default headers values
func NewDeleteServerNoContent() *DeleteServerNoContent {
	return &DeleteServerNoContent{}
}

/*DeleteServerNoContent handles this case with default header values.

server deleted
*/
type DeleteServerNoContent struct {
}

func (o *DeleteServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}][%d] deleteServerNoContent ", 204)
}

func (o *DeleteServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerUnauthorized creates a DeleteServerUnauthorized with default headers values
func NewDeleteServerUnauthorized() *DeleteServerUnauthorized {
	return &DeleteServerUnauthorized{}
}

/*DeleteServerUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteServerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}][%d] deleteServerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServerNotFound creates a DeleteServerNotFound with default headers values
func NewDeleteServerNotFound() *DeleteServerNotFound {
	return &DeleteServerNotFound{}
}

/*DeleteServerNotFound handles this case with default header values.

server not found
*/
type DeleteServerNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteServerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}][%d] deleteServerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServerNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServerInternalServerError creates a DeleteServerInternalServerError with default headers values
func NewDeleteServerInternalServerError() *DeleteServerInternalServerError {
	return &DeleteServerInternalServerError{}
}

/*DeleteServerInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteServerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}][%d] deleteServerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// DeleteImageReader is a Reader for the DeleteImage structure.
type DeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteImageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteImageUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteImageNoContent creates a DeleteImageNoContent with default headers values
func NewDeleteImageNoContent() *DeleteImageNoContent {
	return &DeleteImageNoContent{}
}

/*DeleteImageNoContent handles this case with default header values.

image deleted
*/
type DeleteImageNoContent struct {
}

func (o *DeleteImageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/images][%d] deleteImageNoContent ", 204)
}

func (o *DeleteImageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteImageBadRequest creates a DeleteImageBadRequest with default headers values
func NewDeleteImageBadRequest() *DeleteImageBadRequest {
	return &DeleteImageBadRequest{}
}

/*DeleteImageBadRequest handles this case with default header values.

malformed request
*/
type DeleteImageBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteImageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/images][%d] deleteImageBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteImageBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageUnauthorized creates a DeleteImageUnauthorized with default headers values
func NewDeleteImageUnauthorized() *DeleteImageUnauthorized {
	return &DeleteImageUnauthorized{}
}

/*DeleteImageUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteImageUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteImageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/images][%d] deleteImageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteImageUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageNotFound creates a DeleteImageNotFound with default headers values
func NewDeleteImageNotFound() *DeleteImageNotFound {
	return &DeleteImageNotFound{}
}

/*DeleteImageNotFound handles this case with default header values.

malformed request
*/
type DeleteImageNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteImageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/images][%d] deleteImageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteImageNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageUnprocessableEntity creates a DeleteImageUnprocessableEntity with default headers values
func NewDeleteImageUnprocessableEntity() *DeleteImageUnprocessableEntity {
	return &DeleteImageUnprocessableEntity{}
}

/*DeleteImageUnprocessableEntity handles this case with default header values.

unprocessable entity
*/
type DeleteImageUnprocessableEntity struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteImageUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/images][%d] deleteImageUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteImageUnprocessableEntity) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteImageUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteImageInternalServerError creates a DeleteImageInternalServerError with default headers values
func NewDeleteImageInternalServerError() *DeleteImageInternalServerError {
	return &DeleteImageInternalServerError{}
}

/*DeleteImageInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteImageInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteImageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/images][%d] deleteImageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteImageInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// CountServerDetailedReader is a Reader for the CountServerDetailed structure.
type CountServerDetailedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountServerDetailedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountServerDetailedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCountServerDetailedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCountServerDetailedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCountServerDetailedOK creates a CountServerDetailedOK with default headers values
func NewCountServerDetailedOK() *CountServerDetailedOK {
	return &CountServerDetailedOK{}
}

/*CountServerDetailedOK handles this case with default header values.

servers listed
*/
type CountServerDetailedOK struct {
	Payload *dsmcclientmodels.ModelsDetailedCountServerResponse
}

func (o *CountServerDetailedOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/count/detailed][%d] countServerDetailedOK  %+v", 200, o.Payload)
}

func (o *CountServerDetailedOK) GetPayload() *dsmcclientmodels.ModelsDetailedCountServerResponse {
	return o.Payload
}

func (o *CountServerDetailedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsDetailedCountServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountServerDetailedUnauthorized creates a CountServerDetailedUnauthorized with default headers values
func NewCountServerDetailedUnauthorized() *CountServerDetailedUnauthorized {
	return &CountServerDetailedUnauthorized{}
}

/*CountServerDetailedUnauthorized handles this case with default header values.

Unauthorized
*/
type CountServerDetailedUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CountServerDetailedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/count/detailed][%d] countServerDetailedUnauthorized  %+v", 401, o.Payload)
}

func (o *CountServerDetailedUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CountServerDetailedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountServerDetailedInternalServerError creates a CountServerDetailedInternalServerError with default headers values
func NewCountServerDetailedInternalServerError() *CountServerDetailedInternalServerError {
	return &CountServerDetailedInternalServerError{}
}

/*CountServerDetailedInternalServerError handles this case with default header values.

Internal Server Error
*/
type CountServerDetailedInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *CountServerDetailedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/count/detailed][%d] countServerDetailedInternalServerError  %+v", 500, o.Payload)
}

func (o *CountServerDetailedInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *CountServerDetailedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

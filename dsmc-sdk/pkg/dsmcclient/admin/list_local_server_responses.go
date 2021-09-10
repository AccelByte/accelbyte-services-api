// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// ListLocalServerReader is a Reader for the ListLocalServer structure.
type ListLocalServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLocalServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLocalServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListLocalServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListLocalServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/admin/namespaces/{namespace}/servers/local returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListLocalServerOK creates a ListLocalServerOK with default headers values
func NewListLocalServerOK() *ListLocalServerOK {
	return &ListLocalServerOK{}
}

/*ListLocalServerOK handles this case with default header values.

  servers listed
*/
type ListLocalServerOK struct {
	Payload *dsmcclientmodels.ModelsListServerResponse
}

func (o *ListLocalServerOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/local][%d] listLocalServerOK  %+v", 200, o.Payload)
}

func (o *ListLocalServerOK) GetPayload() *dsmcclientmodels.ModelsListServerResponse {
	return o.Payload
}

func (o *ListLocalServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsListServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalServerUnauthorized creates a ListLocalServerUnauthorized with default headers values
func NewListLocalServerUnauthorized() *ListLocalServerUnauthorized {
	return &ListLocalServerUnauthorized{}
}

/*ListLocalServerUnauthorized handles this case with default header values.

  Unauthorized
*/
type ListLocalServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ListLocalServerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/local][%d] listLocalServerUnauthorized  %+v", 401, o.Payload)
}

func (o *ListLocalServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ListLocalServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalServerInternalServerError creates a ListLocalServerInternalServerError with default headers values
func NewListLocalServerInternalServerError() *ListLocalServerInternalServerError {
	return &ListLocalServerInternalServerError{}
}

/*ListLocalServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type ListLocalServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *ListLocalServerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/local][%d] listLocalServerInternalServerError  %+v", 500, o.Payload)
}

func (o *ListLocalServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *ListLocalServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

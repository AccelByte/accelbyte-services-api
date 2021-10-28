// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// GetNamespacePublisherReader is a Reader for the GetNamespacePublisher structure.
type GetNamespacePublisherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespacePublisherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespacePublisherOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNamespacePublisherBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNamespacePublisherUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetNamespacePublisherForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNamespacePublisherNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/admin/namespaces/{namespace}/publisher returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetNamespacePublisherOK creates a GetNamespacePublisherOK with default headers values
func NewGetNamespacePublisherOK() *GetNamespacePublisherOK {
	return &GetNamespacePublisherOK{}
}

/*GetNamespacePublisherOK handles this case with default header values.

  Successful operation
*/
type GetNamespacePublisherOK struct {
	Payload *basicclientmodels.NamespacePublisherInfo
}

func (o *GetNamespacePublisherOK) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/publisher][%d] getNamespacePublisherOK  %+v", 200, o.Payload)
}

func (o *GetNamespacePublisherOK) GetPayload() *basicclientmodels.NamespacePublisherInfo {
	return o.Payload
}

func (o *GetNamespacePublisherOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.NamespacePublisherInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacePublisherBadRequest creates a GetNamespacePublisherBadRequest with default headers values
func NewGetNamespacePublisherBadRequest() *GetNamespacePublisherBadRequest {
	return &GetNamespacePublisherBadRequest{}
}

/*GetNamespacePublisherBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type GetNamespacePublisherBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *GetNamespacePublisherBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/publisher][%d] getNamespacePublisherBadRequest  %+v", 400, o.Payload)
}

func (o *GetNamespacePublisherBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *GetNamespacePublisherBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacePublisherUnauthorized creates a GetNamespacePublisherUnauthorized with default headers values
func NewGetNamespacePublisherUnauthorized() *GetNamespacePublisherUnauthorized {
	return &GetNamespacePublisherUnauthorized{}
}

/*GetNamespacePublisherUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type GetNamespacePublisherUnauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetNamespacePublisherUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/publisher][%d] getNamespacePublisherUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNamespacePublisherUnauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetNamespacePublisherUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacePublisherForbidden creates a GetNamespacePublisherForbidden with default headers values
func NewGetNamespacePublisherForbidden() *GetNamespacePublisherForbidden {
	return &GetNamespacePublisherForbidden{}
}

/*GetNamespacePublisherForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permission</td></tr></table>
*/
type GetNamespacePublisherForbidden struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetNamespacePublisherForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/publisher][%d] getNamespacePublisherForbidden  %+v", 403, o.Payload)
}

func (o *GetNamespacePublisherForbidden) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetNamespacePublisherForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacePublisherNotFound creates a GetNamespacePublisherNotFound with default headers values
func NewGetNamespacePublisherNotFound() *GetNamespacePublisherNotFound {
	return &GetNamespacePublisherNotFound{}
}

/*GetNamespacePublisherNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>11337</td><td>Unable to {action}: Namespace not found</td></tr></table>
*/
type GetNamespacePublisherNotFound struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetNamespacePublisherNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/publisher][%d] getNamespacePublisherNotFound  %+v", 404, o.Payload)
}

func (o *GetNamespacePublisherNotFound) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetNamespacePublisherNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package concurrent_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// PutPlayerPublicRecordConcurrentHandlerV1Reader is a Reader for the PutPlayerPublicRecordConcurrentHandlerV1 structure.
type PutPlayerPublicRecordConcurrentHandlerV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPlayerPublicRecordConcurrentHandlerV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutPlayerPublicRecordConcurrentHandlerV1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutPlayerPublicRecordConcurrentHandlerV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutPlayerPublicRecordConcurrentHandlerV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 412:
		result := NewPutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutPlayerPublicRecordConcurrentHandlerV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /cloudsave/v1/namespaces/{namespace}/users/{userId}/concurrent/records/{key}/public returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPutPlayerPublicRecordConcurrentHandlerV1NoContent creates a PutPlayerPublicRecordConcurrentHandlerV1NoContent with default headers values
func NewPutPlayerPublicRecordConcurrentHandlerV1NoContent() *PutPlayerPublicRecordConcurrentHandlerV1NoContent {
	return &PutPlayerPublicRecordConcurrentHandlerV1NoContent{}
}

/*PutPlayerPublicRecordConcurrentHandlerV1NoContent handles this case with default header values.

  Record saved
*/
type PutPlayerPublicRecordConcurrentHandlerV1NoContent struct {
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1NoContent) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/users/{userId}/concurrent/records/{key}/public][%d] putPlayerPublicRecordConcurrentHandlerV1NoContent ", 204)
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutPlayerPublicRecordConcurrentHandlerV1BadRequest creates a PutPlayerPublicRecordConcurrentHandlerV1BadRequest with default headers values
func NewPutPlayerPublicRecordConcurrentHandlerV1BadRequest() *PutPlayerPublicRecordConcurrentHandlerV1BadRequest {
	return &PutPlayerPublicRecordConcurrentHandlerV1BadRequest{}
}

/*PutPlayerPublicRecordConcurrentHandlerV1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>18201</td><td>invalid record operator, expect [%s] but actual [%s]</td></tr></table>
*/
type PutPlayerPublicRecordConcurrentHandlerV1BadRequest struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/users/{userId}/concurrent/records/{key}/public][%d] putPlayerPublicRecordConcurrentHandlerV1BadRequest  %+v", 400, o.Payload)
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1BadRequest) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPlayerPublicRecordConcurrentHandlerV1Unauthorized creates a PutPlayerPublicRecordConcurrentHandlerV1Unauthorized with default headers values
func NewPutPlayerPublicRecordConcurrentHandlerV1Unauthorized() *PutPlayerPublicRecordConcurrentHandlerV1Unauthorized {
	return &PutPlayerPublicRecordConcurrentHandlerV1Unauthorized{}
}

/*PutPlayerPublicRecordConcurrentHandlerV1Unauthorized handles this case with default header values.

  Unauthorized
*/
type PutPlayerPublicRecordConcurrentHandlerV1Unauthorized struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/users/{userId}/concurrent/records/{key}/public][%d] putPlayerPublicRecordConcurrentHandlerV1Unauthorized  %+v", 401, o.Payload)
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1Unauthorized) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed creates a PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed with default headers values
func NewPutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed() *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed {
	return &PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed{}
}

/*PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed handles this case with default header values.

  Precondition Failed
*/
type PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/users/{userId}/concurrent/records/{key}/public][%d] putPlayerPublicRecordConcurrentHandlerV1PreconditionFailed  %+v", 412, o.Payload)
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1PreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutPlayerPublicRecordConcurrentHandlerV1InternalServerError creates a PutPlayerPublicRecordConcurrentHandlerV1InternalServerError with default headers values
func NewPutPlayerPublicRecordConcurrentHandlerV1InternalServerError() *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError {
	return &PutPlayerPublicRecordConcurrentHandlerV1InternalServerError{}
}

/*PutPlayerPublicRecordConcurrentHandlerV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type PutPlayerPublicRecordConcurrentHandlerV1InternalServerError struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/users/{userId}/concurrent/records/{key}/public][%d] putPlayerPublicRecordConcurrentHandlerV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *PutPlayerPublicRecordConcurrentHandlerV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

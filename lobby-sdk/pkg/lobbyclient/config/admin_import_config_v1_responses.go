// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// AdminImportConfigV1Reader is a Reader for the AdminImportConfigV1 structure.
type AdminImportConfigV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminImportConfigV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminImportConfigV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminImportConfigV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminImportConfigV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminImportConfigV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /lobby/v1/admin/config/namespaces/{namespace}/import returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminImportConfigV1OK creates a AdminImportConfigV1OK with default headers values
func NewAdminImportConfigV1OK() *AdminImportConfigV1OK {
	return &AdminImportConfigV1OK{}
}

/*AdminImportConfigV1OK handles this case with default header values.

  OK
*/
type AdminImportConfigV1OK struct {
	Payload *lobbyclientmodels.ModelsImportConfigResponse
}

func (o *AdminImportConfigV1OK) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/config/namespaces/{namespace}/import][%d] adminImportConfigV1OK  %+v", 200, o.Payload)
}

func (o *AdminImportConfigV1OK) GetPayload() *lobbyclientmodels.ModelsImportConfigResponse {
	return o.Payload
}

func (o *AdminImportConfigV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.ModelsImportConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminImportConfigV1Unauthorized creates a AdminImportConfigV1Unauthorized with default headers values
func NewAdminImportConfigV1Unauthorized() *AdminImportConfigV1Unauthorized {
	return &AdminImportConfigV1Unauthorized{}
}

/*AdminImportConfigV1Unauthorized handles this case with default header values.

  Unauthorized
*/
type AdminImportConfigV1Unauthorized struct {
	Payload *lobbyclientmodels.ResponseError
}

func (o *AdminImportConfigV1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/config/namespaces/{namespace}/import][%d] adminImportConfigV1Unauthorized  %+v", 401, o.Payload)
}

func (o *AdminImportConfigV1Unauthorized) GetPayload() *lobbyclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminImportConfigV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminImportConfigV1Forbidden creates a AdminImportConfigV1Forbidden with default headers values
func NewAdminImportConfigV1Forbidden() *AdminImportConfigV1Forbidden {
	return &AdminImportConfigV1Forbidden{}
}

/*AdminImportConfigV1Forbidden handles this case with default header values.

  Forbidden
*/
type AdminImportConfigV1Forbidden struct {
	Payload *lobbyclientmodels.ResponseError
}

func (o *AdminImportConfigV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/config/namespaces/{namespace}/import][%d] adminImportConfigV1Forbidden  %+v", 403, o.Payload)
}

func (o *AdminImportConfigV1Forbidden) GetPayload() *lobbyclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminImportConfigV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminImportConfigV1InternalServerError creates a AdminImportConfigV1InternalServerError with default headers values
func NewAdminImportConfigV1InternalServerError() *AdminImportConfigV1InternalServerError {
	return &AdminImportConfigV1InternalServerError{}
}

/*AdminImportConfigV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminImportConfigV1InternalServerError struct {
	Payload *lobbyclientmodels.ResponseError
}

func (o *AdminImportConfigV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/config/namespaces/{namespace}/import][%d] adminImportConfigV1InternalServerError  %+v", 500, o.Payload)
}

func (o *AdminImportConfigV1InternalServerError) GetPayload() *lobbyclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminImportConfigV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package profanity

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

// AdminAddProfanityFiltersReader is a Reader for the AdminAddProfanityFilters structure.
type AdminAddProfanityFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminAddProfanityFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminAddProfanityFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminAddProfanityFiltersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminAddProfanityFiltersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminAddProfanityFiltersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminAddProfanityFiltersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminAddProfanityFiltersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminAddProfanityFiltersOK creates a AdminAddProfanityFiltersOK with default headers values
func NewAdminAddProfanityFiltersOK() *AdminAddProfanityFiltersOK {
	return &AdminAddProfanityFiltersOK{}
}

/*AdminAddProfanityFiltersOK handles this case with default header values.

  OK
*/
type AdminAddProfanityFiltersOK struct {
}

func (o *AdminAddProfanityFiltersOK) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk][%d] adminAddProfanityFiltersOK ", 200)
}

func (o *AdminAddProfanityFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminAddProfanityFiltersBadRequest creates a AdminAddProfanityFiltersBadRequest with default headers values
func NewAdminAddProfanityFiltersBadRequest() *AdminAddProfanityFiltersBadRequest {
	return &AdminAddProfanityFiltersBadRequest{}
}

/*AdminAddProfanityFiltersBadRequest handles this case with default header values.

  Bad Request
*/
type AdminAddProfanityFiltersBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *AdminAddProfanityFiltersBadRequest) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk][%d] adminAddProfanityFiltersBadRequest  %+v", 400, o.Payload)
}

func (o *AdminAddProfanityFiltersBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *AdminAddProfanityFiltersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddProfanityFiltersUnauthorized creates a AdminAddProfanityFiltersUnauthorized with default headers values
func NewAdminAddProfanityFiltersUnauthorized() *AdminAddProfanityFiltersUnauthorized {
	return &AdminAddProfanityFiltersUnauthorized{}
}

/*AdminAddProfanityFiltersUnauthorized handles this case with default header values.

  Unauthorized
*/
type AdminAddProfanityFiltersUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *AdminAddProfanityFiltersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk][%d] adminAddProfanityFiltersUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminAddProfanityFiltersUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *AdminAddProfanityFiltersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddProfanityFiltersForbidden creates a AdminAddProfanityFiltersForbidden with default headers values
func NewAdminAddProfanityFiltersForbidden() *AdminAddProfanityFiltersForbidden {
	return &AdminAddProfanityFiltersForbidden{}
}

/*AdminAddProfanityFiltersForbidden handles this case with default header values.

  Forbidden
*/
type AdminAddProfanityFiltersForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *AdminAddProfanityFiltersForbidden) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk][%d] adminAddProfanityFiltersForbidden  %+v", 403, o.Payload)
}

func (o *AdminAddProfanityFiltersForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *AdminAddProfanityFiltersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddProfanityFiltersNotFound creates a AdminAddProfanityFiltersNotFound with default headers values
func NewAdminAddProfanityFiltersNotFound() *AdminAddProfanityFiltersNotFound {
	return &AdminAddProfanityFiltersNotFound{}
}

/*AdminAddProfanityFiltersNotFound handles this case with default header values.

  Not Found
*/
type AdminAddProfanityFiltersNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *AdminAddProfanityFiltersNotFound) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk][%d] adminAddProfanityFiltersNotFound  %+v", 404, o.Payload)
}

func (o *AdminAddProfanityFiltersNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *AdminAddProfanityFiltersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminAddProfanityFiltersInternalServerError creates a AdminAddProfanityFiltersInternalServerError with default headers values
func NewAdminAddProfanityFiltersInternalServerError() *AdminAddProfanityFiltersInternalServerError {
	return &AdminAddProfanityFiltersInternalServerError{}
}

/*AdminAddProfanityFiltersInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminAddProfanityFiltersInternalServerError struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *AdminAddProfanityFiltersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/profanity/namespaces/{namespace}/list/{list}/filters/bulk][%d] adminAddProfanityFiltersInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminAddProfanityFiltersInternalServerError) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *AdminAddProfanityFiltersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

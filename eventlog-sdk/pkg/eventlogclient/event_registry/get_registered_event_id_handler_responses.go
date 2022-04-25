// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package event_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclientmodels"
)

// GetRegisteredEventIDHandlerReader is a Reader for the GetRegisteredEventIDHandler structure.
type GetRegisteredEventIDHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegisteredEventIDHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegisteredEventIDHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegisteredEventIDHandlerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegisteredEventIDHandlerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRegisteredEventIDHandlerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRegisteredEventIDHandlerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetRegisteredEventIDHandlerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /event/registry/eventIds/{eventId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetRegisteredEventIDHandlerOK creates a GetRegisteredEventIDHandlerOK with default headers values
func NewGetRegisteredEventIDHandlerOK() *GetRegisteredEventIDHandlerOK {
	return &GetRegisteredEventIDHandlerOK{}
}

/*GetRegisteredEventIDHandlerOK handles this case with default header values.

  OK
*/
type GetRegisteredEventIDHandlerOK struct {
	Payload *eventlogclientmodels.ModelsEventRegistry
}

func (o *GetRegisteredEventIDHandlerOK) Error() string {
	return fmt.Sprintf("[GET /event/registry/eventIds/{eventId}][%d] getRegisteredEventIdHandlerOK  %+v", 200, o.Payload)
}

func (o *GetRegisteredEventIDHandlerOK) GetPayload() *eventlogclientmodels.ModelsEventRegistry {
	return o.Payload
}

func (o *GetRegisteredEventIDHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(eventlogclientmodels.ModelsEventRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegisteredEventIDHandlerBadRequest creates a GetRegisteredEventIDHandlerBadRequest with default headers values
func NewGetRegisteredEventIDHandlerBadRequest() *GetRegisteredEventIDHandlerBadRequest {
	return &GetRegisteredEventIDHandlerBadRequest{}
}

/*GetRegisteredEventIDHandlerBadRequest handles this case with default header values.

  Bad Request
*/
type GetRegisteredEventIDHandlerBadRequest struct {
}

func (o *GetRegisteredEventIDHandlerBadRequest) Error() string {
	return fmt.Sprintf("[GET /event/registry/eventIds/{eventId}][%d] getRegisteredEventIdHandlerBadRequest ", 400)
}

func (o *GetRegisteredEventIDHandlerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegisteredEventIDHandlerUnauthorized creates a GetRegisteredEventIDHandlerUnauthorized with default headers values
func NewGetRegisteredEventIDHandlerUnauthorized() *GetRegisteredEventIDHandlerUnauthorized {
	return &GetRegisteredEventIDHandlerUnauthorized{}
}

/*GetRegisteredEventIDHandlerUnauthorized handles this case with default header values.

  Unauthorized
*/
type GetRegisteredEventIDHandlerUnauthorized struct {
}

func (o *GetRegisteredEventIDHandlerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /event/registry/eventIds/{eventId}][%d] getRegisteredEventIdHandlerUnauthorized ", 401)
}

func (o *GetRegisteredEventIDHandlerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegisteredEventIDHandlerForbidden creates a GetRegisteredEventIDHandlerForbidden with default headers values
func NewGetRegisteredEventIDHandlerForbidden() *GetRegisteredEventIDHandlerForbidden {
	return &GetRegisteredEventIDHandlerForbidden{}
}

/*GetRegisteredEventIDHandlerForbidden handles this case with default header values.

  Forbidden
*/
type GetRegisteredEventIDHandlerForbidden struct {
}

func (o *GetRegisteredEventIDHandlerForbidden) Error() string {
	return fmt.Sprintf("[GET /event/registry/eventIds/{eventId}][%d] getRegisteredEventIdHandlerForbidden ", 403)
}

func (o *GetRegisteredEventIDHandlerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegisteredEventIDHandlerNotFound creates a GetRegisteredEventIDHandlerNotFound with default headers values
func NewGetRegisteredEventIDHandlerNotFound() *GetRegisteredEventIDHandlerNotFound {
	return &GetRegisteredEventIDHandlerNotFound{}
}

/*GetRegisteredEventIDHandlerNotFound handles this case with default header values.

  Not Found
*/
type GetRegisteredEventIDHandlerNotFound struct {
}

func (o *GetRegisteredEventIDHandlerNotFound) Error() string {
	return fmt.Sprintf("[GET /event/registry/eventIds/{eventId}][%d] getRegisteredEventIdHandlerNotFound ", 404)
}

func (o *GetRegisteredEventIDHandlerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegisteredEventIDHandlerInternalServerError creates a GetRegisteredEventIDHandlerInternalServerError with default headers values
func NewGetRegisteredEventIDHandlerInternalServerError() *GetRegisteredEventIDHandlerInternalServerError {
	return &GetRegisteredEventIDHandlerInternalServerError{}
}

/*GetRegisteredEventIDHandlerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetRegisteredEventIDHandlerInternalServerError struct {
}

func (o *GetRegisteredEventIDHandlerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /event/registry/eventIds/{eventId}][%d] getRegisteredEventIdHandlerInternalServerError ", 500)
}

func (o *GetRegisteredEventIDHandlerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

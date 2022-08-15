// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// DeleteTemplateSlugReader is a Reader for the DeleteTemplateSlug structure.
type DeleteTemplateSlugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTemplateSlugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTemplateSlugNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTemplateSlugBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTemplateSlugUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteTemplateSlugForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTemplateSlugNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /notification/namespaces/{namespace}/templates/{templateSlug} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteTemplateSlugNoContent creates a DeleteTemplateSlugNoContent with default headers values
func NewDeleteTemplateSlugNoContent() *DeleteTemplateSlugNoContent {
	return &DeleteTemplateSlugNoContent{}
}

/*DeleteTemplateSlugNoContent handles this case with default header values.

  No Content
*/
type DeleteTemplateSlugNoContent struct {
}

func (o *DeleteTemplateSlugNoContent) Error() string {
	return fmt.Sprintf("[DELETE /notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteTemplateSlugNoContent ", 204)
}

func (o *DeleteTemplateSlugNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTemplateSlugBadRequest creates a DeleteTemplateSlugBadRequest with default headers values
func NewDeleteTemplateSlugBadRequest() *DeleteTemplateSlugBadRequest {
	return &DeleteTemplateSlugBadRequest{}
}

/*DeleteTemplateSlugBadRequest handles this case with default header values.

  Bad Request
*/
type DeleteTemplateSlugBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *DeleteTemplateSlugBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteTemplateSlugBadRequest  %+v", 400, o.ToString())
}

func (o *DeleteTemplateSlugBadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteTemplateSlugBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *DeleteTemplateSlugBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTemplateSlugUnauthorized creates a DeleteTemplateSlugUnauthorized with default headers values
func NewDeleteTemplateSlugUnauthorized() *DeleteTemplateSlugUnauthorized {
	return &DeleteTemplateSlugUnauthorized{}
}

/*DeleteTemplateSlugUnauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteTemplateSlugUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *DeleteTemplateSlugUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteTemplateSlugUnauthorized  %+v", 401, o.ToString())
}

func (o *DeleteTemplateSlugUnauthorized) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteTemplateSlugUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *DeleteTemplateSlugUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTemplateSlugForbidden creates a DeleteTemplateSlugForbidden with default headers values
func NewDeleteTemplateSlugForbidden() *DeleteTemplateSlugForbidden {
	return &DeleteTemplateSlugForbidden{}
}

/*DeleteTemplateSlugForbidden handles this case with default header values.

  Forbidden
*/
type DeleteTemplateSlugForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *DeleteTemplateSlugForbidden) Error() string {
	return fmt.Sprintf("[DELETE /notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteTemplateSlugForbidden  %+v", 403, o.ToString())
}

func (o *DeleteTemplateSlugForbidden) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteTemplateSlugForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *DeleteTemplateSlugForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTemplateSlugNotFound creates a DeleteTemplateSlugNotFound with default headers values
func NewDeleteTemplateSlugNotFound() *DeleteTemplateSlugNotFound {
	return &DeleteTemplateSlugNotFound{}
}

/*DeleteTemplateSlugNotFound handles this case with default header values.

  Not Found
*/
type DeleteTemplateSlugNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *DeleteTemplateSlugNotFound) Error() string {
	return fmt.Sprintf("[DELETE /notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteTemplateSlugNotFound  %+v", 404, o.ToString())
}

func (o *DeleteTemplateSlugNotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteTemplateSlugNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *DeleteTemplateSlugNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin_player_record

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

// ListPlayerRecordHandlerV1Reader is a Reader for the ListPlayerRecordHandlerV1 structure.
type ListPlayerRecordHandlerV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPlayerRecordHandlerV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPlayerRecordHandlerV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPlayerRecordHandlerV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListPlayerRecordHandlerV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListPlayerRecordHandlerV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /cloudsave/v1/admin/namespaces/{namespace}/users/records returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListPlayerRecordHandlerV1OK creates a ListPlayerRecordHandlerV1OK with default headers values
func NewListPlayerRecordHandlerV1OK() *ListPlayerRecordHandlerV1OK {
	return &ListPlayerRecordHandlerV1OK{}
}

/*ListPlayerRecordHandlerV1OK handles this case with default header values.

  Successful operation
*/
type ListPlayerRecordHandlerV1OK struct {
	Payload *cloudsaveclientmodels.ModelsListPlayerRecordKeysResponse
}

func (o *ListPlayerRecordHandlerV1OK) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/users/records][%d] listPlayerRecordHandlerV1OK  %+v", 200, o.Payload)
}

func (o *ListPlayerRecordHandlerV1OK) GetPayload() *cloudsaveclientmodels.ModelsListPlayerRecordKeysResponse {
	return o.Payload
}

func (o *ListPlayerRecordHandlerV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsListPlayerRecordKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPlayerRecordHandlerV1BadRequest creates a ListPlayerRecordHandlerV1BadRequest with default headers values
func NewListPlayerRecordHandlerV1BadRequest() *ListPlayerRecordHandlerV1BadRequest {
	return &ListPlayerRecordHandlerV1BadRequest{}
}

/*ListPlayerRecordHandlerV1BadRequest handles this case with default header values.

  Bad Request
*/
type ListPlayerRecordHandlerV1BadRequest struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *ListPlayerRecordHandlerV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/users/records][%d] listPlayerRecordHandlerV1BadRequest  %+v", 400, o.Payload)
}

func (o *ListPlayerRecordHandlerV1BadRequest) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *ListPlayerRecordHandlerV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPlayerRecordHandlerV1Unauthorized creates a ListPlayerRecordHandlerV1Unauthorized with default headers values
func NewListPlayerRecordHandlerV1Unauthorized() *ListPlayerRecordHandlerV1Unauthorized {
	return &ListPlayerRecordHandlerV1Unauthorized{}
}

/*ListPlayerRecordHandlerV1Unauthorized handles this case with default header values.

  Unauthorized
*/
type ListPlayerRecordHandlerV1Unauthorized struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *ListPlayerRecordHandlerV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/users/records][%d] listPlayerRecordHandlerV1Unauthorized  %+v", 401, o.Payload)
}

func (o *ListPlayerRecordHandlerV1Unauthorized) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *ListPlayerRecordHandlerV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPlayerRecordHandlerV1InternalServerError creates a ListPlayerRecordHandlerV1InternalServerError with default headers values
func NewListPlayerRecordHandlerV1InternalServerError() *ListPlayerRecordHandlerV1InternalServerError {
	return &ListPlayerRecordHandlerV1InternalServerError{}
}

/*ListPlayerRecordHandlerV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type ListPlayerRecordHandlerV1InternalServerError struct {
	Payload *cloudsaveclientmodels.ModelsResponseError
}

func (o *ListPlayerRecordHandlerV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudsave/v1/admin/namespaces/{namespace}/users/records][%d] listPlayerRecordHandlerV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ListPlayerRecordHandlerV1InternalServerError) GetPayload() *cloudsaveclientmodels.ModelsResponseError {
	return o.Payload
}

func (o *ListPlayerRecordHandlerV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ModelsResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// GetServerLogsReader is a Reader for the GetServerLogs structure.
type GetServerLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServerLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetServerLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetServerLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetServerLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}/logs returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetServerLogsOK creates a GetServerLogsOK with default headers values
func NewGetServerLogsOK() *GetServerLogsOK {
	return &GetServerLogsOK{}
}

/*GetServerLogsOK handles this case with default header values.

  server logs queried
*/
type GetServerLogsOK struct {
	Payload *dsmcclientmodels.ModelsServerLogs
}

func (o *GetServerLogsOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}/logs][%d] getServerLogsOK  %+v", 200, o.ToString())
}

func (o *GetServerLogsOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetServerLogsOK) GetPayload() *dsmcclientmodels.ModelsServerLogs {
	return o.Payload
}

func (o *GetServerLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsServerLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerLogsBadRequest creates a GetServerLogsBadRequest with default headers values
func NewGetServerLogsBadRequest() *GetServerLogsBadRequest {
	return &GetServerLogsBadRequest{}
}

/*GetServerLogsBadRequest handles this case with default header values.

  malformed request
*/
type GetServerLogsBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetServerLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}/logs][%d] getServerLogsBadRequest  %+v", 400, o.ToString())
}

func (o *GetServerLogsBadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetServerLogsBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetServerLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerLogsUnauthorized creates a GetServerLogsUnauthorized with default headers values
func NewGetServerLogsUnauthorized() *GetServerLogsUnauthorized {
	return &GetServerLogsUnauthorized{}
}

/*GetServerLogsUnauthorized handles this case with default header values.

  Unauthorized
*/
type GetServerLogsUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetServerLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}/logs][%d] getServerLogsUnauthorized  %+v", 401, o.ToString())
}

func (o *GetServerLogsUnauthorized) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetServerLogsUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetServerLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerLogsNotFound creates a GetServerLogsNotFound with default headers values
func NewGetServerLogsNotFound() *GetServerLogsNotFound {
	return &GetServerLogsNotFound{}
}

/*GetServerLogsNotFound handles this case with default header values.

  server not found
*/
type GetServerLogsNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetServerLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}/logs][%d] getServerLogsNotFound  %+v", 404, o.ToString())
}

func (o *GetServerLogsNotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetServerLogsNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetServerLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerLogsInternalServerError creates a GetServerLogsInternalServerError with default headers values
func NewGetServerLogsInternalServerError() *GetServerLogsInternalServerError {
	return &GetServerLogsInternalServerError{}
}

/*GetServerLogsInternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetServerLogsInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetServerLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/servers/{podName}/logs][%d] getServerLogsInternalServerError  %+v", 500, o.ToString())
}

func (o *GetServerLogsInternalServerError) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetServerLogsInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetServerLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

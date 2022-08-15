// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/qosm-sdk/pkg/qosmclientmodels"
)

// ListServerReader is a Reader for the ListServer structure.
type ListServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /qosm/public/qos returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListServerOK creates a ListServerOK with default headers values
func NewListServerOK() *ListServerOK {
	return &ListServerOK{}
}

/*ListServerOK handles this case with default header values.

  list of QoS services returned
*/
type ListServerOK struct {
	Payload *qosmclientmodels.ModelsListServerResponse
}

func (o *ListServerOK) Error() string {
	return fmt.Sprintf("[GET /qosm/public/qos][%d] listServerOK  %+v", 200, o.ToString())
}

func (o *ListServerOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ListServerOK) GetPayload() *qosmclientmodels.ModelsListServerResponse {
	return o.Payload
}

func (o *ListServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(qosmclientmodels.ModelsListServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServerInternalServerError creates a ListServerInternalServerError with default headers values
func NewListServerInternalServerError() *ListServerInternalServerError {
	return &ListServerInternalServerError{}
}

/*ListServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type ListServerInternalServerError struct {
	Payload *qosmclientmodels.ResponseError
}

func (o *ListServerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /qosm/public/qos][%d] listServerInternalServerError  %+v", 500, o.ToString())
}

func (o *ListServerInternalServerError) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ListServerInternalServerError) GetPayload() *qosmclientmodels.ResponseError {
	return o.Payload
}

func (o *ListServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(qosmclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

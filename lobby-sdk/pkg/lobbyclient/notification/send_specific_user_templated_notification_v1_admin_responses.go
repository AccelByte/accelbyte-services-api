// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package notification

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// SendSpecificUserTemplatedNotificationV1AdminReader is a Reader for the SendSpecificUserTemplatedNotificationV1Admin structure.
type SendSpecificUserTemplatedNotificationV1AdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendSpecificUserTemplatedNotificationV1AdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSendSpecificUserTemplatedNotificationV1AdminNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendSpecificUserTemplatedNotificationV1AdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSendSpecificUserTemplatedNotificationV1AdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSendSpecificUserTemplatedNotificationV1AdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSendSpecificUserTemplatedNotificationV1AdminNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /lobby/v1/admin/notification/namespaces/{namespace}/users/{userId}/templates/notify returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSendSpecificUserTemplatedNotificationV1AdminNoContent creates a SendSpecificUserTemplatedNotificationV1AdminNoContent with default headers values
func NewSendSpecificUserTemplatedNotificationV1AdminNoContent() *SendSpecificUserTemplatedNotificationV1AdminNoContent {
	return &SendSpecificUserTemplatedNotificationV1AdminNoContent{}
}

/*SendSpecificUserTemplatedNotificationV1AdminNoContent handles this case with default header values.

  No Content
*/
type SendSpecificUserTemplatedNotificationV1AdminNoContent struct {
}

func (o *SendSpecificUserTemplatedNotificationV1AdminNoContent) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/notification/namespaces/{namespace}/users/{userId}/templates/notify][%d] sendSpecificUserTemplatedNotificationV1AdminNoContent ", 204)
}

func (o *SendSpecificUserTemplatedNotificationV1AdminNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewSendSpecificUserTemplatedNotificationV1AdminBadRequest creates a SendSpecificUserTemplatedNotificationV1AdminBadRequest with default headers values
func NewSendSpecificUserTemplatedNotificationV1AdminBadRequest() *SendSpecificUserTemplatedNotificationV1AdminBadRequest {
	return &SendSpecificUserTemplatedNotificationV1AdminBadRequest{}
}

/*SendSpecificUserTemplatedNotificationV1AdminBadRequest handles this case with default header values.

  Bad Request
*/
type SendSpecificUserTemplatedNotificationV1AdminBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *SendSpecificUserTemplatedNotificationV1AdminBadRequest) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/notification/namespaces/{namespace}/users/{userId}/templates/notify][%d] sendSpecificUserTemplatedNotificationV1AdminBadRequest  %+v", 400, o.ToJSONString())
}

func (o *SendSpecificUserTemplatedNotificationV1AdminBadRequest) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *SendSpecificUserTemplatedNotificationV1AdminBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *SendSpecificUserTemplatedNotificationV1AdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendSpecificUserTemplatedNotificationV1AdminUnauthorized creates a SendSpecificUserTemplatedNotificationV1AdminUnauthorized with default headers values
func NewSendSpecificUserTemplatedNotificationV1AdminUnauthorized() *SendSpecificUserTemplatedNotificationV1AdminUnauthorized {
	return &SendSpecificUserTemplatedNotificationV1AdminUnauthorized{}
}

/*SendSpecificUserTemplatedNotificationV1AdminUnauthorized handles this case with default header values.

  Unauthorized
*/
type SendSpecificUserTemplatedNotificationV1AdminUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *SendSpecificUserTemplatedNotificationV1AdminUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/notification/namespaces/{namespace}/users/{userId}/templates/notify][%d] sendSpecificUserTemplatedNotificationV1AdminUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *SendSpecificUserTemplatedNotificationV1AdminUnauthorized) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *SendSpecificUserTemplatedNotificationV1AdminUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *SendSpecificUserTemplatedNotificationV1AdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendSpecificUserTemplatedNotificationV1AdminForbidden creates a SendSpecificUserTemplatedNotificationV1AdminForbidden with default headers values
func NewSendSpecificUserTemplatedNotificationV1AdminForbidden() *SendSpecificUserTemplatedNotificationV1AdminForbidden {
	return &SendSpecificUserTemplatedNotificationV1AdminForbidden{}
}

/*SendSpecificUserTemplatedNotificationV1AdminForbidden handles this case with default header values.

  Forbidden
*/
type SendSpecificUserTemplatedNotificationV1AdminForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *SendSpecificUserTemplatedNotificationV1AdminForbidden) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/notification/namespaces/{namespace}/users/{userId}/templates/notify][%d] sendSpecificUserTemplatedNotificationV1AdminForbidden  %+v", 403, o.ToJSONString())
}

func (o *SendSpecificUserTemplatedNotificationV1AdminForbidden) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *SendSpecificUserTemplatedNotificationV1AdminForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *SendSpecificUserTemplatedNotificationV1AdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendSpecificUserTemplatedNotificationV1AdminNotFound creates a SendSpecificUserTemplatedNotificationV1AdminNotFound with default headers values
func NewSendSpecificUserTemplatedNotificationV1AdminNotFound() *SendSpecificUserTemplatedNotificationV1AdminNotFound {
	return &SendSpecificUserTemplatedNotificationV1AdminNotFound{}
}

/*SendSpecificUserTemplatedNotificationV1AdminNotFound handles this case with default header values.

  Not Found
*/
type SendSpecificUserTemplatedNotificationV1AdminNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *SendSpecificUserTemplatedNotificationV1AdminNotFound) Error() string {
	return fmt.Sprintf("[POST /lobby/v1/admin/notification/namespaces/{namespace}/users/{userId}/templates/notify][%d] sendSpecificUserTemplatedNotificationV1AdminNotFound  %+v", 404, o.ToJSONString())
}

func (o *SendSpecificUserTemplatedNotificationV1AdminNotFound) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *SendSpecificUserTemplatedNotificationV1AdminNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *SendSpecificUserTemplatedNotificationV1AdminNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

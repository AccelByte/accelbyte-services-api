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

// DeleteNotificationTemplateSlugV1AdminReader is a Reader for the DeleteNotificationTemplateSlugV1Admin structure.
type DeleteNotificationTemplateSlugV1AdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNotificationTemplateSlugV1AdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNotificationTemplateSlugV1AdminNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNotificationTemplateSlugV1AdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNotificationTemplateSlugV1AdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteNotificationTemplateSlugV1AdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteNotificationTemplateSlugV1AdminNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /lobby/v1/admin/notification/namespaces/{namespace}/templates/{templateSlug} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteNotificationTemplateSlugV1AdminNoContent creates a DeleteNotificationTemplateSlugV1AdminNoContent with default headers values
func NewDeleteNotificationTemplateSlugV1AdminNoContent() *DeleteNotificationTemplateSlugV1AdminNoContent {
	return &DeleteNotificationTemplateSlugV1AdminNoContent{}
}

/*DeleteNotificationTemplateSlugV1AdminNoContent handles this case with default header values.

  No Content
*/
type DeleteNotificationTemplateSlugV1AdminNoContent struct {
}

func (o *DeleteNotificationTemplateSlugV1AdminNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lobby/v1/admin/notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteNotificationTemplateSlugV1AdminNoContent ", 204)
}

func (o *DeleteNotificationTemplateSlugV1AdminNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteNotificationTemplateSlugV1AdminBadRequest creates a DeleteNotificationTemplateSlugV1AdminBadRequest with default headers values
func NewDeleteNotificationTemplateSlugV1AdminBadRequest() *DeleteNotificationTemplateSlugV1AdminBadRequest {
	return &DeleteNotificationTemplateSlugV1AdminBadRequest{}
}

/*DeleteNotificationTemplateSlugV1AdminBadRequest handles this case with default header values.

  Bad Request
*/
type DeleteNotificationTemplateSlugV1AdminBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *DeleteNotificationTemplateSlugV1AdminBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /lobby/v1/admin/notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteNotificationTemplateSlugV1AdminBadRequest  %+v", 400, o.ToJSONString())
}

func (o *DeleteNotificationTemplateSlugV1AdminBadRequest) ToJSONString() string {
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

func (o *DeleteNotificationTemplateSlugV1AdminBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *DeleteNotificationTemplateSlugV1AdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNotificationTemplateSlugV1AdminUnauthorized creates a DeleteNotificationTemplateSlugV1AdminUnauthorized with default headers values
func NewDeleteNotificationTemplateSlugV1AdminUnauthorized() *DeleteNotificationTemplateSlugV1AdminUnauthorized {
	return &DeleteNotificationTemplateSlugV1AdminUnauthorized{}
}

/*DeleteNotificationTemplateSlugV1AdminUnauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteNotificationTemplateSlugV1AdminUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *DeleteNotificationTemplateSlugV1AdminUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lobby/v1/admin/notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteNotificationTemplateSlugV1AdminUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *DeleteNotificationTemplateSlugV1AdminUnauthorized) ToJSONString() string {
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

func (o *DeleteNotificationTemplateSlugV1AdminUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *DeleteNotificationTemplateSlugV1AdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNotificationTemplateSlugV1AdminForbidden creates a DeleteNotificationTemplateSlugV1AdminForbidden with default headers values
func NewDeleteNotificationTemplateSlugV1AdminForbidden() *DeleteNotificationTemplateSlugV1AdminForbidden {
	return &DeleteNotificationTemplateSlugV1AdminForbidden{}
}

/*DeleteNotificationTemplateSlugV1AdminForbidden handles this case with default header values.

  Forbidden
*/
type DeleteNotificationTemplateSlugV1AdminForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *DeleteNotificationTemplateSlugV1AdminForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lobby/v1/admin/notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteNotificationTemplateSlugV1AdminForbidden  %+v", 403, o.ToJSONString())
}

func (o *DeleteNotificationTemplateSlugV1AdminForbidden) ToJSONString() string {
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

func (o *DeleteNotificationTemplateSlugV1AdminForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *DeleteNotificationTemplateSlugV1AdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNotificationTemplateSlugV1AdminNotFound creates a DeleteNotificationTemplateSlugV1AdminNotFound with default headers values
func NewDeleteNotificationTemplateSlugV1AdminNotFound() *DeleteNotificationTemplateSlugV1AdminNotFound {
	return &DeleteNotificationTemplateSlugV1AdminNotFound{}
}

/*DeleteNotificationTemplateSlugV1AdminNotFound handles this case with default header values.

  Not Found
*/
type DeleteNotificationTemplateSlugV1AdminNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *DeleteNotificationTemplateSlugV1AdminNotFound) Error() string {
	return fmt.Sprintf("[DELETE /lobby/v1/admin/notification/namespaces/{namespace}/templates/{templateSlug}][%d] deleteNotificationTemplateSlugV1AdminNotFound  %+v", 404, o.ToJSONString())
}

func (o *DeleteNotificationTemplateSlugV1AdminNotFound) ToJSONString() string {
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

func (o *DeleteNotificationTemplateSlugV1AdminNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *DeleteNotificationTemplateSlugV1AdminNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package admin_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// AdminUpdateContentS3Reader is a Reader for the AdminUpdateContentS3 structure.
type AdminUpdateContentS3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateContentS3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateContentS3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateContentS3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminUpdateContentS3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminUpdateContentS3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminUpdateContentS3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/s3/{contentId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminUpdateContentS3OK creates a AdminUpdateContentS3OK with default headers values
func NewAdminUpdateContentS3OK() *AdminUpdateContentS3OK {
	return &AdminUpdateContentS3OK{}
}

/*AdminUpdateContentS3OK handles this case with default header values.

  OK
*/
type AdminUpdateContentS3OK struct {
	Payload *ugcclientmodels.ModelsCreateContentResponse
}

func (o *AdminUpdateContentS3OK) Error() string {
	return fmt.Sprintf("[PUT /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/s3/{contentId}][%d] adminUpdateContentS3OK  %+v", 200, o.ToJSONString())
}

func (o *AdminUpdateContentS3OK) ToJSONString() string {
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

func (o *AdminUpdateContentS3OK) GetPayload() *ugcclientmodels.ModelsCreateContentResponse {
	return o.Payload
}

func (o *AdminUpdateContentS3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ModelsCreateContentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateContentS3BadRequest creates a AdminUpdateContentS3BadRequest with default headers values
func NewAdminUpdateContentS3BadRequest() *AdminUpdateContentS3BadRequest {
	return &AdminUpdateContentS3BadRequest{}
}

/*AdminUpdateContentS3BadRequest handles this case with default header values.

  Bad Request
*/
type AdminUpdateContentS3BadRequest struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUpdateContentS3BadRequest) Error() string {
	return fmt.Sprintf("[PUT /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/s3/{contentId}][%d] adminUpdateContentS3BadRequest  %+v", 400, o.ToJSONString())
}

func (o *AdminUpdateContentS3BadRequest) ToJSONString() string {
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

func (o *AdminUpdateContentS3BadRequest) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUpdateContentS3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateContentS3Unauthorized creates a AdminUpdateContentS3Unauthorized with default headers values
func NewAdminUpdateContentS3Unauthorized() *AdminUpdateContentS3Unauthorized {
	return &AdminUpdateContentS3Unauthorized{}
}

/*AdminUpdateContentS3Unauthorized handles this case with default header values.

  Unauthorized
*/
type AdminUpdateContentS3Unauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUpdateContentS3Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/s3/{contentId}][%d] adminUpdateContentS3Unauthorized  %+v", 401, o.ToJSONString())
}

func (o *AdminUpdateContentS3Unauthorized) ToJSONString() string {
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

func (o *AdminUpdateContentS3Unauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUpdateContentS3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateContentS3NotFound creates a AdminUpdateContentS3NotFound with default headers values
func NewAdminUpdateContentS3NotFound() *AdminUpdateContentS3NotFound {
	return &AdminUpdateContentS3NotFound{}
}

/*AdminUpdateContentS3NotFound handles this case with default header values.

  Not Found
*/
type AdminUpdateContentS3NotFound struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUpdateContentS3NotFound) Error() string {
	return fmt.Sprintf("[PUT /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/s3/{contentId}][%d] adminUpdateContentS3NotFound  %+v", 404, o.ToJSONString())
}

func (o *AdminUpdateContentS3NotFound) ToJSONString() string {
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

func (o *AdminUpdateContentS3NotFound) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUpdateContentS3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateContentS3InternalServerError creates a AdminUpdateContentS3InternalServerError with default headers values
func NewAdminUpdateContentS3InternalServerError() *AdminUpdateContentS3InternalServerError {
	return &AdminUpdateContentS3InternalServerError{}
}

/*AdminUpdateContentS3InternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminUpdateContentS3InternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUpdateContentS3InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/s3/{contentId}][%d] adminUpdateContentS3InternalServerError  %+v", 500, o.ToJSONString())
}

func (o *AdminUpdateContentS3InternalServerError) ToJSONString() string {
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

func (o *AdminUpdateContentS3InternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUpdateContentS3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

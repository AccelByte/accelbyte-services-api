// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package achievements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/achievement-sdk/pkg/achievementclientmodels"
)

// PublicUnlockAchievementReader is a Reader for the PublicUnlockAchievement structure.
type PublicUnlockAchievementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicUnlockAchievementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPublicUnlockAchievementNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicUnlockAchievementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicUnlockAchievementUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPublicUnlockAchievementInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /achievement/v1/public/namespaces/{namespace}/users/{userId}/achievements/{achievementCode}/unlock returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicUnlockAchievementNoContent creates a PublicUnlockAchievementNoContent with default headers values
func NewPublicUnlockAchievementNoContent() *PublicUnlockAchievementNoContent {
	return &PublicUnlockAchievementNoContent{}
}

/*PublicUnlockAchievementNoContent handles this case with default header values.

  No Content
*/
type PublicUnlockAchievementNoContent struct {
}

func (o *PublicUnlockAchievementNoContent) Error() string {
	return fmt.Sprintf("[PUT /achievement/v1/public/namespaces/{namespace}/users/{userId}/achievements/{achievementCode}/unlock][%d] publicUnlockAchievementNoContent ", 204)
}

func (o *PublicUnlockAchievementNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicUnlockAchievementBadRequest creates a PublicUnlockAchievementBadRequest with default headers values
func NewPublicUnlockAchievementBadRequest() *PublicUnlockAchievementBadRequest {
	return &PublicUnlockAchievementBadRequest{}
}

/*PublicUnlockAchievementBadRequest handles this case with default header values.

  Bad Request
*/
type PublicUnlockAchievementBadRequest struct {
	Payload *achievementclientmodels.ResponseError
}

func (o *PublicUnlockAchievementBadRequest) Error() string {
	return fmt.Sprintf("[PUT /achievement/v1/public/namespaces/{namespace}/users/{userId}/achievements/{achievementCode}/unlock][%d] publicUnlockAchievementBadRequest  %+v", 400, o.ToString())
}

func (o *PublicUnlockAchievementBadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicUnlockAchievementBadRequest) GetPayload() *achievementclientmodels.ResponseError {
	return o.Payload
}

func (o *PublicUnlockAchievementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(achievementclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicUnlockAchievementUnauthorized creates a PublicUnlockAchievementUnauthorized with default headers values
func NewPublicUnlockAchievementUnauthorized() *PublicUnlockAchievementUnauthorized {
	return &PublicUnlockAchievementUnauthorized{}
}

/*PublicUnlockAchievementUnauthorized handles this case with default header values.

  Unauthorized
*/
type PublicUnlockAchievementUnauthorized struct {
	Payload *achievementclientmodels.ResponseError
}

func (o *PublicUnlockAchievementUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /achievement/v1/public/namespaces/{namespace}/users/{userId}/achievements/{achievementCode}/unlock][%d] publicUnlockAchievementUnauthorized  %+v", 401, o.ToString())
}

func (o *PublicUnlockAchievementUnauthorized) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicUnlockAchievementUnauthorized) GetPayload() *achievementclientmodels.ResponseError {
	return o.Payload
}

func (o *PublicUnlockAchievementUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(achievementclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicUnlockAchievementInternalServerError creates a PublicUnlockAchievementInternalServerError with default headers values
func NewPublicUnlockAchievementInternalServerError() *PublicUnlockAchievementInternalServerError {
	return &PublicUnlockAchievementInternalServerError{}
}

/*PublicUnlockAchievementInternalServerError handles this case with default header values.

  Internal Server Error
*/
type PublicUnlockAchievementInternalServerError struct {
	Payload *achievementclientmodels.ResponseError
}

func (o *PublicUnlockAchievementInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /achievement/v1/public/namespaces/{namespace}/users/{userId}/achievements/{achievementCode}/unlock][%d] publicUnlockAchievementInternalServerError  %+v", 500, o.ToString())
}

func (o *PublicUnlockAchievementInternalServerError) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicUnlockAchievementInternalServerError) GetPayload() *achievementclientmodels.ResponseError {
	return o.Payload
}

func (o *PublicUnlockAchievementInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(achievementclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

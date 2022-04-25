// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package game_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// PublicDeleteProfileReader is a Reader for the PublicDeleteProfile structure.
type PublicDeleteProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicDeleteProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPublicDeleteProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicDeleteProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicDeleteProfileNoContent creates a PublicDeleteProfileNoContent with default headers values
func NewPublicDeleteProfileNoContent() *PublicDeleteProfileNoContent {
	return &PublicDeleteProfileNoContent{}
}

/*PublicDeleteProfileNoContent handles this case with default header values.

  Game profile has been deleted
*/
type PublicDeleteProfileNoContent struct {
}

func (o *PublicDeleteProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}][%d] publicDeleteProfileNoContent ", 204)
}

func (o *PublicDeleteProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicDeleteProfileNotFound creates a PublicDeleteProfileNotFound with default headers values
func NewPublicDeleteProfileNotFound() *PublicDeleteProfileNotFound {
	return &PublicDeleteProfileNotFound{}
}

/*PublicDeleteProfileNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12041</td><td>Game profile with id [{profileId}] is not found</td></tr></table>
*/
type PublicDeleteProfileNotFound struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *PublicDeleteProfileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}][%d] publicDeleteProfileNotFound  %+v", 404, o.Payload)
}

func (o *PublicDeleteProfileNotFound) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicDeleteProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

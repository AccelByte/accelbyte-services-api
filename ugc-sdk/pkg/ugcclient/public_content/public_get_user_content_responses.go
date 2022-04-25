// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package public_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// PublicGetUserContentReader is a Reader for the PublicGetUserContent structure.
type PublicGetUserContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetUserContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetUserContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicGetUserContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicGetUserContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPublicGetUserContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetUserContentOK creates a PublicGetUserContentOK with default headers values
func NewPublicGetUserContentOK() *PublicGetUserContentOK {
	return &PublicGetUserContentOK{}
}

/*PublicGetUserContentOK handles this case with default header values.

  OK
*/
type PublicGetUserContentOK struct {
	Payload *ugcclientmodels.ModelsPaginatedContentDownloadResponse
}

func (o *PublicGetUserContentOK) Error() string {
	return fmt.Sprintf("[GET /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents][%d] publicGetUserContentOK  %+v", 200, o.Payload)
}

func (o *PublicGetUserContentOK) GetPayload() *ugcclientmodels.ModelsPaginatedContentDownloadResponse {
	return o.Payload
}

func (o *PublicGetUserContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ModelsPaginatedContentDownloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetUserContentUnauthorized creates a PublicGetUserContentUnauthorized with default headers values
func NewPublicGetUserContentUnauthorized() *PublicGetUserContentUnauthorized {
	return &PublicGetUserContentUnauthorized{}
}

/*PublicGetUserContentUnauthorized handles this case with default header values.

  Unauthorized
*/
type PublicGetUserContentUnauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *PublicGetUserContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents][%d] publicGetUserContentUnauthorized  %+v", 401, o.Payload)
}

func (o *PublicGetUserContentUnauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *PublicGetUserContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetUserContentNotFound creates a PublicGetUserContentNotFound with default headers values
func NewPublicGetUserContentNotFound() *PublicGetUserContentNotFound {
	return &PublicGetUserContentNotFound{}
}

/*PublicGetUserContentNotFound handles this case with default header values.

  Not Found
*/
type PublicGetUserContentNotFound struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *PublicGetUserContentNotFound) Error() string {
	return fmt.Sprintf("[GET /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents][%d] publicGetUserContentNotFound  %+v", 404, o.Payload)
}

func (o *PublicGetUserContentNotFound) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *PublicGetUserContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetUserContentInternalServerError creates a PublicGetUserContentInternalServerError with default headers values
func NewPublicGetUserContentInternalServerError() *PublicGetUserContentInternalServerError {
	return &PublicGetUserContentInternalServerError{}
}

/*PublicGetUserContentInternalServerError handles this case with default header values.

  Internal Server Error
*/
type PublicGetUserContentInternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *PublicGetUserContentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents][%d] publicGetUserContentInternalServerError  %+v", 500, o.Payload)
}

func (o *PublicGetUserContentInternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *PublicGetUserContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

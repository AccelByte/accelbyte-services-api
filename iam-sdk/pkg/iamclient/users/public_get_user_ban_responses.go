// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// PublicGetUserBanReader is a Reader for the PublicGetUserBan structure.
type PublicGetUserBanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetUserBanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetUserBanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicGetUserBanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicGetUserBanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicGetUserBanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v2/public/namespaces/{namespace}/users/{userId}/bans returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetUserBanOK creates a PublicGetUserBanOK with default headers values
func NewPublicGetUserBanOK() *PublicGetUserBanOK {
	return &PublicGetUserBanOK{}
}

/*PublicGetUserBanOK handles this case with default header values.

  OK
*/
type PublicGetUserBanOK struct {
	Payload []*iamclientmodels.ModelUserBanResponse
}

func (o *PublicGetUserBanOK) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/users/{userId}/bans][%d] publicGetUserBanOK  %+v", 200, o.Payload)
}

func (o *PublicGetUserBanOK) GetPayload() []*iamclientmodels.ModelUserBanResponse {
	return o.Payload
}

func (o *PublicGetUserBanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetUserBanUnauthorized creates a PublicGetUserBanUnauthorized with default headers values
func NewPublicGetUserBanUnauthorized() *PublicGetUserBanUnauthorized {
	return &PublicGetUserBanUnauthorized{}
}

/*PublicGetUserBanUnauthorized handles this case with default header values.

  Unauthorized access
*/
type PublicGetUserBanUnauthorized struct {
}

func (o *PublicGetUserBanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/users/{userId}/bans][%d] publicGetUserBanUnauthorized ", 401)
}

func (o *PublicGetUserBanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicGetUserBanForbidden creates a PublicGetUserBanForbidden with default headers values
func NewPublicGetUserBanForbidden() *PublicGetUserBanForbidden {
	return &PublicGetUserBanForbidden{}
}

/*PublicGetUserBanForbidden handles this case with default header values.

  Forbidden
*/
type PublicGetUserBanForbidden struct {
}

func (o *PublicGetUserBanForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/users/{userId}/bans][%d] publicGetUserBanForbidden ", 403)
}

func (o *PublicGetUserBanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicGetUserBanNotFound creates a PublicGetUserBanNotFound with default headers values
func NewPublicGetUserBanNotFound() *PublicGetUserBanNotFound {
	return &PublicGetUserBanNotFound{}
}

/*PublicGetUserBanNotFound handles this case with default header values.

  Data not found
*/
type PublicGetUserBanNotFound struct {
}

func (o *PublicGetUserBanNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v2/public/namespaces/{namespace}/users/{userId}/bans][%d] publicGetUserBanNotFound ", 404)
}

func (o *PublicGetUserBanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
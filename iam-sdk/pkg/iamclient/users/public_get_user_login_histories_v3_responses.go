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

// PublicGetUserLoginHistoriesV3Reader is a Reader for the PublicGetUserLoginHistoriesV3 structure.
type PublicGetUserLoginHistoriesV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetUserLoginHistoriesV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetUserLoginHistoriesV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicGetUserLoginHistoriesV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicGetUserLoginHistoriesV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicGetUserLoginHistoriesV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/public/namespaces/{namespace}/users/{userId}/logins/histories returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetUserLoginHistoriesV3OK creates a PublicGetUserLoginHistoriesV3OK with default headers values
func NewPublicGetUserLoginHistoriesV3OK() *PublicGetUserLoginHistoriesV3OK {
	return &PublicGetUserLoginHistoriesV3OK{}
}

/*PublicGetUserLoginHistoriesV3OK handles this case with default header values.

  OK
*/
type PublicGetUserLoginHistoriesV3OK struct {
	Payload *iamclientmodels.ModelLoginHistoriesResponse
}

func (o *PublicGetUserLoginHistoriesV3OK) Error() string {
	return fmt.Sprintf("[GET /iam/v3/public/namespaces/{namespace}/users/{userId}/logins/histories][%d] publicGetUserLoginHistoriesV3OK  %+v", 200, o.Payload)
}

func (o *PublicGetUserLoginHistoriesV3OK) GetPayload() *iamclientmodels.ModelLoginHistoriesResponse {
	return o.Payload
}

func (o *PublicGetUserLoginHistoriesV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelLoginHistoriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetUserLoginHistoriesV3Unauthorized creates a PublicGetUserLoginHistoriesV3Unauthorized with default headers values
func NewPublicGetUserLoginHistoriesV3Unauthorized() *PublicGetUserLoginHistoriesV3Unauthorized {
	return &PublicGetUserLoginHistoriesV3Unauthorized{}
}

/*PublicGetUserLoginHistoriesV3Unauthorized handles this case with default header values.

  Unauthorized access
*/
type PublicGetUserLoginHistoriesV3Unauthorized struct {
}

func (o *PublicGetUserLoginHistoriesV3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v3/public/namespaces/{namespace}/users/{userId}/logins/histories][%d] publicGetUserLoginHistoriesV3Unauthorized ", 401)
}

func (o *PublicGetUserLoginHistoriesV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicGetUserLoginHistoriesV3Forbidden creates a PublicGetUserLoginHistoriesV3Forbidden with default headers values
func NewPublicGetUserLoginHistoriesV3Forbidden() *PublicGetUserLoginHistoriesV3Forbidden {
	return &PublicGetUserLoginHistoriesV3Forbidden{}
}

/*PublicGetUserLoginHistoriesV3Forbidden handles this case with default header values.

  Forbidden
*/
type PublicGetUserLoginHistoriesV3Forbidden struct {
}

func (o *PublicGetUserLoginHistoriesV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v3/public/namespaces/{namespace}/users/{userId}/logins/histories][%d] publicGetUserLoginHistoriesV3Forbidden ", 403)
}

func (o *PublicGetUserLoginHistoriesV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicGetUserLoginHistoriesV3NotFound creates a PublicGetUserLoginHistoriesV3NotFound with default headers values
func NewPublicGetUserLoginHistoriesV3NotFound() *PublicGetUserLoginHistoriesV3NotFound {
	return &PublicGetUserLoginHistoriesV3NotFound{}
}

/*PublicGetUserLoginHistoriesV3NotFound handles this case with default header values.

  Data not found
*/
type PublicGetUserLoginHistoriesV3NotFound struct {
}

func (o *PublicGetUserLoginHistoriesV3NotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v3/public/namespaces/{namespace}/users/{userId}/logins/histories][%d] publicGetUserLoginHistoriesV3NotFound ", 404)
}

func (o *PublicGetUserLoginHistoriesV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

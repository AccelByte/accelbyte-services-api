// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_credential

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

// UpdateSSOPlatformCredentialReader is a Reader for the UpdateSSOPlatformCredential structure.
type UpdateSSOPlatformCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSSOPlatformCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSSOPlatformCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSSOPlatformCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateSSOPlatformCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateSSOPlatformCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSSOPlatformCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateSSOPlatformCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateSSOPlatformCredentialOK creates a UpdateSSOPlatformCredentialOK with default headers values
func NewUpdateSSOPlatformCredentialOK() *UpdateSSOPlatformCredentialOK {
	return &UpdateSSOPlatformCredentialOK{}
}

/*UpdateSSOPlatformCredentialOK handles this case with default header values.

  OK
*/
type UpdateSSOPlatformCredentialOK struct {
	Payload *iamclientmodels.ModelSSOPlatformCredentialResponse
}

func (o *UpdateSSOPlatformCredentialOK) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] updateSSOPlatformCredentialOK  %+v", 200, o.Payload)
}

func (o *UpdateSSOPlatformCredentialOK) GetPayload() *iamclientmodels.ModelSSOPlatformCredentialResponse {
	return o.Payload
}

func (o *UpdateSSOPlatformCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelSSOPlatformCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSSOPlatformCredentialBadRequest creates a UpdateSSOPlatformCredentialBadRequest with default headers values
func NewUpdateSSOPlatformCredentialBadRequest() *UpdateSSOPlatformCredentialBadRequest {
	return &UpdateSSOPlatformCredentialBadRequest{}
}

/*UpdateSSOPlatformCredentialBadRequest handles this case with default header values.

  Bad Request
*/
type UpdateSSOPlatformCredentialBadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *UpdateSSOPlatformCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] updateSSOPlatformCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSSOPlatformCredentialBadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *UpdateSSOPlatformCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSSOPlatformCredentialUnauthorized creates a UpdateSSOPlatformCredentialUnauthorized with default headers values
func NewUpdateSSOPlatformCredentialUnauthorized() *UpdateSSOPlatformCredentialUnauthorized {
	return &UpdateSSOPlatformCredentialUnauthorized{}
}

/*UpdateSSOPlatformCredentialUnauthorized handles this case with default header values.

  Unauthorized
*/
type UpdateSSOPlatformCredentialUnauthorized struct {
}

func (o *UpdateSSOPlatformCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] updateSSOPlatformCredentialUnauthorized ", 401)
}

func (o *UpdateSSOPlatformCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSSOPlatformCredentialForbidden creates a UpdateSSOPlatformCredentialForbidden with default headers values
func NewUpdateSSOPlatformCredentialForbidden() *UpdateSSOPlatformCredentialForbidden {
	return &UpdateSSOPlatformCredentialForbidden{}
}

/*UpdateSSOPlatformCredentialForbidden handles this case with default header values.

  Forbidden
*/
type UpdateSSOPlatformCredentialForbidden struct {
}

func (o *UpdateSSOPlatformCredentialForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] updateSSOPlatformCredentialForbidden ", 403)
}

func (o *UpdateSSOPlatformCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSSOPlatformCredentialNotFound creates a UpdateSSOPlatformCredentialNotFound with default headers values
func NewUpdateSSOPlatformCredentialNotFound() *UpdateSSOPlatformCredentialNotFound {
	return &UpdateSSOPlatformCredentialNotFound{}
}

/*UpdateSSOPlatformCredentialNotFound handles this case with default header values.

  SSO Credential Not Found
*/
type UpdateSSOPlatformCredentialNotFound struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *UpdateSSOPlatformCredentialNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] updateSSOPlatformCredentialNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSSOPlatformCredentialNotFound) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *UpdateSSOPlatformCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSSOPlatformCredentialInternalServerError creates a UpdateSSOPlatformCredentialInternalServerError with default headers values
func NewUpdateSSOPlatformCredentialInternalServerError() *UpdateSSOPlatformCredentialInternalServerError {
	return &UpdateSSOPlatformCredentialInternalServerError{}
}

/*UpdateSSOPlatformCredentialInternalServerError handles this case with default header values.

  Internal Server Error
*/
type UpdateSSOPlatformCredentialInternalServerError struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *UpdateSSOPlatformCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /iam/v3/admin/namespaces/{namespace}/platforms/{platformId}/sso][%d] updateSSOPlatformCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSSOPlatformCredentialInternalServerError) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *UpdateSSOPlatformCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

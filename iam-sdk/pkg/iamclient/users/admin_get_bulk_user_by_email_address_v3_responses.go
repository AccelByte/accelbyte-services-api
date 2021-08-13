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

// AdminGetBulkUserByEmailAddressV3Reader is a Reader for the AdminGetBulkUserByEmailAddressV3 structure.
type AdminGetBulkUserByEmailAddressV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetBulkUserByEmailAddressV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetBulkUserByEmailAddressV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminGetBulkUserByEmailAddressV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminGetBulkUserByEmailAddressV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminGetBulkUserByEmailAddressV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminGetBulkUserByEmailAddressV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminGetBulkUserByEmailAddressV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminGetBulkUserByEmailAddressV3OK creates a AdminGetBulkUserByEmailAddressV3OK with default headers values
func NewAdminGetBulkUserByEmailAddressV3OK() *AdminGetBulkUserByEmailAddressV3OK {
	return &AdminGetBulkUserByEmailAddressV3OK{}
}

/*AdminGetBulkUserByEmailAddressV3OK handles this case with default header values.

  OK
*/
type AdminGetBulkUserByEmailAddressV3OK struct {
	Payload *iamclientmodels.ModelListUserResponseV3
}

func (o *AdminGetBulkUserByEmailAddressV3OK) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk][%d] adminGetBulkUserByEmailAddressV3OK  %+v", 200, o.Payload)
}

func (o *AdminGetBulkUserByEmailAddressV3OK) GetPayload() *iamclientmodels.ModelListUserResponseV3 {
	return o.Payload
}

func (o *AdminGetBulkUserByEmailAddressV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelListUserResponseV3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBulkUserByEmailAddressV3BadRequest creates a AdminGetBulkUserByEmailAddressV3BadRequest with default headers values
func NewAdminGetBulkUserByEmailAddressV3BadRequest() *AdminGetBulkUserByEmailAddressV3BadRequest {
	return &AdminGetBulkUserByEmailAddressV3BadRequest{}
}

/*AdminGetBulkUserByEmailAddressV3BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type AdminGetBulkUserByEmailAddressV3BadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminGetBulkUserByEmailAddressV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk][%d] adminGetBulkUserByEmailAddressV3BadRequest  %+v", 400, o.Payload)
}

func (o *AdminGetBulkUserByEmailAddressV3BadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminGetBulkUserByEmailAddressV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBulkUserByEmailAddressV3Unauthorized creates a AdminGetBulkUserByEmailAddressV3Unauthorized with default headers values
func NewAdminGetBulkUserByEmailAddressV3Unauthorized() *AdminGetBulkUserByEmailAddressV3Unauthorized {
	return &AdminGetBulkUserByEmailAddressV3Unauthorized{}
}

/*AdminGetBulkUserByEmailAddressV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type AdminGetBulkUserByEmailAddressV3Unauthorized struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminGetBulkUserByEmailAddressV3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk][%d] adminGetBulkUserByEmailAddressV3Unauthorized  %+v", 401, o.Payload)
}

func (o *AdminGetBulkUserByEmailAddressV3Unauthorized) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminGetBulkUserByEmailAddressV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBulkUserByEmailAddressV3Forbidden creates a AdminGetBulkUserByEmailAddressV3Forbidden with default headers values
func NewAdminGetBulkUserByEmailAddressV3Forbidden() *AdminGetBulkUserByEmailAddressV3Forbidden {
	return &AdminGetBulkUserByEmailAddressV3Forbidden{}
}

/*AdminGetBulkUserByEmailAddressV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr></table>
*/
type AdminGetBulkUserByEmailAddressV3Forbidden struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminGetBulkUserByEmailAddressV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk][%d] adminGetBulkUserByEmailAddressV3Forbidden  %+v", 403, o.Payload)
}

func (o *AdminGetBulkUserByEmailAddressV3Forbidden) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminGetBulkUserByEmailAddressV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBulkUserByEmailAddressV3NotFound creates a AdminGetBulkUserByEmailAddressV3NotFound with default headers values
func NewAdminGetBulkUserByEmailAddressV3NotFound() *AdminGetBulkUserByEmailAddressV3NotFound {
	return &AdminGetBulkUserByEmailAddressV3NotFound{}
}

/*AdminGetBulkUserByEmailAddressV3NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type AdminGetBulkUserByEmailAddressV3NotFound struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminGetBulkUserByEmailAddressV3NotFound) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk][%d] adminGetBulkUserByEmailAddressV3NotFound  %+v", 404, o.Payload)
}

func (o *AdminGetBulkUserByEmailAddressV3NotFound) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminGetBulkUserByEmailAddressV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetBulkUserByEmailAddressV3InternalServerError creates a AdminGetBulkUserByEmailAddressV3InternalServerError with default headers values
func NewAdminGetBulkUserByEmailAddressV3InternalServerError() *AdminGetBulkUserByEmailAddressV3InternalServerError {
	return &AdminGetBulkUserByEmailAddressV3InternalServerError{}
}

/*AdminGetBulkUserByEmailAddressV3InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type AdminGetBulkUserByEmailAddressV3InternalServerError struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *AdminGetBulkUserByEmailAddressV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/search/bulk][%d] adminGetBulkUserByEmailAddressV3InternalServerError  %+v", 500, o.Payload)
}

func (o *AdminGetBulkUserByEmailAddressV3InternalServerError) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *AdminGetBulkUserByEmailAddressV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

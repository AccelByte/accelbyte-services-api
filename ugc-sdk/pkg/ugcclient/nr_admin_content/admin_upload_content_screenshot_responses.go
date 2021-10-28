// Code generated by go-swagger; DO NOT EDIT.

package nr_admin_content

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

// AdminUploadContentScreenshotReader is a Reader for the AdminUploadContentScreenshot structure.
type AdminUploadContentScreenshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUploadContentScreenshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAdminUploadContentScreenshotCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUploadContentScreenshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminUploadContentScreenshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminUploadContentScreenshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /ugc/v1/admin/namespaces/{namespace}/contents/{contentId}/screenshots returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminUploadContentScreenshotCreated creates a AdminUploadContentScreenshotCreated with default headers values
func NewAdminUploadContentScreenshotCreated() *AdminUploadContentScreenshotCreated {
	return &AdminUploadContentScreenshotCreated{}
}

/*AdminUploadContentScreenshotCreated handles this case with default header values.

  Created
*/
type AdminUploadContentScreenshotCreated struct {
	Payload *ugcclientmodels.ModelsCreateScreenshotResponse
}

func (o *AdminUploadContentScreenshotCreated) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/admin/namespaces/{namespace}/contents/{contentId}/screenshots][%d] adminUploadContentScreenshotCreated  %+v", 201, o.Payload)
}

func (o *AdminUploadContentScreenshotCreated) GetPayload() *ugcclientmodels.ModelsCreateScreenshotResponse {
	return o.Payload
}

func (o *AdminUploadContentScreenshotCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ModelsCreateScreenshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUploadContentScreenshotBadRequest creates a AdminUploadContentScreenshotBadRequest with default headers values
func NewAdminUploadContentScreenshotBadRequest() *AdminUploadContentScreenshotBadRequest {
	return &AdminUploadContentScreenshotBadRequest{}
}

/*AdminUploadContentScreenshotBadRequest handles this case with default header values.

  Bad Request
*/
type AdminUploadContentScreenshotBadRequest struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUploadContentScreenshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/admin/namespaces/{namespace}/contents/{contentId}/screenshots][%d] adminUploadContentScreenshotBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUploadContentScreenshotBadRequest) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUploadContentScreenshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUploadContentScreenshotUnauthorized creates a AdminUploadContentScreenshotUnauthorized with default headers values
func NewAdminUploadContentScreenshotUnauthorized() *AdminUploadContentScreenshotUnauthorized {
	return &AdminUploadContentScreenshotUnauthorized{}
}

/*AdminUploadContentScreenshotUnauthorized handles this case with default header values.

  Unauthorized
*/
type AdminUploadContentScreenshotUnauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUploadContentScreenshotUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/admin/namespaces/{namespace}/contents/{contentId}/screenshots][%d] adminUploadContentScreenshotUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUploadContentScreenshotUnauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUploadContentScreenshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUploadContentScreenshotInternalServerError creates a AdminUploadContentScreenshotInternalServerError with default headers values
func NewAdminUploadContentScreenshotInternalServerError() *AdminUploadContentScreenshotInternalServerError {
	return &AdminUploadContentScreenshotInternalServerError{}
}

/*AdminUploadContentScreenshotInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminUploadContentScreenshotInternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminUploadContentScreenshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/admin/namespaces/{namespace}/contents/{contentId}/screenshots][%d] adminUploadContentScreenshotInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminUploadContentScreenshotInternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminUploadContentScreenshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package nr_public_content

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

// DeleteContentScreenshotReader is a Reader for the DeleteContentScreenshot structure.
type DeleteContentScreenshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentScreenshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteContentScreenshotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteContentScreenshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteContentScreenshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteContentScreenshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteContentScreenshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents/{contentId}/screenshots/{screenshotId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteContentScreenshotNoContent creates a DeleteContentScreenshotNoContent with default headers values
func NewDeleteContentScreenshotNoContent() *DeleteContentScreenshotNoContent {
	return &DeleteContentScreenshotNoContent{}
}

/*DeleteContentScreenshotNoContent handles this case with default header values.

  No Content
*/
type DeleteContentScreenshotNoContent struct {
}

func (o *DeleteContentScreenshotNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents/{contentId}/screenshots/{screenshotId}][%d] deleteContentScreenshotNoContent ", 204)
}

func (o *DeleteContentScreenshotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteContentScreenshotBadRequest creates a DeleteContentScreenshotBadRequest with default headers values
func NewDeleteContentScreenshotBadRequest() *DeleteContentScreenshotBadRequest {
	return &DeleteContentScreenshotBadRequest{}
}

/*DeleteContentScreenshotBadRequest handles this case with default header values.

  Bad Request
*/
type DeleteContentScreenshotBadRequest struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteContentScreenshotBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents/{contentId}/screenshots/{screenshotId}][%d] deleteContentScreenshotBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContentScreenshotBadRequest) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteContentScreenshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentScreenshotUnauthorized creates a DeleteContentScreenshotUnauthorized with default headers values
func NewDeleteContentScreenshotUnauthorized() *DeleteContentScreenshotUnauthorized {
	return &DeleteContentScreenshotUnauthorized{}
}

/*DeleteContentScreenshotUnauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteContentScreenshotUnauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteContentScreenshotUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents/{contentId}/screenshots/{screenshotId}][%d] deleteContentScreenshotUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteContentScreenshotUnauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteContentScreenshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentScreenshotNotFound creates a DeleteContentScreenshotNotFound with default headers values
func NewDeleteContentScreenshotNotFound() *DeleteContentScreenshotNotFound {
	return &DeleteContentScreenshotNotFound{}
}

/*DeleteContentScreenshotNotFound handles this case with default header values.

  Not Found
*/
type DeleteContentScreenshotNotFound struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteContentScreenshotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents/{contentId}/screenshots/{screenshotId}][%d] deleteContentScreenshotNotFound  %+v", 404, o.Payload)
}

func (o *DeleteContentScreenshotNotFound) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteContentScreenshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentScreenshotInternalServerError creates a DeleteContentScreenshotInternalServerError with default headers values
func NewDeleteContentScreenshotInternalServerError() *DeleteContentScreenshotInternalServerError {
	return &DeleteContentScreenshotInternalServerError{}
}

/*DeleteContentScreenshotInternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteContentScreenshotInternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *DeleteContentScreenshotInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/public/namespaces/{namespace}/users/{userId}/contents/{contentId}/screenshots/{screenshotId}][%d] deleteContentScreenshotInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteContentScreenshotInternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteContentScreenshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

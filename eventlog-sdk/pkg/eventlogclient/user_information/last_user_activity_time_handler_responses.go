// Code generated by go-swagger; DO NOT EDIT.

package user_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclientmodels"
)

// LastUserActivityTimeHandlerReader is a Reader for the LastUserActivityTimeHandler structure.
type LastUserActivityTimeHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LastUserActivityTimeHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLastUserActivityTimeHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLastUserActivityTimeHandlerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewLastUserActivityTimeHandlerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLastUserActivityTimeHandlerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewLastUserActivityTimeHandlerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /event/namespaces/{namespace}/users/{userId}/lastActivityTime returns an error %d: %s", response.Code(), string(data))
	}
}

// NewLastUserActivityTimeHandlerOK creates a LastUserActivityTimeHandlerOK with default headers values
func NewLastUserActivityTimeHandlerOK() *LastUserActivityTimeHandlerOK {
	return &LastUserActivityTimeHandlerOK{}
}

/*LastUserActivityTimeHandlerOK handles this case with default header values.

  OK
*/
type LastUserActivityTimeHandlerOK struct {
	Payload *eventlogclientmodels.ModelsUserLastActivity
}

func (o *LastUserActivityTimeHandlerOK) Error() string {
	return fmt.Sprintf("[GET /event/namespaces/{namespace}/users/{userId}/lastActivityTime][%d] lastUserActivityTimeHandlerOK  %+v", 200, o.Payload)
}

func (o *LastUserActivityTimeHandlerOK) GetPayload() *eventlogclientmodels.ModelsUserLastActivity {
	return o.Payload
}

func (o *LastUserActivityTimeHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(eventlogclientmodels.ModelsUserLastActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLastUserActivityTimeHandlerUnauthorized creates a LastUserActivityTimeHandlerUnauthorized with default headers values
func NewLastUserActivityTimeHandlerUnauthorized() *LastUserActivityTimeHandlerUnauthorized {
	return &LastUserActivityTimeHandlerUnauthorized{}
}

/*LastUserActivityTimeHandlerUnauthorized handles this case with default header values.

  Unauthorized
*/
type LastUserActivityTimeHandlerUnauthorized struct {
}

func (o *LastUserActivityTimeHandlerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /event/namespaces/{namespace}/users/{userId}/lastActivityTime][%d] lastUserActivityTimeHandlerUnauthorized ", 401)
}

func (o *LastUserActivityTimeHandlerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLastUserActivityTimeHandlerForbidden creates a LastUserActivityTimeHandlerForbidden with default headers values
func NewLastUserActivityTimeHandlerForbidden() *LastUserActivityTimeHandlerForbidden {
	return &LastUserActivityTimeHandlerForbidden{}
}

/*LastUserActivityTimeHandlerForbidden handles this case with default header values.

  Forbidden
*/
type LastUserActivityTimeHandlerForbidden struct {
}

func (o *LastUserActivityTimeHandlerForbidden) Error() string {
	return fmt.Sprintf("[GET /event/namespaces/{namespace}/users/{userId}/lastActivityTime][%d] lastUserActivityTimeHandlerForbidden ", 403)
}

func (o *LastUserActivityTimeHandlerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLastUserActivityTimeHandlerNotFound creates a LastUserActivityTimeHandlerNotFound with default headers values
func NewLastUserActivityTimeHandlerNotFound() *LastUserActivityTimeHandlerNotFound {
	return &LastUserActivityTimeHandlerNotFound{}
}

/*LastUserActivityTimeHandlerNotFound handles this case with default header values.

  Not Found
*/
type LastUserActivityTimeHandlerNotFound struct {
}

func (o *LastUserActivityTimeHandlerNotFound) Error() string {
	return fmt.Sprintf("[GET /event/namespaces/{namespace}/users/{userId}/lastActivityTime][%d] lastUserActivityTimeHandlerNotFound ", 404)
}

func (o *LastUserActivityTimeHandlerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLastUserActivityTimeHandlerInternalServerError creates a LastUserActivityTimeHandlerInternalServerError with default headers values
func NewLastUserActivityTimeHandlerInternalServerError() *LastUserActivityTimeHandlerInternalServerError {
	return &LastUserActivityTimeHandlerInternalServerError{}
}

/*LastUserActivityTimeHandlerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type LastUserActivityTimeHandlerInternalServerError struct {
}

func (o *LastUserActivityTimeHandlerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /event/namespaces/{namespace}/users/{userId}/lastActivityTime][%d] lastUserActivityTimeHandlerInternalServerError ", 500)
}

func (o *LastUserActivityTimeHandlerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package event_descriptions

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

// EventIDDescriptionHandlerReader is a Reader for the EventIDDescriptionHandler structure.
type EventIDDescriptionHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventIDDescriptionHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEventIDDescriptionHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /event/descriptions/eventId returns an error %d: %s", response.Code(), string(data))
	}
}

// NewEventIDDescriptionHandlerOK creates a EventIDDescriptionHandlerOK with default headers values
func NewEventIDDescriptionHandlerOK() *EventIDDescriptionHandlerOK {
	return &EventIDDescriptionHandlerOK{}
}

/*EventIDDescriptionHandlerOK handles this case with default header values.

  OK
*/
type EventIDDescriptionHandlerOK struct {
	Payload *eventlogclientmodels.ModelsMultipleEventID
}

func (o *EventIDDescriptionHandlerOK) Error() string {
	return fmt.Sprintf("[GET /event/descriptions/eventId][%d] eventIdDescriptionHandlerOK  %+v", 200, o.Payload)
}

func (o *EventIDDescriptionHandlerOK) GetPayload() *eventlogclientmodels.ModelsMultipleEventID {
	return o.Payload
}

func (o *EventIDDescriptionHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(eventlogclientmodels.ModelsMultipleEventID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

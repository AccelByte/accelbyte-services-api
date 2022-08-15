// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package event_descriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclientmodels"
)

// UXNameDescriptionHandlerReader is a Reader for the UXNameDescriptionHandler structure.
type UXNameDescriptionHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UXNameDescriptionHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUXNameDescriptionHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /event/descriptions/ux returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUXNameDescriptionHandlerOK creates a UXNameDescriptionHandlerOK with default headers values
func NewUXNameDescriptionHandlerOK() *UXNameDescriptionHandlerOK {
	return &UXNameDescriptionHandlerOK{}
}

/*UXNameDescriptionHandlerOK handles this case with default header values.

  OK
*/
type UXNameDescriptionHandlerOK struct {
	Payload *eventlogclientmodels.ModelsMultipleUX
}

func (o *UXNameDescriptionHandlerOK) Error() string {
	return fmt.Sprintf("[GET /event/descriptions/ux][%d] uXNameDescriptionHandlerOK  %+v", 200, o.ToString())
}

func (o *UXNameDescriptionHandlerOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *UXNameDescriptionHandlerOK) GetPayload() *eventlogclientmodels.ModelsMultipleUX {
	return o.Payload
}

func (o *UXNameDescriptionHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(eventlogclientmodels.ModelsMultipleUX)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

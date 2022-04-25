// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

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

// AgentTypeDescriptionHandlerReader is a Reader for the AgentTypeDescriptionHandler structure.
type AgentTypeDescriptionHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AgentTypeDescriptionHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAgentTypeDescriptionHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /event/descriptions/agentType returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAgentTypeDescriptionHandlerOK creates a AgentTypeDescriptionHandlerOK with default headers values
func NewAgentTypeDescriptionHandlerOK() *AgentTypeDescriptionHandlerOK {
	return &AgentTypeDescriptionHandlerOK{}
}

/*AgentTypeDescriptionHandlerOK handles this case with default header values.

  OK
*/
type AgentTypeDescriptionHandlerOK struct {
	Payload *eventlogclientmodels.ModelsMultipleAgentType
}

func (o *AgentTypeDescriptionHandlerOK) Error() string {
	return fmt.Sprintf("[GET /event/descriptions/agentType][%d] agentTypeDescriptionHandlerOK  %+v", 200, o.Payload)
}

func (o *AgentTypeDescriptionHandlerOK) GetPayload() *eventlogclientmodels.ModelsMultipleAgentType {
	return o.Payload
}

func (o *AgentTypeDescriptionHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(eventlogclientmodels.ModelsMultipleAgentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

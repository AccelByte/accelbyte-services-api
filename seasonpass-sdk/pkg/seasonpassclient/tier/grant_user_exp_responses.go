// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package tier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
)

// GrantUserExpReader is a Reader for the GrantUserExp structure.
type GrantUserExpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GrantUserExpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGrantUserExpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGrantUserExpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /seasonpass/admin/namespaces/{namespace}/users/{userId}/seasons/current/exp returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGrantUserExpOK creates a GrantUserExpOK with default headers values
func NewGrantUserExpOK() *GrantUserExpOK {
	return &GrantUserExpOK{}
}

/*GrantUserExpOK handles this case with default header values.

  successful operation
*/
type GrantUserExpOK struct {
	Payload *seasonpassclientmodels.UserSeasonSummary
}

func (o *GrantUserExpOK) Error() string {
	return fmt.Sprintf("[POST /seasonpass/admin/namespaces/{namespace}/users/{userId}/seasons/current/exp][%d] grantUserExpOK  %+v", 200, o.Payload)
}

func (o *GrantUserExpOK) GetPayload() *seasonpassclientmodels.UserSeasonSummary {
	return o.Payload
}

func (o *GrantUserExpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(seasonpassclientmodels.UserSeasonSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserExpBadRequest creates a GrantUserExpBadRequest with default headers values
func NewGrantUserExpBadRequest() *GrantUserExpBadRequest {
	return &GrantUserExpBadRequest{}
}

/*GrantUserExpBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20026</td><td>publisher namespace not allowed</td></tr></table>
*/
type GrantUserExpBadRequest struct {
	Payload *seasonpassclientmodels.ErrorEntity
}

func (o *GrantUserExpBadRequest) Error() string {
	return fmt.Sprintf("[POST /seasonpass/admin/namespaces/{namespace}/users/{userId}/seasons/current/exp][%d] grantUserExpBadRequest  %+v", 400, o.Payload)
}

func (o *GrantUserExpBadRequest) GetPayload() *seasonpassclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GrantUserExpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(seasonpassclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

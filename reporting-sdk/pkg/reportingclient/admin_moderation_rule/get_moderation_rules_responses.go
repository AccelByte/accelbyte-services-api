// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package admin_moderation_rule

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/reporting-sdk/pkg/reportingclientmodels"
)

// GetModerationRulesReader is a Reader for the GetModerationRules structure.
type GetModerationRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModerationRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModerationRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetModerationRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetModerationRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetModerationRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /reporting/v1/admin/namespaces/{namespace}/rules returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetModerationRulesOK creates a GetModerationRulesOK with default headers values
func NewGetModerationRulesOK() *GetModerationRulesOK {
	return &GetModerationRulesOK{}
}

/*GetModerationRulesOK handles this case with default header values.

  OK
*/
type GetModerationRulesOK struct {
	Payload *reportingclientmodels.RestapiModerationRulesList
}

func (o *GetModerationRulesOK) Error() string {
	return fmt.Sprintf("[GET /reporting/v1/admin/namespaces/{namespace}/rules][%d] getModerationRulesOK  %+v", 200, o.ToJSONString())
}

func (o *GetModerationRulesOK) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetModerationRulesOK) GetPayload() *reportingclientmodels.RestapiModerationRulesList {
	return o.Payload
}

func (o *GetModerationRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(reportingclientmodels.RestapiModerationRulesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModerationRulesBadRequest creates a GetModerationRulesBadRequest with default headers values
func NewGetModerationRulesBadRequest() *GetModerationRulesBadRequest {
	return &GetModerationRulesBadRequest{}
}

/*GetModerationRulesBadRequest handles this case with default header values.

  Bad Request
*/
type GetModerationRulesBadRequest struct {
	Payload *reportingclientmodels.RestapiErrorResponse
}

func (o *GetModerationRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting/v1/admin/namespaces/{namespace}/rules][%d] getModerationRulesBadRequest  %+v", 400, o.ToJSONString())
}

func (o *GetModerationRulesBadRequest) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetModerationRulesBadRequest) GetPayload() *reportingclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetModerationRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(reportingclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModerationRulesNotFound creates a GetModerationRulesNotFound with default headers values
func NewGetModerationRulesNotFound() *GetModerationRulesNotFound {
	return &GetModerationRulesNotFound{}
}

/*GetModerationRulesNotFound handles this case with default header values.

  Not Found
*/
type GetModerationRulesNotFound struct {
	Payload *reportingclientmodels.RestapiErrorResponse
}

func (o *GetModerationRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v1/admin/namespaces/{namespace}/rules][%d] getModerationRulesNotFound  %+v", 404, o.ToJSONString())
}

func (o *GetModerationRulesNotFound) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetModerationRulesNotFound) GetPayload() *reportingclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetModerationRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(reportingclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModerationRulesInternalServerError creates a GetModerationRulesInternalServerError with default headers values
func NewGetModerationRulesInternalServerError() *GetModerationRulesInternalServerError {
	return &GetModerationRulesInternalServerError{}
}

/*GetModerationRulesInternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetModerationRulesInternalServerError struct {
	Payload *reportingclientmodels.RestapiErrorResponse
}

func (o *GetModerationRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reporting/v1/admin/namespaces/{namespace}/rules][%d] getModerationRulesInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *GetModerationRulesInternalServerError) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetModerationRulesInternalServerError) GetPayload() *reportingclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *GetModerationRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(reportingclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

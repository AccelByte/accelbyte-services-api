// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package policy_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
)

// PublishPolicyVersionReader is a Reader for the PublishPolicyVersion structure.
type PublishPolicyVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublishPolicyVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublishPolicyVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublishPolicyVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /agreement/admin/policies/versions/{policyVersionId}/latest returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublishPolicyVersionOK creates a PublishPolicyVersionOK with default headers values
func NewPublishPolicyVersionOK() *PublishPolicyVersionOK {
	return &PublishPolicyVersionOK{}
}

/*PublishPolicyVersionOK handles this case with default header values.

  successful operation
*/
type PublishPolicyVersionOK struct {
}

func (o *PublishPolicyVersionOK) Error() string {
	return fmt.Sprintf("[PATCH /agreement/admin/policies/versions/{policyVersionId}/latest][%d] publishPolicyVersionOK ", 200)
}

func (o *PublishPolicyVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublishPolicyVersionBadRequest creates a PublishPolicyVersionBadRequest with default headers values
func NewPublishPolicyVersionBadRequest() *PublishPolicyVersionBadRequest {
	return &PublishPolicyVersionBadRequest{}
}

/*PublishPolicyVersionBadRequest handles this case with default header values.

  <table><tr><td>NumericErrorCode</td><td>ErrorCode</td></tr><tr><td>40033</td><td>errors.net.accelbyte.platform.legal.invalid_policy_version</td></tr></table>
*/
type PublishPolicyVersionBadRequest struct {
	Payload *legalclientmodels.ErrorEntity
}

func (o *PublishPolicyVersionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /agreement/admin/policies/versions/{policyVersionId}/latest][%d] publishPolicyVersionBadRequest  %+v", 400, o.ToString())
}

func (o *PublishPolicyVersionBadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublishPolicyVersionBadRequest) GetPayload() *legalclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublishPolicyVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legalclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

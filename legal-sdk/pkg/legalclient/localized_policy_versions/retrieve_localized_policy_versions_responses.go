// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package localized_policy_versions

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

// RetrieveLocalizedPolicyVersionsReader is a Reader for the RetrieveLocalizedPolicyVersions structure.
type RetrieveLocalizedPolicyVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveLocalizedPolicyVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveLocalizedPolicyVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /agreement/admin/localized-policy-versions/versions/{policyVersionId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRetrieveLocalizedPolicyVersionsOK creates a RetrieveLocalizedPolicyVersionsOK with default headers values
func NewRetrieveLocalizedPolicyVersionsOK() *RetrieveLocalizedPolicyVersionsOK {
	return &RetrieveLocalizedPolicyVersionsOK{}
}

/*RetrieveLocalizedPolicyVersionsOK handles this case with default header values.

  successful operation
*/
type RetrieveLocalizedPolicyVersionsOK struct {
	Payload []*legalclientmodels.RetrieveLocalizedPolicyVersionResponse
}

func (o *RetrieveLocalizedPolicyVersionsOK) Error() string {
	return fmt.Sprintf("[GET /agreement/admin/localized-policy-versions/versions/{policyVersionId}][%d] retrieveLocalizedPolicyVersionsOK  %+v", 200, o.ToString())
}

func (o *RetrieveLocalizedPolicyVersionsOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *RetrieveLocalizedPolicyVersionsOK) GetPayload() []*legalclientmodels.RetrieveLocalizedPolicyVersionResponse {
	return o.Payload
}

func (o *RetrieveLocalizedPolicyVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

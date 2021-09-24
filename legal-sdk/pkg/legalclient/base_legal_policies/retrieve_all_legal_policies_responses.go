// Code generated by go-swagger; DO NOT EDIT.

package base_legal_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
)

// RetrieveAllLegalPoliciesReader is a Reader for the RetrieveAllLegalPolicies structure.
type RetrieveAllLegalPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveAllLegalPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveAllLegalPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/base-policies returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRetrieveAllLegalPoliciesOK creates a RetrieveAllLegalPoliciesOK with default headers values
func NewRetrieveAllLegalPoliciesOK() *RetrieveAllLegalPoliciesOK {
	return &RetrieveAllLegalPoliciesOK{}
}

/*RetrieveAllLegalPoliciesOK handles this case with default header values.

  successful operation
*/
type RetrieveAllLegalPoliciesOK struct {
	Payload []*legalclientmodels.RetrieveBasePolicyResponse
}

func (o *RetrieveAllLegalPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /admin/base-policies][%d] retrieveAllLegalPoliciesOK  %+v", 200, o.Payload)
}

func (o *RetrieveAllLegalPoliciesOK) GetPayload() []*legalclientmodels.RetrieveBasePolicyResponse {
	return o.Payload
}

func (o *RetrieveAllLegalPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

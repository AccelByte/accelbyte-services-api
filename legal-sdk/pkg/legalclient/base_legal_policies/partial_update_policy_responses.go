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

// PartialUpdatePolicyReader is a Reader for the PartialUpdatePolicy structure.
type PartialUpdatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartialUpdatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartialUpdatePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPartialUpdatePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /admin/base-policies/{basePolicyId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPartialUpdatePolicyOK creates a PartialUpdatePolicyOK with default headers values
func NewPartialUpdatePolicyOK() *PartialUpdatePolicyOK {
	return &PartialUpdatePolicyOK{}
}

/*PartialUpdatePolicyOK handles this case with default header values.

  successful operation
*/
type PartialUpdatePolicyOK struct {
	Payload *legalclientmodels.UpdateBasePolicyResponse
}

func (o *PartialUpdatePolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /admin/base-policies/{basePolicyId}][%d] partialUpdatePolicyOK  %+v", 200, o.Payload)
}

func (o *PartialUpdatePolicyOK) GetPayload() *legalclientmodels.UpdateBasePolicyResponse {
	return o.Payload
}

func (o *PartialUpdatePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legalclientmodels.UpdateBasePolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartialUpdatePolicyBadRequest creates a PartialUpdatePolicyBadRequest with default headers values
func NewPartialUpdatePolicyBadRequest() *PartialUpdatePolicyBadRequest {
	return &PartialUpdatePolicyBadRequest{}
}

/*PartialUpdatePolicyBadRequest handles this case with default header values.

  <table><tr><td>NumericErrorCode</td><td>ErrorCode</td></tr><tr><td>40032</td><td>errors.net.accelbyte.platform.legal.invalid_base_policy</td></tr></table>
*/
type PartialUpdatePolicyBadRequest struct {
	Payload *legalclientmodels.ErrorEntity
}

func (o *PartialUpdatePolicyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /admin/base-policies/{basePolicyId}][%d] partialUpdatePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PartialUpdatePolicyBadRequest) GetPayload() *legalclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PartialUpdatePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legalclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

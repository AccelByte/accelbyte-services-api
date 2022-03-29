// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// CreditUserWalletReader is a Reader for the CreditUserWallet structure.
type CreditUserWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreditUserWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreditUserWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreditUserWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreditUserWalletUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/credit returns an error %d: %s", response.Code(), string(data))
	}
}

// NewCreditUserWalletOK creates a CreditUserWalletOK with default headers values
func NewCreditUserWalletOK() *CreditUserWalletOK {
	return &CreditUserWalletOK{}
}

/*CreditUserWalletOK handles this case with default header values.

  successful operation
*/
type CreditUserWalletOK struct {
	Payload *platformclientmodels.WalletInfo
}

func (o *CreditUserWalletOK) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/credit][%d] creditUserWalletOK  %+v", 200, o.Payload)
}

func (o *CreditUserWalletOK) GetPayload() *platformclientmodels.WalletInfo {
	return o.Payload
}

func (o *CreditUserWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.WalletInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreditUserWalletBadRequest creates a CreditUserWalletBadRequest with default headers values
func NewCreditUserWalletBadRequest() *CreditUserWalletBadRequest {
	return &CreditUserWalletBadRequest{}
}

/*CreditUserWalletBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>35123</td><td>Wallet [{walletId}] is inactive</td></tr></table>
*/
type CreditUserWalletBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *CreditUserWalletBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/credit][%d] creditUserWalletBadRequest  %+v", 400, o.Payload)
}

func (o *CreditUserWalletBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *CreditUserWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreditUserWalletUnprocessableEntity creates a CreditUserWalletUnprocessableEntity with default headers values
func NewCreditUserWalletUnprocessableEntity() *CreditUserWalletUnprocessableEntity {
	return &CreditUserWalletUnprocessableEntity{}
}

/*CreditUserWalletUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type CreditUserWalletUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *CreditUserWalletUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/credit][%d] creditUserWalletUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreditUserWalletUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *CreditUserWalletUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

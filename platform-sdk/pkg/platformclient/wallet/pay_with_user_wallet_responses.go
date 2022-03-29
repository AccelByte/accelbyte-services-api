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

// PayWithUserWalletReader is a Reader for the PayWithUserWallet structure.
type PayWithUserWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PayWithUserWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPayWithUserWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPayWithUserWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewPayWithUserWalletUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/payment returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPayWithUserWalletOK creates a PayWithUserWalletOK with default headers values
func NewPayWithUserWalletOK() *PayWithUserWalletOK {
	return &PayWithUserWalletOK{}
}

/*PayWithUserWalletOK handles this case with default header values.

  successful operation
*/
type PayWithUserWalletOK struct {
	Payload *platformclientmodels.WalletInfo
}

func (o *PayWithUserWalletOK) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/payment][%d] payWithUserWalletOK  %+v", 200, o.Payload)
}

func (o *PayWithUserWalletOK) GetPayload() *platformclientmodels.WalletInfo {
	return o.Payload
}

func (o *PayWithUserWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.WalletInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPayWithUserWalletBadRequest creates a PayWithUserWalletBadRequest with default headers values
func NewPayWithUserWalletBadRequest() *PayWithUserWalletBadRequest {
	return &PayWithUserWalletBadRequest{}
}

/*PayWithUserWalletBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>35123</td><td>Wallet [{walletId}] is inactive</td></tr><tr><td>35124</td><td>Wallet [{currencyCode}] has insufficient balance</td></tr></table>
*/
type PayWithUserWalletBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *PayWithUserWalletBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/payment][%d] payWithUserWalletBadRequest  %+v", 400, o.Payload)
}

func (o *PayWithUserWalletBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PayWithUserWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPayWithUserWalletUnprocessableEntity creates a PayWithUserWalletUnprocessableEntity with default headers values
func NewPayWithUserWalletUnprocessableEntity() *PayWithUserWalletUnprocessableEntity {
	return &PayWithUserWalletUnprocessableEntity{}
}

/*PayWithUserWalletUnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type PayWithUserWalletUnprocessableEntity struct {
	Payload *platformclientmodels.ValidationErrorEntity
}

func (o *PayWithUserWalletUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/payment][%d] payWithUserWalletUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PayWithUserWalletUnprocessableEntity) GetPayload() *platformclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *PayWithUserWalletUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

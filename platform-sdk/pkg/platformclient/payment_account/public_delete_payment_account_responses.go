// Code generated by go-swagger; DO NOT EDIT.

package payment_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PublicDeletePaymentAccountReader is a Reader for the PublicDeletePaymentAccount structure.
type PublicDeletePaymentAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicDeletePaymentAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPublicDeletePaymentAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /public/namespaces/{namespace}/users/{userId}/payment/accounts/{type}/{id} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicDeletePaymentAccountNoContent creates a PublicDeletePaymentAccountNoContent with default headers values
func NewPublicDeletePaymentAccountNoContent() *PublicDeletePaymentAccountNoContent {
	return &PublicDeletePaymentAccountNoContent{}
}

/*PublicDeletePaymentAccountNoContent handles this case with default header values.

  delete a payment account successfully
*/
type PublicDeletePaymentAccountNoContent struct {
}

func (o *PublicDeletePaymentAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/namespaces/{namespace}/users/{userId}/payment/accounts/{type}/{id}][%d] publicDeletePaymentAccountNoContent ", 204)
}

func (o *PublicDeletePaymentAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
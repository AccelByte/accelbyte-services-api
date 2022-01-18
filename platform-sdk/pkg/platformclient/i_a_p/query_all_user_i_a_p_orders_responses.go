// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// QueryAllUserIAPOrdersReader is a Reader for the QueryAllUserIAPOrders structure.
type QueryAllUserIAPOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryAllUserIAPOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryAllUserIAPOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/users/{userId}/iap/all returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueryAllUserIAPOrdersOK creates a QueryAllUserIAPOrdersOK with default headers values
func NewQueryAllUserIAPOrdersOK() *QueryAllUserIAPOrdersOK {
	return &QueryAllUserIAPOrdersOK{}
}

/*QueryAllUserIAPOrdersOK handles this case with default header values.

  successful operation
*/
type QueryAllUserIAPOrdersOK struct {
	Payload *platformclientmodels.IAPOrderPagingSlicedResult
}

func (o *QueryAllUserIAPOrdersOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/users/{userId}/iap/all][%d] queryAllUserIAPOrdersOK  %+v", 200, o.Payload)
}

func (o *QueryAllUserIAPOrdersOK) GetPayload() *platformclientmodels.IAPOrderPagingSlicedResult {
	return o.Payload
}

func (o *QueryAllUserIAPOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.IAPOrderPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

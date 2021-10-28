// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// QueryUserSubscriptionsReader is a Reader for the QueryUserSubscriptions structure.
type QueryUserSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryUserSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryUserSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/namespaces/{namespace}/users/{userId}/subscriptions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueryUserSubscriptionsOK creates a QueryUserSubscriptionsOK with default headers values
func NewQueryUserSubscriptionsOK() *QueryUserSubscriptionsOK {
	return &QueryUserSubscriptionsOK{}
}

/*QueryUserSubscriptionsOK handles this case with default header values.

  successful operation
*/
type QueryUserSubscriptionsOK struct {
	Payload *platformclientmodels.SubscriptionPagingSlicedResult
}

func (o *QueryUserSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /admin/namespaces/{namespace}/users/{userId}/subscriptions][%d] queryUserSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *QueryUserSubscriptionsOK) GetPayload() *platformclientmodels.SubscriptionPagingSlicedResult {
	return o.Payload
}

func (o *QueryUserSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.SubscriptionPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

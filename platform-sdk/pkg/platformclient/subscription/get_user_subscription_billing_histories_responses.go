// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package subscription

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// GetUserSubscriptionBillingHistoriesReader is a Reader for the GetUserSubscriptionBillingHistories structure.
type GetUserSubscriptionBillingHistoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSubscriptionBillingHistoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSubscriptionBillingHistoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/history returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserSubscriptionBillingHistoriesOK creates a GetUserSubscriptionBillingHistoriesOK with default headers values
func NewGetUserSubscriptionBillingHistoriesOK() *GetUserSubscriptionBillingHistoriesOK {
	return &GetUserSubscriptionBillingHistoriesOK{}
}

/*GetUserSubscriptionBillingHistoriesOK handles this case with default header values.

  successful operation
*/
type GetUserSubscriptionBillingHistoriesOK struct {
	Payload *platformclientmodels.BillingHistoryPagingSlicedResult
}

func (o *GetUserSubscriptionBillingHistoriesOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/history][%d] getUserSubscriptionBillingHistoriesOK  %+v", 200, o.ToJSONString())
}

func (o *GetUserSubscriptionBillingHistoriesOK) ToJSONString() string {
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

func (o *GetUserSubscriptionBillingHistoriesOK) GetPayload() *platformclientmodels.BillingHistoryPagingSlicedResult {
	return o.Payload
}

func (o *GetUserSubscriptionBillingHistoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(platformclientmodels.BillingHistoryPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

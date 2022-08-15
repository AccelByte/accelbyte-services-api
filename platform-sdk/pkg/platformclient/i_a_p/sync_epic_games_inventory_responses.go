// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// SyncEpicGamesInventoryReader is a Reader for the SyncEpicGamesInventory structure.
type SyncEpicGamesInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncEpicGamesInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncEpicGamesInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/public/namespaces/{namespace}/users/{userId}/iap/epicgames/sync returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSyncEpicGamesInventoryOK creates a SyncEpicGamesInventoryOK with default headers values
func NewSyncEpicGamesInventoryOK() *SyncEpicGamesInventoryOK {
	return &SyncEpicGamesInventoryOK{}
}

/*SyncEpicGamesInventoryOK handles this case with default header values.

  successful operation
*/
type SyncEpicGamesInventoryOK struct {
	Payload []*platformclientmodels.EpicGamesReconcileResult
}

func (o *SyncEpicGamesInventoryOK) Error() string {
	return fmt.Sprintf("[PUT /platform/public/namespaces/{namespace}/users/{userId}/iap/epicgames/sync][%d] syncEpicGamesInventoryOK  %+v", 200, o.ToJSONString())
}

func (o *SyncEpicGamesInventoryOK) ToJSONString() string {
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

func (o *SyncEpicGamesInventoryOK) GetPayload() []*platformclientmodels.EpicGamesReconcileResult {
	return o.Payload
}

func (o *SyncEpicGamesInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

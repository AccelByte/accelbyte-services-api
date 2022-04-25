// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package d_l_c

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SyncXboxDLCReader is a Reader for the SyncXboxDLC structure.
type SyncXboxDLCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncXboxDLCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSyncXboxDLCNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/public/namespaces/{namespace}/users/{userId}/dlc/xbl/sync returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSyncXboxDLCNoContent creates a SyncXboxDLCNoContent with default headers values
func NewSyncXboxDLCNoContent() *SyncXboxDLCNoContent {
	return &SyncXboxDLCNoContent{}
}

/*SyncXboxDLCNoContent handles this case with default header values.

  Successful operation
*/
type SyncXboxDLCNoContent struct {
}

func (o *SyncXboxDLCNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/public/namespaces/{namespace}/users/{userId}/dlc/xbl/sync][%d] syncXboxDLCNoContent ", 204)
}

func (o *SyncXboxDLCNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

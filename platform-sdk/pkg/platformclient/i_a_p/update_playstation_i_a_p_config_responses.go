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

// UpdatePlaystationIAPConfigReader is a Reader for the UpdatePlaystationIAPConfig structure.
type UpdatePlaystationIAPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlaystationIAPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePlaystationIAPConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/admin/namespaces/{namespace}/iap/config/playstation returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdatePlaystationIAPConfigOK creates a UpdatePlaystationIAPConfigOK with default headers values
func NewUpdatePlaystationIAPConfigOK() *UpdatePlaystationIAPConfigOK {
	return &UpdatePlaystationIAPConfigOK{}
}

/*UpdatePlaystationIAPConfigOK handles this case with default header values.

  successful operation
*/
type UpdatePlaystationIAPConfigOK struct {
	Payload *platformclientmodels.PlayStationIAPConfigInfo
}

func (o *UpdatePlaystationIAPConfigOK) Error() string {
	return fmt.Sprintf("[PUT /platform/admin/namespaces/{namespace}/iap/config/playstation][%d] updatePlaystationIAPConfigOK  %+v", 200, o.ToString())
}

func (o *UpdatePlaystationIAPConfigOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *UpdatePlaystationIAPConfigOK) GetPayload() *platformclientmodels.PlayStationIAPConfigInfo {
	return o.Payload
}

func (o *UpdatePlaystationIAPConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.PlayStationIAPConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

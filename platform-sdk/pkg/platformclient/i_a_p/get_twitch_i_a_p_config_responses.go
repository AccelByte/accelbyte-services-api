// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

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

// GetTwitchIAPConfigReader is a Reader for the GetTwitchIAPConfig structure.
type GetTwitchIAPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTwitchIAPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTwitchIAPConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/iap/config/twitch returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetTwitchIAPConfigOK creates a GetTwitchIAPConfigOK with default headers values
func NewGetTwitchIAPConfigOK() *GetTwitchIAPConfigOK {
	return &GetTwitchIAPConfigOK{}
}

/*GetTwitchIAPConfigOK handles this case with default header values.

  successful operation
*/
type GetTwitchIAPConfigOK struct {
	Payload *platformclientmodels.TwitchIAPConfigInfo
}

func (o *GetTwitchIAPConfigOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/iap/config/twitch][%d] getTwitchIAPConfigOK  %+v", 200, o.Payload)
}

func (o *GetTwitchIAPConfigOK) GetPayload() *platformclientmodels.TwitchIAPConfigInfo {
	return o.Payload
}

func (o *GetTwitchIAPConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.TwitchIAPConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

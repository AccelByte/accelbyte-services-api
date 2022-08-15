// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package game_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// PublicGetUserGameProfilesReader is a Reader for the PublicGetUserGameProfiles structure.
type PublicGetUserGameProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetUserGameProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetUserGameProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicGetUserGameProfilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /social/public/namespaces/{namespace}/profiles returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetUserGameProfilesOK creates a PublicGetUserGameProfilesOK with default headers values
func NewPublicGetUserGameProfilesOK() *PublicGetUserGameProfilesOK {
	return &PublicGetUserGameProfilesOK{}
}

/*PublicGetUserGameProfilesOK handles this case with default header values.

  successful operation
*/
type PublicGetUserGameProfilesOK struct {
	Payload []*socialclientmodels.UserGameProfiles
}

func (o *PublicGetUserGameProfilesOK) Error() string {
	return fmt.Sprintf("[GET /social/public/namespaces/{namespace}/profiles][%d] publicGetUserGameProfilesOK  %+v", 200, o.ToString())
}

func (o *PublicGetUserGameProfilesOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicGetUserGameProfilesOK) GetPayload() []*socialclientmodels.UserGameProfiles {
	return o.Payload
}

func (o *PublicGetUserGameProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetUserGameProfilesBadRequest creates a PublicGetUserGameProfilesBadRequest with default headers values
func NewPublicGetUserGameProfilesBadRequest() *PublicGetUserGameProfilesBadRequest {
	return &PublicGetUserGameProfilesBadRequest{}
}

/*PublicGetUserGameProfilesBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12021</td><td>{totalUser} users is requested. Cannot retrieve more than {limitUser} users at once</td></tr></table>
*/
type PublicGetUserGameProfilesBadRequest struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *PublicGetUserGameProfilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /social/public/namespaces/{namespace}/profiles][%d] publicGetUserGameProfilesBadRequest  %+v", 400, o.ToString())
}

func (o *PublicGetUserGameProfilesBadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *PublicGetUserGameProfilesBadRequest) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicGetUserGameProfilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package agreement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
)

// ChangePreferenceConsent1Reader is a Reader for the ChangePreferenceConsent1 structure.
type ChangePreferenceConsent1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePreferenceConsent1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePreferenceConsent1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangePreferenceConsent1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /agreement/public/agreements/localized-policy-versions/preferences returns an error %d: %s", response.Code(), string(data))
	}
}

// NewChangePreferenceConsent1OK creates a ChangePreferenceConsent1OK with default headers values
func NewChangePreferenceConsent1OK() *ChangePreferenceConsent1OK {
	return &ChangePreferenceConsent1OK{}
}

/*ChangePreferenceConsent1OK handles this case with default header values.

  successful operation
*/
type ChangePreferenceConsent1OK struct {
}

func (o *ChangePreferenceConsent1OK) Error() string {
	return fmt.Sprintf("[PATCH /agreement/public/agreements/localized-policy-versions/preferences][%d] changePreferenceConsent1OK ", 200)
}

func (o *ChangePreferenceConsent1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePreferenceConsent1BadRequest creates a ChangePreferenceConsent1BadRequest with default headers values
func NewChangePreferenceConsent1BadRequest() *ChangePreferenceConsent1BadRequest {
	return &ChangePreferenceConsent1BadRequest{}
}

/*ChangePreferenceConsent1BadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>40017</td><td>Policy with id : [{policyId}] is not marketing preference</td></tr></table>
*/
type ChangePreferenceConsent1BadRequest struct {
	Payload *legalclientmodels.ErrorEntity
}

func (o *ChangePreferenceConsent1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /agreement/public/agreements/localized-policy-versions/preferences][%d] changePreferenceConsent1BadRequest  %+v", 400, o.ToString())
}

func (o *ChangePreferenceConsent1BadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ChangePreferenceConsent1BadRequest) GetPayload() *legalclientmodels.ErrorEntity {
	return o.Payload
}

func (o *ChangePreferenceConsent1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legalclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

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

// ResetUserStatItemValue1Reader is a Reader for the ResetUserStatItemValue1 structure.
type ResetUserStatItemValue1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetUserStatItemValue1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetUserStatItemValue1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResetUserStatItemValue1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewResetUserStatItemValue1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value/reset returns an error %d: %s", response.Code(), string(data))
	}
}

// NewResetUserStatItemValue1OK creates a ResetUserStatItemValue1OK with default headers values
func NewResetUserStatItemValue1OK() *ResetUserStatItemValue1OK {
	return &ResetUserStatItemValue1OK{}
}

/*ResetUserStatItemValue1OK handles this case with default header values.

  successful operation
*/
type ResetUserStatItemValue1OK struct {
	Payload *socialclientmodels.StatItemIncResult
}

func (o *ResetUserStatItemValue1OK) Error() string {
	return fmt.Sprintf("[PUT /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value/reset][%d] resetUserStatItemValue1OK  %+v", 200, o.ToString())
}

func (o *ResetUserStatItemValue1OK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ResetUserStatItemValue1OK) GetPayload() *socialclientmodels.StatItemIncResult {
	return o.Payload
}

func (o *ResetUserStatItemValue1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.StatItemIncResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetUserStatItemValue1BadRequest creates a ResetUserStatItemValue1BadRequest with default headers values
func NewResetUserStatItemValue1BadRequest() *ResetUserStatItemValue1BadRequest {
	return &ResetUserStatItemValue1BadRequest{}
}

/*ResetUserStatItemValue1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>12221</td><td>Invalid stat operator, expect [{expected}] but actual [{actual}]</td></tr></table>
*/
type ResetUserStatItemValue1BadRequest struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *ResetUserStatItemValue1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value/reset][%d] resetUserStatItemValue1BadRequest  %+v", 400, o.ToString())
}

func (o *ResetUserStatItemValue1BadRequest) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ResetUserStatItemValue1BadRequest) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *ResetUserStatItemValue1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetUserStatItemValue1NotFound creates a ResetUserStatItemValue1NotFound with default headers values
func NewResetUserStatItemValue1NotFound() *ResetUserStatItemValue1NotFound {
	return &ResetUserStatItemValue1NotFound{}
}

/*ResetUserStatItemValue1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>12241</td><td>Stat [{statCode}] cannot be found in namespace [{namespace}]</td></tr></table>
*/
type ResetUserStatItemValue1NotFound struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *ResetUserStatItemValue1NotFound) Error() string {
	return fmt.Sprintf("[PUT /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value/reset][%d] resetUserStatItemValue1NotFound  %+v", 404, o.ToString())
}

func (o *ResetUserStatItemValue1NotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *ResetUserStatItemValue1NotFound) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *ResetUserStatItemValue1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

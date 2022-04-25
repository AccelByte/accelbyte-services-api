// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// BulkResetUserStatItem1Reader is a Reader for the BulkResetUserStatItem1 structure.
type BulkResetUserStatItem1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkResetUserStatItem1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkResetUserStatItem1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewBulkResetUserStatItem1UnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/value/reset/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBulkResetUserStatItem1OK creates a BulkResetUserStatItem1OK with default headers values
func NewBulkResetUserStatItem1OK() *BulkResetUserStatItem1OK {
	return &BulkResetUserStatItem1OK{}
}

/*BulkResetUserStatItem1OK handles this case with default header values.

  successful operation
*/
type BulkResetUserStatItem1OK struct {
	Payload []*socialclientmodels.BulkStatItemOperationResult
}

func (o *BulkResetUserStatItem1OK) Error() string {
	return fmt.Sprintf("[PUT /social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/value/reset/bulk][%d] bulkResetUserStatItem1OK  %+v", 200, o.Payload)
}

func (o *BulkResetUserStatItem1OK) GetPayload() []*socialclientmodels.BulkStatItemOperationResult {
	return o.Payload
}

func (o *BulkResetUserStatItem1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkResetUserStatItem1UnprocessableEntity creates a BulkResetUserStatItem1UnprocessableEntity with default headers values
func NewBulkResetUserStatItem1UnprocessableEntity() *BulkResetUserStatItem1UnprocessableEntity {
	return &BulkResetUserStatItem1UnprocessableEntity{}
}

/*BulkResetUserStatItem1UnprocessableEntity handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type BulkResetUserStatItem1UnprocessableEntity struct {
	Payload *socialclientmodels.ValidationErrorEntity
}

func (o *BulkResetUserStatItem1UnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/value/reset/bulk][%d] bulkResetUserStatItem1UnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BulkResetUserStatItem1UnprocessableEntity) GetPayload() *socialclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *BulkResetUserStatItem1UnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

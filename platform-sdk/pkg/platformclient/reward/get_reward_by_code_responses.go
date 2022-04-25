// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package reward

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

// GetRewardByCodeReader is a Reader for the GetRewardByCode structure.
type GetRewardByCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRewardByCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRewardByCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRewardByCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/public/namespaces/{namespace}/rewards/byCode returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetRewardByCodeOK creates a GetRewardByCodeOK with default headers values
func NewGetRewardByCodeOK() *GetRewardByCodeOK {
	return &GetRewardByCodeOK{}
}

/*GetRewardByCodeOK handles this case with default header values.

  successful operation
*/
type GetRewardByCodeOK struct {
	Payload *platformclientmodels.RewardInfo
}

func (o *GetRewardByCodeOK) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/rewards/byCode][%d] getRewardByCodeOK  %+v", 200, o.Payload)
}

func (o *GetRewardByCodeOK) GetPayload() *platformclientmodels.RewardInfo {
	return o.Payload
}

func (o *GetRewardByCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.RewardInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRewardByCodeNotFound creates a GetRewardByCodeNotFound with default headers values
func NewGetRewardByCodeNotFound() *GetRewardByCodeNotFound {
	return &GetRewardByCodeNotFound{}
}

/*GetRewardByCodeNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>34043</td><td>Reward with code [{rewardCode}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type GetRewardByCodeNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GetRewardByCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/public/namespaces/{namespace}/rewards/byCode][%d] getRewardByCodeNotFound  %+v", 404, o.Payload)
}

func (o *GetRewardByCodeNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetRewardByCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

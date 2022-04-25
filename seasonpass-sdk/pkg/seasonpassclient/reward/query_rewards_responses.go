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

	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
)

// QueryRewardsReader is a Reader for the QueryRewards structure.
type QueryRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryRewardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewQueryRewardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/rewards returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueryRewardsOK creates a QueryRewardsOK with default headers values
func NewQueryRewardsOK() *QueryRewardsOK {
	return &QueryRewardsOK{}
}

/*QueryRewardsOK handles this case with default header values.

  successful operation
*/
type QueryRewardsOK struct {
	Payload []*seasonpassclientmodels.RewardInfo
}

func (o *QueryRewardsOK) Error() string {
	return fmt.Sprintf("[GET /seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/rewards][%d] queryRewardsOK  %+v", 200, o.Payload)
}

func (o *QueryRewardsOK) GetPayload() []*seasonpassclientmodels.RewardInfo {
	return o.Payload
}

func (o *QueryRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRewardsBadRequest creates a QueryRewardsBadRequest with default headers values
func NewQueryRewardsBadRequest() *QueryRewardsBadRequest {
	return &QueryRewardsBadRequest{}
}

/*QueryRewardsBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20026</td><td>publisher namespace not allowed</td></tr></table>
*/
type QueryRewardsBadRequest struct {
	Payload *seasonpassclientmodels.ErrorEntity
}

func (o *QueryRewardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/rewards][%d] queryRewardsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryRewardsBadRequest) GetPayload() *seasonpassclientmodels.ErrorEntity {
	return o.Payload
}

func (o *QueryRewardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(seasonpassclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRewardsNotFound creates a QueryRewardsNotFound with default headers values
func NewQueryRewardsNotFound() *QueryRewardsNotFound {
	return &QueryRewardsNotFound{}
}

/*QueryRewardsNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>49143</td><td>Season [{seasonId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type QueryRewardsNotFound struct {
	Payload *seasonpassclientmodels.ErrorEntity
}

func (o *QueryRewardsNotFound) Error() string {
	return fmt.Sprintf("[GET /seasonpass/admin/namespaces/{namespace}/seasons/{seasonId}/rewards][%d] queryRewardsNotFound  %+v", 404, o.Payload)
}

func (o *QueryRewardsNotFound) GetPayload() *seasonpassclientmodels.ErrorEntity {
	return o.Payload
}

func (o *QueryRewardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(seasonpassclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package leaderboard_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/leaderboard-sdk/pkg/leaderboardclientmodels"
)

// GetAllTimeLeaderboardRankingPublicV1Reader is a Reader for the GetAllTimeLeaderboardRankingPublicV1 structure.
type GetAllTimeLeaderboardRankingPublicV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTimeLeaderboardRankingPublicV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTimeLeaderboardRankingPublicV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllTimeLeaderboardRankingPublicV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllTimeLeaderboardRankingPublicV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAllTimeLeaderboardRankingPublicV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /leaderboard/v1/public/namespaces/{namespace}/leaderboards/{leaderboardCode}/alltime returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetAllTimeLeaderboardRankingPublicV1OK creates a GetAllTimeLeaderboardRankingPublicV1OK with default headers values
func NewGetAllTimeLeaderboardRankingPublicV1OK() *GetAllTimeLeaderboardRankingPublicV1OK {
	return &GetAllTimeLeaderboardRankingPublicV1OK{}
}

/*GetAllTimeLeaderboardRankingPublicV1OK handles this case with default header values.

  OK
*/
type GetAllTimeLeaderboardRankingPublicV1OK struct {
	Payload *leaderboardclientmodels.ModelsGetLeaderboardRankingResp
}

func (o *GetAllTimeLeaderboardRankingPublicV1OK) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/public/namespaces/{namespace}/leaderboards/{leaderboardCode}/alltime][%d] getAllTimeLeaderboardRankingPublicV1OK  %+v", 200, o.Payload)
}

func (o *GetAllTimeLeaderboardRankingPublicV1OK) GetPayload() *leaderboardclientmodels.ModelsGetLeaderboardRankingResp {
	return o.Payload
}

func (o *GetAllTimeLeaderboardRankingPublicV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ModelsGetLeaderboardRankingResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTimeLeaderboardRankingPublicV1BadRequest creates a GetAllTimeLeaderboardRankingPublicV1BadRequest with default headers values
func NewGetAllTimeLeaderboardRankingPublicV1BadRequest() *GetAllTimeLeaderboardRankingPublicV1BadRequest {
	return &GetAllTimeLeaderboardRankingPublicV1BadRequest{}
}

/*GetAllTimeLeaderboardRankingPublicV1BadRequest handles this case with default header values.

  Bad Request
*/
type GetAllTimeLeaderboardRankingPublicV1BadRequest struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetAllTimeLeaderboardRankingPublicV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/public/namespaces/{namespace}/leaderboards/{leaderboardCode}/alltime][%d] getAllTimeLeaderboardRankingPublicV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetAllTimeLeaderboardRankingPublicV1BadRequest) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetAllTimeLeaderboardRankingPublicV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTimeLeaderboardRankingPublicV1NotFound creates a GetAllTimeLeaderboardRankingPublicV1NotFound with default headers values
func NewGetAllTimeLeaderboardRankingPublicV1NotFound() *GetAllTimeLeaderboardRankingPublicV1NotFound {
	return &GetAllTimeLeaderboardRankingPublicV1NotFound{}
}

/*GetAllTimeLeaderboardRankingPublicV1NotFound handles this case with default header values.

  Not Found
*/
type GetAllTimeLeaderboardRankingPublicV1NotFound struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetAllTimeLeaderboardRankingPublicV1NotFound) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/public/namespaces/{namespace}/leaderboards/{leaderboardCode}/alltime][%d] getAllTimeLeaderboardRankingPublicV1NotFound  %+v", 404, o.Payload)
}

func (o *GetAllTimeLeaderboardRankingPublicV1NotFound) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetAllTimeLeaderboardRankingPublicV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTimeLeaderboardRankingPublicV1InternalServerError creates a GetAllTimeLeaderboardRankingPublicV1InternalServerError with default headers values
func NewGetAllTimeLeaderboardRankingPublicV1InternalServerError() *GetAllTimeLeaderboardRankingPublicV1InternalServerError {
	return &GetAllTimeLeaderboardRankingPublicV1InternalServerError{}
}

/*GetAllTimeLeaderboardRankingPublicV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetAllTimeLeaderboardRankingPublicV1InternalServerError struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetAllTimeLeaderboardRankingPublicV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/public/namespaces/{namespace}/leaderboards/{leaderboardCode}/alltime][%d] getAllTimeLeaderboardRankingPublicV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllTimeLeaderboardRankingPublicV1InternalServerError) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetAllTimeLeaderboardRankingPublicV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

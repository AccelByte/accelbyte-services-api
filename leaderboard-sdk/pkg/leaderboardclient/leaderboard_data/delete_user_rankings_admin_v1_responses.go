// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package leaderboard_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/leaderboard-sdk/pkg/leaderboardclientmodels"
)

// DeleteUserRankingsAdminV1Reader is a Reader for the DeleteUserRankingsAdminV1 structure.
type DeleteUserRankingsAdminV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserRankingsAdminV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserRankingsAdminV1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserRankingsAdminV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteUserRankingsAdminV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteUserRankingsAdminV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteUserRankingsAdminV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /leaderboard/v1/admin/namespaces/{namespace}/users/{userId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteUserRankingsAdminV1NoContent creates a DeleteUserRankingsAdminV1NoContent with default headers values
func NewDeleteUserRankingsAdminV1NoContent() *DeleteUserRankingsAdminV1NoContent {
	return &DeleteUserRankingsAdminV1NoContent{}
}

/*DeleteUserRankingsAdminV1NoContent handles this case with default header values.

  No Content
*/
type DeleteUserRankingsAdminV1NoContent struct {
}

func (o *DeleteUserRankingsAdminV1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /leaderboard/v1/admin/namespaces/{namespace}/users/{userId}][%d] deleteUserRankingsAdminV1NoContent ", 204)
}

func (o *DeleteUserRankingsAdminV1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserRankingsAdminV1Unauthorized creates a DeleteUserRankingsAdminV1Unauthorized with default headers values
func NewDeleteUserRankingsAdminV1Unauthorized() *DeleteUserRankingsAdminV1Unauthorized {
	return &DeleteUserRankingsAdminV1Unauthorized{}
}

/*DeleteUserRankingsAdminV1Unauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteUserRankingsAdminV1Unauthorized struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *DeleteUserRankingsAdminV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /leaderboard/v1/admin/namespaces/{namespace}/users/{userId}][%d] deleteUserRankingsAdminV1Unauthorized  %+v", 401, o.ToString())
}

func (o *DeleteUserRankingsAdminV1Unauthorized) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteUserRankingsAdminV1Unauthorized) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteUserRankingsAdminV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRankingsAdminV1Forbidden creates a DeleteUserRankingsAdminV1Forbidden with default headers values
func NewDeleteUserRankingsAdminV1Forbidden() *DeleteUserRankingsAdminV1Forbidden {
	return &DeleteUserRankingsAdminV1Forbidden{}
}

/*DeleteUserRankingsAdminV1Forbidden handles this case with default header values.

  Forbidden
*/
type DeleteUserRankingsAdminV1Forbidden struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *DeleteUserRankingsAdminV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /leaderboard/v1/admin/namespaces/{namespace}/users/{userId}][%d] deleteUserRankingsAdminV1Forbidden  %+v", 403, o.ToString())
}

func (o *DeleteUserRankingsAdminV1Forbidden) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteUserRankingsAdminV1Forbidden) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteUserRankingsAdminV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRankingsAdminV1NotFound creates a DeleteUserRankingsAdminV1NotFound with default headers values
func NewDeleteUserRankingsAdminV1NotFound() *DeleteUserRankingsAdminV1NotFound {
	return &DeleteUserRankingsAdminV1NotFound{}
}

/*DeleteUserRankingsAdminV1NotFound handles this case with default header values.

  Not Found
*/
type DeleteUserRankingsAdminV1NotFound struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *DeleteUserRankingsAdminV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /leaderboard/v1/admin/namespaces/{namespace}/users/{userId}][%d] deleteUserRankingsAdminV1NotFound  %+v", 404, o.ToString())
}

func (o *DeleteUserRankingsAdminV1NotFound) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteUserRankingsAdminV1NotFound) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteUserRankingsAdminV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRankingsAdminV1InternalServerError creates a DeleteUserRankingsAdminV1InternalServerError with default headers values
func NewDeleteUserRankingsAdminV1InternalServerError() *DeleteUserRankingsAdminV1InternalServerError {
	return &DeleteUserRankingsAdminV1InternalServerError{}
}

/*DeleteUserRankingsAdminV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteUserRankingsAdminV1InternalServerError struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *DeleteUserRankingsAdminV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /leaderboard/v1/admin/namespaces/{namespace}/users/{userId}][%d] deleteUserRankingsAdminV1InternalServerError  %+v", 500, o.ToString())
}

func (o *DeleteUserRankingsAdminV1InternalServerError) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteUserRankingsAdminV1InternalServerError) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *DeleteUserRankingsAdminV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

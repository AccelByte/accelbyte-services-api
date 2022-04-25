// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package leaderboard_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetTodayLeaderboardRankingPublicV1Params creates a new GetTodayLeaderboardRankingPublicV1Params object
// with the default values initialized.
func NewGetTodayLeaderboardRankingPublicV1Params() *GetTodayLeaderboardRankingPublicV1Params {
	var ()
	return &GetTodayLeaderboardRankingPublicV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTodayLeaderboardRankingPublicV1ParamsWithTimeout creates a new GetTodayLeaderboardRankingPublicV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTodayLeaderboardRankingPublicV1ParamsWithTimeout(timeout time.Duration) *GetTodayLeaderboardRankingPublicV1Params {
	var ()
	return &GetTodayLeaderboardRankingPublicV1Params{

		timeout: timeout,
	}
}

// NewGetTodayLeaderboardRankingPublicV1ParamsWithContext creates a new GetTodayLeaderboardRankingPublicV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetTodayLeaderboardRankingPublicV1ParamsWithContext(ctx context.Context) *GetTodayLeaderboardRankingPublicV1Params {
	var ()
	return &GetTodayLeaderboardRankingPublicV1Params{

		Context: ctx,
	}
}

// NewGetTodayLeaderboardRankingPublicV1ParamsWithHTTPClient creates a new GetTodayLeaderboardRankingPublicV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTodayLeaderboardRankingPublicV1ParamsWithHTTPClient(client *http.Client) *GetTodayLeaderboardRankingPublicV1Params {
	var ()
	return &GetTodayLeaderboardRankingPublicV1Params{
		HTTPClient: client,
	}
}

/*GetTodayLeaderboardRankingPublicV1Params contains all the parameters to send to the API endpoint
for the get today leaderboard ranking public v1 operation typically these are written to a http.Request
*/
type GetTodayLeaderboardRankingPublicV1Params struct {

	/*LeaderboardCode
	  the human readable unique code to identify the leaderboard. Must be lowercase and maximum length is 48

	*/
	LeaderboardCode string
	/*Limit
	  size of displayed data

	*/
	Limit *int64
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Offset
	  The start position that points to query data

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithTimeout(timeout time.Duration) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithContext(ctx context.Context) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithHTTPClient(client *http.Client) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLeaderboardCode adds the leaderboardCode to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithLeaderboardCode(leaderboardCode string) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetLeaderboardCode(leaderboardCode)
	return o
}

// SetLeaderboardCode adds the leaderboardCode to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetLeaderboardCode(leaderboardCode string) {
	o.LeaderboardCode = leaderboardCode
}

// WithLimit adds the limit to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithLimit(limit *int64) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithNamespace(namespace string) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) WithOffset(offset *int64) *GetTodayLeaderboardRankingPublicV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get today leaderboard ranking public v1 params
func (o *GetTodayLeaderboardRankingPublicV1Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetTodayLeaderboardRankingPublicV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param leaderboardCode
	if err := r.SetPathParam("leaderboardCode", o.LeaderboardCode); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

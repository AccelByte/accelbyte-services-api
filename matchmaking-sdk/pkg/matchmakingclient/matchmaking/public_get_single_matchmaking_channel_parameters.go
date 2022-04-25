// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package matchmaking

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
)

// NewPublicGetSingleMatchmakingChannelParams creates a new PublicGetSingleMatchmakingChannelParams object
// with the default values initialized.
func NewPublicGetSingleMatchmakingChannelParams() *PublicGetSingleMatchmakingChannelParams {
	var ()
	return &PublicGetSingleMatchmakingChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetSingleMatchmakingChannelParamsWithTimeout creates a new PublicGetSingleMatchmakingChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetSingleMatchmakingChannelParamsWithTimeout(timeout time.Duration) *PublicGetSingleMatchmakingChannelParams {
	var ()
	return &PublicGetSingleMatchmakingChannelParams{

		timeout: timeout,
	}
}

// NewPublicGetSingleMatchmakingChannelParamsWithContext creates a new PublicGetSingleMatchmakingChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetSingleMatchmakingChannelParamsWithContext(ctx context.Context) *PublicGetSingleMatchmakingChannelParams {
	var ()
	return &PublicGetSingleMatchmakingChannelParams{

		Context: ctx,
	}
}

// NewPublicGetSingleMatchmakingChannelParamsWithHTTPClient creates a new PublicGetSingleMatchmakingChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetSingleMatchmakingChannelParamsWithHTTPClient(client *http.Client) *PublicGetSingleMatchmakingChannelParams {
	var ()
	return &PublicGetSingleMatchmakingChannelParams{
		HTTPClient: client,
	}
}

/*PublicGetSingleMatchmakingChannelParams contains all the parameters to send to the API endpoint
for the public get single matchmaking channel operation typically these are written to a http.Request
*/
type PublicGetSingleMatchmakingChannelParams struct {

	/*ChannelName
	  channel name, accept snake_case, lowercase, and numeric

	*/
	ChannelName string
	/*Namespace
	  namespace of the game, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) WithTimeout(timeout time.Duration) *PublicGetSingleMatchmakingChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) WithContext(ctx context.Context) *PublicGetSingleMatchmakingChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) WithHTTPClient(client *http.Client) *PublicGetSingleMatchmakingChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelName adds the channelName to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) WithChannelName(channelName string) *PublicGetSingleMatchmakingChannelParams {
	o.SetChannelName(channelName)
	return o
}

// SetChannelName adds the channelName to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) SetChannelName(channelName string) {
	o.ChannelName = channelName
}

// WithNamespace adds the namespace to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) WithNamespace(namespace string) *PublicGetSingleMatchmakingChannelParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get single matchmaking channel params
func (o *PublicGetSingleMatchmakingChannelParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetSingleMatchmakingChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channelName
	if err := r.SetPathParam("channelName", o.ChannelName); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

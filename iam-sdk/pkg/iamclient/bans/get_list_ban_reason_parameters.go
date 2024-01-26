// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package bans

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetListBanReasonParams creates a new GetListBanReasonParams object
// with the default values initialized.
func NewGetListBanReasonParams() *GetListBanReasonParams {
	var ()
	return &GetListBanReasonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetListBanReasonParamsWithTimeout creates a new GetListBanReasonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetListBanReasonParamsWithTimeout(timeout time.Duration) *GetListBanReasonParams {
	var ()
	return &GetListBanReasonParams{

		timeout: timeout,
	}
}

// NewGetListBanReasonParamsWithContext creates a new GetListBanReasonParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetListBanReasonParamsWithContext(ctx context.Context) *GetListBanReasonParams {
	var ()
	return &GetListBanReasonParams{

		Context: ctx,
	}
}

// NewGetListBanReasonParamsWithHTTPClient creates a new GetListBanReasonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetListBanReasonParamsWithHTTPClient(client *http.Client) *GetListBanReasonParams {
	var ()
	return &GetListBanReasonParams{
		HTTPClient: client,
	}
}

/*GetListBanReasonParams contains all the parameters to send to the API endpoint
for the get list ban reason operation typically these are written to a http.Request
*/
type GetListBanReasonParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the get list ban reason params
func (o *GetListBanReasonParams) WithTimeout(timeout time.Duration) *GetListBanReasonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get list ban reason params
func (o *GetListBanReasonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get list ban reason params
func (o *GetListBanReasonParams) WithContext(ctx context.Context) *GetListBanReasonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get list ban reason params
func (o *GetListBanReasonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the get list ban reason params
func (o *GetListBanReasonParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the get list ban reason params
func (o *GetListBanReasonParams) WithHTTPClient(client *http.Client) *GetListBanReasonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get list ban reason params
func (o *GetListBanReasonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the get list ban reason params
func (o *GetListBanReasonParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *GetListBanReasonParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WriteToRequest writes these params to a swagger request
func (o *GetListBanReasonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// setting the default header value
	if err := r.SetHeaderParam("User-Agent", utils.UserAgentGen()); err != nil {
		return err
	}

	if err := r.SetHeaderParam("X-Amzn-Trace-Id", utils.AmazonTraceIDGen()); err != nil {
		return err
	}

	if o.XFlightId == nil {
		if err := r.SetHeaderParam("X-Flight-Id", utils.GetDefaultFlightID().Value); err != nil {
			return err
		}
	} else {
		if err := r.SetHeaderParam("X-Flight-Id", *o.XFlightId); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

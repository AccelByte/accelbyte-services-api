// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package fulfillment

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewFulfillRewardsParams creates a new FulfillRewardsParams object
// with the default values initialized.
func NewFulfillRewardsParams() *FulfillRewardsParams {
	var ()
	return &FulfillRewardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFulfillRewardsParamsWithTimeout creates a new FulfillRewardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFulfillRewardsParamsWithTimeout(timeout time.Duration) *FulfillRewardsParams {
	var ()
	return &FulfillRewardsParams{

		timeout: timeout,
	}
}

// NewFulfillRewardsParamsWithContext creates a new FulfillRewardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFulfillRewardsParamsWithContext(ctx context.Context) *FulfillRewardsParams {
	var ()
	return &FulfillRewardsParams{

		Context: ctx,
	}
}

// NewFulfillRewardsParamsWithHTTPClient creates a new FulfillRewardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFulfillRewardsParamsWithHTTPClient(client *http.Client) *FulfillRewardsParams {
	var ()
	return &FulfillRewardsParams{
		HTTPClient: client,
	}
}

/*FulfillRewardsParams contains all the parameters to send to the API endpoint
for the fulfill rewards operation typically these are written to a http.Request
*/
type FulfillRewardsParams struct {

	/*Body*/
	Body *platformclientmodels.RewardsRequest
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the fulfill rewards params
func (o *FulfillRewardsParams) WithTimeout(timeout time.Duration) *FulfillRewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fulfill rewards params
func (o *FulfillRewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fulfill rewards params
func (o *FulfillRewardsParams) WithContext(ctx context.Context) *FulfillRewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fulfill rewards params
func (o *FulfillRewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fulfill rewards params
func (o *FulfillRewardsParams) WithHTTPClient(client *http.Client) *FulfillRewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fulfill rewards params
func (o *FulfillRewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the fulfill rewards params
func (o *FulfillRewardsParams) WithBody(body *platformclientmodels.RewardsRequest) *FulfillRewardsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the fulfill rewards params
func (o *FulfillRewardsParams) SetBody(body *platformclientmodels.RewardsRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the fulfill rewards params
func (o *FulfillRewardsParams) WithNamespace(namespace string) *FulfillRewardsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the fulfill rewards params
func (o *FulfillRewardsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the fulfill rewards params
func (o *FulfillRewardsParams) WithUserID(userID string) *FulfillRewardsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the fulfill rewards params
func (o *FulfillRewardsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *FulfillRewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

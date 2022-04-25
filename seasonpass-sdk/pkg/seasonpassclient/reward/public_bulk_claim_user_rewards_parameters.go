// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package reward

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

// NewPublicBulkClaimUserRewardsParams creates a new PublicBulkClaimUserRewardsParams object
// with the default values initialized.
func NewPublicBulkClaimUserRewardsParams() *PublicBulkClaimUserRewardsParams {
	var ()
	return &PublicBulkClaimUserRewardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicBulkClaimUserRewardsParamsWithTimeout creates a new PublicBulkClaimUserRewardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicBulkClaimUserRewardsParamsWithTimeout(timeout time.Duration) *PublicBulkClaimUserRewardsParams {
	var ()
	return &PublicBulkClaimUserRewardsParams{

		timeout: timeout,
	}
}

// NewPublicBulkClaimUserRewardsParamsWithContext creates a new PublicBulkClaimUserRewardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicBulkClaimUserRewardsParamsWithContext(ctx context.Context) *PublicBulkClaimUserRewardsParams {
	var ()
	return &PublicBulkClaimUserRewardsParams{

		Context: ctx,
	}
}

// NewPublicBulkClaimUserRewardsParamsWithHTTPClient creates a new PublicBulkClaimUserRewardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicBulkClaimUserRewardsParamsWithHTTPClient(client *http.Client) *PublicBulkClaimUserRewardsParams {
	var ()
	return &PublicBulkClaimUserRewardsParams{
		HTTPClient: client,
	}
}

/*PublicBulkClaimUserRewardsParams contains all the parameters to send to the API endpoint
for the public bulk claim user rewards operation typically these are written to a http.Request
*/
type PublicBulkClaimUserRewardsParams struct {

	/*Namespace
	  Namespace

	*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) WithTimeout(timeout time.Duration) *PublicBulkClaimUserRewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) WithContext(ctx context.Context) *PublicBulkClaimUserRewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) WithHTTPClient(client *http.Client) *PublicBulkClaimUserRewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) WithNamespace(namespace string) *PublicBulkClaimUserRewardsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) WithUserID(userID string) *PublicBulkClaimUserRewardsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public bulk claim user rewards params
func (o *PublicBulkClaimUserRewardsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicBulkClaimUserRewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

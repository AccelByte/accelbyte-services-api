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

// NewDeleteRewardParams creates a new DeleteRewardParams object
// with the default values initialized.
func NewDeleteRewardParams() *DeleteRewardParams {
	var ()
	return &DeleteRewardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRewardParamsWithTimeout creates a new DeleteRewardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRewardParamsWithTimeout(timeout time.Duration) *DeleteRewardParams {
	var ()
	return &DeleteRewardParams{

		timeout: timeout,
	}
}

// NewDeleteRewardParamsWithContext creates a new DeleteRewardParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRewardParamsWithContext(ctx context.Context) *DeleteRewardParams {
	var ()
	return &DeleteRewardParams{

		Context: ctx,
	}
}

// NewDeleteRewardParamsWithHTTPClient creates a new DeleteRewardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRewardParamsWithHTTPClient(client *http.Client) *DeleteRewardParams {
	var ()
	return &DeleteRewardParams{
		HTTPClient: client,
	}
}

/*DeleteRewardParams contains all the parameters to send to the API endpoint
for the delete reward operation typically these are written to a http.Request
*/
type DeleteRewardParams struct {

	/*Namespace
	  Namespace

	*/
	Namespace string
	/*RewardID*/
	RewardID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete reward params
func (o *DeleteRewardParams) WithTimeout(timeout time.Duration) *DeleteRewardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete reward params
func (o *DeleteRewardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete reward params
func (o *DeleteRewardParams) WithContext(ctx context.Context) *DeleteRewardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete reward params
func (o *DeleteRewardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete reward params
func (o *DeleteRewardParams) WithHTTPClient(client *http.Client) *DeleteRewardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete reward params
func (o *DeleteRewardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete reward params
func (o *DeleteRewardParams) WithNamespace(namespace string) *DeleteRewardParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete reward params
func (o *DeleteRewardParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRewardID adds the rewardID to the delete reward params
func (o *DeleteRewardParams) WithRewardID(rewardID string) *DeleteRewardParams {
	o.SetRewardID(rewardID)
	return o
}

// SetRewardID adds the rewardId to the delete reward params
func (o *DeleteRewardParams) SetRewardID(rewardID string) {
	o.RewardID = rewardID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRewardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param rewardId
	if err := r.SetPathParam("rewardId", o.RewardID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

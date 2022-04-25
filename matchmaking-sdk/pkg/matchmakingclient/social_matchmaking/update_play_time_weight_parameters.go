// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package social_matchmaking

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

	"github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclientmodels"
)

// NewUpdatePlayTimeWeightParams creates a new UpdatePlayTimeWeightParams object
// with the default values initialized.
func NewUpdatePlayTimeWeightParams() *UpdatePlayTimeWeightParams {
	var ()
	return &UpdatePlayTimeWeightParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePlayTimeWeightParamsWithTimeout creates a new UpdatePlayTimeWeightParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePlayTimeWeightParamsWithTimeout(timeout time.Duration) *UpdatePlayTimeWeightParams {
	var ()
	return &UpdatePlayTimeWeightParams{

		timeout: timeout,
	}
}

// NewUpdatePlayTimeWeightParamsWithContext creates a new UpdatePlayTimeWeightParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePlayTimeWeightParamsWithContext(ctx context.Context) *UpdatePlayTimeWeightParams {
	var ()
	return &UpdatePlayTimeWeightParams{

		Context: ctx,
	}
}

// NewUpdatePlayTimeWeightParamsWithHTTPClient creates a new UpdatePlayTimeWeightParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePlayTimeWeightParamsWithHTTPClient(client *http.Client) *UpdatePlayTimeWeightParams {
	var ()
	return &UpdatePlayTimeWeightParams{
		HTTPClient: client,
	}
}

/*UpdatePlayTimeWeightParams contains all the parameters to send to the API endpoint
for the update play time weight operation typically these are written to a http.Request
*/
type UpdatePlayTimeWeightParams struct {

	/*Body*/
	Body *matchmakingclientmodels.ModelsUpdatePlayTimeWeightRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update play time weight params
func (o *UpdatePlayTimeWeightParams) WithTimeout(timeout time.Duration) *UpdatePlayTimeWeightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update play time weight params
func (o *UpdatePlayTimeWeightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update play time weight params
func (o *UpdatePlayTimeWeightParams) WithContext(ctx context.Context) *UpdatePlayTimeWeightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update play time weight params
func (o *UpdatePlayTimeWeightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update play time weight params
func (o *UpdatePlayTimeWeightParams) WithHTTPClient(client *http.Client) *UpdatePlayTimeWeightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update play time weight params
func (o *UpdatePlayTimeWeightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update play time weight params
func (o *UpdatePlayTimeWeightParams) WithBody(body *matchmakingclientmodels.ModelsUpdatePlayTimeWeightRequest) *UpdatePlayTimeWeightParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update play time weight params
func (o *UpdatePlayTimeWeightParams) SetBody(body *matchmakingclientmodels.ModelsUpdatePlayTimeWeightRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the update play time weight params
func (o *UpdatePlayTimeWeightParams) WithNamespace(namespace string) *UpdatePlayTimeWeightParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update play time weight params
func (o *UpdatePlayTimeWeightParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePlayTimeWeightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

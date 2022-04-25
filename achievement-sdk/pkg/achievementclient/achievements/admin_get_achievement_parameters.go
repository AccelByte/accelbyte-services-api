// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package achievements

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

// NewAdminGetAchievementParams creates a new AdminGetAchievementParams object
// with the default values initialized.
func NewAdminGetAchievementParams() *AdminGetAchievementParams {
	var ()
	return &AdminGetAchievementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetAchievementParamsWithTimeout creates a new AdminGetAchievementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetAchievementParamsWithTimeout(timeout time.Duration) *AdminGetAchievementParams {
	var ()
	return &AdminGetAchievementParams{

		timeout: timeout,
	}
}

// NewAdminGetAchievementParamsWithContext creates a new AdminGetAchievementParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetAchievementParamsWithContext(ctx context.Context) *AdminGetAchievementParams {
	var ()
	return &AdminGetAchievementParams{

		Context: ctx,
	}
}

// NewAdminGetAchievementParamsWithHTTPClient creates a new AdminGetAchievementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetAchievementParamsWithHTTPClient(client *http.Client) *AdminGetAchievementParams {
	var ()
	return &AdminGetAchievementParams{
		HTTPClient: client,
	}
}

/*AdminGetAchievementParams contains all the parameters to send to the API endpoint
for the admin get achievement operation typically these are written to a http.Request
*/
type AdminGetAchievementParams struct {

	/*AchievementCode
	  achievement code

	*/
	AchievementCode string
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get achievement params
func (o *AdminGetAchievementParams) WithTimeout(timeout time.Duration) *AdminGetAchievementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get achievement params
func (o *AdminGetAchievementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get achievement params
func (o *AdminGetAchievementParams) WithContext(ctx context.Context) *AdminGetAchievementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get achievement params
func (o *AdminGetAchievementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get achievement params
func (o *AdminGetAchievementParams) WithHTTPClient(client *http.Client) *AdminGetAchievementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get achievement params
func (o *AdminGetAchievementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAchievementCode adds the achievementCode to the admin get achievement params
func (o *AdminGetAchievementParams) WithAchievementCode(achievementCode string) *AdminGetAchievementParams {
	o.SetAchievementCode(achievementCode)
	return o
}

// SetAchievementCode adds the achievementCode to the admin get achievement params
func (o *AdminGetAchievementParams) SetAchievementCode(achievementCode string) {
	o.AchievementCode = achievementCode
}

// WithNamespace adds the namespace to the admin get achievement params
func (o *AdminGetAchievementParams) WithNamespace(namespace string) *AdminGetAchievementParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get achievement params
func (o *AdminGetAchievementParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetAchievementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param achievementCode
	if err := r.SetPathParam("achievementCode", o.AchievementCode); err != nil {
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

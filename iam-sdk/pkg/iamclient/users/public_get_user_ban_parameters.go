// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPublicGetUserBanParams creates a new PublicGetUserBanParams object
// with the default values initialized.
func NewPublicGetUserBanParams() *PublicGetUserBanParams {
	var ()
	return &PublicGetUserBanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserBanParamsWithTimeout creates a new PublicGetUserBanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserBanParamsWithTimeout(timeout time.Duration) *PublicGetUserBanParams {
	var ()
	return &PublicGetUserBanParams{

		timeout: timeout,
	}
}

// NewPublicGetUserBanParamsWithContext creates a new PublicGetUserBanParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserBanParamsWithContext(ctx context.Context) *PublicGetUserBanParams {
	var ()
	return &PublicGetUserBanParams{

		Context: ctx,
	}
}

// NewPublicGetUserBanParamsWithHTTPClient creates a new PublicGetUserBanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserBanParamsWithHTTPClient(client *http.Client) *PublicGetUserBanParams {
	var ()
	return &PublicGetUserBanParams{
		HTTPClient: client,
	}
}

/*PublicGetUserBanParams contains all the parameters to send to the API endpoint
for the public get user ban operation typically these are written to a http.Request
*/
type PublicGetUserBanParams struct {

	/*ActiveOnly
	  Filter ban to only returns the active one

	*/
	ActiveOnly *bool
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get user ban params
func (o *PublicGetUserBanParams) WithTimeout(timeout time.Duration) *PublicGetUserBanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user ban params
func (o *PublicGetUserBanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user ban params
func (o *PublicGetUserBanParams) WithContext(ctx context.Context) *PublicGetUserBanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user ban params
func (o *PublicGetUserBanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user ban params
func (o *PublicGetUserBanParams) WithHTTPClient(client *http.Client) *PublicGetUserBanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user ban params
func (o *PublicGetUserBanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActiveOnly adds the activeOnly to the public get user ban params
func (o *PublicGetUserBanParams) WithActiveOnly(activeOnly *bool) *PublicGetUserBanParams {
	o.SetActiveOnly(activeOnly)
	return o
}

// SetActiveOnly adds the activeOnly to the public get user ban params
func (o *PublicGetUserBanParams) SetActiveOnly(activeOnly *bool) {
	o.ActiveOnly = activeOnly
}

// WithNamespace adds the namespace to the public get user ban params
func (o *PublicGetUserBanParams) WithNamespace(namespace string) *PublicGetUserBanParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user ban params
func (o *PublicGetUserBanParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user ban params
func (o *PublicGetUserBanParams) WithUserID(userID string) *PublicGetUserBanParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user ban params
func (o *PublicGetUserBanParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserBanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActiveOnly != nil {

		// query param activeOnly
		var qrActiveOnly bool
		if o.ActiveOnly != nil {
			qrActiveOnly = *o.ActiveOnly
		}
		qActiveOnly := swag.FormatBool(qrActiveOnly)
		if qActiveOnly != "" {
			if err := r.SetQueryParam("activeOnly", qActiveOnly); err != nil {
				return err
			}
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

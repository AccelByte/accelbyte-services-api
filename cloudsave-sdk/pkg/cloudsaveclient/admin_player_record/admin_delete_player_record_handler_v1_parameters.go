// Code generated by go-swagger; DO NOT EDIT.

package admin_player_record

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

// NewAdminDeletePlayerRecordHandlerV1Params creates a new AdminDeletePlayerRecordHandlerV1Params object
// with the default values initialized.
func NewAdminDeletePlayerRecordHandlerV1Params() *AdminDeletePlayerRecordHandlerV1Params {
	var ()
	return &AdminDeletePlayerRecordHandlerV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminDeletePlayerRecordHandlerV1ParamsWithTimeout creates a new AdminDeletePlayerRecordHandlerV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminDeletePlayerRecordHandlerV1ParamsWithTimeout(timeout time.Duration) *AdminDeletePlayerRecordHandlerV1Params {
	var ()
	return &AdminDeletePlayerRecordHandlerV1Params{

		timeout: timeout,
	}
}

// NewAdminDeletePlayerRecordHandlerV1ParamsWithContext creates a new AdminDeletePlayerRecordHandlerV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminDeletePlayerRecordHandlerV1ParamsWithContext(ctx context.Context) *AdminDeletePlayerRecordHandlerV1Params {
	var ()
	return &AdminDeletePlayerRecordHandlerV1Params{

		Context: ctx,
	}
}

// NewAdminDeletePlayerRecordHandlerV1ParamsWithHTTPClient creates a new AdminDeletePlayerRecordHandlerV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminDeletePlayerRecordHandlerV1ParamsWithHTTPClient(client *http.Client) *AdminDeletePlayerRecordHandlerV1Params {
	var ()
	return &AdminDeletePlayerRecordHandlerV1Params{
		HTTPClient: client,
	}
}

/*AdminDeletePlayerRecordHandlerV1Params contains all the parameters to send to the API endpoint
for the admin delete player record handler v1 operation typically these are written to a http.Request
*/
type AdminDeletePlayerRecordHandlerV1Params struct {

	/*Key
	  key of record

	*/
	Key string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*UserID
	  user ID who own the record

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) WithTimeout(timeout time.Duration) *AdminDeletePlayerRecordHandlerV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) WithContext(ctx context.Context) *AdminDeletePlayerRecordHandlerV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) WithHTTPClient(client *http.Client) *AdminDeletePlayerRecordHandlerV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) WithKey(key string) *AdminDeletePlayerRecordHandlerV1Params {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) SetKey(key string) {
	o.Key = key
}

// WithNamespace adds the namespace to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) WithNamespace(namespace string) *AdminDeletePlayerRecordHandlerV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) WithUserID(userID string) *AdminDeletePlayerRecordHandlerV1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin delete player record handler v1 params
func (o *AdminDeletePlayerRecordHandlerV1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminDeletePlayerRecordHandlerV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userID
	if err := r.SetPathParam("userID", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
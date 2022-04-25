// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package friends

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

// NewGetUserFriendsUpdatedParams creates a new GetUserFriendsUpdatedParams object
// with the default values initialized.
func NewGetUserFriendsUpdatedParams() *GetUserFriendsUpdatedParams {
	var ()
	return &GetUserFriendsUpdatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserFriendsUpdatedParamsWithTimeout creates a new GetUserFriendsUpdatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserFriendsUpdatedParamsWithTimeout(timeout time.Duration) *GetUserFriendsUpdatedParams {
	var ()
	return &GetUserFriendsUpdatedParams{

		timeout: timeout,
	}
}

// NewGetUserFriendsUpdatedParamsWithContext creates a new GetUserFriendsUpdatedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserFriendsUpdatedParamsWithContext(ctx context.Context) *GetUserFriendsUpdatedParams {
	var ()
	return &GetUserFriendsUpdatedParams{

		Context: ctx,
	}
}

// NewGetUserFriendsUpdatedParamsWithHTTPClient creates a new GetUserFriendsUpdatedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserFriendsUpdatedParamsWithHTTPClient(client *http.Client) *GetUserFriendsUpdatedParams {
	var ()
	return &GetUserFriendsUpdatedParams{
		HTTPClient: client,
	}
}

/*GetUserFriendsUpdatedParams contains all the parameters to send to the API endpoint
for the get user friends updated operation typically these are written to a http.Request
*/
type GetUserFriendsUpdatedParams struct {

	/*Limit
	  maximum number of data

	*/
	Limit *int64
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Offset
	  numbers of row to skip within the result

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) WithTimeout(timeout time.Duration) *GetUserFriendsUpdatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) WithContext(ctx context.Context) *GetUserFriendsUpdatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) WithHTTPClient(client *http.Client) *GetUserFriendsUpdatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) WithLimit(limit *int64) *GetUserFriendsUpdatedParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) WithNamespace(namespace string) *GetUserFriendsUpdatedParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) WithOffset(offset *int64) *GetUserFriendsUpdatedParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get user friends updated params
func (o *GetUserFriendsUpdatedParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserFriendsUpdatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

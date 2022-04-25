// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package public_follow

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

// NewGetFollowedContentParams creates a new GetFollowedContentParams object
// with the default values initialized.
func NewGetFollowedContentParams() *GetFollowedContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &GetFollowedContentParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFollowedContentParamsWithTimeout creates a new GetFollowedContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFollowedContentParamsWithTimeout(timeout time.Duration) *GetFollowedContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &GetFollowedContentParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetFollowedContentParamsWithContext creates a new GetFollowedContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFollowedContentParamsWithContext(ctx context.Context) *GetFollowedContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &GetFollowedContentParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetFollowedContentParamsWithHTTPClient creates a new GetFollowedContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFollowedContentParamsWithHTTPClient(client *http.Client) *GetFollowedContentParams {
	var (
		limitDefault  = int64(1000)
		offsetDefault = int64(0)
	)
	return &GetFollowedContentParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetFollowedContentParams contains all the parameters to send to the API endpoint
for the get followed content operation typically these are written to a http.Request
*/
type GetFollowedContentParams struct {

	/*Limit
	  number of content per page

	*/
	Limit *int64
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Offset
	  the offset number to retrieve

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get followed content params
func (o *GetFollowedContentParams) WithTimeout(timeout time.Duration) *GetFollowedContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get followed content params
func (o *GetFollowedContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get followed content params
func (o *GetFollowedContentParams) WithContext(ctx context.Context) *GetFollowedContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get followed content params
func (o *GetFollowedContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get followed content params
func (o *GetFollowedContentParams) WithHTTPClient(client *http.Client) *GetFollowedContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get followed content params
func (o *GetFollowedContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get followed content params
func (o *GetFollowedContentParams) WithLimit(limit *int64) *GetFollowedContentParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get followed content params
func (o *GetFollowedContentParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the get followed content params
func (o *GetFollowedContentParams) WithNamespace(namespace string) *GetFollowedContentParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get followed content params
func (o *GetFollowedContentParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get followed content params
func (o *GetFollowedContentParams) WithOffset(offset *int64) *GetFollowedContentParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get followed content params
func (o *GetFollowedContentParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetFollowedContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

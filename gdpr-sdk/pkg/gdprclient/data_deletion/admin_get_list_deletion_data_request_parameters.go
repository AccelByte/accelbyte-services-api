// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package data_deletion

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

// NewAdminGetListDeletionDataRequestParams creates a new AdminGetListDeletionDataRequestParams object
// with the default values initialized.
func NewAdminGetListDeletionDataRequestParams() *AdminGetListDeletionDataRequestParams {
	var ()
	return &AdminGetListDeletionDataRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetListDeletionDataRequestParamsWithTimeout creates a new AdminGetListDeletionDataRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetListDeletionDataRequestParamsWithTimeout(timeout time.Duration) *AdminGetListDeletionDataRequestParams {
	var ()
	return &AdminGetListDeletionDataRequestParams{

		timeout: timeout,
	}
}

// NewAdminGetListDeletionDataRequestParamsWithContext creates a new AdminGetListDeletionDataRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetListDeletionDataRequestParamsWithContext(ctx context.Context) *AdminGetListDeletionDataRequestParams {
	var ()
	return &AdminGetListDeletionDataRequestParams{

		Context: ctx,
	}
}

// NewAdminGetListDeletionDataRequestParamsWithHTTPClient creates a new AdminGetListDeletionDataRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetListDeletionDataRequestParamsWithHTTPClient(client *http.Client) *AdminGetListDeletionDataRequestParams {
	var ()
	return &AdminGetListDeletionDataRequestParams{
		HTTPClient: client,
	}
}

/*AdminGetListDeletionDataRequestParams contains all the parameters to send to the API endpoint
for the admin get list deletion data request operation typically these are written to a http.Request
*/
type AdminGetListDeletionDataRequestParams struct {

	/*After
	  The day in UTC format to get deletion request list after that. format : YYYY-MM-DD

	*/
	After *string
	/*Before
	  The day in UTC format to get deletion request list before that. format : YYYY-MM-DD

	*/
	Before *string
	/*Limit
	  the maximum number of data that may be returned (1...100)

	*/
	Limit *int64
	/*Namespace
	  namespace of the user

	*/
	Namespace string
	/*Offset
	  The start position that points to query data

	*/
	Offset *int64
	/*RequestDate
	  The day in UTC of the deletion request. format : YYYY-MM-DD. Default : today time in UTC.

	*/
	RequestDate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithTimeout(timeout time.Duration) *AdminGetListDeletionDataRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithContext(ctx context.Context) *AdminGetListDeletionDataRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithHTTPClient(client *http.Client) *AdminGetListDeletionDataRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithAfter(after *string) *AdminGetListDeletionDataRequestParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithBefore(before *string) *AdminGetListDeletionDataRequestParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetBefore(before *string) {
	o.Before = before
}

// WithLimit adds the limit to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithLimit(limit *int64) *AdminGetListDeletionDataRequestParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNamespace adds the namespace to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithNamespace(namespace string) *AdminGetListDeletionDataRequestParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithOffset(offset *int64) *AdminGetListDeletionDataRequestParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRequestDate adds the requestDate to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) WithRequestDate(requestDate *string) *AdminGetListDeletionDataRequestParams {
	o.SetRequestDate(requestDate)
	return o
}

// SetRequestDate adds the requestDate to the admin get list deletion data request params
func (o *AdminGetListDeletionDataRequestParams) SetRequestDate(requestDate *string) {
	o.RequestDate = requestDate
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetListDeletionDataRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string
		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {
			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}

	}

	if o.Before != nil {

		// query param before
		var qrBefore string
		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {
			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}

	}

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

	if o.RequestDate != nil {

		// query param requestDate
		var qrRequestDate string
		if o.RequestDate != nil {
			qrRequestDate = *o.RequestDate
		}
		qRequestDate := qrRequestDate
		if qRequestDate != "" {
			if err := r.SetQueryParam("requestDate", qRequestDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

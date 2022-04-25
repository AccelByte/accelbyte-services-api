// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package event_v2

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

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclientmodels"
)

// NewQueryEventStreamHandlerParams creates a new QueryEventStreamHandlerParams object
// with the default values initialized.
func NewQueryEventStreamHandlerParams() *QueryEventStreamHandlerParams {
	var ()
	return &QueryEventStreamHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryEventStreamHandlerParamsWithTimeout creates a new QueryEventStreamHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryEventStreamHandlerParamsWithTimeout(timeout time.Duration) *QueryEventStreamHandlerParams {
	var ()
	return &QueryEventStreamHandlerParams{

		timeout: timeout,
	}
}

// NewQueryEventStreamHandlerParamsWithContext creates a new QueryEventStreamHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryEventStreamHandlerParamsWithContext(ctx context.Context) *QueryEventStreamHandlerParams {
	var ()
	return &QueryEventStreamHandlerParams{

		Context: ctx,
	}
}

// NewQueryEventStreamHandlerParamsWithHTTPClient creates a new QueryEventStreamHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryEventStreamHandlerParamsWithHTTPClient(client *http.Client) *QueryEventStreamHandlerParams {
	var ()
	return &QueryEventStreamHandlerParams{
		HTTPClient: client,
	}
}

/*QueryEventStreamHandlerParams contains all the parameters to send to the API endpoint
for the query event stream handler operation typically these are written to a http.Request
*/
type QueryEventStreamHandlerParams struct {

	/*Body*/
	Body *eventlogclientmodels.ModelsGenericQueryPayload
	/*EndDate
	  Ending date. e.g. 2015-03-20T12:27:06Z. Default : Current time in UTC

	*/
	EndDate *string
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*Offset
	  Offset to query. Default : 0

	*/
	Offset *int64
	/*PageSize
	  Number of result in a page. Default : 100. Max : 500

	*/
	PageSize *int64
	/*StartDate
	  Starting date. e.g. 2015-03-20T12:27:06Z. Default : 1970-01-01T00:00:00Z

	*/
	StartDate *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithTimeout(timeout time.Duration) *QueryEventStreamHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithContext(ctx context.Context) *QueryEventStreamHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithHTTPClient(client *http.Client) *QueryEventStreamHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithBody(body *eventlogclientmodels.ModelsGenericQueryPayload) *QueryEventStreamHandlerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetBody(body *eventlogclientmodels.ModelsGenericQueryPayload) {
	o.Body = body
}

// WithEndDate adds the endDate to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithEndDate(endDate *string) *QueryEventStreamHandlerParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetEndDate(endDate *string) {
	o.EndDate = endDate
}

// WithNamespace adds the namespace to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithNamespace(namespace string) *QueryEventStreamHandlerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithOffset(offset *int64) *QueryEventStreamHandlerParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageSize adds the pageSize to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithPageSize(pageSize *int64) *QueryEventStreamHandlerParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithStartDate adds the startDate to the query event stream handler params
func (o *QueryEventStreamHandlerParams) WithStartDate(startDate *string) *QueryEventStreamHandlerParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the query event stream handler params
func (o *QueryEventStreamHandlerParams) SetStartDate(startDate *string) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *QueryEventStreamHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate string
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
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

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate string
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

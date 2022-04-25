// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewGetEventByEventIDHandlerParams creates a new GetEventByEventIDHandlerParams object
// with the default values initialized.
func NewGetEventByEventIDHandlerParams() *GetEventByEventIDHandlerParams {
	var ()
	return &GetEventByEventIDHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventByEventIDHandlerParamsWithTimeout creates a new GetEventByEventIDHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventByEventIDHandlerParamsWithTimeout(timeout time.Duration) *GetEventByEventIDHandlerParams {
	var ()
	return &GetEventByEventIDHandlerParams{

		timeout: timeout,
	}
}

// NewGetEventByEventIDHandlerParamsWithContext creates a new GetEventByEventIDHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventByEventIDHandlerParamsWithContext(ctx context.Context) *GetEventByEventIDHandlerParams {
	var ()
	return &GetEventByEventIDHandlerParams{

		Context: ctx,
	}
}

// NewGetEventByEventIDHandlerParamsWithHTTPClient creates a new GetEventByEventIDHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventByEventIDHandlerParamsWithHTTPClient(client *http.Client) *GetEventByEventIDHandlerParams {
	var ()
	return &GetEventByEventIDHandlerParams{
		HTTPClient: client,
	}
}

/*GetEventByEventIDHandlerParams contains all the parameters to send to the API endpoint
for the get event by event ID handler operation typically these are written to a http.Request
*/
type GetEventByEventIDHandlerParams struct {

	/*EndDate
	  Ending date. e.g. 2015-03-20T12:27:06.351Z

	*/
	EndDate string
	/*EventID
	  Event's ID

	*/
	EventID float64
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*Offset
	  Offset to query

	*/
	Offset *int64
	/*PageSize
	  Number of result in a page

	*/
	PageSize int64
	/*StartDate
	  Starting date. e.g. 2015-03-20T12:27:06.351Z

	*/
	StartDate string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithTimeout(timeout time.Duration) *GetEventByEventIDHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithContext(ctx context.Context) *GetEventByEventIDHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithHTTPClient(client *http.Client) *GetEventByEventIDHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithEndDate(endDate string) *GetEventByEventIDHandlerParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetEndDate(endDate string) {
	o.EndDate = endDate
}

// WithEventID adds the eventID to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithEventID(eventID float64) *GetEventByEventIDHandlerParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetEventID(eventID float64) {
	o.EventID = eventID
}

// WithNamespace adds the namespace to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithNamespace(namespace string) *GetEventByEventIDHandlerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithOffset(offset *int64) *GetEventByEventIDHandlerParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageSize adds the pageSize to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithPageSize(pageSize int64) *GetEventByEventIDHandlerParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetPageSize(pageSize int64) {
	o.PageSize = pageSize
}

// WithStartDate adds the startDate to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) WithStartDate(startDate string) *GetEventByEventIDHandlerParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get event by event ID handler params
func (o *GetEventByEventIDHandlerParams) SetStartDate(startDate string) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventByEventIDHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endDate
	qrEndDate := o.EndDate
	qEndDate := qrEndDate
	if qEndDate != "" {
		if err := r.SetQueryParam("endDate", qEndDate); err != nil {
			return err
		}
	}

	// path param eventId
	if err := r.SetPathParam("eventId", swag.FormatFloat64(o.EventID)); err != nil {
		return err
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

	// query param pageSize
	qrPageSize := o.PageSize
	qPageSize := swag.FormatInt64(qrPageSize)
	if qPageSize != "" {
		if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
			return err
		}
	}

	// query param startDate
	qrStartDate := o.StartDate
	qStartDate := qrStartDate
	if qStartDate != "" {
		if err := r.SetQueryParam("startDate", qStartDate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

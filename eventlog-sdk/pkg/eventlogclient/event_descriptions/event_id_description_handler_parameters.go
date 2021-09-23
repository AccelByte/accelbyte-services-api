// Code generated by go-swagger; DO NOT EDIT.

package event_descriptions

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

// NewEventIDDescriptionHandlerParams creates a new EventIDDescriptionHandlerParams object
// with the default values initialized.
func NewEventIDDescriptionHandlerParams() *EventIDDescriptionHandlerParams {

	return &EventIDDescriptionHandlerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEventIDDescriptionHandlerParamsWithTimeout creates a new EventIDDescriptionHandlerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEventIDDescriptionHandlerParamsWithTimeout(timeout time.Duration) *EventIDDescriptionHandlerParams {

	return &EventIDDescriptionHandlerParams{

		timeout: timeout,
	}
}

// NewEventIDDescriptionHandlerParamsWithContext creates a new EventIDDescriptionHandlerParams object
// with the default values initialized, and the ability to set a context for a request
func NewEventIDDescriptionHandlerParamsWithContext(ctx context.Context) *EventIDDescriptionHandlerParams {

	return &EventIDDescriptionHandlerParams{

		Context: ctx,
	}
}

// NewEventIDDescriptionHandlerParamsWithHTTPClient creates a new EventIDDescriptionHandlerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEventIDDescriptionHandlerParamsWithHTTPClient(client *http.Client) *EventIDDescriptionHandlerParams {

	return &EventIDDescriptionHandlerParams{
		HTTPClient: client,
	}
}

/*EventIDDescriptionHandlerParams contains all the parameters to send to the API endpoint
for the event ID description handler operation typically these are written to a http.Request
*/
type EventIDDescriptionHandlerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the event ID description handler params
func (o *EventIDDescriptionHandlerParams) WithTimeout(timeout time.Duration) *EventIDDescriptionHandlerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the event ID description handler params
func (o *EventIDDescriptionHandlerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the event ID description handler params
func (o *EventIDDescriptionHandlerParams) WithContext(ctx context.Context) *EventIDDescriptionHandlerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the event ID description handler params
func (o *EventIDDescriptionHandlerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the event ID description handler params
func (o *EventIDDescriptionHandlerParams) WithHTTPClient(client *http.Client) *EventIDDescriptionHandlerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the event ID description handler params
func (o *EventIDDescriptionHandlerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EventIDDescriptionHandlerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

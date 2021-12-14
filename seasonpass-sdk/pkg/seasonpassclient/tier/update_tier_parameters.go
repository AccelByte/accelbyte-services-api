// Code generated by go-swagger; DO NOT EDIT.

package tier

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

	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
)

// NewUpdateTierParams creates a new UpdateTierParams object
// with the default values initialized.
func NewUpdateTierParams() *UpdateTierParams {
	var ()
	return &UpdateTierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTierParamsWithTimeout creates a new UpdateTierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateTierParamsWithTimeout(timeout time.Duration) *UpdateTierParams {
	var ()
	return &UpdateTierParams{

		timeout: timeout,
	}
}

// NewUpdateTierParamsWithContext creates a new UpdateTierParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateTierParamsWithContext(ctx context.Context) *UpdateTierParams {
	var ()
	return &UpdateTierParams{

		Context: ctx,
	}
}

// NewUpdateTierParamsWithHTTPClient creates a new UpdateTierParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateTierParamsWithHTTPClient(client *http.Client) *UpdateTierParams {
	var ()
	return &UpdateTierParams{
		HTTPClient: client,
	}
}

/*UpdateTierParams contains all the parameters to send to the API endpoint
for the update tier operation typically these are written to a http.Request
*/
type UpdateTierParams struct {

	/*Body*/
	Body *seasonpassclientmodels.TierInput
	/*ID*/
	ID string
	/*Namespace*/
	Namespace string
	/*SeasonID*/
	SeasonID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update tier params
func (o *UpdateTierParams) WithTimeout(timeout time.Duration) *UpdateTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tier params
func (o *UpdateTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tier params
func (o *UpdateTierParams) WithContext(ctx context.Context) *UpdateTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tier params
func (o *UpdateTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tier params
func (o *UpdateTierParams) WithHTTPClient(client *http.Client) *UpdateTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tier params
func (o *UpdateTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tier params
func (o *UpdateTierParams) WithBody(body *seasonpassclientmodels.TierInput) *UpdateTierParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tier params
func (o *UpdateTierParams) SetBody(body *seasonpassclientmodels.TierInput) {
	o.Body = body
}

// WithID adds the id to the update tier params
func (o *UpdateTierParams) WithID(id string) *UpdateTierParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update tier params
func (o *UpdateTierParams) SetID(id string) {
	o.ID = id
}

// WithNamespace adds the namespace to the update tier params
func (o *UpdateTierParams) WithNamespace(namespace string) *UpdateTierParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update tier params
func (o *UpdateTierParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithSeasonID adds the seasonID to the update tier params
func (o *UpdateTierParams) WithSeasonID(seasonID string) *UpdateTierParams {
	o.SetSeasonID(seasonID)
	return o
}

// SetSeasonID adds the seasonId to the update tier params
func (o *UpdateTierParams) SetSeasonID(seasonID string) {
	o.SeasonID = seasonID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param seasonId
	if err := r.SetPathParam("seasonId", o.SeasonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

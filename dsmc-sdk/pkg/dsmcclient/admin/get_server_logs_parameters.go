// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewGetServerLogsParams creates a new GetServerLogsParams object
// with the default values initialized.
func NewGetServerLogsParams() *GetServerLogsParams {
	var ()
	return &GetServerLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerLogsParamsWithTimeout creates a new GetServerLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerLogsParamsWithTimeout(timeout time.Duration) *GetServerLogsParams {
	var ()
	return &GetServerLogsParams{

		timeout: timeout,
	}
}

// NewGetServerLogsParamsWithContext creates a new GetServerLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServerLogsParamsWithContext(ctx context.Context) *GetServerLogsParams {
	var ()
	return &GetServerLogsParams{

		Context: ctx,
	}
}

// NewGetServerLogsParamsWithHTTPClient creates a new GetServerLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServerLogsParamsWithHTTPClient(client *http.Client) *GetServerLogsParams {
	var ()
	return &GetServerLogsParams{
		HTTPClient: client,
	}
}

/*GetServerLogsParams contains all the parameters to send to the API endpoint
for the get server logs operation typically these are written to a http.Request
*/
type GetServerLogsParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*PodName
	  name of the DS pod

	*/
	PodName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get server logs params
func (o *GetServerLogsParams) WithTimeout(timeout time.Duration) *GetServerLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server logs params
func (o *GetServerLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server logs params
func (o *GetServerLogsParams) WithContext(ctx context.Context) *GetServerLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server logs params
func (o *GetServerLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server logs params
func (o *GetServerLogsParams) WithHTTPClient(client *http.Client) *GetServerLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server logs params
func (o *GetServerLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get server logs params
func (o *GetServerLogsParams) WithNamespace(namespace string) *GetServerLogsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get server logs params
func (o *GetServerLogsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPodName adds the podName to the get server logs params
func (o *GetServerLogsParams) WithPodName(podName string) *GetServerLogsParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the get server logs params
func (o *GetServerLogsParams) SetPodName(podName string) {
	o.PodName = podName
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
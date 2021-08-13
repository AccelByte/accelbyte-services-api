// Code generated by go-swagger; DO NOT EDIT.

package clients

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

// NewDeleteClientByNamespaceParams creates a new DeleteClientByNamespaceParams object
// with the default values initialized.
func NewDeleteClientByNamespaceParams() *DeleteClientByNamespaceParams {
	var ()
	return &DeleteClientByNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClientByNamespaceParamsWithTimeout creates a new DeleteClientByNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClientByNamespaceParamsWithTimeout(timeout time.Duration) *DeleteClientByNamespaceParams {
	var ()
	return &DeleteClientByNamespaceParams{

		timeout: timeout,
	}
}

// NewDeleteClientByNamespaceParamsWithContext creates a new DeleteClientByNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClientByNamespaceParamsWithContext(ctx context.Context) *DeleteClientByNamespaceParams {
	var ()
	return &DeleteClientByNamespaceParams{

		Context: ctx,
	}
}

// NewDeleteClientByNamespaceParamsWithHTTPClient creates a new DeleteClientByNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClientByNamespaceParamsWithHTTPClient(client *http.Client) *DeleteClientByNamespaceParams {
	var ()
	return &DeleteClientByNamespaceParams{
		HTTPClient: client,
	}
}

/*DeleteClientByNamespaceParams contains all the parameters to send to the API endpoint
for the delete client by namespace operation typically these are written to a http.Request
*/
type DeleteClientByNamespaceParams struct {

	/*ClientID
	  Client ID

	*/
	ClientID string
	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) WithTimeout(timeout time.Duration) *DeleteClientByNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) WithContext(ctx context.Context) *DeleteClientByNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) WithHTTPClient(client *http.Client) *DeleteClientByNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) WithClientID(clientID string) *DeleteClientByNamespaceParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithNamespace adds the namespace to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) WithNamespace(namespace string) *DeleteClientByNamespaceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete client by namespace params
func (o *DeleteClientByNamespaceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClientByNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

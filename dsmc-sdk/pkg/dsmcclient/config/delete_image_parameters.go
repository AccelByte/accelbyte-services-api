// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewDeleteImageParams creates a new DeleteImageParams object
// with the default values initialized.
func NewDeleteImageParams() *DeleteImageParams {
	var ()
	return &DeleteImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImageParamsWithTimeout creates a new DeleteImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteImageParamsWithTimeout(timeout time.Duration) *DeleteImageParams {
	var ()
	return &DeleteImageParams{

		timeout: timeout,
	}
}

// NewDeleteImageParamsWithContext creates a new DeleteImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteImageParamsWithContext(ctx context.Context) *DeleteImageParams {
	var ()
	return &DeleteImageParams{

		Context: ctx,
	}
}

// NewDeleteImageParamsWithHTTPClient creates a new DeleteImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteImageParamsWithHTTPClient(client *http.Client) *DeleteImageParams {
	var ()
	return &DeleteImageParams{
		HTTPClient: client,
	}
}

/*DeleteImageParams contains all the parameters to send to the API endpoint
for the delete image operation typically these are written to a http.Request
*/
type DeleteImageParams struct {

	/*ImageURI
	  registry image URI that will be deleted

	*/
	ImageURI string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Version
	  image version that will be deleted

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete image params
func (o *DeleteImageParams) WithTimeout(timeout time.Duration) *DeleteImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete image params
func (o *DeleteImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete image params
func (o *DeleteImageParams) WithContext(ctx context.Context) *DeleteImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete image params
func (o *DeleteImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete image params
func (o *DeleteImageParams) WithHTTPClient(client *http.Client) *DeleteImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete image params
func (o *DeleteImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageURI adds the imageURI to the delete image params
func (o *DeleteImageParams) WithImageURI(imageURI string) *DeleteImageParams {
	o.SetImageURI(imageURI)
	return o
}

// SetImageURI adds the imageUri to the delete image params
func (o *DeleteImageParams) SetImageURI(imageURI string) {
	o.ImageURI = imageURI
}

// WithNamespace adds the namespace to the delete image params
func (o *DeleteImageParams) WithNamespace(namespace string) *DeleteImageParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete image params
func (o *DeleteImageParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVersion adds the version to the delete image params
func (o *DeleteImageParams) WithVersion(version string) *DeleteImageParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete image params
func (o *DeleteImageParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param imageURI
	qrImageURI := o.ImageURI
	qImageURI := qrImageURI
	if qImageURI != "" {
		if err := r.SetQueryParam("imageURI", qImageURI); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

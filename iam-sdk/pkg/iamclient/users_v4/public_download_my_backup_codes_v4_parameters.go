// Code generated by go-swagger; DO NOT EDIT.

package users_v4

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

// NewPublicDownloadMyBackupCodesV4Params creates a new PublicDownloadMyBackupCodesV4Params object
// with the default values initialized.
func NewPublicDownloadMyBackupCodesV4Params() *PublicDownloadMyBackupCodesV4Params {
	var ()
	return &PublicDownloadMyBackupCodesV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicDownloadMyBackupCodesV4ParamsWithTimeout creates a new PublicDownloadMyBackupCodesV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicDownloadMyBackupCodesV4ParamsWithTimeout(timeout time.Duration) *PublicDownloadMyBackupCodesV4Params {
	var ()
	return &PublicDownloadMyBackupCodesV4Params{

		timeout: timeout,
	}
}

// NewPublicDownloadMyBackupCodesV4ParamsWithContext creates a new PublicDownloadMyBackupCodesV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicDownloadMyBackupCodesV4ParamsWithContext(ctx context.Context) *PublicDownloadMyBackupCodesV4Params {
	var ()
	return &PublicDownloadMyBackupCodesV4Params{

		Context: ctx,
	}
}

// NewPublicDownloadMyBackupCodesV4ParamsWithHTTPClient creates a new PublicDownloadMyBackupCodesV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicDownloadMyBackupCodesV4ParamsWithHTTPClient(client *http.Client) *PublicDownloadMyBackupCodesV4Params {
	var ()
	return &PublicDownloadMyBackupCodesV4Params{
		HTTPClient: client,
	}
}

/*PublicDownloadMyBackupCodesV4Params contains all the parameters to send to the API endpoint
for the public download my backup codes v4 operation typically these are written to a http.Request
*/
type PublicDownloadMyBackupCodesV4Params struct {

	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) WithTimeout(timeout time.Duration) *PublicDownloadMyBackupCodesV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) WithContext(ctx context.Context) *PublicDownloadMyBackupCodesV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) WithHTTPClient(client *http.Client) *PublicDownloadMyBackupCodesV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) WithNamespace(namespace string) *PublicDownloadMyBackupCodesV4Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public download my backup codes v4 params
func (o *PublicDownloadMyBackupCodesV4Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicDownloadMyBackupCodesV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

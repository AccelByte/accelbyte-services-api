// Code generated by go-swagger; DO NOT EDIT.

package nr_admin_content

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

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// NewAdminUploadContentScreenshotParams creates a new AdminUploadContentScreenshotParams object
// with the default values initialized.
func NewAdminUploadContentScreenshotParams() *AdminUploadContentScreenshotParams {
	var ()
	return &AdminUploadContentScreenshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUploadContentScreenshotParamsWithTimeout creates a new AdminUploadContentScreenshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminUploadContentScreenshotParamsWithTimeout(timeout time.Duration) *AdminUploadContentScreenshotParams {
	var ()
	return &AdminUploadContentScreenshotParams{

		timeout: timeout,
	}
}

// NewAdminUploadContentScreenshotParamsWithContext creates a new AdminUploadContentScreenshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminUploadContentScreenshotParamsWithContext(ctx context.Context) *AdminUploadContentScreenshotParams {
	var ()
	return &AdminUploadContentScreenshotParams{

		Context: ctx,
	}
}

// NewAdminUploadContentScreenshotParamsWithHTTPClient creates a new AdminUploadContentScreenshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminUploadContentScreenshotParamsWithHTTPClient(client *http.Client) *AdminUploadContentScreenshotParams {
	var ()
	return &AdminUploadContentScreenshotParams{
		HTTPClient: client,
	}
}

/*AdminUploadContentScreenshotParams contains all the parameters to send to the API endpoint
for the admin upload content screenshot operation typically these are written to a http.Request
*/
type AdminUploadContentScreenshotParams struct {

	/*Body*/
	Body *ugcclientmodels.ModelsCreateScreenshotRequest
	/*ContentID
	  content ID

	*/
	ContentID string
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) WithTimeout(timeout time.Duration) *AdminUploadContentScreenshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) WithContext(ctx context.Context) *AdminUploadContentScreenshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) WithHTTPClient(client *http.Client) *AdminUploadContentScreenshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) WithBody(body *ugcclientmodels.ModelsCreateScreenshotRequest) *AdminUploadContentScreenshotParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) SetBody(body *ugcclientmodels.ModelsCreateScreenshotRequest) {
	o.Body = body
}

// WithContentID adds the contentID to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) WithContentID(contentID string) *AdminUploadContentScreenshotParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithNamespace adds the namespace to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) WithNamespace(namespace string) *AdminUploadContentScreenshotParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin upload content screenshot params
func (o *AdminUploadContentScreenshotParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUploadContentScreenshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
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

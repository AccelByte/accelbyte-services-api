// Code generated by go-swagger; DO NOT EDIT.

package nr_public_content

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

// NewDeleteContentScreenshotParams creates a new DeleteContentScreenshotParams object
// with the default values initialized.
func NewDeleteContentScreenshotParams() *DeleteContentScreenshotParams {
	var ()
	return &DeleteContentScreenshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentScreenshotParamsWithTimeout creates a new DeleteContentScreenshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteContentScreenshotParamsWithTimeout(timeout time.Duration) *DeleteContentScreenshotParams {
	var ()
	return &DeleteContentScreenshotParams{

		timeout: timeout,
	}
}

// NewDeleteContentScreenshotParamsWithContext creates a new DeleteContentScreenshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteContentScreenshotParamsWithContext(ctx context.Context) *DeleteContentScreenshotParams {
	var ()
	return &DeleteContentScreenshotParams{

		Context: ctx,
	}
}

// NewDeleteContentScreenshotParamsWithHTTPClient creates a new DeleteContentScreenshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteContentScreenshotParamsWithHTTPClient(client *http.Client) *DeleteContentScreenshotParams {
	var ()
	return &DeleteContentScreenshotParams{
		HTTPClient: client,
	}
}

/*DeleteContentScreenshotParams contains all the parameters to send to the API endpoint
for the delete content screenshot operation typically these are written to a http.Request
*/
type DeleteContentScreenshotParams struct {

	/*ContentID
	  content ID

	*/
	ContentID string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*ScreenshotID
	  screenshot ID

	*/
	ScreenshotID string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithTimeout(timeout time.Duration) *DeleteContentScreenshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithContext(ctx context.Context) *DeleteContentScreenshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithHTTPClient(client *http.Client) *DeleteContentScreenshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentID adds the contentID to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithContentID(contentID string) *DeleteContentScreenshotParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithNamespace adds the namespace to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithNamespace(namespace string) *DeleteContentScreenshotParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithScreenshotID adds the screenshotID to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithScreenshotID(screenshotID string) *DeleteContentScreenshotParams {
	o.SetScreenshotID(screenshotID)
	return o
}

// SetScreenshotID adds the screenshotId to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetScreenshotID(screenshotID string) {
	o.ScreenshotID = screenshotID
}

// WithUserID adds the userID to the delete content screenshot params
func (o *DeleteContentScreenshotParams) WithUserID(userID string) *DeleteContentScreenshotParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete content screenshot params
func (o *DeleteContentScreenshotParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentScreenshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contentId
	if err := r.SetPathParam("contentId", o.ContentID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param screenshotId
	if err := r.SetPathParam("screenshotId", o.ScreenshotID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

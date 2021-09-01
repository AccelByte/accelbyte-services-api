// Code generated by go-swagger; DO NOT EDIT.

package friends

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

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NewUserAcceptFriendRequestParams creates a new UserAcceptFriendRequestParams object
// with the default values initialized.
func NewUserAcceptFriendRequestParams() *UserAcceptFriendRequestParams {
	var ()
	return &UserAcceptFriendRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserAcceptFriendRequestParamsWithTimeout creates a new UserAcceptFriendRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserAcceptFriendRequestParamsWithTimeout(timeout time.Duration) *UserAcceptFriendRequestParams {
	var ()
	return &UserAcceptFriendRequestParams{

		timeout: timeout,
	}
}

// NewUserAcceptFriendRequestParamsWithContext creates a new UserAcceptFriendRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserAcceptFriendRequestParamsWithContext(ctx context.Context) *UserAcceptFriendRequestParams {
	var ()
	return &UserAcceptFriendRequestParams{

		Context: ctx,
	}
}

// NewUserAcceptFriendRequestParamsWithHTTPClient creates a new UserAcceptFriendRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserAcceptFriendRequestParamsWithHTTPClient(client *http.Client) *UserAcceptFriendRequestParams {
	var ()
	return &UserAcceptFriendRequestParams{
		HTTPClient: client,
	}
}

/*UserAcceptFriendRequestParams contains all the parameters to send to the API endpoint
for the user accept friend request operation typically these are written to a http.Request
*/
type UserAcceptFriendRequestParams struct {

	/*Body
	  accept friend content

	*/
	Body *lobbyclientmodels.ModelUserAcceptFriendRequest
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user accept friend request params
func (o *UserAcceptFriendRequestParams) WithTimeout(timeout time.Duration) *UserAcceptFriendRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user accept friend request params
func (o *UserAcceptFriendRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user accept friend request params
func (o *UserAcceptFriendRequestParams) WithContext(ctx context.Context) *UserAcceptFriendRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user accept friend request params
func (o *UserAcceptFriendRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user accept friend request params
func (o *UserAcceptFriendRequestParams) WithHTTPClient(client *http.Client) *UserAcceptFriendRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user accept friend request params
func (o *UserAcceptFriendRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user accept friend request params
func (o *UserAcceptFriendRequestParams) WithBody(body *lobbyclientmodels.ModelUserAcceptFriendRequest) *UserAcceptFriendRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user accept friend request params
func (o *UserAcceptFriendRequestParams) SetBody(body *lobbyclientmodels.ModelUserAcceptFriendRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the user accept friend request params
func (o *UserAcceptFriendRequestParams) WithNamespace(namespace string) *UserAcceptFriendRequestParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the user accept friend request params
func (o *UserAcceptFriendRequestParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *UserAcceptFriendRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
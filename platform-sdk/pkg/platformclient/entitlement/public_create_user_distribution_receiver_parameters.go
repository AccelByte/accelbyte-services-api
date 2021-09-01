// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewPublicCreateUserDistributionReceiverParams creates a new PublicCreateUserDistributionReceiverParams object
// with the default values initialized.
func NewPublicCreateUserDistributionReceiverParams() *PublicCreateUserDistributionReceiverParams {
	var ()
	return &PublicCreateUserDistributionReceiverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicCreateUserDistributionReceiverParamsWithTimeout creates a new PublicCreateUserDistributionReceiverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicCreateUserDistributionReceiverParamsWithTimeout(timeout time.Duration) *PublicCreateUserDistributionReceiverParams {
	var ()
	return &PublicCreateUserDistributionReceiverParams{

		timeout: timeout,
	}
}

// NewPublicCreateUserDistributionReceiverParamsWithContext creates a new PublicCreateUserDistributionReceiverParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicCreateUserDistributionReceiverParamsWithContext(ctx context.Context) *PublicCreateUserDistributionReceiverParams {
	var ()
	return &PublicCreateUserDistributionReceiverParams{

		Context: ctx,
	}
}

// NewPublicCreateUserDistributionReceiverParamsWithHTTPClient creates a new PublicCreateUserDistributionReceiverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicCreateUserDistributionReceiverParamsWithHTTPClient(client *http.Client) *PublicCreateUserDistributionReceiverParams {
	var ()
	return &PublicCreateUserDistributionReceiverParams{
		HTTPClient: client,
	}
}

/*PublicCreateUserDistributionReceiverParams contains all the parameters to send to the API endpoint
for the public create user distribution receiver operation typically these are written to a http.Request
*/
type PublicCreateUserDistributionReceiverParams struct {

	/*Body*/
	Body *platformclientmodels.DistributionReceiverCreate
	/*ExtUserID*/
	ExtUserID string
	/*Namespace
	  game namespace

	*/
	Namespace string
	/*UserID
	  game user id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithTimeout(timeout time.Duration) *PublicCreateUserDistributionReceiverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithContext(ctx context.Context) *PublicCreateUserDistributionReceiverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithHTTPClient(client *http.Client) *PublicCreateUserDistributionReceiverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithBody(body *platformclientmodels.DistributionReceiverCreate) *PublicCreateUserDistributionReceiverParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetBody(body *platformclientmodels.DistributionReceiverCreate) {
	o.Body = body
}

// WithExtUserID adds the extUserID to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithExtUserID(extUserID string) *PublicCreateUserDistributionReceiverParams {
	o.SetExtUserID(extUserID)
	return o
}

// SetExtUserID adds the extUserId to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetExtUserID(extUserID string) {
	o.ExtUserID = extUserID
}

// WithNamespace adds the namespace to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithNamespace(namespace string) *PublicCreateUserDistributionReceiverParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) WithUserID(userID string) *PublicCreateUserDistributionReceiverParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public create user distribution receiver params
func (o *PublicCreateUserDistributionReceiverParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicCreateUserDistributionReceiverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param extUserId
	if err := r.SetPathParam("extUserId", o.ExtUserID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
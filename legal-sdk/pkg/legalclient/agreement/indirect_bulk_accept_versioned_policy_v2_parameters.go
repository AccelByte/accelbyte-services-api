// Code generated by go-swagger; DO NOT EDIT.

package agreement

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

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
)

// NewIndirectBulkAcceptVersionedPolicyV2Params creates a new IndirectBulkAcceptVersionedPolicyV2Params object
// with the default values initialized.
func NewIndirectBulkAcceptVersionedPolicyV2Params() *IndirectBulkAcceptVersionedPolicyV2Params {
	var ()
	return &IndirectBulkAcceptVersionedPolicyV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndirectBulkAcceptVersionedPolicyV2ParamsWithTimeout creates a new IndirectBulkAcceptVersionedPolicyV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndirectBulkAcceptVersionedPolicyV2ParamsWithTimeout(timeout time.Duration) *IndirectBulkAcceptVersionedPolicyV2Params {
	var ()
	return &IndirectBulkAcceptVersionedPolicyV2Params{

		timeout: timeout,
	}
}

// NewIndirectBulkAcceptVersionedPolicyV2ParamsWithContext creates a new IndirectBulkAcceptVersionedPolicyV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewIndirectBulkAcceptVersionedPolicyV2ParamsWithContext(ctx context.Context) *IndirectBulkAcceptVersionedPolicyV2Params {
	var ()
	return &IndirectBulkAcceptVersionedPolicyV2Params{

		Context: ctx,
	}
}

// NewIndirectBulkAcceptVersionedPolicyV2ParamsWithHTTPClient creates a new IndirectBulkAcceptVersionedPolicyV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndirectBulkAcceptVersionedPolicyV2ParamsWithHTTPClient(client *http.Client) *IndirectBulkAcceptVersionedPolicyV2Params {
	var ()
	return &IndirectBulkAcceptVersionedPolicyV2Params{
		HTTPClient: client,
	}
}

/*IndirectBulkAcceptVersionedPolicyV2Params contains all the parameters to send to the API endpoint
for the indirect bulk accept versioned policy v2 operation typically these are written to a http.Request
*/
type IndirectBulkAcceptVersionedPolicyV2Params struct {

	/*Body*/
	Body []*legalclientmodels.AcceptAgreementRequest
	/*ClientID
	  Client ID

	*/
	ClientID string
	/*CountryCode
	  Country Code

	*/
	CountryCode string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*UserID
	  User Id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithTimeout(timeout time.Duration) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithContext(ctx context.Context) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithHTTPClient(client *http.Client) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithBody(body []*legalclientmodels.AcceptAgreementRequest) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetBody(body []*legalclientmodels.AcceptAgreementRequest) {
	o.Body = body
}

// WithClientID adds the clientID to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithClientID(clientID string) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithCountryCode adds the countryCode to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithCountryCode(countryCode string) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetCountryCode(countryCode)
	return o
}

// SetCountryCode adds the countryCode to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetCountryCode(countryCode string) {
	o.CountryCode = countryCode
}

// WithNamespace adds the namespace to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithNamespace(namespace string) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WithUserID(userID string) *IndirectBulkAcceptVersionedPolicyV2Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the indirect bulk accept versioned policy v2 params
func (o *IndirectBulkAcceptVersionedPolicyV2Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *IndirectBulkAcceptVersionedPolicyV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	// path param countryCode
	if err := r.SetPathParam("countryCode", o.CountryCode); err != nil {
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

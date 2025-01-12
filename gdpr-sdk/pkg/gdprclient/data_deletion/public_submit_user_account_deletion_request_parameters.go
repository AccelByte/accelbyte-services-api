// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package data_deletion

import (
	"context"
	"net/http"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPublicSubmitUserAccountDeletionRequestParams creates a new PublicSubmitUserAccountDeletionRequestParams object
// with the default values initialized.
func NewPublicSubmitUserAccountDeletionRequestParams() *PublicSubmitUserAccountDeletionRequestParams {
	var ()
	return &PublicSubmitUserAccountDeletionRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicSubmitUserAccountDeletionRequestParamsWithTimeout creates a new PublicSubmitUserAccountDeletionRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicSubmitUserAccountDeletionRequestParamsWithTimeout(timeout time.Duration) *PublicSubmitUserAccountDeletionRequestParams {
	var ()
	return &PublicSubmitUserAccountDeletionRequestParams{

		timeout: timeout,
	}
}

// NewPublicSubmitUserAccountDeletionRequestParamsWithContext creates a new PublicSubmitUserAccountDeletionRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicSubmitUserAccountDeletionRequestParamsWithContext(ctx context.Context) *PublicSubmitUserAccountDeletionRequestParams {
	var ()
	return &PublicSubmitUserAccountDeletionRequestParams{

		Context: ctx,
	}
}

// NewPublicSubmitUserAccountDeletionRequestParamsWithHTTPClient creates a new PublicSubmitUserAccountDeletionRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicSubmitUserAccountDeletionRequestParamsWithHTTPClient(client *http.Client) *PublicSubmitUserAccountDeletionRequestParams {
	var ()
	return &PublicSubmitUserAccountDeletionRequestParams{
		HTTPClient: client,
	}
}

/*PublicSubmitUserAccountDeletionRequestParams contains all the parameters to send to the API endpoint
for the public submit user account deletion request operation typically these are written to a http.Request
*/
type PublicSubmitUserAccountDeletionRequestParams struct {

	/*RetryPolicy*/
	RetryPolicy *utils.Retry
	/*LanguageTag
	  Language tag

	*/
	LanguageTag *string
	/*Password
	  IAM password of the user

	*/
	Password string
	/*Namespace
	  namespace of the user

	*/
	Namespace string
	/*UserID
	  IAM ID of the user

	*/
	UserID string

	timeout        time.Duration
	AuthInfoWriter runtime.ClientAuthInfoWriter
	Context        context.Context
	HTTPClient     *http.Client

	// XFlightId is an optional parameter from this SDK
	XFlightId *string
}

// WithTimeout adds the timeout to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithTimeout(timeout time.Duration) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithContext(ctx context.Context) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// SetAuthInfoWriter adds the authInfoWriter to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetAuthInfoWriter(authInfoWriter runtime.ClientAuthInfoWriter) {
	o.AuthInfoWriter = authInfoWriter
}

// WithHTTPClient adds the HTTPClient to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithHTTPClient(client *http.Client) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// SetHTTPClient adds the HTTPClient Transport to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetHTTPClientTransport(roundTripper http.RoundTripper) {
	if o.HTTPClient != nil {
		o.HTTPClient.Transport = roundTripper
	} else {
		o.HTTPClient = &http.Client{Transport: roundTripper}
	}
}

// SetFlightId adds the flightId as the header value for this specific endpoint
func (o *PublicSubmitUserAccountDeletionRequestParams) SetFlightId(flightId string) {
	if o.XFlightId != nil {
		o.XFlightId = &flightId
	} else {
		o.XFlightId = &utils.GetDefaultFlightID().Value
	}
}

// WithLanguageTag adds the languageTag to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithLanguageTag(languageTag *string) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetLanguageTag(languageTag)
	return o
}

// SetLanguageTag adds the languageTag to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetLanguageTag(languageTag *string) {
	o.LanguageTag = languageTag
}

// WithPassword adds the password to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithPassword(password string) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetPassword(password string) {
	o.Password = password
}

// WithNamespace adds the namespace to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithNamespace(namespace string) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) WithUserID(userID string) *PublicSubmitUserAccountDeletionRequestParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public submit user account deletion request params
func (o *PublicSubmitUserAccountDeletionRequestParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicSubmitUserAccountDeletionRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LanguageTag != nil {

		// form param languageTag
		var frLanguageTag string
		if o.LanguageTag != nil {
			frLanguageTag = *o.LanguageTag
		}
		fLanguageTag := frLanguageTag
		if fLanguageTag != "" {
			if err := r.SetFormParam("languageTag", fLanguageTag); err != nil {
				return err
			}
		}

	}

	// form param password
	frPassword := o.Password
	fPassword := frPassword
	if fPassword != "" {
		if err := r.SetFormParam("password", fPassword); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	// setting the default header value
	if err := r.SetHeaderParam("User-Agent", utils.UserAgentGen()); err != nil {
		return err
	}

	if err := r.SetHeaderParam("X-Amzn-Trace-Id", utils.AmazonTraceIDGen()); err != nil {
		return err
	}

	if o.XFlightId == nil {
		if err := r.SetHeaderParam("X-Flight-Id", utils.GetDefaultFlightID().Value); err != nil {
			return err
		}
	} else {
		if err := r.SetHeaderParam("X-Flight-Id", *o.XFlightId); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

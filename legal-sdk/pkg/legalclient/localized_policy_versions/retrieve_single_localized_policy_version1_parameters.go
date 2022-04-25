// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package localized_policy_versions

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

// NewRetrieveSingleLocalizedPolicyVersion1Params creates a new RetrieveSingleLocalizedPolicyVersion1Params object
// with the default values initialized.
func NewRetrieveSingleLocalizedPolicyVersion1Params() *RetrieveSingleLocalizedPolicyVersion1Params {
	var ()
	return &RetrieveSingleLocalizedPolicyVersion1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveSingleLocalizedPolicyVersion1ParamsWithTimeout creates a new RetrieveSingleLocalizedPolicyVersion1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveSingleLocalizedPolicyVersion1ParamsWithTimeout(timeout time.Duration) *RetrieveSingleLocalizedPolicyVersion1Params {
	var ()
	return &RetrieveSingleLocalizedPolicyVersion1Params{

		timeout: timeout,
	}
}

// NewRetrieveSingleLocalizedPolicyVersion1ParamsWithContext creates a new RetrieveSingleLocalizedPolicyVersion1Params object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveSingleLocalizedPolicyVersion1ParamsWithContext(ctx context.Context) *RetrieveSingleLocalizedPolicyVersion1Params {
	var ()
	return &RetrieveSingleLocalizedPolicyVersion1Params{

		Context: ctx,
	}
}

// NewRetrieveSingleLocalizedPolicyVersion1ParamsWithHTTPClient creates a new RetrieveSingleLocalizedPolicyVersion1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveSingleLocalizedPolicyVersion1ParamsWithHTTPClient(client *http.Client) *RetrieveSingleLocalizedPolicyVersion1Params {
	var ()
	return &RetrieveSingleLocalizedPolicyVersion1Params{
		HTTPClient: client,
	}
}

/*RetrieveSingleLocalizedPolicyVersion1Params contains all the parameters to send to the API endpoint
for the retrieve single localized policy version 1 operation typically these are written to a http.Request
*/
type RetrieveSingleLocalizedPolicyVersion1Params struct {

	/*LocalizedPolicyVersionID
	  Localized Policy Version Id

	*/
	LocalizedPolicyVersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) WithTimeout(timeout time.Duration) *RetrieveSingleLocalizedPolicyVersion1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) WithContext(ctx context.Context) *RetrieveSingleLocalizedPolicyVersion1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) WithHTTPClient(client *http.Client) *RetrieveSingleLocalizedPolicyVersion1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocalizedPolicyVersionID adds the localizedPolicyVersionID to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) WithLocalizedPolicyVersionID(localizedPolicyVersionID string) *RetrieveSingleLocalizedPolicyVersion1Params {
	o.SetLocalizedPolicyVersionID(localizedPolicyVersionID)
	return o
}

// SetLocalizedPolicyVersionID adds the localizedPolicyVersionId to the retrieve single localized policy version 1 params
func (o *RetrieveSingleLocalizedPolicyVersion1Params) SetLocalizedPolicyVersionID(localizedPolicyVersionID string) {
	o.LocalizedPolicyVersionID = localizedPolicyVersionID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveSingleLocalizedPolicyVersion1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param localizedPolicyVersionId
	if err := r.SetPathParam("localizedPolicyVersionId", o.LocalizedPolicyVersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

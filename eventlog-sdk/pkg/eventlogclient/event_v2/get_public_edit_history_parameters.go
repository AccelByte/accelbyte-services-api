// Code generated by go-swagger; DO NOT EDIT.

package event_v2

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
	"github.com/go-openapi/swag"
)

// NewGetPublicEditHistoryParams creates a new GetPublicEditHistoryParams object
// with the default values initialized.
func NewGetPublicEditHistoryParams() *GetPublicEditHistoryParams {
	var ()
	return &GetPublicEditHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicEditHistoryParamsWithTimeout creates a new GetPublicEditHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicEditHistoryParamsWithTimeout(timeout time.Duration) *GetPublicEditHistoryParams {
	var ()
	return &GetPublicEditHistoryParams{

		timeout: timeout,
	}
}

// NewGetPublicEditHistoryParamsWithContext creates a new GetPublicEditHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicEditHistoryParamsWithContext(ctx context.Context) *GetPublicEditHistoryParams {
	var ()
	return &GetPublicEditHistoryParams{

		Context: ctx,
	}
}

// NewGetPublicEditHistoryParamsWithHTTPClient creates a new GetPublicEditHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicEditHistoryParamsWithHTTPClient(client *http.Client) *GetPublicEditHistoryParams {
	var ()
	return &GetPublicEditHistoryParams{
		HTTPClient: client,
	}
}

/*GetPublicEditHistoryParams contains all the parameters to send to the API endpoint
for the get public edit history operation typically these are written to a http.Request
*/
type GetPublicEditHistoryParams struct {

	/*EndDate
	  Ending date. e.g. 2015-03-20T12:27:06Z. Default : Current time in UTC

	*/
	EndDate *string
	/*Namespace
	  Namespace

	*/
	Namespace string
	/*Offset
	  Offset to query. Default : 0

	*/
	Offset *float64
	/*PageSize
	  Number of result in a page. Default : 100. Max : 500

	*/
	PageSize *float64
	/*StartDate
	  Starting date. e.g. 2015-03-20T12:27:06Z. Default : 1970-01-01T00:00:00Z

	*/
	StartDate *string
	/*Type
	  Edit History type to fetch

	*/
	Type *string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public edit history params
func (o *GetPublicEditHistoryParams) WithTimeout(timeout time.Duration) *GetPublicEditHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public edit history params
func (o *GetPublicEditHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public edit history params
func (o *GetPublicEditHistoryParams) WithContext(ctx context.Context) *GetPublicEditHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public edit history params
func (o *GetPublicEditHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public edit history params
func (o *GetPublicEditHistoryParams) WithHTTPClient(client *http.Client) *GetPublicEditHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public edit history params
func (o *GetPublicEditHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get public edit history params
func (o *GetPublicEditHistoryParams) WithEndDate(endDate *string) *GetPublicEditHistoryParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get public edit history params
func (o *GetPublicEditHistoryParams) SetEndDate(endDate *string) {
	o.EndDate = endDate
}

// WithNamespace adds the namespace to the get public edit history params
func (o *GetPublicEditHistoryParams) WithNamespace(namespace string) *GetPublicEditHistoryParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get public edit history params
func (o *GetPublicEditHistoryParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOffset adds the offset to the get public edit history params
func (o *GetPublicEditHistoryParams) WithOffset(offset *float64) *GetPublicEditHistoryParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get public edit history params
func (o *GetPublicEditHistoryParams) SetOffset(offset *float64) {
	o.Offset = offset
}

// WithPageSize adds the pageSize to the get public edit history params
func (o *GetPublicEditHistoryParams) WithPageSize(pageSize *float64) *GetPublicEditHistoryParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get public edit history params
func (o *GetPublicEditHistoryParams) SetPageSize(pageSize *float64) {
	o.PageSize = pageSize
}

// WithStartDate adds the startDate to the get public edit history params
func (o *GetPublicEditHistoryParams) WithStartDate(startDate *string) *GetPublicEditHistoryParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get public edit history params
func (o *GetPublicEditHistoryParams) SetStartDate(startDate *string) {
	o.StartDate = startDate
}

// WithType adds the typeVar to the get public edit history params
func (o *GetPublicEditHistoryParams) WithType(typeVar *string) *GetPublicEditHistoryParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get public edit history params
func (o *GetPublicEditHistoryParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUserID adds the userID to the get public edit history params
func (o *GetPublicEditHistoryParams) WithUserID(userID string) *GetPublicEditHistoryParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get public edit history params
func (o *GetPublicEditHistoryParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicEditHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate string
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset float64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatFloat64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize float64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatFloat64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate string
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

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

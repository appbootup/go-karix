// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWebhookParams creates a new GetWebhookParams object
// with the default values initialized.
func NewGetWebhookParams() *GetWebhookParams {
	var (
		aPIVersionDefault = string("1.0")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetWebhookParams{
		APIVersion: &aPIVersionDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebhookParamsWithTimeout creates a new GetWebhookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebhookParamsWithTimeout(timeout time.Duration) *GetWebhookParams {
	var (
		aPIVersionDefault = string("1.0")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetWebhookParams{
		APIVersion: &aPIVersionDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: timeout,
	}
}

// NewGetWebhookParamsWithContext creates a new GetWebhookParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebhookParamsWithContext(ctx context.Context) *GetWebhookParams {
	var (
		apiVersionDefault = string("1.0")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetWebhookParams{
		APIVersion: &apiVersionDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		Context: ctx,
	}
}

// NewGetWebhookParamsWithHTTPClient creates a new GetWebhookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebhookParamsWithHTTPClient(client *http.Client) *GetWebhookParams {
	var (
		apiVersionDefault = string("1.0")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetWebhookParams{
		APIVersion: &apiVersionDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetWebhookParams contains all the parameters to send to the API endpoint
for the get webhook operation typically these are written to a http.Request
*/
type GetWebhookParams struct {

	/*APIVersion
	  API Version. If not specified your pinned verison is used.

	*/
	APIVersion *string
	/*Limit
	  The numbers of items to return.

	*/
	Limit *int64
	/*Offset
	  The number of items to skip before starting to collect the result set.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get webhook params
func (o *GetWebhookParams) WithTimeout(timeout time.Duration) *GetWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webhook params
func (o *GetWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webhook params
func (o *GetWebhookParams) WithContext(ctx context.Context) *GetWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webhook params
func (o *GetWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webhook params
func (o *GetWebhookParams) WithHTTPClient(client *http.Client) *GetWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webhook params
func (o *GetWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get webhook params
func (o *GetWebhookParams) WithAPIVersion(aPIVersion *string) *GetWebhookParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get webhook params
func (o *GetWebhookParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithLimit adds the limit to the get webhook params
func (o *GetWebhookParams) WithLimit(limit *int64) *GetWebhookParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get webhook params
func (o *GetWebhookParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get webhook params
func (o *GetWebhookParams) WithOffset(offset *int64) *GetWebhookParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get webhook params
func (o *GetWebhookParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// header param api-version
		if err := r.SetHeaderParam("api-version", *o.APIVersion); err != nil {
			return err
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package number

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

// NewDeleteNumberNumParams creates a new DeleteNumberNumParams object
// with the default values initialized.
func NewDeleteNumberNumParams() *DeleteNumberNumParams {
	var (
		aPIVersionDefault = string("1.0")
	)
	return &DeleteNumberNumParams{
		APIVersion: &aPIVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNumberNumParamsWithTimeout creates a new DeleteNumberNumParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNumberNumParamsWithTimeout(timeout time.Duration) *DeleteNumberNumParams {
	var (
		aPIVersionDefault = string("1.0")
	)
	return &DeleteNumberNumParams{
		APIVersion: &aPIVersionDefault,

		timeout: timeout,
	}
}

// NewDeleteNumberNumParamsWithContext creates a new DeleteNumberNumParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNumberNumParamsWithContext(ctx context.Context) *DeleteNumberNumParams {
	var (
		apiVersionDefault = string("1.0")
	)
	return &DeleteNumberNumParams{
		APIVersion: &apiVersionDefault,

		Context: ctx,
	}
}

// NewDeleteNumberNumParamsWithHTTPClient creates a new DeleteNumberNumParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNumberNumParamsWithHTTPClient(client *http.Client) *DeleteNumberNumParams {
	var (
		apiVersionDefault = string("1.0")
	)
	return &DeleteNumberNumParams{
		APIVersion: &apiVersionDefault,
		HTTPClient: client,
	}
}

/*DeleteNumberNumParams contains all the parameters to send to the API endpoint
for the delete number num operation typically these are written to a http.Request
*/
type DeleteNumberNumParams struct {

	/*APIVersion
	  API Version. If not specified your pinned verison is used.

	*/
	APIVersion *string
	/*Num
	  Number which needs to be unrented

	*/
	Num int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete number num params
func (o *DeleteNumberNumParams) WithTimeout(timeout time.Duration) *DeleteNumberNumParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete number num params
func (o *DeleteNumberNumParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete number num params
func (o *DeleteNumberNumParams) WithContext(ctx context.Context) *DeleteNumberNumParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete number num params
func (o *DeleteNumberNumParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete number num params
func (o *DeleteNumberNumParams) WithHTTPClient(client *http.Client) *DeleteNumberNumParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete number num params
func (o *DeleteNumberNumParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete number num params
func (o *DeleteNumberNumParams) WithAPIVersion(aPIVersion *string) *DeleteNumberNumParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete number num params
func (o *DeleteNumberNumParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithNum adds the num to the delete number num params
func (o *DeleteNumberNumParams) WithNum(num int64) *DeleteNumberNumParams {
	o.SetNum(num)
	return o
}

// SetNum adds the num to the delete number num params
func (o *DeleteNumberNumParams) SetNum(num int64) {
	o.Num = num
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNumberNumParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param num
	if err := r.SetPathParam("num", swag.FormatInt64(o.Num)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

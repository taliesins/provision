package dhcp_subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETDhcpSubnetParams creates a new GETDhcpSubnetParams object
// with the default values initialized.
func NewGETDhcpSubnetParams() *GETDhcpSubnetParams {
	var ()
	return &GETDhcpSubnetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETDhcpSubnetParamsWithTimeout creates a new GETDhcpSubnetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETDhcpSubnetParamsWithTimeout(timeout time.Duration) *GETDhcpSubnetParams {
	var ()
	return &GETDhcpSubnetParams{

		timeout: timeout,
	}
}

// NewGETDhcpSubnetParamsWithContext creates a new GETDhcpSubnetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETDhcpSubnetParamsWithContext(ctx context.Context) *GETDhcpSubnetParams {
	var ()
	return &GETDhcpSubnetParams{

		Context: ctx,
	}
}

// NewGETDhcpSubnetParamsWithHTTPClient creates a new GETDhcpSubnetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETDhcpSubnetParamsWithHTTPClient(client *http.Client) *GETDhcpSubnetParams {
	var ()
	return &GETDhcpSubnetParams{
		HTTPClient: client,
	}
}

/*GETDhcpSubnetParams contains all the parameters to send to the API endpoint
for the g e t dhcp subnet operation typically these are written to a http.Request
*/
type GETDhcpSubnetParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) WithTimeout(timeout time.Duration) *GETDhcpSubnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) WithContext(ctx context.Context) *GETDhcpSubnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) WithHTTPClient(client *http.Client) *GETDhcpSubnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) WithID(id string) *GETDhcpSubnetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the g e t dhcp subnet params
func (o *GETDhcpSubnetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETDhcpSubnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
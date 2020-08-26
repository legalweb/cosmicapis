// Code generated by go-swagger; DO NOT EDIT.

package customer_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCustomerGatewayUpdateOwnershipParams creates a new CustomerGatewayUpdateOwnershipParams object
// with the default values initialized.
func NewCustomerGatewayUpdateOwnershipParams() *CustomerGatewayUpdateOwnershipParams {
	var ()
	return &CustomerGatewayUpdateOwnershipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayUpdateOwnershipParamsWithTimeout creates a new CustomerGatewayUpdateOwnershipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayUpdateOwnershipParamsWithTimeout(timeout time.Duration) *CustomerGatewayUpdateOwnershipParams {
	var ()
	return &CustomerGatewayUpdateOwnershipParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayUpdateOwnershipParamsWithContext creates a new CustomerGatewayUpdateOwnershipParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayUpdateOwnershipParamsWithContext(ctx context.Context) *CustomerGatewayUpdateOwnershipParams {
	var ()
	return &CustomerGatewayUpdateOwnershipParams{

		Context: ctx,
	}
}

// NewCustomerGatewayUpdateOwnershipParamsWithHTTPClient creates a new CustomerGatewayUpdateOwnershipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayUpdateOwnershipParamsWithHTTPClient(client *http.Client) *CustomerGatewayUpdateOwnershipParams {
	var ()
	return &CustomerGatewayUpdateOwnershipParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayUpdateOwnershipParams contains all the parameters to send to the API endpoint
for the customer gateway update ownership operation typically these are written to a http.Request
*/
type CustomerGatewayUpdateOwnershipParams struct {

	/*OwnershipID*/
	OwnershipID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) WithTimeout(timeout time.Duration) *CustomerGatewayUpdateOwnershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) WithContext(ctx context.Context) *CustomerGatewayUpdateOwnershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) WithHTTPClient(client *http.Client) *CustomerGatewayUpdateOwnershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwnershipID adds the ownershipID to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) WithOwnershipID(ownershipID string) *CustomerGatewayUpdateOwnershipParams {
	o.SetOwnershipID(ownershipID)
	return o
}

// SetOwnershipID adds the ownershipId to the customer gateway update ownership params
func (o *CustomerGatewayUpdateOwnershipParams) SetOwnershipID(ownershipID string) {
	o.OwnershipID = ownershipID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayUpdateOwnershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ownership_id
	if err := r.SetPathParam("ownership_id", o.OwnershipID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

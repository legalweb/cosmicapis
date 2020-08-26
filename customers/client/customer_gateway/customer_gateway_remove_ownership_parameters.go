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

// NewCustomerGatewayRemoveOwnershipParams creates a new CustomerGatewayRemoveOwnershipParams object
// with the default values initialized.
func NewCustomerGatewayRemoveOwnershipParams() *CustomerGatewayRemoveOwnershipParams {
	var ()
	return &CustomerGatewayRemoveOwnershipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerGatewayRemoveOwnershipParamsWithTimeout creates a new CustomerGatewayRemoveOwnershipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomerGatewayRemoveOwnershipParamsWithTimeout(timeout time.Duration) *CustomerGatewayRemoveOwnershipParams {
	var ()
	return &CustomerGatewayRemoveOwnershipParams{

		timeout: timeout,
	}
}

// NewCustomerGatewayRemoveOwnershipParamsWithContext creates a new CustomerGatewayRemoveOwnershipParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomerGatewayRemoveOwnershipParamsWithContext(ctx context.Context) *CustomerGatewayRemoveOwnershipParams {
	var ()
	return &CustomerGatewayRemoveOwnershipParams{

		Context: ctx,
	}
}

// NewCustomerGatewayRemoveOwnershipParamsWithHTTPClient creates a new CustomerGatewayRemoveOwnershipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomerGatewayRemoveOwnershipParamsWithHTTPClient(client *http.Client) *CustomerGatewayRemoveOwnershipParams {
	var ()
	return &CustomerGatewayRemoveOwnershipParams{
		HTTPClient: client,
	}
}

/*CustomerGatewayRemoveOwnershipParams contains all the parameters to send to the API endpoint
for the customer gateway remove ownership operation typically these are written to a http.Request
*/
type CustomerGatewayRemoveOwnershipParams struct {

	/*OwnershipID*/
	OwnershipID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) WithTimeout(timeout time.Duration) *CustomerGatewayRemoveOwnershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) WithContext(ctx context.Context) *CustomerGatewayRemoveOwnershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) WithHTTPClient(client *http.Client) *CustomerGatewayRemoveOwnershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwnershipID adds the ownershipID to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) WithOwnershipID(ownershipID string) *CustomerGatewayRemoveOwnershipParams {
	o.SetOwnershipID(ownershipID)
	return o
}

// SetOwnershipID adds the ownershipId to the customer gateway remove ownership params
func (o *CustomerGatewayRemoveOwnershipParams) SetOwnershipID(ownershipID string) {
	o.OwnershipID = ownershipID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerGatewayRemoveOwnershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

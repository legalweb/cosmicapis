// Code generated by go-swagger; DO NOT EDIT.

package identities

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

	"lwebco.de/cosmic-apis-spec/identity/models"
)

// NewAddOrganisationMemberParams creates a new AddOrganisationMemberParams object
// with the default values initialized.
func NewAddOrganisationMemberParams() *AddOrganisationMemberParams {
	var ()
	return &AddOrganisationMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOrganisationMemberParamsWithTimeout creates a new AddOrganisationMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOrganisationMemberParamsWithTimeout(timeout time.Duration) *AddOrganisationMemberParams {
	var ()
	return &AddOrganisationMemberParams{

		timeout: timeout,
	}
}

// NewAddOrganisationMemberParamsWithContext creates a new AddOrganisationMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOrganisationMemberParamsWithContext(ctx context.Context) *AddOrganisationMemberParams {
	var ()
	return &AddOrganisationMemberParams{

		Context: ctx,
	}
}

// NewAddOrganisationMemberParamsWithHTTPClient creates a new AddOrganisationMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOrganisationMemberParamsWithHTTPClient(client *http.Client) *AddOrganisationMemberParams {
	var ()
	return &AddOrganisationMemberParams{
		HTTPClient: client,
	}
}

/*AddOrganisationMemberParams contains all the parameters to send to the API endpoint
for the add organisation member operation typically these are written to a http.Request
*/
type AddOrganisationMemberParams struct {

	/*Body*/
	Body *models.V1AddOrganisationMemberRequest
	/*OrganisationID*/
	OrganisationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add organisation member params
func (o *AddOrganisationMemberParams) WithTimeout(timeout time.Duration) *AddOrganisationMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add organisation member params
func (o *AddOrganisationMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add organisation member params
func (o *AddOrganisationMemberParams) WithContext(ctx context.Context) *AddOrganisationMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add organisation member params
func (o *AddOrganisationMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add organisation member params
func (o *AddOrganisationMemberParams) WithHTTPClient(client *http.Client) *AddOrganisationMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add organisation member params
func (o *AddOrganisationMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add organisation member params
func (o *AddOrganisationMemberParams) WithBody(body *models.V1AddOrganisationMemberRequest) *AddOrganisationMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add organisation member params
func (o *AddOrganisationMemberParams) SetBody(body *models.V1AddOrganisationMemberRequest) {
	o.Body = body
}

// WithOrganisationID adds the organisationID to the add organisation member params
func (o *AddOrganisationMemberParams) WithOrganisationID(organisationID string) *AddOrganisationMemberParams {
	o.SetOrganisationID(organisationID)
	return o
}

// SetOrganisationID adds the organisationId to the add organisation member params
func (o *AddOrganisationMemberParams) SetOrganisationID(organisationID string) {
	o.OrganisationID = organisationID
}

// WriteToRequest writes these params to a swagger request
func (o *AddOrganisationMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organisation_id
	if err := r.SetPathParam("organisation_id", o.OrganisationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

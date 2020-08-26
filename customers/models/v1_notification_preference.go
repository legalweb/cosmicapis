// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NotificationPreference v1 notification preference
// swagger:model v1NotificationPreference
type V1NotificationPreference struct {

	// A notification category.
	// @todo(bweston92): list available notification categories.
	Category string `json:"category,omitempty"`

	// Whether or not the customer would like to receive
	// notifications via e-mail for this category.
	EnabledEmail bool `json:"enabled_email,omitempty"`

	// Whether or not the customer would like to receive
	// notifications via SMS for this category.
	EnabledSms bool `json:"enabled_sms,omitempty"`
}

// Validate validates this v1 notification preference
func (m *V1NotificationPreference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NotificationPreference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NotificationPreference) UnmarshalBinary(b []byte) error {
	var res V1NotificationPreference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssetGeneric @todo(bweston92): add a better phrase for this category
// this just prevents us from having an Asset category of
// Asset which it clearly is.
// swagger:model AssetGeneric
type AssetGeneric struct {

	// age in years
	AgeInYears int64 `json:"age_in_years,omitempty"`

	// chassis no
	ChassisNo string `json:"chassis_no,omitempty"`

	// electric
	Electric bool `json:"electric,omitempty"`

	// make
	Make string `json:"make,omitempty"`

	// model
	Model string `json:"model,omitempty"`

	// subtype
	Subtype string `json:"subtype,omitempty"`

	// type
	Type GenericGenericType `json:"type,omitempty"`
}

// Validate validates this asset generic
func (m *AssetGeneric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetGeneric) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetGeneric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetGeneric) UnmarshalBinary(b []byte) error {
	var res AssetGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

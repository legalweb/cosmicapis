// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Customersv1Asset customersv1 asset
// swagger:model customersv1Asset
type Customersv1Asset struct {

	// bank account
	BankAccount *AssetBankAccount `json:"bank_account,omitempty"`

	// generic
	Generic *AssetGeneric `json:"generic,omitempty"`

	// other
	Other *AssetOther `json:"other,omitempty"`

	// ownership percentage
	OwnershipPercentage float32 `json:"ownership_percentage,omitempty"`

	// pension
	Pension *AssetPension `json:"pension,omitempty"`

	// property
	Property *V1AssetProperty `json:"property,omitempty"`

	// shares
	Shares *AssetShares `json:"shares,omitempty"`

	// value
	Value *GoogletypeMoney `json:"value,omitempty"`
}

// Validate validates this customersv1 asset
func (m *Customersv1Asset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOther(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Customersv1Asset) validateBankAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BankAccount) { // not required
		return nil
	}

	if m.BankAccount != nil {
		if err := m.BankAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank_account")
			}
			return err
		}
	}

	return nil
}

func (m *Customersv1Asset) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(m.Generic) { // not required
		return nil
	}

	if m.Generic != nil {
		if err := m.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generic")
			}
			return err
		}
	}

	return nil
}

func (m *Customersv1Asset) validateOther(formats strfmt.Registry) error {

	if swag.IsZero(m.Other) { // not required
		return nil
	}

	if m.Other != nil {
		if err := m.Other.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other")
			}
			return err
		}
	}

	return nil
}

func (m *Customersv1Asset) validatePension(formats strfmt.Registry) error {

	if swag.IsZero(m.Pension) { // not required
		return nil
	}

	if m.Pension != nil {
		if err := m.Pension.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pension")
			}
			return err
		}
	}

	return nil
}

func (m *Customersv1Asset) validateProperty(formats strfmt.Registry) error {

	if swag.IsZero(m.Property) { // not required
		return nil
	}

	if m.Property != nil {
		if err := m.Property.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("property")
			}
			return err
		}
	}

	return nil
}

func (m *Customersv1Asset) validateShares(formats strfmt.Registry) error {

	if swag.IsZero(m.Shares) { // not required
		return nil
	}

	if m.Shares != nil {
		if err := m.Shares.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shares")
			}
			return err
		}
	}

	return nil
}

func (m *Customersv1Asset) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Customersv1Asset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customersv1Asset) UnmarshalBinary(b []byte) error {
	var res Customersv1Asset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

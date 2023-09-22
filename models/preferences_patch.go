// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreferencesPatch preferences patch
//
// swagger:model PreferencesPatch
type PreferencesPatch struct {

	// content languages
	ContentLanguages []string `json:"contentLanguages"`

	// ui languages
	UILanguages []string `json:"uiLanguages"`

	// visibility
	Visibility *VisibilityPreferencesPatch `json:"visibility,omitempty"`
}

// Validate validates this preferences patch
func (m *PreferencesPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreferencesPatch) validateVisibility(formats strfmt.Registry) error {
	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	if m.Visibility != nil {
		if err := m.Visibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this preferences patch based on the context it is used
func (m *PreferencesPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVisibility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreferencesPatch) contextValidateVisibility(ctx context.Context, formats strfmt.Registry) error {

	if m.Visibility != nil {

		if swag.IsZero(m.Visibility) { // not required
			return nil
		}

		if err := m.Visibility.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreferencesPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreferencesPatch) UnmarshalBinary(b []byte) error {
	var res PreferencesPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

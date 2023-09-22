// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrgWrite org write
//
// swagger:model OrgWrite
type OrgWrite struct {

	// active status
	// Required: true
	ActiveStatus struct {
		OrgActiveStatus
	} `json:"activeStatus"`

	// billing address
	BillingAddress *Address `json:"billingAddress,omitempty"`

	// district name
	// Required: true
	// Max Length: 255
	DistrictName *string `json:"districtName"`

	// name
	// Required: true
	// Max Length: 255
	Name *string `json:"name"`

	// postal address
	PostalAddress *Address `json:"postalAddress,omitempty"`

	// type
	// Required: true
	Type struct {
		OrgType
	} `json:"type"`

	// visiting address
	VisitingAddress *Address `json:"visitingAddress,omitempty"`
}

// Validate validates this org write
func (m *OrgWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistrictName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisitingAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrgWrite) validateActiveStatus(formats strfmt.Registry) error {

	return nil
}

func (m *OrgWrite) validateBillingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingAddress) { // not required
		return nil
	}

	if m.BillingAddress != nil {
		if err := m.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OrgWrite) validateDistrictName(formats strfmt.Registry) error {

	if err := validate.Required("districtName", "body", m.DistrictName); err != nil {
		return err
	}

	if err := validate.MaxLength("districtName", "body", *m.DistrictName, 255); err != nil {
		return err
	}

	return nil
}

func (m *OrgWrite) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *OrgWrite) validatePostalAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.PostalAddress) { // not required
		return nil
	}

	if m.PostalAddress != nil {
		if err := m.PostalAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postalAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OrgWrite) validateType(formats strfmt.Registry) error {

	return nil
}

func (m *OrgWrite) validateVisitingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.VisitingAddress) { // not required
		return nil
	}

	if m.VisitingAddress != nil {
		if err := m.VisitingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visitingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visitingAddress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this org write based on the context it is used
func (m *OrgWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActiveStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBillingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostalAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVisitingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrgWrite) contextValidateActiveStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *OrgWrite) contextValidateBillingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.BillingAddress != nil {

		if swag.IsZero(m.BillingAddress) { // not required
			return nil
		}

		if err := m.BillingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billingAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OrgWrite) contextValidatePostalAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.PostalAddress != nil {

		if swag.IsZero(m.PostalAddress) { // not required
			return nil
		}

		if err := m.PostalAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postalAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OrgWrite) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *OrgWrite) contextValidateVisitingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.VisitingAddress != nil {

		if swag.IsZero(m.VisitingAddress) { // not required
			return nil
		}

		if err := m.VisitingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visitingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("visitingAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrgWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrgWrite) UnmarshalBinary(b []byte) error {
	var res OrgWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

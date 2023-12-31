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

// Org org
//
// swagger:model Org
type Org struct {

	// active status
	// Required: true
	ActiveStatus OrgActiveStatus `json:"activeStatus"`

	// billing address
	BillingAddress *Address `json:"billingAddress,omitempty"`

	// district name
	// Required: true
	// Max Length: 255
	DistrictName *string `json:"districtName"`

	// last changed date
	// Required: true
	// Format: date-time
	LastChangedDate *strfmt.DateTime `json:"lastChangedDate"`

	// name
	// Required: true
	// Max Length: 255
	Name *string `json:"name"`

	// org ID
	// Required: true
	OrgID *int64 `json:"orgID"`

	// postal address
	PostalAddress *Address `json:"postalAddress,omitempty"`

	// type
	// Required: true
	Type OrgType `json:"type"`

	// uid
	// Required: true
	// Format: uuid
	UID *strfmt.UUID `json:"uid"`

	// visiting address
	VisitingAddress *Address `json:"visitingAddress,omitempty"`
}

// Validate validates this org
func (m *Org) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLastChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
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

func (m *Org) validateActiveStatus(formats strfmt.Registry) error {

	return nil
}

func (m *Org) validateBillingAddress(formats strfmt.Registry) error {
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

func (m *Org) validateDistrictName(formats strfmt.Registry) error {

	if err := validate.Required("districtName", "body", m.DistrictName); err != nil {
		return err
	}

	if err := validate.MaxLength("districtName", "body", *m.DistrictName, 255); err != nil {
		return err
	}

	return nil
}

func (m *Org) validateLastChangedDate(formats strfmt.Registry) error {

	if err := validate.Required("lastChangedDate", "body", m.LastChangedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastChangedDate", "body", "date-time", m.LastChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Org) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *Org) validateOrgID(formats strfmt.Registry) error {

	if err := validate.Required("orgID", "body", m.OrgID); err != nil {
		return err
	}

	return nil
}

func (m *Org) validatePostalAddress(formats strfmt.Registry) error {
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

func (m *Org) validateType(formats strfmt.Registry) error {

	return nil
}

func (m *Org) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Org) validateVisitingAddress(formats strfmt.Registry) error {
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

// ContextValidate validate this org based on the context it is used
func (m *Org) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *Org) contextValidateActiveStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Org) contextValidateBillingAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Org) contextValidatePostalAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Org) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Org) contextValidateVisitingAddress(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Org) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Org) UnmarshalBinary(b []byte) error {
	var res Org
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

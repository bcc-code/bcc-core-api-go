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

// RoleAssignment role assignment
//
// swagger:model RoleAssignment
type RoleAssignment struct {

	// last changed date
	// Required: true
	// Format: date-time
	LastChangedDate *strfmt.DateTime `json:"lastChangedDate"`

	// org
	Org *Org `json:"org,omitempty"`

	// org Uid
	// Format: uuid
	OrgUID strfmt.UUID `json:"orgUid,omitempty"`

	// person
	Person *Person `json:"person,omitempty"`

	// person Uid
	// Required: true
	// Format: uuid
	PersonUID *strfmt.UUID `json:"personUid"`

	// role
	Role *Role `json:"role,omitempty"`

	// role Uid
	// Required: true
	// Format: uuid
	RoleUID *strfmt.UUID `json:"roleUid"`

	// uid
	// Required: true
	// Format: uuid
	UID *strfmt.UUID `json:"uid"`

	// valid from
	// Required: true
	// Format: date-time
	ValidFrom *strfmt.DateTime `json:"validFrom"`

	// valid to
	// Format: date-time
	ValidTo *strfmt.DateTime `json:"validTo,omitempty"`
}

// Validate validates this role assignment
func (m *RoleAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerson(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleAssignment) validateLastChangedDate(formats strfmt.Registry) error {

	if err := validate.Required("lastChangedDate", "body", m.LastChangedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastChangedDate", "body", "date-time", m.LastChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoleAssignment) validateOrg(formats strfmt.Registry) error {
	if swag.IsZero(m.Org) { // not required
		return nil
	}

	if m.Org != nil {
		if err := m.Org.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

func (m *RoleAssignment) validateOrgUID(formats strfmt.Registry) error {
	if swag.IsZero(m.OrgUID) { // not required
		return nil
	}

	if err := validate.FormatOf("orgUid", "body", "uuid", m.OrgUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoleAssignment) validatePerson(formats strfmt.Registry) error {
	if swag.IsZero(m.Person) { // not required
		return nil
	}

	if m.Person != nil {
		if err := m.Person.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("person")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("person")
			}
			return err
		}
	}

	return nil
}

func (m *RoleAssignment) validatePersonUID(formats strfmt.Registry) error {

	if err := validate.Required("personUid", "body", m.PersonUID); err != nil {
		return err
	}

	if err := validate.FormatOf("personUid", "body", "uuid", m.PersonUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoleAssignment) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *RoleAssignment) validateRoleUID(formats strfmt.Registry) error {

	if err := validate.Required("roleUid", "body", m.RoleUID); err != nil {
		return err
	}

	if err := validate.FormatOf("roleUid", "body", "uuid", m.RoleUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoleAssignment) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoleAssignment) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoleAssignment) validateValidTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this role assignment based on the context it is used
func (m *RoleAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerson(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleAssignment) contextValidateOrg(ctx context.Context, formats strfmt.Registry) error {

	if m.Org != nil {

		if swag.IsZero(m.Org) { // not required
			return nil
		}

		if err := m.Org.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

func (m *RoleAssignment) contextValidatePerson(ctx context.Context, formats strfmt.Registry) error {

	if m.Person != nil {

		if swag.IsZero(m.Person) { // not required
			return nil
		}

		if err := m.Person.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("person")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("person")
			}
			return err
		}
	}

	return nil
}

func (m *RoleAssignment) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {

		if swag.IsZero(m.Role) { // not required
			return nil
		}

		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleAssignment) UnmarshalBinary(b []byte) error {
	var res RoleAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// Consent consent
//
// swagger:model Consent
type Consent struct {

	// last changed date
	// Required: true
	// Format: date-time
	LastChangedDate *strfmt.DateTime `json:"lastChangedDate"`

	// org
	Org *Org `json:"org,omitempty"`

	// org Uid
	// Required: true
	// Format: uuid
	OrgUID *strfmt.UUID `json:"orgUid"`

	// person
	Person *Person `json:"person,omitempty"`

	// person Uid
	// Required: true
	// Format: uuid
	PersonUID *strfmt.UUID `json:"personUid"`

	// purpose
	// Required: true
	Purpose *ConsentPurpose `json:"purpose"`

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

// Validate validates this consent
func (m *Consent) Validate(formats strfmt.Registry) error {
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

	if err := m.validatePurpose(formats); err != nil {
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

func (m *Consent) validateLastChangedDate(formats strfmt.Registry) error {

	if err := validate.Required("lastChangedDate", "body", m.LastChangedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastChangedDate", "body", "date-time", m.LastChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Consent) validateOrg(formats strfmt.Registry) error {
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

func (m *Consent) validateOrgUID(formats strfmt.Registry) error {

	if err := validate.Required("orgUid", "body", m.OrgUID); err != nil {
		return err
	}

	if err := validate.FormatOf("orgUid", "body", "uuid", m.OrgUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Consent) validatePerson(formats strfmt.Registry) error {
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

func (m *Consent) validatePersonUID(formats strfmt.Registry) error {

	if err := validate.Required("personUid", "body", m.PersonUID); err != nil {
		return err
	}

	if err := validate.FormatOf("personUid", "body", "uuid", m.PersonUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Consent) validatePurpose(formats strfmt.Registry) error {

	if err := validate.Required("purpose", "body", m.Purpose); err != nil {
		return err
	}

	if err := validate.Required("purpose", "body", m.Purpose); err != nil {
		return err
	}

	if m.Purpose != nil {
		if err := m.Purpose.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purpose")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("purpose")
			}
			return err
		}
	}

	return nil
}

func (m *Consent) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Consent) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Consent) validateValidTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this consent based on the context it is used
func (m *Consent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerson(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePurpose(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Consent) contextValidateOrg(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Consent) contextValidatePerson(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Consent) contextValidatePurpose(ctx context.Context, formats strfmt.Registry) error {

	if m.Purpose != nil {

		if err := m.Purpose.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purpose")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("purpose")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Consent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Consent) UnmarshalBinary(b []byte) error {
	var res Consent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
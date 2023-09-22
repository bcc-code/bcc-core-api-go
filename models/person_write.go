// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PersonWrite person write
//
// swagger:model PersonWrite
type PersonWrite struct {

	// address
	Address *Address `json:"address,omitempty"`

	// birth date
	// Required: true
	// Format: date
	BirthDate *strfmt.Date `json:"birthDate"`

	// cell phone
	CellPhone string `json:"cellPhone,omitempty"`

	// cell phone verified
	CellPhoneVerified bool `json:"cellPhoneVerified,omitempty"`

	// deceased date
	// Format: date
	DeceasedDate *strfmt.Date `json:"deceasedDate,omitempty"`

	// display name
	// Required: true
	// Max Length: 255
	DisplayName *string `json:"displayName"`

	// email
	Email string `json:"email,omitempty"`

	// email verified
	EmailVerified bool `json:"emailVerified,omitempty"`

	// first name
	// Required: true
	// Max Length: 255
	FirstName *string `json:"firstName"`

	// gender
	// Required: true
	Gender *Gender `json:"gender"`

	// home phone
	HomePhone string `json:"homePhone,omitempty"`

	// last name
	// Required: true
	// Max Length: 255
	LastName *string `json:"lastName"`

	// marital status
	// Required: true
	MaritalStatus *MaritalStatus `json:"maritalStatus"`

	// middle name
	// Max Length: 255
	MiddleName string `json:"middleName,omitempty"`

	// national ids
	NationalIds []*NationalID `json:"nationalIds"`

	// preferences
	Preferences *Preferences `json:"preferences,omitempty"`

	// URL of person's profile picture
	// Max Length: 512
	ProfilePicture string `json:"profilePicture,omitempty"`
}

// Validate validates this person write
func (m *PersonWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeceasedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaritalStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMiddleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNationalIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfilePicture(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonWrite) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) validateBirthDate(formats strfmt.Registry) error {

	if err := validate.Required("birthDate", "body", m.BirthDate); err != nil {
		return err
	}

	if err := validate.FormatOf("birthDate", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonWrite) validateDeceasedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DeceasedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deceasedDate", "body", "date", m.DeceasedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PersonWrite) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", *m.DisplayName, 255); err != nil {
		return err
	}

	return nil
}

func (m *PersonWrite) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MaxLength("firstName", "body", *m.FirstName, 255); err != nil {
		return err
	}

	return nil
}

func (m *PersonWrite) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}

	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}

	if m.Gender != nil {
		if err := m.Gender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gender")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	if err := validate.MaxLength("lastName", "body", *m.LastName, 255); err != nil {
		return err
	}

	return nil
}

func (m *PersonWrite) validateMaritalStatus(formats strfmt.Registry) error {

	if err := validate.Required("maritalStatus", "body", m.MaritalStatus); err != nil {
		return err
	}

	if err := validate.Required("maritalStatus", "body", m.MaritalStatus); err != nil {
		return err
	}

	if m.MaritalStatus != nil {
		if err := m.MaritalStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maritalStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maritalStatus")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) validateMiddleName(formats strfmt.Registry) error {
	if swag.IsZero(m.MiddleName) { // not required
		return nil
	}

	if err := validate.MaxLength("middleName", "body", m.MiddleName, 255); err != nil {
		return err
	}

	return nil
}

func (m *PersonWrite) validateNationalIds(formats strfmt.Registry) error {
	if swag.IsZero(m.NationalIds) { // not required
		return nil
	}

	for i := 0; i < len(m.NationalIds); i++ {
		if swag.IsZero(m.NationalIds[i]) { // not required
			continue
		}

		if m.NationalIds[i] != nil {
			if err := m.NationalIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nationalIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nationalIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PersonWrite) validatePreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	if m.Preferences != nil {
		if err := m.Preferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) validateProfilePicture(formats strfmt.Registry) error {
	if swag.IsZero(m.ProfilePicture) { // not required
		return nil
	}

	if err := validate.MaxLength("profilePicture", "body", m.ProfilePicture, 512); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this person write based on the context it is used
func (m *PersonWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaritalStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNationalIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonWrite) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {

		if swag.IsZero(m.Address) { // not required
			return nil
		}

		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) contextValidateGender(ctx context.Context, formats strfmt.Registry) error {

	if m.Gender != nil {

		if err := m.Gender.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gender")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gender")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) contextValidateMaritalStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.MaritalStatus != nil {

		if err := m.MaritalStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maritalStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maritalStatus")
			}
			return err
		}
	}

	return nil
}

func (m *PersonWrite) contextValidateNationalIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NationalIds); i++ {

		if m.NationalIds[i] != nil {

			if swag.IsZero(m.NationalIds[i]) { // not required
				return nil
			}

			if err := m.NationalIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nationalIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nationalIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PersonWrite) contextValidatePreferences(ctx context.Context, formats strfmt.Registry) error {

	if m.Preferences != nil {

		if swag.IsZero(m.Preferences) { // not required
			return nil
		}

		if err := m.Preferences.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preferences")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PersonWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonWrite) UnmarshalBinary(b []byte) error {
	var res PersonWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

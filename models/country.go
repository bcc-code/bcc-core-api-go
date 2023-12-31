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

// Country country
//
// swagger:model Country
type Country struct {

	// iso2 code
	// Required: true
	Iso2Code *string `json:"iso2Code"`

	// last changed date
	// Required: true
	// Format: date-time
	LastChangedDate *strfmt.DateTime `json:"lastChangedDate"`

	// name en
	// Required: true
	// Max Length: 255
	NameEn *string `json:"nameEn"`

	// name native
	// Required: true
	// Max Length: 255
	NameNative *string `json:"nameNative"`

	// name no
	// Required: true
	// Max Length: 255
	NameNo *string `json:"nameNo"`

	// uid
	// Required: true
	// Format: uuid
	UID *strfmt.UUID `json:"uid"`
}

// Validate validates this country
func (m *Country) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIso2Code(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameEn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameNative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Country) validateIso2Code(formats strfmt.Registry) error {

	if err := validate.Required("iso2Code", "body", m.Iso2Code); err != nil {
		return err
	}

	return nil
}

func (m *Country) validateLastChangedDate(formats strfmt.Registry) error {

	if err := validate.Required("lastChangedDate", "body", m.LastChangedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastChangedDate", "body", "date-time", m.LastChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Country) validateNameEn(formats strfmt.Registry) error {

	if err := validate.Required("nameEn", "body", m.NameEn); err != nil {
		return err
	}

	if err := validate.MaxLength("nameEn", "body", *m.NameEn, 255); err != nil {
		return err
	}

	return nil
}

func (m *Country) validateNameNative(formats strfmt.Registry) error {

	if err := validate.Required("nameNative", "body", m.NameNative); err != nil {
		return err
	}

	if err := validate.MaxLength("nameNative", "body", *m.NameNative, 255); err != nil {
		return err
	}

	return nil
}

func (m *Country) validateNameNo(formats strfmt.Registry) error {

	if err := validate.Required("nameNo", "body", m.NameNo); err != nil {
		return err
	}

	if err := validate.MaxLength("nameNo", "body", *m.NameNo, 255); err != nil {
		return err
	}

	return nil
}

func (m *Country) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this country based on context it is used
func (m *Country) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Country) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Country) UnmarshalBinary(b []byte) error {
	var res Country
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

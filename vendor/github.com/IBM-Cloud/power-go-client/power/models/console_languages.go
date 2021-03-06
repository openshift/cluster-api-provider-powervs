// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsoleLanguages console languages
// swagger:model ConsoleLanguages
type ConsoleLanguages struct {

	// console languages
	// Required: true
	ConsoleLanguages []*ConsoleLanguage `json:"consoleLanguages"`
}

// Validate validates this console languages
func (m *ConsoleLanguages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsoleLanguages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleLanguages) validateConsoleLanguages(formats strfmt.Registry) error {

	if err := validate.Required("consoleLanguages", "body", m.ConsoleLanguages); err != nil {
		return err
	}

	for i := 0; i < len(m.ConsoleLanguages); i++ {
		if swag.IsZero(m.ConsoleLanguages[i]) { // not required
			continue
		}

		if m.ConsoleLanguages[i] != nil {
			if err := m.ConsoleLanguages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consoleLanguages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleLanguages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleLanguages) UnmarshalBinary(b []byte) error {
	var res ConsoleLanguages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

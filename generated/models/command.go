// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Command command
//
// swagger:model Command
type Command struct {

	// arguments
	// Required: true
	Arguments []interface{} `json:"arguments"`

	// capability
	// Required: true
	// Max Length: 30
	// Pattern: ^st\.
	Capability *string `json:"capability"`

	// command
	// Required: true
	// Max Length: 30
	Command *string `json:"command"`

	// component
	// Required: true
	// Max Length: 30
	Component *string `json:"component"`
}

// Validate validates this command
func (m *Command) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Command) validateArguments(formats strfmt.Registry) error {

	if err := validate.Required("arguments", "body", m.Arguments); err != nil {
		return err
	}

	return nil
}

func (m *Command) validateCapability(formats strfmt.Registry) error {

	if err := validate.Required("capability", "body", m.Capability); err != nil {
		return err
	}

	if err := validate.MaxLength("capability", "body", string(*m.Capability), 30); err != nil {
		return err
	}

	if err := validate.Pattern("capability", "body", string(*m.Capability), `^st\.`); err != nil {
		return err
	}

	return nil
}

func (m *Command) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	if err := validate.MaxLength("command", "body", string(*m.Command), 30); err != nil {
		return err
	}

	return nil
}

func (m *Command) validateComponent(formats strfmt.Registry) error {

	if err := validate.Required("component", "body", m.Component); err != nil {
		return err
	}

	if err := validate.MaxLength("component", "body", string(*m.Component), 30); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Command) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Command) UnmarshalBinary(b []byte) error {
	var res Command
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

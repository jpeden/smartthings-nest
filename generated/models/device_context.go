// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceContext device context
//
// swagger:model DeviceContext
type DeviceContext struct {

	// categories
	Categories []string `json:"categories"`

	// groups
	// Required: true
	Groups []string `json:"groups"`

	// room name
	// Required: true
	// Max Length: 100
	RoomName *string `json:"roomName"`
}

// Validate validates this device context
func (m *DeviceContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoomName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceContext) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {

		if err := validate.MaxLength("categories"+"."+strconv.Itoa(i), "body", string(m.Categories[i]), 100); err != nil {
			return err
		}

	}

	return nil
}

func (m *DeviceContext) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {

		if err := validate.MaxLength("groups"+"."+strconv.Itoa(i), "body", string(m.Groups[i]), 100); err != nil {
			return err
		}

	}

	return nil
}

func (m *DeviceContext) validateRoomName(formats strfmt.Registry) error {

	if err := validate.Required("roomName", "body", m.RoomName); err != nil {
		return err
	}

	if err := validate.MaxLength("roomName", "body", string(*m.RoomName), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceContext) UnmarshalBinary(b []byte) error {
	var res DeviceContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

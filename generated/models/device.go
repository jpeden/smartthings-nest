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

// Device device
//
// swagger:model Device
type Device struct {

	// device context
	DeviceContext *DeviceContext `json:"deviceContext,omitempty"`

	// device cookie
	DeviceCookie Cookie `json:"deviceCookie,omitempty"`

	// device handler type
	// Max Length: 30
	DeviceHandlerType string `json:"deviceHandlerType,omitempty"`

	// device unique Id
	// Max Length: 30
	DeviceUniqueID string `json:"deviceUniqueId,omitempty"`

	// external device Id
	// Max Length: 256
	ExternalDeviceID string `json:"externalDeviceId,omitempty"`

	// friendly name
	// Max Length: 100
	FriendlyName string `json:"friendlyName,omitempty"`

	// manufacturer info
	ManufacturerInfo *Manufacturer `json:"manufacturerInfo,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceHandlerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceUniqueID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFriendlyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturerInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateDeviceContext(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceContext) { // not required
		return nil
	}

	if m.DeviceContext != nil {
		if err := m.DeviceContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceContext")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateDeviceCookie(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceCookie) { // not required
		return nil
	}

	if err := m.DeviceCookie.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deviceCookie")
		}
		return err
	}

	return nil
}

func (m *Device) validateDeviceHandlerType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceHandlerType) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceHandlerType", "body", string(m.DeviceHandlerType), 30); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateDeviceUniqueID(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceUniqueID) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceUniqueId", "body", string(m.DeviceUniqueID), 30); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateExternalDeviceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDeviceID) { // not required
		return nil
	}

	if err := validate.MaxLength("externalDeviceId", "body", string(m.ExternalDeviceID), 256); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateFriendlyName(formats strfmt.Registry) error {

	if swag.IsZero(m.FriendlyName) { // not required
		return nil
	}

	if err := validate.MaxLength("friendlyName", "body", string(m.FriendlyName), 100); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateManufacturerInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ManufacturerInfo) { // not required
		return nil
	}

	if m.ManufacturerInfo != nil {
		if err := m.ManufacturerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturerInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

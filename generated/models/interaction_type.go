// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// InteractionType interaction type
//
// swagger:model InteractionType
type InteractionType string

const (

	// InteractionTypeDiscoveryRequest captures enum value "discoveryRequest"
	InteractionTypeDiscoveryRequest InteractionType = "discoveryRequest"

	// InteractionTypeDiscoveryResponse captures enum value "discoveryResponse"
	InteractionTypeDiscoveryResponse InteractionType = "discoveryResponse"

	// InteractionTypeStateRefreshRequest captures enum value "stateRefreshRequest"
	InteractionTypeStateRefreshRequest InteractionType = "stateRefreshRequest"

	// InteractionTypeStateRefreshResponse captures enum value "stateRefreshResponse"
	InteractionTypeStateRefreshResponse InteractionType = "stateRefreshResponse"

	// InteractionTypeCommandRequest captures enum value "commandRequest"
	InteractionTypeCommandRequest InteractionType = "commandRequest"

	// InteractionTypeCommandResponse captures enum value "commandResponse"
	InteractionTypeCommandResponse InteractionType = "commandResponse"

	// InteractionTypeGrantCallbackAccess captures enum value "grantCallbackAccess"
	InteractionTypeGrantCallbackAccess InteractionType = "grantCallbackAccess"

	// InteractionTypeAccessTokenRequest captures enum value "accessTokenRequest"
	InteractionTypeAccessTokenRequest InteractionType = "accessTokenRequest"

	// InteractionTypeAccessTokenResponse captures enum value "accessTokenResponse"
	InteractionTypeAccessTokenResponse InteractionType = "accessTokenResponse"

	// InteractionTypeRefreshAccessTokens captures enum value "refreshAccessTokens"
	InteractionTypeRefreshAccessTokens InteractionType = "refreshAccessTokens"

	// InteractionTypeStateCallback captures enum value "stateCallback"
	InteractionTypeStateCallback InteractionType = "stateCallback"

	// InteractionTypeDiscoveryCallback captures enum value "discoveryCallback"
	InteractionTypeDiscoveryCallback InteractionType = "discoveryCallback"

	// InteractionTypeInteractionResult captures enum value "interactionResult"
	InteractionTypeInteractionResult InteractionType = "interactionResult"

	// InteractionTypeIntegrationDeleted captures enum value "integrationDeleted"
	InteractionTypeIntegrationDeleted InteractionType = "integrationDeleted"
)

// for schema
var interactionTypeEnum []interface{}

func init() {
	var res []InteractionType
	if err := json.Unmarshal([]byte(`["discoveryRequest","discoveryResponse","stateRefreshRequest","stateRefreshResponse","commandRequest","commandResponse","grantCallbackAccess","accessTokenRequest","accessTokenResponse","refreshAccessTokens","stateCallback","discoveryCallback","interactionResult","integrationDeleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionTypeEnum = append(interactionTypeEnum, v)
	}
}

func (m InteractionType) validateInteractionTypeEnum(path, location string, value InteractionType) error {
	if err := validate.EnumCase(path, location, value, interactionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this interaction type
func (m InteractionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInteractionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
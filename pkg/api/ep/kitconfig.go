// Code generated by go-swagger; DO NOT EDIT.

//
//   Copyright (c) 2022 Intel Corporation.
//
//   SPDX-License-Identifier: Apache-2.0
//
//
//

package ep

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Kitconfig kitconfig
//
// swagger:model kitconfig
type Kitconfig struct {

	// cluster
	Cluster *KitconfigCluster `json:"Cluster,omitempty"`

	// components
	Components *KitconfigComponents `json:"Components,omitempty"`

	// o s
	OS *KitconfigOS `json:"OS,omitempty"`

	// parameters
	Parameters *KitconfigParameters `json:"Parameters,omitempty"`

	// use
	Use []string `json:"Use"`
}

// Validate validates this kitconfig
func (m *Kitconfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Kitconfig) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cluster")
			}
			return err
		}
	}

	return nil
}

func (m *Kitconfig) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	if m.Components != nil {
		if err := m.Components.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Components")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Components")
			}
			return err
		}
	}

	return nil
}

func (m *Kitconfig) validateOS(formats strfmt.Registry) error {
	if swag.IsZero(m.OS) { // not required
		return nil
	}

	if m.OS != nil {
		if err := m.OS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OS")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OS")
			}
			return err
		}
	}

	return nil
}

func (m *Kitconfig) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Parameters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kitconfig based on the context it is used
func (m *Kitconfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Kitconfig) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cluster")
			}
			return err
		}
	}

	return nil
}

func (m *Kitconfig) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.Components != nil {
		if err := m.Components.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Components")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Components")
			}
			return err
		}
	}

	return nil
}

func (m *Kitconfig) contextValidateOS(ctx context.Context, formats strfmt.Registry) error {

	if m.OS != nil {
		if err := m.OS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OS")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OS")
			}
			return err
		}
	}

	return nil
}

func (m *Kitconfig) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Kitconfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Kitconfig) UnmarshalBinary(b []byte) error {
	var res Kitconfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KitconfigCluster kitconfig cluster
//
// swagger:model KitconfigCluster
type KitconfigCluster struct {

	// config
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Config string `json:"config,omitempty"`

	// export config folder
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	ExportConfigFolder string `json:"export_config_folder,omitempty"`

	// manifests
	Manifests []string `json:"manifests"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this kitconfig cluster
func (m *KitconfigCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportConfigFolder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigCluster) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if err := validate.Pattern("Cluster"+"."+"config", "body", m.Config, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigCluster) validateExportConfigFolder(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportConfigFolder) { // not required
		return nil
	}

	if err := validate.Pattern("Cluster"+"."+"export_config_folder", "body", m.ExportConfigFolder, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigCluster) validateManifests(formats strfmt.Registry) error {
	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {

		if err := validate.Pattern("Cluster"+"."+"manifests"+"."+strconv.Itoa(i), "body", m.Manifests[i], `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this kitconfig cluster based on context it is used
func (m *KitconfigCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KitconfigCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KitconfigCluster) UnmarshalBinary(b []byte) error {
	var res KitconfigCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KitconfigComponents kitconfig components
//
// swagger:model KitconfigComponents
type KitconfigComponents struct {

	// manifests
	Manifests []string `json:"manifests"`

	// selector
	Selector []*KitconfigComponentsSelectorItems0 `json:"selector"`
}

// Validate validates this kitconfig components
func (m *KitconfigComponents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigComponents) validateManifests(formats strfmt.Registry) error {
	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {

		if err := validate.Pattern("Components"+"."+"manifests"+"."+strconv.Itoa(i), "body", m.Manifests[i], `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *KitconfigComponents) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	for i := 0; i < len(m.Selector); i++ {
		if swag.IsZero(m.Selector[i]) { // not required
			continue
		}

		if m.Selector[i] != nil {
			if err := m.Selector[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Components" + "." + "selector" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Components" + "." + "selector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this kitconfig components based on the context it is used
func (m *KitconfigComponents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigComponents) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Selector); i++ {

		if m.Selector[i] != nil {
			if err := m.Selector[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Components" + "." + "selector" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Components" + "." + "selector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KitconfigComponents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KitconfigComponents) UnmarshalBinary(b []byte) error {
	var res KitconfigComponents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KitconfigComponentsSelectorItems0 kitconfig components selector items0
//
// swagger:model KitconfigComponentsSelectorItems0
type KitconfigComponentsSelectorItems0 struct {

	// override yaml
	OverrideYaml string `json:"__override_yaml__,omitempty"`

	// name
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this kitconfig components selector items0
func (m *KitconfigComponentsSelectorItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigComponentsSelectorItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kitconfig components selector items0 based on context it is used
func (m *KitconfigComponentsSelectorItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KitconfigComponentsSelectorItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KitconfigComponentsSelectorItems0) UnmarshalBinary(b []byte) error {
	var res KitconfigComponentsSelectorItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KitconfigOS kitconfig o s
//
// swagger:model KitconfigOS
type KitconfigOS struct {

	// config
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Config string `json:"config,omitempty"`

	// distro
	// Pattern: ^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$
	Distro string `json:"distro,omitempty"`

	// manifests
	Manifests []string `json:"manifests"`

	// provider
	// Enum: [esp none]
	Provider string `json:"provider,omitempty"`
}

// Validate validates this kitconfig o s
func (m *KitconfigOS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistro(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigOS) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if err := validate.Pattern("OS"+"."+"config", "body", m.Config, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigOS) validateDistro(formats strfmt.Registry) error {
	if swag.IsZero(m.Distro) { // not required
		return nil
	}

	if err := validate.Pattern("OS"+"."+"distro", "body", m.Distro, `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigOS) validateManifests(formats strfmt.Registry) error {
	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {

		if err := validate.Pattern("OS"+"."+"manifests"+"."+strconv.Itoa(i), "body", m.Manifests[i], `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
			return err
		}

	}

	return nil
}

var kitconfigOSTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["esp","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kitconfigOSTypeProviderPropEnum = append(kitconfigOSTypeProviderPropEnum, v)
	}
}

const (

	// KitconfigOSProviderEsp captures enum value "esp"
	KitconfigOSProviderEsp string = "esp"

	// KitconfigOSProviderNone captures enum value "none"
	KitconfigOSProviderNone string = "none"
)

// prop value enum
func (m *KitconfigOS) validateProviderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kitconfigOSTypeProviderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KitconfigOS) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("OS"+"."+"provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kitconfig o s based on context it is used
func (m *KitconfigOS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KitconfigOS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KitconfigOS) UnmarshalBinary(b []byte) error {
	var res KitconfigOS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KitconfigParameters kitconfig parameters
//
// swagger:model KitconfigParameters
type KitconfigParameters struct {

	// customconfig
	Customconfig *Customconfig `json:"customconfig,omitempty"`

	// default ssh key path
	DefaultSSHKeyPath string `json:"default_ssh_key_path,omitempty"`

	// extensions
	Extensions []string `json:"extensions"`

	// global settings
	GlobalSettings *KitconfigParametersGlobalSettings `json:"global_settings,omitempty"`

	// nodes
	Nodes []*Node `json:"nodes"`
}

// Validate validates this kitconfig parameters
func (m *KitconfigParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomconfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigParameters) validateCustomconfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Customconfig) { // not required
		return nil
	}

	if m.Customconfig != nil {
		if err := m.Customconfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parameters" + "." + "customconfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Parameters" + "." + "customconfig")
			}
			return err
		}
	}

	return nil
}

func (m *KitconfigParameters) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {

		if err := validate.Pattern("Parameters"+"."+"extensions"+"."+strconv.Itoa(i), "body", m.Extensions[i], `^[a-zA-Z.\/][a-zA-Z0-9-_.\/]*$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *KitconfigParameters) validateGlobalSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalSettings) { // not required
		return nil
	}

	if m.GlobalSettings != nil {
		if err := m.GlobalSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parameters" + "." + "global_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Parameters" + "." + "global_settings")
			}
			return err
		}
	}

	return nil
}

func (m *KitconfigParameters) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	for i := 0; i < len(m.Nodes); i++ {
		if swag.IsZero(m.Nodes[i]) { // not required
			continue
		}

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Parameters" + "." + "nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Parameters" + "." + "nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this kitconfig parameters based on the context it is used
func (m *KitconfigParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomconfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlobalSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigParameters) contextValidateCustomconfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Customconfig != nil {
		if err := m.Customconfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parameters" + "." + "customconfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Parameters" + "." + "customconfig")
			}
			return err
		}
	}

	return nil
}

func (m *KitconfigParameters) contextValidateGlobalSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalSettings != nil {
		if err := m.GlobalSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parameters" + "." + "global_settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Parameters" + "." + "global_settings")
			}
			return err
		}
	}

	return nil
}

func (m *KitconfigParameters) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nodes); i++ {

		if m.Nodes[i] != nil {
			if err := m.Nodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Parameters" + "." + "nodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Parameters" + "." + "nodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KitconfigParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KitconfigParameters) UnmarshalBinary(b []byte) error {
	var res KitconfigParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KitconfigParametersGlobalSettings kitconfig parameters global settings
//
// swagger:model KitconfigParametersGlobalSettings
type KitconfigParametersGlobalSettings struct {

	// dns server
	DNSServer []string `json:"dns_server"`

	// http proxy
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	HTTPProxy string `json:"http_proxy,omitempty"`

	// https proxy
	// Pattern: (?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])
	HTTPSProxy string `json:"https_proxy,omitempty"`

	// no proxy
	NoProxy string `json:"no_proxy,omitempty"`

	// ntp server
	// Pattern: ^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$
	NtpServer string `json:"ntp_server,omitempty"`

	// provider ip
	// Pattern: (\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b
	ProviderIP string `json:"provider_ip,omitempty"`

	// registry port
	// Pattern: ^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$
	RegistryPort string `json:"registry_port,omitempty"`

	// workflow port
	// Pattern: ^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$
	WorkflowPort string `json:"workflow_port,omitempty"`
}

// Validate validates this kitconfig parameters global settings
func (m *KitconfigParametersGlobalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPSProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtpServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KitconfigParametersGlobalSettings) validateHTTPProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPProxy) { // not required
		return nil
	}

	if err := validate.Pattern("Parameters"+"."+"global_settings"+"."+"http_proxy", "body", m.HTTPProxy, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigParametersGlobalSettings) validateHTTPSProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPSProxy) { // not required
		return nil
	}

	if err := validate.Pattern("Parameters"+"."+"global_settings"+"."+"https_proxy", "body", m.HTTPSProxy, `(?:(?:https?|http|ftp|file|oci)://|www.|ftp.)(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:([-A-Z0-9+&@#/%=~_|$?!:,.]*)|[A-Z0-9+&@#/%=~_|$])`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigParametersGlobalSettings) validateNtpServer(formats strfmt.Registry) error {
	if swag.IsZero(m.NtpServer) { // not required
		return nil
	}

	if err := validate.Pattern("Parameters"+"."+"global_settings"+"."+"ntp_server", "body", m.NtpServer, `^[a-zA-Z_$][a-zA-Z_.\-$0-9]*$`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigParametersGlobalSettings) validateProviderIP(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderIP) { // not required
		return nil
	}

	if err := validate.Pattern("Parameters"+"."+"global_settings"+"."+"provider_ip", "body", m.ProviderIP, `(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigParametersGlobalSettings) validateRegistryPort(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistryPort) { // not required
		return nil
	}

	if err := validate.Pattern("Parameters"+"."+"global_settings"+"."+"registry_port", "body", m.RegistryPort, `^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$`); err != nil {
		return err
	}

	return nil
}

func (m *KitconfigParametersGlobalSettings) validateWorkflowPort(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowPort) { // not required
		return nil
	}

	if err := validate.Pattern("Parameters"+"."+"global_settings"+"."+"workflow_port", "body", m.WorkflowPort, `^((6553[0-5])|(655[0-2][0-9])|(65[0-4][0-9]{2})|(6[0-4][0-9]{3})|([1-5][0-9]{4})|([0-5]{0,5})|([0-9]{1,4}))$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kitconfig parameters global settings based on context it is used
func (m *KitconfigParametersGlobalSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KitconfigParametersGlobalSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KitconfigParametersGlobalSettings) UnmarshalBinary(b []byte) error {
	var res KitconfigParametersGlobalSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ServerInstanceProfileNetworkInterfacesConfig - struct for ServerInstanceProfileNetworkInterfacesConfig
type ServerInstanceProfileNetworkInterfacesConfig struct {
	ServerInstanceInterfaceConfiguration *ServerInstanceInterfaceConfiguration
}

// ServerInstanceInterfaceConfigurationAsServerInstanceProfileNetworkInterfacesConfig is a convenience function that returns ServerInstanceInterfaceConfiguration wrapped in ServerInstanceProfileNetworkInterfacesConfig
func ServerInstanceInterfaceConfigurationAsServerInstanceProfileNetworkInterfacesConfig(v *ServerInstanceInterfaceConfiguration) ServerInstanceProfileNetworkInterfacesConfig {
	return ServerInstanceProfileNetworkInterfacesConfig{
		ServerInstanceInterfaceConfiguration: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerInstanceProfileNetworkInterfacesConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServerInstanceInterfaceConfiguration
	err = newStrictDecoder(data).Decode(&dst.ServerInstanceInterfaceConfiguration)
	if err == nil {
		jsonServerInstanceInterfaceConfiguration, _ := json.Marshal(dst.ServerInstanceInterfaceConfiguration)
		if string(jsonServerInstanceInterfaceConfiguration) == "{}" { // empty struct
			dst.ServerInstanceInterfaceConfiguration = nil
		} else {
			if err = validator.Validate(dst.ServerInstanceInterfaceConfiguration); err != nil {
				dst.ServerInstanceInterfaceConfiguration = nil
			} else {
				match++
			}
		}
	} else {
		dst.ServerInstanceInterfaceConfiguration = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ServerInstanceInterfaceConfiguration = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServerInstanceProfileNetworkInterfacesConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServerInstanceProfileNetworkInterfacesConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerInstanceProfileNetworkInterfacesConfig) MarshalJSON() ([]byte, error) {
	if src.ServerInstanceInterfaceConfiguration != nil {
		return json.Marshal(&src.ServerInstanceInterfaceConfiguration)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerInstanceProfileNetworkInterfacesConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ServerInstanceInterfaceConfiguration != nil {
		return obj.ServerInstanceInterfaceConfiguration
	}

	// all schemas are nil
	return nil
}

type NullableServerInstanceProfileNetworkInterfacesConfig struct {
	value *ServerInstanceProfileNetworkInterfacesConfig
	isSet bool
}

func (v NullableServerInstanceProfileNetworkInterfacesConfig) Get() *ServerInstanceProfileNetworkInterfacesConfig {
	return v.value
}

func (v *NullableServerInstanceProfileNetworkInterfacesConfig) Set(val *ServerInstanceProfileNetworkInterfacesConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceProfileNetworkInterfacesConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceProfileNetworkInterfacesConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceProfileNetworkInterfacesConfig(val *ServerInstanceProfileNetworkInterfacesConfig) *NullableServerInstanceProfileNetworkInterfacesConfig {
	return &NullableServerInstanceProfileNetworkInterfacesConfig{value: val, isSet: true}
}

func (v NullableServerInstanceProfileNetworkInterfacesConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceProfileNetworkInterfacesConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



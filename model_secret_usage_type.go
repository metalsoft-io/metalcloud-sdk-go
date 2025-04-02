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
	"fmt"
)

// SecretUsageType the model 'SecretUsageType'
type SecretUsageType string

// List of SecretUsageType
const (
	HTTP_REQUEST_SECRET SecretUsageType = "HTTPRequestSecret"
	JAVA_SCRIPT_SECRET SecretUsageType = "JavaScriptSecret"
	API_CALL_SECRET SecretUsageType = "APICallSecret"
	ANSIBLE_BUNDLE_SECRET SecretUsageType = "AnsibleBundleSecret"
	SSH_EXEC_SECRET SecretUsageType = "SSHExecSecret"
	COPY_SECRET SecretUsageType = "CopySecret"
	OS_ASSET_SECRET SecretUsageType = "OSAssetSecret"
)

// All allowed values of SecretUsageType enum
var AllowedSecretUsageTypeEnumValues = []SecretUsageType{
	"HTTPRequestSecret",
	"JavaScriptSecret",
	"APICallSecret",
	"AnsibleBundleSecret",
	"SSHExecSecret",
	"CopySecret",
	"OSAssetSecret",
}

func (v *SecretUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecretUsageType(value)
	for _, existing := range AllowedSecretUsageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecretUsageType", value)
}

// NewSecretUsageTypeFromValue returns a pointer to a valid SecretUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecretUsageTypeFromValue(v string) (*SecretUsageType, error) {
	ev := SecretUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecretUsageType: valid values are %v", v, AllowedSecretUsageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecretUsageType) IsValid() bool {
	for _, existing := range AllowedSecretUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecretUsageType value
func (v SecretUsageType) Ptr() *SecretUsageType {
	return &v
}

type NullableSecretUsageType struct {
	value *SecretUsageType
	isSet bool
}

func (v NullableSecretUsageType) Get() *SecretUsageType {
	return v.value
}

func (v *NullableSecretUsageType) Set(val *SecretUsageType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretUsageType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretUsageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretUsageType(val *SecretUsageType) *NullableSecretUsageType {
	return &NullableSecretUsageType{value: val, isSet: true}
}

func (v NullableSecretUsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretUsageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


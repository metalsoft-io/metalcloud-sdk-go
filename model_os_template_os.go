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

// checks if the OSTemplateOs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSTemplateOs{}

// OSTemplateOs struct for OSTemplateOs
type OSTemplateOs struct {
	// The name of the operating system that this template will install
	Name string `json:"name"`
	// The version of the operating system that this template will install
	Version string `json:"version"`
	Credential OSTemplateOsCredential `json:"credential"`
	// The initial operating system SSH port.
	SshPort *int32 `json:"sshPort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSTemplateOs OSTemplateOs

// NewOSTemplateOs instantiates a new OSTemplateOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSTemplateOs(name string, version string, credential OSTemplateOsCredential) *OSTemplateOs {
	this := OSTemplateOs{}
	this.Name = name
	this.Version = version
	this.Credential = credential
	return &this
}

// NewOSTemplateOsWithDefaults instantiates a new OSTemplateOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSTemplateOsWithDefaults() *OSTemplateOs {
	this := OSTemplateOs{}
	return &this
}

// GetName returns the Name field value
func (o *OSTemplateOs) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OSTemplateOs) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OSTemplateOs) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *OSTemplateOs) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OSTemplateOs) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OSTemplateOs) SetVersion(v string) {
	o.Version = v
}

// GetCredential returns the Credential field value
func (o *OSTemplateOs) GetCredential() OSTemplateOsCredential {
	if o == nil {
		var ret OSTemplateOsCredential
		return ret
	}

	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value
// and a boolean to check if the value has been set.
func (o *OSTemplateOs) GetCredentialOk() (*OSTemplateOsCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credential, true
}

// SetCredential sets field value
func (o *OSTemplateOs) SetCredential(v OSTemplateOsCredential) {
	o.Credential = v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *OSTemplateOs) GetSshPort() int32 {
	if o == nil || IsNil(o.SshPort) {
		var ret int32
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSTemplateOs) GetSshPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *OSTemplateOs) HasSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int32 and assigns it to the SshPort field.
func (o *OSTemplateOs) SetSshPort(v int32) {
	o.SshPort = &v
}

func (o OSTemplateOs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSTemplateOs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	toSerialize["credential"] = o.Credential
	if !IsNil(o.SshPort) {
		toSerialize["sshPort"] = o.SshPort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSTemplateOs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
		"credential",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOSTemplateOs := _OSTemplateOs{}

	err = json.Unmarshal(data, &varOSTemplateOs)

	if err != nil {
		return err
	}

	*o = OSTemplateOs(varOSTemplateOs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		delete(additionalProperties, "credential")
		delete(additionalProperties, "sshPort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSTemplateOs struct {
	value *OSTemplateOs
	isSet bool
}

func (v NullableOSTemplateOs) Get() *OSTemplateOs {
	return v.value
}

func (v *NullableOSTemplateOs) Set(val *OSTemplateOs) {
	v.value = val
	v.isSet = true
}

func (v NullableOSTemplateOs) IsSet() bool {
	return v.isSet
}

func (v *NullableOSTemplateOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSTemplateOs(val *OSTemplateOs) *NullableOSTemplateOs {
	return &NullableOSTemplateOs{value: val, isSet: true}
}

func (v NullableOSTemplateOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSTemplateOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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
)

// checks if the UpdateInfrastructureMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateInfrastructureMeta{}

// UpdateInfrastructureMeta struct for UpdateInfrastructureMeta
type UpdateInfrastructureMeta struct {
	GuiSettings *GenericGUISettings `json:"guiSettings,omitempty"`
	// Tags for the Infrastructure.
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateInfrastructureMeta UpdateInfrastructureMeta

// NewUpdateInfrastructureMeta instantiates a new UpdateInfrastructureMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInfrastructureMeta() *UpdateInfrastructureMeta {
	this := UpdateInfrastructureMeta{}
	return &this
}

// NewUpdateInfrastructureMetaWithDefaults instantiates a new UpdateInfrastructureMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInfrastructureMetaWithDefaults() *UpdateInfrastructureMeta {
	this := UpdateInfrastructureMeta{}
	return &this
}

// GetGuiSettings returns the GuiSettings field value if set, zero value otherwise.
func (o *UpdateInfrastructureMeta) GetGuiSettings() GenericGUISettings {
	if o == nil || IsNil(o.GuiSettings) {
		var ret GenericGUISettings
		return ret
	}
	return *o.GuiSettings
}

// GetGuiSettingsOk returns a tuple with the GuiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfrastructureMeta) GetGuiSettingsOk() (*GenericGUISettings, bool) {
	if o == nil || IsNil(o.GuiSettings) {
		return nil, false
	}
	return o.GuiSettings, true
}

// HasGuiSettings returns a boolean if a field has been set.
func (o *UpdateInfrastructureMeta) HasGuiSettings() bool {
	if o != nil && !IsNil(o.GuiSettings) {
		return true
	}

	return false
}

// SetGuiSettings gets a reference to the given GenericGUISettings and assigns it to the GuiSettings field.
func (o *UpdateInfrastructureMeta) SetGuiSettings(v GenericGUISettings) {
	o.GuiSettings = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateInfrastructureMeta) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInfrastructureMeta) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateInfrastructureMeta) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateInfrastructureMeta) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateInfrastructureMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateInfrastructureMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuiSettings) {
		toSerialize["guiSettings"] = o.GuiSettings
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateInfrastructureMeta) UnmarshalJSON(data []byte) (err error) {
	varUpdateInfrastructureMeta := _UpdateInfrastructureMeta{}

	err = json.Unmarshal(data, &varUpdateInfrastructureMeta)

	if err != nil {
		return err
	}

	*o = UpdateInfrastructureMeta(varUpdateInfrastructureMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "guiSettings")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateInfrastructureMeta struct {
	value *UpdateInfrastructureMeta
	isSet bool
}

func (v NullableUpdateInfrastructureMeta) Get() *UpdateInfrastructureMeta {
	return v.value
}

func (v *NullableUpdateInfrastructureMeta) Set(val *UpdateInfrastructureMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInfrastructureMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInfrastructureMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInfrastructureMeta(val *UpdateInfrastructureMeta) *NullableUpdateInfrastructureMeta {
	return &NullableUpdateInfrastructureMeta{value: val, isSet: true}
}

func (v NullableUpdateInfrastructureMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInfrastructureMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



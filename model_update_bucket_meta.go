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

// checks if the UpdateBucketMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBucketMeta{}

// UpdateBucketMeta struct for UpdateBucketMeta
type UpdateBucketMeta struct {
	GuiSettings *GenericGUISettings `json:"guiSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateBucketMeta UpdateBucketMeta

// NewUpdateBucketMeta instantiates a new UpdateBucketMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBucketMeta() *UpdateBucketMeta {
	this := UpdateBucketMeta{}
	return &this
}

// NewUpdateBucketMetaWithDefaults instantiates a new UpdateBucketMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBucketMetaWithDefaults() *UpdateBucketMeta {
	this := UpdateBucketMeta{}
	return &this
}

// GetGuiSettings returns the GuiSettings field value if set, zero value otherwise.
func (o *UpdateBucketMeta) GetGuiSettings() GenericGUISettings {
	if o == nil || IsNil(o.GuiSettings) {
		var ret GenericGUISettings
		return ret
	}
	return *o.GuiSettings
}

// GetGuiSettingsOk returns a tuple with the GuiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBucketMeta) GetGuiSettingsOk() (*GenericGUISettings, bool) {
	if o == nil || IsNil(o.GuiSettings) {
		return nil, false
	}
	return o.GuiSettings, true
}

// HasGuiSettings returns a boolean if a field has been set.
func (o *UpdateBucketMeta) HasGuiSettings() bool {
	if o != nil && !IsNil(o.GuiSettings) {
		return true
	}

	return false
}

// SetGuiSettings gets a reference to the given GenericGUISettings and assigns it to the GuiSettings field.
func (o *UpdateBucketMeta) SetGuiSettings(v GenericGUISettings) {
	o.GuiSettings = &v
}

func (o UpdateBucketMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBucketMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuiSettings) {
		toSerialize["guiSettings"] = o.GuiSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateBucketMeta) UnmarshalJSON(data []byte) (err error) {
	varUpdateBucketMeta := _UpdateBucketMeta{}

	err = json.Unmarshal(data, &varUpdateBucketMeta)

	if err != nil {
		return err
	}

	*o = UpdateBucketMeta(varUpdateBucketMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "guiSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateBucketMeta struct {
	value *UpdateBucketMeta
	isSet bool
}

func (v NullableUpdateBucketMeta) Get() *UpdateBucketMeta {
	return v.value
}

func (v *NullableUpdateBucketMeta) Set(val *UpdateBucketMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBucketMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBucketMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBucketMeta(val *UpdateBucketMeta) *NullableUpdateBucketMeta {
	return &NullableUpdateBucketMeta{value: val, isSet: true}
}

func (v NullableUpdateBucketMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBucketMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



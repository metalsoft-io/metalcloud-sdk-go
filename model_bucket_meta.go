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

// checks if the BucketMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketMeta{}

// BucketMeta struct for BucketMeta
type BucketMeta struct {
	GuiSettings *GenericGUISettings `json:"guiSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BucketMeta BucketMeta

// NewBucketMeta instantiates a new BucketMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketMeta() *BucketMeta {
	this := BucketMeta{}
	return &this
}

// NewBucketMetaWithDefaults instantiates a new BucketMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketMetaWithDefaults() *BucketMeta {
	this := BucketMeta{}
	return &this
}

// GetGuiSettings returns the GuiSettings field value if set, zero value otherwise.
func (o *BucketMeta) GetGuiSettings() GenericGUISettings {
	if o == nil || IsNil(o.GuiSettings) {
		var ret GenericGUISettings
		return ret
	}
	return *o.GuiSettings
}

// GetGuiSettingsOk returns a tuple with the GuiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketMeta) GetGuiSettingsOk() (*GenericGUISettings, bool) {
	if o == nil || IsNil(o.GuiSettings) {
		return nil, false
	}
	return o.GuiSettings, true
}

// HasGuiSettings returns a boolean if a field has been set.
func (o *BucketMeta) HasGuiSettings() bool {
	if o != nil && !IsNil(o.GuiSettings) {
		return true
	}

	return false
}

// SetGuiSettings gets a reference to the given GenericGUISettings and assigns it to the GuiSettings field.
func (o *BucketMeta) SetGuiSettings(v GenericGUISettings) {
	o.GuiSettings = &v
}

func (o BucketMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuiSettings) {
		toSerialize["guiSettings"] = o.GuiSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BucketMeta) UnmarshalJSON(data []byte) (err error) {
	varBucketMeta := _BucketMeta{}

	err = json.Unmarshal(data, &varBucketMeta)

	if err != nil {
		return err
	}

	*o = BucketMeta(varBucketMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "guiSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBucketMeta struct {
	value *BucketMeta
	isSet bool
}

func (v NullableBucketMeta) Get() *BucketMeta {
	return v.value
}

func (v *NullableBucketMeta) Set(val *BucketMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketMeta(val *BucketMeta) *NullableBucketMeta {
	return &NullableBucketMeta{value: val, isSet: true}
}

func (v NullableBucketMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



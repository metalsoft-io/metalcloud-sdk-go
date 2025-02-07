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

// checks if the OSTemplateImageBuild type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSTemplateImageBuild{}

// OSTemplateImageBuild struct for OSTemplateImageBuild
type OSTemplateImageBuild struct {
	// Indicates whether the OS template requires an image build
	Required bool `json:"required"`
	// The OS template image build provider
	Provider *string `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSTemplateImageBuild OSTemplateImageBuild

// NewOSTemplateImageBuild instantiates a new OSTemplateImageBuild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSTemplateImageBuild(required bool) *OSTemplateImageBuild {
	this := OSTemplateImageBuild{}
	this.Required = required
	return &this
}

// NewOSTemplateImageBuildWithDefaults instantiates a new OSTemplateImageBuild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSTemplateImageBuildWithDefaults() *OSTemplateImageBuild {
	this := OSTemplateImageBuild{}
	var required bool = true
	this.Required = required
	return &this
}

// GetRequired returns the Required field value
func (o *OSTemplateImageBuild) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *OSTemplateImageBuild) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *OSTemplateImageBuild) SetRequired(v bool) {
	o.Required = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *OSTemplateImageBuild) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSTemplateImageBuild) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *OSTemplateImageBuild) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *OSTemplateImageBuild) SetProvider(v string) {
	o.Provider = &v
}

func (o OSTemplateImageBuild) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSTemplateImageBuild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["required"] = o.Required
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSTemplateImageBuild) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"required",
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

	varOSTemplateImageBuild := _OSTemplateImageBuild{}

	err = json.Unmarshal(data, &varOSTemplateImageBuild)

	if err != nil {
		return err
	}

	*o = OSTemplateImageBuild(varOSTemplateImageBuild)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "required")
		delete(additionalProperties, "provider")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSTemplateImageBuild struct {
	value *OSTemplateImageBuild
	isSet bool
}

func (v NullableOSTemplateImageBuild) Get() *OSTemplateImageBuild {
	return v.value
}

func (v *NullableOSTemplateImageBuild) Set(val *OSTemplateImageBuild) {
	v.value = val
	v.isSet = true
}

func (v NullableOSTemplateImageBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableOSTemplateImageBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSTemplateImageBuild(val *OSTemplateImageBuild) *NullableOSTemplateImageBuild {
	return &NullableOSTemplateImageBuild{value: val, isSet: true}
}

func (v NullableOSTemplateImageBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSTemplateImageBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



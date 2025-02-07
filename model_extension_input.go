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

// checks if the ExtensionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionInput{}

// ExtensionInput struct for ExtensionInput
type ExtensionInput struct {
	// Label of the input.
	Label string `json:"label"`
	// Name of the input.
	Name string `json:"name"`
	// Type of the input.
	InputType string `json:"inputType"`
	Options ExtensionInputOptions `json:"options"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionInput ExtensionInput

// NewExtensionInput instantiates a new ExtensionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionInput(label string, name string, inputType string, options ExtensionInputOptions) *ExtensionInput {
	this := ExtensionInput{}
	this.Label = label
	this.Name = name
	this.InputType = inputType
	this.Options = options
	return &this
}

// NewExtensionInputWithDefaults instantiates a new ExtensionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionInputWithDefaults() *ExtensionInput {
	this := ExtensionInput{}
	return &this
}

// GetLabel returns the Label field value
func (o *ExtensionInput) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ExtensionInput) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ExtensionInput) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *ExtensionInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtensionInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtensionInput) SetName(v string) {
	o.Name = v
}

// GetInputType returns the InputType field value
func (o *ExtensionInput) GetInputType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputType
}

// GetInputTypeOk returns a tuple with the InputType field value
// and a boolean to check if the value has been set.
func (o *ExtensionInput) GetInputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputType, true
}

// SetInputType sets field value
func (o *ExtensionInput) SetInputType(v string) {
	o.InputType = v
}

// GetOptions returns the Options field value
func (o *ExtensionInput) GetOptions() ExtensionInputOptions {
	if o == nil {
		var ret ExtensionInputOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ExtensionInput) GetOptionsOk() (*ExtensionInputOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ExtensionInput) SetOptions(v ExtensionInputOptions) {
	o.Options = v
}

func (o ExtensionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["inputType"] = o.InputType
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
		"inputType",
		"options",
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

	varExtensionInput := _ExtensionInput{}

	err = json.Unmarshal(data, &varExtensionInput)

	if err != nil {
		return err
	}

	*o = ExtensionInput(varExtensionInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "inputType")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionInput struct {
	value *ExtensionInput
	isSet bool
}

func (v NullableExtensionInput) Get() *ExtensionInput {
	return v.value
}

func (v *NullableExtensionInput) Set(val *ExtensionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInput(val *ExtensionInput) *NullableExtensionInput {
	return &NullableExtensionInput{value: val, isSet: true}
}

func (v NullableExtensionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



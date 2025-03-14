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

// checks if the ExtensionOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionOutput{}

// ExtensionOutput struct for ExtensionOutput
type ExtensionOutput struct {
	// Label of the output.
	Label string `json:"label"`
	// Name of the output.
	Name string `json:"name"`
	// Type of the output.
	OutputType string `json:"outputType"`
	// Template of the output.
	Template *string `json:"template,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionOutput ExtensionOutput

// NewExtensionOutput instantiates a new ExtensionOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionOutput(label string, name string, outputType string) *ExtensionOutput {
	this := ExtensionOutput{}
	this.Label = label
	this.Name = name
	this.OutputType = outputType
	return &this
}

// NewExtensionOutputWithDefaults instantiates a new ExtensionOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionOutputWithDefaults() *ExtensionOutput {
	this := ExtensionOutput{}
	return &this
}

// GetLabel returns the Label field value
func (o *ExtensionOutput) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ExtensionOutput) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ExtensionOutput) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *ExtensionOutput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtensionOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtensionOutput) SetName(v string) {
	o.Name = v
}

// GetOutputType returns the OutputType field value
func (o *ExtensionOutput) GetOutputType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputType
}

// GetOutputTypeOk returns a tuple with the OutputType field value
// and a boolean to check if the value has been set.
func (o *ExtensionOutput) GetOutputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputType, true
}

// SetOutputType sets field value
func (o *ExtensionOutput) SetOutputType(v string) {
	o.OutputType = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ExtensionOutput) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionOutput) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ExtensionOutput) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ExtensionOutput) SetTemplate(v string) {
	o.Template = &v
}

func (o ExtensionOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["outputType"] = o.OutputType
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
		"outputType",
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

	varExtensionOutput := _ExtensionOutput{}

	err = json.Unmarshal(data, &varExtensionOutput)

	if err != nil {
		return err
	}

	*o = ExtensionOutput(varExtensionOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "outputType")
		delete(additionalProperties, "template")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionOutput struct {
	value *ExtensionOutput
	isSet bool
}

func (v NullableExtensionOutput) Get() *ExtensionOutput {
	return v.value
}

func (v *NullableExtensionOutput) Set(val *ExtensionOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionOutput(val *ExtensionOutput) *NullableExtensionOutput {
	return &NullableExtensionOutput{value: val, isSet: true}
}

func (v NullableExtensionOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



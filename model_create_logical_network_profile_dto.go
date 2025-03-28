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

// checks if the CreateLogicalNetworkProfileDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogicalNetworkProfileDto{}

// CreateLogicalNetworkProfileDto struct for CreateLogicalNetworkProfileDto
type CreateLogicalNetworkProfileDto struct {
	// Label for the logical network profile
	Label *string `json:"label,omitempty"`
	// Name of the logical network profile
	Name *string `json:"name,omitempty"`
	// Description of the logical network profile
	Description *string `json:"description,omitempty"`
	// Annotations for the logical network profile
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// Type of the logical network profile
	LogicalNetworkType string `json:"logicalNetworkType"`
	AdditionalProperties map[string]interface{}
}

type _CreateLogicalNetworkProfileDto CreateLogicalNetworkProfileDto

// NewCreateLogicalNetworkProfileDto instantiates a new CreateLogicalNetworkProfileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogicalNetworkProfileDto(logicalNetworkType string) *CreateLogicalNetworkProfileDto {
	this := CreateLogicalNetworkProfileDto{}
	this.LogicalNetworkType = logicalNetworkType
	return &this
}

// NewCreateLogicalNetworkProfileDtoWithDefaults instantiates a new CreateLogicalNetworkProfileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogicalNetworkProfileDtoWithDefaults() *CreateLogicalNetworkProfileDto {
	this := CreateLogicalNetworkProfileDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateLogicalNetworkProfileDto) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetworkProfileDto) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateLogicalNetworkProfileDto) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateLogicalNetworkProfileDto) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateLogicalNetworkProfileDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetworkProfileDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateLogicalNetworkProfileDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateLogicalNetworkProfileDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateLogicalNetworkProfileDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetworkProfileDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateLogicalNetworkProfileDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateLogicalNetworkProfileDto) SetDescription(v string) {
	o.Description = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *CreateLogicalNetworkProfileDto) GetAnnotations() map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetworkProfileDto) GetAnnotationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *CreateLogicalNetworkProfileDto) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]interface{} and assigns it to the Annotations field.
func (o *CreateLogicalNetworkProfileDto) SetAnnotations(v map[string]interface{}) {
	o.Annotations = v
}

// GetLogicalNetworkType returns the LogicalNetworkType field value
func (o *CreateLogicalNetworkProfileDto) GetLogicalNetworkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogicalNetworkType
}

// GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field value
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetworkProfileDto) GetLogicalNetworkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalNetworkType, true
}

// SetLogicalNetworkType sets field value
func (o *CreateLogicalNetworkProfileDto) SetLogicalNetworkType(v string) {
	o.LogicalNetworkType = v
}

func (o CreateLogicalNetworkProfileDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogicalNetworkProfileDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	toSerialize["logicalNetworkType"] = o.LogicalNetworkType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateLogicalNetworkProfileDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logicalNetworkType",
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

	varCreateLogicalNetworkProfileDto := _CreateLogicalNetworkProfileDto{}

	err = json.Unmarshal(data, &varCreateLogicalNetworkProfileDto)

	if err != nil {
		return err
	}

	*o = CreateLogicalNetworkProfileDto(varCreateLogicalNetworkProfileDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "annotations")
		delete(additionalProperties, "logicalNetworkType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateLogicalNetworkProfileDto struct {
	value *CreateLogicalNetworkProfileDto
	isSet bool
}

func (v NullableCreateLogicalNetworkProfileDto) Get() *CreateLogicalNetworkProfileDto {
	return v.value
}

func (v *NullableCreateLogicalNetworkProfileDto) Set(val *CreateLogicalNetworkProfileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogicalNetworkProfileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogicalNetworkProfileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogicalNetworkProfileDto(val *CreateLogicalNetworkProfileDto) *NullableCreateLogicalNetworkProfileDto {
	return &NullableCreateLogicalNetworkProfileDto{value: val, isSet: true}
}

func (v NullableCreateLogicalNetworkProfileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogicalNetworkProfileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



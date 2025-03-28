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

// checks if the UpdateLogicalNetworkProfileDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLogicalNetworkProfileDto{}

// UpdateLogicalNetworkProfileDto struct for UpdateLogicalNetworkProfileDto
type UpdateLogicalNetworkProfileDto struct {
	// Label for the logical network profile
	Label *string `json:"label,omitempty"`
	// Name of the logical network profile
	Name *string `json:"name,omitempty"`
	// Description of the logical network profile
	Description *string `json:"description,omitempty"`
	// Annotations for the logical network profile
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// Type of the logical network profile
	LogicalNetworkType *string `json:"logicalNetworkType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLogicalNetworkProfileDto UpdateLogicalNetworkProfileDto

// NewUpdateLogicalNetworkProfileDto instantiates a new UpdateLogicalNetworkProfileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLogicalNetworkProfileDto() *UpdateLogicalNetworkProfileDto {
	this := UpdateLogicalNetworkProfileDto{}
	return &this
}

// NewUpdateLogicalNetworkProfileDtoWithDefaults instantiates a new UpdateLogicalNetworkProfileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLogicalNetworkProfileDtoWithDefaults() *UpdateLogicalNetworkProfileDto {
	this := UpdateLogicalNetworkProfileDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UpdateLogicalNetworkProfileDto) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkProfileDto) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UpdateLogicalNetworkProfileDto) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UpdateLogicalNetworkProfileDto) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateLogicalNetworkProfileDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkProfileDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateLogicalNetworkProfileDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateLogicalNetworkProfileDto) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateLogicalNetworkProfileDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkProfileDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateLogicalNetworkProfileDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateLogicalNetworkProfileDto) SetDescription(v string) {
	o.Description = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *UpdateLogicalNetworkProfileDto) GetAnnotations() map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkProfileDto) GetAnnotationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *UpdateLogicalNetworkProfileDto) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]interface{} and assigns it to the Annotations field.
func (o *UpdateLogicalNetworkProfileDto) SetAnnotations(v map[string]interface{}) {
	o.Annotations = v
}

// GetLogicalNetworkType returns the LogicalNetworkType field value if set, zero value otherwise.
func (o *UpdateLogicalNetworkProfileDto) GetLogicalNetworkType() string {
	if o == nil || IsNil(o.LogicalNetworkType) {
		var ret string
		return ret
	}
	return *o.LogicalNetworkType
}

// GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkProfileDto) GetLogicalNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LogicalNetworkType) {
		return nil, false
	}
	return o.LogicalNetworkType, true
}

// HasLogicalNetworkType returns a boolean if a field has been set.
func (o *UpdateLogicalNetworkProfileDto) HasLogicalNetworkType() bool {
	if o != nil && !IsNil(o.LogicalNetworkType) {
		return true
	}

	return false
}

// SetLogicalNetworkType gets a reference to the given string and assigns it to the LogicalNetworkType field.
func (o *UpdateLogicalNetworkProfileDto) SetLogicalNetworkType(v string) {
	o.LogicalNetworkType = &v
}

func (o UpdateLogicalNetworkProfileDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLogicalNetworkProfileDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LogicalNetworkType) {
		toSerialize["logicalNetworkType"] = o.LogicalNetworkType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateLogicalNetworkProfileDto) UnmarshalJSON(data []byte) (err error) {
	varUpdateLogicalNetworkProfileDto := _UpdateLogicalNetworkProfileDto{}

	err = json.Unmarshal(data, &varUpdateLogicalNetworkProfileDto)

	if err != nil {
		return err
	}

	*o = UpdateLogicalNetworkProfileDto(varUpdateLogicalNetworkProfileDto)

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

type NullableUpdateLogicalNetworkProfileDto struct {
	value *UpdateLogicalNetworkProfileDto
	isSet bool
}

func (v NullableUpdateLogicalNetworkProfileDto) Get() *UpdateLogicalNetworkProfileDto {
	return v.value
}

func (v *NullableUpdateLogicalNetworkProfileDto) Set(val *UpdateLogicalNetworkProfileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLogicalNetworkProfileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLogicalNetworkProfileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLogicalNetworkProfileDto(val *UpdateLogicalNetworkProfileDto) *NullableUpdateLogicalNetworkProfileDto {
	return &NullableUpdateLogicalNetworkProfileDto{value: val, isSet: true}
}

func (v NullableUpdateLogicalNetworkProfileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLogicalNetworkProfileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



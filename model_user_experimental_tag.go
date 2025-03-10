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

// checks if the UserExperimentalTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserExperimentalTag{}

// UserExperimentalTag struct for UserExperimentalTag
type UserExperimentalTag struct {
	// The experimental tag for the user
	ExperimentalTag *string `json:"experimentalTag,omitempty"`
	// Whether to remove the experimental tag
	RemoveExperimentalTag *bool `json:"removeExperimentalTag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserExperimentalTag UserExperimentalTag

// NewUserExperimentalTag instantiates a new UserExperimentalTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExperimentalTag() *UserExperimentalTag {
	this := UserExperimentalTag{}
	var removeExperimentalTag bool = false
	this.RemoveExperimentalTag = &removeExperimentalTag
	return &this
}

// NewUserExperimentalTagWithDefaults instantiates a new UserExperimentalTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExperimentalTagWithDefaults() *UserExperimentalTag {
	this := UserExperimentalTag{}
	var removeExperimentalTag bool = false
	this.RemoveExperimentalTag = &removeExperimentalTag
	return &this
}

// GetExperimentalTag returns the ExperimentalTag field value if set, zero value otherwise.
func (o *UserExperimentalTag) GetExperimentalTag() string {
	if o == nil || IsNil(o.ExperimentalTag) {
		var ret string
		return ret
	}
	return *o.ExperimentalTag
}

// GetExperimentalTagOk returns a tuple with the ExperimentalTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExperimentalTag) GetExperimentalTagOk() (*string, bool) {
	if o == nil || IsNil(o.ExperimentalTag) {
		return nil, false
	}
	return o.ExperimentalTag, true
}

// HasExperimentalTag returns a boolean if a field has been set.
func (o *UserExperimentalTag) HasExperimentalTag() bool {
	if o != nil && !IsNil(o.ExperimentalTag) {
		return true
	}

	return false
}

// SetExperimentalTag gets a reference to the given string and assigns it to the ExperimentalTag field.
func (o *UserExperimentalTag) SetExperimentalTag(v string) {
	o.ExperimentalTag = &v
}

// GetRemoveExperimentalTag returns the RemoveExperimentalTag field value if set, zero value otherwise.
func (o *UserExperimentalTag) GetRemoveExperimentalTag() bool {
	if o == nil || IsNil(o.RemoveExperimentalTag) {
		var ret bool
		return ret
	}
	return *o.RemoveExperimentalTag
}

// GetRemoveExperimentalTagOk returns a tuple with the RemoveExperimentalTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExperimentalTag) GetRemoveExperimentalTagOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveExperimentalTag) {
		return nil, false
	}
	return o.RemoveExperimentalTag, true
}

// HasRemoveExperimentalTag returns a boolean if a field has been set.
func (o *UserExperimentalTag) HasRemoveExperimentalTag() bool {
	if o != nil && !IsNil(o.RemoveExperimentalTag) {
		return true
	}

	return false
}

// SetRemoveExperimentalTag gets a reference to the given bool and assigns it to the RemoveExperimentalTag field.
func (o *UserExperimentalTag) SetRemoveExperimentalTag(v bool) {
	o.RemoveExperimentalTag = &v
}

func (o UserExperimentalTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExperimentalTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExperimentalTag) {
		toSerialize["experimentalTag"] = o.ExperimentalTag
	}
	if !IsNil(o.RemoveExperimentalTag) {
		toSerialize["removeExperimentalTag"] = o.RemoveExperimentalTag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserExperimentalTag) UnmarshalJSON(data []byte) (err error) {
	varUserExperimentalTag := _UserExperimentalTag{}

	err = json.Unmarshal(data, &varUserExperimentalTag)

	if err != nil {
		return err
	}

	*o = UserExperimentalTag(varUserExperimentalTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "experimentalTag")
		delete(additionalProperties, "removeExperimentalTag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserExperimentalTag struct {
	value *UserExperimentalTag
	isSet bool
}

func (v NullableUserExperimentalTag) Get() *UserExperimentalTag {
	return v.value
}

func (v *NullableUserExperimentalTag) Set(val *UserExperimentalTag) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExperimentalTag) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExperimentalTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserExperimentalTag(val *UserExperimentalTag) *NullableUserExperimentalTag {
	return &NullableUserExperimentalTag{value: val, isSet: true}
}

func (v NullableUserExperimentalTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExperimentalTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



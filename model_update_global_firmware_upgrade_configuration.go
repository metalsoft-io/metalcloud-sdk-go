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

// checks if the UpdateGlobalFirmwareUpgradeConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGlobalFirmwareUpgradeConfiguration{}

// UpdateGlobalFirmwareUpgradeConfiguration struct for UpdateGlobalFirmwareUpgradeConfiguration
type UpdateGlobalFirmwareUpgradeConfiguration struct {
	// Indicates whether the firmware upgrade policies are activated globally.
	Activated *bool `json:"activated,omitempty"`
	// The timestamp when the firmware upgrade policies are allowed to be executed. Only the hours and minutes are considered.
	UpgradeStartTime *string `json:"upgradeStartTime,omitempty"`
	// The timestamp when the firmware upgrade policies are no longer allowed to be executed. Only the hours and minutes are considered.
	UpgradeEndTime *string `json:"upgradeEndTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateGlobalFirmwareUpgradeConfiguration UpdateGlobalFirmwareUpgradeConfiguration

// NewUpdateGlobalFirmwareUpgradeConfiguration instantiates a new UpdateGlobalFirmwareUpgradeConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGlobalFirmwareUpgradeConfiguration() *UpdateGlobalFirmwareUpgradeConfiguration {
	this := UpdateGlobalFirmwareUpgradeConfiguration{}
	return &this
}

// NewUpdateGlobalFirmwareUpgradeConfigurationWithDefaults instantiates a new UpdateGlobalFirmwareUpgradeConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGlobalFirmwareUpgradeConfigurationWithDefaults() *UpdateGlobalFirmwareUpgradeConfiguration {
	this := UpdateGlobalFirmwareUpgradeConfiguration{}
	return &this
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) GetActivated() bool {
	if o == nil || IsNil(o.Activated) {
		var ret bool
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) GetActivatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Activated) {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) HasActivated() bool {
	if o != nil && !IsNil(o.Activated) {
		return true
	}

	return false
}

// SetActivated gets a reference to the given bool and assigns it to the Activated field.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) SetActivated(v bool) {
	o.Activated = &v
}

// GetUpgradeStartTime returns the UpgradeStartTime field value if set, zero value otherwise.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) GetUpgradeStartTime() string {
	if o == nil || IsNil(o.UpgradeStartTime) {
		var ret string
		return ret
	}
	return *o.UpgradeStartTime
}

// GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) GetUpgradeStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeStartTime) {
		return nil, false
	}
	return o.UpgradeStartTime, true
}

// HasUpgradeStartTime returns a boolean if a field has been set.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) HasUpgradeStartTime() bool {
	if o != nil && !IsNil(o.UpgradeStartTime) {
		return true
	}

	return false
}

// SetUpgradeStartTime gets a reference to the given string and assigns it to the UpgradeStartTime field.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) SetUpgradeStartTime(v string) {
	o.UpgradeStartTime = &v
}

// GetUpgradeEndTime returns the UpgradeEndTime field value if set, zero value otherwise.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) GetUpgradeEndTime() string {
	if o == nil || IsNil(o.UpgradeEndTime) {
		var ret string
		return ret
	}
	return *o.UpgradeEndTime
}

// GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) GetUpgradeEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeEndTime) {
		return nil, false
	}
	return o.UpgradeEndTime, true
}

// HasUpgradeEndTime returns a boolean if a field has been set.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) HasUpgradeEndTime() bool {
	if o != nil && !IsNil(o.UpgradeEndTime) {
		return true
	}

	return false
}

// SetUpgradeEndTime gets a reference to the given string and assigns it to the UpgradeEndTime field.
func (o *UpdateGlobalFirmwareUpgradeConfiguration) SetUpgradeEndTime(v string) {
	o.UpgradeEndTime = &v
}

func (o UpdateGlobalFirmwareUpgradeConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGlobalFirmwareUpgradeConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activated) {
		toSerialize["activated"] = o.Activated
	}
	if !IsNil(o.UpgradeStartTime) {
		toSerialize["upgradeStartTime"] = o.UpgradeStartTime
	}
	if !IsNil(o.UpgradeEndTime) {
		toSerialize["upgradeEndTime"] = o.UpgradeEndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateGlobalFirmwareUpgradeConfiguration) UnmarshalJSON(data []byte) (err error) {
	varUpdateGlobalFirmwareUpgradeConfiguration := _UpdateGlobalFirmwareUpgradeConfiguration{}

	err = json.Unmarshal(data, &varUpdateGlobalFirmwareUpgradeConfiguration)

	if err != nil {
		return err
	}

	*o = UpdateGlobalFirmwareUpgradeConfiguration(varUpdateGlobalFirmwareUpgradeConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activated")
		delete(additionalProperties, "upgradeStartTime")
		delete(additionalProperties, "upgradeEndTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateGlobalFirmwareUpgradeConfiguration struct {
	value *UpdateGlobalFirmwareUpgradeConfiguration
	isSet bool
}

func (v NullableUpdateGlobalFirmwareUpgradeConfiguration) Get() *UpdateGlobalFirmwareUpgradeConfiguration {
	return v.value
}

func (v *NullableUpdateGlobalFirmwareUpgradeConfiguration) Set(val *UpdateGlobalFirmwareUpgradeConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGlobalFirmwareUpgradeConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGlobalFirmwareUpgradeConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGlobalFirmwareUpgradeConfiguration(val *UpdateGlobalFirmwareUpgradeConfiguration) *NullableUpdateGlobalFirmwareUpgradeConfiguration {
	return &NullableUpdateGlobalFirmwareUpgradeConfiguration{value: val, isSet: true}
}

func (v NullableUpdateGlobalFirmwareUpgradeConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGlobalFirmwareUpgradeConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



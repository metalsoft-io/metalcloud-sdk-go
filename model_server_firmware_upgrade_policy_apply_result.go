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

// checks if the ServerFirmwareUpgradePolicyApplyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerFirmwareUpgradePolicyApplyResult{}

// ServerFirmwareUpgradePolicyApplyResult struct for ServerFirmwareUpgradePolicyApplyResult
type ServerFirmwareUpgradePolicyApplyResult struct {
	// The list of server component ids that have been successfully scheduled for firmware upgrade.
	Scheduled []float32 `json:"scheduled"`
	// The list of server component ids that have failed to be scheduled for firmware upgrade.
	FailedToSchedule []float32 `json:"failedToSchedule"`
	AdditionalProperties map[string]interface{}
}

type _ServerFirmwareUpgradePolicyApplyResult ServerFirmwareUpgradePolicyApplyResult

// NewServerFirmwareUpgradePolicyApplyResult instantiates a new ServerFirmwareUpgradePolicyApplyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerFirmwareUpgradePolicyApplyResult(scheduled []float32, failedToSchedule []float32) *ServerFirmwareUpgradePolicyApplyResult {
	this := ServerFirmwareUpgradePolicyApplyResult{}
	this.Scheduled = scheduled
	this.FailedToSchedule = failedToSchedule
	return &this
}

// NewServerFirmwareUpgradePolicyApplyResultWithDefaults instantiates a new ServerFirmwareUpgradePolicyApplyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerFirmwareUpgradePolicyApplyResultWithDefaults() *ServerFirmwareUpgradePolicyApplyResult {
	this := ServerFirmwareUpgradePolicyApplyResult{}
	return &this
}

// GetScheduled returns the Scheduled field value
func (o *ServerFirmwareUpgradePolicyApplyResult) GetScheduled() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value
// and a boolean to check if the value has been set.
func (o *ServerFirmwareUpgradePolicyApplyResult) GetScheduledOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheduled, true
}

// SetScheduled sets field value
func (o *ServerFirmwareUpgradePolicyApplyResult) SetScheduled(v []float32) {
	o.Scheduled = v
}

// GetFailedToSchedule returns the FailedToSchedule field value
func (o *ServerFirmwareUpgradePolicyApplyResult) GetFailedToSchedule() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.FailedToSchedule
}

// GetFailedToScheduleOk returns a tuple with the FailedToSchedule field value
// and a boolean to check if the value has been set.
func (o *ServerFirmwareUpgradePolicyApplyResult) GetFailedToScheduleOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedToSchedule, true
}

// SetFailedToSchedule sets field value
func (o *ServerFirmwareUpgradePolicyApplyResult) SetFailedToSchedule(v []float32) {
	o.FailedToSchedule = v
}

func (o ServerFirmwareUpgradePolicyApplyResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerFirmwareUpgradePolicyApplyResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scheduled"] = o.Scheduled
	toSerialize["failedToSchedule"] = o.FailedToSchedule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerFirmwareUpgradePolicyApplyResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scheduled",
		"failedToSchedule",
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

	varServerFirmwareUpgradePolicyApplyResult := _ServerFirmwareUpgradePolicyApplyResult{}

	err = json.Unmarshal(data, &varServerFirmwareUpgradePolicyApplyResult)

	if err != nil {
		return err
	}

	*o = ServerFirmwareUpgradePolicyApplyResult(varServerFirmwareUpgradePolicyApplyResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scheduled")
		delete(additionalProperties, "failedToSchedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerFirmwareUpgradePolicyApplyResult struct {
	value *ServerFirmwareUpgradePolicyApplyResult
	isSet bool
}

func (v NullableServerFirmwareUpgradePolicyApplyResult) Get() *ServerFirmwareUpgradePolicyApplyResult {
	return v.value
}

func (v *NullableServerFirmwareUpgradePolicyApplyResult) Set(val *ServerFirmwareUpgradePolicyApplyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableServerFirmwareUpgradePolicyApplyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableServerFirmwareUpgradePolicyApplyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerFirmwareUpgradePolicyApplyResult(val *ServerFirmwareUpgradePolicyApplyResult) *NullableServerFirmwareUpgradePolicyApplyResult {
	return &NullableServerFirmwareUpgradePolicyApplyResult{value: val, isSet: true}
}

func (v NullableServerFirmwareUpgradePolicyApplyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerFirmwareUpgradePolicyApplyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



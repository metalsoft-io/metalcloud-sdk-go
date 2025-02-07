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

// checks if the ScheduleFirmwareUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleFirmwareUpgrade{}

// ScheduleFirmwareUpgrade struct for ScheduleFirmwareUpgrade
type ScheduleFirmwareUpgrade struct {
	// The time when the firmware update is scheduled to run.
	ScheduleUpdateTimestamp *string `json:"scheduleUpdateTimestamp,omitempty"`
	// Flag to indicate if the firmware update requires confirmation.
	ConfirmationRequired *bool `json:"confirmationRequired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleFirmwareUpgrade ScheduleFirmwareUpgrade

// NewScheduleFirmwareUpgrade instantiates a new ScheduleFirmwareUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleFirmwareUpgrade() *ScheduleFirmwareUpgrade {
	this := ScheduleFirmwareUpgrade{}
	return &this
}

// NewScheduleFirmwareUpgradeWithDefaults instantiates a new ScheduleFirmwareUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleFirmwareUpgradeWithDefaults() *ScheduleFirmwareUpgrade {
	this := ScheduleFirmwareUpgrade{}
	return &this
}

// GetScheduleUpdateTimestamp returns the ScheduleUpdateTimestamp field value if set, zero value otherwise.
func (o *ScheduleFirmwareUpgrade) GetScheduleUpdateTimestamp() string {
	if o == nil || IsNil(o.ScheduleUpdateTimestamp) {
		var ret string
		return ret
	}
	return *o.ScheduleUpdateTimestamp
}

// GetScheduleUpdateTimestampOk returns a tuple with the ScheduleUpdateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleFirmwareUpgrade) GetScheduleUpdateTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleUpdateTimestamp) {
		return nil, false
	}
	return o.ScheduleUpdateTimestamp, true
}

// HasScheduleUpdateTimestamp returns a boolean if a field has been set.
func (o *ScheduleFirmwareUpgrade) HasScheduleUpdateTimestamp() bool {
	if o != nil && !IsNil(o.ScheduleUpdateTimestamp) {
		return true
	}

	return false
}

// SetScheduleUpdateTimestamp gets a reference to the given string and assigns it to the ScheduleUpdateTimestamp field.
func (o *ScheduleFirmwareUpgrade) SetScheduleUpdateTimestamp(v string) {
	o.ScheduleUpdateTimestamp = &v
}

// GetConfirmationRequired returns the ConfirmationRequired field value if set, zero value otherwise.
func (o *ScheduleFirmwareUpgrade) GetConfirmationRequired() bool {
	if o == nil || IsNil(o.ConfirmationRequired) {
		var ret bool
		return ret
	}
	return *o.ConfirmationRequired
}

// GetConfirmationRequiredOk returns a tuple with the ConfirmationRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleFirmwareUpgrade) GetConfirmationRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfirmationRequired) {
		return nil, false
	}
	return o.ConfirmationRequired, true
}

// HasConfirmationRequired returns a boolean if a field has been set.
func (o *ScheduleFirmwareUpgrade) HasConfirmationRequired() bool {
	if o != nil && !IsNil(o.ConfirmationRequired) {
		return true
	}

	return false
}

// SetConfirmationRequired gets a reference to the given bool and assigns it to the ConfirmationRequired field.
func (o *ScheduleFirmwareUpgrade) SetConfirmationRequired(v bool) {
	o.ConfirmationRequired = &v
}

func (o ScheduleFirmwareUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleFirmwareUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScheduleUpdateTimestamp) {
		toSerialize["scheduleUpdateTimestamp"] = o.ScheduleUpdateTimestamp
	}
	if !IsNil(o.ConfirmationRequired) {
		toSerialize["confirmationRequired"] = o.ConfirmationRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduleFirmwareUpgrade) UnmarshalJSON(data []byte) (err error) {
	varScheduleFirmwareUpgrade := _ScheduleFirmwareUpgrade{}

	err = json.Unmarshal(data, &varScheduleFirmwareUpgrade)

	if err != nil {
		return err
	}

	*o = ScheduleFirmwareUpgrade(varScheduleFirmwareUpgrade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scheduleUpdateTimestamp")
		delete(additionalProperties, "confirmationRequired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduleFirmwareUpgrade struct {
	value *ScheduleFirmwareUpgrade
	isSet bool
}

func (v NullableScheduleFirmwareUpgrade) Get() *ScheduleFirmwareUpgrade {
	return v.value
}

func (v *NullableScheduleFirmwareUpgrade) Set(val *ScheduleFirmwareUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleFirmwareUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleFirmwareUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleFirmwareUpgrade(val *ScheduleFirmwareUpgrade) *NullableScheduleFirmwareUpgrade {
	return &NullableScheduleFirmwareUpgrade{value: val, isSet: true}
}

func (v NullableScheduleFirmwareUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleFirmwareUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the JobCommandInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobCommandInfo{}

// JobCommandInfo struct for JobCommandInfo
type JobCommandInfo struct {
	// Mark the job for death
	Command *string `json:"command,omitempty"`
	// Execute the command immediately
	ExecuteImmediately *bool `json:"executeImmediately,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JobCommandInfo JobCommandInfo

// NewJobCommandInfo instantiates a new JobCommandInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobCommandInfo() *JobCommandInfo {
	this := JobCommandInfo{}
	var executeImmediately bool = false
	this.ExecuteImmediately = &executeImmediately
	return &this
}

// NewJobCommandInfoWithDefaults instantiates a new JobCommandInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobCommandInfoWithDefaults() *JobCommandInfo {
	this := JobCommandInfo{}
	var executeImmediately bool = false
	this.ExecuteImmediately = &executeImmediately
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *JobCommandInfo) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobCommandInfo) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *JobCommandInfo) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *JobCommandInfo) SetCommand(v string) {
	o.Command = &v
}

// GetExecuteImmediately returns the ExecuteImmediately field value if set, zero value otherwise.
func (o *JobCommandInfo) GetExecuteImmediately() bool {
	if o == nil || IsNil(o.ExecuteImmediately) {
		var ret bool
		return ret
	}
	return *o.ExecuteImmediately
}

// GetExecuteImmediatelyOk returns a tuple with the ExecuteImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobCommandInfo) GetExecuteImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecuteImmediately) {
		return nil, false
	}
	return o.ExecuteImmediately, true
}

// HasExecuteImmediately returns a boolean if a field has been set.
func (o *JobCommandInfo) HasExecuteImmediately() bool {
	if o != nil && !IsNil(o.ExecuteImmediately) {
		return true
	}

	return false
}

// SetExecuteImmediately gets a reference to the given bool and assigns it to the ExecuteImmediately field.
func (o *JobCommandInfo) SetExecuteImmediately(v bool) {
	o.ExecuteImmediately = &v
}

func (o JobCommandInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobCommandInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.ExecuteImmediately) {
		toSerialize["executeImmediately"] = o.ExecuteImmediately
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobCommandInfo) UnmarshalJSON(data []byte) (err error) {
	varJobCommandInfo := _JobCommandInfo{}

	err = json.Unmarshal(data, &varJobCommandInfo)

	if err != nil {
		return err
	}

	*o = JobCommandInfo(varJobCommandInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "command")
		delete(additionalProperties, "executeImmediately")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobCommandInfo struct {
	value *JobCommandInfo
	isSet bool
}

func (v NullableJobCommandInfo) Get() *JobCommandInfo {
	return v.value
}

func (v *NullableJobCommandInfo) Set(val *JobCommandInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJobCommandInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJobCommandInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobCommandInfo(val *JobCommandInfo) *NullableJobCommandInfo {
	return &NullableJobCommandInfo{value: val, isSet: true}
}

func (v NullableJobCommandInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobCommandInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



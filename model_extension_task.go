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

// checks if the ExtensionTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionTask{}

// ExtensionTask struct for ExtensionTask
type ExtensionTask struct {
	// Label of the task.
	Label string `json:"label"`
	// Type of the task.
	TaskType string `json:"taskType"`
	Options ExtensionTaskOptions `json:"options"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionTask ExtensionTask

// NewExtensionTask instantiates a new ExtensionTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionTask(label string, taskType string, options ExtensionTaskOptions) *ExtensionTask {
	this := ExtensionTask{}
	this.Label = label
	this.TaskType = taskType
	this.Options = options
	return &this
}

// NewExtensionTaskWithDefaults instantiates a new ExtensionTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionTaskWithDefaults() *ExtensionTask {
	this := ExtensionTask{}
	return &this
}

// GetLabel returns the Label field value
func (o *ExtensionTask) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ExtensionTask) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ExtensionTask) SetLabel(v string) {
	o.Label = v
}

// GetTaskType returns the TaskType field value
func (o *ExtensionTask) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *ExtensionTask) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value
func (o *ExtensionTask) SetTaskType(v string) {
	o.TaskType = v
}

// GetOptions returns the Options field value
func (o *ExtensionTask) GetOptions() ExtensionTaskOptions {
	if o == nil {
		var ret ExtensionTaskOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ExtensionTask) GetOptionsOk() (*ExtensionTaskOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ExtensionTask) SetOptions(v ExtensionTaskOptions) {
	o.Options = v
}

func (o ExtensionTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["taskType"] = o.TaskType
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"taskType",
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

	varExtensionTask := _ExtensionTask{}

	err = json.Unmarshal(data, &varExtensionTask)

	if err != nil {
		return err
	}

	*o = ExtensionTask(varExtensionTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "taskType")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionTask struct {
	value *ExtensionTask
	isSet bool
}

func (v NullableExtensionTask) Get() *ExtensionTask {
	return v.value
}

func (v *NullableExtensionTask) Set(val *ExtensionTask) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionTask) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionTask(val *ExtensionTask) *NullableExtensionTask {
	return &NullableExtensionTask{value: val, isSet: true}
}

func (v NullableExtensionTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



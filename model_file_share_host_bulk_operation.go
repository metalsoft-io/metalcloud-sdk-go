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

// checks if the FileShareHostBulkOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileShareHostBulkOperation{}

// FileShareHostBulkOperation struct for FileShareHostBulkOperation
type FileShareHostBulkOperation struct {
	// Id of the Instance Array Host that will be modified
	InstanceArrayId float32 `json:"instanceArrayId"`
	// Operation type for the Instance Array Host
	OperationType string `json:"operationType"`
	AdditionalProperties map[string]interface{}
}

type _FileShareHostBulkOperation FileShareHostBulkOperation

// NewFileShareHostBulkOperation instantiates a new FileShareHostBulkOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileShareHostBulkOperation(instanceArrayId float32, operationType string) *FileShareHostBulkOperation {
	this := FileShareHostBulkOperation{}
	this.InstanceArrayId = instanceArrayId
	this.OperationType = operationType
	return &this
}

// NewFileShareHostBulkOperationWithDefaults instantiates a new FileShareHostBulkOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileShareHostBulkOperationWithDefaults() *FileShareHostBulkOperation {
	this := FileShareHostBulkOperation{}
	return &this
}

// GetInstanceArrayId returns the InstanceArrayId field value
func (o *FileShareHostBulkOperation) GetInstanceArrayId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InstanceArrayId
}

// GetInstanceArrayIdOk returns a tuple with the InstanceArrayId field value
// and a boolean to check if the value has been set.
func (o *FileShareHostBulkOperation) GetInstanceArrayIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceArrayId, true
}

// SetInstanceArrayId sets field value
func (o *FileShareHostBulkOperation) SetInstanceArrayId(v float32) {
	o.InstanceArrayId = v
}

// GetOperationType returns the OperationType field value
func (o *FileShareHostBulkOperation) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *FileShareHostBulkOperation) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *FileShareHostBulkOperation) SetOperationType(v string) {
	o.OperationType = v
}

func (o FileShareHostBulkOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileShareHostBulkOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceArrayId"] = o.InstanceArrayId
	toSerialize["operationType"] = o.OperationType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileShareHostBulkOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instanceArrayId",
		"operationType",
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

	varFileShareHostBulkOperation := _FileShareHostBulkOperation{}

	err = json.Unmarshal(data, &varFileShareHostBulkOperation)

	if err != nil {
		return err
	}

	*o = FileShareHostBulkOperation(varFileShareHostBulkOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instanceArrayId")
		delete(additionalProperties, "operationType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileShareHostBulkOperation struct {
	value *FileShareHostBulkOperation
	isSet bool
}

func (v NullableFileShareHostBulkOperation) Get() *FileShareHostBulkOperation {
	return v.value
}

func (v *NullableFileShareHostBulkOperation) Set(val *FileShareHostBulkOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableFileShareHostBulkOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableFileShareHostBulkOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileShareHostBulkOperation(val *FileShareHostBulkOperation) *NullableFileShareHostBulkOperation {
	return &NullableFileShareHostBulkOperation{value: val, isSet: true}
}

func (v NullableFileShareHostBulkOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileShareHostBulkOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



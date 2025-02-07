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

// checks if the ServerTypeDiskGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerTypeDiskGroup{}

// ServerTypeDiskGroup struct for ServerTypeDiskGroup
type ServerTypeDiskGroup struct {
	// The storage controllers for the server type. Key is the controller name and value is the controller information.
	StorageControllers map[string]interface{} `json:"storageControllers"`
	AdditionalProperties map[string]interface{}
}

type _ServerTypeDiskGroup ServerTypeDiskGroup

// NewServerTypeDiskGroup instantiates a new ServerTypeDiskGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerTypeDiskGroup(storageControllers map[string]interface{}) *ServerTypeDiskGroup {
	this := ServerTypeDiskGroup{}
	this.StorageControllers = storageControllers
	return &this
}

// NewServerTypeDiskGroupWithDefaults instantiates a new ServerTypeDiskGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerTypeDiskGroupWithDefaults() *ServerTypeDiskGroup {
	this := ServerTypeDiskGroup{}
	return &this
}

// GetStorageControllers returns the StorageControllers field value
func (o *ServerTypeDiskGroup) GetStorageControllers() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.StorageControllers
}

// GetStorageControllersOk returns a tuple with the StorageControllers field value
// and a boolean to check if the value has been set.
func (o *ServerTypeDiskGroup) GetStorageControllersOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.StorageControllers, true
}

// SetStorageControllers sets field value
func (o *ServerTypeDiskGroup) SetStorageControllers(v map[string]interface{}) {
	o.StorageControllers = v
}

func (o ServerTypeDiskGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerTypeDiskGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storageControllers"] = o.StorageControllers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerTypeDiskGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storageControllers",
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

	varServerTypeDiskGroup := _ServerTypeDiskGroup{}

	err = json.Unmarshal(data, &varServerTypeDiskGroup)

	if err != nil {
		return err
	}

	*o = ServerTypeDiskGroup(varServerTypeDiskGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "storageControllers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerTypeDiskGroup struct {
	value *ServerTypeDiskGroup
	isSet bool
}

func (v NullableServerTypeDiskGroup) Get() *ServerTypeDiskGroup {
	return v.value
}

func (v *NullableServerTypeDiskGroup) Set(val *ServerTypeDiskGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableServerTypeDiskGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableServerTypeDiskGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerTypeDiskGroup(val *ServerTypeDiskGroup) *NullableServerTypeDiskGroup {
	return &NullableServerTypeDiskGroup{value: val, isSet: true}
}

func (v NullableServerTypeDiskGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerTypeDiskGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



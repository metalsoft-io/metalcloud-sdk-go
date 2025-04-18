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

// checks if the StorageStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageStatistics{}

// StorageStatistics struct for StorageStatistics
type StorageStatistics struct {
	// Number of storages in maintenance
	MaintenanceCount float32 `json:"maintenanceCount"`
	// Number of experimental storages
	ExperimentalCount float32 `json:"experimentalCount"`
	// Number of storages low on space
	LowSpaceCount float32 `json:"lowSpaceCount"`
	// Total used space across all storages
	UsedSpace float32 `json:"usedSpace"`
	// Total free space across all storages
	FreeSpace float32 `json:"freeSpace"`
	// Number of storages by type
	Types map[string]interface{} `json:"types"`
	// Number of storages in pending state
	PendingCount float32 `json:"pendingCount"`
	// Number of storages in ready state
	ReadyCount float32 `json:"readyCount"`
	// Number of storages in active state
	ActiveCount float32 `json:"activeCount"`
	AdditionalProperties map[string]interface{}
}

type _StorageStatistics StorageStatistics

// NewStorageStatistics instantiates a new StorageStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStatistics(maintenanceCount float32, experimentalCount float32, lowSpaceCount float32, usedSpace float32, freeSpace float32, types map[string]interface{}, pendingCount float32, readyCount float32, activeCount float32) *StorageStatistics {
	this := StorageStatistics{}
	this.MaintenanceCount = maintenanceCount
	this.ExperimentalCount = experimentalCount
	this.LowSpaceCount = lowSpaceCount
	this.UsedSpace = usedSpace
	this.FreeSpace = freeSpace
	this.Types = types
	this.PendingCount = pendingCount
	this.ReadyCount = readyCount
	this.ActiveCount = activeCount
	return &this
}

// NewStorageStatisticsWithDefaults instantiates a new StorageStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStatisticsWithDefaults() *StorageStatistics {
	this := StorageStatistics{}
	return &this
}

// GetMaintenanceCount returns the MaintenanceCount field value
func (o *StorageStatistics) GetMaintenanceCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaintenanceCount
}

// GetMaintenanceCountOk returns a tuple with the MaintenanceCount field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetMaintenanceCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenanceCount, true
}

// SetMaintenanceCount sets field value
func (o *StorageStatistics) SetMaintenanceCount(v float32) {
	o.MaintenanceCount = v
}

// GetExperimentalCount returns the ExperimentalCount field value
func (o *StorageStatistics) GetExperimentalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExperimentalCount
}

// GetExperimentalCountOk returns a tuple with the ExperimentalCount field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetExperimentalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperimentalCount, true
}

// SetExperimentalCount sets field value
func (o *StorageStatistics) SetExperimentalCount(v float32) {
	o.ExperimentalCount = v
}

// GetLowSpaceCount returns the LowSpaceCount field value
func (o *StorageStatistics) GetLowSpaceCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LowSpaceCount
}

// GetLowSpaceCountOk returns a tuple with the LowSpaceCount field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetLowSpaceCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowSpaceCount, true
}

// SetLowSpaceCount sets field value
func (o *StorageStatistics) SetLowSpaceCount(v float32) {
	o.LowSpaceCount = v
}

// GetUsedSpace returns the UsedSpace field value
func (o *StorageStatistics) GetUsedSpace() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsedSpace
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetUsedSpaceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedSpace, true
}

// SetUsedSpace sets field value
func (o *StorageStatistics) SetUsedSpace(v float32) {
	o.UsedSpace = v
}

// GetFreeSpace returns the FreeSpace field value
func (o *StorageStatistics) GetFreeSpace() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FreeSpace
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetFreeSpaceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeSpace, true
}

// SetFreeSpace sets field value
func (o *StorageStatistics) SetFreeSpace(v float32) {
	o.FreeSpace = v
}

// GetTypes returns the Types field value
func (o *StorageStatistics) GetTypes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetTypesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *StorageStatistics) SetTypes(v map[string]interface{}) {
	o.Types = v
}

// GetPendingCount returns the PendingCount field value
func (o *StorageStatistics) GetPendingCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetPendingCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingCount, true
}

// SetPendingCount sets field value
func (o *StorageStatistics) SetPendingCount(v float32) {
	o.PendingCount = v
}

// GetReadyCount returns the ReadyCount field value
func (o *StorageStatistics) GetReadyCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ReadyCount
}

// GetReadyCountOk returns a tuple with the ReadyCount field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetReadyCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadyCount, true
}

// SetReadyCount sets field value
func (o *StorageStatistics) SetReadyCount(v float32) {
	o.ReadyCount = v
}

// GetActiveCount returns the ActiveCount field value
func (o *StorageStatistics) GetActiveCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ActiveCount
}

// GetActiveCountOk returns a tuple with the ActiveCount field value
// and a boolean to check if the value has been set.
func (o *StorageStatistics) GetActiveCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveCount, true
}

// SetActiveCount sets field value
func (o *StorageStatistics) SetActiveCount(v float32) {
	o.ActiveCount = v
}

func (o StorageStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maintenanceCount"] = o.MaintenanceCount
	toSerialize["experimentalCount"] = o.ExperimentalCount
	toSerialize["lowSpaceCount"] = o.LowSpaceCount
	toSerialize["usedSpace"] = o.UsedSpace
	toSerialize["freeSpace"] = o.FreeSpace
	toSerialize["types"] = o.Types
	toSerialize["pendingCount"] = o.PendingCount
	toSerialize["readyCount"] = o.ReadyCount
	toSerialize["activeCount"] = o.ActiveCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"maintenanceCount",
		"experimentalCount",
		"lowSpaceCount",
		"usedSpace",
		"freeSpace",
		"types",
		"pendingCount",
		"readyCount",
		"activeCount",
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

	varStorageStatistics := _StorageStatistics{}

	err = json.Unmarshal(data, &varStorageStatistics)

	if err != nil {
		return err
	}

	*o = StorageStatistics(varStorageStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maintenanceCount")
		delete(additionalProperties, "experimentalCount")
		delete(additionalProperties, "lowSpaceCount")
		delete(additionalProperties, "usedSpace")
		delete(additionalProperties, "freeSpace")
		delete(additionalProperties, "types")
		delete(additionalProperties, "pendingCount")
		delete(additionalProperties, "readyCount")
		delete(additionalProperties, "activeCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageStatistics struct {
	value *StorageStatistics
	isSet bool
}

func (v NullableStorageStatistics) Get() *StorageStatistics {
	return v.value
}

func (v *NullableStorageStatistics) Set(val *StorageStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStatistics(val *StorageStatistics) *NullableStorageStatistics {
	return &NullableStorageStatistics{value: val, isSet: true}
}

func (v NullableStorageStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



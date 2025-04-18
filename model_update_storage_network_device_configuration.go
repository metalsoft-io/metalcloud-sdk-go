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

// checks if the UpdateStorageNetworkDeviceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorageNetworkDeviceConfiguration{}

// UpdateStorageNetworkDeviceConfiguration struct for UpdateStorageNetworkDeviceConfiguration
type UpdateStorageNetworkDeviceConfiguration struct {
	// Identifier of the storage physical interface
	StoragePhysicalInterfaceIdentifier *string `json:"storagePhysicalInterfaceIdentifier,omitempty"`
	// Identifier of the network device interface
	NetworkDeviceInterfaceIdentifier *string `json:"networkDeviceInterfaceIdentifier,omitempty"`
	// Array of VLANS for the network device interface
	NetworkDeviceInterfaceVlans []float32 `json:"networkDeviceInterfaceVlans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateStorageNetworkDeviceConfiguration UpdateStorageNetworkDeviceConfiguration

// NewUpdateStorageNetworkDeviceConfiguration instantiates a new UpdateStorageNetworkDeviceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorageNetworkDeviceConfiguration() *UpdateStorageNetworkDeviceConfiguration {
	this := UpdateStorageNetworkDeviceConfiguration{}
	return &this
}

// NewUpdateStorageNetworkDeviceConfigurationWithDefaults instantiates a new UpdateStorageNetworkDeviceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageNetworkDeviceConfigurationWithDefaults() *UpdateStorageNetworkDeviceConfiguration {
	this := UpdateStorageNetworkDeviceConfiguration{}
	return &this
}

// GetStoragePhysicalInterfaceIdentifier returns the StoragePhysicalInterfaceIdentifier field value if set, zero value otherwise.
func (o *UpdateStorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifier() string {
	if o == nil || IsNil(o.StoragePhysicalInterfaceIdentifier) {
		var ret string
		return ret
	}
	return *o.StoragePhysicalInterfaceIdentifier
}

// GetStoragePhysicalInterfaceIdentifierOk returns a tuple with the StoragePhysicalInterfaceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.StoragePhysicalInterfaceIdentifier) {
		return nil, false
	}
	return o.StoragePhysicalInterfaceIdentifier, true
}

// HasStoragePhysicalInterfaceIdentifier returns a boolean if a field has been set.
func (o *UpdateStorageNetworkDeviceConfiguration) HasStoragePhysicalInterfaceIdentifier() bool {
	if o != nil && !IsNil(o.StoragePhysicalInterfaceIdentifier) {
		return true
	}

	return false
}

// SetStoragePhysicalInterfaceIdentifier gets a reference to the given string and assigns it to the StoragePhysicalInterfaceIdentifier field.
func (o *UpdateStorageNetworkDeviceConfiguration) SetStoragePhysicalInterfaceIdentifier(v string) {
	o.StoragePhysicalInterfaceIdentifier = &v
}

// GetNetworkDeviceInterfaceIdentifier returns the NetworkDeviceInterfaceIdentifier field value if set, zero value otherwise.
func (o *UpdateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceIdentifier() string {
	if o == nil || IsNil(o.NetworkDeviceInterfaceIdentifier) {
		var ret string
		return ret
	}
	return *o.NetworkDeviceInterfaceIdentifier
}

// GetNetworkDeviceInterfaceIdentifierOk returns a tuple with the NetworkDeviceInterfaceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkDeviceInterfaceIdentifier) {
		return nil, false
	}
	return o.NetworkDeviceInterfaceIdentifier, true
}

// HasNetworkDeviceInterfaceIdentifier returns a boolean if a field has been set.
func (o *UpdateStorageNetworkDeviceConfiguration) HasNetworkDeviceInterfaceIdentifier() bool {
	if o != nil && !IsNil(o.NetworkDeviceInterfaceIdentifier) {
		return true
	}

	return false
}

// SetNetworkDeviceInterfaceIdentifier gets a reference to the given string and assigns it to the NetworkDeviceInterfaceIdentifier field.
func (o *UpdateStorageNetworkDeviceConfiguration) SetNetworkDeviceInterfaceIdentifier(v string) {
	o.NetworkDeviceInterfaceIdentifier = &v
}

// GetNetworkDeviceInterfaceVlans returns the NetworkDeviceInterfaceVlans field value if set, zero value otherwise.
func (o *UpdateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceVlans() []float32 {
	if o == nil || IsNil(o.NetworkDeviceInterfaceVlans) {
		var ret []float32
		return ret
	}
	return o.NetworkDeviceInterfaceVlans
}

// GetNetworkDeviceInterfaceVlansOk returns a tuple with the NetworkDeviceInterfaceVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceVlansOk() ([]float32, bool) {
	if o == nil || IsNil(o.NetworkDeviceInterfaceVlans) {
		return nil, false
	}
	return o.NetworkDeviceInterfaceVlans, true
}

// HasNetworkDeviceInterfaceVlans returns a boolean if a field has been set.
func (o *UpdateStorageNetworkDeviceConfiguration) HasNetworkDeviceInterfaceVlans() bool {
	if o != nil && !IsNil(o.NetworkDeviceInterfaceVlans) {
		return true
	}

	return false
}

// SetNetworkDeviceInterfaceVlans gets a reference to the given []float32 and assigns it to the NetworkDeviceInterfaceVlans field.
func (o *UpdateStorageNetworkDeviceConfiguration) SetNetworkDeviceInterfaceVlans(v []float32) {
	o.NetworkDeviceInterfaceVlans = v
}

func (o UpdateStorageNetworkDeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorageNetworkDeviceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StoragePhysicalInterfaceIdentifier) {
		toSerialize["storagePhysicalInterfaceIdentifier"] = o.StoragePhysicalInterfaceIdentifier
	}
	if !IsNil(o.NetworkDeviceInterfaceIdentifier) {
		toSerialize["networkDeviceInterfaceIdentifier"] = o.NetworkDeviceInterfaceIdentifier
	}
	if !IsNil(o.NetworkDeviceInterfaceVlans) {
		toSerialize["networkDeviceInterfaceVlans"] = o.NetworkDeviceInterfaceVlans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateStorageNetworkDeviceConfiguration) UnmarshalJSON(data []byte) (err error) {
	varUpdateStorageNetworkDeviceConfiguration := _UpdateStorageNetworkDeviceConfiguration{}

	err = json.Unmarshal(data, &varUpdateStorageNetworkDeviceConfiguration)

	if err != nil {
		return err
	}

	*o = UpdateStorageNetworkDeviceConfiguration(varUpdateStorageNetworkDeviceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "storagePhysicalInterfaceIdentifier")
		delete(additionalProperties, "networkDeviceInterfaceIdentifier")
		delete(additionalProperties, "networkDeviceInterfaceVlans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateStorageNetworkDeviceConfiguration struct {
	value *UpdateStorageNetworkDeviceConfiguration
	isSet bool
}

func (v NullableUpdateStorageNetworkDeviceConfiguration) Get() *UpdateStorageNetworkDeviceConfiguration {
	return v.value
}

func (v *NullableUpdateStorageNetworkDeviceConfiguration) Set(val *UpdateStorageNetworkDeviceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStorageNetworkDeviceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStorageNetworkDeviceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStorageNetworkDeviceConfiguration(val *UpdateStorageNetworkDeviceConfiguration) *NullableUpdateStorageNetworkDeviceConfiguration {
	return &NullableUpdateStorageNetworkDeviceConfiguration{value: val, isSet: true}
}

func (v NullableUpdateStorageNetworkDeviceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStorageNetworkDeviceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



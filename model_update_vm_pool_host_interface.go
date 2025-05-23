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

// checks if the UpdateVMPoolHostInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVMPoolHostInterface{}

// UpdateVMPoolHostInterface struct for UpdateVMPoolHostInterface
type UpdateVMPoolHostInterface struct {
	// Fabric of the switch interface
	Fabric *string `json:"fabric,omitempty"`
	// Id of the switch
	SwitchId *float32 `json:"switchId,omitempty"`
	// Name of the switch interface
	SwitchInterfaceName *float32 `json:"switchInterfaceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateVMPoolHostInterface UpdateVMPoolHostInterface

// NewUpdateVMPoolHostInterface instantiates a new UpdateVMPoolHostInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVMPoolHostInterface() *UpdateVMPoolHostInterface {
	this := UpdateVMPoolHostInterface{}
	return &this
}

// NewUpdateVMPoolHostInterfaceWithDefaults instantiates a new UpdateVMPoolHostInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVMPoolHostInterfaceWithDefaults() *UpdateVMPoolHostInterface {
	this := UpdateVMPoolHostInterface{}
	return &this
}

// GetFabric returns the Fabric field value if set, zero value otherwise.
func (o *UpdateVMPoolHostInterface) GetFabric() string {
	if o == nil || IsNil(o.Fabric) {
		var ret string
		return ret
	}
	return *o.Fabric
}

// GetFabricOk returns a tuple with the Fabric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPoolHostInterface) GetFabricOk() (*string, bool) {
	if o == nil || IsNil(o.Fabric) {
		return nil, false
	}
	return o.Fabric, true
}

// HasFabric returns a boolean if a field has been set.
func (o *UpdateVMPoolHostInterface) HasFabric() bool {
	if o != nil && !IsNil(o.Fabric) {
		return true
	}

	return false
}

// SetFabric gets a reference to the given string and assigns it to the Fabric field.
func (o *UpdateVMPoolHostInterface) SetFabric(v string) {
	o.Fabric = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *UpdateVMPoolHostInterface) GetSwitchId() float32 {
	if o == nil || IsNil(o.SwitchId) {
		var ret float32
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPoolHostInterface) GetSwitchIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *UpdateVMPoolHostInterface) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given float32 and assigns it to the SwitchId field.
func (o *UpdateVMPoolHostInterface) SetSwitchId(v float32) {
	o.SwitchId = &v
}

// GetSwitchInterfaceName returns the SwitchInterfaceName field value if set, zero value otherwise.
func (o *UpdateVMPoolHostInterface) GetSwitchInterfaceName() float32 {
	if o == nil || IsNil(o.SwitchInterfaceName) {
		var ret float32
		return ret
	}
	return *o.SwitchInterfaceName
}

// GetSwitchInterfaceNameOk returns a tuple with the SwitchInterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPoolHostInterface) GetSwitchInterfaceNameOk() (*float32, bool) {
	if o == nil || IsNil(o.SwitchInterfaceName) {
		return nil, false
	}
	return o.SwitchInterfaceName, true
}

// HasSwitchInterfaceName returns a boolean if a field has been set.
func (o *UpdateVMPoolHostInterface) HasSwitchInterfaceName() bool {
	if o != nil && !IsNil(o.SwitchInterfaceName) {
		return true
	}

	return false
}

// SetSwitchInterfaceName gets a reference to the given float32 and assigns it to the SwitchInterfaceName field.
func (o *UpdateVMPoolHostInterface) SetSwitchInterfaceName(v float32) {
	o.SwitchInterfaceName = &v
}

func (o UpdateVMPoolHostInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVMPoolHostInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fabric) {
		toSerialize["fabric"] = o.Fabric
	}
	if !IsNil(o.SwitchId) {
		toSerialize["switchId"] = o.SwitchId
	}
	if !IsNil(o.SwitchInterfaceName) {
		toSerialize["switchInterfaceName"] = o.SwitchInterfaceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateVMPoolHostInterface) UnmarshalJSON(data []byte) (err error) {
	varUpdateVMPoolHostInterface := _UpdateVMPoolHostInterface{}

	err = json.Unmarshal(data, &varUpdateVMPoolHostInterface)

	if err != nil {
		return err
	}

	*o = UpdateVMPoolHostInterface(varUpdateVMPoolHostInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fabric")
		delete(additionalProperties, "switchId")
		delete(additionalProperties, "switchInterfaceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateVMPoolHostInterface struct {
	value *UpdateVMPoolHostInterface
	isSet bool
}

func (v NullableUpdateVMPoolHostInterface) Get() *UpdateVMPoolHostInterface {
	return v.value
}

func (v *NullableUpdateVMPoolHostInterface) Set(val *UpdateVMPoolHostInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVMPoolHostInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVMPoolHostInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVMPoolHostInterface(val *UpdateVMPoolHostInterface) *NullableUpdateVMPoolHostInterface {
	return &NullableUpdateVMPoolHostInterface{value: val, isSet: true}
}

func (v NullableUpdateVMPoolHostInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVMPoolHostInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



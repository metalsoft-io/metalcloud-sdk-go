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

// checks if the CreateVMInstanceGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMInstanceGroup{}

// CreateVMInstanceGroup struct for CreateVMInstanceGroup
type CreateVMInstanceGroup struct {
	InstanceCount *float32 `json:"instanceCount,omitempty"`
	// Disk size in GB for each VM Instance in the VM Instance Group.
	DiskSizeGB float32 `json:"diskSizeGB"`
	// Id of the VM Type.
	TypeId float32 `json:"typeId"`
	// Id of the template used by the VM Instance Group.
	OsTemplateId float32 `json:"osTemplateId"`
	// Tags for the VM Instance Group.
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateVMInstanceGroup CreateVMInstanceGroup

// NewCreateVMInstanceGroup instantiates a new CreateVMInstanceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMInstanceGroup(diskSizeGB float32, typeId float32, osTemplateId float32) *CreateVMInstanceGroup {
	this := CreateVMInstanceGroup{}
	var instanceCount float32 = 1
	this.InstanceCount = &instanceCount
	this.DiskSizeGB = diskSizeGB
	this.TypeId = typeId
	this.OsTemplateId = osTemplateId
	return &this
}

// NewCreateVMInstanceGroupWithDefaults instantiates a new CreateVMInstanceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMInstanceGroupWithDefaults() *CreateVMInstanceGroup {
	this := CreateVMInstanceGroup{}
	var instanceCount float32 = 1
	this.InstanceCount = &instanceCount
	return &this
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *CreateVMInstanceGroup) GetInstanceCount() float32 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret float32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMInstanceGroup) GetInstanceCountOk() (*float32, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *CreateVMInstanceGroup) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given float32 and assigns it to the InstanceCount field.
func (o *CreateVMInstanceGroup) SetInstanceCount(v float32) {
	o.InstanceCount = &v
}

// GetDiskSizeGB returns the DiskSizeGB field value
func (o *CreateVMInstanceGroup) GetDiskSizeGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value
// and a boolean to check if the value has been set.
func (o *CreateVMInstanceGroup) GetDiskSizeGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskSizeGB, true
}

// SetDiskSizeGB sets field value
func (o *CreateVMInstanceGroup) SetDiskSizeGB(v float32) {
	o.DiskSizeGB = v
}

// GetTypeId returns the TypeId field value
func (o *CreateVMInstanceGroup) GetTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *CreateVMInstanceGroup) GetTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *CreateVMInstanceGroup) SetTypeId(v float32) {
	o.TypeId = v
}

// GetOsTemplateId returns the OsTemplateId field value
func (o *CreateVMInstanceGroup) GetOsTemplateId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OsTemplateId
}

// GetOsTemplateIdOk returns a tuple with the OsTemplateId field value
// and a boolean to check if the value has been set.
func (o *CreateVMInstanceGroup) GetOsTemplateIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsTemplateId, true
}

// SetOsTemplateId sets field value
func (o *CreateVMInstanceGroup) SetOsTemplateId(v float32) {
	o.OsTemplateId = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateVMInstanceGroup) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMInstanceGroup) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateVMInstanceGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateVMInstanceGroup) SetTags(v []string) {
	o.Tags = v
}

func (o CreateVMInstanceGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMInstanceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceCount) {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	toSerialize["diskSizeGB"] = o.DiskSizeGB
	toSerialize["typeId"] = o.TypeId
	toSerialize["osTemplateId"] = o.OsTemplateId
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateVMInstanceGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"diskSizeGB",
		"typeId",
		"osTemplateId",
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

	varCreateVMInstanceGroup := _CreateVMInstanceGroup{}

	err = json.Unmarshal(data, &varCreateVMInstanceGroup)

	if err != nil {
		return err
	}

	*o = CreateVMInstanceGroup(varCreateVMInstanceGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "diskSizeGB")
		delete(additionalProperties, "typeId")
		delete(additionalProperties, "osTemplateId")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateVMInstanceGroup struct {
	value *CreateVMInstanceGroup
	isSet bool
}

func (v NullableCreateVMInstanceGroup) Get() *CreateVMInstanceGroup {
	return v.value
}

func (v *NullableCreateVMInstanceGroup) Set(val *CreateVMInstanceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMInstanceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMInstanceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMInstanceGroup(val *CreateVMInstanceGroup) *NullableCreateVMInstanceGroup {
	return &NullableCreateVMInstanceGroup{value: val, isSet: true}
}

func (v NullableCreateVMInstanceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMInstanceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



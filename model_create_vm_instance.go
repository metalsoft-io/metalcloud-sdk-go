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

// checks if the CreateVMInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMInstance{}

// CreateVMInstance struct for CreateVMInstance
type CreateVMInstance struct {
	// Id of the VM Instance Group.
	GroupId float32 `json:"groupId"`
	// Id of the VM Type.
	TypeId float32 `json:"typeId"`
	// Tags for the VM Instance.
	Tags []string `json:"tags,omitempty"`
	// Disk size in GB of the VM Instance. If not passed, the default disk size from the group will be used
	DiskSizeGB *float32 `json:"diskSizeGB,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateVMInstance CreateVMInstance

// NewCreateVMInstance instantiates a new CreateVMInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMInstance(groupId float32, typeId float32) *CreateVMInstance {
	this := CreateVMInstance{}
	this.GroupId = groupId
	this.TypeId = typeId
	return &this
}

// NewCreateVMInstanceWithDefaults instantiates a new CreateVMInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMInstanceWithDefaults() *CreateVMInstance {
	this := CreateVMInstance{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *CreateVMInstance) GetGroupId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *CreateVMInstance) GetGroupIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *CreateVMInstance) SetGroupId(v float32) {
	o.GroupId = v
}

// GetTypeId returns the TypeId field value
func (o *CreateVMInstance) GetTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *CreateVMInstance) GetTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *CreateVMInstance) SetTypeId(v float32) {
	o.TypeId = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateVMInstance) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMInstance) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateVMInstance) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateVMInstance) SetTags(v []string) {
	o.Tags = v
}

// GetDiskSizeGB returns the DiskSizeGB field value if set, zero value otherwise.
func (o *CreateVMInstance) GetDiskSizeGB() float32 {
	if o == nil || IsNil(o.DiskSizeGB) {
		var ret float32
		return ret
	}
	return *o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMInstance) GetDiskSizeGBOk() (*float32, bool) {
	if o == nil || IsNil(o.DiskSizeGB) {
		return nil, false
	}
	return o.DiskSizeGB, true
}

// HasDiskSizeGB returns a boolean if a field has been set.
func (o *CreateVMInstance) HasDiskSizeGB() bool {
	if o != nil && !IsNil(o.DiskSizeGB) {
		return true
	}

	return false
}

// SetDiskSizeGB gets a reference to the given float32 and assigns it to the DiskSizeGB field.
func (o *CreateVMInstance) SetDiskSizeGB(v float32) {
	o.DiskSizeGB = &v
}

func (o CreateVMInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["typeId"] = o.TypeId
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.DiskSizeGB) {
		toSerialize["diskSizeGB"] = o.DiskSizeGB
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateVMInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"typeId",
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

	varCreateVMInstance := _CreateVMInstance{}

	err = json.Unmarshal(data, &varCreateVMInstance)

	if err != nil {
		return err
	}

	*o = CreateVMInstance(varCreateVMInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "typeId")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "diskSizeGB")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateVMInstance struct {
	value *CreateVMInstance
	isSet bool
}

func (v NullableCreateVMInstance) Get() *CreateVMInstance {
	return v.value
}

func (v *NullableCreateVMInstance) Set(val *CreateVMInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMInstance(val *CreateVMInstance) *NullableCreateVMInstance {
	return &NullableCreateVMInstance{value: val, isSet: true}
}

func (v NullableCreateVMInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



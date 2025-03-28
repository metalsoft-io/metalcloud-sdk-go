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

// checks if the CreateVMType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMType{}

// CreateVMType struct for CreateVMType
type CreateVMType struct {
	// Name of the VM Pool type
	Name string `json:"name"`
	// Display name of the VM Pool type
	DisplayName *string `json:"displayName,omitempty"`
	// Label of the VM Pool type
	Label *string `json:"label,omitempty"`
	// Number of CPU cores for the VM Pool type
	CpuCores float32 `json:"cpuCores"`
	// RAM in GB for the VM Pool type
	RamGB float32 `json:"ramGB"`
	// Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0.
	IsExperimental *float32 `json:"isExperimental,omitempty"`
	// Tags for the VM Type.
	Tags []string `json:"tags,omitempty"`
	// Flag to indicate if the VM Pool is for unmanaged VMs only. 1 for true, 0 for false. Default is 0.
	ForUnmanagedVMsOnly *float32 `json:"forUnmanagedVMsOnly,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateVMType CreateVMType

// NewCreateVMType instantiates a new CreateVMType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMType(name string, cpuCores float32, ramGB float32) *CreateVMType {
	this := CreateVMType{}
	this.Name = name
	this.CpuCores = cpuCores
	this.RamGB = ramGB
	return &this
}

// NewCreateVMTypeWithDefaults instantiates a new CreateVMType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMTypeWithDefaults() *CreateVMType {
	this := CreateVMType{}
	return &this
}

// GetName returns the Name field value
func (o *CreateVMType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateVMType) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateVMType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateVMType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateVMType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateVMType) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateVMType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateVMType) SetLabel(v string) {
	o.Label = &v
}

// GetCpuCores returns the CpuCores field value
func (o *CreateVMType) GetCpuCores() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetCpuCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuCores, true
}

// SetCpuCores sets field value
func (o *CreateVMType) SetCpuCores(v float32) {
	o.CpuCores = v
}

// GetRamGB returns the RamGB field value
func (o *CreateVMType) GetRamGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamGB
}

// GetRamGBOk returns a tuple with the RamGB field value
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetRamGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamGB, true
}

// SetRamGB sets field value
func (o *CreateVMType) SetRamGB(v float32) {
	o.RamGB = v
}

// GetIsExperimental returns the IsExperimental field value if set, zero value otherwise.
func (o *CreateVMType) GetIsExperimental() float32 {
	if o == nil || IsNil(o.IsExperimental) {
		var ret float32
		return ret
	}
	return *o.IsExperimental
}

// GetIsExperimentalOk returns a tuple with the IsExperimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetIsExperimentalOk() (*float32, bool) {
	if o == nil || IsNil(o.IsExperimental) {
		return nil, false
	}
	return o.IsExperimental, true
}

// HasIsExperimental returns a boolean if a field has been set.
func (o *CreateVMType) HasIsExperimental() bool {
	if o != nil && !IsNil(o.IsExperimental) {
		return true
	}

	return false
}

// SetIsExperimental gets a reference to the given float32 and assigns it to the IsExperimental field.
func (o *CreateVMType) SetIsExperimental(v float32) {
	o.IsExperimental = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateVMType) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateVMType) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateVMType) SetTags(v []string) {
	o.Tags = v
}

// GetForUnmanagedVMsOnly returns the ForUnmanagedVMsOnly field value if set, zero value otherwise.
func (o *CreateVMType) GetForUnmanagedVMsOnly() float32 {
	if o == nil || IsNil(o.ForUnmanagedVMsOnly) {
		var ret float32
		return ret
	}
	return *o.ForUnmanagedVMsOnly
}

// GetForUnmanagedVMsOnlyOk returns a tuple with the ForUnmanagedVMsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMType) GetForUnmanagedVMsOnlyOk() (*float32, bool) {
	if o == nil || IsNil(o.ForUnmanagedVMsOnly) {
		return nil, false
	}
	return o.ForUnmanagedVMsOnly, true
}

// HasForUnmanagedVMsOnly returns a boolean if a field has been set.
func (o *CreateVMType) HasForUnmanagedVMsOnly() bool {
	if o != nil && !IsNil(o.ForUnmanagedVMsOnly) {
		return true
	}

	return false
}

// SetForUnmanagedVMsOnly gets a reference to the given float32 and assigns it to the ForUnmanagedVMsOnly field.
func (o *CreateVMType) SetForUnmanagedVMsOnly(v float32) {
	o.ForUnmanagedVMsOnly = &v
}

func (o CreateVMType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["cpuCores"] = o.CpuCores
	toSerialize["ramGB"] = o.RamGB
	if !IsNil(o.IsExperimental) {
		toSerialize["isExperimental"] = o.IsExperimental
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ForUnmanagedVMsOnly) {
		toSerialize["forUnmanagedVMsOnly"] = o.ForUnmanagedVMsOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateVMType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"cpuCores",
		"ramGB",
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

	varCreateVMType := _CreateVMType{}

	err = json.Unmarshal(data, &varCreateVMType)

	if err != nil {
		return err
	}

	*o = CreateVMType(varCreateVMType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "label")
		delete(additionalProperties, "cpuCores")
		delete(additionalProperties, "ramGB")
		delete(additionalProperties, "isExperimental")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "forUnmanagedVMsOnly")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateVMType struct {
	value *CreateVMType
	isSet bool
}

func (v NullableCreateVMType) Get() *CreateVMType {
	return v.value
}

func (v *NullableCreateVMType) Set(val *CreateVMType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMType(val *CreateVMType) *NullableCreateVMType {
	return &NullableCreateVMType{value: val, isSet: true}
}

func (v NullableCreateVMType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



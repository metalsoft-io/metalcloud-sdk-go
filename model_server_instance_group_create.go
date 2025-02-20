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

// checks if the ServerInstanceGroupCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceGroupCreate{}

// ServerInstanceGroupCreate struct for ServerInstanceGroupCreate
type ServerInstanceGroupCreate struct {
	// The server instance group label. Will be automatically generated if not provided.
	Label *string `json:"label,omitempty"`
	ExtensionInstanceId *int32 `json:"extensionInstanceId,omitempty"`
	InstanceCount *int32 `json:"instanceCount,omitempty"`
	VolumeTemplateId *int32 `json:"volumeTemplateId,omitempty"`
	CustomVariables map[string]interface{} `json:"customVariables,omitempty"`
	DiskCount *int32 `json:"diskCount,omitempty"`
	DiskSizeMbytes *int32 `json:"diskSizeMbytes,omitempty"`
	DiskTypes []string `json:"diskTypes,omitempty"`
	// The group's default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied.
	DefaultServerProfileID *int32 `json:"defaultServerProfileID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceGroupCreate ServerInstanceGroupCreate

// NewServerInstanceGroupCreate instantiates a new ServerInstanceGroupCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceGroupCreate() *ServerInstanceGroupCreate {
	this := ServerInstanceGroupCreate{}
	var instanceCount int32 = 1
	this.InstanceCount = &instanceCount
	return &this
}

// NewServerInstanceGroupCreateWithDefaults instantiates a new ServerInstanceGroupCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceGroupCreateWithDefaults() *ServerInstanceGroupCreate {
	this := ServerInstanceGroupCreate{}
	var instanceCount int32 = 1
	this.InstanceCount = &instanceCount
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServerInstanceGroupCreate) SetLabel(v string) {
	o.Label = &v
}

// GetExtensionInstanceId returns the ExtensionInstanceId field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetExtensionInstanceId() int32 {
	if o == nil || IsNil(o.ExtensionInstanceId) {
		var ret int32
		return ret
	}
	return *o.ExtensionInstanceId
}

// GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetExtensionInstanceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ExtensionInstanceId) {
		return nil, false
	}
	return o.ExtensionInstanceId, true
}

// HasExtensionInstanceId returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasExtensionInstanceId() bool {
	if o != nil && !IsNil(o.ExtensionInstanceId) {
		return true
	}

	return false
}

// SetExtensionInstanceId gets a reference to the given int32 and assigns it to the ExtensionInstanceId field.
func (o *ServerInstanceGroupCreate) SetExtensionInstanceId(v int32) {
	o.ExtensionInstanceId = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetInstanceCount() int32 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetInstanceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *ServerInstanceGroupCreate) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetVolumeTemplateId returns the VolumeTemplateId field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetVolumeTemplateId() int32 {
	if o == nil || IsNil(o.VolumeTemplateId) {
		var ret int32
		return ret
	}
	return *o.VolumeTemplateId
}

// GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetVolumeTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VolumeTemplateId) {
		return nil, false
	}
	return o.VolumeTemplateId, true
}

// HasVolumeTemplateId returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasVolumeTemplateId() bool {
	if o != nil && !IsNil(o.VolumeTemplateId) {
		return true
	}

	return false
}

// SetVolumeTemplateId gets a reference to the given int32 and assigns it to the VolumeTemplateId field.
func (o *ServerInstanceGroupCreate) SetVolumeTemplateId(v int32) {
	o.VolumeTemplateId = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetCustomVariables() map[string]interface{} {
	if o == nil || IsNil(o.CustomVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetCustomVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return map[string]interface{}{}, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given map[string]interface{} and assigns it to the CustomVariables field.
func (o *ServerInstanceGroupCreate) SetCustomVariables(v map[string]interface{}) {
	o.CustomVariables = v
}

// GetDiskCount returns the DiskCount field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetDiskCount() int32 {
	if o == nil || IsNil(o.DiskCount) {
		var ret int32
		return ret
	}
	return *o.DiskCount
}

// GetDiskCountOk returns a tuple with the DiskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetDiskCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskCount) {
		return nil, false
	}
	return o.DiskCount, true
}

// HasDiskCount returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasDiskCount() bool {
	if o != nil && !IsNil(o.DiskCount) {
		return true
	}

	return false
}

// SetDiskCount gets a reference to the given int32 and assigns it to the DiskCount field.
func (o *ServerInstanceGroupCreate) SetDiskCount(v int32) {
	o.DiskCount = &v
}

// GetDiskSizeMbytes returns the DiskSizeMbytes field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetDiskSizeMbytes() int32 {
	if o == nil || IsNil(o.DiskSizeMbytes) {
		var ret int32
		return ret
	}
	return *o.DiskSizeMbytes
}

// GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetDiskSizeMbytesOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskSizeMbytes) {
		return nil, false
	}
	return o.DiskSizeMbytes, true
}

// HasDiskSizeMbytes returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasDiskSizeMbytes() bool {
	if o != nil && !IsNil(o.DiskSizeMbytes) {
		return true
	}

	return false
}

// SetDiskSizeMbytes gets a reference to the given int32 and assigns it to the DiskSizeMbytes field.
func (o *ServerInstanceGroupCreate) SetDiskSizeMbytes(v int32) {
	o.DiskSizeMbytes = &v
}

// GetDiskTypes returns the DiskTypes field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetDiskTypes() []string {
	if o == nil || IsNil(o.DiskTypes) {
		var ret []string
		return ret
	}
	return o.DiskTypes
}

// GetDiskTypesOk returns a tuple with the DiskTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetDiskTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DiskTypes) {
		return nil, false
	}
	return o.DiskTypes, true
}

// HasDiskTypes returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasDiskTypes() bool {
	if o != nil && !IsNil(o.DiskTypes) {
		return true
	}

	return false
}

// SetDiskTypes gets a reference to the given []string and assigns it to the DiskTypes field.
func (o *ServerInstanceGroupCreate) SetDiskTypes(v []string) {
	o.DiskTypes = v
}

// GetDefaultServerProfileID returns the DefaultServerProfileID field value if set, zero value otherwise.
func (o *ServerInstanceGroupCreate) GetDefaultServerProfileID() int32 {
	if o == nil || IsNil(o.DefaultServerProfileID) {
		var ret int32
		return ret
	}
	return *o.DefaultServerProfileID
}

// GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupCreate) GetDefaultServerProfileIDOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultServerProfileID) {
		return nil, false
	}
	return o.DefaultServerProfileID, true
}

// HasDefaultServerProfileID returns a boolean if a field has been set.
func (o *ServerInstanceGroupCreate) HasDefaultServerProfileID() bool {
	if o != nil && !IsNil(o.DefaultServerProfileID) {
		return true
	}

	return false
}

// SetDefaultServerProfileID gets a reference to the given int32 and assigns it to the DefaultServerProfileID field.
func (o *ServerInstanceGroupCreate) SetDefaultServerProfileID(v int32) {
	o.DefaultServerProfileID = &v
}

func (o ServerInstanceGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceGroupCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ExtensionInstanceId) {
		toSerialize["extensionInstanceId"] = o.ExtensionInstanceId
	}
	if !IsNil(o.InstanceCount) {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if !IsNil(o.VolumeTemplateId) {
		toSerialize["volumeTemplateId"] = o.VolumeTemplateId
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	if !IsNil(o.DiskCount) {
		toSerialize["diskCount"] = o.DiskCount
	}
	if !IsNil(o.DiskSizeMbytes) {
		toSerialize["diskSizeMbytes"] = o.DiskSizeMbytes
	}
	if !IsNil(o.DiskTypes) {
		toSerialize["diskTypes"] = o.DiskTypes
	}
	if !IsNil(o.DefaultServerProfileID) {
		toSerialize["defaultServerProfileID"] = o.DefaultServerProfileID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceGroupCreate) UnmarshalJSON(data []byte) (err error) {
	varServerInstanceGroupCreate := _ServerInstanceGroupCreate{}

	err = json.Unmarshal(data, &varServerInstanceGroupCreate)

	if err != nil {
		return err
	}

	*o = ServerInstanceGroupCreate(varServerInstanceGroupCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "extensionInstanceId")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "volumeTemplateId")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "diskCount")
		delete(additionalProperties, "diskSizeMbytes")
		delete(additionalProperties, "diskTypes")
		delete(additionalProperties, "defaultServerProfileID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceGroupCreate struct {
	value *ServerInstanceGroupCreate
	isSet bool
}

func (v NullableServerInstanceGroupCreate) Get() *ServerInstanceGroupCreate {
	return v.value
}

func (v *NullableServerInstanceGroupCreate) Set(val *ServerInstanceGroupCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceGroupCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceGroupCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceGroupCreate(val *ServerInstanceGroupCreate) *NullableServerInstanceGroupCreate {
	return &NullableServerInstanceGroupCreate{value: val, isSet: true}
}

func (v NullableServerInstanceGroupCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceGroupCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



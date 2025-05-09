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

// checks if the UpdateStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorage{}

// UpdateStorage struct for UpdateStorage
type UpdateStorage struct {
	// Specifies if the storage is in maintenance
	InMaintenance *float32 `json:"inMaintenance,omitempty"`
	// Specifies if the storage is experimental
	IsExperimental *float32 `json:"isExperimental,omitempty"`
	// Specifies the drive priority
	DrivePriority *float32 `json:"drivePriority,omitempty"`
	// Specifies the shared drive priority
	SharedDrivePriority *float32 `json:"sharedDrivePriority,omitempty"`
	// Tags
	Tags []string `json:"tags,omitempty"`
	// Options for the storage
	Options *UpdateStorageOptions `json:"options,omitempty"`
	// The password to use.
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateStorage UpdateStorage

// NewUpdateStorage instantiates a new UpdateStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorage() *UpdateStorage {
	this := UpdateStorage{}
	return &this
}

// NewUpdateStorageWithDefaults instantiates a new UpdateStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageWithDefaults() *UpdateStorage {
	this := UpdateStorage{}
	return &this
}

// GetInMaintenance returns the InMaintenance field value if set, zero value otherwise.
func (o *UpdateStorage) GetInMaintenance() float32 {
	if o == nil || IsNil(o.InMaintenance) {
		var ret float32
		return ret
	}
	return *o.InMaintenance
}

// GetInMaintenanceOk returns a tuple with the InMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetInMaintenanceOk() (*float32, bool) {
	if o == nil || IsNil(o.InMaintenance) {
		return nil, false
	}
	return o.InMaintenance, true
}

// HasInMaintenance returns a boolean if a field has been set.
func (o *UpdateStorage) HasInMaintenance() bool {
	if o != nil && !IsNil(o.InMaintenance) {
		return true
	}

	return false
}

// SetInMaintenance gets a reference to the given float32 and assigns it to the InMaintenance field.
func (o *UpdateStorage) SetInMaintenance(v float32) {
	o.InMaintenance = &v
}

// GetIsExperimental returns the IsExperimental field value if set, zero value otherwise.
func (o *UpdateStorage) GetIsExperimental() float32 {
	if o == nil || IsNil(o.IsExperimental) {
		var ret float32
		return ret
	}
	return *o.IsExperimental
}

// GetIsExperimentalOk returns a tuple with the IsExperimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetIsExperimentalOk() (*float32, bool) {
	if o == nil || IsNil(o.IsExperimental) {
		return nil, false
	}
	return o.IsExperimental, true
}

// HasIsExperimental returns a boolean if a field has been set.
func (o *UpdateStorage) HasIsExperimental() bool {
	if o != nil && !IsNil(o.IsExperimental) {
		return true
	}

	return false
}

// SetIsExperimental gets a reference to the given float32 and assigns it to the IsExperimental field.
func (o *UpdateStorage) SetIsExperimental(v float32) {
	o.IsExperimental = &v
}

// GetDrivePriority returns the DrivePriority field value if set, zero value otherwise.
func (o *UpdateStorage) GetDrivePriority() float32 {
	if o == nil || IsNil(o.DrivePriority) {
		var ret float32
		return ret
	}
	return *o.DrivePriority
}

// GetDrivePriorityOk returns a tuple with the DrivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetDrivePriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.DrivePriority) {
		return nil, false
	}
	return o.DrivePriority, true
}

// HasDrivePriority returns a boolean if a field has been set.
func (o *UpdateStorage) HasDrivePriority() bool {
	if o != nil && !IsNil(o.DrivePriority) {
		return true
	}

	return false
}

// SetDrivePriority gets a reference to the given float32 and assigns it to the DrivePriority field.
func (o *UpdateStorage) SetDrivePriority(v float32) {
	o.DrivePriority = &v
}

// GetSharedDrivePriority returns the SharedDrivePriority field value if set, zero value otherwise.
func (o *UpdateStorage) GetSharedDrivePriority() float32 {
	if o == nil || IsNil(o.SharedDrivePriority) {
		var ret float32
		return ret
	}
	return *o.SharedDrivePriority
}

// GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetSharedDrivePriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.SharedDrivePriority) {
		return nil, false
	}
	return o.SharedDrivePriority, true
}

// HasSharedDrivePriority returns a boolean if a field has been set.
func (o *UpdateStorage) HasSharedDrivePriority() bool {
	if o != nil && !IsNil(o.SharedDrivePriority) {
		return true
	}

	return false
}

// SetSharedDrivePriority gets a reference to the given float32 and assigns it to the SharedDrivePriority field.
func (o *UpdateStorage) SetSharedDrivePriority(v float32) {
	o.SharedDrivePriority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateStorage) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateStorage) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateStorage) SetTags(v []string) {
	o.Tags = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *UpdateStorage) GetOptions() UpdateStorageOptions {
	if o == nil || IsNil(o.Options) {
		var ret UpdateStorageOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetOptionsOk() (*UpdateStorageOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *UpdateStorage) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given UpdateStorageOptions and assigns it to the Options field.
func (o *UpdateStorage) SetOptions(v UpdateStorageOptions) {
	o.Options = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateStorage) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorage) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateStorage) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateStorage) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InMaintenance) {
		toSerialize["inMaintenance"] = o.InMaintenance
	}
	if !IsNil(o.IsExperimental) {
		toSerialize["isExperimental"] = o.IsExperimental
	}
	if !IsNil(o.DrivePriority) {
		toSerialize["drivePriority"] = o.DrivePriority
	}
	if !IsNil(o.SharedDrivePriority) {
		toSerialize["sharedDrivePriority"] = o.SharedDrivePriority
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateStorage) UnmarshalJSON(data []byte) (err error) {
	varUpdateStorage := _UpdateStorage{}

	err = json.Unmarshal(data, &varUpdateStorage)

	if err != nil {
		return err
	}

	*o = UpdateStorage(varUpdateStorage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "inMaintenance")
		delete(additionalProperties, "isExperimental")
		delete(additionalProperties, "drivePriority")
		delete(additionalProperties, "sharedDrivePriority")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "options")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateStorage struct {
	value *UpdateStorage
	isSet bool
}

func (v NullableUpdateStorage) Get() *UpdateStorage {
	return v.value
}

func (v *NullableUpdateStorage) Set(val *UpdateStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStorage(val *UpdateStorage) *NullableUpdateStorage {
	return &NullableUpdateStorage{value: val, isSet: true}
}

func (v NullableUpdateStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



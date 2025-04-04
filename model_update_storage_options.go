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

// checks if the UpdateStorageOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStorageOptions{}

// UpdateStorageOptions struct for UpdateStorageOptions
type UpdateStorageOptions struct {
	// Enable data reduction
	EnableDataReduction *float32 `json:"enableDataReduction,omitempty"`
	// Enable advanced deduplication
	EnableAdvancedDeduplication *float32 `json:"enableAdvancedDeduplication,omitempty"`
	// Volume name
	VolumeName *string `json:"volumeName,omitempty"`
	// Array id to use (for certain storage drivers)
	ArrayId *string `json:"arrayId,omitempty"`
	// Director id to use (for certain storage drivers)
	DirectorId *string `json:"directorId,omitempty"`
	// S3 Hostname to use (for certain storage drivers)
	S3Hostname *string `json:"s3Hostname,omitempty"`
	// Enable advanced deduplication
	S3Port *float32 `json:"s3Port,omitempty"`
	// Fibre channel enabled
	FibreChannelEnabled *float32 `json:"fibreChannelEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateStorageOptions UpdateStorageOptions

// NewUpdateStorageOptions instantiates a new UpdateStorageOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStorageOptions() *UpdateStorageOptions {
	this := UpdateStorageOptions{}
	return &this
}

// NewUpdateStorageOptionsWithDefaults instantiates a new UpdateStorageOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStorageOptionsWithDefaults() *UpdateStorageOptions {
	this := UpdateStorageOptions{}
	return &this
}

// GetEnableDataReduction returns the EnableDataReduction field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetEnableDataReduction() float32 {
	if o == nil || IsNil(o.EnableDataReduction) {
		var ret float32
		return ret
	}
	return *o.EnableDataReduction
}

// GetEnableDataReductionOk returns a tuple with the EnableDataReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetEnableDataReductionOk() (*float32, bool) {
	if o == nil || IsNil(o.EnableDataReduction) {
		return nil, false
	}
	return o.EnableDataReduction, true
}

// HasEnableDataReduction returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasEnableDataReduction() bool {
	if o != nil && !IsNil(o.EnableDataReduction) {
		return true
	}

	return false
}

// SetEnableDataReduction gets a reference to the given float32 and assigns it to the EnableDataReduction field.
func (o *UpdateStorageOptions) SetEnableDataReduction(v float32) {
	o.EnableDataReduction = &v
}

// GetEnableAdvancedDeduplication returns the EnableAdvancedDeduplication field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetEnableAdvancedDeduplication() float32 {
	if o == nil || IsNil(o.EnableAdvancedDeduplication) {
		var ret float32
		return ret
	}
	return *o.EnableAdvancedDeduplication
}

// GetEnableAdvancedDeduplicationOk returns a tuple with the EnableAdvancedDeduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetEnableAdvancedDeduplicationOk() (*float32, bool) {
	if o == nil || IsNil(o.EnableAdvancedDeduplication) {
		return nil, false
	}
	return o.EnableAdvancedDeduplication, true
}

// HasEnableAdvancedDeduplication returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasEnableAdvancedDeduplication() bool {
	if o != nil && !IsNil(o.EnableAdvancedDeduplication) {
		return true
	}

	return false
}

// SetEnableAdvancedDeduplication gets a reference to the given float32 and assigns it to the EnableAdvancedDeduplication field.
func (o *UpdateStorageOptions) SetEnableAdvancedDeduplication(v float32) {
	o.EnableAdvancedDeduplication = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *UpdateStorageOptions) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetArrayId returns the ArrayId field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetArrayId() string {
	if o == nil || IsNil(o.ArrayId) {
		var ret string
		return ret
	}
	return *o.ArrayId
}

// GetArrayIdOk returns a tuple with the ArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetArrayIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArrayId) {
		return nil, false
	}
	return o.ArrayId, true
}

// HasArrayId returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasArrayId() bool {
	if o != nil && !IsNil(o.ArrayId) {
		return true
	}

	return false
}

// SetArrayId gets a reference to the given string and assigns it to the ArrayId field.
func (o *UpdateStorageOptions) SetArrayId(v string) {
	o.ArrayId = &v
}

// GetDirectorId returns the DirectorId field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetDirectorId() string {
	if o == nil || IsNil(o.DirectorId) {
		var ret string
		return ret
	}
	return *o.DirectorId
}

// GetDirectorIdOk returns a tuple with the DirectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetDirectorIdOk() (*string, bool) {
	if o == nil || IsNil(o.DirectorId) {
		return nil, false
	}
	return o.DirectorId, true
}

// HasDirectorId returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasDirectorId() bool {
	if o != nil && !IsNil(o.DirectorId) {
		return true
	}

	return false
}

// SetDirectorId gets a reference to the given string and assigns it to the DirectorId field.
func (o *UpdateStorageOptions) SetDirectorId(v string) {
	o.DirectorId = &v
}

// GetS3Hostname returns the S3Hostname field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetS3Hostname() string {
	if o == nil || IsNil(o.S3Hostname) {
		var ret string
		return ret
	}
	return *o.S3Hostname
}

// GetS3HostnameOk returns a tuple with the S3Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetS3HostnameOk() (*string, bool) {
	if o == nil || IsNil(o.S3Hostname) {
		return nil, false
	}
	return o.S3Hostname, true
}

// HasS3Hostname returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasS3Hostname() bool {
	if o != nil && !IsNil(o.S3Hostname) {
		return true
	}

	return false
}

// SetS3Hostname gets a reference to the given string and assigns it to the S3Hostname field.
func (o *UpdateStorageOptions) SetS3Hostname(v string) {
	o.S3Hostname = &v
}

// GetS3Port returns the S3Port field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetS3Port() float32 {
	if o == nil || IsNil(o.S3Port) {
		var ret float32
		return ret
	}
	return *o.S3Port
}

// GetS3PortOk returns a tuple with the S3Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetS3PortOk() (*float32, bool) {
	if o == nil || IsNil(o.S3Port) {
		return nil, false
	}
	return o.S3Port, true
}

// HasS3Port returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasS3Port() bool {
	if o != nil && !IsNil(o.S3Port) {
		return true
	}

	return false
}

// SetS3Port gets a reference to the given float32 and assigns it to the S3Port field.
func (o *UpdateStorageOptions) SetS3Port(v float32) {
	o.S3Port = &v
}

// GetFibreChannelEnabled returns the FibreChannelEnabled field value if set, zero value otherwise.
func (o *UpdateStorageOptions) GetFibreChannelEnabled() float32 {
	if o == nil || IsNil(o.FibreChannelEnabled) {
		var ret float32
		return ret
	}
	return *o.FibreChannelEnabled
}

// GetFibreChannelEnabledOk returns a tuple with the FibreChannelEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStorageOptions) GetFibreChannelEnabledOk() (*float32, bool) {
	if o == nil || IsNil(o.FibreChannelEnabled) {
		return nil, false
	}
	return o.FibreChannelEnabled, true
}

// HasFibreChannelEnabled returns a boolean if a field has been set.
func (o *UpdateStorageOptions) HasFibreChannelEnabled() bool {
	if o != nil && !IsNil(o.FibreChannelEnabled) {
		return true
	}

	return false
}

// SetFibreChannelEnabled gets a reference to the given float32 and assigns it to the FibreChannelEnabled field.
func (o *UpdateStorageOptions) SetFibreChannelEnabled(v float32) {
	o.FibreChannelEnabled = &v
}

func (o UpdateStorageOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStorageOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableDataReduction) {
		toSerialize["enableDataReduction"] = o.EnableDataReduction
	}
	if !IsNil(o.EnableAdvancedDeduplication) {
		toSerialize["enableAdvancedDeduplication"] = o.EnableAdvancedDeduplication
	}
	if !IsNil(o.VolumeName) {
		toSerialize["volumeName"] = o.VolumeName
	}
	if !IsNil(o.ArrayId) {
		toSerialize["arrayId"] = o.ArrayId
	}
	if !IsNil(o.DirectorId) {
		toSerialize["directorId"] = o.DirectorId
	}
	if !IsNil(o.S3Hostname) {
		toSerialize["s3Hostname"] = o.S3Hostname
	}
	if !IsNil(o.S3Port) {
		toSerialize["s3Port"] = o.S3Port
	}
	if !IsNil(o.FibreChannelEnabled) {
		toSerialize["fibreChannelEnabled"] = o.FibreChannelEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateStorageOptions) UnmarshalJSON(data []byte) (err error) {
	varUpdateStorageOptions := _UpdateStorageOptions{}

	err = json.Unmarshal(data, &varUpdateStorageOptions)

	if err != nil {
		return err
	}

	*o = UpdateStorageOptions(varUpdateStorageOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enableDataReduction")
		delete(additionalProperties, "enableAdvancedDeduplication")
		delete(additionalProperties, "volumeName")
		delete(additionalProperties, "arrayId")
		delete(additionalProperties, "directorId")
		delete(additionalProperties, "s3Hostname")
		delete(additionalProperties, "s3Port")
		delete(additionalProperties, "fibreChannelEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateStorageOptions struct {
	value *UpdateStorageOptions
	isSet bool
}

func (v NullableUpdateStorageOptions) Get() *UpdateStorageOptions {
	return v.value
}

func (v *NullableUpdateStorageOptions) Set(val *UpdateStorageOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStorageOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStorageOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStorageOptions(val *UpdateStorageOptions) *NullableUpdateStorageOptions {
	return &NullableUpdateStorageOptions{value: val, isSet: true}
}

func (v NullableUpdateStorageOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStorageOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



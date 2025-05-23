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

// checks if the UpdateFileShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFileShare{}

// UpdateFileShare struct for UpdateFileShare
type UpdateFileShare struct {
	// Disk size in GB for File Share
	SizeGB *float32 `json:"sizeGB,omitempty"`
	// Label of the File Share.
	Label *string `json:"label,omitempty"`
	// Id of the Logical Network for the File Share.
	LogicalNetworkId *float32 `json:"logicalNetworkId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFileShare UpdateFileShare

// NewUpdateFileShare instantiates a new UpdateFileShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFileShare() *UpdateFileShare {
	this := UpdateFileShare{}
	return &this
}

// NewUpdateFileShareWithDefaults instantiates a new UpdateFileShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileShareWithDefaults() *UpdateFileShare {
	this := UpdateFileShare{}
	return &this
}

// GetSizeGB returns the SizeGB field value if set, zero value otherwise.
func (o *UpdateFileShare) GetSizeGB() float32 {
	if o == nil || IsNil(o.SizeGB) {
		var ret float32
		return ret
	}
	return *o.SizeGB
}

// GetSizeGBOk returns a tuple with the SizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileShare) GetSizeGBOk() (*float32, bool) {
	if o == nil || IsNil(o.SizeGB) {
		return nil, false
	}
	return o.SizeGB, true
}

// HasSizeGB returns a boolean if a field has been set.
func (o *UpdateFileShare) HasSizeGB() bool {
	if o != nil && !IsNil(o.SizeGB) {
		return true
	}

	return false
}

// SetSizeGB gets a reference to the given float32 and assigns it to the SizeGB field.
func (o *UpdateFileShare) SetSizeGB(v float32) {
	o.SizeGB = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UpdateFileShare) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileShare) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UpdateFileShare) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UpdateFileShare) SetLabel(v string) {
	o.Label = &v
}

// GetLogicalNetworkId returns the LogicalNetworkId field value if set, zero value otherwise.
func (o *UpdateFileShare) GetLogicalNetworkId() float32 {
	if o == nil || IsNil(o.LogicalNetworkId) {
		var ret float32
		return ret
	}
	return *o.LogicalNetworkId
}

// GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFileShare) GetLogicalNetworkIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LogicalNetworkId) {
		return nil, false
	}
	return o.LogicalNetworkId, true
}

// HasLogicalNetworkId returns a boolean if a field has been set.
func (o *UpdateFileShare) HasLogicalNetworkId() bool {
	if o != nil && !IsNil(o.LogicalNetworkId) {
		return true
	}

	return false
}

// SetLogicalNetworkId gets a reference to the given float32 and assigns it to the LogicalNetworkId field.
func (o *UpdateFileShare) SetLogicalNetworkId(v float32) {
	o.LogicalNetworkId = &v
}

func (o UpdateFileShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFileShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SizeGB) {
		toSerialize["sizeGB"] = o.SizeGB
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.LogicalNetworkId) {
		toSerialize["logicalNetworkId"] = o.LogicalNetworkId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFileShare) UnmarshalJSON(data []byte) (err error) {
	varUpdateFileShare := _UpdateFileShare{}

	err = json.Unmarshal(data, &varUpdateFileShare)

	if err != nil {
		return err
	}

	*o = UpdateFileShare(varUpdateFileShare)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sizeGB")
		delete(additionalProperties, "label")
		delete(additionalProperties, "logicalNetworkId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFileShare struct {
	value *UpdateFileShare
	isSet bool
}

func (v NullableUpdateFileShare) Get() *UpdateFileShare {
	return v.value
}

func (v *NullableUpdateFileShare) Set(val *UpdateFileShare) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFileShare) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFileShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFileShare(val *UpdateFileShare) *NullableUpdateFileShare {
	return &NullableUpdateFileShare{value: val, isSet: true}
}

func (v NullableUpdateFileShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFileShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



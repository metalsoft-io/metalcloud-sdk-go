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

// checks if the UpdateBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBucket{}

// UpdateBucket struct for UpdateBucket
type UpdateBucket struct {
	// Disk size in GB for Bucket
	SizeGB *float32 `json:"sizeGB,omitempty"`
	// Label of the Bucket.
	Label *string `json:"label,omitempty"`
	// Id of the Logical Network for the Bucket.
	LogicalNetworkId *float32 `json:"logicalNetworkId,omitempty"`
	GuiSettings *GenericGUISettings `json:"guiSettings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateBucket UpdateBucket

// NewUpdateBucket instantiates a new UpdateBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBucket() *UpdateBucket {
	this := UpdateBucket{}
	return &this
}

// NewUpdateBucketWithDefaults instantiates a new UpdateBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBucketWithDefaults() *UpdateBucket {
	this := UpdateBucket{}
	return &this
}

// GetSizeGB returns the SizeGB field value if set, zero value otherwise.
func (o *UpdateBucket) GetSizeGB() float32 {
	if o == nil || IsNil(o.SizeGB) {
		var ret float32
		return ret
	}
	return *o.SizeGB
}

// GetSizeGBOk returns a tuple with the SizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBucket) GetSizeGBOk() (*float32, bool) {
	if o == nil || IsNil(o.SizeGB) {
		return nil, false
	}
	return o.SizeGB, true
}

// HasSizeGB returns a boolean if a field has been set.
func (o *UpdateBucket) HasSizeGB() bool {
	if o != nil && !IsNil(o.SizeGB) {
		return true
	}

	return false
}

// SetSizeGB gets a reference to the given float32 and assigns it to the SizeGB field.
func (o *UpdateBucket) SetSizeGB(v float32) {
	o.SizeGB = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UpdateBucket) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBucket) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UpdateBucket) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UpdateBucket) SetLabel(v string) {
	o.Label = &v
}

// GetLogicalNetworkId returns the LogicalNetworkId field value if set, zero value otherwise.
func (o *UpdateBucket) GetLogicalNetworkId() float32 {
	if o == nil || IsNil(o.LogicalNetworkId) {
		var ret float32
		return ret
	}
	return *o.LogicalNetworkId
}

// GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBucket) GetLogicalNetworkIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LogicalNetworkId) {
		return nil, false
	}
	return o.LogicalNetworkId, true
}

// HasLogicalNetworkId returns a boolean if a field has been set.
func (o *UpdateBucket) HasLogicalNetworkId() bool {
	if o != nil && !IsNil(o.LogicalNetworkId) {
		return true
	}

	return false
}

// SetLogicalNetworkId gets a reference to the given float32 and assigns it to the LogicalNetworkId field.
func (o *UpdateBucket) SetLogicalNetworkId(v float32) {
	o.LogicalNetworkId = &v
}

// GetGuiSettings returns the GuiSettings field value if set, zero value otherwise.
func (o *UpdateBucket) GetGuiSettings() GenericGUISettings {
	if o == nil || IsNil(o.GuiSettings) {
		var ret GenericGUISettings
		return ret
	}
	return *o.GuiSettings
}

// GetGuiSettingsOk returns a tuple with the GuiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBucket) GetGuiSettingsOk() (*GenericGUISettings, bool) {
	if o == nil || IsNil(o.GuiSettings) {
		return nil, false
	}
	return o.GuiSettings, true
}

// HasGuiSettings returns a boolean if a field has been set.
func (o *UpdateBucket) HasGuiSettings() bool {
	if o != nil && !IsNil(o.GuiSettings) {
		return true
	}

	return false
}

// SetGuiSettings gets a reference to the given GenericGUISettings and assigns it to the GuiSettings field.
func (o *UpdateBucket) SetGuiSettings(v GenericGUISettings) {
	o.GuiSettings = &v
}

func (o UpdateBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBucket) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GuiSettings) {
		toSerialize["guiSettings"] = o.GuiSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateBucket) UnmarshalJSON(data []byte) (err error) {
	varUpdateBucket := _UpdateBucket{}

	err = json.Unmarshal(data, &varUpdateBucket)

	if err != nil {
		return err
	}

	*o = UpdateBucket(varUpdateBucket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sizeGB")
		delete(additionalProperties, "label")
		delete(additionalProperties, "logicalNetworkId")
		delete(additionalProperties, "guiSettings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateBucket struct {
	value *UpdateBucket
	isSet bool
}

func (v NullableUpdateBucket) Get() *UpdateBucket {
	return v.value
}

func (v *NullableUpdateBucket) Set(val *UpdateBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBucket(val *UpdateBucket) *NullableUpdateBucket {
	return &NullableUpdateBucket{value: val, isSet: true}
}

func (v NullableUpdateBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



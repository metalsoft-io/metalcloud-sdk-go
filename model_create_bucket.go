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

// checks if the CreateBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBucket{}

// CreateBucket struct for CreateBucket
type CreateBucket struct {
	// Disk size in GB for Bucket
	SizeGB float32 `json:"sizeGB"`
	// Id of the Logical Network for the Bucket.
	LogicalNetworkId *float32 `json:"logicalNetworkId,omitempty"`
	// Label of the Bucket.
	Label *string `json:"label,omitempty"`
	Meta *BucketMeta `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateBucket CreateBucket

// NewCreateBucket instantiates a new CreateBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBucket(sizeGB float32) *CreateBucket {
	this := CreateBucket{}
	this.SizeGB = sizeGB
	return &this
}

// NewCreateBucketWithDefaults instantiates a new CreateBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBucketWithDefaults() *CreateBucket {
	this := CreateBucket{}
	return &this
}

// GetSizeGB returns the SizeGB field value
func (o *CreateBucket) GetSizeGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SizeGB
}

// GetSizeGBOk returns a tuple with the SizeGB field value
// and a boolean to check if the value has been set.
func (o *CreateBucket) GetSizeGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeGB, true
}

// SetSizeGB sets field value
func (o *CreateBucket) SetSizeGB(v float32) {
	o.SizeGB = v
}

// GetLogicalNetworkId returns the LogicalNetworkId field value if set, zero value otherwise.
func (o *CreateBucket) GetLogicalNetworkId() float32 {
	if o == nil || IsNil(o.LogicalNetworkId) {
		var ret float32
		return ret
	}
	return *o.LogicalNetworkId
}

// GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBucket) GetLogicalNetworkIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LogicalNetworkId) {
		return nil, false
	}
	return o.LogicalNetworkId, true
}

// HasLogicalNetworkId returns a boolean if a field has been set.
func (o *CreateBucket) HasLogicalNetworkId() bool {
	if o != nil && !IsNil(o.LogicalNetworkId) {
		return true
	}

	return false
}

// SetLogicalNetworkId gets a reference to the given float32 and assigns it to the LogicalNetworkId field.
func (o *CreateBucket) SetLogicalNetworkId(v float32) {
	o.LogicalNetworkId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateBucket) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBucket) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateBucket) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateBucket) SetLabel(v string) {
	o.Label = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CreateBucket) GetMeta() BucketMeta {
	if o == nil || IsNil(o.Meta) {
		var ret BucketMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBucket) GetMetaOk() (*BucketMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CreateBucket) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given BucketMeta and assigns it to the Meta field.
func (o *CreateBucket) SetMeta(v BucketMeta) {
	o.Meta = &v
}

func (o CreateBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sizeGB"] = o.SizeGB
	if !IsNil(o.LogicalNetworkId) {
		toSerialize["logicalNetworkId"] = o.LogicalNetworkId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sizeGB",
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

	varCreateBucket := _CreateBucket{}

	err = json.Unmarshal(data, &varCreateBucket)

	if err != nil {
		return err
	}

	*o = CreateBucket(varCreateBucket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sizeGB")
		delete(additionalProperties, "logicalNetworkId")
		delete(additionalProperties, "label")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateBucket struct {
	value *CreateBucket
	isSet bool
}

func (v NullableCreateBucket) Get() *CreateBucket {
	return v.value
}

func (v *NullableCreateBucket) Set(val *CreateBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBucket(val *CreateBucket) *NullableCreateBucket {
	return &NullableCreateBucket{value: val, isSet: true}
}

func (v NullableCreateBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



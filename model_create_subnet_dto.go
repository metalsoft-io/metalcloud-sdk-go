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

// checks if the CreateSubnetDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubnetDto{}

// CreateSubnetDto struct for CreateSubnetDto
type CreateSubnetDto struct {
	Label *string `json:"label,omitempty"`
	// Name of the Subnet
	Name string `json:"name"`
	// ID of the parent subnet
	ParentSubnetId float32 `json:"parentSubnetId"`
	NetworkAddress string `json:"networkAddress"`
	PrefixLength float32 `json:"prefixLength"`
	IsPool bool `json:"isPool"`
	VrfId float32 `json:"vrfId"`
	AllocationDenylist []string `json:"allocationDenylist"`
	AllowedChildOverlapConditions []string `json:"allowedChildOverlapConditions"`
	Tags map[string]interface{} `json:"tags"`
	Metadata map[string]interface{} `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _CreateSubnetDto CreateSubnetDto

// NewCreateSubnetDto instantiates a new CreateSubnetDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubnetDto(name string, parentSubnetId float32, networkAddress string, prefixLength float32, isPool bool, vrfId float32, allocationDenylist []string, allowedChildOverlapConditions []string, tags map[string]interface{}, metadata map[string]interface{}) *CreateSubnetDto {
	this := CreateSubnetDto{}
	this.Name = name
	this.ParentSubnetId = parentSubnetId
	this.NetworkAddress = networkAddress
	this.PrefixLength = prefixLength
	this.IsPool = isPool
	this.VrfId = vrfId
	this.AllocationDenylist = allocationDenylist
	this.AllowedChildOverlapConditions = allowedChildOverlapConditions
	this.Tags = tags
	this.Metadata = metadata
	return &this
}

// NewCreateSubnetDtoWithDefaults instantiates a new CreateSubnetDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubnetDtoWithDefaults() *CreateSubnetDto {
	this := CreateSubnetDto{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateSubnetDto) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateSubnetDto) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateSubnetDto) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value
func (o *CreateSubnetDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSubnetDto) SetName(v string) {
	o.Name = v
}

// GetParentSubnetId returns the ParentSubnetId field value
func (o *CreateSubnetDto) GetParentSubnetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ParentSubnetId
}

// GetParentSubnetIdOk returns a tuple with the ParentSubnetId field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetParentSubnetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentSubnetId, true
}

// SetParentSubnetId sets field value
func (o *CreateSubnetDto) SetParentSubnetId(v float32) {
	o.ParentSubnetId = v
}

// GetNetworkAddress returns the NetworkAddress field value
func (o *CreateSubnetDto) GetNetworkAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkAddress
}

// GetNetworkAddressOk returns a tuple with the NetworkAddress field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetNetworkAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkAddress, true
}

// SetNetworkAddress sets field value
func (o *CreateSubnetDto) SetNetworkAddress(v string) {
	o.NetworkAddress = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *CreateSubnetDto) GetPrefixLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetPrefixLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *CreateSubnetDto) SetPrefixLength(v float32) {
	o.PrefixLength = v
}

// GetIsPool returns the IsPool field value
func (o *CreateSubnetDto) GetIsPool() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPool
}

// GetIsPoolOk returns a tuple with the IsPool field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetIsPoolOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPool, true
}

// SetIsPool sets field value
func (o *CreateSubnetDto) SetIsPool(v bool) {
	o.IsPool = v
}

// GetVrfId returns the VrfId field value
func (o *CreateSubnetDto) GetVrfId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VrfId
}

// GetVrfIdOk returns a tuple with the VrfId field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetVrfIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrfId, true
}

// SetVrfId sets field value
func (o *CreateSubnetDto) SetVrfId(v float32) {
	o.VrfId = v
}

// GetAllocationDenylist returns the AllocationDenylist field value
func (o *CreateSubnetDto) GetAllocationDenylist() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllocationDenylist
}

// GetAllocationDenylistOk returns a tuple with the AllocationDenylist field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetAllocationDenylistOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocationDenylist, true
}

// SetAllocationDenylist sets field value
func (o *CreateSubnetDto) SetAllocationDenylist(v []string) {
	o.AllocationDenylist = v
}

// GetAllowedChildOverlapConditions returns the AllowedChildOverlapConditions field value
func (o *CreateSubnetDto) GetAllowedChildOverlapConditions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedChildOverlapConditions
}

// GetAllowedChildOverlapConditionsOk returns a tuple with the AllowedChildOverlapConditions field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetAllowedChildOverlapConditionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedChildOverlapConditions, true
}

// SetAllowedChildOverlapConditions sets field value
func (o *CreateSubnetDto) SetAllowedChildOverlapConditions(v []string) {
	o.AllowedChildOverlapConditions = v
}

// GetTags returns the Tags field value
func (o *CreateSubnetDto) GetTags() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CreateSubnetDto) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetMetadata returns the Metadata field value
func (o *CreateSubnetDto) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CreateSubnetDto) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *CreateSubnetDto) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CreateSubnetDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubnetDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["name"] = o.Name
	toSerialize["parentSubnetId"] = o.ParentSubnetId
	toSerialize["networkAddress"] = o.NetworkAddress
	toSerialize["prefixLength"] = o.PrefixLength
	toSerialize["isPool"] = o.IsPool
	toSerialize["vrfId"] = o.VrfId
	toSerialize["allocationDenylist"] = o.AllocationDenylist
	toSerialize["allowedChildOverlapConditions"] = o.AllowedChildOverlapConditions
	toSerialize["tags"] = o.Tags
	toSerialize["metadata"] = o.Metadata

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSubnetDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"parentSubnetId",
		"networkAddress",
		"prefixLength",
		"isPool",
		"vrfId",
		"allocationDenylist",
		"allowedChildOverlapConditions",
		"tags",
		"metadata",
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

	varCreateSubnetDto := _CreateSubnetDto{}

	err = json.Unmarshal(data, &varCreateSubnetDto)

	if err != nil {
		return err
	}

	*o = CreateSubnetDto(varCreateSubnetDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "parentSubnetId")
		delete(additionalProperties, "networkAddress")
		delete(additionalProperties, "prefixLength")
		delete(additionalProperties, "isPool")
		delete(additionalProperties, "vrfId")
		delete(additionalProperties, "allocationDenylist")
		delete(additionalProperties, "allowedChildOverlapConditions")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSubnetDto struct {
	value *CreateSubnetDto
	isSet bool
}

func (v NullableCreateSubnetDto) Get() *CreateSubnetDto {
	return v.value
}

func (v *NullableCreateSubnetDto) Set(val *CreateSubnetDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubnetDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubnetDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubnetDto(val *CreateSubnetDto) *NullableCreateSubnetDto {
	return &NullableCreateSubnetDto{value: val, isSet: true}
}

func (v NullableCreateSubnetDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubnetDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



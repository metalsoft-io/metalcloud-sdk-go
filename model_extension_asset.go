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

// checks if the ExtensionAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionAsset{}

// ExtensionAsset struct for ExtensionAsset
type ExtensionAsset struct {
	// Label of the asset.
	Label string `json:"label"`
	// Name of the asset.
	Name string `json:"name"`
	// Type of the asset.
	AssetType string `json:"assetType"`
	// URL of the asset.
	Url string `json:"url"`
	// Required assets by this asset.
	RequiredAssets []string `json:"requiredAssets,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionAsset ExtensionAsset

// NewExtensionAsset instantiates a new ExtensionAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionAsset(label string, name string, assetType string, url string) *ExtensionAsset {
	this := ExtensionAsset{}
	this.Label = label
	this.Name = name
	this.AssetType = assetType
	this.Url = url
	return &this
}

// NewExtensionAssetWithDefaults instantiates a new ExtensionAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionAssetWithDefaults() *ExtensionAsset {
	this := ExtensionAsset{}
	return &this
}

// GetLabel returns the Label field value
func (o *ExtensionAsset) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ExtensionAsset) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ExtensionAsset) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *ExtensionAsset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtensionAsset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtensionAsset) SetName(v string) {
	o.Name = v
}

// GetAssetType returns the AssetType field value
func (o *ExtensionAsset) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *ExtensionAsset) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *ExtensionAsset) SetAssetType(v string) {
	o.AssetType = v
}

// GetUrl returns the Url field value
func (o *ExtensionAsset) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExtensionAsset) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExtensionAsset) SetUrl(v string) {
	o.Url = v
}

// GetRequiredAssets returns the RequiredAssets field value if set, zero value otherwise.
func (o *ExtensionAsset) GetRequiredAssets() []string {
	if o == nil || IsNil(o.RequiredAssets) {
		var ret []string
		return ret
	}
	return o.RequiredAssets
}

// GetRequiredAssetsOk returns a tuple with the RequiredAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAsset) GetRequiredAssetsOk() ([]string, bool) {
	if o == nil || IsNil(o.RequiredAssets) {
		return nil, false
	}
	return o.RequiredAssets, true
}

// HasRequiredAssets returns a boolean if a field has been set.
func (o *ExtensionAsset) HasRequiredAssets() bool {
	if o != nil && !IsNil(o.RequiredAssets) {
		return true
	}

	return false
}

// SetRequiredAssets gets a reference to the given []string and assigns it to the RequiredAssets field.
func (o *ExtensionAsset) SetRequiredAssets(v []string) {
	o.RequiredAssets = v
}

func (o ExtensionAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["assetType"] = o.AssetType
	toSerialize["url"] = o.Url
	if !IsNil(o.RequiredAssets) {
		toSerialize["requiredAssets"] = o.RequiredAssets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionAsset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
		"assetType",
		"url",
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

	varExtensionAsset := _ExtensionAsset{}

	err = json.Unmarshal(data, &varExtensionAsset)

	if err != nil {
		return err
	}

	*o = ExtensionAsset(varExtensionAsset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "assetType")
		delete(additionalProperties, "url")
		delete(additionalProperties, "requiredAssets")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionAsset struct {
	value *ExtensionAsset
	isSet bool
}

func (v NullableExtensionAsset) Get() *ExtensionAsset {
	return v.value
}

func (v *NullableExtensionAsset) Set(val *ExtensionAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionAsset(val *ExtensionAsset) *NullableExtensionAsset {
	return &NullableExtensionAsset{value: val, isSet: true}
}

func (v NullableExtensionAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



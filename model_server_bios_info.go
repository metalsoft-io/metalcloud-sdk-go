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

// checks if the ServerBiosInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerBiosInfo{}

// ServerBiosInfo struct for ServerBiosInfo
type ServerBiosInfo struct {
	// The BIOS vendor of the server.
	Vendor *string `json:"vendor,omitempty"`
	// The BIOS version of the server.
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerBiosInfo ServerBiosInfo

// NewServerBiosInfo instantiates a new ServerBiosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerBiosInfo() *ServerBiosInfo {
	this := ServerBiosInfo{}
	return &this
}

// NewServerBiosInfoWithDefaults instantiates a new ServerBiosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerBiosInfoWithDefaults() *ServerBiosInfo {
	this := ServerBiosInfo{}
	return &this
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ServerBiosInfo) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBiosInfo) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ServerBiosInfo) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ServerBiosInfo) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServerBiosInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBiosInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServerBiosInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServerBiosInfo) SetVersion(v string) {
	o.Version = &v
}

func (o ServerBiosInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerBiosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerBiosInfo) UnmarshalJSON(data []byte) (err error) {
	varServerBiosInfo := _ServerBiosInfo{}

	err = json.Unmarshal(data, &varServerBiosInfo)

	if err != nil {
		return err
	}

	*o = ServerBiosInfo(varServerBiosInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerBiosInfo struct {
	value *ServerBiosInfo
	isSet bool
}

func (v NullableServerBiosInfo) Get() *ServerBiosInfo {
	return v.value
}

func (v *NullableServerBiosInfo) Set(val *ServerBiosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServerBiosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServerBiosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerBiosInfo(val *ServerBiosInfo) *NullableServerBiosInfo {
	return &NullableServerBiosInfo{value: val, isSet: true}
}

func (v NullableServerBiosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerBiosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



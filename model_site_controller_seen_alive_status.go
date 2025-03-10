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

// checks if the SiteControllerSeenAliveStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteControllerSeenAliveStatus{}

// SiteControllerSeenAliveStatus struct for SiteControllerSeenAliveStatus
type SiteControllerSeenAliveStatus struct {
	// The site ID
	Id int32 `json:"id"`
	// Number servers
	ServersCount float32 `json:"serversCount"`
	// Number network equipment
	NetworksCount float32 `json:"networksCount"`
	// Number of storage pools
	StoragesCount float32 `json:"storagesCount"`
	// Number of infrastructures
	InfrastructuresCount float32 `json:"infrastructuresCount"`
	AdditionalProperties map[string]interface{}
}

type _SiteControllerSeenAliveStatus SiteControllerSeenAliveStatus

// NewSiteControllerSeenAliveStatus instantiates a new SiteControllerSeenAliveStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteControllerSeenAliveStatus(id int32, serversCount float32, networksCount float32, storagesCount float32, infrastructuresCount float32) *SiteControllerSeenAliveStatus {
	this := SiteControllerSeenAliveStatus{}
	this.Id = id
	this.ServersCount = serversCount
	this.NetworksCount = networksCount
	this.StoragesCount = storagesCount
	this.InfrastructuresCount = infrastructuresCount
	return &this
}

// NewSiteControllerSeenAliveStatusWithDefaults instantiates a new SiteControllerSeenAliveStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteControllerSeenAliveStatusWithDefaults() *SiteControllerSeenAliveStatus {
	this := SiteControllerSeenAliveStatus{}
	return &this
}

// GetId returns the Id field value
func (o *SiteControllerSeenAliveStatus) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SiteControllerSeenAliveStatus) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SiteControllerSeenAliveStatus) SetId(v int32) {
	o.Id = v
}

// GetServersCount returns the ServersCount field value
func (o *SiteControllerSeenAliveStatus) GetServersCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServersCount
}

// GetServersCountOk returns a tuple with the ServersCount field value
// and a boolean to check if the value has been set.
func (o *SiteControllerSeenAliveStatus) GetServersCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServersCount, true
}

// SetServersCount sets field value
func (o *SiteControllerSeenAliveStatus) SetServersCount(v float32) {
	o.ServersCount = v
}

// GetNetworksCount returns the NetworksCount field value
func (o *SiteControllerSeenAliveStatus) GetNetworksCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworksCount
}

// GetNetworksCountOk returns a tuple with the NetworksCount field value
// and a boolean to check if the value has been set.
func (o *SiteControllerSeenAliveStatus) GetNetworksCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworksCount, true
}

// SetNetworksCount sets field value
func (o *SiteControllerSeenAliveStatus) SetNetworksCount(v float32) {
	o.NetworksCount = v
}

// GetStoragesCount returns the StoragesCount field value
func (o *SiteControllerSeenAliveStatus) GetStoragesCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StoragesCount
}

// GetStoragesCountOk returns a tuple with the StoragesCount field value
// and a boolean to check if the value has been set.
func (o *SiteControllerSeenAliveStatus) GetStoragesCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoragesCount, true
}

// SetStoragesCount sets field value
func (o *SiteControllerSeenAliveStatus) SetStoragesCount(v float32) {
	o.StoragesCount = v
}

// GetInfrastructuresCount returns the InfrastructuresCount field value
func (o *SiteControllerSeenAliveStatus) GetInfrastructuresCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfrastructuresCount
}

// GetInfrastructuresCountOk returns a tuple with the InfrastructuresCount field value
// and a boolean to check if the value has been set.
func (o *SiteControllerSeenAliveStatus) GetInfrastructuresCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructuresCount, true
}

// SetInfrastructuresCount sets field value
func (o *SiteControllerSeenAliveStatus) SetInfrastructuresCount(v float32) {
	o.InfrastructuresCount = v
}

func (o SiteControllerSeenAliveStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteControllerSeenAliveStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serversCount"] = o.ServersCount
	toSerialize["networksCount"] = o.NetworksCount
	toSerialize["storagesCount"] = o.StoragesCount
	toSerialize["infrastructuresCount"] = o.InfrastructuresCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteControllerSeenAliveStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serversCount",
		"networksCount",
		"storagesCount",
		"infrastructuresCount",
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

	varSiteControllerSeenAliveStatus := _SiteControllerSeenAliveStatus{}

	err = json.Unmarshal(data, &varSiteControllerSeenAliveStatus)

	if err != nil {
		return err
	}

	*o = SiteControllerSeenAliveStatus(varSiteControllerSeenAliveStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serversCount")
		delete(additionalProperties, "networksCount")
		delete(additionalProperties, "storagesCount")
		delete(additionalProperties, "infrastructuresCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteControllerSeenAliveStatus struct {
	value *SiteControllerSeenAliveStatus
	isSet bool
}

func (v NullableSiteControllerSeenAliveStatus) Get() *SiteControllerSeenAliveStatus {
	return v.value
}

func (v *NullableSiteControllerSeenAliveStatus) Set(val *SiteControllerSeenAliveStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteControllerSeenAliveStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteControllerSeenAliveStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteControllerSeenAliveStatus(val *SiteControllerSeenAliveStatus) *NullableSiteControllerSeenAliveStatus {
	return &NullableSiteControllerSeenAliveStatus{value: val, isSet: true}
}

func (v NullableSiteControllerSeenAliveStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteControllerSeenAliveStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



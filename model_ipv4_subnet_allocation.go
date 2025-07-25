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

// checks if the Ipv4SubnetAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv4SubnetAllocation{}

// Ipv4SubnetAllocation struct for Ipv4SubnetAllocation
type Ipv4SubnetAllocation struct {
	Id int32 `json:"id"`
	Scope ResourceScope `json:"scope"`
	Status ResourceAllocationStatus `json:"status"`
	NetworkAddress string `json:"networkAddress"`
	PrefixLength int32 `json:"prefixLength"`
	Gateway string `json:"gateway"`
	GatewayPlacement NullableSubnetGatewayPlacement `json:"gatewayPlacement"`
	AdditionalProperties map[string]interface{}
}

type _Ipv4SubnetAllocation Ipv4SubnetAllocation

// NewIpv4SubnetAllocation instantiates a new Ipv4SubnetAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv4SubnetAllocation(id int32, scope ResourceScope, status ResourceAllocationStatus, networkAddress string, prefixLength int32, gateway string, gatewayPlacement NullableSubnetGatewayPlacement) *Ipv4SubnetAllocation {
	this := Ipv4SubnetAllocation{}
	this.Id = id
	this.Scope = scope
	this.Status = status
	this.NetworkAddress = networkAddress
	this.PrefixLength = prefixLength
	this.Gateway = gateway
	this.GatewayPlacement = gatewayPlacement
	return &this
}

// NewIpv4SubnetAllocationWithDefaults instantiates a new Ipv4SubnetAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv4SubnetAllocationWithDefaults() *Ipv4SubnetAllocation {
	this := Ipv4SubnetAllocation{}
	return &this
}

// GetId returns the Id field value
func (o *Ipv4SubnetAllocation) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Ipv4SubnetAllocation) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Ipv4SubnetAllocation) SetId(v int32) {
	o.Id = v
}

// GetScope returns the Scope field value
func (o *Ipv4SubnetAllocation) GetScope() ResourceScope {
	if o == nil {
		var ret ResourceScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *Ipv4SubnetAllocation) GetScopeOk() (*ResourceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *Ipv4SubnetAllocation) SetScope(v ResourceScope) {
	o.Scope = v
}

// GetStatus returns the Status field value
func (o *Ipv4SubnetAllocation) GetStatus() ResourceAllocationStatus {
	if o == nil {
		var ret ResourceAllocationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Ipv4SubnetAllocation) GetStatusOk() (*ResourceAllocationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Ipv4SubnetAllocation) SetStatus(v ResourceAllocationStatus) {
	o.Status = v
}

// GetNetworkAddress returns the NetworkAddress field value
func (o *Ipv4SubnetAllocation) GetNetworkAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkAddress
}

// GetNetworkAddressOk returns a tuple with the NetworkAddress field value
// and a boolean to check if the value has been set.
func (o *Ipv4SubnetAllocation) GetNetworkAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkAddress, true
}

// SetNetworkAddress sets field value
func (o *Ipv4SubnetAllocation) SetNetworkAddress(v string) {
	o.NetworkAddress = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *Ipv4SubnetAllocation) GetPrefixLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *Ipv4SubnetAllocation) GetPrefixLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *Ipv4SubnetAllocation) SetPrefixLength(v int32) {
	o.PrefixLength = v
}

// GetGateway returns the Gateway field value
func (o *Ipv4SubnetAllocation) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *Ipv4SubnetAllocation) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *Ipv4SubnetAllocation) SetGateway(v string) {
	o.Gateway = v
}

// GetGatewayPlacement returns the GatewayPlacement field value
// If the value is explicit nil, the zero value for SubnetGatewayPlacement will be returned
func (o *Ipv4SubnetAllocation) GetGatewayPlacement() SubnetGatewayPlacement {
	if o == nil || o.GatewayPlacement.Get() == nil {
		var ret SubnetGatewayPlacement
		return ret
	}

	return *o.GatewayPlacement.Get()
}

// GetGatewayPlacementOk returns a tuple with the GatewayPlacement field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ipv4SubnetAllocation) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayPlacement.Get(), o.GatewayPlacement.IsSet()
}

// SetGatewayPlacement sets field value
func (o *Ipv4SubnetAllocation) SetGatewayPlacement(v SubnetGatewayPlacement) {
	o.GatewayPlacement.Set(&v)
}

func (o Ipv4SubnetAllocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv4SubnetAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["scope"] = o.Scope
	toSerialize["status"] = o.Status
	toSerialize["networkAddress"] = o.NetworkAddress
	toSerialize["prefixLength"] = o.PrefixLength
	toSerialize["gateway"] = o.Gateway
	toSerialize["gatewayPlacement"] = o.GatewayPlacement.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ipv4SubnetAllocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"scope",
		"status",
		"networkAddress",
		"prefixLength",
		"gateway",
		"gatewayPlacement",
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

	varIpv4SubnetAllocation := _Ipv4SubnetAllocation{}

	err = json.Unmarshal(data, &varIpv4SubnetAllocation)

	if err != nil {
		return err
	}

	*o = Ipv4SubnetAllocation(varIpv4SubnetAllocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "status")
		delete(additionalProperties, "networkAddress")
		delete(additionalProperties, "prefixLength")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gatewayPlacement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpv4SubnetAllocation struct {
	value *Ipv4SubnetAllocation
	isSet bool
}

func (v NullableIpv4SubnetAllocation) Get() *Ipv4SubnetAllocation {
	return v.value
}

func (v *NullableIpv4SubnetAllocation) Set(val *Ipv4SubnetAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv4SubnetAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv4SubnetAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv4SubnetAllocation(val *Ipv4SubnetAllocation) *NullableIpv4SubnetAllocation {
	return &NullableIpv4SubnetAllocation{value: val, isSet: true}
}

func (v NullableIpv4SubnetAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv4SubnetAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



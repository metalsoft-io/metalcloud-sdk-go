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

// checks if the CreateManualIpv6SubnetAllocationStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateManualIpv6SubnetAllocationStrategy{}

// CreateManualIpv6SubnetAllocationStrategy struct for CreateManualIpv6SubnetAllocationStrategy
type CreateManualIpv6SubnetAllocationStrategy struct {
	Kind AllocationStrategyKind `json:"kind"`
	Scope CreateResourceScope `json:"scope"`
	GatewayPlacement *SubnetGatewayPlacement `json:"gatewayPlacement,omitempty"`
	SubnetId int32 `json:"subnetId"`
	AdditionalProperties map[string]interface{}
}

type _CreateManualIpv6SubnetAllocationStrategy CreateManualIpv6SubnetAllocationStrategy

// NewCreateManualIpv6SubnetAllocationStrategy instantiates a new CreateManualIpv6SubnetAllocationStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateManualIpv6SubnetAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, subnetId int32) *CreateManualIpv6SubnetAllocationStrategy {
	this := CreateManualIpv6SubnetAllocationStrategy{}
	this.Kind = kind
	this.Scope = scope
	var gatewayPlacement SubnetGatewayPlacement = SUBNETGATEWAYPLACEMENT_DEFAULT
	this.GatewayPlacement = &gatewayPlacement
	this.SubnetId = subnetId
	return &this
}

// NewCreateManualIpv6SubnetAllocationStrategyWithDefaults instantiates a new CreateManualIpv6SubnetAllocationStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateManualIpv6SubnetAllocationStrategyWithDefaults() *CreateManualIpv6SubnetAllocationStrategy {
	this := CreateManualIpv6SubnetAllocationStrategy{}
	var gatewayPlacement SubnetGatewayPlacement = SUBNETGATEWAYPLACEMENT_DEFAULT
	this.GatewayPlacement = &gatewayPlacement
	return &this
}

// GetKind returns the Kind field value
func (o *CreateManualIpv6SubnetAllocationStrategy) GetKind() AllocationStrategyKind {
	if o == nil {
		var ret AllocationStrategyKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CreateManualIpv6SubnetAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *CreateManualIpv6SubnetAllocationStrategy) SetKind(v AllocationStrategyKind) {
	o.Kind = v
}

// GetScope returns the Scope field value
func (o *CreateManualIpv6SubnetAllocationStrategy) GetScope() CreateResourceScope {
	if o == nil {
		var ret CreateResourceScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CreateManualIpv6SubnetAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CreateManualIpv6SubnetAllocationStrategy) SetScope(v CreateResourceScope) {
	o.Scope = v
}

// GetGatewayPlacement returns the GatewayPlacement field value if set, zero value otherwise.
func (o *CreateManualIpv6SubnetAllocationStrategy) GetGatewayPlacement() SubnetGatewayPlacement {
	if o == nil || IsNil(o.GatewayPlacement) {
		var ret SubnetGatewayPlacement
		return ret
	}
	return *o.GatewayPlacement
}

// GetGatewayPlacementOk returns a tuple with the GatewayPlacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateManualIpv6SubnetAllocationStrategy) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool) {
	if o == nil || IsNil(o.GatewayPlacement) {
		return nil, false
	}
	return o.GatewayPlacement, true
}

// HasGatewayPlacement returns a boolean if a field has been set.
func (o *CreateManualIpv6SubnetAllocationStrategy) HasGatewayPlacement() bool {
	if o != nil && !IsNil(o.GatewayPlacement) {
		return true
	}

	return false
}

// SetGatewayPlacement gets a reference to the given SubnetGatewayPlacement and assigns it to the GatewayPlacement field.
func (o *CreateManualIpv6SubnetAllocationStrategy) SetGatewayPlacement(v SubnetGatewayPlacement) {
	o.GatewayPlacement = &v
}

// GetSubnetId returns the SubnetId field value
func (o *CreateManualIpv6SubnetAllocationStrategy) GetSubnetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *CreateManualIpv6SubnetAllocationStrategy) GetSubnetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *CreateManualIpv6SubnetAllocationStrategy) SetSubnetId(v int32) {
	o.SubnetId = v
}

func (o CreateManualIpv6SubnetAllocationStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateManualIpv6SubnetAllocationStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["scope"] = o.Scope
	if !IsNil(o.GatewayPlacement) {
		toSerialize["gatewayPlacement"] = o.GatewayPlacement
	}
	toSerialize["subnetId"] = o.SubnetId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateManualIpv6SubnetAllocationStrategy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"scope",
		"subnetId",
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

	varCreateManualIpv6SubnetAllocationStrategy := _CreateManualIpv6SubnetAllocationStrategy{}

	err = json.Unmarshal(data, &varCreateManualIpv6SubnetAllocationStrategy)

	if err != nil {
		return err
	}

	*o = CreateManualIpv6SubnetAllocationStrategy(varCreateManualIpv6SubnetAllocationStrategy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kind")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "gatewayPlacement")
		delete(additionalProperties, "subnetId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateManualIpv6SubnetAllocationStrategy struct {
	value *CreateManualIpv6SubnetAllocationStrategy
	isSet bool
}

func (v NullableCreateManualIpv6SubnetAllocationStrategy) Get() *CreateManualIpv6SubnetAllocationStrategy {
	return v.value
}

func (v *NullableCreateManualIpv6SubnetAllocationStrategy) Set(val *CreateManualIpv6SubnetAllocationStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateManualIpv6SubnetAllocationStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateManualIpv6SubnetAllocationStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateManualIpv6SubnetAllocationStrategy(val *CreateManualIpv6SubnetAllocationStrategy) *NullableCreateManualIpv6SubnetAllocationStrategy {
	return &NullableCreateManualIpv6SubnetAllocationStrategy{value: val, isSet: true}
}

func (v NullableCreateManualIpv6SubnetAllocationStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateManualIpv6SubnetAllocationStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



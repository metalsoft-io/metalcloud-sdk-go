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

// checks if the ServerInstanceInterfaceUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceInterfaceUpdate{}

// ServerInstanceInterfaceUpdate struct for ServerInstanceInterfaceUpdate
type ServerInstanceInterfaceUpdate struct {
	CapacityMbps *int32 `json:"capacityMbps,omitempty"`
	// The ID of the network to which this interface is to be attached to.
	NetworkId *int32 `json:"networkId,omitempty"`
	// The server instance interface label.
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceInterfaceUpdate ServerInstanceInterfaceUpdate

// NewServerInstanceInterfaceUpdate instantiates a new ServerInstanceInterfaceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceInterfaceUpdate() *ServerInstanceInterfaceUpdate {
	this := ServerInstanceInterfaceUpdate{}
	return &this
}

// NewServerInstanceInterfaceUpdateWithDefaults instantiates a new ServerInstanceInterfaceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceInterfaceUpdateWithDefaults() *ServerInstanceInterfaceUpdate {
	this := ServerInstanceInterfaceUpdate{}
	return &this
}

// GetCapacityMbps returns the CapacityMbps field value if set, zero value otherwise.
func (o *ServerInstanceInterfaceUpdate) GetCapacityMbps() int32 {
	if o == nil || IsNil(o.CapacityMbps) {
		var ret int32
		return ret
	}
	return *o.CapacityMbps
}

// GetCapacityMbpsOk returns a tuple with the CapacityMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceInterfaceUpdate) GetCapacityMbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.CapacityMbps) {
		return nil, false
	}
	return o.CapacityMbps, true
}

// HasCapacityMbps returns a boolean if a field has been set.
func (o *ServerInstanceInterfaceUpdate) HasCapacityMbps() bool {
	if o != nil && !IsNil(o.CapacityMbps) {
		return true
	}

	return false
}

// SetCapacityMbps gets a reference to the given int32 and assigns it to the CapacityMbps field.
func (o *ServerInstanceInterfaceUpdate) SetCapacityMbps(v int32) {
	o.CapacityMbps = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *ServerInstanceInterfaceUpdate) GetNetworkId() int32 {
	if o == nil || IsNil(o.NetworkId) {
		var ret int32
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceInterfaceUpdate) GetNetworkIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *ServerInstanceInterfaceUpdate) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given int32 and assigns it to the NetworkId field.
func (o *ServerInstanceInterfaceUpdate) SetNetworkId(v int32) {
	o.NetworkId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServerInstanceInterfaceUpdate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceInterfaceUpdate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServerInstanceInterfaceUpdate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServerInstanceInterfaceUpdate) SetLabel(v string) {
	o.Label = &v
}

func (o ServerInstanceInterfaceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceInterfaceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CapacityMbps) {
		toSerialize["capacityMbps"] = o.CapacityMbps
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceInterfaceUpdate) UnmarshalJSON(data []byte) (err error) {
	varServerInstanceInterfaceUpdate := _ServerInstanceInterfaceUpdate{}

	err = json.Unmarshal(data, &varServerInstanceInterfaceUpdate)

	if err != nil {
		return err
	}

	*o = ServerInstanceInterfaceUpdate(varServerInstanceInterfaceUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capacityMbps")
		delete(additionalProperties, "networkId")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceInterfaceUpdate struct {
	value *ServerInstanceInterfaceUpdate
	isSet bool
}

func (v NullableServerInstanceInterfaceUpdate) Get() *ServerInstanceInterfaceUpdate {
	return v.value
}

func (v *NullableServerInstanceInterfaceUpdate) Set(val *ServerInstanceInterfaceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceInterfaceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceInterfaceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceInterfaceUpdate(val *ServerInstanceInterfaceUpdate) *NullableServerInstanceInterfaceUpdate {
	return &NullableServerInstanceInterfaceUpdate{value: val, isSet: true}
}

func (v NullableServerInstanceInterfaceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceInterfaceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



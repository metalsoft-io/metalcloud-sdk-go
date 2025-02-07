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

// checks if the EthernetFlatL3FabricDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthernetFlatL3FabricDto{}

// EthernetFlatL3FabricDto struct for EthernetFlatL3FabricDto
type EthernetFlatL3FabricDto struct {
	DefaultNetworkProfileId float32 `json:"defaultNetworkProfileId"`
	GnmiMonitoringEnabled *bool `json:"gnmiMonitoringEnabled,omitempty"`
	SyslogMonitoringEnabled *bool `json:"syslogMonitoringEnabled,omitempty"`
	ZeroTouchEnabled *bool `json:"zeroTouchEnabled,omitempty"`
	AllocateDefaultVlan *bool `json:"allocateDefaultVlan,omitempty"`
	AsnRanges []string `json:"asnRanges,omitempty"`
	DefaultVlan *float32 `json:"defaultVlan,omitempty"`
	ExtraInternalIPsPerSubnet *float32 `json:"extraInternalIPsPerSubnet,omitempty"`
	LagRanges []string `json:"lagRanges,omitempty"`
	LeafSwitchesHaveMlagPairs *bool `json:"leafSwitchesHaveMlagPairs,omitempty"`
	// MLAG ID ranges. Values must be between 1 and 4096.
	MlagRanges []string `json:"mlagRanges,omitempty"`
	// Number of spines next to leaf switches
	NumberOfSpinesNextToLeafSwitches float32 `json:"numberOfSpinesNextToLeafSwitches"`
	PreventVlanCleanup []string `json:"preventVlanCleanup,omitempty"`
	PreventCleanupFromUplinks *bool `json:"preventCleanupFromUplinks,omitempty"`
	ReservedVlans []string `json:"reservedVlans,omitempty"`
	// Array of VLAN range strings in format \"start-end\"
	VlanRanges []string `json:"vlanRanges,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EthernetFlatL3FabricDto EthernetFlatL3FabricDto

// NewEthernetFlatL3FabricDto instantiates a new EthernetFlatL3FabricDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthernetFlatL3FabricDto(defaultNetworkProfileId float32, numberOfSpinesNextToLeafSwitches float32) *EthernetFlatL3FabricDto {
	this := EthernetFlatL3FabricDto{}
	this.DefaultNetworkProfileId = defaultNetworkProfileId
	this.NumberOfSpinesNextToLeafSwitches = numberOfSpinesNextToLeafSwitches
	return &this
}

// NewEthernetFlatL3FabricDtoWithDefaults instantiates a new EthernetFlatL3FabricDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthernetFlatL3FabricDtoWithDefaults() *EthernetFlatL3FabricDto {
	this := EthernetFlatL3FabricDto{}
	return &this
}

// GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field value
func (o *EthernetFlatL3FabricDto) GetDefaultNetworkProfileId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DefaultNetworkProfileId
}

// GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field value
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetDefaultNetworkProfileIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultNetworkProfileId, true
}

// SetDefaultNetworkProfileId sets field value
func (o *EthernetFlatL3FabricDto) SetDefaultNetworkProfileId(v float32) {
	o.DefaultNetworkProfileId = v
}

// GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetGnmiMonitoringEnabled() bool {
	if o == nil || IsNil(o.GnmiMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.GnmiMonitoringEnabled
}

// GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetGnmiMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GnmiMonitoringEnabled) {
		return nil, false
	}
	return o.GnmiMonitoringEnabled, true
}

// HasGnmiMonitoringEnabled returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasGnmiMonitoringEnabled() bool {
	if o != nil && !IsNil(o.GnmiMonitoringEnabled) {
		return true
	}

	return false
}

// SetGnmiMonitoringEnabled gets a reference to the given bool and assigns it to the GnmiMonitoringEnabled field.
func (o *EthernetFlatL3FabricDto) SetGnmiMonitoringEnabled(v bool) {
	o.GnmiMonitoringEnabled = &v
}

// GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetSyslogMonitoringEnabled() bool {
	if o == nil || IsNil(o.SyslogMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.SyslogMonitoringEnabled
}

// GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetSyslogMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyslogMonitoringEnabled) {
		return nil, false
	}
	return o.SyslogMonitoringEnabled, true
}

// HasSyslogMonitoringEnabled returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasSyslogMonitoringEnabled() bool {
	if o != nil && !IsNil(o.SyslogMonitoringEnabled) {
		return true
	}

	return false
}

// SetSyslogMonitoringEnabled gets a reference to the given bool and assigns it to the SyslogMonitoringEnabled field.
func (o *EthernetFlatL3FabricDto) SetSyslogMonitoringEnabled(v bool) {
	o.SyslogMonitoringEnabled = &v
}

// GetZeroTouchEnabled returns the ZeroTouchEnabled field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetZeroTouchEnabled() bool {
	if o == nil || IsNil(o.ZeroTouchEnabled) {
		var ret bool
		return ret
	}
	return *o.ZeroTouchEnabled
}

// GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetZeroTouchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ZeroTouchEnabled) {
		return nil, false
	}
	return o.ZeroTouchEnabled, true
}

// HasZeroTouchEnabled returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasZeroTouchEnabled() bool {
	if o != nil && !IsNil(o.ZeroTouchEnabled) {
		return true
	}

	return false
}

// SetZeroTouchEnabled gets a reference to the given bool and assigns it to the ZeroTouchEnabled field.
func (o *EthernetFlatL3FabricDto) SetZeroTouchEnabled(v bool) {
	o.ZeroTouchEnabled = &v
}

// GetAllocateDefaultVlan returns the AllocateDefaultVlan field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetAllocateDefaultVlan() bool {
	if o == nil || IsNil(o.AllocateDefaultVlan) {
		var ret bool
		return ret
	}
	return *o.AllocateDefaultVlan
}

// GetAllocateDefaultVlanOk returns a tuple with the AllocateDefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetAllocateDefaultVlanOk() (*bool, bool) {
	if o == nil || IsNil(o.AllocateDefaultVlan) {
		return nil, false
	}
	return o.AllocateDefaultVlan, true
}

// HasAllocateDefaultVlan returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasAllocateDefaultVlan() bool {
	if o != nil && !IsNil(o.AllocateDefaultVlan) {
		return true
	}

	return false
}

// SetAllocateDefaultVlan gets a reference to the given bool and assigns it to the AllocateDefaultVlan field.
func (o *EthernetFlatL3FabricDto) SetAllocateDefaultVlan(v bool) {
	o.AllocateDefaultVlan = &v
}

// GetAsnRanges returns the AsnRanges field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetAsnRanges() []string {
	if o == nil || IsNil(o.AsnRanges) {
		var ret []string
		return ret
	}
	return o.AsnRanges
}

// GetAsnRangesOk returns a tuple with the AsnRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetAsnRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.AsnRanges) {
		return nil, false
	}
	return o.AsnRanges, true
}

// HasAsnRanges returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasAsnRanges() bool {
	if o != nil && !IsNil(o.AsnRanges) {
		return true
	}

	return false
}

// SetAsnRanges gets a reference to the given []string and assigns it to the AsnRanges field.
func (o *EthernetFlatL3FabricDto) SetAsnRanges(v []string) {
	o.AsnRanges = v
}

// GetDefaultVlan returns the DefaultVlan field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetDefaultVlan() float32 {
	if o == nil || IsNil(o.DefaultVlan) {
		var ret float32
		return ret
	}
	return *o.DefaultVlan
}

// GetDefaultVlanOk returns a tuple with the DefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetDefaultVlanOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultVlan) {
		return nil, false
	}
	return o.DefaultVlan, true
}

// HasDefaultVlan returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasDefaultVlan() bool {
	if o != nil && !IsNil(o.DefaultVlan) {
		return true
	}

	return false
}

// SetDefaultVlan gets a reference to the given float32 and assigns it to the DefaultVlan field.
func (o *EthernetFlatL3FabricDto) SetDefaultVlan(v float32) {
	o.DefaultVlan = &v
}

// GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetExtraInternalIPsPerSubnet() float32 {
	if o == nil || IsNil(o.ExtraInternalIPsPerSubnet) {
		var ret float32
		return ret
	}
	return *o.ExtraInternalIPsPerSubnet
}

// GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetExtraInternalIPsPerSubnetOk() (*float32, bool) {
	if o == nil || IsNil(o.ExtraInternalIPsPerSubnet) {
		return nil, false
	}
	return o.ExtraInternalIPsPerSubnet, true
}

// HasExtraInternalIPsPerSubnet returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasExtraInternalIPsPerSubnet() bool {
	if o != nil && !IsNil(o.ExtraInternalIPsPerSubnet) {
		return true
	}

	return false
}

// SetExtraInternalIPsPerSubnet gets a reference to the given float32 and assigns it to the ExtraInternalIPsPerSubnet field.
func (o *EthernetFlatL3FabricDto) SetExtraInternalIPsPerSubnet(v float32) {
	o.ExtraInternalIPsPerSubnet = &v
}

// GetLagRanges returns the LagRanges field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetLagRanges() []string {
	if o == nil || IsNil(o.LagRanges) {
		var ret []string
		return ret
	}
	return o.LagRanges
}

// GetLagRangesOk returns a tuple with the LagRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetLagRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.LagRanges) {
		return nil, false
	}
	return o.LagRanges, true
}

// HasLagRanges returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasLagRanges() bool {
	if o != nil && !IsNil(o.LagRanges) {
		return true
	}

	return false
}

// SetLagRanges gets a reference to the given []string and assigns it to the LagRanges field.
func (o *EthernetFlatL3FabricDto) SetLagRanges(v []string) {
	o.LagRanges = v
}

// GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetLeafSwitchesHaveMlagPairs() bool {
	if o == nil || IsNil(o.LeafSwitchesHaveMlagPairs) {
		var ret bool
		return ret
	}
	return *o.LeafSwitchesHaveMlagPairs
}

// GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool) {
	if o == nil || IsNil(o.LeafSwitchesHaveMlagPairs) {
		return nil, false
	}
	return o.LeafSwitchesHaveMlagPairs, true
}

// HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasLeafSwitchesHaveMlagPairs() bool {
	if o != nil && !IsNil(o.LeafSwitchesHaveMlagPairs) {
		return true
	}

	return false
}

// SetLeafSwitchesHaveMlagPairs gets a reference to the given bool and assigns it to the LeafSwitchesHaveMlagPairs field.
func (o *EthernetFlatL3FabricDto) SetLeafSwitchesHaveMlagPairs(v bool) {
	o.LeafSwitchesHaveMlagPairs = &v
}

// GetMlagRanges returns the MlagRanges field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetMlagRanges() []string {
	if o == nil || IsNil(o.MlagRanges) {
		var ret []string
		return ret
	}
	return o.MlagRanges
}

// GetMlagRangesOk returns a tuple with the MlagRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetMlagRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.MlagRanges) {
		return nil, false
	}
	return o.MlagRanges, true
}

// HasMlagRanges returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasMlagRanges() bool {
	if o != nil && !IsNil(o.MlagRanges) {
		return true
	}

	return false
}

// SetMlagRanges gets a reference to the given []string and assigns it to the MlagRanges field.
func (o *EthernetFlatL3FabricDto) SetMlagRanges(v []string) {
	o.MlagRanges = v
}

// GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field value
func (o *EthernetFlatL3FabricDto) GetNumberOfSpinesNextToLeafSwitches() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumberOfSpinesNextToLeafSwitches
}

// GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field value
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetNumberOfSpinesNextToLeafSwitchesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfSpinesNextToLeafSwitches, true
}

// SetNumberOfSpinesNextToLeafSwitches sets field value
func (o *EthernetFlatL3FabricDto) SetNumberOfSpinesNextToLeafSwitches(v float32) {
	o.NumberOfSpinesNextToLeafSwitches = v
}

// GetPreventVlanCleanup returns the PreventVlanCleanup field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetPreventVlanCleanup() []string {
	if o == nil || IsNil(o.PreventVlanCleanup) {
		var ret []string
		return ret
	}
	return o.PreventVlanCleanup
}

// GetPreventVlanCleanupOk returns a tuple with the PreventVlanCleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetPreventVlanCleanupOk() ([]string, bool) {
	if o == nil || IsNil(o.PreventVlanCleanup) {
		return nil, false
	}
	return o.PreventVlanCleanup, true
}

// HasPreventVlanCleanup returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasPreventVlanCleanup() bool {
	if o != nil && !IsNil(o.PreventVlanCleanup) {
		return true
	}

	return false
}

// SetPreventVlanCleanup gets a reference to the given []string and assigns it to the PreventVlanCleanup field.
func (o *EthernetFlatL3FabricDto) SetPreventVlanCleanup(v []string) {
	o.PreventVlanCleanup = v
}

// GetPreventCleanupFromUplinks returns the PreventCleanupFromUplinks field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetPreventCleanupFromUplinks() bool {
	if o == nil || IsNil(o.PreventCleanupFromUplinks) {
		var ret bool
		return ret
	}
	return *o.PreventCleanupFromUplinks
}

// GetPreventCleanupFromUplinksOk returns a tuple with the PreventCleanupFromUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetPreventCleanupFromUplinksOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventCleanupFromUplinks) {
		return nil, false
	}
	return o.PreventCleanupFromUplinks, true
}

// HasPreventCleanupFromUplinks returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasPreventCleanupFromUplinks() bool {
	if o != nil && !IsNil(o.PreventCleanupFromUplinks) {
		return true
	}

	return false
}

// SetPreventCleanupFromUplinks gets a reference to the given bool and assigns it to the PreventCleanupFromUplinks field.
func (o *EthernetFlatL3FabricDto) SetPreventCleanupFromUplinks(v bool) {
	o.PreventCleanupFromUplinks = &v
}

// GetReservedVlans returns the ReservedVlans field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetReservedVlans() []string {
	if o == nil || IsNil(o.ReservedVlans) {
		var ret []string
		return ret
	}
	return o.ReservedVlans
}

// GetReservedVlansOk returns a tuple with the ReservedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetReservedVlansOk() ([]string, bool) {
	if o == nil || IsNil(o.ReservedVlans) {
		return nil, false
	}
	return o.ReservedVlans, true
}

// HasReservedVlans returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasReservedVlans() bool {
	if o != nil && !IsNil(o.ReservedVlans) {
		return true
	}

	return false
}

// SetReservedVlans gets a reference to the given []string and assigns it to the ReservedVlans field.
func (o *EthernetFlatL3FabricDto) SetReservedVlans(v []string) {
	o.ReservedVlans = v
}

// GetVlanRanges returns the VlanRanges field value if set, zero value otherwise.
func (o *EthernetFlatL3FabricDto) GetVlanRanges() []string {
	if o == nil || IsNil(o.VlanRanges) {
		var ret []string
		return ret
	}
	return o.VlanRanges
}

// GetVlanRangesOk returns a tuple with the VlanRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetFlatL3FabricDto) GetVlanRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.VlanRanges) {
		return nil, false
	}
	return o.VlanRanges, true
}

// HasVlanRanges returns a boolean if a field has been set.
func (o *EthernetFlatL3FabricDto) HasVlanRanges() bool {
	if o != nil && !IsNil(o.VlanRanges) {
		return true
	}

	return false
}

// SetVlanRanges gets a reference to the given []string and assigns it to the VlanRanges field.
func (o *EthernetFlatL3FabricDto) SetVlanRanges(v []string) {
	o.VlanRanges = v
}

func (o EthernetFlatL3FabricDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthernetFlatL3FabricDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultNetworkProfileId"] = o.DefaultNetworkProfileId
	if !IsNil(o.GnmiMonitoringEnabled) {
		toSerialize["gnmiMonitoringEnabled"] = o.GnmiMonitoringEnabled
	}
	if !IsNil(o.SyslogMonitoringEnabled) {
		toSerialize["syslogMonitoringEnabled"] = o.SyslogMonitoringEnabled
	}
	if !IsNil(o.ZeroTouchEnabled) {
		toSerialize["zeroTouchEnabled"] = o.ZeroTouchEnabled
	}
	if !IsNil(o.AllocateDefaultVlan) {
		toSerialize["allocateDefaultVlan"] = o.AllocateDefaultVlan
	}
	if !IsNil(o.AsnRanges) {
		toSerialize["asnRanges"] = o.AsnRanges
	}
	if !IsNil(o.DefaultVlan) {
		toSerialize["defaultVlan"] = o.DefaultVlan
	}
	if !IsNil(o.ExtraInternalIPsPerSubnet) {
		toSerialize["extraInternalIPsPerSubnet"] = o.ExtraInternalIPsPerSubnet
	}
	if !IsNil(o.LagRanges) {
		toSerialize["lagRanges"] = o.LagRanges
	}
	if !IsNil(o.LeafSwitchesHaveMlagPairs) {
		toSerialize["leafSwitchesHaveMlagPairs"] = o.LeafSwitchesHaveMlagPairs
	}
	if !IsNil(o.MlagRanges) {
		toSerialize["mlagRanges"] = o.MlagRanges
	}
	toSerialize["numberOfSpinesNextToLeafSwitches"] = o.NumberOfSpinesNextToLeafSwitches
	if !IsNil(o.PreventVlanCleanup) {
		toSerialize["preventVlanCleanup"] = o.PreventVlanCleanup
	}
	if !IsNil(o.PreventCleanupFromUplinks) {
		toSerialize["preventCleanupFromUplinks"] = o.PreventCleanupFromUplinks
	}
	if !IsNil(o.ReservedVlans) {
		toSerialize["reservedVlans"] = o.ReservedVlans
	}
	if !IsNil(o.VlanRanges) {
		toSerialize["vlanRanges"] = o.VlanRanges
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EthernetFlatL3FabricDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultNetworkProfileId",
		"numberOfSpinesNextToLeafSwitches",
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

	varEthernetFlatL3FabricDto := _EthernetFlatL3FabricDto{}

	err = json.Unmarshal(data, &varEthernetFlatL3FabricDto)

	if err != nil {
		return err
	}

	*o = EthernetFlatL3FabricDto(varEthernetFlatL3FabricDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultNetworkProfileId")
		delete(additionalProperties, "gnmiMonitoringEnabled")
		delete(additionalProperties, "syslogMonitoringEnabled")
		delete(additionalProperties, "zeroTouchEnabled")
		delete(additionalProperties, "allocateDefaultVlan")
		delete(additionalProperties, "asnRanges")
		delete(additionalProperties, "defaultVlan")
		delete(additionalProperties, "extraInternalIPsPerSubnet")
		delete(additionalProperties, "lagRanges")
		delete(additionalProperties, "leafSwitchesHaveMlagPairs")
		delete(additionalProperties, "mlagRanges")
		delete(additionalProperties, "numberOfSpinesNextToLeafSwitches")
		delete(additionalProperties, "preventVlanCleanup")
		delete(additionalProperties, "preventCleanupFromUplinks")
		delete(additionalProperties, "reservedVlans")
		delete(additionalProperties, "vlanRanges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEthernetFlatL3FabricDto struct {
	value *EthernetFlatL3FabricDto
	isSet bool
}

func (v NullableEthernetFlatL3FabricDto) Get() *EthernetFlatL3FabricDto {
	return v.value
}

func (v *NullableEthernetFlatL3FabricDto) Set(val *EthernetFlatL3FabricDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEthernetFlatL3FabricDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEthernetFlatL3FabricDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthernetFlatL3FabricDto(val *EthernetFlatL3FabricDto) *NullableEthernetFlatL3FabricDto {
	return &NullableEthernetFlatL3FabricDto{value: val, isSet: true}
}

func (v NullableEthernetFlatL3FabricDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthernetFlatL3FabricDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



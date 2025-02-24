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

// checks if the EthernetEvpnVxlanL3Fabric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthernetEvpnVxlanL3Fabric{}

// EthernetEvpnVxlanL3Fabric struct for EthernetEvpnVxlanL3Fabric
type EthernetEvpnVxlanL3Fabric struct {
	// Unique identifier for the default network profile. Must be a positive integer (minimum: 1) corresponding to an existing profile.
	DefaultNetworkProfileId int32 `json:"defaultNetworkProfileId"`
	// Enables gNMI monitoring for telemetry data collection using the gNMI protocol.
	GnmiMonitoringEnabled *bool `json:"gnmiMonitoringEnabled,omitempty"`
	// Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting.
	SyslogMonitoringEnabled *bool `json:"syslogMonitoringEnabled,omitempty"`
	// Enables zero-touch provisioning for automatic device configuration.
	ZeroTouchEnabled *bool `json:"zeroTouchEnabled,omitempty"`
	// Indicates whether to automatically allocate a default VLAN.
	AllocateDefaultVlan *bool `json:"allocateDefaultVlan,omitempty"`
	// ASN ranges in the format \"start-end\", where each range is an ordered pair with values between 1 and 4294967295.
	AsnRanges []string `json:"asnRanges,omitempty"`
	// Default VLAN ID. Must be a number between 1 and 4096.
	DefaultVlan *int32 `json:"defaultVlan,omitempty"`
	// Extra internal IPs allocated per subnet; valid range is between 1 and 1000.
	ExtraInternalIPsPerSubnet *int32 `json:"extraInternalIPsPerSubnet,omitempty"`
	// Link Aggregation (LAG) ranges in the format \"start-end\"; each range must be within the bounds of 1 to 4096.
	LagRanges []string `json:"lagRanges,omitempty"`
	// Indicates if leaf switches have MLAG pairs.
	LeafSwitchesHaveMlagPairs *bool `json:"leafSwitchesHaveMlagPairs,omitempty"`
	// MLAG ID ranges. Each range must be provided in \"start-end\" format with values between 1 and 4096.
	MlagRanges []string `json:"mlagRanges,omitempty"`
	// Number of spines adjacent to leaf switches. Must be a positive number.
	NumberOfSpinesNextToLeafSwitches int32 `json:"numberOfSpinesNextToLeafSwitches"`
	// VLAN ranges that should be prevented from automatic cleanup. Format must be \"start-end\".
	PreventVlanCleanup []string `json:"preventVlanCleanup,omitempty"`
	// Flag indicating whether cleanup from uplink interfaces should be prevented.
	PreventCleanupFromUplinks *bool `json:"preventCleanupFromUplinks,omitempty"`
	// Reserved VLAN ranges that are excluded from general allocation. Must follow the \"start-end\" format.
	ReservedVlans []string `json:"reservedVlans,omitempty"`
	// Array of VLAN range strings in \"start-end\" format to be used in configuration.
	VlanRanges []string `json:"vlanRanges,omitempty"`
	// The VNI prefix for the EVPN VXLAN fabric.
	VniPrefix *int32 `json:"vniPrefix,omitempty"`
	// VRF ID ranges to be preserved from automatic cleanup. Each range must follow the \"start-end\" format.
	PreventVrfCleanup []string `json:"preventVrfCleanup,omitempty"`
	// Reserved VRF ID ranges that are set aside exclusively for specific network functions. Each range must be provided in the \"start-end\" format.
	ReservedVrfs []string `json:"reservedVrfs,omitempty"`
	// VLAN ranges to be associated with VRF instances. Each value must be an ordered pair specified in the \"start-end\" format.
	VrfVlanRanges []string `json:"vrfVlanRanges,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EthernetEvpnVxlanL3Fabric EthernetEvpnVxlanL3Fabric

// NewEthernetEvpnVxlanL3Fabric instantiates a new EthernetEvpnVxlanL3Fabric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthernetEvpnVxlanL3Fabric(defaultNetworkProfileId int32, numberOfSpinesNextToLeafSwitches int32) *EthernetEvpnVxlanL3Fabric {
	this := EthernetEvpnVxlanL3Fabric{}
	this.DefaultNetworkProfileId = defaultNetworkProfileId
	this.NumberOfSpinesNextToLeafSwitches = numberOfSpinesNextToLeafSwitches
	return &this
}

// NewEthernetEvpnVxlanL3FabricWithDefaults instantiates a new EthernetEvpnVxlanL3Fabric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthernetEvpnVxlanL3FabricWithDefaults() *EthernetEvpnVxlanL3Fabric {
	this := EthernetEvpnVxlanL3Fabric{}
	return &this
}

// GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field value
func (o *EthernetEvpnVxlanL3Fabric) GetDefaultNetworkProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultNetworkProfileId
}

// GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field value
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetDefaultNetworkProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultNetworkProfileId, true
}

// SetDefaultNetworkProfileId sets field value
func (o *EthernetEvpnVxlanL3Fabric) SetDefaultNetworkProfileId(v int32) {
	o.DefaultNetworkProfileId = v
}

// GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetGnmiMonitoringEnabled() bool {
	if o == nil || IsNil(o.GnmiMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.GnmiMonitoringEnabled
}

// GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetGnmiMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GnmiMonitoringEnabled) {
		return nil, false
	}
	return o.GnmiMonitoringEnabled, true
}

// HasGnmiMonitoringEnabled returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasGnmiMonitoringEnabled() bool {
	if o != nil && !IsNil(o.GnmiMonitoringEnabled) {
		return true
	}

	return false
}

// SetGnmiMonitoringEnabled gets a reference to the given bool and assigns it to the GnmiMonitoringEnabled field.
func (o *EthernetEvpnVxlanL3Fabric) SetGnmiMonitoringEnabled(v bool) {
	o.GnmiMonitoringEnabled = &v
}

// GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetSyslogMonitoringEnabled() bool {
	if o == nil || IsNil(o.SyslogMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.SyslogMonitoringEnabled
}

// GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetSyslogMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyslogMonitoringEnabled) {
		return nil, false
	}
	return o.SyslogMonitoringEnabled, true
}

// HasSyslogMonitoringEnabled returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasSyslogMonitoringEnabled() bool {
	if o != nil && !IsNil(o.SyslogMonitoringEnabled) {
		return true
	}

	return false
}

// SetSyslogMonitoringEnabled gets a reference to the given bool and assigns it to the SyslogMonitoringEnabled field.
func (o *EthernetEvpnVxlanL3Fabric) SetSyslogMonitoringEnabled(v bool) {
	o.SyslogMonitoringEnabled = &v
}

// GetZeroTouchEnabled returns the ZeroTouchEnabled field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetZeroTouchEnabled() bool {
	if o == nil || IsNil(o.ZeroTouchEnabled) {
		var ret bool
		return ret
	}
	return *o.ZeroTouchEnabled
}

// GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetZeroTouchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ZeroTouchEnabled) {
		return nil, false
	}
	return o.ZeroTouchEnabled, true
}

// HasZeroTouchEnabled returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasZeroTouchEnabled() bool {
	if o != nil && !IsNil(o.ZeroTouchEnabled) {
		return true
	}

	return false
}

// SetZeroTouchEnabled gets a reference to the given bool and assigns it to the ZeroTouchEnabled field.
func (o *EthernetEvpnVxlanL3Fabric) SetZeroTouchEnabled(v bool) {
	o.ZeroTouchEnabled = &v
}

// GetAllocateDefaultVlan returns the AllocateDefaultVlan field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetAllocateDefaultVlan() bool {
	if o == nil || IsNil(o.AllocateDefaultVlan) {
		var ret bool
		return ret
	}
	return *o.AllocateDefaultVlan
}

// GetAllocateDefaultVlanOk returns a tuple with the AllocateDefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetAllocateDefaultVlanOk() (*bool, bool) {
	if o == nil || IsNil(o.AllocateDefaultVlan) {
		return nil, false
	}
	return o.AllocateDefaultVlan, true
}

// HasAllocateDefaultVlan returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasAllocateDefaultVlan() bool {
	if o != nil && !IsNil(o.AllocateDefaultVlan) {
		return true
	}

	return false
}

// SetAllocateDefaultVlan gets a reference to the given bool and assigns it to the AllocateDefaultVlan field.
func (o *EthernetEvpnVxlanL3Fabric) SetAllocateDefaultVlan(v bool) {
	o.AllocateDefaultVlan = &v
}

// GetAsnRanges returns the AsnRanges field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetAsnRanges() []string {
	if o == nil || IsNil(o.AsnRanges) {
		var ret []string
		return ret
	}
	return o.AsnRanges
}

// GetAsnRangesOk returns a tuple with the AsnRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetAsnRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.AsnRanges) {
		return nil, false
	}
	return o.AsnRanges, true
}

// HasAsnRanges returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasAsnRanges() bool {
	if o != nil && !IsNil(o.AsnRanges) {
		return true
	}

	return false
}

// SetAsnRanges gets a reference to the given []string and assigns it to the AsnRanges field.
func (o *EthernetEvpnVxlanL3Fabric) SetAsnRanges(v []string) {
	o.AsnRanges = v
}

// GetDefaultVlan returns the DefaultVlan field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetDefaultVlan() int32 {
	if o == nil || IsNil(o.DefaultVlan) {
		var ret int32
		return ret
	}
	return *o.DefaultVlan
}

// GetDefaultVlanOk returns a tuple with the DefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetDefaultVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultVlan) {
		return nil, false
	}
	return o.DefaultVlan, true
}

// HasDefaultVlan returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasDefaultVlan() bool {
	if o != nil && !IsNil(o.DefaultVlan) {
		return true
	}

	return false
}

// SetDefaultVlan gets a reference to the given int32 and assigns it to the DefaultVlan field.
func (o *EthernetEvpnVxlanL3Fabric) SetDefaultVlan(v int32) {
	o.DefaultVlan = &v
}

// GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetExtraInternalIPsPerSubnet() int32 {
	if o == nil || IsNil(o.ExtraInternalIPsPerSubnet) {
		var ret int32
		return ret
	}
	return *o.ExtraInternalIPsPerSubnet
}

// GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetExtraInternalIPsPerSubnetOk() (*int32, bool) {
	if o == nil || IsNil(o.ExtraInternalIPsPerSubnet) {
		return nil, false
	}
	return o.ExtraInternalIPsPerSubnet, true
}

// HasExtraInternalIPsPerSubnet returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasExtraInternalIPsPerSubnet() bool {
	if o != nil && !IsNil(o.ExtraInternalIPsPerSubnet) {
		return true
	}

	return false
}

// SetExtraInternalIPsPerSubnet gets a reference to the given int32 and assigns it to the ExtraInternalIPsPerSubnet field.
func (o *EthernetEvpnVxlanL3Fabric) SetExtraInternalIPsPerSubnet(v int32) {
	o.ExtraInternalIPsPerSubnet = &v
}

// GetLagRanges returns the LagRanges field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetLagRanges() []string {
	if o == nil || IsNil(o.LagRanges) {
		var ret []string
		return ret
	}
	return o.LagRanges
}

// GetLagRangesOk returns a tuple with the LagRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetLagRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.LagRanges) {
		return nil, false
	}
	return o.LagRanges, true
}

// HasLagRanges returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasLagRanges() bool {
	if o != nil && !IsNil(o.LagRanges) {
		return true
	}

	return false
}

// SetLagRanges gets a reference to the given []string and assigns it to the LagRanges field.
func (o *EthernetEvpnVxlanL3Fabric) SetLagRanges(v []string) {
	o.LagRanges = v
}

// GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetLeafSwitchesHaveMlagPairs() bool {
	if o == nil || IsNil(o.LeafSwitchesHaveMlagPairs) {
		var ret bool
		return ret
	}
	return *o.LeafSwitchesHaveMlagPairs
}

// GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool) {
	if o == nil || IsNil(o.LeafSwitchesHaveMlagPairs) {
		return nil, false
	}
	return o.LeafSwitchesHaveMlagPairs, true
}

// HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasLeafSwitchesHaveMlagPairs() bool {
	if o != nil && !IsNil(o.LeafSwitchesHaveMlagPairs) {
		return true
	}

	return false
}

// SetLeafSwitchesHaveMlagPairs gets a reference to the given bool and assigns it to the LeafSwitchesHaveMlagPairs field.
func (o *EthernetEvpnVxlanL3Fabric) SetLeafSwitchesHaveMlagPairs(v bool) {
	o.LeafSwitchesHaveMlagPairs = &v
}

// GetMlagRanges returns the MlagRanges field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetMlagRanges() []string {
	if o == nil || IsNil(o.MlagRanges) {
		var ret []string
		return ret
	}
	return o.MlagRanges
}

// GetMlagRangesOk returns a tuple with the MlagRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetMlagRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.MlagRanges) {
		return nil, false
	}
	return o.MlagRanges, true
}

// HasMlagRanges returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasMlagRanges() bool {
	if o != nil && !IsNil(o.MlagRanges) {
		return true
	}

	return false
}

// SetMlagRanges gets a reference to the given []string and assigns it to the MlagRanges field.
func (o *EthernetEvpnVxlanL3Fabric) SetMlagRanges(v []string) {
	o.MlagRanges = v
}

// GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field value
func (o *EthernetEvpnVxlanL3Fabric) GetNumberOfSpinesNextToLeafSwitches() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfSpinesNextToLeafSwitches
}

// GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field value
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetNumberOfSpinesNextToLeafSwitchesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfSpinesNextToLeafSwitches, true
}

// SetNumberOfSpinesNextToLeafSwitches sets field value
func (o *EthernetEvpnVxlanL3Fabric) SetNumberOfSpinesNextToLeafSwitches(v int32) {
	o.NumberOfSpinesNextToLeafSwitches = v
}

// GetPreventVlanCleanup returns the PreventVlanCleanup field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetPreventVlanCleanup() []string {
	if o == nil || IsNil(o.PreventVlanCleanup) {
		var ret []string
		return ret
	}
	return o.PreventVlanCleanup
}

// GetPreventVlanCleanupOk returns a tuple with the PreventVlanCleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetPreventVlanCleanupOk() ([]string, bool) {
	if o == nil || IsNil(o.PreventVlanCleanup) {
		return nil, false
	}
	return o.PreventVlanCleanup, true
}

// HasPreventVlanCleanup returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasPreventVlanCleanup() bool {
	if o != nil && !IsNil(o.PreventVlanCleanup) {
		return true
	}

	return false
}

// SetPreventVlanCleanup gets a reference to the given []string and assigns it to the PreventVlanCleanup field.
func (o *EthernetEvpnVxlanL3Fabric) SetPreventVlanCleanup(v []string) {
	o.PreventVlanCleanup = v
}

// GetPreventCleanupFromUplinks returns the PreventCleanupFromUplinks field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetPreventCleanupFromUplinks() bool {
	if o == nil || IsNil(o.PreventCleanupFromUplinks) {
		var ret bool
		return ret
	}
	return *o.PreventCleanupFromUplinks
}

// GetPreventCleanupFromUplinksOk returns a tuple with the PreventCleanupFromUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetPreventCleanupFromUplinksOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventCleanupFromUplinks) {
		return nil, false
	}
	return o.PreventCleanupFromUplinks, true
}

// HasPreventCleanupFromUplinks returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasPreventCleanupFromUplinks() bool {
	if o != nil && !IsNil(o.PreventCleanupFromUplinks) {
		return true
	}

	return false
}

// SetPreventCleanupFromUplinks gets a reference to the given bool and assigns it to the PreventCleanupFromUplinks field.
func (o *EthernetEvpnVxlanL3Fabric) SetPreventCleanupFromUplinks(v bool) {
	o.PreventCleanupFromUplinks = &v
}

// GetReservedVlans returns the ReservedVlans field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetReservedVlans() []string {
	if o == nil || IsNil(o.ReservedVlans) {
		var ret []string
		return ret
	}
	return o.ReservedVlans
}

// GetReservedVlansOk returns a tuple with the ReservedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetReservedVlansOk() ([]string, bool) {
	if o == nil || IsNil(o.ReservedVlans) {
		return nil, false
	}
	return o.ReservedVlans, true
}

// HasReservedVlans returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasReservedVlans() bool {
	if o != nil && !IsNil(o.ReservedVlans) {
		return true
	}

	return false
}

// SetReservedVlans gets a reference to the given []string and assigns it to the ReservedVlans field.
func (o *EthernetEvpnVxlanL3Fabric) SetReservedVlans(v []string) {
	o.ReservedVlans = v
}

// GetVlanRanges returns the VlanRanges field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetVlanRanges() []string {
	if o == nil || IsNil(o.VlanRanges) {
		var ret []string
		return ret
	}
	return o.VlanRanges
}

// GetVlanRangesOk returns a tuple with the VlanRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetVlanRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.VlanRanges) {
		return nil, false
	}
	return o.VlanRanges, true
}

// HasVlanRanges returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasVlanRanges() bool {
	if o != nil && !IsNil(o.VlanRanges) {
		return true
	}

	return false
}

// SetVlanRanges gets a reference to the given []string and assigns it to the VlanRanges field.
func (o *EthernetEvpnVxlanL3Fabric) SetVlanRanges(v []string) {
	o.VlanRanges = v
}

// GetVniPrefix returns the VniPrefix field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetVniPrefix() int32 {
	if o == nil || IsNil(o.VniPrefix) {
		var ret int32
		return ret
	}
	return *o.VniPrefix
}

// GetVniPrefixOk returns a tuple with the VniPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetVniPrefixOk() (*int32, bool) {
	if o == nil || IsNil(o.VniPrefix) {
		return nil, false
	}
	return o.VniPrefix, true
}

// HasVniPrefix returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasVniPrefix() bool {
	if o != nil && !IsNil(o.VniPrefix) {
		return true
	}

	return false
}

// SetVniPrefix gets a reference to the given int32 and assigns it to the VniPrefix field.
func (o *EthernetEvpnVxlanL3Fabric) SetVniPrefix(v int32) {
	o.VniPrefix = &v
}

// GetPreventVrfCleanup returns the PreventVrfCleanup field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetPreventVrfCleanup() []string {
	if o == nil || IsNil(o.PreventVrfCleanup) {
		var ret []string
		return ret
	}
	return o.PreventVrfCleanup
}

// GetPreventVrfCleanupOk returns a tuple with the PreventVrfCleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetPreventVrfCleanupOk() ([]string, bool) {
	if o == nil || IsNil(o.PreventVrfCleanup) {
		return nil, false
	}
	return o.PreventVrfCleanup, true
}

// HasPreventVrfCleanup returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasPreventVrfCleanup() bool {
	if o != nil && !IsNil(o.PreventVrfCleanup) {
		return true
	}

	return false
}

// SetPreventVrfCleanup gets a reference to the given []string and assigns it to the PreventVrfCleanup field.
func (o *EthernetEvpnVxlanL3Fabric) SetPreventVrfCleanup(v []string) {
	o.PreventVrfCleanup = v
}

// GetReservedVrfs returns the ReservedVrfs field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetReservedVrfs() []string {
	if o == nil || IsNil(o.ReservedVrfs) {
		var ret []string
		return ret
	}
	return o.ReservedVrfs
}

// GetReservedVrfsOk returns a tuple with the ReservedVrfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetReservedVrfsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReservedVrfs) {
		return nil, false
	}
	return o.ReservedVrfs, true
}

// HasReservedVrfs returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasReservedVrfs() bool {
	if o != nil && !IsNil(o.ReservedVrfs) {
		return true
	}

	return false
}

// SetReservedVrfs gets a reference to the given []string and assigns it to the ReservedVrfs field.
func (o *EthernetEvpnVxlanL3Fabric) SetReservedVrfs(v []string) {
	o.ReservedVrfs = v
}

// GetVrfVlanRanges returns the VrfVlanRanges field value if set, zero value otherwise.
func (o *EthernetEvpnVxlanL3Fabric) GetVrfVlanRanges() []string {
	if o == nil || IsNil(o.VrfVlanRanges) {
		var ret []string
		return ret
	}
	return o.VrfVlanRanges
}

// GetVrfVlanRangesOk returns a tuple with the VrfVlanRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthernetEvpnVxlanL3Fabric) GetVrfVlanRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.VrfVlanRanges) {
		return nil, false
	}
	return o.VrfVlanRanges, true
}

// HasVrfVlanRanges returns a boolean if a field has been set.
func (o *EthernetEvpnVxlanL3Fabric) HasVrfVlanRanges() bool {
	if o != nil && !IsNil(o.VrfVlanRanges) {
		return true
	}

	return false
}

// SetVrfVlanRanges gets a reference to the given []string and assigns it to the VrfVlanRanges field.
func (o *EthernetEvpnVxlanL3Fabric) SetVrfVlanRanges(v []string) {
	o.VrfVlanRanges = v
}

func (o EthernetEvpnVxlanL3Fabric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthernetEvpnVxlanL3Fabric) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.VniPrefix) {
		toSerialize["vniPrefix"] = o.VniPrefix
	}
	if !IsNil(o.PreventVrfCleanup) {
		toSerialize["preventVrfCleanup"] = o.PreventVrfCleanup
	}
	if !IsNil(o.ReservedVrfs) {
		toSerialize["reservedVrfs"] = o.ReservedVrfs
	}
	if !IsNil(o.VrfVlanRanges) {
		toSerialize["vrfVlanRanges"] = o.VrfVlanRanges
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EthernetEvpnVxlanL3Fabric) UnmarshalJSON(data []byte) (err error) {
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

	varEthernetEvpnVxlanL3Fabric := _EthernetEvpnVxlanL3Fabric{}

	err = json.Unmarshal(data, &varEthernetEvpnVxlanL3Fabric)

	if err != nil {
		return err
	}

	*o = EthernetEvpnVxlanL3Fabric(varEthernetEvpnVxlanL3Fabric)

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
		delete(additionalProperties, "vniPrefix")
		delete(additionalProperties, "preventVrfCleanup")
		delete(additionalProperties, "reservedVrfs")
		delete(additionalProperties, "vrfVlanRanges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEthernetEvpnVxlanL3Fabric struct {
	value *EthernetEvpnVxlanL3Fabric
	isSet bool
}

func (v NullableEthernetEvpnVxlanL3Fabric) Get() *EthernetEvpnVxlanL3Fabric {
	return v.value
}

func (v *NullableEthernetEvpnVxlanL3Fabric) Set(val *EthernetEvpnVxlanL3Fabric) {
	v.value = val
	v.isSet = true
}

func (v NullableEthernetEvpnVxlanL3Fabric) IsSet() bool {
	return v.isSet
}

func (v *NullableEthernetEvpnVxlanL3Fabric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthernetEvpnVxlanL3Fabric(val *EthernetEvpnVxlanL3Fabric) *NullableEthernetEvpnVxlanL3Fabric {
	return &NullableEthernetEvpnVxlanL3Fabric{value: val, isSet: true}
}

func (v NullableEthernetEvpnVxlanL3Fabric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthernetEvpnVxlanL3Fabric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ServerInstanceGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceGroupUpdate{}

// ServerInstanceGroupUpdate struct for ServerInstanceGroupUpdate
type ServerInstanceGroupUpdate struct {
	// The server instance group label. Will be automatically generated if not provided.
	Label *string `json:"label,omitempty"`
	ServerGroupName *string `json:"serverGroupName,omitempty"`
	// The number of instances to be created on the InstanceArray.
	InstanceCount *int32 `json:"instanceCount,omitempty"`
	// Automatically allocate IP addresses to child Instance`s InstanceInterface elements.
	IpAllocateAuto *int32 `json:"ipAllocateAuto,omitempty"`
	// Automatically create or expand Subnet elements until the necessary IPv4 addresses are allocated.
	Ipv4SubnetCreateAuto *int32 `json:"ipv4SubnetCreateAuto,omitempty"`
	// The volume template ID (or name) to use if the servers in the InstanceArray have local disks.
	VolumeTemplateId *int32 `json:"volumeTemplateId,omitempty"`
	// Id of the bootable drive for the Server Instance Group.
	DriveArrayIdBoot *int32 `json:"driveArrayIdBoot,omitempty"`
	// Object containing custom variables and variable overrides.
	CustomVariables map[string]interface{} `json:"customVariables,omitempty"`
	// The CPU count on each instance.
	ProcessorCount *int32 `json:"processorCount,omitempty"`
	// The minimum cores of a CPU.
	ProcessorCoreCount *int32 `json:"processorCoreCount,omitempty"`
	// The minimum clock speed of a CPU.
	ProcessorCoreMhz *int32 `json:"processorCoreMhz,omitempty"`
	// The minimum RAM capacity of each instance.
	RamGbytes *int32 `json:"ramGbytes,omitempty"`
	// The minimum number of physical disks.
	DiskCount *int32 `json:"diskCount,omitempty"`
	// The minimum size of a single disk.
	DiskSizeMbytes *int32 `json:"diskSizeMbytes,omitempty"`
	// The types of physical disks.
	DiskTypes []string `json:"diskTypes,omitempty"`
	// Enable virtual interfaces
	VirtualInterfacesEnabled *int32 `json:"virtualInterfacesEnabled,omitempty"`
	// Contains info about additional ips to be assigned to the WAN interfaces.
	AdditionalWanIpv4Json map[string]interface{} `json:"additionalWanIpv4Json,omitempty"`
	// The ipv4 vlan that should override the default from the WAN Network for the primary ip.
	OverrideIpv4WanVlanId *int32 `json:"overrideIpv4WanVlanId,omitempty"`
	// ID of a ipv4 WAN subnet-pool from which to force the subnet allocation for the InstanceInterfaces associated with this InstanceArray.
	NetworkEquipmentForceSubnetPoolIpv4WanId *int32 `json:"networkEquipmentForceSubnetPoolIpv4WanId,omitempty"`
	// The group's default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied.
	DefaultServerProfileID *int32 `json:"defaultServerProfileID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceGroupUpdate ServerInstanceGroupUpdate

// NewServerInstanceGroupUpdate instantiates a new ServerInstanceGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceGroupUpdate() *ServerInstanceGroupUpdate {
	this := ServerInstanceGroupUpdate{}
	var instanceCount int32 = 1
	this.InstanceCount = &instanceCount
	var ipAllocateAuto int32 = 1
	this.IpAllocateAuto = &ipAllocateAuto
	var ipv4SubnetCreateAuto int32 = 1
	this.Ipv4SubnetCreateAuto = &ipv4SubnetCreateAuto
	var processorCount int32 = 1
	this.ProcessorCount = &processorCount
	var processorCoreCount int32 = 1
	this.ProcessorCoreCount = &processorCoreCount
	var processorCoreMhz int32 = 1000
	this.ProcessorCoreMhz = &processorCoreMhz
	var ramGbytes int32 = 1
	this.RamGbytes = &ramGbytes
	var diskCount int32 = 0
	this.DiskCount = &diskCount
	var diskSizeMbytes int32 = 0
	this.DiskSizeMbytes = &diskSizeMbytes
	var virtualInterfacesEnabled int32 = 0
	this.VirtualInterfacesEnabled = &virtualInterfacesEnabled
	return &this
}

// NewServerInstanceGroupUpdateWithDefaults instantiates a new ServerInstanceGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceGroupUpdateWithDefaults() *ServerInstanceGroupUpdate {
	this := ServerInstanceGroupUpdate{}
	var instanceCount int32 = 1
	this.InstanceCount = &instanceCount
	var ipAllocateAuto int32 = 1
	this.IpAllocateAuto = &ipAllocateAuto
	var ipv4SubnetCreateAuto int32 = 1
	this.Ipv4SubnetCreateAuto = &ipv4SubnetCreateAuto
	var processorCount int32 = 1
	this.ProcessorCount = &processorCount
	var processorCoreCount int32 = 1
	this.ProcessorCoreCount = &processorCoreCount
	var processorCoreMhz int32 = 1000
	this.ProcessorCoreMhz = &processorCoreMhz
	var ramGbytes int32 = 1
	this.RamGbytes = &ramGbytes
	var diskCount int32 = 0
	this.DiskCount = &diskCount
	var diskSizeMbytes int32 = 0
	this.DiskSizeMbytes = &diskSizeMbytes
	var virtualInterfacesEnabled int32 = 0
	this.VirtualInterfacesEnabled = &virtualInterfacesEnabled
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServerInstanceGroupUpdate) SetLabel(v string) {
	o.Label = &v
}

// GetServerGroupName returns the ServerGroupName field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetServerGroupName() string {
	if o == nil || IsNil(o.ServerGroupName) {
		var ret string
		return ret
	}
	return *o.ServerGroupName
}

// GetServerGroupNameOk returns a tuple with the ServerGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetServerGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerGroupName) {
		return nil, false
	}
	return o.ServerGroupName, true
}

// HasServerGroupName returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasServerGroupName() bool {
	if o != nil && !IsNil(o.ServerGroupName) {
		return true
	}

	return false
}

// SetServerGroupName gets a reference to the given string and assigns it to the ServerGroupName field.
func (o *ServerInstanceGroupUpdate) SetServerGroupName(v string) {
	o.ServerGroupName = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetInstanceCount() int32 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetInstanceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *ServerInstanceGroupUpdate) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

// GetIpAllocateAuto returns the IpAllocateAuto field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetIpAllocateAuto() int32 {
	if o == nil || IsNil(o.IpAllocateAuto) {
		var ret int32
		return ret
	}
	return *o.IpAllocateAuto
}

// GetIpAllocateAutoOk returns a tuple with the IpAllocateAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetIpAllocateAutoOk() (*int32, bool) {
	if o == nil || IsNil(o.IpAllocateAuto) {
		return nil, false
	}
	return o.IpAllocateAuto, true
}

// HasIpAllocateAuto returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasIpAllocateAuto() bool {
	if o != nil && !IsNil(o.IpAllocateAuto) {
		return true
	}

	return false
}

// SetIpAllocateAuto gets a reference to the given int32 and assigns it to the IpAllocateAuto field.
func (o *ServerInstanceGroupUpdate) SetIpAllocateAuto(v int32) {
	o.IpAllocateAuto = &v
}

// GetIpv4SubnetCreateAuto returns the Ipv4SubnetCreateAuto field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetIpv4SubnetCreateAuto() int32 {
	if o == nil || IsNil(o.Ipv4SubnetCreateAuto) {
		var ret int32
		return ret
	}
	return *o.Ipv4SubnetCreateAuto
}

// GetIpv4SubnetCreateAutoOk returns a tuple with the Ipv4SubnetCreateAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetIpv4SubnetCreateAutoOk() (*int32, bool) {
	if o == nil || IsNil(o.Ipv4SubnetCreateAuto) {
		return nil, false
	}
	return o.Ipv4SubnetCreateAuto, true
}

// HasIpv4SubnetCreateAuto returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasIpv4SubnetCreateAuto() bool {
	if o != nil && !IsNil(o.Ipv4SubnetCreateAuto) {
		return true
	}

	return false
}

// SetIpv4SubnetCreateAuto gets a reference to the given int32 and assigns it to the Ipv4SubnetCreateAuto field.
func (o *ServerInstanceGroupUpdate) SetIpv4SubnetCreateAuto(v int32) {
	o.Ipv4SubnetCreateAuto = &v
}

// GetVolumeTemplateId returns the VolumeTemplateId field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetVolumeTemplateId() int32 {
	if o == nil || IsNil(o.VolumeTemplateId) {
		var ret int32
		return ret
	}
	return *o.VolumeTemplateId
}

// GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetVolumeTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VolumeTemplateId) {
		return nil, false
	}
	return o.VolumeTemplateId, true
}

// HasVolumeTemplateId returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasVolumeTemplateId() bool {
	if o != nil && !IsNil(o.VolumeTemplateId) {
		return true
	}

	return false
}

// SetVolumeTemplateId gets a reference to the given int32 and assigns it to the VolumeTemplateId field.
func (o *ServerInstanceGroupUpdate) SetVolumeTemplateId(v int32) {
	o.VolumeTemplateId = &v
}

// GetDriveArrayIdBoot returns the DriveArrayIdBoot field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetDriveArrayIdBoot() int32 {
	if o == nil || IsNil(o.DriveArrayIdBoot) {
		var ret int32
		return ret
	}
	return *o.DriveArrayIdBoot
}

// GetDriveArrayIdBootOk returns a tuple with the DriveArrayIdBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetDriveArrayIdBootOk() (*int32, bool) {
	if o == nil || IsNil(o.DriveArrayIdBoot) {
		return nil, false
	}
	return o.DriveArrayIdBoot, true
}

// HasDriveArrayIdBoot returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasDriveArrayIdBoot() bool {
	if o != nil && !IsNil(o.DriveArrayIdBoot) {
		return true
	}

	return false
}

// SetDriveArrayIdBoot gets a reference to the given int32 and assigns it to the DriveArrayIdBoot field.
func (o *ServerInstanceGroupUpdate) SetDriveArrayIdBoot(v int32) {
	o.DriveArrayIdBoot = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetCustomVariables() map[string]interface{} {
	if o == nil || IsNil(o.CustomVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetCustomVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return map[string]interface{}{}, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given map[string]interface{} and assigns it to the CustomVariables field.
func (o *ServerInstanceGroupUpdate) SetCustomVariables(v map[string]interface{}) {
	o.CustomVariables = v
}

// GetProcessorCount returns the ProcessorCount field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetProcessorCount() int32 {
	if o == nil || IsNil(o.ProcessorCount) {
		var ret int32
		return ret
	}
	return *o.ProcessorCount
}

// GetProcessorCountOk returns a tuple with the ProcessorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetProcessorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessorCount) {
		return nil, false
	}
	return o.ProcessorCount, true
}

// HasProcessorCount returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasProcessorCount() bool {
	if o != nil && !IsNil(o.ProcessorCount) {
		return true
	}

	return false
}

// SetProcessorCount gets a reference to the given int32 and assigns it to the ProcessorCount field.
func (o *ServerInstanceGroupUpdate) SetProcessorCount(v int32) {
	o.ProcessorCount = &v
}

// GetProcessorCoreCount returns the ProcessorCoreCount field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetProcessorCoreCount() int32 {
	if o == nil || IsNil(o.ProcessorCoreCount) {
		var ret int32
		return ret
	}
	return *o.ProcessorCoreCount
}

// GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetProcessorCoreCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessorCoreCount) {
		return nil, false
	}
	return o.ProcessorCoreCount, true
}

// HasProcessorCoreCount returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasProcessorCoreCount() bool {
	if o != nil && !IsNil(o.ProcessorCoreCount) {
		return true
	}

	return false
}

// SetProcessorCoreCount gets a reference to the given int32 and assigns it to the ProcessorCoreCount field.
func (o *ServerInstanceGroupUpdate) SetProcessorCoreCount(v int32) {
	o.ProcessorCoreCount = &v
}

// GetProcessorCoreMhz returns the ProcessorCoreMhz field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetProcessorCoreMhz() int32 {
	if o == nil || IsNil(o.ProcessorCoreMhz) {
		var ret int32
		return ret
	}
	return *o.ProcessorCoreMhz
}

// GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetProcessorCoreMhzOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessorCoreMhz) {
		return nil, false
	}
	return o.ProcessorCoreMhz, true
}

// HasProcessorCoreMhz returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasProcessorCoreMhz() bool {
	if o != nil && !IsNil(o.ProcessorCoreMhz) {
		return true
	}

	return false
}

// SetProcessorCoreMhz gets a reference to the given int32 and assigns it to the ProcessorCoreMhz field.
func (o *ServerInstanceGroupUpdate) SetProcessorCoreMhz(v int32) {
	o.ProcessorCoreMhz = &v
}

// GetRamGbytes returns the RamGbytes field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetRamGbytes() int32 {
	if o == nil || IsNil(o.RamGbytes) {
		var ret int32
		return ret
	}
	return *o.RamGbytes
}

// GetRamGbytesOk returns a tuple with the RamGbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetRamGbytesOk() (*int32, bool) {
	if o == nil || IsNil(o.RamGbytes) {
		return nil, false
	}
	return o.RamGbytes, true
}

// HasRamGbytes returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasRamGbytes() bool {
	if o != nil && !IsNil(o.RamGbytes) {
		return true
	}

	return false
}

// SetRamGbytes gets a reference to the given int32 and assigns it to the RamGbytes field.
func (o *ServerInstanceGroupUpdate) SetRamGbytes(v int32) {
	o.RamGbytes = &v
}

// GetDiskCount returns the DiskCount field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetDiskCount() int32 {
	if o == nil || IsNil(o.DiskCount) {
		var ret int32
		return ret
	}
	return *o.DiskCount
}

// GetDiskCountOk returns a tuple with the DiskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetDiskCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskCount) {
		return nil, false
	}
	return o.DiskCount, true
}

// HasDiskCount returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasDiskCount() bool {
	if o != nil && !IsNil(o.DiskCount) {
		return true
	}

	return false
}

// SetDiskCount gets a reference to the given int32 and assigns it to the DiskCount field.
func (o *ServerInstanceGroupUpdate) SetDiskCount(v int32) {
	o.DiskCount = &v
}

// GetDiskSizeMbytes returns the DiskSizeMbytes field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetDiskSizeMbytes() int32 {
	if o == nil || IsNil(o.DiskSizeMbytes) {
		var ret int32
		return ret
	}
	return *o.DiskSizeMbytes
}

// GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetDiskSizeMbytesOk() (*int32, bool) {
	if o == nil || IsNil(o.DiskSizeMbytes) {
		return nil, false
	}
	return o.DiskSizeMbytes, true
}

// HasDiskSizeMbytes returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasDiskSizeMbytes() bool {
	if o != nil && !IsNil(o.DiskSizeMbytes) {
		return true
	}

	return false
}

// SetDiskSizeMbytes gets a reference to the given int32 and assigns it to the DiskSizeMbytes field.
func (o *ServerInstanceGroupUpdate) SetDiskSizeMbytes(v int32) {
	o.DiskSizeMbytes = &v
}

// GetDiskTypes returns the DiskTypes field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetDiskTypes() []string {
	if o == nil || IsNil(o.DiskTypes) {
		var ret []string
		return ret
	}
	return o.DiskTypes
}

// GetDiskTypesOk returns a tuple with the DiskTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetDiskTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DiskTypes) {
		return nil, false
	}
	return o.DiskTypes, true
}

// HasDiskTypes returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasDiskTypes() bool {
	if o != nil && !IsNil(o.DiskTypes) {
		return true
	}

	return false
}

// SetDiskTypes gets a reference to the given []string and assigns it to the DiskTypes field.
func (o *ServerInstanceGroupUpdate) SetDiskTypes(v []string) {
	o.DiskTypes = v
}

// GetVirtualInterfacesEnabled returns the VirtualInterfacesEnabled field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetVirtualInterfacesEnabled() int32 {
	if o == nil || IsNil(o.VirtualInterfacesEnabled) {
		var ret int32
		return ret
	}
	return *o.VirtualInterfacesEnabled
}

// GetVirtualInterfacesEnabledOk returns a tuple with the VirtualInterfacesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetVirtualInterfacesEnabledOk() (*int32, bool) {
	if o == nil || IsNil(o.VirtualInterfacesEnabled) {
		return nil, false
	}
	return o.VirtualInterfacesEnabled, true
}

// HasVirtualInterfacesEnabled returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasVirtualInterfacesEnabled() bool {
	if o != nil && !IsNil(o.VirtualInterfacesEnabled) {
		return true
	}

	return false
}

// SetVirtualInterfacesEnabled gets a reference to the given int32 and assigns it to the VirtualInterfacesEnabled field.
func (o *ServerInstanceGroupUpdate) SetVirtualInterfacesEnabled(v int32) {
	o.VirtualInterfacesEnabled = &v
}

// GetAdditionalWanIpv4Json returns the AdditionalWanIpv4Json field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetAdditionalWanIpv4Json() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalWanIpv4Json) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalWanIpv4Json
}

// GetAdditionalWanIpv4JsonOk returns a tuple with the AdditionalWanIpv4Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetAdditionalWanIpv4JsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalWanIpv4Json) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalWanIpv4Json, true
}

// HasAdditionalWanIpv4Json returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasAdditionalWanIpv4Json() bool {
	if o != nil && !IsNil(o.AdditionalWanIpv4Json) {
		return true
	}

	return false
}

// SetAdditionalWanIpv4Json gets a reference to the given map[string]interface{} and assigns it to the AdditionalWanIpv4Json field.
func (o *ServerInstanceGroupUpdate) SetAdditionalWanIpv4Json(v map[string]interface{}) {
	o.AdditionalWanIpv4Json = v
}

// GetOverrideIpv4WanVlanId returns the OverrideIpv4WanVlanId field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetOverrideIpv4WanVlanId() int32 {
	if o == nil || IsNil(o.OverrideIpv4WanVlanId) {
		var ret int32
		return ret
	}
	return *o.OverrideIpv4WanVlanId
}

// GetOverrideIpv4WanVlanIdOk returns a tuple with the OverrideIpv4WanVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetOverrideIpv4WanVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OverrideIpv4WanVlanId) {
		return nil, false
	}
	return o.OverrideIpv4WanVlanId, true
}

// HasOverrideIpv4WanVlanId returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasOverrideIpv4WanVlanId() bool {
	if o != nil && !IsNil(o.OverrideIpv4WanVlanId) {
		return true
	}

	return false
}

// SetOverrideIpv4WanVlanId gets a reference to the given int32 and assigns it to the OverrideIpv4WanVlanId field.
func (o *ServerInstanceGroupUpdate) SetOverrideIpv4WanVlanId(v int32) {
	o.OverrideIpv4WanVlanId = &v
}

// GetNetworkEquipmentForceSubnetPoolIpv4WanId returns the NetworkEquipmentForceSubnetPoolIpv4WanId field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetNetworkEquipmentForceSubnetPoolIpv4WanId() int32 {
	if o == nil || IsNil(o.NetworkEquipmentForceSubnetPoolIpv4WanId) {
		var ret int32
		return ret
	}
	return *o.NetworkEquipmentForceSubnetPoolIpv4WanId
}

// GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk returns a tuple with the NetworkEquipmentForceSubnetPoolIpv4WanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkEquipmentForceSubnetPoolIpv4WanId) {
		return nil, false
	}
	return o.NetworkEquipmentForceSubnetPoolIpv4WanId, true
}

// HasNetworkEquipmentForceSubnetPoolIpv4WanId returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasNetworkEquipmentForceSubnetPoolIpv4WanId() bool {
	if o != nil && !IsNil(o.NetworkEquipmentForceSubnetPoolIpv4WanId) {
		return true
	}

	return false
}

// SetNetworkEquipmentForceSubnetPoolIpv4WanId gets a reference to the given int32 and assigns it to the NetworkEquipmentForceSubnetPoolIpv4WanId field.
func (o *ServerInstanceGroupUpdate) SetNetworkEquipmentForceSubnetPoolIpv4WanId(v int32) {
	o.NetworkEquipmentForceSubnetPoolIpv4WanId = &v
}

// GetDefaultServerProfileID returns the DefaultServerProfileID field value if set, zero value otherwise.
func (o *ServerInstanceGroupUpdate) GetDefaultServerProfileID() int32 {
	if o == nil || IsNil(o.DefaultServerProfileID) {
		var ret int32
		return ret
	}
	return *o.DefaultServerProfileID
}

// GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupUpdate) GetDefaultServerProfileIDOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultServerProfileID) {
		return nil, false
	}
	return o.DefaultServerProfileID, true
}

// HasDefaultServerProfileID returns a boolean if a field has been set.
func (o *ServerInstanceGroupUpdate) HasDefaultServerProfileID() bool {
	if o != nil && !IsNil(o.DefaultServerProfileID) {
		return true
	}

	return false
}

// SetDefaultServerProfileID gets a reference to the given int32 and assigns it to the DefaultServerProfileID field.
func (o *ServerInstanceGroupUpdate) SetDefaultServerProfileID(v int32) {
	o.DefaultServerProfileID = &v
}

func (o ServerInstanceGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceGroupUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ServerGroupName) {
		toSerialize["serverGroupName"] = o.ServerGroupName
	}
	if !IsNil(o.InstanceCount) {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if !IsNil(o.IpAllocateAuto) {
		toSerialize["ipAllocateAuto"] = o.IpAllocateAuto
	}
	if !IsNil(o.Ipv4SubnetCreateAuto) {
		toSerialize["ipv4SubnetCreateAuto"] = o.Ipv4SubnetCreateAuto
	}
	if !IsNil(o.VolumeTemplateId) {
		toSerialize["volumeTemplateId"] = o.VolumeTemplateId
	}
	if !IsNil(o.DriveArrayIdBoot) {
		toSerialize["driveArrayIdBoot"] = o.DriveArrayIdBoot
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	if !IsNil(o.ProcessorCount) {
		toSerialize["processorCount"] = o.ProcessorCount
	}
	if !IsNil(o.ProcessorCoreCount) {
		toSerialize["processorCoreCount"] = o.ProcessorCoreCount
	}
	if !IsNil(o.ProcessorCoreMhz) {
		toSerialize["processorCoreMhz"] = o.ProcessorCoreMhz
	}
	if !IsNil(o.RamGbytes) {
		toSerialize["ramGbytes"] = o.RamGbytes
	}
	if !IsNil(o.DiskCount) {
		toSerialize["diskCount"] = o.DiskCount
	}
	if !IsNil(o.DiskSizeMbytes) {
		toSerialize["diskSizeMbytes"] = o.DiskSizeMbytes
	}
	if !IsNil(o.DiskTypes) {
		toSerialize["diskTypes"] = o.DiskTypes
	}
	if !IsNil(o.VirtualInterfacesEnabled) {
		toSerialize["virtualInterfacesEnabled"] = o.VirtualInterfacesEnabled
	}
	if !IsNil(o.AdditionalWanIpv4Json) {
		toSerialize["additionalWanIpv4Json"] = o.AdditionalWanIpv4Json
	}
	if !IsNil(o.OverrideIpv4WanVlanId) {
		toSerialize["overrideIpv4WanVlanId"] = o.OverrideIpv4WanVlanId
	}
	if !IsNil(o.NetworkEquipmentForceSubnetPoolIpv4WanId) {
		toSerialize["networkEquipmentForceSubnetPoolIpv4WanId"] = o.NetworkEquipmentForceSubnetPoolIpv4WanId
	}
	if !IsNil(o.DefaultServerProfileID) {
		toSerialize["defaultServerProfileID"] = o.DefaultServerProfileID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceGroupUpdate) UnmarshalJSON(data []byte) (err error) {
	varServerInstanceGroupUpdate := _ServerInstanceGroupUpdate{}

	err = json.Unmarshal(data, &varServerInstanceGroupUpdate)

	if err != nil {
		return err
	}

	*o = ServerInstanceGroupUpdate(varServerInstanceGroupUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "serverGroupName")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "ipAllocateAuto")
		delete(additionalProperties, "ipv4SubnetCreateAuto")
		delete(additionalProperties, "volumeTemplateId")
		delete(additionalProperties, "driveArrayIdBoot")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "processorCount")
		delete(additionalProperties, "processorCoreCount")
		delete(additionalProperties, "processorCoreMhz")
		delete(additionalProperties, "ramGbytes")
		delete(additionalProperties, "diskCount")
		delete(additionalProperties, "diskSizeMbytes")
		delete(additionalProperties, "diskTypes")
		delete(additionalProperties, "virtualInterfacesEnabled")
		delete(additionalProperties, "additionalWanIpv4Json")
		delete(additionalProperties, "overrideIpv4WanVlanId")
		delete(additionalProperties, "networkEquipmentForceSubnetPoolIpv4WanId")
		delete(additionalProperties, "defaultServerProfileID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceGroupUpdate struct {
	value *ServerInstanceGroupUpdate
	isSet bool
}

func (v NullableServerInstanceGroupUpdate) Get() *ServerInstanceGroupUpdate {
	return v.value
}

func (v *NullableServerInstanceGroupUpdate) Set(val *ServerInstanceGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceGroupUpdate(val *ServerInstanceGroupUpdate) *NullableServerInstanceGroupUpdate {
	return &NullableServerInstanceGroupUpdate{value: val, isSet: true}
}

func (v NullableServerInstanceGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



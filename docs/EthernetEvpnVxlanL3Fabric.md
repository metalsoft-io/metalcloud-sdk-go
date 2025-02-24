# EthernetEvpnVxlanL3Fabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultNetworkProfileId** | **int32** | Unique identifier for the default network profile. Must be a positive integer (minimum: 1) corresponding to an existing profile. | 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 
**ZeroTouchEnabled** | Pointer to **bool** | Enables zero-touch provisioning for automatic device configuration. | [optional] 
**AllocateDefaultVlan** | Pointer to **bool** | Indicates whether to automatically allocate a default VLAN. | [optional] 
**AsnRanges** | Pointer to **[]string** | ASN ranges in the format \&quot;start-end\&quot;, where each range is an ordered pair with values between 1 and 4294967295. | [optional] 
**DefaultVlan** | Pointer to **int32** | Default VLAN ID. Must be a number between 1 and 4096. | [optional] 
**ExtraInternalIPsPerSubnet** | Pointer to **int32** | Extra internal IPs allocated per subnet; valid range is between 1 and 1000. | [optional] 
**LagRanges** | Pointer to **[]string** | Link Aggregation (LAG) ranges in the format \&quot;start-end\&quot;; each range must be within the bounds of 1 to 4096. | [optional] 
**LeafSwitchesHaveMlagPairs** | Pointer to **bool** | Indicates if leaf switches have MLAG pairs. | [optional] 
**MlagRanges** | Pointer to **[]string** | MLAG ID ranges. Each range must be provided in \&quot;start-end\&quot; format with values between 1 and 4096. | [optional] 
**NumberOfSpinesNextToLeafSwitches** | **int32** | Number of spines adjacent to leaf switches. Must be a positive number. | 
**PreventVlanCleanup** | Pointer to **[]string** | VLAN ranges that should be prevented from automatic cleanup. Format must be \&quot;start-end\&quot;. | [optional] 
**PreventCleanupFromUplinks** | Pointer to **bool** | Flag indicating whether cleanup from uplink interfaces should be prevented. | [optional] 
**ReservedVlans** | Pointer to **[]string** | Reserved VLAN ranges that are excluded from general allocation. Must follow the \&quot;start-end\&quot; format. | [optional] 
**VlanRanges** | Pointer to **[]string** | Array of VLAN range strings in \&quot;start-end\&quot; format to be used in configuration. | [optional] 
**VniPrefix** | Pointer to **int32** | The VNI prefix for the EVPN VXLAN fabric. | [optional] 
**PreventVrfCleanup** | Pointer to **[]string** | VRF ID ranges to be preserved from automatic cleanup. Each range must follow the \&quot;start-end\&quot; format. | [optional] 
**ReservedVrfs** | Pointer to **[]string** | Reserved VRF ID ranges that are set aside exclusively for specific network functions. Each range must be provided in the \&quot;start-end\&quot; format. | [optional] 
**VrfVlanRanges** | Pointer to **[]string** | VLAN ranges to be associated with VRF instances. Each value must be an ordered pair specified in the \&quot;start-end\&quot; format. | [optional] 

## Methods

### NewEthernetEvpnVxlanL3Fabric

`func NewEthernetEvpnVxlanL3Fabric(defaultNetworkProfileId int32, numberOfSpinesNextToLeafSwitches int32, ) *EthernetEvpnVxlanL3Fabric`

NewEthernetEvpnVxlanL3Fabric instantiates a new EthernetEvpnVxlanL3Fabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetEvpnVxlanL3FabricWithDefaults

`func NewEthernetEvpnVxlanL3FabricWithDefaults() *EthernetEvpnVxlanL3Fabric`

NewEthernetEvpnVxlanL3FabricWithDefaults instantiates a new EthernetEvpnVxlanL3Fabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultNetworkProfileId

`func (o *EthernetEvpnVxlanL3Fabric) GetDefaultNetworkProfileId() int32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *EthernetEvpnVxlanL3Fabric) GetDefaultNetworkProfileIdOk() (*int32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *EthernetEvpnVxlanL3Fabric) SetDefaultNetworkProfileId(v int32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.


### GetGnmiMonitoringEnabled

`func (o *EthernetEvpnVxlanL3Fabric) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *EthernetEvpnVxlanL3Fabric) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *EthernetEvpnVxlanL3Fabric) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *EthernetEvpnVxlanL3Fabric) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *EthernetEvpnVxlanL3Fabric) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *EthernetEvpnVxlanL3Fabric) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *EthernetEvpnVxlanL3Fabric) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *EthernetEvpnVxlanL3Fabric) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetZeroTouchEnabled

`func (o *EthernetEvpnVxlanL3Fabric) GetZeroTouchEnabled() bool`

GetZeroTouchEnabled returns the ZeroTouchEnabled field if non-nil, zero value otherwise.

### GetZeroTouchEnabledOk

`func (o *EthernetEvpnVxlanL3Fabric) GetZeroTouchEnabledOk() (*bool, bool)`

GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchEnabled

`func (o *EthernetEvpnVxlanL3Fabric) SetZeroTouchEnabled(v bool)`

SetZeroTouchEnabled sets ZeroTouchEnabled field to given value.

### HasZeroTouchEnabled

`func (o *EthernetEvpnVxlanL3Fabric) HasZeroTouchEnabled() bool`

HasZeroTouchEnabled returns a boolean if a field has been set.

### GetAllocateDefaultVlan

`func (o *EthernetEvpnVxlanL3Fabric) GetAllocateDefaultVlan() bool`

GetAllocateDefaultVlan returns the AllocateDefaultVlan field if non-nil, zero value otherwise.

### GetAllocateDefaultVlanOk

`func (o *EthernetEvpnVxlanL3Fabric) GetAllocateDefaultVlanOk() (*bool, bool)`

GetAllocateDefaultVlanOk returns a tuple with the AllocateDefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateDefaultVlan

`func (o *EthernetEvpnVxlanL3Fabric) SetAllocateDefaultVlan(v bool)`

SetAllocateDefaultVlan sets AllocateDefaultVlan field to given value.

### HasAllocateDefaultVlan

`func (o *EthernetEvpnVxlanL3Fabric) HasAllocateDefaultVlan() bool`

HasAllocateDefaultVlan returns a boolean if a field has been set.

### GetAsnRanges

`func (o *EthernetEvpnVxlanL3Fabric) GetAsnRanges() []string`

GetAsnRanges returns the AsnRanges field if non-nil, zero value otherwise.

### GetAsnRangesOk

`func (o *EthernetEvpnVxlanL3Fabric) GetAsnRangesOk() (*[]string, bool)`

GetAsnRangesOk returns a tuple with the AsnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnRanges

`func (o *EthernetEvpnVxlanL3Fabric) SetAsnRanges(v []string)`

SetAsnRanges sets AsnRanges field to given value.

### HasAsnRanges

`func (o *EthernetEvpnVxlanL3Fabric) HasAsnRanges() bool`

HasAsnRanges returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *EthernetEvpnVxlanL3Fabric) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *EthernetEvpnVxlanL3Fabric) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *EthernetEvpnVxlanL3Fabric) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *EthernetEvpnVxlanL3Fabric) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetExtraInternalIPsPerSubnet

`func (o *EthernetEvpnVxlanL3Fabric) GetExtraInternalIPsPerSubnet() int32`

GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field if non-nil, zero value otherwise.

### GetExtraInternalIPsPerSubnetOk

`func (o *EthernetEvpnVxlanL3Fabric) GetExtraInternalIPsPerSubnetOk() (*int32, bool)`

GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInternalIPsPerSubnet

`func (o *EthernetEvpnVxlanL3Fabric) SetExtraInternalIPsPerSubnet(v int32)`

SetExtraInternalIPsPerSubnet sets ExtraInternalIPsPerSubnet field to given value.

### HasExtraInternalIPsPerSubnet

`func (o *EthernetEvpnVxlanL3Fabric) HasExtraInternalIPsPerSubnet() bool`

HasExtraInternalIPsPerSubnet returns a boolean if a field has been set.

### GetLagRanges

`func (o *EthernetEvpnVxlanL3Fabric) GetLagRanges() []string`

GetLagRanges returns the LagRanges field if non-nil, zero value otherwise.

### GetLagRangesOk

`func (o *EthernetEvpnVxlanL3Fabric) GetLagRangesOk() (*[]string, bool)`

GetLagRangesOk returns a tuple with the LagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagRanges

`func (o *EthernetEvpnVxlanL3Fabric) SetLagRanges(v []string)`

SetLagRanges sets LagRanges field to given value.

### HasLagRanges

`func (o *EthernetEvpnVxlanL3Fabric) HasLagRanges() bool`

HasLagRanges returns a boolean if a field has been set.

### GetLeafSwitchesHaveMlagPairs

`func (o *EthernetEvpnVxlanL3Fabric) GetLeafSwitchesHaveMlagPairs() bool`

GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field if non-nil, zero value otherwise.

### GetLeafSwitchesHaveMlagPairsOk

`func (o *EthernetEvpnVxlanL3Fabric) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool)`

GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafSwitchesHaveMlagPairs

`func (o *EthernetEvpnVxlanL3Fabric) SetLeafSwitchesHaveMlagPairs(v bool)`

SetLeafSwitchesHaveMlagPairs sets LeafSwitchesHaveMlagPairs field to given value.

### HasLeafSwitchesHaveMlagPairs

`func (o *EthernetEvpnVxlanL3Fabric) HasLeafSwitchesHaveMlagPairs() bool`

HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.

### GetMlagRanges

`func (o *EthernetEvpnVxlanL3Fabric) GetMlagRanges() []string`

GetMlagRanges returns the MlagRanges field if non-nil, zero value otherwise.

### GetMlagRangesOk

`func (o *EthernetEvpnVxlanL3Fabric) GetMlagRangesOk() (*[]string, bool)`

GetMlagRangesOk returns a tuple with the MlagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagRanges

`func (o *EthernetEvpnVxlanL3Fabric) SetMlagRanges(v []string)`

SetMlagRanges sets MlagRanges field to given value.

### HasMlagRanges

`func (o *EthernetEvpnVxlanL3Fabric) HasMlagRanges() bool`

HasMlagRanges returns a boolean if a field has been set.

### GetNumberOfSpinesNextToLeafSwitches

`func (o *EthernetEvpnVxlanL3Fabric) GetNumberOfSpinesNextToLeafSwitches() int32`

GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field if non-nil, zero value otherwise.

### GetNumberOfSpinesNextToLeafSwitchesOk

`func (o *EthernetEvpnVxlanL3Fabric) GetNumberOfSpinesNextToLeafSwitchesOk() (*int32, bool)`

GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesNextToLeafSwitches

`func (o *EthernetEvpnVxlanL3Fabric) SetNumberOfSpinesNextToLeafSwitches(v int32)`

SetNumberOfSpinesNextToLeafSwitches sets NumberOfSpinesNextToLeafSwitches field to given value.


### GetPreventVlanCleanup

`func (o *EthernetEvpnVxlanL3Fabric) GetPreventVlanCleanup() []string`

GetPreventVlanCleanup returns the PreventVlanCleanup field if non-nil, zero value otherwise.

### GetPreventVlanCleanupOk

`func (o *EthernetEvpnVxlanL3Fabric) GetPreventVlanCleanupOk() (*[]string, bool)`

GetPreventVlanCleanupOk returns a tuple with the PreventVlanCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVlanCleanup

`func (o *EthernetEvpnVxlanL3Fabric) SetPreventVlanCleanup(v []string)`

SetPreventVlanCleanup sets PreventVlanCleanup field to given value.

### HasPreventVlanCleanup

`func (o *EthernetEvpnVxlanL3Fabric) HasPreventVlanCleanup() bool`

HasPreventVlanCleanup returns a boolean if a field has been set.

### GetPreventCleanupFromUplinks

`func (o *EthernetEvpnVxlanL3Fabric) GetPreventCleanupFromUplinks() bool`

GetPreventCleanupFromUplinks returns the PreventCleanupFromUplinks field if non-nil, zero value otherwise.

### GetPreventCleanupFromUplinksOk

`func (o *EthernetEvpnVxlanL3Fabric) GetPreventCleanupFromUplinksOk() (*bool, bool)`

GetPreventCleanupFromUplinksOk returns a tuple with the PreventCleanupFromUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventCleanupFromUplinks

`func (o *EthernetEvpnVxlanL3Fabric) SetPreventCleanupFromUplinks(v bool)`

SetPreventCleanupFromUplinks sets PreventCleanupFromUplinks field to given value.

### HasPreventCleanupFromUplinks

`func (o *EthernetEvpnVxlanL3Fabric) HasPreventCleanupFromUplinks() bool`

HasPreventCleanupFromUplinks returns a boolean if a field has been set.

### GetReservedVlans

`func (o *EthernetEvpnVxlanL3Fabric) GetReservedVlans() []string`

GetReservedVlans returns the ReservedVlans field if non-nil, zero value otherwise.

### GetReservedVlansOk

`func (o *EthernetEvpnVxlanL3Fabric) GetReservedVlansOk() (*[]string, bool)`

GetReservedVlansOk returns a tuple with the ReservedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVlans

`func (o *EthernetEvpnVxlanL3Fabric) SetReservedVlans(v []string)`

SetReservedVlans sets ReservedVlans field to given value.

### HasReservedVlans

`func (o *EthernetEvpnVxlanL3Fabric) HasReservedVlans() bool`

HasReservedVlans returns a boolean if a field has been set.

### GetVlanRanges

`func (o *EthernetEvpnVxlanL3Fabric) GetVlanRanges() []string`

GetVlanRanges returns the VlanRanges field if non-nil, zero value otherwise.

### GetVlanRangesOk

`func (o *EthernetEvpnVxlanL3Fabric) GetVlanRangesOk() (*[]string, bool)`

GetVlanRangesOk returns a tuple with the VlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRanges

`func (o *EthernetEvpnVxlanL3Fabric) SetVlanRanges(v []string)`

SetVlanRanges sets VlanRanges field to given value.

### HasVlanRanges

`func (o *EthernetEvpnVxlanL3Fabric) HasVlanRanges() bool`

HasVlanRanges returns a boolean if a field has been set.

### GetVniPrefix

`func (o *EthernetEvpnVxlanL3Fabric) GetVniPrefix() int32`

GetVniPrefix returns the VniPrefix field if non-nil, zero value otherwise.

### GetVniPrefixOk

`func (o *EthernetEvpnVxlanL3Fabric) GetVniPrefixOk() (*int32, bool)`

GetVniPrefixOk returns a tuple with the VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniPrefix

`func (o *EthernetEvpnVxlanL3Fabric) SetVniPrefix(v int32)`

SetVniPrefix sets VniPrefix field to given value.

### HasVniPrefix

`func (o *EthernetEvpnVxlanL3Fabric) HasVniPrefix() bool`

HasVniPrefix returns a boolean if a field has been set.

### GetPreventVrfCleanup

`func (o *EthernetEvpnVxlanL3Fabric) GetPreventVrfCleanup() []string`

GetPreventVrfCleanup returns the PreventVrfCleanup field if non-nil, zero value otherwise.

### GetPreventVrfCleanupOk

`func (o *EthernetEvpnVxlanL3Fabric) GetPreventVrfCleanupOk() (*[]string, bool)`

GetPreventVrfCleanupOk returns a tuple with the PreventVrfCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVrfCleanup

`func (o *EthernetEvpnVxlanL3Fabric) SetPreventVrfCleanup(v []string)`

SetPreventVrfCleanup sets PreventVrfCleanup field to given value.

### HasPreventVrfCleanup

`func (o *EthernetEvpnVxlanL3Fabric) HasPreventVrfCleanup() bool`

HasPreventVrfCleanup returns a boolean if a field has been set.

### GetReservedVrfs

`func (o *EthernetEvpnVxlanL3Fabric) GetReservedVrfs() []string`

GetReservedVrfs returns the ReservedVrfs field if non-nil, zero value otherwise.

### GetReservedVrfsOk

`func (o *EthernetEvpnVxlanL3Fabric) GetReservedVrfsOk() (*[]string, bool)`

GetReservedVrfsOk returns a tuple with the ReservedVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVrfs

`func (o *EthernetEvpnVxlanL3Fabric) SetReservedVrfs(v []string)`

SetReservedVrfs sets ReservedVrfs field to given value.

### HasReservedVrfs

`func (o *EthernetEvpnVxlanL3Fabric) HasReservedVrfs() bool`

HasReservedVrfs returns a boolean if a field has been set.

### GetVrfVlanRanges

`func (o *EthernetEvpnVxlanL3Fabric) GetVrfVlanRanges() []string`

GetVrfVlanRanges returns the VrfVlanRanges field if non-nil, zero value otherwise.

### GetVrfVlanRangesOk

`func (o *EthernetEvpnVxlanL3Fabric) GetVrfVlanRangesOk() (*[]string, bool)`

GetVrfVlanRangesOk returns a tuple with the VrfVlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfVlanRanges

`func (o *EthernetEvpnVxlanL3Fabric) SetVrfVlanRanges(v []string)`

SetVrfVlanRanges sets VrfVlanRanges field to given value.

### HasVrfVlanRanges

`func (o *EthernetEvpnVxlanL3Fabric) HasVrfVlanRanges() bool`

HasVrfVlanRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



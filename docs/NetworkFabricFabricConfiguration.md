# NetworkFabricFabricConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | **string** | The type of network fabric | 
**DefaultNetworkProfileId** | Pointer to **int32** | Unique identifier for the default network profile. Must be a positive integer (minimum: 1) corresponding to an existing profile. | [optional] 
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
**NumberOfSpinesNextToLeafSwitches** | Pointer to **int32** | Number of spines adjacent to leaf switches. Must be a positive number. | [optional] 
**PreventVlanCleanup** | Pointer to **[]string** | VLAN ranges that should be prevented from automatic cleanup. Format must be \&quot;start-end\&quot;. | [optional] 
**PreventCleanupFromUplinks** | Pointer to **bool** | Flag indicating whether cleanup from uplink interfaces should be prevented. | [optional] 
**ReservedVlans** | Pointer to **[]string** | Reserved VLAN ranges that are excluded from general allocation. Must follow the \&quot;start-end\&quot; format. | [optional] 
**VlanRanges** | Pointer to **[]string** | Array of VLAN range strings in \&quot;start-end\&quot; format to be used in configuration. | [optional] 
**VniPrefix** | Pointer to **int32** | The VNI prefix for the EVPN VXLAN fabric. | [optional] 
**PreventVrfCleanup** | Pointer to **[]string** | VRF ID ranges to be preserved from automatic cleanup. Each range must follow the \&quot;start-end\&quot; format. | [optional] 
**ReservedVrfs** | Pointer to **[]string** | Reserved VRF ID ranges that are set aside exclusively for specific network functions. Each range must be provided in the \&quot;start-end\&quot; format. | [optional] 
**VrfVlanRanges** | Pointer to **[]string** | VLAN ranges to be associated with VRF instances. Each value must be an ordered pair specified in the \&quot;start-end\&quot; format. | [optional] 
**VsanId** | Pointer to **int32** | VSAN ID for the Fibre Channel fabric | [optional] 
**TopologyType** | Pointer to **string** | Fabric topology type | [optional] 
**Mtu** | Pointer to **float32** | Maximum transmission unit (MTU) size in bytes | [optional] 
**ZoningConfiguration** | Pointer to **map[string]interface{}** | Zoning configuration for the fabric | [optional] 
**InteropMode** | Pointer to **string** | Interoperability mode for multi-vendor environments | [optional] 
**QosConfiguration** | Pointer to **map[string]interface{}** | Quality of Service (QoS) configuration | [optional] 
**TrunkingConfiguration** | Pointer to **map[string]interface{}** | Trunking configuration for ISLs (Inter-Switch Links) | [optional] 
**PortChannelConfiguration** | Pointer to **map[string]interface{}** | Port channel configuration for ISLs | [optional] 

## Methods

### NewNetworkFabricFabricConfiguration

`func NewNetworkFabricFabricConfiguration(fabricType string, ) *NetworkFabricFabricConfiguration`

NewNetworkFabricFabricConfiguration instantiates a new NetworkFabricFabricConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricFabricConfigurationWithDefaults

`func NewNetworkFabricFabricConfigurationWithDefaults() *NetworkFabricFabricConfiguration`

NewNetworkFabricFabricConfigurationWithDefaults instantiates a new NetworkFabricFabricConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFabricType

`func (o *NetworkFabricFabricConfiguration) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *NetworkFabricFabricConfiguration) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *NetworkFabricFabricConfiguration) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.


### GetDefaultNetworkProfileId

`func (o *NetworkFabricFabricConfiguration) GetDefaultNetworkProfileId() int32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *NetworkFabricFabricConfiguration) GetDefaultNetworkProfileIdOk() (*int32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *NetworkFabricFabricConfiguration) SetDefaultNetworkProfileId(v int32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.

### HasDefaultNetworkProfileId

`func (o *NetworkFabricFabricConfiguration) HasDefaultNetworkProfileId() bool`

HasDefaultNetworkProfileId returns a boolean if a field has been set.

### GetGnmiMonitoringEnabled

`func (o *NetworkFabricFabricConfiguration) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *NetworkFabricFabricConfiguration) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *NetworkFabricFabricConfiguration) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *NetworkFabricFabricConfiguration) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *NetworkFabricFabricConfiguration) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *NetworkFabricFabricConfiguration) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *NetworkFabricFabricConfiguration) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *NetworkFabricFabricConfiguration) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetZeroTouchEnabled

`func (o *NetworkFabricFabricConfiguration) GetZeroTouchEnabled() bool`

GetZeroTouchEnabled returns the ZeroTouchEnabled field if non-nil, zero value otherwise.

### GetZeroTouchEnabledOk

`func (o *NetworkFabricFabricConfiguration) GetZeroTouchEnabledOk() (*bool, bool)`

GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchEnabled

`func (o *NetworkFabricFabricConfiguration) SetZeroTouchEnabled(v bool)`

SetZeroTouchEnabled sets ZeroTouchEnabled field to given value.

### HasZeroTouchEnabled

`func (o *NetworkFabricFabricConfiguration) HasZeroTouchEnabled() bool`

HasZeroTouchEnabled returns a boolean if a field has been set.

### GetAllocateDefaultVlan

`func (o *NetworkFabricFabricConfiguration) GetAllocateDefaultVlan() bool`

GetAllocateDefaultVlan returns the AllocateDefaultVlan field if non-nil, zero value otherwise.

### GetAllocateDefaultVlanOk

`func (o *NetworkFabricFabricConfiguration) GetAllocateDefaultVlanOk() (*bool, bool)`

GetAllocateDefaultVlanOk returns a tuple with the AllocateDefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateDefaultVlan

`func (o *NetworkFabricFabricConfiguration) SetAllocateDefaultVlan(v bool)`

SetAllocateDefaultVlan sets AllocateDefaultVlan field to given value.

### HasAllocateDefaultVlan

`func (o *NetworkFabricFabricConfiguration) HasAllocateDefaultVlan() bool`

HasAllocateDefaultVlan returns a boolean if a field has been set.

### GetAsnRanges

`func (o *NetworkFabricFabricConfiguration) GetAsnRanges() []string`

GetAsnRanges returns the AsnRanges field if non-nil, zero value otherwise.

### GetAsnRangesOk

`func (o *NetworkFabricFabricConfiguration) GetAsnRangesOk() (*[]string, bool)`

GetAsnRangesOk returns a tuple with the AsnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnRanges

`func (o *NetworkFabricFabricConfiguration) SetAsnRanges(v []string)`

SetAsnRanges sets AsnRanges field to given value.

### HasAsnRanges

`func (o *NetworkFabricFabricConfiguration) HasAsnRanges() bool`

HasAsnRanges returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *NetworkFabricFabricConfiguration) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *NetworkFabricFabricConfiguration) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *NetworkFabricFabricConfiguration) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *NetworkFabricFabricConfiguration) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetExtraInternalIPsPerSubnet

`func (o *NetworkFabricFabricConfiguration) GetExtraInternalIPsPerSubnet() int32`

GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field if non-nil, zero value otherwise.

### GetExtraInternalIPsPerSubnetOk

`func (o *NetworkFabricFabricConfiguration) GetExtraInternalIPsPerSubnetOk() (*int32, bool)`

GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInternalIPsPerSubnet

`func (o *NetworkFabricFabricConfiguration) SetExtraInternalIPsPerSubnet(v int32)`

SetExtraInternalIPsPerSubnet sets ExtraInternalIPsPerSubnet field to given value.

### HasExtraInternalIPsPerSubnet

`func (o *NetworkFabricFabricConfiguration) HasExtraInternalIPsPerSubnet() bool`

HasExtraInternalIPsPerSubnet returns a boolean if a field has been set.

### GetLagRanges

`func (o *NetworkFabricFabricConfiguration) GetLagRanges() []string`

GetLagRanges returns the LagRanges field if non-nil, zero value otherwise.

### GetLagRangesOk

`func (o *NetworkFabricFabricConfiguration) GetLagRangesOk() (*[]string, bool)`

GetLagRangesOk returns a tuple with the LagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagRanges

`func (o *NetworkFabricFabricConfiguration) SetLagRanges(v []string)`

SetLagRanges sets LagRanges field to given value.

### HasLagRanges

`func (o *NetworkFabricFabricConfiguration) HasLagRanges() bool`

HasLagRanges returns a boolean if a field has been set.

### GetLeafSwitchesHaveMlagPairs

`func (o *NetworkFabricFabricConfiguration) GetLeafSwitchesHaveMlagPairs() bool`

GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field if non-nil, zero value otherwise.

### GetLeafSwitchesHaveMlagPairsOk

`func (o *NetworkFabricFabricConfiguration) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool)`

GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafSwitchesHaveMlagPairs

`func (o *NetworkFabricFabricConfiguration) SetLeafSwitchesHaveMlagPairs(v bool)`

SetLeafSwitchesHaveMlagPairs sets LeafSwitchesHaveMlagPairs field to given value.

### HasLeafSwitchesHaveMlagPairs

`func (o *NetworkFabricFabricConfiguration) HasLeafSwitchesHaveMlagPairs() bool`

HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.

### GetMlagRanges

`func (o *NetworkFabricFabricConfiguration) GetMlagRanges() []string`

GetMlagRanges returns the MlagRanges field if non-nil, zero value otherwise.

### GetMlagRangesOk

`func (o *NetworkFabricFabricConfiguration) GetMlagRangesOk() (*[]string, bool)`

GetMlagRangesOk returns a tuple with the MlagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagRanges

`func (o *NetworkFabricFabricConfiguration) SetMlagRanges(v []string)`

SetMlagRanges sets MlagRanges field to given value.

### HasMlagRanges

`func (o *NetworkFabricFabricConfiguration) HasMlagRanges() bool`

HasMlagRanges returns a boolean if a field has been set.

### GetNumberOfSpinesNextToLeafSwitches

`func (o *NetworkFabricFabricConfiguration) GetNumberOfSpinesNextToLeafSwitches() int32`

GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field if non-nil, zero value otherwise.

### GetNumberOfSpinesNextToLeafSwitchesOk

`func (o *NetworkFabricFabricConfiguration) GetNumberOfSpinesNextToLeafSwitchesOk() (*int32, bool)`

GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesNextToLeafSwitches

`func (o *NetworkFabricFabricConfiguration) SetNumberOfSpinesNextToLeafSwitches(v int32)`

SetNumberOfSpinesNextToLeafSwitches sets NumberOfSpinesNextToLeafSwitches field to given value.

### HasNumberOfSpinesNextToLeafSwitches

`func (o *NetworkFabricFabricConfiguration) HasNumberOfSpinesNextToLeafSwitches() bool`

HasNumberOfSpinesNextToLeafSwitches returns a boolean if a field has been set.

### GetPreventVlanCleanup

`func (o *NetworkFabricFabricConfiguration) GetPreventVlanCleanup() []string`

GetPreventVlanCleanup returns the PreventVlanCleanup field if non-nil, zero value otherwise.

### GetPreventVlanCleanupOk

`func (o *NetworkFabricFabricConfiguration) GetPreventVlanCleanupOk() (*[]string, bool)`

GetPreventVlanCleanupOk returns a tuple with the PreventVlanCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVlanCleanup

`func (o *NetworkFabricFabricConfiguration) SetPreventVlanCleanup(v []string)`

SetPreventVlanCleanup sets PreventVlanCleanup field to given value.

### HasPreventVlanCleanup

`func (o *NetworkFabricFabricConfiguration) HasPreventVlanCleanup() bool`

HasPreventVlanCleanup returns a boolean if a field has been set.

### GetPreventCleanupFromUplinks

`func (o *NetworkFabricFabricConfiguration) GetPreventCleanupFromUplinks() bool`

GetPreventCleanupFromUplinks returns the PreventCleanupFromUplinks field if non-nil, zero value otherwise.

### GetPreventCleanupFromUplinksOk

`func (o *NetworkFabricFabricConfiguration) GetPreventCleanupFromUplinksOk() (*bool, bool)`

GetPreventCleanupFromUplinksOk returns a tuple with the PreventCleanupFromUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventCleanupFromUplinks

`func (o *NetworkFabricFabricConfiguration) SetPreventCleanupFromUplinks(v bool)`

SetPreventCleanupFromUplinks sets PreventCleanupFromUplinks field to given value.

### HasPreventCleanupFromUplinks

`func (o *NetworkFabricFabricConfiguration) HasPreventCleanupFromUplinks() bool`

HasPreventCleanupFromUplinks returns a boolean if a field has been set.

### GetReservedVlans

`func (o *NetworkFabricFabricConfiguration) GetReservedVlans() []string`

GetReservedVlans returns the ReservedVlans field if non-nil, zero value otherwise.

### GetReservedVlansOk

`func (o *NetworkFabricFabricConfiguration) GetReservedVlansOk() (*[]string, bool)`

GetReservedVlansOk returns a tuple with the ReservedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVlans

`func (o *NetworkFabricFabricConfiguration) SetReservedVlans(v []string)`

SetReservedVlans sets ReservedVlans field to given value.

### HasReservedVlans

`func (o *NetworkFabricFabricConfiguration) HasReservedVlans() bool`

HasReservedVlans returns a boolean if a field has been set.

### GetVlanRanges

`func (o *NetworkFabricFabricConfiguration) GetVlanRanges() []string`

GetVlanRanges returns the VlanRanges field if non-nil, zero value otherwise.

### GetVlanRangesOk

`func (o *NetworkFabricFabricConfiguration) GetVlanRangesOk() (*[]string, bool)`

GetVlanRangesOk returns a tuple with the VlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRanges

`func (o *NetworkFabricFabricConfiguration) SetVlanRanges(v []string)`

SetVlanRanges sets VlanRanges field to given value.

### HasVlanRanges

`func (o *NetworkFabricFabricConfiguration) HasVlanRanges() bool`

HasVlanRanges returns a boolean if a field has been set.

### GetVniPrefix

`func (o *NetworkFabricFabricConfiguration) GetVniPrefix() int32`

GetVniPrefix returns the VniPrefix field if non-nil, zero value otherwise.

### GetVniPrefixOk

`func (o *NetworkFabricFabricConfiguration) GetVniPrefixOk() (*int32, bool)`

GetVniPrefixOk returns a tuple with the VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniPrefix

`func (o *NetworkFabricFabricConfiguration) SetVniPrefix(v int32)`

SetVniPrefix sets VniPrefix field to given value.

### HasVniPrefix

`func (o *NetworkFabricFabricConfiguration) HasVniPrefix() bool`

HasVniPrefix returns a boolean if a field has been set.

### GetPreventVrfCleanup

`func (o *NetworkFabricFabricConfiguration) GetPreventVrfCleanup() []string`

GetPreventVrfCleanup returns the PreventVrfCleanup field if non-nil, zero value otherwise.

### GetPreventVrfCleanupOk

`func (o *NetworkFabricFabricConfiguration) GetPreventVrfCleanupOk() (*[]string, bool)`

GetPreventVrfCleanupOk returns a tuple with the PreventVrfCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVrfCleanup

`func (o *NetworkFabricFabricConfiguration) SetPreventVrfCleanup(v []string)`

SetPreventVrfCleanup sets PreventVrfCleanup field to given value.

### HasPreventVrfCleanup

`func (o *NetworkFabricFabricConfiguration) HasPreventVrfCleanup() bool`

HasPreventVrfCleanup returns a boolean if a field has been set.

### GetReservedVrfs

`func (o *NetworkFabricFabricConfiguration) GetReservedVrfs() []string`

GetReservedVrfs returns the ReservedVrfs field if non-nil, zero value otherwise.

### GetReservedVrfsOk

`func (o *NetworkFabricFabricConfiguration) GetReservedVrfsOk() (*[]string, bool)`

GetReservedVrfsOk returns a tuple with the ReservedVrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVrfs

`func (o *NetworkFabricFabricConfiguration) SetReservedVrfs(v []string)`

SetReservedVrfs sets ReservedVrfs field to given value.

### HasReservedVrfs

`func (o *NetworkFabricFabricConfiguration) HasReservedVrfs() bool`

HasReservedVrfs returns a boolean if a field has been set.

### GetVrfVlanRanges

`func (o *NetworkFabricFabricConfiguration) GetVrfVlanRanges() []string`

GetVrfVlanRanges returns the VrfVlanRanges field if non-nil, zero value otherwise.

### GetVrfVlanRangesOk

`func (o *NetworkFabricFabricConfiguration) GetVrfVlanRangesOk() (*[]string, bool)`

GetVrfVlanRangesOk returns a tuple with the VrfVlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfVlanRanges

`func (o *NetworkFabricFabricConfiguration) SetVrfVlanRanges(v []string)`

SetVrfVlanRanges sets VrfVlanRanges field to given value.

### HasVrfVlanRanges

`func (o *NetworkFabricFabricConfiguration) HasVrfVlanRanges() bool`

HasVrfVlanRanges returns a boolean if a field has been set.

### GetVsanId

`func (o *NetworkFabricFabricConfiguration) GetVsanId() int32`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *NetworkFabricFabricConfiguration) GetVsanIdOk() (*int32, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *NetworkFabricFabricConfiguration) SetVsanId(v int32)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *NetworkFabricFabricConfiguration) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.

### GetTopologyType

`func (o *NetworkFabricFabricConfiguration) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *NetworkFabricFabricConfiguration) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *NetworkFabricFabricConfiguration) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.

### HasTopologyType

`func (o *NetworkFabricFabricConfiguration) HasTopologyType() bool`

HasTopologyType returns a boolean if a field has been set.

### GetMtu

`func (o *NetworkFabricFabricConfiguration) GetMtu() float32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkFabricFabricConfiguration) GetMtuOk() (*float32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkFabricFabricConfiguration) SetMtu(v float32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkFabricFabricConfiguration) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetZoningConfiguration

`func (o *NetworkFabricFabricConfiguration) GetZoningConfiguration() map[string]interface{}`

GetZoningConfiguration returns the ZoningConfiguration field if non-nil, zero value otherwise.

### GetZoningConfigurationOk

`func (o *NetworkFabricFabricConfiguration) GetZoningConfigurationOk() (*map[string]interface{}, bool)`

GetZoningConfigurationOk returns a tuple with the ZoningConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoningConfiguration

`func (o *NetworkFabricFabricConfiguration) SetZoningConfiguration(v map[string]interface{})`

SetZoningConfiguration sets ZoningConfiguration field to given value.

### HasZoningConfiguration

`func (o *NetworkFabricFabricConfiguration) HasZoningConfiguration() bool`

HasZoningConfiguration returns a boolean if a field has been set.

### GetInteropMode

`func (o *NetworkFabricFabricConfiguration) GetInteropMode() string`

GetInteropMode returns the InteropMode field if non-nil, zero value otherwise.

### GetInteropModeOk

`func (o *NetworkFabricFabricConfiguration) GetInteropModeOk() (*string, bool)`

GetInteropModeOk returns a tuple with the InteropMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteropMode

`func (o *NetworkFabricFabricConfiguration) SetInteropMode(v string)`

SetInteropMode sets InteropMode field to given value.

### HasInteropMode

`func (o *NetworkFabricFabricConfiguration) HasInteropMode() bool`

HasInteropMode returns a boolean if a field has been set.

### GetQosConfiguration

`func (o *NetworkFabricFabricConfiguration) GetQosConfiguration() map[string]interface{}`

GetQosConfiguration returns the QosConfiguration field if non-nil, zero value otherwise.

### GetQosConfigurationOk

`func (o *NetworkFabricFabricConfiguration) GetQosConfigurationOk() (*map[string]interface{}, bool)`

GetQosConfigurationOk returns a tuple with the QosConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfiguration

`func (o *NetworkFabricFabricConfiguration) SetQosConfiguration(v map[string]interface{})`

SetQosConfiguration sets QosConfiguration field to given value.

### HasQosConfiguration

`func (o *NetworkFabricFabricConfiguration) HasQosConfiguration() bool`

HasQosConfiguration returns a boolean if a field has been set.

### GetTrunkingConfiguration

`func (o *NetworkFabricFabricConfiguration) GetTrunkingConfiguration() map[string]interface{}`

GetTrunkingConfiguration returns the TrunkingConfiguration field if non-nil, zero value otherwise.

### GetTrunkingConfigurationOk

`func (o *NetworkFabricFabricConfiguration) GetTrunkingConfigurationOk() (*map[string]interface{}, bool)`

GetTrunkingConfigurationOk returns a tuple with the TrunkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkingConfiguration

`func (o *NetworkFabricFabricConfiguration) SetTrunkingConfiguration(v map[string]interface{})`

SetTrunkingConfiguration sets TrunkingConfiguration field to given value.

### HasTrunkingConfiguration

`func (o *NetworkFabricFabricConfiguration) HasTrunkingConfiguration() bool`

HasTrunkingConfiguration returns a boolean if a field has been set.

### GetPortChannelConfiguration

`func (o *NetworkFabricFabricConfiguration) GetPortChannelConfiguration() map[string]interface{}`

GetPortChannelConfiguration returns the PortChannelConfiguration field if non-nil, zero value otherwise.

### GetPortChannelConfigurationOk

`func (o *NetworkFabricFabricConfiguration) GetPortChannelConfigurationOk() (*map[string]interface{}, bool)`

GetPortChannelConfigurationOk returns a tuple with the PortChannelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelConfiguration

`func (o *NetworkFabricFabricConfiguration) SetPortChannelConfiguration(v map[string]interface{})`

SetPortChannelConfiguration sets PortChannelConfiguration field to given value.

### HasPortChannelConfiguration

`func (o *NetworkFabricFabricConfiguration) HasPortChannelConfiguration() bool`

HasPortChannelConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



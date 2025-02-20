# EthernetEvpnVxlanL2Fabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultNetworkProfileId** | **float32** |  | 
**GnmiMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**ZeroTouchEnabled** | Pointer to **bool** |  | [optional] 
**AllocateDefaultVlan** | Pointer to **bool** |  | [optional] 
**AsnRanges** | Pointer to **[]string** |  | [optional] 
**DefaultVlan** | Pointer to **float32** |  | [optional] 
**ExtraInternalIPsPerSubnet** | Pointer to **float32** |  | [optional] 
**LagRanges** | Pointer to **[]string** |  | [optional] 
**LeafSwitchesHaveMlagPairs** | Pointer to **bool** |  | [optional] 
**MlagRanges** | Pointer to **[]string** | MLAG ID ranges. Values must be between 1 and 4096. | [optional] 
**NumberOfSpinesNextToLeafSwitches** | **float32** | Number of spines next to leaf switches | 
**PreventVlanCleanup** | Pointer to **[]string** |  | [optional] 
**PreventCleanupFromUplinks** | Pointer to **bool** |  | [optional] 
**ReservedVlans** | Pointer to **[]string** |  | [optional] 
**VlanRanges** | Pointer to **[]string** | Array of VLAN range strings in format \&quot;start-end\&quot; | [optional] 
**VniPrefix** | Pointer to **float32** |  | [optional] 

## Methods

### NewEthernetEvpnVxlanL2Fabric

`func NewEthernetEvpnVxlanL2Fabric(defaultNetworkProfileId float32, numberOfSpinesNextToLeafSwitches float32, ) *EthernetEvpnVxlanL2Fabric`

NewEthernetEvpnVxlanL2Fabric instantiates a new EthernetEvpnVxlanL2Fabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetEvpnVxlanL2FabricWithDefaults

`func NewEthernetEvpnVxlanL2FabricWithDefaults() *EthernetEvpnVxlanL2Fabric`

NewEthernetEvpnVxlanL2FabricWithDefaults instantiates a new EthernetEvpnVxlanL2Fabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultNetworkProfileId

`func (o *EthernetEvpnVxlanL2Fabric) GetDefaultNetworkProfileId() float32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *EthernetEvpnVxlanL2Fabric) GetDefaultNetworkProfileIdOk() (*float32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *EthernetEvpnVxlanL2Fabric) SetDefaultNetworkProfileId(v float32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.


### GetGnmiMonitoringEnabled

`func (o *EthernetEvpnVxlanL2Fabric) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *EthernetEvpnVxlanL2Fabric) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *EthernetEvpnVxlanL2Fabric) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *EthernetEvpnVxlanL2Fabric) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *EthernetEvpnVxlanL2Fabric) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *EthernetEvpnVxlanL2Fabric) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *EthernetEvpnVxlanL2Fabric) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *EthernetEvpnVxlanL2Fabric) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetZeroTouchEnabled

`func (o *EthernetEvpnVxlanL2Fabric) GetZeroTouchEnabled() bool`

GetZeroTouchEnabled returns the ZeroTouchEnabled field if non-nil, zero value otherwise.

### GetZeroTouchEnabledOk

`func (o *EthernetEvpnVxlanL2Fabric) GetZeroTouchEnabledOk() (*bool, bool)`

GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchEnabled

`func (o *EthernetEvpnVxlanL2Fabric) SetZeroTouchEnabled(v bool)`

SetZeroTouchEnabled sets ZeroTouchEnabled field to given value.

### HasZeroTouchEnabled

`func (o *EthernetEvpnVxlanL2Fabric) HasZeroTouchEnabled() bool`

HasZeroTouchEnabled returns a boolean if a field has been set.

### GetAllocateDefaultVlan

`func (o *EthernetEvpnVxlanL2Fabric) GetAllocateDefaultVlan() bool`

GetAllocateDefaultVlan returns the AllocateDefaultVlan field if non-nil, zero value otherwise.

### GetAllocateDefaultVlanOk

`func (o *EthernetEvpnVxlanL2Fabric) GetAllocateDefaultVlanOk() (*bool, bool)`

GetAllocateDefaultVlanOk returns a tuple with the AllocateDefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateDefaultVlan

`func (o *EthernetEvpnVxlanL2Fabric) SetAllocateDefaultVlan(v bool)`

SetAllocateDefaultVlan sets AllocateDefaultVlan field to given value.

### HasAllocateDefaultVlan

`func (o *EthernetEvpnVxlanL2Fabric) HasAllocateDefaultVlan() bool`

HasAllocateDefaultVlan returns a boolean if a field has been set.

### GetAsnRanges

`func (o *EthernetEvpnVxlanL2Fabric) GetAsnRanges() []string`

GetAsnRanges returns the AsnRanges field if non-nil, zero value otherwise.

### GetAsnRangesOk

`func (o *EthernetEvpnVxlanL2Fabric) GetAsnRangesOk() (*[]string, bool)`

GetAsnRangesOk returns a tuple with the AsnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnRanges

`func (o *EthernetEvpnVxlanL2Fabric) SetAsnRanges(v []string)`

SetAsnRanges sets AsnRanges field to given value.

### HasAsnRanges

`func (o *EthernetEvpnVxlanL2Fabric) HasAsnRanges() bool`

HasAsnRanges returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *EthernetEvpnVxlanL2Fabric) GetDefaultVlan() float32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *EthernetEvpnVxlanL2Fabric) GetDefaultVlanOk() (*float32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *EthernetEvpnVxlanL2Fabric) SetDefaultVlan(v float32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *EthernetEvpnVxlanL2Fabric) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetExtraInternalIPsPerSubnet

`func (o *EthernetEvpnVxlanL2Fabric) GetExtraInternalIPsPerSubnet() float32`

GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field if non-nil, zero value otherwise.

### GetExtraInternalIPsPerSubnetOk

`func (o *EthernetEvpnVxlanL2Fabric) GetExtraInternalIPsPerSubnetOk() (*float32, bool)`

GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInternalIPsPerSubnet

`func (o *EthernetEvpnVxlanL2Fabric) SetExtraInternalIPsPerSubnet(v float32)`

SetExtraInternalIPsPerSubnet sets ExtraInternalIPsPerSubnet field to given value.

### HasExtraInternalIPsPerSubnet

`func (o *EthernetEvpnVxlanL2Fabric) HasExtraInternalIPsPerSubnet() bool`

HasExtraInternalIPsPerSubnet returns a boolean if a field has been set.

### GetLagRanges

`func (o *EthernetEvpnVxlanL2Fabric) GetLagRanges() []string`

GetLagRanges returns the LagRanges field if non-nil, zero value otherwise.

### GetLagRangesOk

`func (o *EthernetEvpnVxlanL2Fabric) GetLagRangesOk() (*[]string, bool)`

GetLagRangesOk returns a tuple with the LagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagRanges

`func (o *EthernetEvpnVxlanL2Fabric) SetLagRanges(v []string)`

SetLagRanges sets LagRanges field to given value.

### HasLagRanges

`func (o *EthernetEvpnVxlanL2Fabric) HasLagRanges() bool`

HasLagRanges returns a boolean if a field has been set.

### GetLeafSwitchesHaveMlagPairs

`func (o *EthernetEvpnVxlanL2Fabric) GetLeafSwitchesHaveMlagPairs() bool`

GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field if non-nil, zero value otherwise.

### GetLeafSwitchesHaveMlagPairsOk

`func (o *EthernetEvpnVxlanL2Fabric) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool)`

GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafSwitchesHaveMlagPairs

`func (o *EthernetEvpnVxlanL2Fabric) SetLeafSwitchesHaveMlagPairs(v bool)`

SetLeafSwitchesHaveMlagPairs sets LeafSwitchesHaveMlagPairs field to given value.

### HasLeafSwitchesHaveMlagPairs

`func (o *EthernetEvpnVxlanL2Fabric) HasLeafSwitchesHaveMlagPairs() bool`

HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.

### GetMlagRanges

`func (o *EthernetEvpnVxlanL2Fabric) GetMlagRanges() []string`

GetMlagRanges returns the MlagRanges field if non-nil, zero value otherwise.

### GetMlagRangesOk

`func (o *EthernetEvpnVxlanL2Fabric) GetMlagRangesOk() (*[]string, bool)`

GetMlagRangesOk returns a tuple with the MlagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagRanges

`func (o *EthernetEvpnVxlanL2Fabric) SetMlagRanges(v []string)`

SetMlagRanges sets MlagRanges field to given value.

### HasMlagRanges

`func (o *EthernetEvpnVxlanL2Fabric) HasMlagRanges() bool`

HasMlagRanges returns a boolean if a field has been set.

### GetNumberOfSpinesNextToLeafSwitches

`func (o *EthernetEvpnVxlanL2Fabric) GetNumberOfSpinesNextToLeafSwitches() float32`

GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field if non-nil, zero value otherwise.

### GetNumberOfSpinesNextToLeafSwitchesOk

`func (o *EthernetEvpnVxlanL2Fabric) GetNumberOfSpinesNextToLeafSwitchesOk() (*float32, bool)`

GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesNextToLeafSwitches

`func (o *EthernetEvpnVxlanL2Fabric) SetNumberOfSpinesNextToLeafSwitches(v float32)`

SetNumberOfSpinesNextToLeafSwitches sets NumberOfSpinesNextToLeafSwitches field to given value.


### GetPreventVlanCleanup

`func (o *EthernetEvpnVxlanL2Fabric) GetPreventVlanCleanup() []string`

GetPreventVlanCleanup returns the PreventVlanCleanup field if non-nil, zero value otherwise.

### GetPreventVlanCleanupOk

`func (o *EthernetEvpnVxlanL2Fabric) GetPreventVlanCleanupOk() (*[]string, bool)`

GetPreventVlanCleanupOk returns a tuple with the PreventVlanCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVlanCleanup

`func (o *EthernetEvpnVxlanL2Fabric) SetPreventVlanCleanup(v []string)`

SetPreventVlanCleanup sets PreventVlanCleanup field to given value.

### HasPreventVlanCleanup

`func (o *EthernetEvpnVxlanL2Fabric) HasPreventVlanCleanup() bool`

HasPreventVlanCleanup returns a boolean if a field has been set.

### GetPreventCleanupFromUplinks

`func (o *EthernetEvpnVxlanL2Fabric) GetPreventCleanupFromUplinks() bool`

GetPreventCleanupFromUplinks returns the PreventCleanupFromUplinks field if non-nil, zero value otherwise.

### GetPreventCleanupFromUplinksOk

`func (o *EthernetEvpnVxlanL2Fabric) GetPreventCleanupFromUplinksOk() (*bool, bool)`

GetPreventCleanupFromUplinksOk returns a tuple with the PreventCleanupFromUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventCleanupFromUplinks

`func (o *EthernetEvpnVxlanL2Fabric) SetPreventCleanupFromUplinks(v bool)`

SetPreventCleanupFromUplinks sets PreventCleanupFromUplinks field to given value.

### HasPreventCleanupFromUplinks

`func (o *EthernetEvpnVxlanL2Fabric) HasPreventCleanupFromUplinks() bool`

HasPreventCleanupFromUplinks returns a boolean if a field has been set.

### GetReservedVlans

`func (o *EthernetEvpnVxlanL2Fabric) GetReservedVlans() []string`

GetReservedVlans returns the ReservedVlans field if non-nil, zero value otherwise.

### GetReservedVlansOk

`func (o *EthernetEvpnVxlanL2Fabric) GetReservedVlansOk() (*[]string, bool)`

GetReservedVlansOk returns a tuple with the ReservedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVlans

`func (o *EthernetEvpnVxlanL2Fabric) SetReservedVlans(v []string)`

SetReservedVlans sets ReservedVlans field to given value.

### HasReservedVlans

`func (o *EthernetEvpnVxlanL2Fabric) HasReservedVlans() bool`

HasReservedVlans returns a boolean if a field has been set.

### GetVlanRanges

`func (o *EthernetEvpnVxlanL2Fabric) GetVlanRanges() []string`

GetVlanRanges returns the VlanRanges field if non-nil, zero value otherwise.

### GetVlanRangesOk

`func (o *EthernetEvpnVxlanL2Fabric) GetVlanRangesOk() (*[]string, bool)`

GetVlanRangesOk returns a tuple with the VlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRanges

`func (o *EthernetEvpnVxlanL2Fabric) SetVlanRanges(v []string)`

SetVlanRanges sets VlanRanges field to given value.

### HasVlanRanges

`func (o *EthernetEvpnVxlanL2Fabric) HasVlanRanges() bool`

HasVlanRanges returns a boolean if a field has been set.

### GetVniPrefix

`func (o *EthernetEvpnVxlanL2Fabric) GetVniPrefix() float32`

GetVniPrefix returns the VniPrefix field if non-nil, zero value otherwise.

### GetVniPrefixOk

`func (o *EthernetEvpnVxlanL2Fabric) GetVniPrefixOk() (*float32, bool)`

GetVniPrefixOk returns a tuple with the VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniPrefix

`func (o *EthernetEvpnVxlanL2Fabric) SetVniPrefix(v float32)`

SetVniPrefix sets VniPrefix field to given value.

### HasVniPrefix

`func (o *EthernetEvpnVxlanL2Fabric) HasVniPrefix() bool`

HasVniPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



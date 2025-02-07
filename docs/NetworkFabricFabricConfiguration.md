# NetworkFabricFabricConfiguration

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
**PreventVrfCleanup** | Pointer to **[]string** |  | [optional] 
**ReservedVrfs** | Pointer to **[]string** |  | [optional] 
**VrfVlanRanges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkFabricFabricConfiguration

`func NewNetworkFabricFabricConfiguration(defaultNetworkProfileId float32, numberOfSpinesNextToLeafSwitches float32, ) *NetworkFabricFabricConfiguration`

NewNetworkFabricFabricConfiguration instantiates a new NetworkFabricFabricConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricFabricConfigurationWithDefaults

`func NewNetworkFabricFabricConfigurationWithDefaults() *NetworkFabricFabricConfiguration`

NewNetworkFabricFabricConfigurationWithDefaults instantiates a new NetworkFabricFabricConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultNetworkProfileId

`func (o *NetworkFabricFabricConfiguration) GetDefaultNetworkProfileId() float32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *NetworkFabricFabricConfiguration) GetDefaultNetworkProfileIdOk() (*float32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *NetworkFabricFabricConfiguration) SetDefaultNetworkProfileId(v float32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.


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

`func (o *NetworkFabricFabricConfiguration) GetDefaultVlan() float32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *NetworkFabricFabricConfiguration) GetDefaultVlanOk() (*float32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *NetworkFabricFabricConfiguration) SetDefaultVlan(v float32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *NetworkFabricFabricConfiguration) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetExtraInternalIPsPerSubnet

`func (o *NetworkFabricFabricConfiguration) GetExtraInternalIPsPerSubnet() float32`

GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field if non-nil, zero value otherwise.

### GetExtraInternalIPsPerSubnetOk

`func (o *NetworkFabricFabricConfiguration) GetExtraInternalIPsPerSubnetOk() (*float32, bool)`

GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInternalIPsPerSubnet

`func (o *NetworkFabricFabricConfiguration) SetExtraInternalIPsPerSubnet(v float32)`

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

`func (o *NetworkFabricFabricConfiguration) GetNumberOfSpinesNextToLeafSwitches() float32`

GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field if non-nil, zero value otherwise.

### GetNumberOfSpinesNextToLeafSwitchesOk

`func (o *NetworkFabricFabricConfiguration) GetNumberOfSpinesNextToLeafSwitchesOk() (*float32, bool)`

GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesNextToLeafSwitches

`func (o *NetworkFabricFabricConfiguration) SetNumberOfSpinesNextToLeafSwitches(v float32)`

SetNumberOfSpinesNextToLeafSwitches sets NumberOfSpinesNextToLeafSwitches field to given value.


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

`func (o *NetworkFabricFabricConfiguration) GetVniPrefix() float32`

GetVniPrefix returns the VniPrefix field if non-nil, zero value otherwise.

### GetVniPrefixOk

`func (o *NetworkFabricFabricConfiguration) GetVniPrefixOk() (*float32, bool)`

GetVniPrefixOk returns a tuple with the VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniPrefix

`func (o *NetworkFabricFabricConfiguration) SetVniPrefix(v float32)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



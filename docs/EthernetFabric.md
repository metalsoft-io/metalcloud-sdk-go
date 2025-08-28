# EthernetFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | [**FabricType**](FabricType.md) | The type of network fabric | 
**DefaultNetworkProfileId** | Pointer to **int32** | Unique identifier for the default network profile. Must be a positive integer (minimum: 1) corresponding to an existing profile. | [optional] 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**ServerOnlyOperationEnabled** | Pointer to **bool** | Enables server-only operation mode on the network fabric. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 
**ZeroTouchEnabled** | Pointer to **bool** | Enables zero-touch provisioning for automatic device configuration. | [optional] 
**AsnRanges** | Pointer to **[]string** | ASN ranges in the format \&quot;start-end\&quot;, where each range is an ordered pair with values between 1 and 4294967295. | [optional] 
**AsnAllocationStrategy** | Pointer to [**AsnAllocationStrategy**](AsnAllocationStrategy.md) |  | [optional] 
**BgpNumbering** | Pointer to [**BgpNumberingType**](BgpNumberingType.md) |  | [optional] 
**LibraryLabel** | Pointer to **string** | Library label for the Ethernet fabric configuration, used to identify the BGP templates that can be used for deploy. | [optional] 
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
**VrfVlanRanges** | Pointer to **[]string** | VLAN ranges to be associated with VRF instances. Each value must be an ordered pair specified in the \&quot;start-end\&quot; format. | [optional] 

## Methods

### NewEthernetFabric

`func NewEthernetFabric(fabricType FabricType, ) *EthernetFabric`

NewEthernetFabric instantiates a new EthernetFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthernetFabricWithDefaults

`func NewEthernetFabricWithDefaults() *EthernetFabric`

NewEthernetFabricWithDefaults instantiates a new EthernetFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFabricType

`func (o *EthernetFabric) GetFabricType() FabricType`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *EthernetFabric) GetFabricTypeOk() (*FabricType, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *EthernetFabric) SetFabricType(v FabricType)`

SetFabricType sets FabricType field to given value.


### GetDefaultNetworkProfileId

`func (o *EthernetFabric) GetDefaultNetworkProfileId() int32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *EthernetFabric) GetDefaultNetworkProfileIdOk() (*int32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *EthernetFabric) SetDefaultNetworkProfileId(v int32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.

### HasDefaultNetworkProfileId

`func (o *EthernetFabric) HasDefaultNetworkProfileId() bool`

HasDefaultNetworkProfileId returns a boolean if a field has been set.

### GetGnmiMonitoringEnabled

`func (o *EthernetFabric) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *EthernetFabric) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *EthernetFabric) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *EthernetFabric) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetServerOnlyOperationEnabled

`func (o *EthernetFabric) GetServerOnlyOperationEnabled() bool`

GetServerOnlyOperationEnabled returns the ServerOnlyOperationEnabled field if non-nil, zero value otherwise.

### GetServerOnlyOperationEnabledOk

`func (o *EthernetFabric) GetServerOnlyOperationEnabledOk() (*bool, bool)`

GetServerOnlyOperationEnabledOk returns a tuple with the ServerOnlyOperationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOnlyOperationEnabled

`func (o *EthernetFabric) SetServerOnlyOperationEnabled(v bool)`

SetServerOnlyOperationEnabled sets ServerOnlyOperationEnabled field to given value.

### HasServerOnlyOperationEnabled

`func (o *EthernetFabric) HasServerOnlyOperationEnabled() bool`

HasServerOnlyOperationEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *EthernetFabric) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *EthernetFabric) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *EthernetFabric) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *EthernetFabric) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetZeroTouchEnabled

`func (o *EthernetFabric) GetZeroTouchEnabled() bool`

GetZeroTouchEnabled returns the ZeroTouchEnabled field if non-nil, zero value otherwise.

### GetZeroTouchEnabledOk

`func (o *EthernetFabric) GetZeroTouchEnabledOk() (*bool, bool)`

GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchEnabled

`func (o *EthernetFabric) SetZeroTouchEnabled(v bool)`

SetZeroTouchEnabled sets ZeroTouchEnabled field to given value.

### HasZeroTouchEnabled

`func (o *EthernetFabric) HasZeroTouchEnabled() bool`

HasZeroTouchEnabled returns a boolean if a field has been set.

### GetAsnRanges

`func (o *EthernetFabric) GetAsnRanges() []string`

GetAsnRanges returns the AsnRanges field if non-nil, zero value otherwise.

### GetAsnRangesOk

`func (o *EthernetFabric) GetAsnRangesOk() (*[]string, bool)`

GetAsnRangesOk returns a tuple with the AsnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnRanges

`func (o *EthernetFabric) SetAsnRanges(v []string)`

SetAsnRanges sets AsnRanges field to given value.

### HasAsnRanges

`func (o *EthernetFabric) HasAsnRanges() bool`

HasAsnRanges returns a boolean if a field has been set.

### GetAsnAllocationStrategy

`func (o *EthernetFabric) GetAsnAllocationStrategy() AsnAllocationStrategy`

GetAsnAllocationStrategy returns the AsnAllocationStrategy field if non-nil, zero value otherwise.

### GetAsnAllocationStrategyOk

`func (o *EthernetFabric) GetAsnAllocationStrategyOk() (*AsnAllocationStrategy, bool)`

GetAsnAllocationStrategyOk returns a tuple with the AsnAllocationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnAllocationStrategy

`func (o *EthernetFabric) SetAsnAllocationStrategy(v AsnAllocationStrategy)`

SetAsnAllocationStrategy sets AsnAllocationStrategy field to given value.

### HasAsnAllocationStrategy

`func (o *EthernetFabric) HasAsnAllocationStrategy() bool`

HasAsnAllocationStrategy returns a boolean if a field has been set.

### GetBgpNumbering

`func (o *EthernetFabric) GetBgpNumbering() BgpNumberingType`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *EthernetFabric) GetBgpNumberingOk() (*BgpNumberingType, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *EthernetFabric) SetBgpNumbering(v BgpNumberingType)`

SetBgpNumbering sets BgpNumbering field to given value.

### HasBgpNumbering

`func (o *EthernetFabric) HasBgpNumbering() bool`

HasBgpNumbering returns a boolean if a field has been set.

### GetLibraryLabel

`func (o *EthernetFabric) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *EthernetFabric) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *EthernetFabric) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.

### HasLibraryLabel

`func (o *EthernetFabric) HasLibraryLabel() bool`

HasLibraryLabel returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *EthernetFabric) GetDefaultVlan() int32`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *EthernetFabric) GetDefaultVlanOk() (*int32, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *EthernetFabric) SetDefaultVlan(v int32)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *EthernetFabric) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetExtraInternalIPsPerSubnet

`func (o *EthernetFabric) GetExtraInternalIPsPerSubnet() int32`

GetExtraInternalIPsPerSubnet returns the ExtraInternalIPsPerSubnet field if non-nil, zero value otherwise.

### GetExtraInternalIPsPerSubnetOk

`func (o *EthernetFabric) GetExtraInternalIPsPerSubnetOk() (*int32, bool)`

GetExtraInternalIPsPerSubnetOk returns a tuple with the ExtraInternalIPsPerSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInternalIPsPerSubnet

`func (o *EthernetFabric) SetExtraInternalIPsPerSubnet(v int32)`

SetExtraInternalIPsPerSubnet sets ExtraInternalIPsPerSubnet field to given value.

### HasExtraInternalIPsPerSubnet

`func (o *EthernetFabric) HasExtraInternalIPsPerSubnet() bool`

HasExtraInternalIPsPerSubnet returns a boolean if a field has been set.

### GetLagRanges

`func (o *EthernetFabric) GetLagRanges() []string`

GetLagRanges returns the LagRanges field if non-nil, zero value otherwise.

### GetLagRangesOk

`func (o *EthernetFabric) GetLagRangesOk() (*[]string, bool)`

GetLagRangesOk returns a tuple with the LagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagRanges

`func (o *EthernetFabric) SetLagRanges(v []string)`

SetLagRanges sets LagRanges field to given value.

### HasLagRanges

`func (o *EthernetFabric) HasLagRanges() bool`

HasLagRanges returns a boolean if a field has been set.

### GetLeafSwitchesHaveMlagPairs

`func (o *EthernetFabric) GetLeafSwitchesHaveMlagPairs() bool`

GetLeafSwitchesHaveMlagPairs returns the LeafSwitchesHaveMlagPairs field if non-nil, zero value otherwise.

### GetLeafSwitchesHaveMlagPairsOk

`func (o *EthernetFabric) GetLeafSwitchesHaveMlagPairsOk() (*bool, bool)`

GetLeafSwitchesHaveMlagPairsOk returns a tuple with the LeafSwitchesHaveMlagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafSwitchesHaveMlagPairs

`func (o *EthernetFabric) SetLeafSwitchesHaveMlagPairs(v bool)`

SetLeafSwitchesHaveMlagPairs sets LeafSwitchesHaveMlagPairs field to given value.

### HasLeafSwitchesHaveMlagPairs

`func (o *EthernetFabric) HasLeafSwitchesHaveMlagPairs() bool`

HasLeafSwitchesHaveMlagPairs returns a boolean if a field has been set.

### GetMlagRanges

`func (o *EthernetFabric) GetMlagRanges() []string`

GetMlagRanges returns the MlagRanges field if non-nil, zero value otherwise.

### GetMlagRangesOk

`func (o *EthernetFabric) GetMlagRangesOk() (*[]string, bool)`

GetMlagRangesOk returns a tuple with the MlagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagRanges

`func (o *EthernetFabric) SetMlagRanges(v []string)`

SetMlagRanges sets MlagRanges field to given value.

### HasMlagRanges

`func (o *EthernetFabric) HasMlagRanges() bool`

HasMlagRanges returns a boolean if a field has been set.

### GetNumberOfSpinesNextToLeafSwitches

`func (o *EthernetFabric) GetNumberOfSpinesNextToLeafSwitches() int32`

GetNumberOfSpinesNextToLeafSwitches returns the NumberOfSpinesNextToLeafSwitches field if non-nil, zero value otherwise.

### GetNumberOfSpinesNextToLeafSwitchesOk

`func (o *EthernetFabric) GetNumberOfSpinesNextToLeafSwitchesOk() (*int32, bool)`

GetNumberOfSpinesNextToLeafSwitchesOk returns a tuple with the NumberOfSpinesNextToLeafSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesNextToLeafSwitches

`func (o *EthernetFabric) SetNumberOfSpinesNextToLeafSwitches(v int32)`

SetNumberOfSpinesNextToLeafSwitches sets NumberOfSpinesNextToLeafSwitches field to given value.

### HasNumberOfSpinesNextToLeafSwitches

`func (o *EthernetFabric) HasNumberOfSpinesNextToLeafSwitches() bool`

HasNumberOfSpinesNextToLeafSwitches returns a boolean if a field has been set.

### GetPreventVlanCleanup

`func (o *EthernetFabric) GetPreventVlanCleanup() []string`

GetPreventVlanCleanup returns the PreventVlanCleanup field if non-nil, zero value otherwise.

### GetPreventVlanCleanupOk

`func (o *EthernetFabric) GetPreventVlanCleanupOk() (*[]string, bool)`

GetPreventVlanCleanupOk returns a tuple with the PreventVlanCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVlanCleanup

`func (o *EthernetFabric) SetPreventVlanCleanup(v []string)`

SetPreventVlanCleanup sets PreventVlanCleanup field to given value.

### HasPreventVlanCleanup

`func (o *EthernetFabric) HasPreventVlanCleanup() bool`

HasPreventVlanCleanup returns a boolean if a field has been set.

### GetPreventCleanupFromUplinks

`func (o *EthernetFabric) GetPreventCleanupFromUplinks() bool`

GetPreventCleanupFromUplinks returns the PreventCleanupFromUplinks field if non-nil, zero value otherwise.

### GetPreventCleanupFromUplinksOk

`func (o *EthernetFabric) GetPreventCleanupFromUplinksOk() (*bool, bool)`

GetPreventCleanupFromUplinksOk returns a tuple with the PreventCleanupFromUplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventCleanupFromUplinks

`func (o *EthernetFabric) SetPreventCleanupFromUplinks(v bool)`

SetPreventCleanupFromUplinks sets PreventCleanupFromUplinks field to given value.

### HasPreventCleanupFromUplinks

`func (o *EthernetFabric) HasPreventCleanupFromUplinks() bool`

HasPreventCleanupFromUplinks returns a boolean if a field has been set.

### GetReservedVlans

`func (o *EthernetFabric) GetReservedVlans() []string`

GetReservedVlans returns the ReservedVlans field if non-nil, zero value otherwise.

### GetReservedVlansOk

`func (o *EthernetFabric) GetReservedVlansOk() (*[]string, bool)`

GetReservedVlansOk returns a tuple with the ReservedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedVlans

`func (o *EthernetFabric) SetReservedVlans(v []string)`

SetReservedVlans sets ReservedVlans field to given value.

### HasReservedVlans

`func (o *EthernetFabric) HasReservedVlans() bool`

HasReservedVlans returns a boolean if a field has been set.

### GetVlanRanges

`func (o *EthernetFabric) GetVlanRanges() []string`

GetVlanRanges returns the VlanRanges field if non-nil, zero value otherwise.

### GetVlanRangesOk

`func (o *EthernetFabric) GetVlanRangesOk() (*[]string, bool)`

GetVlanRangesOk returns a tuple with the VlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRanges

`func (o *EthernetFabric) SetVlanRanges(v []string)`

SetVlanRanges sets VlanRanges field to given value.

### HasVlanRanges

`func (o *EthernetFabric) HasVlanRanges() bool`

HasVlanRanges returns a boolean if a field has been set.

### GetVniPrefix

`func (o *EthernetFabric) GetVniPrefix() int32`

GetVniPrefix returns the VniPrefix field if non-nil, zero value otherwise.

### GetVniPrefixOk

`func (o *EthernetFabric) GetVniPrefixOk() (*int32, bool)`

GetVniPrefixOk returns a tuple with the VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniPrefix

`func (o *EthernetFabric) SetVniPrefix(v int32)`

SetVniPrefix sets VniPrefix field to given value.

### HasVniPrefix

`func (o *EthernetFabric) HasVniPrefix() bool`

HasVniPrefix returns a boolean if a field has been set.

### GetVrfVlanRanges

`func (o *EthernetFabric) GetVrfVlanRanges() []string`

GetVrfVlanRanges returns the VrfVlanRanges field if non-nil, zero value otherwise.

### GetVrfVlanRangesOk

`func (o *EthernetFabric) GetVrfVlanRangesOk() (*[]string, bool)`

GetVrfVlanRangesOk returns a tuple with the VrfVlanRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfVlanRanges

`func (o *EthernetFabric) SetVrfVlanRanges(v []string)`

SetVrfVlanRanges sets VrfVlanRanges field to given value.

### HasVrfVlanRanges

`func (o *EthernetFabric) HasVrfVlanRanges() bool`

HasVrfVlanRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



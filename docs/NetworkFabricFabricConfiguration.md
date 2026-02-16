# NetworkFabricFabricConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | [**FabricType**](FabricType.md) | The type of network fabric | 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**ServerOnlyOperationEnabled** | Pointer to **bool** | Enables server-only operation mode on the network fabric. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 
**BgpNumbering** | Pointer to [**BgpNumberingType**](BgpNumberingType.md) |  | [optional] 
**LibraryLabel** | Pointer to **string** | Library label for the Ethernet fabric configuration, used to identify the Network Device Configuration Templates that can be used for deploy. | [optional] 
**LagRanges** | Pointer to **[]string** | Link Aggregation (LAG) ranges in the format \&quot;start-end\&quot;; each range must be within the bounds of 1 to 4096. | [optional] 
**MlagRanges** | Pointer to **[]string** | MLAG ID ranges. Each range must be provided in \&quot;start-end\&quot; format with values between 1 and 4096. | [optional] 
**PreventVlanCleanup** | Pointer to **[]string** | VLAN ranges that should be prevented from automatic cleanup. Format must be \&quot;start-end\&quot;. | [optional] 
**ReservedVlans** | Pointer to **[]string** | Reserved VLAN ranges that are excluded from general allocation. Must follow the \&quot;start-end\&quot; format. | [optional] 
**VlanRanges** | Pointer to **[]string** | Array of VLAN range strings in \&quot;start-end\&quot; format to be used in configuration. | [optional] 
**VrfVlanRanges** | Pointer to **[]string** | VLAN ranges to be associated with VRF instances. Each value must be an ordered pair specified in the \&quot;start-end\&quot; format. | [optional] 
**VniPrefix** | Pointer to **int32** | The VNI prefix for the EVPN VXLAN fabric. | [optional] 
**L3VniPrefix** | Pointer to **int32** | The L3 VNI prefix for the EVPN VXLAN fabric. | [optional] 
**PreventPKeyCleanup** | Pointer to **[]string** | PKey ranges that should be prevented from automatic cleanup. Format must be \&quot;start-end\&quot;. | [optional] 
**ReservedPkeys** | Pointer to **[]string** | Reserved PKey ranges that are excluded from general allocation. Must follow the \&quot;start-end\&quot; format. | [optional] 
**PkeyRanges** | Pointer to **[]string** | Array of PKey range strings in \&quot;start-end\&quot; format to be used in configuration. | [optional] 

## Methods

### NewNetworkFabricFabricConfiguration

`func NewNetworkFabricFabricConfiguration(fabricType FabricType, ) *NetworkFabricFabricConfiguration`

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

`func (o *NetworkFabricFabricConfiguration) GetFabricType() FabricType`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *NetworkFabricFabricConfiguration) GetFabricTypeOk() (*FabricType, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *NetworkFabricFabricConfiguration) SetFabricType(v FabricType)`

SetFabricType sets FabricType field to given value.


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

### GetServerOnlyOperationEnabled

`func (o *NetworkFabricFabricConfiguration) GetServerOnlyOperationEnabled() bool`

GetServerOnlyOperationEnabled returns the ServerOnlyOperationEnabled field if non-nil, zero value otherwise.

### GetServerOnlyOperationEnabledOk

`func (o *NetworkFabricFabricConfiguration) GetServerOnlyOperationEnabledOk() (*bool, bool)`

GetServerOnlyOperationEnabledOk returns a tuple with the ServerOnlyOperationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOnlyOperationEnabled

`func (o *NetworkFabricFabricConfiguration) SetServerOnlyOperationEnabled(v bool)`

SetServerOnlyOperationEnabled sets ServerOnlyOperationEnabled field to given value.

### HasServerOnlyOperationEnabled

`func (o *NetworkFabricFabricConfiguration) HasServerOnlyOperationEnabled() bool`

HasServerOnlyOperationEnabled returns a boolean if a field has been set.

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

### GetBgpNumbering

`func (o *NetworkFabricFabricConfiguration) GetBgpNumbering() BgpNumberingType`

GetBgpNumbering returns the BgpNumbering field if non-nil, zero value otherwise.

### GetBgpNumberingOk

`func (o *NetworkFabricFabricConfiguration) GetBgpNumberingOk() (*BgpNumberingType, bool)`

GetBgpNumberingOk returns a tuple with the BgpNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNumbering

`func (o *NetworkFabricFabricConfiguration) SetBgpNumbering(v BgpNumberingType)`

SetBgpNumbering sets BgpNumbering field to given value.

### HasBgpNumbering

`func (o *NetworkFabricFabricConfiguration) HasBgpNumbering() bool`

HasBgpNumbering returns a boolean if a field has been set.

### GetLibraryLabel

`func (o *NetworkFabricFabricConfiguration) GetLibraryLabel() string`

GetLibraryLabel returns the LibraryLabel field if non-nil, zero value otherwise.

### GetLibraryLabelOk

`func (o *NetworkFabricFabricConfiguration) GetLibraryLabelOk() (*string, bool)`

GetLibraryLabelOk returns a tuple with the LibraryLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryLabel

`func (o *NetworkFabricFabricConfiguration) SetLibraryLabel(v string)`

SetLibraryLabel sets LibraryLabel field to given value.

### HasLibraryLabel

`func (o *NetworkFabricFabricConfiguration) HasLibraryLabel() bool`

HasLibraryLabel returns a boolean if a field has been set.

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

### GetL3VniPrefix

`func (o *NetworkFabricFabricConfiguration) GetL3VniPrefix() int32`

GetL3VniPrefix returns the L3VniPrefix field if non-nil, zero value otherwise.

### GetL3VniPrefixOk

`func (o *NetworkFabricFabricConfiguration) GetL3VniPrefixOk() (*int32, bool)`

GetL3VniPrefixOk returns a tuple with the L3VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VniPrefix

`func (o *NetworkFabricFabricConfiguration) SetL3VniPrefix(v int32)`

SetL3VniPrefix sets L3VniPrefix field to given value.

### HasL3VniPrefix

`func (o *NetworkFabricFabricConfiguration) HasL3VniPrefix() bool`

HasL3VniPrefix returns a boolean if a field has been set.

### GetPreventPKeyCleanup

`func (o *NetworkFabricFabricConfiguration) GetPreventPKeyCleanup() []string`

GetPreventPKeyCleanup returns the PreventPKeyCleanup field if non-nil, zero value otherwise.

### GetPreventPKeyCleanupOk

`func (o *NetworkFabricFabricConfiguration) GetPreventPKeyCleanupOk() (*[]string, bool)`

GetPreventPKeyCleanupOk returns a tuple with the PreventPKeyCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventPKeyCleanup

`func (o *NetworkFabricFabricConfiguration) SetPreventPKeyCleanup(v []string)`

SetPreventPKeyCleanup sets PreventPKeyCleanup field to given value.

### HasPreventPKeyCleanup

`func (o *NetworkFabricFabricConfiguration) HasPreventPKeyCleanup() bool`

HasPreventPKeyCleanup returns a boolean if a field has been set.

### GetReservedPkeys

`func (o *NetworkFabricFabricConfiguration) GetReservedPkeys() []string`

GetReservedPkeys returns the ReservedPkeys field if non-nil, zero value otherwise.

### GetReservedPkeysOk

`func (o *NetworkFabricFabricConfiguration) GetReservedPkeysOk() (*[]string, bool)`

GetReservedPkeysOk returns a tuple with the ReservedPkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPkeys

`func (o *NetworkFabricFabricConfiguration) SetReservedPkeys(v []string)`

SetReservedPkeys sets ReservedPkeys field to given value.

### HasReservedPkeys

`func (o *NetworkFabricFabricConfiguration) HasReservedPkeys() bool`

HasReservedPkeys returns a boolean if a field has been set.

### GetPkeyRanges

`func (o *NetworkFabricFabricConfiguration) GetPkeyRanges() []string`

GetPkeyRanges returns the PkeyRanges field if non-nil, zero value otherwise.

### GetPkeyRangesOk

`func (o *NetworkFabricFabricConfiguration) GetPkeyRangesOk() (*[]string, bool)`

GetPkeyRangesOk returns a tuple with the PkeyRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeyRanges

`func (o *NetworkFabricFabricConfiguration) SetPkeyRanges(v []string)`

SetPkeyRanges sets PkeyRanges field to given value.

### HasPkeyRanges

`func (o *NetworkFabricFabricConfiguration) HasPkeyRanges() bool`

HasPkeyRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



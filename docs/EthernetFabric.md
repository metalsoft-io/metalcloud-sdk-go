# EthernetFabric

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

### GetL3VniPrefix

`func (o *EthernetFabric) GetL3VniPrefix() int32`

GetL3VniPrefix returns the L3VniPrefix field if non-nil, zero value otherwise.

### GetL3VniPrefixOk

`func (o *EthernetFabric) GetL3VniPrefixOk() (*int32, bool)`

GetL3VniPrefixOk returns a tuple with the L3VniPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VniPrefix

`func (o *EthernetFabric) SetL3VniPrefix(v int32)`

SetL3VniPrefix sets L3VniPrefix field to given value.

### HasL3VniPrefix

`func (o *EthernetFabric) HasL3VniPrefix() bool`

HasL3VniPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InfinibandFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | [**FabricType**](FabricType.md) | The type of network fabric | 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**ServerOnlyOperationEnabled** | Pointer to **bool** | Enables server-only operation mode on the network fabric. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 
**LagRanges** | Pointer to **[]string** | Link Aggregation (LAG) ranges in the format \&quot;start-end\&quot;; each range must be within the bounds of 1 to 4096. | [optional] 
**PreventPKeyCleanup** | Pointer to **[]string** | PKey ranges that should be prevented from automatic cleanup. Format must be \&quot;start-end\&quot;. | [optional] 
**ReservedPkeys** | Pointer to **[]string** | Reserved PKey ranges that are excluded from general allocation. Must follow the \&quot;start-end\&quot; format. | [optional] 
**PkeyRanges** | Pointer to **[]string** | Array of PKey range strings in \&quot;start-end\&quot; format to be used in configuration. | [optional] 

## Methods

### NewInfinibandFabric

`func NewInfinibandFabric(fabricType FabricType, ) *InfinibandFabric`

NewInfinibandFabric instantiates a new InfinibandFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfinibandFabricWithDefaults

`func NewInfinibandFabricWithDefaults() *InfinibandFabric`

NewInfinibandFabricWithDefaults instantiates a new InfinibandFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFabricType

`func (o *InfinibandFabric) GetFabricType() FabricType`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *InfinibandFabric) GetFabricTypeOk() (*FabricType, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *InfinibandFabric) SetFabricType(v FabricType)`

SetFabricType sets FabricType field to given value.


### GetGnmiMonitoringEnabled

`func (o *InfinibandFabric) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *InfinibandFabric) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *InfinibandFabric) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *InfinibandFabric) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetServerOnlyOperationEnabled

`func (o *InfinibandFabric) GetServerOnlyOperationEnabled() bool`

GetServerOnlyOperationEnabled returns the ServerOnlyOperationEnabled field if non-nil, zero value otherwise.

### GetServerOnlyOperationEnabledOk

`func (o *InfinibandFabric) GetServerOnlyOperationEnabledOk() (*bool, bool)`

GetServerOnlyOperationEnabledOk returns a tuple with the ServerOnlyOperationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOnlyOperationEnabled

`func (o *InfinibandFabric) SetServerOnlyOperationEnabled(v bool)`

SetServerOnlyOperationEnabled sets ServerOnlyOperationEnabled field to given value.

### HasServerOnlyOperationEnabled

`func (o *InfinibandFabric) HasServerOnlyOperationEnabled() bool`

HasServerOnlyOperationEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *InfinibandFabric) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *InfinibandFabric) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *InfinibandFabric) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *InfinibandFabric) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetLagRanges

`func (o *InfinibandFabric) GetLagRanges() []string`

GetLagRanges returns the LagRanges field if non-nil, zero value otherwise.

### GetLagRangesOk

`func (o *InfinibandFabric) GetLagRangesOk() (*[]string, bool)`

GetLagRangesOk returns a tuple with the LagRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagRanges

`func (o *InfinibandFabric) SetLagRanges(v []string)`

SetLagRanges sets LagRanges field to given value.

### HasLagRanges

`func (o *InfinibandFabric) HasLagRanges() bool`

HasLagRanges returns a boolean if a field has been set.

### GetPreventPKeyCleanup

`func (o *InfinibandFabric) GetPreventPKeyCleanup() []string`

GetPreventPKeyCleanup returns the PreventPKeyCleanup field if non-nil, zero value otherwise.

### GetPreventPKeyCleanupOk

`func (o *InfinibandFabric) GetPreventPKeyCleanupOk() (*[]string, bool)`

GetPreventPKeyCleanupOk returns a tuple with the PreventPKeyCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventPKeyCleanup

`func (o *InfinibandFabric) SetPreventPKeyCleanup(v []string)`

SetPreventPKeyCleanup sets PreventPKeyCleanup field to given value.

### HasPreventPKeyCleanup

`func (o *InfinibandFabric) HasPreventPKeyCleanup() bool`

HasPreventPKeyCleanup returns a boolean if a field has been set.

### GetReservedPkeys

`func (o *InfinibandFabric) GetReservedPkeys() []string`

GetReservedPkeys returns the ReservedPkeys field if non-nil, zero value otherwise.

### GetReservedPkeysOk

`func (o *InfinibandFabric) GetReservedPkeysOk() (*[]string, bool)`

GetReservedPkeysOk returns a tuple with the ReservedPkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedPkeys

`func (o *InfinibandFabric) SetReservedPkeys(v []string)`

SetReservedPkeys sets ReservedPkeys field to given value.

### HasReservedPkeys

`func (o *InfinibandFabric) HasReservedPkeys() bool`

HasReservedPkeys returns a boolean if a field has been set.

### GetPkeyRanges

`func (o *InfinibandFabric) GetPkeyRanges() []string`

GetPkeyRanges returns the PkeyRanges field if non-nil, zero value otherwise.

### GetPkeyRangesOk

`func (o *InfinibandFabric) GetPkeyRangesOk() (*[]string, bool)`

GetPkeyRangesOk returns a tuple with the PkeyRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeyRanges

`func (o *InfinibandFabric) SetPkeyRanges(v []string)`

SetPkeyRanges sets PkeyRanges field to given value.

### HasPkeyRanges

`func (o *InfinibandFabric) HasPkeyRanges() bool`

HasPkeyRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



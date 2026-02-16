# FibreChannelFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | [**FabricType**](FabricType.md) | The type of network fabric | 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**ServerOnlyOperationEnabled** | Pointer to **bool** | Enables server-only operation mode on the network fabric. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 

## Methods

### NewFibreChannelFabric

`func NewFibreChannelFabric(fabricType FabricType, ) *FibreChannelFabric`

NewFibreChannelFabric instantiates a new FibreChannelFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFibreChannelFabricWithDefaults

`func NewFibreChannelFabricWithDefaults() *FibreChannelFabric`

NewFibreChannelFabricWithDefaults instantiates a new FibreChannelFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFabricType

`func (o *FibreChannelFabric) GetFabricType() FabricType`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *FibreChannelFabric) GetFabricTypeOk() (*FabricType, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *FibreChannelFabric) SetFabricType(v FabricType)`

SetFabricType sets FabricType field to given value.


### GetGnmiMonitoringEnabled

`func (o *FibreChannelFabric) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *FibreChannelFabric) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *FibreChannelFabric) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *FibreChannelFabric) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetServerOnlyOperationEnabled

`func (o *FibreChannelFabric) GetServerOnlyOperationEnabled() bool`

GetServerOnlyOperationEnabled returns the ServerOnlyOperationEnabled field if non-nil, zero value otherwise.

### GetServerOnlyOperationEnabledOk

`func (o *FibreChannelFabric) GetServerOnlyOperationEnabledOk() (*bool, bool)`

GetServerOnlyOperationEnabledOk returns a tuple with the ServerOnlyOperationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOnlyOperationEnabled

`func (o *FibreChannelFabric) SetServerOnlyOperationEnabled(v bool)`

SetServerOnlyOperationEnabled sets ServerOnlyOperationEnabled field to given value.

### HasServerOnlyOperationEnabled

`func (o *FibreChannelFabric) HasServerOnlyOperationEnabled() bool`

HasServerOnlyOperationEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *FibreChannelFabric) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *FibreChannelFabric) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *FibreChannelFabric) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *FibreChannelFabric) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



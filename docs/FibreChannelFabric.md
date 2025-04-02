# FibreChannelFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | **string** | The type of network fabric | 
**DefaultNetworkProfileId** | Pointer to **int32** | Unique identifier for the default network profile. Must be a positive integer (minimum: 1) corresponding to an existing profile. | [optional] 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 
**ZeroTouchEnabled** | Pointer to **bool** | Enables zero-touch provisioning for automatic device configuration. | [optional] 
**VsanId** | Pointer to **int32** | VSAN ID for the Fibre Channel fabric | [optional] 
**TopologyType** | **string** | Fabric topology type | 
**Mtu** | Pointer to **float32** | Maximum transmission unit (MTU) size in bytes | [optional] 
**ZoningConfiguration** | Pointer to **map[string]interface{}** | Zoning configuration for the fabric | [optional] 
**InteropMode** | Pointer to **string** | Interoperability mode for multi-vendor environments | [optional] 
**QosConfiguration** | Pointer to **map[string]interface{}** | Quality of Service (QoS) configuration | [optional] 
**TrunkingConfiguration** | Pointer to **map[string]interface{}** | Trunking configuration for ISLs (Inter-Switch Links) | [optional] 
**PortChannelConfiguration** | Pointer to **map[string]interface{}** | Port channel configuration for ISLs | [optional] 

## Methods

### NewFibreChannelFabric

`func NewFibreChannelFabric(fabricType string, topologyType string, ) *FibreChannelFabric`

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

`func (o *FibreChannelFabric) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *FibreChannelFabric) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *FibreChannelFabric) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.


### GetDefaultNetworkProfileId

`func (o *FibreChannelFabric) GetDefaultNetworkProfileId() int32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *FibreChannelFabric) GetDefaultNetworkProfileIdOk() (*int32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *FibreChannelFabric) SetDefaultNetworkProfileId(v int32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.

### HasDefaultNetworkProfileId

`func (o *FibreChannelFabric) HasDefaultNetworkProfileId() bool`

HasDefaultNetworkProfileId returns a boolean if a field has been set.

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

### GetZeroTouchEnabled

`func (o *FibreChannelFabric) GetZeroTouchEnabled() bool`

GetZeroTouchEnabled returns the ZeroTouchEnabled field if non-nil, zero value otherwise.

### GetZeroTouchEnabledOk

`func (o *FibreChannelFabric) GetZeroTouchEnabledOk() (*bool, bool)`

GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchEnabled

`func (o *FibreChannelFabric) SetZeroTouchEnabled(v bool)`

SetZeroTouchEnabled sets ZeroTouchEnabled field to given value.

### HasZeroTouchEnabled

`func (o *FibreChannelFabric) HasZeroTouchEnabled() bool`

HasZeroTouchEnabled returns a boolean if a field has been set.

### GetVsanId

`func (o *FibreChannelFabric) GetVsanId() int32`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FibreChannelFabric) GetVsanIdOk() (*int32, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FibreChannelFabric) SetVsanId(v int32)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FibreChannelFabric) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.

### GetTopologyType

`func (o *FibreChannelFabric) GetTopologyType() string`

GetTopologyType returns the TopologyType field if non-nil, zero value otherwise.

### GetTopologyTypeOk

`func (o *FibreChannelFabric) GetTopologyTypeOk() (*string, bool)`

GetTopologyTypeOk returns a tuple with the TopologyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyType

`func (o *FibreChannelFabric) SetTopologyType(v string)`

SetTopologyType sets TopologyType field to given value.


### GetMtu

`func (o *FibreChannelFabric) GetMtu() float32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *FibreChannelFabric) GetMtuOk() (*float32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *FibreChannelFabric) SetMtu(v float32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *FibreChannelFabric) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetZoningConfiguration

`func (o *FibreChannelFabric) GetZoningConfiguration() map[string]interface{}`

GetZoningConfiguration returns the ZoningConfiguration field if non-nil, zero value otherwise.

### GetZoningConfigurationOk

`func (o *FibreChannelFabric) GetZoningConfigurationOk() (*map[string]interface{}, bool)`

GetZoningConfigurationOk returns a tuple with the ZoningConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoningConfiguration

`func (o *FibreChannelFabric) SetZoningConfiguration(v map[string]interface{})`

SetZoningConfiguration sets ZoningConfiguration field to given value.

### HasZoningConfiguration

`func (o *FibreChannelFabric) HasZoningConfiguration() bool`

HasZoningConfiguration returns a boolean if a field has been set.

### GetInteropMode

`func (o *FibreChannelFabric) GetInteropMode() string`

GetInteropMode returns the InteropMode field if non-nil, zero value otherwise.

### GetInteropModeOk

`func (o *FibreChannelFabric) GetInteropModeOk() (*string, bool)`

GetInteropModeOk returns a tuple with the InteropMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteropMode

`func (o *FibreChannelFabric) SetInteropMode(v string)`

SetInteropMode sets InteropMode field to given value.

### HasInteropMode

`func (o *FibreChannelFabric) HasInteropMode() bool`

HasInteropMode returns a boolean if a field has been set.

### GetQosConfiguration

`func (o *FibreChannelFabric) GetQosConfiguration() map[string]interface{}`

GetQosConfiguration returns the QosConfiguration field if non-nil, zero value otherwise.

### GetQosConfigurationOk

`func (o *FibreChannelFabric) GetQosConfigurationOk() (*map[string]interface{}, bool)`

GetQosConfigurationOk returns a tuple with the QosConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfiguration

`func (o *FibreChannelFabric) SetQosConfiguration(v map[string]interface{})`

SetQosConfiguration sets QosConfiguration field to given value.

### HasQosConfiguration

`func (o *FibreChannelFabric) HasQosConfiguration() bool`

HasQosConfiguration returns a boolean if a field has been set.

### GetTrunkingConfiguration

`func (o *FibreChannelFabric) GetTrunkingConfiguration() map[string]interface{}`

GetTrunkingConfiguration returns the TrunkingConfiguration field if non-nil, zero value otherwise.

### GetTrunkingConfigurationOk

`func (o *FibreChannelFabric) GetTrunkingConfigurationOk() (*map[string]interface{}, bool)`

GetTrunkingConfigurationOk returns a tuple with the TrunkingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkingConfiguration

`func (o *FibreChannelFabric) SetTrunkingConfiguration(v map[string]interface{})`

SetTrunkingConfiguration sets TrunkingConfiguration field to given value.

### HasTrunkingConfiguration

`func (o *FibreChannelFabric) HasTrunkingConfiguration() bool`

HasTrunkingConfiguration returns a boolean if a field has been set.

### GetPortChannelConfiguration

`func (o *FibreChannelFabric) GetPortChannelConfiguration() map[string]interface{}`

GetPortChannelConfiguration returns the PortChannelConfiguration field if non-nil, zero value otherwise.

### GetPortChannelConfigurationOk

`func (o *FibreChannelFabric) GetPortChannelConfigurationOk() (*map[string]interface{}, bool)`

GetPortChannelConfigurationOk returns a tuple with the PortChannelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelConfiguration

`func (o *FibreChannelFabric) SetPortChannelConfiguration(v map[string]interface{})`

SetPortChannelConfiguration sets PortChannelConfiguration field to given value.

### HasPortChannelConfiguration

`func (o *FibreChannelFabric) HasPortChannelConfiguration() bool`

HasPortChannelConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



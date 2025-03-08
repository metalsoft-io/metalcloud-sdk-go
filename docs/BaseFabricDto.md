# BaseFabricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricType** | **string** | The type of network fabric | 
**DefaultNetworkProfileId** | **int32** | Unique identifier for the default network profile. Must be a positive integer (minimum: 1) corresponding to an existing profile. | 
**GnmiMonitoringEnabled** | Pointer to **bool** | Enables gNMI monitoring for telemetry data collection using the gNMI protocol. | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Enables syslog monitoring for capturing system logs for diagnostics and troubleshooting. | [optional] 
**ZeroTouchEnabled** | Pointer to **bool** | Enables zero-touch provisioning for automatic device configuration. | [optional] 

## Methods

### NewBaseFabricDto

`func NewBaseFabricDto(fabricType string, defaultNetworkProfileId int32, ) *BaseFabricDto`

NewBaseFabricDto instantiates a new BaseFabricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFabricDtoWithDefaults

`func NewBaseFabricDtoWithDefaults() *BaseFabricDto`

NewBaseFabricDtoWithDefaults instantiates a new BaseFabricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFabricType

`func (o *BaseFabricDto) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *BaseFabricDto) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *BaseFabricDto) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.


### GetDefaultNetworkProfileId

`func (o *BaseFabricDto) GetDefaultNetworkProfileId() int32`

GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field if non-nil, zero value otherwise.

### GetDefaultNetworkProfileIdOk

`func (o *BaseFabricDto) GetDefaultNetworkProfileIdOk() (*int32, bool)`

GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkProfileId

`func (o *BaseFabricDto) SetDefaultNetworkProfileId(v int32)`

SetDefaultNetworkProfileId sets DefaultNetworkProfileId field to given value.


### GetGnmiMonitoringEnabled

`func (o *BaseFabricDto) GetGnmiMonitoringEnabled() bool`

GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field if non-nil, zero value otherwise.

### GetGnmiMonitoringEnabledOk

`func (o *BaseFabricDto) GetGnmiMonitoringEnabledOk() (*bool, bool)`

GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnmiMonitoringEnabled

`func (o *BaseFabricDto) SetGnmiMonitoringEnabled(v bool)`

SetGnmiMonitoringEnabled sets GnmiMonitoringEnabled field to given value.

### HasGnmiMonitoringEnabled

`func (o *BaseFabricDto) HasGnmiMonitoringEnabled() bool`

HasGnmiMonitoringEnabled returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *BaseFabricDto) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *BaseFabricDto) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *BaseFabricDto) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *BaseFabricDto) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetZeroTouchEnabled

`func (o *BaseFabricDto) GetZeroTouchEnabled() bool`

GetZeroTouchEnabled returns the ZeroTouchEnabled field if non-nil, zero value otherwise.

### GetZeroTouchEnabledOk

`func (o *BaseFabricDto) GetZeroTouchEnabledOk() (*bool, bool)`

GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroTouchEnabled

`func (o *BaseFabricDto) SetZeroTouchEnabled(v bool)`

SetZeroTouchEnabled sets ZeroTouchEnabled field to given value.

### HasZeroTouchEnabled

`func (o *BaseFabricDto) HasZeroTouchEnabled() bool`

HasZeroTouchEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TunnelSyslogFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPriorities** | Pointer to **[]float32** | Allowed syslog priority values (0–191) | [optional] 
**AllowedFacilities** | Pointer to **[]float32** | Allowed syslog facility values (0–23) | [optional] 
**AllowedSeverities** | Pointer to **[]float32** | Allowed syslog severity values (0–7) | [optional] 

## Methods

### NewTunnelSyslogFilters

`func NewTunnelSyslogFilters() *TunnelSyslogFilters`

NewTunnelSyslogFilters instantiates a new TunnelSyslogFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelSyslogFiltersWithDefaults

`func NewTunnelSyslogFiltersWithDefaults() *TunnelSyslogFilters`

NewTunnelSyslogFiltersWithDefaults instantiates a new TunnelSyslogFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPriorities

`func (o *TunnelSyslogFilters) GetAllowedPriorities() []float32`

GetAllowedPriorities returns the AllowedPriorities field if non-nil, zero value otherwise.

### GetAllowedPrioritiesOk

`func (o *TunnelSyslogFilters) GetAllowedPrioritiesOk() (*[]float32, bool)`

GetAllowedPrioritiesOk returns a tuple with the AllowedPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPriorities

`func (o *TunnelSyslogFilters) SetAllowedPriorities(v []float32)`

SetAllowedPriorities sets AllowedPriorities field to given value.

### HasAllowedPriorities

`func (o *TunnelSyslogFilters) HasAllowedPriorities() bool`

HasAllowedPriorities returns a boolean if a field has been set.

### GetAllowedFacilities

`func (o *TunnelSyslogFilters) GetAllowedFacilities() []float32`

GetAllowedFacilities returns the AllowedFacilities field if non-nil, zero value otherwise.

### GetAllowedFacilitiesOk

`func (o *TunnelSyslogFilters) GetAllowedFacilitiesOk() (*[]float32, bool)`

GetAllowedFacilitiesOk returns a tuple with the AllowedFacilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFacilities

`func (o *TunnelSyslogFilters) SetAllowedFacilities(v []float32)`

SetAllowedFacilities sets AllowedFacilities field to given value.

### HasAllowedFacilities

`func (o *TunnelSyslogFilters) HasAllowedFacilities() bool`

HasAllowedFacilities returns a boolean if a field has been set.

### GetAllowedSeverities

`func (o *TunnelSyslogFilters) GetAllowedSeverities() []float32`

GetAllowedSeverities returns the AllowedSeverities field if non-nil, zero value otherwise.

### GetAllowedSeveritiesOk

`func (o *TunnelSyslogFilters) GetAllowedSeveritiesOk() (*[]float32, bool)`

GetAllowedSeveritiesOk returns a tuple with the AllowedSeverities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSeverities

`func (o *TunnelSyslogFilters) SetAllowedSeverities(v []float32)`

SetAllowedSeverities sets AllowedSeverities field to given value.

### HasAllowedSeverities

`func (o *TunnelSyslogFilters) HasAllowedSeverities() bool`

HasAllowedSeverities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



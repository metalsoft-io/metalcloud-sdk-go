# TunnelSyslogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**TunnelSyslogFiltersDto**](TunnelSyslogFiltersDto.md) | Syslog message filters applied by the tunnel | 

## Methods

### NewTunnelSyslogDto

`func NewTunnelSyslogDto(filters TunnelSyslogFiltersDto, ) *TunnelSyslogDto`

NewTunnelSyslogDto instantiates a new TunnelSyslogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelSyslogDtoWithDefaults

`func NewTunnelSyslogDtoWithDefaults() *TunnelSyslogDto`

NewTunnelSyslogDtoWithDefaults instantiates a new TunnelSyslogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *TunnelSyslogDto) GetFilters() TunnelSyslogFiltersDto`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TunnelSyslogDto) GetFiltersOk() (*TunnelSyslogFiltersDto, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TunnelSyslogDto) SetFilters(v TunnelSyslogFiltersDto)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



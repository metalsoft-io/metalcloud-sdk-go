# TunnelSyslog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**TunnelSyslogFilters**](TunnelSyslogFilters.md) | Syslog message filters applied by the tunnel | 

## Methods

### NewTunnelSyslog

`func NewTunnelSyslog(filters TunnelSyslogFilters, ) *TunnelSyslog`

NewTunnelSyslog instantiates a new TunnelSyslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelSyslogWithDefaults

`func NewTunnelSyslogWithDefaults() *TunnelSyslog`

NewTunnelSyslogWithDefaults instantiates a new TunnelSyslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *TunnelSyslog) GetFilters() TunnelSyslogFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TunnelSyslog) GetFiltersOk() (*TunnelSyslogFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TunnelSyslog) SetFilters(v TunnelSyslogFilters)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



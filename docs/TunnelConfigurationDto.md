# TunnelConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | **string** | Shared secret used for tunnel authentication | 
**Bdk** | [**TunnelBdkDto**](TunnelBdkDto.md) | BDK service credentials | 
**Syslog** | [**TunnelSyslogDto**](TunnelSyslogDto.md) | Syslog message filtering configuration | 

## Methods

### NewTunnelConfigurationDto

`func NewTunnelConfigurationDto(sharedSecret string, bdk TunnelBdkDto, syslog TunnelSyslogDto, ) *TunnelConfigurationDto`

NewTunnelConfigurationDto instantiates a new TunnelConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfigurationDtoWithDefaults

`func NewTunnelConfigurationDtoWithDefaults() *TunnelConfigurationDto`

NewTunnelConfigurationDtoWithDefaults instantiates a new TunnelConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *TunnelConfigurationDto) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *TunnelConfigurationDto) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *TunnelConfigurationDto) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetBdk

`func (o *TunnelConfigurationDto) GetBdk() TunnelBdkDto`

GetBdk returns the Bdk field if non-nil, zero value otherwise.

### GetBdkOk

`func (o *TunnelConfigurationDto) GetBdkOk() (*TunnelBdkDto, bool)`

GetBdkOk returns a tuple with the Bdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdk

`func (o *TunnelConfigurationDto) SetBdk(v TunnelBdkDto)`

SetBdk sets Bdk field to given value.


### GetSyslog

`func (o *TunnelConfigurationDto) GetSyslog() TunnelSyslogDto`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *TunnelConfigurationDto) GetSyslogOk() (*TunnelSyslogDto, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *TunnelConfigurationDto) SetSyslog(v TunnelSyslogDto)`

SetSyslog sets Syslog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



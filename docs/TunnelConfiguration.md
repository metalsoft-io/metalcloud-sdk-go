# TunnelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | **string** | Shared secret used for tunnel authentication | 
**Bdk** | [**TunnelBdk**](TunnelBdk.md) | BDK service credentials | 
**Syslog** | [**TunnelSyslog**](TunnelSyslog.md) | Syslog message filtering configuration | 

## Methods

### NewTunnelConfiguration

`func NewTunnelConfiguration(sharedSecret string, bdk TunnelBdk, syslog TunnelSyslog, ) *TunnelConfiguration`

NewTunnelConfiguration instantiates a new TunnelConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelConfigurationWithDefaults

`func NewTunnelConfigurationWithDefaults() *TunnelConfiguration`

NewTunnelConfigurationWithDefaults instantiates a new TunnelConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *TunnelConfiguration) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *TunnelConfiguration) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *TunnelConfiguration) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetBdk

`func (o *TunnelConfiguration) GetBdk() TunnelBdk`

GetBdk returns the Bdk field if non-nil, zero value otherwise.

### GetBdkOk

`func (o *TunnelConfiguration) GetBdkOk() (*TunnelBdk, bool)`

GetBdkOk returns a tuple with the Bdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdk

`func (o *TunnelConfiguration) SetBdk(v TunnelBdk)`

SetBdk sets Bdk field to given value.


### GetSyslog

`func (o *TunnelConfiguration) GetSyslog() TunnelSyslog`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *TunnelConfiguration) GetSyslogOk() (*TunnelSyslog, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *TunnelConfiguration) SetSyslog(v TunnelSyslog)`

SetSyslog sets Syslog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



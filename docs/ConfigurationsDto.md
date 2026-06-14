# ConfigurationsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**AuthConfigurationDto**](AuthConfigurationDto.md) |  | [optional] 
**Platform** | Pointer to [**PlatformConfigurationDto**](PlatformConfigurationDto.md) |  | [optional] 
**Notification** | Pointer to [**NotificationConfigurationDto**](NotificationConfigurationDto.md) |  | [optional] 
**Tunnel** | Pointer to [**TunnelConfigurationDto**](TunnelConfigurationDto.md) |  | [optional] 
**GatewayApi** | Pointer to [**GatewayConfigurationDto**](GatewayConfigurationDto.md) |  | [optional] 
**ImageBuilder** | Pointer to [**ImageBuilderConfigurationDto**](ImageBuilderConfigurationDto.md) |  | [optional] 

## Methods

### NewConfigurationsDto

`func NewConfigurationsDto() *ConfigurationsDto`

NewConfigurationsDto instantiates a new ConfigurationsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationsDtoWithDefaults

`func NewConfigurationsDtoWithDefaults() *ConfigurationsDto`

NewConfigurationsDtoWithDefaults instantiates a new ConfigurationsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *ConfigurationsDto) GetAuth() AuthConfigurationDto`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ConfigurationsDto) GetAuthOk() (*AuthConfigurationDto, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ConfigurationsDto) SetAuth(v AuthConfigurationDto)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ConfigurationsDto) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetPlatform

`func (o *ConfigurationsDto) GetPlatform() PlatformConfigurationDto`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ConfigurationsDto) GetPlatformOk() (*PlatformConfigurationDto, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ConfigurationsDto) SetPlatform(v PlatformConfigurationDto)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ConfigurationsDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetNotification

`func (o *ConfigurationsDto) GetNotification() NotificationConfigurationDto`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *ConfigurationsDto) GetNotificationOk() (*NotificationConfigurationDto, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *ConfigurationsDto) SetNotification(v NotificationConfigurationDto)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *ConfigurationsDto) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetTunnel

`func (o *ConfigurationsDto) GetTunnel() TunnelConfigurationDto`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *ConfigurationsDto) GetTunnelOk() (*TunnelConfigurationDto, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *ConfigurationsDto) SetTunnel(v TunnelConfigurationDto)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *ConfigurationsDto) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetGatewayApi

`func (o *ConfigurationsDto) GetGatewayApi() GatewayConfigurationDto`

GetGatewayApi returns the GatewayApi field if non-nil, zero value otherwise.

### GetGatewayApiOk

`func (o *ConfigurationsDto) GetGatewayApiOk() (*GatewayConfigurationDto, bool)`

GetGatewayApiOk returns a tuple with the GatewayApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayApi

`func (o *ConfigurationsDto) SetGatewayApi(v GatewayConfigurationDto)`

SetGatewayApi sets GatewayApi field to given value.

### HasGatewayApi

`func (o *ConfigurationsDto) HasGatewayApi() bool`

HasGatewayApi returns a boolean if a field has been set.

### GetImageBuilder

`func (o *ConfigurationsDto) GetImageBuilder() ImageBuilderConfigurationDto`

GetImageBuilder returns the ImageBuilder field if non-nil, zero value otherwise.

### GetImageBuilderOk

`func (o *ConfigurationsDto) GetImageBuilderOk() (*ImageBuilderConfigurationDto, bool)`

GetImageBuilderOk returns a tuple with the ImageBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuilder

`func (o *ConfigurationsDto) SetImageBuilder(v ImageBuilderConfigurationDto)`

SetImageBuilder sets ImageBuilder field to given value.

### HasImageBuilder

`func (o *ConfigurationsDto) HasImageBuilder() bool`

HasImageBuilder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



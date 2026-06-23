# Configurations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**AuthConfiguration**](AuthConfiguration.md) |  | [optional] 
**Platform** | Pointer to [**PlatformConfiguration**](PlatformConfiguration.md) |  | [optional] 
**Notification** | Pointer to [**NotificationConfiguration**](NotificationConfiguration.md) |  | [optional] 
**Tunnel** | Pointer to [**TunnelConfiguration**](TunnelConfiguration.md) |  | [optional] 
**GatewayApi** | Pointer to [**GatewayConfiguration**](GatewayConfiguration.md) |  | [optional] 
**ImageBuilder** | Pointer to [**ImageBuilderConfiguration**](ImageBuilderConfiguration.md) |  | [optional] 

## Methods

### NewConfigurations

`func NewConfigurations() *Configurations`

NewConfigurations instantiates a new Configurations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationsWithDefaults

`func NewConfigurationsWithDefaults() *Configurations`

NewConfigurationsWithDefaults instantiates a new Configurations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *Configurations) GetAuth() AuthConfiguration`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Configurations) GetAuthOk() (*AuthConfiguration, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Configurations) SetAuth(v AuthConfiguration)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Configurations) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetPlatform

`func (o *Configurations) GetPlatform() PlatformConfiguration`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Configurations) GetPlatformOk() (*PlatformConfiguration, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Configurations) SetPlatform(v PlatformConfiguration)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Configurations) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetNotification

`func (o *Configurations) GetNotification() NotificationConfiguration`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *Configurations) GetNotificationOk() (*NotificationConfiguration, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *Configurations) SetNotification(v NotificationConfiguration)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *Configurations) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetTunnel

`func (o *Configurations) GetTunnel() TunnelConfiguration`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *Configurations) GetTunnelOk() (*TunnelConfiguration, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *Configurations) SetTunnel(v TunnelConfiguration)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *Configurations) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetGatewayApi

`func (o *Configurations) GetGatewayApi() GatewayConfiguration`

GetGatewayApi returns the GatewayApi field if non-nil, zero value otherwise.

### GetGatewayApiOk

`func (o *Configurations) GetGatewayApiOk() (*GatewayConfiguration, bool)`

GetGatewayApiOk returns a tuple with the GatewayApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayApi

`func (o *Configurations) SetGatewayApi(v GatewayConfiguration)`

SetGatewayApi sets GatewayApi field to given value.

### HasGatewayApi

`func (o *Configurations) HasGatewayApi() bool`

HasGatewayApi returns a boolean if a field has been set.

### GetImageBuilder

`func (o *Configurations) GetImageBuilder() ImageBuilderConfiguration`

GetImageBuilder returns the ImageBuilder field if non-nil, zero value otherwise.

### GetImageBuilderOk

`func (o *Configurations) GetImageBuilderOk() (*ImageBuilderConfiguration, bool)`

GetImageBuilderOk returns a tuple with the ImageBuilder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuilder

`func (o *Configurations) SetImageBuilder(v ImageBuilderConfiguration)`

SetImageBuilder sets ImageBuilder field to given value.

### HasImageBuilder

`func (o *Configurations) HasImageBuilder() bool`

HasImageBuilder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



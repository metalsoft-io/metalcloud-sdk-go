# UpdateLogicalNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | **map[string]interface{}** |  | 
**Vxlan** | **map[string]interface{}** |  | 
**Ipv4** | **map[string]interface{}** |  | 
**Ipv6** | **map[string]interface{}** |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewUpdateLogicalNetworkConfig

`func NewUpdateLogicalNetworkConfig(vlan map[string]interface{}, vxlan map[string]interface{}, ipv4 map[string]interface{}, ipv6 map[string]interface{}, ) *UpdateLogicalNetworkConfig`

NewUpdateLogicalNetworkConfig instantiates a new UpdateLogicalNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLogicalNetworkConfigWithDefaults

`func NewUpdateLogicalNetworkConfigWithDefaults() *UpdateLogicalNetworkConfig`

NewUpdateLogicalNetworkConfigWithDefaults instantiates a new UpdateLogicalNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *UpdateLogicalNetworkConfig) GetVlan() map[string]interface{}`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UpdateLogicalNetworkConfig) GetVlanOk() (*map[string]interface{}, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UpdateLogicalNetworkConfig) SetVlan(v map[string]interface{})`

SetVlan sets Vlan field to given value.


### GetVxlan

`func (o *UpdateLogicalNetworkConfig) GetVxlan() map[string]interface{}`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *UpdateLogicalNetworkConfig) GetVxlanOk() (*map[string]interface{}, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *UpdateLogicalNetworkConfig) SetVxlan(v map[string]interface{})`

SetVxlan sets Vxlan field to given value.


### GetIpv4

`func (o *UpdateLogicalNetworkConfig) GetIpv4() map[string]interface{}`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *UpdateLogicalNetworkConfig) GetIpv4Ok() (*map[string]interface{}, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *UpdateLogicalNetworkConfig) SetIpv4(v map[string]interface{})`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *UpdateLogicalNetworkConfig) GetIpv6() map[string]interface{}`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UpdateLogicalNetworkConfig) GetIpv6Ok() (*map[string]interface{}, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UpdateLogicalNetworkConfig) SetIpv6(v map[string]interface{})`

SetIpv6 sets Ipv6 field to given value.


### GetMtu

`func (o *UpdateLogicalNetworkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *UpdateLogicalNetworkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *UpdateLogicalNetworkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *UpdateLogicalNetworkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *UpdateLogicalNetworkConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *UpdateLogicalNetworkConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworkEquipmentInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | **int64** |  | 
**SwitchId** | **int64** |  | 
**Kind** | **string** |  | 
**InterfaceName** | **string** |  | 
**ParentInterfaceId** | Pointer to **NullableInt64** |  | [optional] 
**InterfaceDescription** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Speed** | Pointer to **NullableString** |  | [optional] 
**AutoNegotiate** | Pointer to **NullableBool** |  | [optional] 
**ServiceStatus** | **string** |  | 
**PendingDelete** | **bool** |  | 
**Config** | [**NetworkEquipmentInterfaceConfig**](NetworkEquipmentInterfaceConfig.md) |  | 
**Ipv4** | [**NetworkEquipmentInterfaceIpFamily**](NetworkEquipmentInterfaceIpFamily.md) |  | 
**Ipv6** | [**NetworkEquipmentInterfaceIpFamily**](NetworkEquipmentInterfaceIpFamily.md) |  | 

## Methods

### NewNetworkEquipmentInterface

`func NewNetworkEquipmentInterface(interfaceId int64, switchId int64, kind string, interfaceName string, serviceStatus string, pendingDelete bool, config NetworkEquipmentInterfaceConfig, ipv4 NetworkEquipmentInterfaceIpFamily, ipv6 NetworkEquipmentInterfaceIpFamily, ) *NetworkEquipmentInterface`

NewNetworkEquipmentInterface instantiates a new NetworkEquipmentInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentInterfaceWithDefaults

`func NewNetworkEquipmentInterfaceWithDefaults() *NetworkEquipmentInterface`

NewNetworkEquipmentInterfaceWithDefaults instantiates a new NetworkEquipmentInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *NetworkEquipmentInterface) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkEquipmentInterface) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkEquipmentInterface) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetSwitchId

`func (o *NetworkEquipmentInterface) GetSwitchId() int64`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NetworkEquipmentInterface) GetSwitchIdOk() (*int64, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NetworkEquipmentInterface) SetSwitchId(v int64)`

SetSwitchId sets SwitchId field to given value.


### GetKind

`func (o *NetworkEquipmentInterface) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkEquipmentInterface) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkEquipmentInterface) SetKind(v string)`

SetKind sets Kind field to given value.


### GetInterfaceName

`func (o *NetworkEquipmentInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkEquipmentInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkEquipmentInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetParentInterfaceId

`func (o *NetworkEquipmentInterface) GetParentInterfaceId() int64`

GetParentInterfaceId returns the ParentInterfaceId field if non-nil, zero value otherwise.

### GetParentInterfaceIdOk

`func (o *NetworkEquipmentInterface) GetParentInterfaceIdOk() (*int64, bool)`

GetParentInterfaceIdOk returns a tuple with the ParentInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterfaceId

`func (o *NetworkEquipmentInterface) SetParentInterfaceId(v int64)`

SetParentInterfaceId sets ParentInterfaceId field to given value.

### HasParentInterfaceId

`func (o *NetworkEquipmentInterface) HasParentInterfaceId() bool`

HasParentInterfaceId returns a boolean if a field has been set.

### SetParentInterfaceIdNil

`func (o *NetworkEquipmentInterface) SetParentInterfaceIdNil(b bool)`

 SetParentInterfaceIdNil sets the value for ParentInterfaceId to be an explicit nil

### UnsetParentInterfaceId
`func (o *NetworkEquipmentInterface) UnsetParentInterfaceId()`

UnsetParentInterfaceId ensures that no value is present for ParentInterfaceId, not even an explicit nil
### GetInterfaceDescription

`func (o *NetworkEquipmentInterface) GetInterfaceDescription() string`

GetInterfaceDescription returns the InterfaceDescription field if non-nil, zero value otherwise.

### GetInterfaceDescriptionOk

`func (o *NetworkEquipmentInterface) GetInterfaceDescriptionOk() (*string, bool)`

GetInterfaceDescriptionOk returns a tuple with the InterfaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceDescription

`func (o *NetworkEquipmentInterface) SetInterfaceDescription(v string)`

SetInterfaceDescription sets InterfaceDescription field to given value.

### HasInterfaceDescription

`func (o *NetworkEquipmentInterface) HasInterfaceDescription() bool`

HasInterfaceDescription returns a boolean if a field has been set.

### GetMtu

`func (o *NetworkEquipmentInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkEquipmentInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkEquipmentInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkEquipmentInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkEquipmentInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkEquipmentInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkEquipmentInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkEquipmentInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSpeed

`func (o *NetworkEquipmentInterface) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NetworkEquipmentInterface) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NetworkEquipmentInterface) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NetworkEquipmentInterface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *NetworkEquipmentInterface) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *NetworkEquipmentInterface) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetAutoNegotiate

`func (o *NetworkEquipmentInterface) GetAutoNegotiate() bool`

GetAutoNegotiate returns the AutoNegotiate field if non-nil, zero value otherwise.

### GetAutoNegotiateOk

`func (o *NetworkEquipmentInterface) GetAutoNegotiateOk() (*bool, bool)`

GetAutoNegotiateOk returns a tuple with the AutoNegotiate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiate

`func (o *NetworkEquipmentInterface) SetAutoNegotiate(v bool)`

SetAutoNegotiate sets AutoNegotiate field to given value.

### HasAutoNegotiate

`func (o *NetworkEquipmentInterface) HasAutoNegotiate() bool`

HasAutoNegotiate returns a boolean if a field has been set.

### SetAutoNegotiateNil

`func (o *NetworkEquipmentInterface) SetAutoNegotiateNil(b bool)`

 SetAutoNegotiateNil sets the value for AutoNegotiate to be an explicit nil

### UnsetAutoNegotiate
`func (o *NetworkEquipmentInterface) UnsetAutoNegotiate()`

UnsetAutoNegotiate ensures that no value is present for AutoNegotiate, not even an explicit nil
### GetServiceStatus

`func (o *NetworkEquipmentInterface) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *NetworkEquipmentInterface) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *NetworkEquipmentInterface) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetPendingDelete

`func (o *NetworkEquipmentInterface) GetPendingDelete() bool`

GetPendingDelete returns the PendingDelete field if non-nil, zero value otherwise.

### GetPendingDeleteOk

`func (o *NetworkEquipmentInterface) GetPendingDeleteOk() (*bool, bool)`

GetPendingDeleteOk returns a tuple with the PendingDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDelete

`func (o *NetworkEquipmentInterface) SetPendingDelete(v bool)`

SetPendingDelete sets PendingDelete field to given value.


### GetConfig

`func (o *NetworkEquipmentInterface) GetConfig() NetworkEquipmentInterfaceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkEquipmentInterface) GetConfigOk() (*NetworkEquipmentInterfaceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkEquipmentInterface) SetConfig(v NetworkEquipmentInterfaceConfig)`

SetConfig sets Config field to given value.


### GetIpv4

`func (o *NetworkEquipmentInterface) GetIpv4() NetworkEquipmentInterfaceIpFamily`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *NetworkEquipmentInterface) GetIpv4Ok() (*NetworkEquipmentInterfaceIpFamily, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *NetworkEquipmentInterface) SetIpv4(v NetworkEquipmentInterfaceIpFamily)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *NetworkEquipmentInterface) GetIpv6() NetworkEquipmentInterfaceIpFamily`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *NetworkEquipmentInterface) GetIpv6Ok() (*NetworkEquipmentInterfaceIpFamily, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *NetworkEquipmentInterface) SetIpv6(v NetworkEquipmentInterfaceIpFamily)`

SetIpv6 sets Ipv6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



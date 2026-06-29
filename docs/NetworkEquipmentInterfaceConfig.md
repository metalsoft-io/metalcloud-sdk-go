# NetworkEquipmentInterfaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | L1 admin override for interface description. | [optional] 
**Mtu** | Pointer to **NullableInt32** | L1 admin override for interface MTU. | [optional] 
**Enabled** | Pointer to **NullableBool** | L1 admin override for admin enabled (link up/down). | [optional] 
**Speed** | Pointer to **NullableString** | L1 forced port speed. NULL &#x3D; no opinion; the device’s current forced speed is left unchanged. Forced speed only — autonegotiation is controlled by autoNegotiate. | [optional] 
**AutoNegotiate** | Pointer to **NullableBool** | L1 autonegotiation intent. NULL &#x3D; no opinion (leave as-is); true &#x3D; negotiate (advertise all); false &#x3D; force the speed from &#x60;speed&#x60;. | [optional] 
**Revision** | **int64** | Optimistic-lock revision of the config buffer (independent of the main interface revision). | 

## Methods

### NewNetworkEquipmentInterfaceConfig

`func NewNetworkEquipmentInterfaceConfig(revision int64, ) *NetworkEquipmentInterfaceConfig`

NewNetworkEquipmentInterfaceConfig instantiates a new NetworkEquipmentInterfaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentInterfaceConfigWithDefaults

`func NewNetworkEquipmentInterfaceConfigWithDefaults() *NetworkEquipmentInterfaceConfig`

NewNetworkEquipmentInterfaceConfigWithDefaults instantiates a new NetworkEquipmentInterfaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkEquipmentInterfaceConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkEquipmentInterfaceConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkEquipmentInterfaceConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkEquipmentInterfaceConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkEquipmentInterfaceConfig) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkEquipmentInterfaceConfig) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMtu

`func (o *NetworkEquipmentInterfaceConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkEquipmentInterfaceConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkEquipmentInterfaceConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkEquipmentInterfaceConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *NetworkEquipmentInterfaceConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *NetworkEquipmentInterfaceConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetEnabled

`func (o *NetworkEquipmentInterfaceConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkEquipmentInterfaceConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkEquipmentInterfaceConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkEquipmentInterfaceConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *NetworkEquipmentInterfaceConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *NetworkEquipmentInterfaceConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetSpeed

`func (o *NetworkEquipmentInterfaceConfig) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NetworkEquipmentInterfaceConfig) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NetworkEquipmentInterfaceConfig) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NetworkEquipmentInterfaceConfig) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *NetworkEquipmentInterfaceConfig) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *NetworkEquipmentInterfaceConfig) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetAutoNegotiate

`func (o *NetworkEquipmentInterfaceConfig) GetAutoNegotiate() bool`

GetAutoNegotiate returns the AutoNegotiate field if non-nil, zero value otherwise.

### GetAutoNegotiateOk

`func (o *NetworkEquipmentInterfaceConfig) GetAutoNegotiateOk() (*bool, bool)`

GetAutoNegotiateOk returns a tuple with the AutoNegotiate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiate

`func (o *NetworkEquipmentInterfaceConfig) SetAutoNegotiate(v bool)`

SetAutoNegotiate sets AutoNegotiate field to given value.

### HasAutoNegotiate

`func (o *NetworkEquipmentInterfaceConfig) HasAutoNegotiate() bool`

HasAutoNegotiate returns a boolean if a field has been set.

### SetAutoNegotiateNil

`func (o *NetworkEquipmentInterfaceConfig) SetAutoNegotiateNil(b bool)`

 SetAutoNegotiateNil sets the value for AutoNegotiate to be an explicit nil

### UnsetAutoNegotiate
`func (o *NetworkEquipmentInterfaceConfig) UnsetAutoNegotiate()`

UnsetAutoNegotiate ensures that no value is present for AutoNegotiate, not even an explicit nil
### GetRevision

`func (o *NetworkEquipmentInterfaceConfig) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkEquipmentInterfaceConfig) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkEquipmentInterfaceConfig) SetRevision(v int64)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



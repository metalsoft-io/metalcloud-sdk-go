# CreateNetworkEquipmentInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind of interface to create. | 
**Name** | **string** | Device-side interface name (e.g., \&quot;lo\&quot;, \&quot;loopback0\&quot;, \&quot;Loopback1\&quot;). | 
**Description** | Pointer to **string** | Interface description. | [optional] 
**Mtu** | Pointer to **int32** | MTU. | [optional] 
**Enabled** | Pointer to **bool** | Admin enabled (link up/down). | [optional] 
**ParentInterfaceId** | Pointer to **int64** | Parent interface id (same switch). Required for kinds that ride a parent (e.g., sub_interface). Must be omitted for kinds that do not (e.g., loopback). | [optional] 

## Methods

### NewCreateNetworkEquipmentInterface

`func NewCreateNetworkEquipmentInterface(kind string, name string, ) *CreateNetworkEquipmentInterface`

NewCreateNetworkEquipmentInterface instantiates a new CreateNetworkEquipmentInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkEquipmentInterfaceWithDefaults

`func NewCreateNetworkEquipmentInterfaceWithDefaults() *CreateNetworkEquipmentInterface`

NewCreateNetworkEquipmentInterfaceWithDefaults instantiates a new CreateNetworkEquipmentInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateNetworkEquipmentInterface) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateNetworkEquipmentInterface) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateNetworkEquipmentInterface) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *CreateNetworkEquipmentInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkEquipmentInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkEquipmentInterface) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNetworkEquipmentInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkEquipmentInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkEquipmentInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkEquipmentInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMtu

`func (o *CreateNetworkEquipmentInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateNetworkEquipmentInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateNetworkEquipmentInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateNetworkEquipmentInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateNetworkEquipmentInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkEquipmentInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkEquipmentInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkEquipmentInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParentInterfaceId

`func (o *CreateNetworkEquipmentInterface) GetParentInterfaceId() int64`

GetParentInterfaceId returns the ParentInterfaceId field if non-nil, zero value otherwise.

### GetParentInterfaceIdOk

`func (o *CreateNetworkEquipmentInterface) GetParentInterfaceIdOk() (*int64, bool)`

GetParentInterfaceIdOk returns a tuple with the ParentInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInterfaceId

`func (o *CreateNetworkEquipmentInterface) SetParentInterfaceId(v int64)`

SetParentInterfaceId sets ParentInterfaceId field to given value.

### HasParentInterfaceId

`func (o *CreateNetworkEquipmentInterface) HasParentInterfaceId() bool`

HasParentInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



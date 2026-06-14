# NetworkEquipmentInterfaceIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**InterfaceId** | **int64** |  | 
**Kind** | **string** |  | 
**Address** | **string** |  | 
**PrefixLength** | **int32** |  | 
**ServiceStatus** | **string** |  | 
**PendingDelete** | **bool** |  | 
**Vrrp** | Pointer to [**map[string]VrrpGroup**](VrrpGroup.md) | VRRP groups keyed by virtual router id (1..255); null when none. | [optional] 

## Methods

### NewNetworkEquipmentInterfaceIp

`func NewNetworkEquipmentInterfaceIp(id int64, interfaceId int64, kind string, address string, prefixLength int32, serviceStatus string, pendingDelete bool, ) *NetworkEquipmentInterfaceIp`

NewNetworkEquipmentInterfaceIp instantiates a new NetworkEquipmentInterfaceIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEquipmentInterfaceIpWithDefaults

`func NewNetworkEquipmentInterfaceIpWithDefaults() *NetworkEquipmentInterfaceIp`

NewNetworkEquipmentInterfaceIpWithDefaults instantiates a new NetworkEquipmentInterfaceIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkEquipmentInterfaceIp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkEquipmentInterfaceIp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkEquipmentInterfaceIp) SetId(v int64)`

SetId sets Id field to given value.


### GetInterfaceId

`func (o *NetworkEquipmentInterfaceIp) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkEquipmentInterfaceIp) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkEquipmentInterfaceIp) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetKind

`func (o *NetworkEquipmentInterfaceIp) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkEquipmentInterfaceIp) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkEquipmentInterfaceIp) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAddress

`func (o *NetworkEquipmentInterfaceIp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkEquipmentInterfaceIp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkEquipmentInterfaceIp) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPrefixLength

`func (o *NetworkEquipmentInterfaceIp) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *NetworkEquipmentInterfaceIp) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *NetworkEquipmentInterfaceIp) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetServiceStatus

`func (o *NetworkEquipmentInterfaceIp) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *NetworkEquipmentInterfaceIp) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *NetworkEquipmentInterfaceIp) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetPendingDelete

`func (o *NetworkEquipmentInterfaceIp) GetPendingDelete() bool`

GetPendingDelete returns the PendingDelete field if non-nil, zero value otherwise.

### GetPendingDeleteOk

`func (o *NetworkEquipmentInterfaceIp) GetPendingDeleteOk() (*bool, bool)`

GetPendingDeleteOk returns a tuple with the PendingDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDelete

`func (o *NetworkEquipmentInterfaceIp) SetPendingDelete(v bool)`

SetPendingDelete sets PendingDelete field to given value.


### GetVrrp

`func (o *NetworkEquipmentInterfaceIp) GetVrrp() map[string]VrrpGroup`

GetVrrp returns the Vrrp field if non-nil, zero value otherwise.

### GetVrrpOk

`func (o *NetworkEquipmentInterfaceIp) GetVrrpOk() (*map[string]VrrpGroup, bool)`

GetVrrpOk returns a tuple with the Vrrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrrp

`func (o *NetworkEquipmentInterfaceIp) SetVrrp(v map[string]VrrpGroup)`

SetVrrp sets Vrrp field to given value.

### HasVrrp

`func (o *NetworkEquipmentInterfaceIp) HasVrrp() bool`

HasVrrp returns a boolean if a field has been set.

### SetVrrpNil

`func (o *NetworkEquipmentInterfaceIp) SetVrrpNil(b bool)`

 SetVrrpNil sets the value for Vrrp to be an explicit nil

### UnsetVrrp
`func (o *NetworkEquipmentInterfaceIp) UnsetVrrp()`

UnsetVrrp ensures that no value is present for Vrrp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AddNetworkEquipmentInterfaceIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Address (e.g. \&quot;10.0.0.1\&quot; or \&quot;2001:db8::1\&quot;). No CIDR. | 
**PrefixLength** | **int32** | Prefix length. 0..32 for ipv4; 0..128 for ipv6. | 

## Methods

### NewAddNetworkEquipmentInterfaceIp

`func NewAddNetworkEquipmentInterfaceIp(address string, prefixLength int32, ) *AddNetworkEquipmentInterfaceIp`

NewAddNetworkEquipmentInterfaceIp instantiates a new AddNetworkEquipmentInterfaceIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNetworkEquipmentInterfaceIpWithDefaults

`func NewAddNetworkEquipmentInterfaceIpWithDefaults() *AddNetworkEquipmentInterfaceIp`

NewAddNetworkEquipmentInterfaceIpWithDefaults instantiates a new AddNetworkEquipmentInterfaceIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddNetworkEquipmentInterfaceIp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddNetworkEquipmentInterfaceIp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddNetworkEquipmentInterfaceIp) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPrefixLength

`func (o *AddNetworkEquipmentInterfaceIp) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *AddNetworkEquipmentInterfaceIp) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *AddNetworkEquipmentInterfaceIp) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



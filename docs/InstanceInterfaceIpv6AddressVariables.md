# InstanceInterfaceIpv6AddressVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**Cidr** | **string** |  | 
**Gateway** | Pointer to **string** |  | [optional] 
**IsDefaultGateway** | **bool** |  | 
**Netmask** | **string** |  | 
**MaskBits** | **float32** |  | 

## Methods

### NewInstanceInterfaceIpv6AddressVariables

`func NewInstanceInterfaceIpv6AddressVariables(ip string, cidr string, isDefaultGateway bool, netmask string, maskBits float32, ) *InstanceInterfaceIpv6AddressVariables`

NewInstanceInterfaceIpv6AddressVariables instantiates a new InstanceInterfaceIpv6AddressVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInterfaceIpv6AddressVariablesWithDefaults

`func NewInstanceInterfaceIpv6AddressVariablesWithDefaults() *InstanceInterfaceIpv6AddressVariables`

NewInstanceInterfaceIpv6AddressVariablesWithDefaults instantiates a new InstanceInterfaceIpv6AddressVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InstanceInterfaceIpv6AddressVariables) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceInterfaceIpv6AddressVariables) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceInterfaceIpv6AddressVariables) SetIp(v string)`

SetIp sets Ip field to given value.


### GetCidr

`func (o *InstanceInterfaceIpv6AddressVariables) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InstanceInterfaceIpv6AddressVariables) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InstanceInterfaceIpv6AddressVariables) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetGateway

`func (o *InstanceInterfaceIpv6AddressVariables) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InstanceInterfaceIpv6AddressVariables) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InstanceInterfaceIpv6AddressVariables) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InstanceInterfaceIpv6AddressVariables) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIsDefaultGateway

`func (o *InstanceInterfaceIpv6AddressVariables) GetIsDefaultGateway() bool`

GetIsDefaultGateway returns the IsDefaultGateway field if non-nil, zero value otherwise.

### GetIsDefaultGatewayOk

`func (o *InstanceInterfaceIpv6AddressVariables) GetIsDefaultGatewayOk() (*bool, bool)`

GetIsDefaultGatewayOk returns a tuple with the IsDefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultGateway

`func (o *InstanceInterfaceIpv6AddressVariables) SetIsDefaultGateway(v bool)`

SetIsDefaultGateway sets IsDefaultGateway field to given value.


### GetNetmask

`func (o *InstanceInterfaceIpv6AddressVariables) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *InstanceInterfaceIpv6AddressVariables) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *InstanceInterfaceIpv6AddressVariables) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetMaskBits

`func (o *InstanceInterfaceIpv6AddressVariables) GetMaskBits() float32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *InstanceInterfaceIpv6AddressVariables) GetMaskBitsOk() (*float32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *InstanceInterfaceIpv6AddressVariables) SetMaskBits(v float32)`

SetMaskBits sets MaskBits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



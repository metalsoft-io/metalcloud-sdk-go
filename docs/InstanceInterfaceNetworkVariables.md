# InstanceInterfaceNetworkVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceType** | [**InstanceInterfaceType**](InstanceInterfaceType.md) |  | 
**LogicalNetworkId** | Pointer to **float32** |  | [optional] 
**TypeInterfaceId** | **float32** |  | 
**MacAddress** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **float32** |  | [optional] 
**VlanId** | Pointer to **float32** |  | [optional] 
**Ipv4Addresses** | Pointer to [**[]InstanceInterfaceIpv4AddressVariables**](InstanceInterfaceIpv4AddressVariables.md) |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]InstanceInterfaceIpv6AddressVariables**](InstanceInterfaceIpv6AddressVariables.md) |  | [optional] 
**Routes** | Pointer to [**[]InstanceInterfaceRouteVariables**](InstanceInterfaceRouteVariables.md) |  | [optional] 
**Members** | Pointer to [**[]InstanceInterfaceMemberVariables**](InstanceInterfaceMemberVariables.md) |  | [optional] 
**Links** | Pointer to [**[]InstanceInterfaceLinkVariables**](InstanceInterfaceLinkVariables.md) |  | [optional] 

## Methods

### NewInstanceInterfaceNetworkVariables

`func NewInstanceInterfaceNetworkVariables(interfaceType InstanceInterfaceType, typeInterfaceId float32, ) *InstanceInterfaceNetworkVariables`

NewInstanceInterfaceNetworkVariables instantiates a new InstanceInterfaceNetworkVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInterfaceNetworkVariablesWithDefaults

`func NewInstanceInterfaceNetworkVariablesWithDefaults() *InstanceInterfaceNetworkVariables`

NewInstanceInterfaceNetworkVariablesWithDefaults instantiates a new InstanceInterfaceNetworkVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *InstanceInterfaceNetworkVariables) GetInterfaceType() InstanceInterfaceType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *InstanceInterfaceNetworkVariables) GetInterfaceTypeOk() (*InstanceInterfaceType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *InstanceInterfaceNetworkVariables) SetInterfaceType(v InstanceInterfaceType)`

SetInterfaceType sets InterfaceType field to given value.


### GetLogicalNetworkId

`func (o *InstanceInterfaceNetworkVariables) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *InstanceInterfaceNetworkVariables) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *InstanceInterfaceNetworkVariables) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *InstanceInterfaceNetworkVariables) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetTypeInterfaceId

`func (o *InstanceInterfaceNetworkVariables) GetTypeInterfaceId() float32`

GetTypeInterfaceId returns the TypeInterfaceId field if non-nil, zero value otherwise.

### GetTypeInterfaceIdOk

`func (o *InstanceInterfaceNetworkVariables) GetTypeInterfaceIdOk() (*float32, bool)`

GetTypeInterfaceIdOk returns a tuple with the TypeInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInterfaceId

`func (o *InstanceInterfaceNetworkVariables) SetTypeInterfaceId(v float32)`

SetTypeInterfaceId sets TypeInterfaceId field to given value.


### GetMacAddress

`func (o *InstanceInterfaceNetworkVariables) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InstanceInterfaceNetworkVariables) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InstanceInterfaceNetworkVariables) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InstanceInterfaceNetworkVariables) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMtu

`func (o *InstanceInterfaceNetworkVariables) GetMtu() float32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InstanceInterfaceNetworkVariables) GetMtuOk() (*float32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InstanceInterfaceNetworkVariables) SetMtu(v float32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InstanceInterfaceNetworkVariables) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetVlanId

`func (o *InstanceInterfaceNetworkVariables) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InstanceInterfaceNetworkVariables) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InstanceInterfaceNetworkVariables) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InstanceInterfaceNetworkVariables) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *InstanceInterfaceNetworkVariables) GetIpv4Addresses() []InstanceInterfaceIpv4AddressVariables`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *InstanceInterfaceNetworkVariables) GetIpv4AddressesOk() (*[]InstanceInterfaceIpv4AddressVariables, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *InstanceInterfaceNetworkVariables) SetIpv4Addresses(v []InstanceInterfaceIpv4AddressVariables)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *InstanceInterfaceNetworkVariables) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *InstanceInterfaceNetworkVariables) GetIpv6Addresses() []InstanceInterfaceIpv6AddressVariables`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *InstanceInterfaceNetworkVariables) GetIpv6AddressesOk() (*[]InstanceInterfaceIpv6AddressVariables, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *InstanceInterfaceNetworkVariables) SetIpv6Addresses(v []InstanceInterfaceIpv6AddressVariables)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *InstanceInterfaceNetworkVariables) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetRoutes

`func (o *InstanceInterfaceNetworkVariables) GetRoutes() []InstanceInterfaceRouteVariables`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *InstanceInterfaceNetworkVariables) GetRoutesOk() (*[]InstanceInterfaceRouteVariables, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *InstanceInterfaceNetworkVariables) SetRoutes(v []InstanceInterfaceRouteVariables)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *InstanceInterfaceNetworkVariables) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetMembers

`func (o *InstanceInterfaceNetworkVariables) GetMembers() []InstanceInterfaceMemberVariables`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InstanceInterfaceNetworkVariables) GetMembersOk() (*[]InstanceInterfaceMemberVariables, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InstanceInterfaceNetworkVariables) SetMembers(v []InstanceInterfaceMemberVariables)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InstanceInterfaceNetworkVariables) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetLinks

`func (o *InstanceInterfaceNetworkVariables) GetLinks() []InstanceInterfaceLinkVariables`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InstanceInterfaceNetworkVariables) GetLinksOk() (*[]InstanceInterfaceLinkVariables, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InstanceInterfaceNetworkVariables) SetLinks(v []InstanceInterfaceLinkVariables)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InstanceInterfaceNetworkVariables) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



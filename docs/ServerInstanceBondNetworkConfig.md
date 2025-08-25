# ServerInstanceBondNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceType** | **string** |  | 
**TypeInterfaceId** | **float32** |  | 
**LogicalNetworkId** | Pointer to **float32** |  | [optional] 
**LogicalNetworkName** | Pointer to **string** |  | [optional] 
**LogicalNetworkLabel** | Pointer to **string** |  | [optional] 
**Ipv4Addresses** | Pointer to [**[]ServerInstanceNetworkConfigIpInfo**](ServerInstanceNetworkConfigIpInfo.md) | The list of IPv4 addresses. | [optional] 
**Ipv6Addresses** | Pointer to [**[]ServerInstanceNetworkConfigIpInfo**](ServerInstanceNetworkConfigIpInfo.md) | The list of IPv6 addresses. | [optional] 
**Routes** | Pointer to [**[]ServerInstanceNetworkConfigRoute**](ServerInstanceNetworkConfigRoute.md) | The list of static routes. | [optional] 
**Members** | Pointer to [**[]ServerInstanceNetworkConfigMembers**](ServerInstanceNetworkConfigMembers.md) | The list of network configuration members. | [optional] 

## Methods

### NewServerInstanceBondNetworkConfig

`func NewServerInstanceBondNetworkConfig(interfaceType string, typeInterfaceId float32, ) *ServerInstanceBondNetworkConfig`

NewServerInstanceBondNetworkConfig instantiates a new ServerInstanceBondNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceBondNetworkConfigWithDefaults

`func NewServerInstanceBondNetworkConfigWithDefaults() *ServerInstanceBondNetworkConfig`

NewServerInstanceBondNetworkConfigWithDefaults instantiates a new ServerInstanceBondNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *ServerInstanceBondNetworkConfig) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *ServerInstanceBondNetworkConfig) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *ServerInstanceBondNetworkConfig) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetTypeInterfaceId

`func (o *ServerInstanceBondNetworkConfig) GetTypeInterfaceId() float32`

GetTypeInterfaceId returns the TypeInterfaceId field if non-nil, zero value otherwise.

### GetTypeInterfaceIdOk

`func (o *ServerInstanceBondNetworkConfig) GetTypeInterfaceIdOk() (*float32, bool)`

GetTypeInterfaceIdOk returns a tuple with the TypeInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInterfaceId

`func (o *ServerInstanceBondNetworkConfig) SetTypeInterfaceId(v float32)`

SetTypeInterfaceId sets TypeInterfaceId field to given value.


### GetLogicalNetworkId

`func (o *ServerInstanceBondNetworkConfig) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *ServerInstanceBondNetworkConfig) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *ServerInstanceBondNetworkConfig) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *ServerInstanceBondNetworkConfig) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetLogicalNetworkName

`func (o *ServerInstanceBondNetworkConfig) GetLogicalNetworkName() string`

GetLogicalNetworkName returns the LogicalNetworkName field if non-nil, zero value otherwise.

### GetLogicalNetworkNameOk

`func (o *ServerInstanceBondNetworkConfig) GetLogicalNetworkNameOk() (*string, bool)`

GetLogicalNetworkNameOk returns a tuple with the LogicalNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkName

`func (o *ServerInstanceBondNetworkConfig) SetLogicalNetworkName(v string)`

SetLogicalNetworkName sets LogicalNetworkName field to given value.

### HasLogicalNetworkName

`func (o *ServerInstanceBondNetworkConfig) HasLogicalNetworkName() bool`

HasLogicalNetworkName returns a boolean if a field has been set.

### GetLogicalNetworkLabel

`func (o *ServerInstanceBondNetworkConfig) GetLogicalNetworkLabel() string`

GetLogicalNetworkLabel returns the LogicalNetworkLabel field if non-nil, zero value otherwise.

### GetLogicalNetworkLabelOk

`func (o *ServerInstanceBondNetworkConfig) GetLogicalNetworkLabelOk() (*string, bool)`

GetLogicalNetworkLabelOk returns a tuple with the LogicalNetworkLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkLabel

`func (o *ServerInstanceBondNetworkConfig) SetLogicalNetworkLabel(v string)`

SetLogicalNetworkLabel sets LogicalNetworkLabel field to given value.

### HasLogicalNetworkLabel

`func (o *ServerInstanceBondNetworkConfig) HasLogicalNetworkLabel() bool`

HasLogicalNetworkLabel returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *ServerInstanceBondNetworkConfig) GetIpv4Addresses() []ServerInstanceNetworkConfigIpInfo`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *ServerInstanceBondNetworkConfig) GetIpv4AddressesOk() (*[]ServerInstanceNetworkConfigIpInfo, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *ServerInstanceBondNetworkConfig) SetIpv4Addresses(v []ServerInstanceNetworkConfigIpInfo)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *ServerInstanceBondNetworkConfig) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ServerInstanceBondNetworkConfig) GetIpv6Addresses() []ServerInstanceNetworkConfigIpInfo`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ServerInstanceBondNetworkConfig) GetIpv6AddressesOk() (*[]ServerInstanceNetworkConfigIpInfo, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ServerInstanceBondNetworkConfig) SetIpv6Addresses(v []ServerInstanceNetworkConfigIpInfo)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ServerInstanceBondNetworkConfig) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetRoutes

`func (o *ServerInstanceBondNetworkConfig) GetRoutes() []ServerInstanceNetworkConfigRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServerInstanceBondNetworkConfig) GetRoutesOk() (*[]ServerInstanceNetworkConfigRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServerInstanceBondNetworkConfig) SetRoutes(v []ServerInstanceNetworkConfigRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ServerInstanceBondNetworkConfig) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetMembers

`func (o *ServerInstanceBondNetworkConfig) GetMembers() []ServerInstanceNetworkConfigMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ServerInstanceBondNetworkConfig) GetMembersOk() (*[]ServerInstanceNetworkConfigMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ServerInstanceBondNetworkConfig) SetMembers(v []ServerInstanceNetworkConfigMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ServerInstanceBondNetworkConfig) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



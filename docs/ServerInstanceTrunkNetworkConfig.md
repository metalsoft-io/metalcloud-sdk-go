# ServerInstanceTrunkNetworkConfig

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
**VlanId** | **float32** |  | 
**Links** | Pointer to [**[]ServerInstanceNetworkConfigMembers**](ServerInstanceNetworkConfigMembers.md) | The list of network configuration members. | [optional] 

## Methods

### NewServerInstanceTrunkNetworkConfig

`func NewServerInstanceTrunkNetworkConfig(interfaceType string, typeInterfaceId float32, vlanId float32, ) *ServerInstanceTrunkNetworkConfig`

NewServerInstanceTrunkNetworkConfig instantiates a new ServerInstanceTrunkNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceTrunkNetworkConfigWithDefaults

`func NewServerInstanceTrunkNetworkConfigWithDefaults() *ServerInstanceTrunkNetworkConfig`

NewServerInstanceTrunkNetworkConfigWithDefaults instantiates a new ServerInstanceTrunkNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *ServerInstanceTrunkNetworkConfig) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *ServerInstanceTrunkNetworkConfig) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *ServerInstanceTrunkNetworkConfig) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetTypeInterfaceId

`func (o *ServerInstanceTrunkNetworkConfig) GetTypeInterfaceId() float32`

GetTypeInterfaceId returns the TypeInterfaceId field if non-nil, zero value otherwise.

### GetTypeInterfaceIdOk

`func (o *ServerInstanceTrunkNetworkConfig) GetTypeInterfaceIdOk() (*float32, bool)`

GetTypeInterfaceIdOk returns a tuple with the TypeInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInterfaceId

`func (o *ServerInstanceTrunkNetworkConfig) SetTypeInterfaceId(v float32)`

SetTypeInterfaceId sets TypeInterfaceId field to given value.


### GetLogicalNetworkId

`func (o *ServerInstanceTrunkNetworkConfig) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *ServerInstanceTrunkNetworkConfig) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *ServerInstanceTrunkNetworkConfig) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *ServerInstanceTrunkNetworkConfig) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetLogicalNetworkName

`func (o *ServerInstanceTrunkNetworkConfig) GetLogicalNetworkName() string`

GetLogicalNetworkName returns the LogicalNetworkName field if non-nil, zero value otherwise.

### GetLogicalNetworkNameOk

`func (o *ServerInstanceTrunkNetworkConfig) GetLogicalNetworkNameOk() (*string, bool)`

GetLogicalNetworkNameOk returns a tuple with the LogicalNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkName

`func (o *ServerInstanceTrunkNetworkConfig) SetLogicalNetworkName(v string)`

SetLogicalNetworkName sets LogicalNetworkName field to given value.

### HasLogicalNetworkName

`func (o *ServerInstanceTrunkNetworkConfig) HasLogicalNetworkName() bool`

HasLogicalNetworkName returns a boolean if a field has been set.

### GetLogicalNetworkLabel

`func (o *ServerInstanceTrunkNetworkConfig) GetLogicalNetworkLabel() string`

GetLogicalNetworkLabel returns the LogicalNetworkLabel field if non-nil, zero value otherwise.

### GetLogicalNetworkLabelOk

`func (o *ServerInstanceTrunkNetworkConfig) GetLogicalNetworkLabelOk() (*string, bool)`

GetLogicalNetworkLabelOk returns a tuple with the LogicalNetworkLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkLabel

`func (o *ServerInstanceTrunkNetworkConfig) SetLogicalNetworkLabel(v string)`

SetLogicalNetworkLabel sets LogicalNetworkLabel field to given value.

### HasLogicalNetworkLabel

`func (o *ServerInstanceTrunkNetworkConfig) HasLogicalNetworkLabel() bool`

HasLogicalNetworkLabel returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *ServerInstanceTrunkNetworkConfig) GetIpv4Addresses() []ServerInstanceNetworkConfigIpInfo`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *ServerInstanceTrunkNetworkConfig) GetIpv4AddressesOk() (*[]ServerInstanceNetworkConfigIpInfo, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *ServerInstanceTrunkNetworkConfig) SetIpv4Addresses(v []ServerInstanceNetworkConfigIpInfo)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *ServerInstanceTrunkNetworkConfig) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ServerInstanceTrunkNetworkConfig) GetIpv6Addresses() []ServerInstanceNetworkConfigIpInfo`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ServerInstanceTrunkNetworkConfig) GetIpv6AddressesOk() (*[]ServerInstanceNetworkConfigIpInfo, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ServerInstanceTrunkNetworkConfig) SetIpv6Addresses(v []ServerInstanceNetworkConfigIpInfo)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ServerInstanceTrunkNetworkConfig) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetRoutes

`func (o *ServerInstanceTrunkNetworkConfig) GetRoutes() []ServerInstanceNetworkConfigRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServerInstanceTrunkNetworkConfig) GetRoutesOk() (*[]ServerInstanceNetworkConfigRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServerInstanceTrunkNetworkConfig) SetRoutes(v []ServerInstanceNetworkConfigRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ServerInstanceTrunkNetworkConfig) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetVlanId

`func (o *ServerInstanceTrunkNetworkConfig) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ServerInstanceTrunkNetworkConfig) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ServerInstanceTrunkNetworkConfig) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.


### GetLinks

`func (o *ServerInstanceTrunkNetworkConfig) GetLinks() []ServerInstanceNetworkConfigMembers`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceTrunkNetworkConfig) GetLinksOk() (*[]ServerInstanceNetworkConfigMembers, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceTrunkNetworkConfig) SetLinks(v []ServerInstanceNetworkConfigMembers)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstanceTrunkNetworkConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ServerInstancePhysicalNetworkConfig

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
**MacAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewServerInstancePhysicalNetworkConfig

`func NewServerInstancePhysicalNetworkConfig(interfaceType string, typeInterfaceId float32, ) *ServerInstancePhysicalNetworkConfig`

NewServerInstancePhysicalNetworkConfig instantiates a new ServerInstancePhysicalNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstancePhysicalNetworkConfigWithDefaults

`func NewServerInstancePhysicalNetworkConfigWithDefaults() *ServerInstancePhysicalNetworkConfig`

NewServerInstancePhysicalNetworkConfigWithDefaults instantiates a new ServerInstancePhysicalNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *ServerInstancePhysicalNetworkConfig) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *ServerInstancePhysicalNetworkConfig) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *ServerInstancePhysicalNetworkConfig) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetTypeInterfaceId

`func (o *ServerInstancePhysicalNetworkConfig) GetTypeInterfaceId() float32`

GetTypeInterfaceId returns the TypeInterfaceId field if non-nil, zero value otherwise.

### GetTypeInterfaceIdOk

`func (o *ServerInstancePhysicalNetworkConfig) GetTypeInterfaceIdOk() (*float32, bool)`

GetTypeInterfaceIdOk returns a tuple with the TypeInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeInterfaceId

`func (o *ServerInstancePhysicalNetworkConfig) SetTypeInterfaceId(v float32)`

SetTypeInterfaceId sets TypeInterfaceId field to given value.


### GetLogicalNetworkId

`func (o *ServerInstancePhysicalNetworkConfig) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *ServerInstancePhysicalNetworkConfig) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *ServerInstancePhysicalNetworkConfig) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *ServerInstancePhysicalNetworkConfig) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetLogicalNetworkName

`func (o *ServerInstancePhysicalNetworkConfig) GetLogicalNetworkName() string`

GetLogicalNetworkName returns the LogicalNetworkName field if non-nil, zero value otherwise.

### GetLogicalNetworkNameOk

`func (o *ServerInstancePhysicalNetworkConfig) GetLogicalNetworkNameOk() (*string, bool)`

GetLogicalNetworkNameOk returns a tuple with the LogicalNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkName

`func (o *ServerInstancePhysicalNetworkConfig) SetLogicalNetworkName(v string)`

SetLogicalNetworkName sets LogicalNetworkName field to given value.

### HasLogicalNetworkName

`func (o *ServerInstancePhysicalNetworkConfig) HasLogicalNetworkName() bool`

HasLogicalNetworkName returns a boolean if a field has been set.

### GetLogicalNetworkLabel

`func (o *ServerInstancePhysicalNetworkConfig) GetLogicalNetworkLabel() string`

GetLogicalNetworkLabel returns the LogicalNetworkLabel field if non-nil, zero value otherwise.

### GetLogicalNetworkLabelOk

`func (o *ServerInstancePhysicalNetworkConfig) GetLogicalNetworkLabelOk() (*string, bool)`

GetLogicalNetworkLabelOk returns a tuple with the LogicalNetworkLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkLabel

`func (o *ServerInstancePhysicalNetworkConfig) SetLogicalNetworkLabel(v string)`

SetLogicalNetworkLabel sets LogicalNetworkLabel field to given value.

### HasLogicalNetworkLabel

`func (o *ServerInstancePhysicalNetworkConfig) HasLogicalNetworkLabel() bool`

HasLogicalNetworkLabel returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *ServerInstancePhysicalNetworkConfig) GetIpv4Addresses() []ServerInstanceNetworkConfigIpInfo`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *ServerInstancePhysicalNetworkConfig) GetIpv4AddressesOk() (*[]ServerInstanceNetworkConfigIpInfo, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *ServerInstancePhysicalNetworkConfig) SetIpv4Addresses(v []ServerInstanceNetworkConfigIpInfo)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *ServerInstancePhysicalNetworkConfig) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *ServerInstancePhysicalNetworkConfig) GetIpv6Addresses() []ServerInstanceNetworkConfigIpInfo`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *ServerInstancePhysicalNetworkConfig) GetIpv6AddressesOk() (*[]ServerInstanceNetworkConfigIpInfo, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *ServerInstancePhysicalNetworkConfig) SetIpv6Addresses(v []ServerInstanceNetworkConfigIpInfo)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *ServerInstancePhysicalNetworkConfig) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetRoutes

`func (o *ServerInstancePhysicalNetworkConfig) GetRoutes() []ServerInstanceNetworkConfigRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServerInstancePhysicalNetworkConfig) GetRoutesOk() (*[]ServerInstanceNetworkConfigRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServerInstancePhysicalNetworkConfig) SetRoutes(v []ServerInstanceNetworkConfigRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ServerInstancePhysicalNetworkConfig) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetMacAddress

`func (o *ServerInstancePhysicalNetworkConfig) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ServerInstancePhysicalNetworkConfig) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ServerInstancePhysicalNetworkConfig) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ServerInstancePhysicalNetworkConfig) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



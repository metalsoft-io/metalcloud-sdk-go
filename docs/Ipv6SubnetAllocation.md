# Ipv6SubnetAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**Status** | [**ResourceAllocationStatus**](ResourceAllocationStatus.md) |  | 
**NetworkAddress** | **string** |  | 
**PrefixLength** | **int32** |  | 
**Gateway** | **string** |  | 
**GatewayPlacement** | [**NullableSubnetGatewayPlacement**](SubnetGatewayPlacement.md) |  | 

## Methods

### NewIpv6SubnetAllocation

`func NewIpv6SubnetAllocation(id int32, scope ResourceScope, status ResourceAllocationStatus, networkAddress string, prefixLength int32, gateway string, gatewayPlacement NullableSubnetGatewayPlacement, ) *Ipv6SubnetAllocation`

NewIpv6SubnetAllocation instantiates a new Ipv6SubnetAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6SubnetAllocationWithDefaults

`func NewIpv6SubnetAllocationWithDefaults() *Ipv6SubnetAllocation`

NewIpv6SubnetAllocationWithDefaults instantiates a new Ipv6SubnetAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ipv6SubnetAllocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ipv6SubnetAllocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ipv6SubnetAllocation) SetId(v int32)`

SetId sets Id field to given value.


### GetScope

`func (o *Ipv6SubnetAllocation) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Ipv6SubnetAllocation) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Ipv6SubnetAllocation) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *Ipv6SubnetAllocation) GetStatus() ResourceAllocationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ipv6SubnetAllocation) GetStatusOk() (*ResourceAllocationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ipv6SubnetAllocation) SetStatus(v ResourceAllocationStatus)`

SetStatus sets Status field to given value.


### GetNetworkAddress

`func (o *Ipv6SubnetAllocation) GetNetworkAddress() string`

GetNetworkAddress returns the NetworkAddress field if non-nil, zero value otherwise.

### GetNetworkAddressOk

`func (o *Ipv6SubnetAllocation) GetNetworkAddressOk() (*string, bool)`

GetNetworkAddressOk returns a tuple with the NetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddress

`func (o *Ipv6SubnetAllocation) SetNetworkAddress(v string)`

SetNetworkAddress sets NetworkAddress field to given value.


### GetPrefixLength

`func (o *Ipv6SubnetAllocation) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Ipv6SubnetAllocation) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Ipv6SubnetAllocation) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetGateway

`func (o *Ipv6SubnetAllocation) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Ipv6SubnetAllocation) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Ipv6SubnetAllocation) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayPlacement

`func (o *Ipv6SubnetAllocation) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *Ipv6SubnetAllocation) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *Ipv6SubnetAllocation) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.


### SetGatewayPlacementNil

`func (o *Ipv6SubnetAllocation) SetGatewayPlacementNil(b bool)`

 SetGatewayPlacementNil sets the value for GatewayPlacement to be an explicit nil

### UnsetGatewayPlacement
`func (o *Ipv6SubnetAllocation) UnsetGatewayPlacement()`

UnsetGatewayPlacement ensures that no value is present for GatewayPlacement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



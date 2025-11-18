# GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**GatewayPlacement** | [**SubnetGatewayPlacement**](SubnetGatewayPlacement.md) |  | [default to SUBNETGATEWAYPLACEMENT_DEFAULT]
**SubnetId** | **int32** |  | 
**SubnetPoolIds** | **[]int32** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewGetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response

`func NewGetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, gatewayPlacement SubnetGatewayPlacement, subnetId int32, subnetPoolIds []int32, prefixLength int32, ) *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response`

NewGetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response instantiates a new GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogicalNetworkProfileIpv4SubnetAllocationStrategy200ResponseWithDefaults

`func NewGetLogicalNetworkProfileIpv4SubnetAllocationStrategy200ResponseWithDefaults() *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response`

NewGetLogicalNetworkProfileIpv4SubnetAllocationStrategy200ResponseWithDefaults instantiates a new GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.


### GetSubnetId

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetPoolIds

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetSubnetPoolIds() []int32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetSubnetPoolIdsOk() (*[]int32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetSubnetPoolIds(v []int32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *GetLogicalNetworkProfileIpv4SubnetAllocationStrategy200Response) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



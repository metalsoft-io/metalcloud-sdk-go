# Ipv4SubnetAllocationStrategy1DataItem

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

### NewIpv4SubnetAllocationStrategy1DataItem

`func NewIpv4SubnetAllocationStrategy1DataItem(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, gatewayPlacement SubnetGatewayPlacement, subnetId int32, subnetPoolIds []int32, prefixLength int32, ) *Ipv4SubnetAllocationStrategy1DataItem`

NewIpv4SubnetAllocationStrategy1DataItem instantiates a new Ipv4SubnetAllocationStrategy1DataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4SubnetAllocationStrategy1DataItemWithDefaults

`func NewIpv4SubnetAllocationStrategy1DataItemWithDefaults() *Ipv4SubnetAllocationStrategy1DataItem`

NewIpv4SubnetAllocationStrategy1DataItemWithDefaults instantiates a new Ipv4SubnetAllocationStrategy1DataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.


### GetSubnetId

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetPoolIds

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetSubnetPoolIds() []int32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetSubnetPoolIdsOk() (*[]int32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetSubnetPoolIds(v []int32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *Ipv4SubnetAllocationStrategy1DataItem) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *Ipv4SubnetAllocationStrategy1DataItem) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



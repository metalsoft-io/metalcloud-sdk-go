# CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**SubnetId** | **int32** |  | 
**SubnetPoolIds** | **[]int32** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewCreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response

`func NewCreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, subnetId int32, subnetPoolIds []int32, prefixLength int32, ) *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response`

NewCreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response instantiates a new CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201ResponseWithDefaults

`func NewCreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201ResponseWithDefaults() *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response`

NewCreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201ResponseWithDefaults instantiates a new CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetPoolIds

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetSubnetPoolIds() []int32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetSubnetPoolIdsOk() (*[]int32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetSubnetPoolIds(v []int32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateLogicalNetworkProfileIpv6SubnetAllocationStrategy201Response) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



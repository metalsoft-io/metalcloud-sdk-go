# AutoIpv6PointToPointAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**PointToPointAllocationStrategyKind**](PointToPointAllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**SubnetPoolIds** | **[]int64** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewAutoIpv6PointToPointAllocationStrategy

`func NewAutoIpv6PointToPointAllocationStrategy(id int64, createdAt time.Time, updatedAt time.Time, kind PointToPointAllocationStrategyKind, scope ResourceScope, subnetPoolIds []int64, prefixLength int32, ) *AutoIpv6PointToPointAllocationStrategy`

NewAutoIpv6PointToPointAllocationStrategy instantiates a new AutoIpv6PointToPointAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoIpv6PointToPointAllocationStrategyWithDefaults

`func NewAutoIpv6PointToPointAllocationStrategyWithDefaults() *AutoIpv6PointToPointAllocationStrategy`

NewAutoIpv6PointToPointAllocationStrategyWithDefaults instantiates a new AutoIpv6PointToPointAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoIpv6PointToPointAllocationStrategy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoIpv6PointToPointAllocationStrategy) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *AutoIpv6PointToPointAllocationStrategy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoIpv6PointToPointAllocationStrategy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AutoIpv6PointToPointAllocationStrategy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutoIpv6PointToPointAllocationStrategy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *AutoIpv6PointToPointAllocationStrategy) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AutoIpv6PointToPointAllocationStrategy) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *AutoIpv6PointToPointAllocationStrategy) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AutoIpv6PointToPointAllocationStrategy) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetPoolIds

`func (o *AutoIpv6PointToPointAllocationStrategy) GetSubnetPoolIds() []int64`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetSubnetPoolIdsOk() (*[]int64, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *AutoIpv6PointToPointAllocationStrategy) SetSubnetPoolIds(v []int64)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *AutoIpv6PointToPointAllocationStrategy) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *AutoIpv6PointToPointAllocationStrategy) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *AutoIpv6PointToPointAllocationStrategy) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ManualIpv6SubnetAllocationStrategy

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

## Methods

### NewManualIpv6SubnetAllocationStrategy

`func NewManualIpv6SubnetAllocationStrategy(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, gatewayPlacement SubnetGatewayPlacement, subnetId int32, ) *ManualIpv6SubnetAllocationStrategy`

NewManualIpv6SubnetAllocationStrategy instantiates a new ManualIpv6SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualIpv6SubnetAllocationStrategyWithDefaults

`func NewManualIpv6SubnetAllocationStrategyWithDefaults() *ManualIpv6SubnetAllocationStrategy`

NewManualIpv6SubnetAllocationStrategyWithDefaults instantiates a new ManualIpv6SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualIpv6SubnetAllocationStrategy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualIpv6SubnetAllocationStrategy) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ManualIpv6SubnetAllocationStrategy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManualIpv6SubnetAllocationStrategy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManualIpv6SubnetAllocationStrategy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManualIpv6SubnetAllocationStrategy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *ManualIpv6SubnetAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManualIpv6SubnetAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *ManualIpv6SubnetAllocationStrategy) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ManualIpv6SubnetAllocationStrategy) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *ManualIpv6SubnetAllocationStrategy) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *ManualIpv6SubnetAllocationStrategy) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.


### GetSubnetId

`func (o *ManualIpv6SubnetAllocationStrategy) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ManualIpv6SubnetAllocationStrategy) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ManualIpv6SubnetAllocationStrategy) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



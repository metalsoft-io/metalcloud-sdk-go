# ManualIpv4SubnetAllocationStrategy

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

### NewManualIpv4SubnetAllocationStrategy

`func NewManualIpv4SubnetAllocationStrategy(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, gatewayPlacement SubnetGatewayPlacement, subnetId int32, ) *ManualIpv4SubnetAllocationStrategy`

NewManualIpv4SubnetAllocationStrategy instantiates a new ManualIpv4SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualIpv4SubnetAllocationStrategyWithDefaults

`func NewManualIpv4SubnetAllocationStrategyWithDefaults() *ManualIpv4SubnetAllocationStrategy`

NewManualIpv4SubnetAllocationStrategyWithDefaults instantiates a new ManualIpv4SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualIpv4SubnetAllocationStrategy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualIpv4SubnetAllocationStrategy) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ManualIpv4SubnetAllocationStrategy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManualIpv4SubnetAllocationStrategy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManualIpv4SubnetAllocationStrategy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManualIpv4SubnetAllocationStrategy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *ManualIpv4SubnetAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManualIpv4SubnetAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *ManualIpv4SubnetAllocationStrategy) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ManualIpv4SubnetAllocationStrategy) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *ManualIpv4SubnetAllocationStrategy) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *ManualIpv4SubnetAllocationStrategy) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.


### GetSubnetId

`func (o *ManualIpv4SubnetAllocationStrategy) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ManualIpv4SubnetAllocationStrategy) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ManualIpv4SubnetAllocationStrategy) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



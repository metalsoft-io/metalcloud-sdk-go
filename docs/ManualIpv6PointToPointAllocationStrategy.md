# ManualIpv6PointToPointAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**PointToPointAllocationStrategyKind**](PointToPointAllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**SubnetId** | **int64** | ID of the IPAM subnet backing the link. | 
**InterfaceABinding** | [**PointToPointInterfaceBinding**](PointToPointInterfaceBinding.md) |  | 

## Methods

### NewManualIpv6PointToPointAllocationStrategy

`func NewManualIpv6PointToPointAllocationStrategy(id int64, createdAt time.Time, updatedAt time.Time, kind PointToPointAllocationStrategyKind, scope ResourceScope, subnetId int64, interfaceABinding PointToPointInterfaceBinding, ) *ManualIpv6PointToPointAllocationStrategy`

NewManualIpv6PointToPointAllocationStrategy instantiates a new ManualIpv6PointToPointAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualIpv6PointToPointAllocationStrategyWithDefaults

`func NewManualIpv6PointToPointAllocationStrategyWithDefaults() *ManualIpv6PointToPointAllocationStrategy`

NewManualIpv6PointToPointAllocationStrategyWithDefaults instantiates a new ManualIpv6PointToPointAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualIpv6PointToPointAllocationStrategy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualIpv6PointToPointAllocationStrategy) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ManualIpv6PointToPointAllocationStrategy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManualIpv6PointToPointAllocationStrategy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManualIpv6PointToPointAllocationStrategy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManualIpv6PointToPointAllocationStrategy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *ManualIpv6PointToPointAllocationStrategy) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManualIpv6PointToPointAllocationStrategy) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *ManualIpv6PointToPointAllocationStrategy) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ManualIpv6PointToPointAllocationStrategy) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *ManualIpv6PointToPointAllocationStrategy) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ManualIpv6PointToPointAllocationStrategy) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetInterfaceABinding

`func (o *ManualIpv6PointToPointAllocationStrategy) GetInterfaceABinding() PointToPointInterfaceBinding`

GetInterfaceABinding returns the InterfaceABinding field if non-nil, zero value otherwise.

### GetInterfaceABindingOk

`func (o *ManualIpv6PointToPointAllocationStrategy) GetInterfaceABindingOk() (*PointToPointInterfaceBinding, bool)`

GetInterfaceABindingOk returns a tuple with the InterfaceABinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceABinding

`func (o *ManualIpv6PointToPointAllocationStrategy) SetInterfaceABinding(v PointToPointInterfaceBinding)`

SetInterfaceABinding sets InterfaceABinding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



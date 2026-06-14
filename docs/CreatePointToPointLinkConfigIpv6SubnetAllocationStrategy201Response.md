# CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response

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
**SubnetPoolIds** | **[]int64** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewCreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response

`func NewCreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response(id int64, createdAt time.Time, updatedAt time.Time, kind PointToPointAllocationStrategyKind, scope ResourceScope, subnetId int64, interfaceABinding PointToPointInterfaceBinding, subnetPoolIds []int64, prefixLength int32, ) *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response`

NewCreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response instantiates a new CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201ResponseWithDefaults

`func NewCreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201ResponseWithDefaults() *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response`

NewCreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201ResponseWithDefaults instantiates a new CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetInterfaceABinding

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetInterfaceABinding() PointToPointInterfaceBinding`

GetInterfaceABinding returns the InterfaceABinding field if non-nil, zero value otherwise.

### GetInterfaceABindingOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetInterfaceABindingOk() (*PointToPointInterfaceBinding, bool)`

GetInterfaceABindingOk returns a tuple with the InterfaceABinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceABinding

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetInterfaceABinding(v PointToPointInterfaceBinding)`

SetInterfaceABinding sets InterfaceABinding field to given value.


### GetSubnetPoolIds

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetSubnetPoolIds() []int64`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetSubnetPoolIdsOk() (*[]int64, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetSubnetPoolIds(v []int64)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateZoneAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**ZoneName** | **string** |  | 

## Methods

### NewCreateZoneAllocationStrategy

`func NewCreateZoneAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, zoneName string, ) *CreateZoneAllocationStrategy`

NewCreateZoneAllocationStrategy instantiates a new CreateZoneAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateZoneAllocationStrategyWithDefaults

`func NewCreateZoneAllocationStrategyWithDefaults() *CreateZoneAllocationStrategy`

NewCreateZoneAllocationStrategyWithDefaults instantiates a new CreateZoneAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateZoneAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateZoneAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateZoneAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateZoneAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateZoneAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateZoneAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetZoneName

`func (o *CreateZoneAllocationStrategy) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *CreateZoneAllocationStrategy) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *CreateZoneAllocationStrategy) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



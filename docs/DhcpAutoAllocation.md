# DhcpAutoAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**SubnetPoolIds** | **[]float32** | Ordered IPAM OOB subnet ids to allocate the next free IP from | 

## Methods

### NewDhcpAutoAllocation

`func NewDhcpAutoAllocation(kind string, subnetPoolIds []float32, ) *DhcpAutoAllocation`

NewDhcpAutoAllocation instantiates a new DhcpAutoAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpAutoAllocationWithDefaults

`func NewDhcpAutoAllocationWithDefaults() *DhcpAutoAllocation`

NewDhcpAutoAllocationWithDefaults instantiates a new DhcpAutoAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DhcpAutoAllocation) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DhcpAutoAllocation) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DhcpAutoAllocation) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSubnetPoolIds

`func (o *DhcpAutoAllocation) GetSubnetPoolIds() []float32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *DhcpAutoAllocation) GetSubnetPoolIdsOk() (*[]float32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *DhcpAutoAllocation) SetSubnetPoolIds(v []float32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



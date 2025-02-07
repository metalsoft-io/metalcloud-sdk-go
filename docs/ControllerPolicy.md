# ControllerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetAllocation** | **[]string** | List of subnets allocated to this datacenter | 

## Methods

### NewControllerPolicy

`func NewControllerPolicy(subnetAllocation []string, ) *ControllerPolicy`

NewControllerPolicy instantiates a new ControllerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerPolicyWithDefaults

`func NewControllerPolicyWithDefaults() *ControllerPolicy`

NewControllerPolicyWithDefaults instantiates a new ControllerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetAllocation

`func (o *ControllerPolicy) GetSubnetAllocation() []string`

GetSubnetAllocation returns the SubnetAllocation field if non-nil, zero value otherwise.

### GetSubnetAllocationOk

`func (o *ControllerPolicy) GetSubnetAllocationOk() (*[]string, bool)`

GetSubnetAllocationOk returns a tuple with the SubnetAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocation

`func (o *ControllerPolicy) SetSubnetAllocation(v []string)`

SetSubnetAllocation sets SubnetAllocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



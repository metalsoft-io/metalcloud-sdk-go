# InfinibandLogicalNetworkIpv6Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**[]Ipv6SubnetAllocation**](Ipv6SubnetAllocation.md) |  | 
**SubnetAllocationStrategies** | [**[]Ipv6SubnetAllocationStrategy**](Ipv6SubnetAllocationStrategy.md) |  | 

## Methods

### NewInfinibandLogicalNetworkIpv6Properties

`func NewInfinibandLogicalNetworkIpv6Properties(subnets []Ipv6SubnetAllocation, subnetAllocationStrategies []Ipv6SubnetAllocationStrategy, ) *InfinibandLogicalNetworkIpv6Properties`

NewInfinibandLogicalNetworkIpv6Properties instantiates a new InfinibandLogicalNetworkIpv6Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfinibandLogicalNetworkIpv6PropertiesWithDefaults

`func NewInfinibandLogicalNetworkIpv6PropertiesWithDefaults() *InfinibandLogicalNetworkIpv6Properties`

NewInfinibandLogicalNetworkIpv6PropertiesWithDefaults instantiates a new InfinibandLogicalNetworkIpv6Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *InfinibandLogicalNetworkIpv6Properties) GetSubnets() []Ipv6SubnetAllocation`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InfinibandLogicalNetworkIpv6Properties) GetSubnetsOk() (*[]Ipv6SubnetAllocation, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InfinibandLogicalNetworkIpv6Properties) SetSubnets(v []Ipv6SubnetAllocation)`

SetSubnets sets Subnets field to given value.


### GetSubnetAllocationStrategies

`func (o *InfinibandLogicalNetworkIpv6Properties) GetSubnetAllocationStrategies() []Ipv6SubnetAllocationStrategy`

GetSubnetAllocationStrategies returns the SubnetAllocationStrategies field if non-nil, zero value otherwise.

### GetSubnetAllocationStrategiesOk

`func (o *InfinibandLogicalNetworkIpv6Properties) GetSubnetAllocationStrategiesOk() (*[]Ipv6SubnetAllocationStrategy, bool)`

GetSubnetAllocationStrategiesOk returns a tuple with the SubnetAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocationStrategies

`func (o *InfinibandLogicalNetworkIpv6Properties) SetSubnetAllocationStrategies(v []Ipv6SubnetAllocationStrategy)`

SetSubnetAllocationStrategies sets SubnetAllocationStrategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LogicalNetworkIpv4Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**[]Ipv4SubnetAllocation**](Ipv4SubnetAllocation.md) |  | 
**SubnetAllocationStrategies** | [**[]Ipv4SubnetAllocationStrategy1**](Ipv4SubnetAllocationStrategy1.md) |  | 

## Methods

### NewLogicalNetworkIpv4Properties

`func NewLogicalNetworkIpv4Properties(subnets []Ipv4SubnetAllocation, subnetAllocationStrategies []Ipv4SubnetAllocationStrategy1, ) *LogicalNetworkIpv4Properties`

NewLogicalNetworkIpv4Properties instantiates a new LogicalNetworkIpv4Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkIpv4PropertiesWithDefaults

`func NewLogicalNetworkIpv4PropertiesWithDefaults() *LogicalNetworkIpv4Properties`

NewLogicalNetworkIpv4PropertiesWithDefaults instantiates a new LogicalNetworkIpv4Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *LogicalNetworkIpv4Properties) GetSubnets() []Ipv4SubnetAllocation`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *LogicalNetworkIpv4Properties) GetSubnetsOk() (*[]Ipv4SubnetAllocation, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *LogicalNetworkIpv4Properties) SetSubnets(v []Ipv4SubnetAllocation)`

SetSubnets sets Subnets field to given value.


### GetSubnetAllocationStrategies

`func (o *LogicalNetworkIpv4Properties) GetSubnetAllocationStrategies() []Ipv4SubnetAllocationStrategy1`

GetSubnetAllocationStrategies returns the SubnetAllocationStrategies field if non-nil, zero value otherwise.

### GetSubnetAllocationStrategiesOk

`func (o *LogicalNetworkIpv4Properties) GetSubnetAllocationStrategiesOk() (*[]Ipv4SubnetAllocationStrategy1, bool)`

GetSubnetAllocationStrategiesOk returns a tuple with the SubnetAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocationStrategies

`func (o *LogicalNetworkIpv4Properties) SetSubnetAllocationStrategies(v []Ipv4SubnetAllocationStrategy1)`

SetSubnetAllocationStrategies sets SubnetAllocationStrategies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



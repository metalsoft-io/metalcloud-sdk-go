# PointToPointLinkIpv6Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**[]Ipv6PointToPointSubnetAllocation**](Ipv6PointToPointSubnetAllocation.md) | Deployed /127 subnet allocations (cardinality 0..1 per link, kept as an array for parity with logical networks). Empty if not yet deployed or unnumbered. | 
**SubnetAllocationStrategies** | [**[]CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response**](CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response.md) |  | 
**StaticRoutes** | [**[]PointToPointStaticRoute**](PointToPointStaticRoute.md) | Deployed IPv6 static routes. | 

## Methods

### NewPointToPointLinkIpv6Properties

`func NewPointToPointLinkIpv6Properties(subnets []Ipv6PointToPointSubnetAllocation, subnetAllocationStrategies []CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response, staticRoutes []PointToPointStaticRoute, ) *PointToPointLinkIpv6Properties`

NewPointToPointLinkIpv6Properties instantiates a new PointToPointLinkIpv6Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointToPointLinkIpv6PropertiesWithDefaults

`func NewPointToPointLinkIpv6PropertiesWithDefaults() *PointToPointLinkIpv6Properties`

NewPointToPointLinkIpv6PropertiesWithDefaults instantiates a new PointToPointLinkIpv6Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *PointToPointLinkIpv6Properties) GetSubnets() []Ipv6PointToPointSubnetAllocation`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *PointToPointLinkIpv6Properties) GetSubnetsOk() (*[]Ipv6PointToPointSubnetAllocation, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *PointToPointLinkIpv6Properties) SetSubnets(v []Ipv6PointToPointSubnetAllocation)`

SetSubnets sets Subnets field to given value.


### GetSubnetAllocationStrategies

`func (o *PointToPointLinkIpv6Properties) GetSubnetAllocationStrategies() []CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response`

GetSubnetAllocationStrategies returns the SubnetAllocationStrategies field if non-nil, zero value otherwise.

### GetSubnetAllocationStrategiesOk

`func (o *PointToPointLinkIpv6Properties) GetSubnetAllocationStrategiesOk() (*[]CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response, bool)`

GetSubnetAllocationStrategiesOk returns a tuple with the SubnetAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocationStrategies

`func (o *PointToPointLinkIpv6Properties) SetSubnetAllocationStrategies(v []CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response)`

SetSubnetAllocationStrategies sets SubnetAllocationStrategies field to given value.


### GetStaticRoutes

`func (o *PointToPointLinkIpv6Properties) GetStaticRoutes() []PointToPointStaticRoute`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PointToPointLinkIpv6Properties) GetStaticRoutesOk() (*[]PointToPointStaticRoute, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PointToPointLinkIpv6Properties) SetStaticRoutes(v []PointToPointStaticRoute)`

SetStaticRoutes sets StaticRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



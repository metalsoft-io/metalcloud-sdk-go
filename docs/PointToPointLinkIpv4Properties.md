# PointToPointLinkIpv4Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**[]Ipv4PointToPointSubnetAllocation**](Ipv4PointToPointSubnetAllocation.md) | Deployed /31 subnet allocations (cardinality 0..1 per link, kept as an array for parity with logical networks). Empty if not yet deployed or unnumbered. | 
**SubnetAllocationStrategies** | [**[]CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response**](CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response.md) |  | 
**StaticRoutes** | [**[]PointToPointStaticRoute**](PointToPointStaticRoute.md) | Deployed IPv4 static routes. | 

## Methods

### NewPointToPointLinkIpv4Properties

`func NewPointToPointLinkIpv4Properties(subnets []Ipv4PointToPointSubnetAllocation, subnetAllocationStrategies []CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response, staticRoutes []PointToPointStaticRoute, ) *PointToPointLinkIpv4Properties`

NewPointToPointLinkIpv4Properties instantiates a new PointToPointLinkIpv4Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointToPointLinkIpv4PropertiesWithDefaults

`func NewPointToPointLinkIpv4PropertiesWithDefaults() *PointToPointLinkIpv4Properties`

NewPointToPointLinkIpv4PropertiesWithDefaults instantiates a new PointToPointLinkIpv4Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *PointToPointLinkIpv4Properties) GetSubnets() []Ipv4PointToPointSubnetAllocation`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *PointToPointLinkIpv4Properties) GetSubnetsOk() (*[]Ipv4PointToPointSubnetAllocation, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *PointToPointLinkIpv4Properties) SetSubnets(v []Ipv4PointToPointSubnetAllocation)`

SetSubnets sets Subnets field to given value.


### GetSubnetAllocationStrategies

`func (o *PointToPointLinkIpv4Properties) GetSubnetAllocationStrategies() []CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response`

GetSubnetAllocationStrategies returns the SubnetAllocationStrategies field if non-nil, zero value otherwise.

### GetSubnetAllocationStrategiesOk

`func (o *PointToPointLinkIpv4Properties) GetSubnetAllocationStrategiesOk() (*[]CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response, bool)`

GetSubnetAllocationStrategiesOk returns a tuple with the SubnetAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocationStrategies

`func (o *PointToPointLinkIpv4Properties) SetSubnetAllocationStrategies(v []CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response)`

SetSubnetAllocationStrategies sets SubnetAllocationStrategies field to given value.


### GetStaticRoutes

`func (o *PointToPointLinkIpv4Properties) GetStaticRoutes() []PointToPointStaticRoute`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PointToPointLinkIpv4Properties) GetStaticRoutesOk() (*[]PointToPointStaticRoute, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PointToPointLinkIpv4Properties) SetStaticRoutes(v []PointToPointStaticRoute)`

SetStaticRoutes sets StaticRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PointToPointLinkConfigIpv6Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetAllocationStrategies** | [**[]CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response**](CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response.md) |  | 
**StaticRoutes** | [**[]PointToPointStaticRoute**](PointToPointStaticRoute.md) | Staged IPv6 static routes. | 

## Methods

### NewPointToPointLinkConfigIpv6Properties

`func NewPointToPointLinkConfigIpv6Properties(subnetAllocationStrategies []CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response, staticRoutes []PointToPointStaticRoute, ) *PointToPointLinkConfigIpv6Properties`

NewPointToPointLinkConfigIpv6Properties instantiates a new PointToPointLinkConfigIpv6Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointToPointLinkConfigIpv6PropertiesWithDefaults

`func NewPointToPointLinkConfigIpv6PropertiesWithDefaults() *PointToPointLinkConfigIpv6Properties`

NewPointToPointLinkConfigIpv6PropertiesWithDefaults instantiates a new PointToPointLinkConfigIpv6Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetAllocationStrategies

`func (o *PointToPointLinkConfigIpv6Properties) GetSubnetAllocationStrategies() []CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response`

GetSubnetAllocationStrategies returns the SubnetAllocationStrategies field if non-nil, zero value otherwise.

### GetSubnetAllocationStrategiesOk

`func (o *PointToPointLinkConfigIpv6Properties) GetSubnetAllocationStrategiesOk() (*[]CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response, bool)`

GetSubnetAllocationStrategiesOk returns a tuple with the SubnetAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocationStrategies

`func (o *PointToPointLinkConfigIpv6Properties) SetSubnetAllocationStrategies(v []CreatePointToPointLinkConfigIpv6SubnetAllocationStrategy201Response)`

SetSubnetAllocationStrategies sets SubnetAllocationStrategies field to given value.


### GetStaticRoutes

`func (o *PointToPointLinkConfigIpv6Properties) GetStaticRoutes() []PointToPointStaticRoute`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PointToPointLinkConfigIpv6Properties) GetStaticRoutesOk() (*[]PointToPointStaticRoute, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PointToPointLinkConfigIpv6Properties) SetStaticRoutes(v []PointToPointStaticRoute)`

SetStaticRoutes sets StaticRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



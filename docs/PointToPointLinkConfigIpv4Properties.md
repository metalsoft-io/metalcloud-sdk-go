# PointToPointLinkConfigIpv4Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetAllocationStrategies** | [**[]CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response**](CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response.md) |  | 
**StaticRoutes** | [**[]PointToPointStaticRoute**](PointToPointStaticRoute.md) | Staged IPv4 static routes. | 

## Methods

### NewPointToPointLinkConfigIpv4Properties

`func NewPointToPointLinkConfigIpv4Properties(subnetAllocationStrategies []CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response, staticRoutes []PointToPointStaticRoute, ) *PointToPointLinkConfigIpv4Properties`

NewPointToPointLinkConfigIpv4Properties instantiates a new PointToPointLinkConfigIpv4Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointToPointLinkConfigIpv4PropertiesWithDefaults

`func NewPointToPointLinkConfigIpv4PropertiesWithDefaults() *PointToPointLinkConfigIpv4Properties`

NewPointToPointLinkConfigIpv4PropertiesWithDefaults instantiates a new PointToPointLinkConfigIpv4Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetAllocationStrategies

`func (o *PointToPointLinkConfigIpv4Properties) GetSubnetAllocationStrategies() []CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response`

GetSubnetAllocationStrategies returns the SubnetAllocationStrategies field if non-nil, zero value otherwise.

### GetSubnetAllocationStrategiesOk

`func (o *PointToPointLinkConfigIpv4Properties) GetSubnetAllocationStrategiesOk() (*[]CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response, bool)`

GetSubnetAllocationStrategiesOk returns a tuple with the SubnetAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAllocationStrategies

`func (o *PointToPointLinkConfigIpv4Properties) SetSubnetAllocationStrategies(v []CreatePointToPointLinkConfigIpv4SubnetAllocationStrategy201Response)`

SetSubnetAllocationStrategies sets SubnetAllocationStrategies field to given value.


### GetStaticRoutes

`func (o *PointToPointLinkConfigIpv4Properties) GetStaticRoutes() []PointToPointStaticRoute`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *PointToPointLinkConfigIpv4Properties) GetStaticRoutesOk() (*[]PointToPointStaticRoute, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *PointToPointLinkConfigIpv4Properties) SetStaticRoutes(v []PointToPointStaticRoute)`

SetStaticRoutes sets StaticRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



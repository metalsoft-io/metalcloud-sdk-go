# CreatePointToPointStaticRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPrefix** | **string** | Destination prefix in CIDR, e.g. \&quot;10.0.0.0/24\&quot; or \&quot;2001:db8::/64\&quot;. The address family must match the family of the collection (ipv4 / ipv6) it is added to. | 

## Methods

### NewCreatePointToPointStaticRoute

`func NewCreatePointToPointStaticRoute(destinationPrefix string, ) *CreatePointToPointStaticRoute`

NewCreatePointToPointStaticRoute instantiates a new CreatePointToPointStaticRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePointToPointStaticRouteWithDefaults

`func NewCreatePointToPointStaticRouteWithDefaults() *CreatePointToPointStaticRoute`

NewCreatePointToPointStaticRouteWithDefaults instantiates a new CreatePointToPointStaticRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationPrefix

`func (o *CreatePointToPointStaticRoute) GetDestinationPrefix() string`

GetDestinationPrefix returns the DestinationPrefix field if non-nil, zero value otherwise.

### GetDestinationPrefixOk

`func (o *CreatePointToPointStaticRoute) GetDestinationPrefixOk() (*string, bool)`

GetDestinationPrefixOk returns a tuple with the DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPrefix

`func (o *CreatePointToPointStaticRoute) SetDestinationPrefix(v string)`

SetDestinationPrefix sets DestinationPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



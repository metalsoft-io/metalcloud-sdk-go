# ResourcePoolStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **float32** | Number of users in the Resource Pool | 
**Servers** | **float32** | Number of servers in the Resource Pool | 
**SubnetPools** | **float32** | Number of subnet pools in the Resource Pool | 

## Methods

### NewResourcePoolStatistics

`func NewResourcePoolStatistics(users float32, servers float32, subnetPools float32, ) *ResourcePoolStatistics`

NewResourcePoolStatistics instantiates a new ResourcePoolStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolStatisticsWithDefaults

`func NewResourcePoolStatisticsWithDefaults() *ResourcePoolStatistics`

NewResourcePoolStatisticsWithDefaults instantiates a new ResourcePoolStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ResourcePoolStatistics) GetUsers() float32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ResourcePoolStatistics) GetUsersOk() (*float32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ResourcePoolStatistics) SetUsers(v float32)`

SetUsers sets Users field to given value.


### GetServers

`func (o *ResourcePoolStatistics) GetServers() float32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ResourcePoolStatistics) GetServersOk() (*float32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ResourcePoolStatistics) SetServers(v float32)`

SetServers sets Servers field to given value.


### GetSubnetPools

`func (o *ResourcePoolStatistics) GetSubnetPools() float32`

GetSubnetPools returns the SubnetPools field if non-nil, zero value otherwise.

### GetSubnetPoolsOk

`func (o *ResourcePoolStatistics) GetSubnetPoolsOk() (*float32, bool)`

GetSubnetPoolsOk returns a tuple with the SubnetPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPools

`func (o *ResourcePoolStatistics) SetSubnetPools(v float32)`

SetSubnetPools sets SubnetPools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



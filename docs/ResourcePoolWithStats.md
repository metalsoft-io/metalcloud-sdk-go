# ResourcePoolWithStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolId** | **float32** | Resource Pool ID | 
**ResourcePoolLabel** | **string** | Label of the Resource Pool | 
**ResourcePoolDescription** | **string** | Description of the Resource Pool | 
**ResourcePoolCreatedTimestamp** | **string** | Resource Pool Created Timestamp | 
**ResourcePoolUpdatedTimestamp** | **string** | Resource Pool Updated Timestamp | 
**Statistics** | [**ResourcePoolStatistics**](ResourcePoolStatistics.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewResourcePoolWithStats

`func NewResourcePoolWithStats(resourcePoolId float32, resourcePoolLabel string, resourcePoolDescription string, resourcePoolCreatedTimestamp string, resourcePoolUpdatedTimestamp string, statistics ResourcePoolStatistics, ) *ResourcePoolWithStats`

NewResourcePoolWithStats instantiates a new ResourcePoolWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolWithStatsWithDefaults

`func NewResourcePoolWithStatsWithDefaults() *ResourcePoolWithStats`

NewResourcePoolWithStatsWithDefaults instantiates a new ResourcePoolWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *ResourcePoolWithStats) GetResourcePoolId() float32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ResourcePoolWithStats) GetResourcePoolIdOk() (*float32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ResourcePoolWithStats) SetResourcePoolId(v float32)`

SetResourcePoolId sets ResourcePoolId field to given value.


### GetResourcePoolLabel

`func (o *ResourcePoolWithStats) GetResourcePoolLabel() string`

GetResourcePoolLabel returns the ResourcePoolLabel field if non-nil, zero value otherwise.

### GetResourcePoolLabelOk

`func (o *ResourcePoolWithStats) GetResourcePoolLabelOk() (*string, bool)`

GetResourcePoolLabelOk returns a tuple with the ResourcePoolLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolLabel

`func (o *ResourcePoolWithStats) SetResourcePoolLabel(v string)`

SetResourcePoolLabel sets ResourcePoolLabel field to given value.


### GetResourcePoolDescription

`func (o *ResourcePoolWithStats) GetResourcePoolDescription() string`

GetResourcePoolDescription returns the ResourcePoolDescription field if non-nil, zero value otherwise.

### GetResourcePoolDescriptionOk

`func (o *ResourcePoolWithStats) GetResourcePoolDescriptionOk() (*string, bool)`

GetResourcePoolDescriptionOk returns a tuple with the ResourcePoolDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolDescription

`func (o *ResourcePoolWithStats) SetResourcePoolDescription(v string)`

SetResourcePoolDescription sets ResourcePoolDescription field to given value.


### GetResourcePoolCreatedTimestamp

`func (o *ResourcePoolWithStats) GetResourcePoolCreatedTimestamp() string`

GetResourcePoolCreatedTimestamp returns the ResourcePoolCreatedTimestamp field if non-nil, zero value otherwise.

### GetResourcePoolCreatedTimestampOk

`func (o *ResourcePoolWithStats) GetResourcePoolCreatedTimestampOk() (*string, bool)`

GetResourcePoolCreatedTimestampOk returns a tuple with the ResourcePoolCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolCreatedTimestamp

`func (o *ResourcePoolWithStats) SetResourcePoolCreatedTimestamp(v string)`

SetResourcePoolCreatedTimestamp sets ResourcePoolCreatedTimestamp field to given value.


### GetResourcePoolUpdatedTimestamp

`func (o *ResourcePoolWithStats) GetResourcePoolUpdatedTimestamp() string`

GetResourcePoolUpdatedTimestamp returns the ResourcePoolUpdatedTimestamp field if non-nil, zero value otherwise.

### GetResourcePoolUpdatedTimestampOk

`func (o *ResourcePoolWithStats) GetResourcePoolUpdatedTimestampOk() (*string, bool)`

GetResourcePoolUpdatedTimestampOk returns a tuple with the ResourcePoolUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolUpdatedTimestamp

`func (o *ResourcePoolWithStats) SetResourcePoolUpdatedTimestamp(v string)`

SetResourcePoolUpdatedTimestamp sets ResourcePoolUpdatedTimestamp field to given value.


### GetStatistics

`func (o *ResourcePoolWithStats) GetStatistics() ResourcePoolStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ResourcePoolWithStats) GetStatisticsOk() (*ResourcePoolStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ResourcePoolWithStats) SetStatistics(v ResourcePoolStatistics)`

SetStatistics sets Statistics field to given value.


### GetLinks

`func (o *ResourcePoolWithStats) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourcePoolWithStats) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourcePoolWithStats) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourcePoolWithStats) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



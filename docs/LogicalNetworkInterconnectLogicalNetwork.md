# LogicalNetworkInterconnectLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the logical network interconnect to logical network relationship | 
**LogicalNetworkId** | **int32** | Foreign key to logical_networks | 
**LogicalNetworkInterconnectId** | **int32** | Foreign key to logical_network_interconnect | 
**Status** | [**LogicalNetworkInterconnectStatus**](LogicalNetworkInterconnectStatus.md) | Status of the logical network interconnect to logical network relationship | 
**CreatedAt** | **time.Time** | The date and time the entity was created | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the entity was last updated | [readonly] 

## Methods

### NewLogicalNetworkInterconnectLogicalNetwork

`func NewLogicalNetworkInterconnectLogicalNetwork(id string, logicalNetworkId int32, logicalNetworkInterconnectId int32, status LogicalNetworkInterconnectStatus, createdAt time.Time, updatedAt time.Time, ) *LogicalNetworkInterconnectLogicalNetwork`

NewLogicalNetworkInterconnectLogicalNetwork instantiates a new LogicalNetworkInterconnectLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkInterconnectLogicalNetworkWithDefaults

`func NewLogicalNetworkInterconnectLogicalNetworkWithDefaults() *LogicalNetworkInterconnectLogicalNetwork`

NewLogicalNetworkInterconnectLogicalNetworkWithDefaults instantiates a new LogicalNetworkInterconnectLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkInterconnectLogicalNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetLogicalNetworkId

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetLogicalNetworkId() int32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetLogicalNetworkIdOk() (*int32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *LogicalNetworkInterconnectLogicalNetwork) SetLogicalNetworkId(v int32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetLogicalNetworkInterconnectId

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetLogicalNetworkInterconnectId() int32`

GetLogicalNetworkInterconnectId returns the LogicalNetworkInterconnectId field if non-nil, zero value otherwise.

### GetLogicalNetworkInterconnectIdOk

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetLogicalNetworkInterconnectIdOk() (*int32, bool)`

GetLogicalNetworkInterconnectIdOk returns a tuple with the LogicalNetworkInterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkInterconnectId

`func (o *LogicalNetworkInterconnectLogicalNetwork) SetLogicalNetworkInterconnectId(v int32)`

SetLogicalNetworkInterconnectId sets LogicalNetworkInterconnectId field to given value.


### GetStatus

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetStatus() LogicalNetworkInterconnectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetStatusOk() (*LogicalNetworkInterconnectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogicalNetworkInterconnectLogicalNetwork) SetStatus(v LogicalNetworkInterconnectStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogicalNetworkInterconnectLogicalNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogicalNetworkInterconnectLogicalNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogicalNetworkInterconnectLogicalNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



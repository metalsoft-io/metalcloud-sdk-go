# ExternalConnectionLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The external connection logical network id. | 
**ExternalConnectionId** | **int32** | The ID of the external connection identifier this entity belongs to. | 
**LogicalNetworkId** | **int32** | The ID of the logical network identifier this entity belongs to. | 
**Status** | **string** | The status of the external connection logical network | 
**CreatedAt** | **time.Time** | The date and time the entity was created | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the entity was last updated | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExternalConnectionLogicalNetwork

`func NewExternalConnectionLogicalNetwork(id float32, externalConnectionId int32, logicalNetworkId int32, status string, createdAt time.Time, updatedAt time.Time, ) *ExternalConnectionLogicalNetwork`

NewExternalConnectionLogicalNetwork instantiates a new ExternalConnectionLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalConnectionLogicalNetworkWithDefaults

`func NewExternalConnectionLogicalNetworkWithDefaults() *ExternalConnectionLogicalNetwork`

NewExternalConnectionLogicalNetworkWithDefaults instantiates a new ExternalConnectionLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalConnectionLogicalNetwork) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalConnectionLogicalNetwork) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalConnectionLogicalNetwork) SetId(v float32)`

SetId sets Id field to given value.


### GetExternalConnectionId

`func (o *ExternalConnectionLogicalNetwork) GetExternalConnectionId() int32`

GetExternalConnectionId returns the ExternalConnectionId field if non-nil, zero value otherwise.

### GetExternalConnectionIdOk

`func (o *ExternalConnectionLogicalNetwork) GetExternalConnectionIdOk() (*int32, bool)`

GetExternalConnectionIdOk returns a tuple with the ExternalConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectionId

`func (o *ExternalConnectionLogicalNetwork) SetExternalConnectionId(v int32)`

SetExternalConnectionId sets ExternalConnectionId field to given value.


### GetLogicalNetworkId

`func (o *ExternalConnectionLogicalNetwork) GetLogicalNetworkId() int32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *ExternalConnectionLogicalNetwork) GetLogicalNetworkIdOk() (*int32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *ExternalConnectionLogicalNetwork) SetLogicalNetworkId(v int32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetStatus

`func (o *ExternalConnectionLogicalNetwork) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalConnectionLogicalNetwork) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalConnectionLogicalNetwork) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ExternalConnectionLogicalNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalConnectionLogicalNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalConnectionLogicalNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExternalConnectionLogicalNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalConnectionLogicalNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalConnectionLogicalNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *ExternalConnectionLogicalNetwork) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalConnectionLogicalNetwork) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalConnectionLogicalNetwork) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalConnectionLogicalNetwork) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



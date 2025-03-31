# InfrastructureServiceStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **float32** | Count of active infrastructures | 
**Ordered** | **float32** | Count of ordered infrastructures | 
**Deleted** | **float32** | Count of deleted infrastructures | 

## Methods

### NewInfrastructureServiceStatusDto

`func NewInfrastructureServiceStatusDto(active float32, ordered float32, deleted float32, ) *InfrastructureServiceStatusDto`

NewInfrastructureServiceStatusDto instantiates a new InfrastructureServiceStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureServiceStatusDtoWithDefaults

`func NewInfrastructureServiceStatusDtoWithDefaults() *InfrastructureServiceStatusDto`

NewInfrastructureServiceStatusDtoWithDefaults instantiates a new InfrastructureServiceStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *InfrastructureServiceStatusDto) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InfrastructureServiceStatusDto) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InfrastructureServiceStatusDto) SetActive(v float32)`

SetActive sets Active field to given value.


### GetOrdered

`func (o *InfrastructureServiceStatusDto) GetOrdered() float32`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *InfrastructureServiceStatusDto) GetOrderedOk() (*float32, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *InfrastructureServiceStatusDto) SetOrdered(v float32)`

SetOrdered sets Ordered field to given value.


### GetDeleted

`func (o *InfrastructureServiceStatusDto) GetDeleted() float32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *InfrastructureServiceStatusDto) GetDeletedOk() (*float32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *InfrastructureServiceStatusDto) SetDeleted(v float32)`

SetDeleted sets Deleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



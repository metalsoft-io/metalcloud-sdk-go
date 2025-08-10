# GetResourceUtilizationSummarized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIdOwner** | **float32** | User ID of the owner | 
**StartTimestamp** | **string** | Start timestamp for the resource utilization | 
**EndTimestamp** | **string** | End timestamp for the resource utilization | 
**InfrastructureIds** | Pointer to **[]float32** | List of infrastructure IDs | [optional] 

## Methods

### NewGetResourceUtilizationSummarized

`func NewGetResourceUtilizationSummarized(userIdOwner float32, startTimestamp string, endTimestamp string, ) *GetResourceUtilizationSummarized`

NewGetResourceUtilizationSummarized instantiates a new GetResourceUtilizationSummarized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourceUtilizationSummarizedWithDefaults

`func NewGetResourceUtilizationSummarizedWithDefaults() *GetResourceUtilizationSummarized`

NewGetResourceUtilizationSummarizedWithDefaults instantiates a new GetResourceUtilizationSummarized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIdOwner

`func (o *GetResourceUtilizationSummarized) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *GetResourceUtilizationSummarized) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *GetResourceUtilizationSummarized) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.


### GetStartTimestamp

`func (o *GetResourceUtilizationSummarized) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *GetResourceUtilizationSummarized) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *GetResourceUtilizationSummarized) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *GetResourceUtilizationSummarized) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *GetResourceUtilizationSummarized) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *GetResourceUtilizationSummarized) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetInfrastructureIds

`func (o *GetResourceUtilizationSummarized) GetInfrastructureIds() []float32`

GetInfrastructureIds returns the InfrastructureIds field if non-nil, zero value otherwise.

### GetInfrastructureIdsOk

`func (o *GetResourceUtilizationSummarized) GetInfrastructureIdsOk() (*[]float32, bool)`

GetInfrastructureIdsOk returns a tuple with the InfrastructureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureIds

`func (o *GetResourceUtilizationSummarized) SetInfrastructureIds(v []float32)`

SetInfrastructureIds sets InfrastructureIds field to given value.

### HasInfrastructureIds

`func (o *GetResourceUtilizationSummarized) HasInfrastructureIds() bool`

HasInfrastructureIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



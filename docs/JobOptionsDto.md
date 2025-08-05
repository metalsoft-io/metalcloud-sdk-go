# JobOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBeRetriedIfSuccessful** | Pointer to **bool** | Whether the job can be retried if it was successful | [optional] [default to true]
**CanBeSkipped** | Pointer to **bool** | Whether the job can be skipped | [optional] [default to true]

## Methods

### NewJobOptionsDto

`func NewJobOptionsDto() *JobOptionsDto`

NewJobOptionsDto instantiates a new JobOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobOptionsDtoWithDefaults

`func NewJobOptionsDtoWithDefaults() *JobOptionsDto`

NewJobOptionsDtoWithDefaults instantiates a new JobOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanBeRetriedIfSuccessful

`func (o *JobOptionsDto) GetCanBeRetriedIfSuccessful() bool`

GetCanBeRetriedIfSuccessful returns the CanBeRetriedIfSuccessful field if non-nil, zero value otherwise.

### GetCanBeRetriedIfSuccessfulOk

`func (o *JobOptionsDto) GetCanBeRetriedIfSuccessfulOk() (*bool, bool)`

GetCanBeRetriedIfSuccessfulOk returns a tuple with the CanBeRetriedIfSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeRetriedIfSuccessful

`func (o *JobOptionsDto) SetCanBeRetriedIfSuccessful(v bool)`

SetCanBeRetriedIfSuccessful sets CanBeRetriedIfSuccessful field to given value.

### HasCanBeRetriedIfSuccessful

`func (o *JobOptionsDto) HasCanBeRetriedIfSuccessful() bool`

HasCanBeRetriedIfSuccessful returns a boolean if a field has been set.

### GetCanBeSkipped

`func (o *JobOptionsDto) GetCanBeSkipped() bool`

GetCanBeSkipped returns the CanBeSkipped field if non-nil, zero value otherwise.

### GetCanBeSkippedOk

`func (o *JobOptionsDto) GetCanBeSkippedOk() (*bool, bool)`

GetCanBeSkippedOk returns a tuple with the CanBeSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeSkipped

`func (o *JobOptionsDto) SetCanBeSkipped(v bool)`

SetCanBeSkipped sets CanBeSkipped field to given value.

### HasCanBeSkipped

`func (o *JobOptionsDto) HasCanBeSkipped() bool`

HasCanBeSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



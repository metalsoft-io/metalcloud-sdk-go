# JobRetryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryEvenIfSuccessful** | Pointer to **bool** | Retry even if the job was successful | [optional] [default to false]

## Methods

### NewJobRetryInfo

`func NewJobRetryInfo() *JobRetryInfo`

NewJobRetryInfo instantiates a new JobRetryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRetryInfoWithDefaults

`func NewJobRetryInfoWithDefaults() *JobRetryInfo`

NewJobRetryInfoWithDefaults instantiates a new JobRetryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryEvenIfSuccessful

`func (o *JobRetryInfo) GetRetryEvenIfSuccessful() bool`

GetRetryEvenIfSuccessful returns the RetryEvenIfSuccessful field if non-nil, zero value otherwise.

### GetRetryEvenIfSuccessfulOk

`func (o *JobRetryInfo) GetRetryEvenIfSuccessfulOk() (*bool, bool)`

GetRetryEvenIfSuccessfulOk returns a tuple with the RetryEvenIfSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryEvenIfSuccessful

`func (o *JobRetryInfo) SetRetryEvenIfSuccessful(v bool)`

SetRetryEvenIfSuccessful sets RetryEvenIfSuccessful field to given value.

### HasRetryEvenIfSuccessful

`func (o *JobRetryInfo) HasRetryEvenIfSuccessful() bool`

HasRetryEvenIfSuccessful returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# JobException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionId** | **int32** | Job exception Id | 
**ArchiveId** | Pointer to **int32** | Job archive Id | [optional] 
**JobId** | Pointer to **int32** | Job Id | [optional] 
**Exception** | Pointer to **map[string]interface{}** | The exception that was thrown by the function | [optional] 
**CreatedTimestamp** | **string** | The timestamp when the exception was created | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewJobException

`func NewJobException(exceptionId int32, createdTimestamp string, links map[string]interface{}, ) *JobException`

NewJobException instantiates a new JobException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobExceptionWithDefaults

`func NewJobExceptionWithDefaults() *JobException`

NewJobExceptionWithDefaults instantiates a new JobException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptionId

`func (o *JobException) GetExceptionId() int32`

GetExceptionId returns the ExceptionId field if non-nil, zero value otherwise.

### GetExceptionIdOk

`func (o *JobException) GetExceptionIdOk() (*int32, bool)`

GetExceptionIdOk returns a tuple with the ExceptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionId

`func (o *JobException) SetExceptionId(v int32)`

SetExceptionId sets ExceptionId field to given value.


### GetArchiveId

`func (o *JobException) GetArchiveId() int32`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *JobException) GetArchiveIdOk() (*int32, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *JobException) SetArchiveId(v int32)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *JobException) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetJobId

`func (o *JobException) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobException) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobException) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *JobException) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetException

`func (o *JobException) GetException() map[string]interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *JobException) GetExceptionOk() (*map[string]interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *JobException) SetException(v map[string]interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *JobException) HasException() bool`

HasException returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *JobException) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *JobException) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *JobException) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLinks

`func (o *JobException) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *JobException) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *JobException) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



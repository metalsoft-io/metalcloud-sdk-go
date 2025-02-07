# JobGroupStatisticsWithoutIdDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupCreatedTimestamp** | **string** | Group created timestamp | 
**GroupCompletedTimestamp** | **string** | Group completed timestamp | 
**JobsThrownError** | **int32** | Total count of jobs with errors | 
**JobsCompleted** | **int32** | Total count of completed jobs | 
**JobsTotal** | **int32** | Total count of jobs | 

## Methods

### NewJobGroupStatisticsWithoutIdDto

`func NewJobGroupStatisticsWithoutIdDto(groupCreatedTimestamp string, groupCompletedTimestamp string, jobsThrownError int32, jobsCompleted int32, jobsTotal int32, ) *JobGroupStatisticsWithoutIdDto`

NewJobGroupStatisticsWithoutIdDto instantiates a new JobGroupStatisticsWithoutIdDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobGroupStatisticsWithoutIdDtoWithDefaults

`func NewJobGroupStatisticsWithoutIdDtoWithDefaults() *JobGroupStatisticsWithoutIdDto`

NewJobGroupStatisticsWithoutIdDtoWithDefaults instantiates a new JobGroupStatisticsWithoutIdDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupCreatedTimestamp

`func (o *JobGroupStatisticsWithoutIdDto) GetGroupCreatedTimestamp() string`

GetGroupCreatedTimestamp returns the GroupCreatedTimestamp field if non-nil, zero value otherwise.

### GetGroupCreatedTimestampOk

`func (o *JobGroupStatisticsWithoutIdDto) GetGroupCreatedTimestampOk() (*string, bool)`

GetGroupCreatedTimestampOk returns a tuple with the GroupCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCreatedTimestamp

`func (o *JobGroupStatisticsWithoutIdDto) SetGroupCreatedTimestamp(v string)`

SetGroupCreatedTimestamp sets GroupCreatedTimestamp field to given value.


### GetGroupCompletedTimestamp

`func (o *JobGroupStatisticsWithoutIdDto) GetGroupCompletedTimestamp() string`

GetGroupCompletedTimestamp returns the GroupCompletedTimestamp field if non-nil, zero value otherwise.

### GetGroupCompletedTimestampOk

`func (o *JobGroupStatisticsWithoutIdDto) GetGroupCompletedTimestampOk() (*string, bool)`

GetGroupCompletedTimestampOk returns a tuple with the GroupCompletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCompletedTimestamp

`func (o *JobGroupStatisticsWithoutIdDto) SetGroupCompletedTimestamp(v string)`

SetGroupCompletedTimestamp sets GroupCompletedTimestamp field to given value.


### GetJobsThrownError

`func (o *JobGroupStatisticsWithoutIdDto) GetJobsThrownError() int32`

GetJobsThrownError returns the JobsThrownError field if non-nil, zero value otherwise.

### GetJobsThrownErrorOk

`func (o *JobGroupStatisticsWithoutIdDto) GetJobsThrownErrorOk() (*int32, bool)`

GetJobsThrownErrorOk returns a tuple with the JobsThrownError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsThrownError

`func (o *JobGroupStatisticsWithoutIdDto) SetJobsThrownError(v int32)`

SetJobsThrownError sets JobsThrownError field to given value.


### GetJobsCompleted

`func (o *JobGroupStatisticsWithoutIdDto) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *JobGroupStatisticsWithoutIdDto) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *JobGroupStatisticsWithoutIdDto) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.


### GetJobsTotal

`func (o *JobGroupStatisticsWithoutIdDto) GetJobsTotal() int32`

GetJobsTotal returns the JobsTotal field if non-nil, zero value otherwise.

### GetJobsTotalOk

`func (o *JobGroupStatisticsWithoutIdDto) GetJobsTotalOk() (*int32, bool)`

GetJobsTotalOk returns a tuple with the JobsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsTotal

`func (o *JobGroupStatisticsWithoutIdDto) SetJobsTotal(v int32)`

SetJobsTotal sets JobsTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



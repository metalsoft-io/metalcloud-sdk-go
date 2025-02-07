# JobGroupStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | Group Id | 
**GroupCreatedTimestamp** | **string** | Group created timestamp | 
**GroupCompletedTimestamp** | **string** | Group completed timestamp | 
**JobsThrownError** | **int32** | Total count of jobs with errors | 
**JobsCompleted** | **int32** | Total count of completed jobs | 
**JobsTotal** | **int32** | Total count of jobs | 

## Methods

### NewJobGroupStatistics

`func NewJobGroupStatistics(groupId int32, groupCreatedTimestamp string, groupCompletedTimestamp string, jobsThrownError int32, jobsCompleted int32, jobsTotal int32, ) *JobGroupStatistics`

NewJobGroupStatistics instantiates a new JobGroupStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobGroupStatisticsWithDefaults

`func NewJobGroupStatisticsWithDefaults() *JobGroupStatistics`

NewJobGroupStatisticsWithDefaults instantiates a new JobGroupStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *JobGroupStatistics) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *JobGroupStatistics) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *JobGroupStatistics) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupCreatedTimestamp

`func (o *JobGroupStatistics) GetGroupCreatedTimestamp() string`

GetGroupCreatedTimestamp returns the GroupCreatedTimestamp field if non-nil, zero value otherwise.

### GetGroupCreatedTimestampOk

`func (o *JobGroupStatistics) GetGroupCreatedTimestampOk() (*string, bool)`

GetGroupCreatedTimestampOk returns a tuple with the GroupCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCreatedTimestamp

`func (o *JobGroupStatistics) SetGroupCreatedTimestamp(v string)`

SetGroupCreatedTimestamp sets GroupCreatedTimestamp field to given value.


### GetGroupCompletedTimestamp

`func (o *JobGroupStatistics) GetGroupCompletedTimestamp() string`

GetGroupCompletedTimestamp returns the GroupCompletedTimestamp field if non-nil, zero value otherwise.

### GetGroupCompletedTimestampOk

`func (o *JobGroupStatistics) GetGroupCompletedTimestampOk() (*string, bool)`

GetGroupCompletedTimestampOk returns a tuple with the GroupCompletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCompletedTimestamp

`func (o *JobGroupStatistics) SetGroupCompletedTimestamp(v string)`

SetGroupCompletedTimestamp sets GroupCompletedTimestamp field to given value.


### GetJobsThrownError

`func (o *JobGroupStatistics) GetJobsThrownError() int32`

GetJobsThrownError returns the JobsThrownError field if non-nil, zero value otherwise.

### GetJobsThrownErrorOk

`func (o *JobGroupStatistics) GetJobsThrownErrorOk() (*int32, bool)`

GetJobsThrownErrorOk returns a tuple with the JobsThrownError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsThrownError

`func (o *JobGroupStatistics) SetJobsThrownError(v int32)`

SetJobsThrownError sets JobsThrownError field to given value.


### GetJobsCompleted

`func (o *JobGroupStatistics) GetJobsCompleted() int32`

GetJobsCompleted returns the JobsCompleted field if non-nil, zero value otherwise.

### GetJobsCompletedOk

`func (o *JobGroupStatistics) GetJobsCompletedOk() (*int32, bool)`

GetJobsCompletedOk returns a tuple with the JobsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCompleted

`func (o *JobGroupStatistics) SetJobsCompleted(v int32)`

SetJobsCompleted sets JobsCompleted field to given value.


### GetJobsTotal

`func (o *JobGroupStatistics) GetJobsTotal() int32`

GetJobsTotal returns the JobsTotal field if non-nil, zero value otherwise.

### GetJobsTotalOk

`func (o *JobGroupStatistics) GetJobsTotalOk() (*int32, bool)`

GetJobsTotalOk returns a tuple with the JobsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsTotal

`func (o *JobGroupStatistics) SetJobsTotal(v int32)`

SetJobsTotal sets JobsTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



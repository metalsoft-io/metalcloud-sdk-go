# JobStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusToCount** | **map[string]interface{}** | The number of jobs for each status | 
**ArchivedCount** | **int32** | The total number of jobs in the archive | 

## Methods

### NewJobStatistics

`func NewJobStatistics(statusToCount map[string]interface{}, archivedCount int32, ) *JobStatistics`

NewJobStatistics instantiates a new JobStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStatisticsWithDefaults

`func NewJobStatisticsWithDefaults() *JobStatistics`

NewJobStatisticsWithDefaults instantiates a new JobStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusToCount

`func (o *JobStatistics) GetStatusToCount() map[string]interface{}`

GetStatusToCount returns the StatusToCount field if non-nil, zero value otherwise.

### GetStatusToCountOk

`func (o *JobStatistics) GetStatusToCountOk() (*map[string]interface{}, bool)`

GetStatusToCountOk returns a tuple with the StatusToCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusToCount

`func (o *JobStatistics) SetStatusToCount(v map[string]interface{})`

SetStatusToCount sets StatusToCount field to given value.


### GetArchivedCount

`func (o *JobStatistics) GetArchivedCount() int32`

GetArchivedCount returns the ArchivedCount field if non-nil, zero value otherwise.

### GetArchivedCountOk

`func (o *JobStatistics) GetArchivedCountOk() (*int32, bool)`

GetArchivedCountOk returns a tuple with the ArchivedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedCount

`func (o *JobStatistics) SetArchivedCount(v int32)`

SetArchivedCount sets ArchivedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



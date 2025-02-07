# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **int32** | Job Id | 
**SiteId** | Pointer to **int32** | Site Id | [optional] 
**InstanceId** | Pointer to **int32** | Instance Id | [optional] 
**ServerId** | Pointer to **int32** | Server Id | [optional] 
**VmPoolId** | Pointer to **int32** | VM Pool Id | [optional] 
**StorageId** | Pointer to **int32** | Storage Pool Id | [optional] 
**NetworkDeviceId** | Pointer to **int32** | Network Equipment Id | [optional] 
**InfrastructureId** | Pointer to **int32** | Infrastructure Id | [optional] 
**JobIdBlocked** | Pointer to **int32** | The id of the next job that should be executed after this job | [optional] 
**JobIdBlockedBy** | Pointer to **int32** | The id of the job that is blocking this job from executing | [optional] 
**JobGroupId** | Pointer to **int32** | The id of the group that this job belongs to | [optional] 
**Type** | **string** | The type of the job | 
**Status** | **string** | The status of the job | 
**FunctionName** | **string** | The name of the function that should be executed | 
**Params** | Pointer to **map[string]interface{}** | The parameters that should be passed to the function | [optional] 
**Response** | Pointer to **map[string]interface{}** | The response of the function | [optional] 
**CallCount** | **int32** | The number of times the job has been called | 
**RetryMax** | **int32** | The maximum number of times the job should be retried | 
**RetryCount** | **int32** | The number of times the job has been retried | 
**RetryMinSeconds** | **int32** | The minimum number of seconds that should pass before the job is retried | 
**Exception** | Pointer to **map[string]interface{}** | The exception that was thrown by the function | [optional] 
**ExtraInfo** | Pointer to **map[string]interface{}** | Extra information about the job | [optional] 
**DurationMilliseconds** | Pointer to **int32** | The time in milliseconds that the job took to execute | [optional] 
**CreatedTimestamp** | **string** | The timestamp when the job was created | 
**UpdatedTimestamp** | **string** | The timestamp when the job was last updated | 
**StartTimestamp** | Pointer to **string** | The timestamp when the job was started | [optional] 
**MarkForDeath** | Pointer to **string** | Mark the job for death | [optional] 
**UniqueIdentifier** | Pointer to **string** | The unique identifier of the job | [optional] 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewJob

`func NewJob(jobId int32, type_ string, status string, functionName string, callCount int32, retryMax int32, retryCount int32, retryMinSeconds int32, createdTimestamp string, updatedTimestamp string, links map[string]interface{}, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *Job) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Job) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Job) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetSiteId

`func (o *Job) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Job) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Job) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Job) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetInstanceId

`func (o *Job) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Job) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Job) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *Job) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetServerId

`func (o *Job) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Job) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Job) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *Job) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetVmPoolId

`func (o *Job) GetVmPoolId() int32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *Job) GetVmPoolIdOk() (*int32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *Job) SetVmPoolId(v int32)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *Job) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetStorageId

`func (o *Job) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *Job) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *Job) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *Job) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetNetworkDeviceId

`func (o *Job) GetNetworkDeviceId() int32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *Job) GetNetworkDeviceIdOk() (*int32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *Job) SetNetworkDeviceId(v int32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *Job) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *Job) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Job) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Job) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *Job) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetJobIdBlocked

`func (o *Job) GetJobIdBlocked() int32`

GetJobIdBlocked returns the JobIdBlocked field if non-nil, zero value otherwise.

### GetJobIdBlockedOk

`func (o *Job) GetJobIdBlockedOk() (*int32, bool)`

GetJobIdBlockedOk returns a tuple with the JobIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdBlocked

`func (o *Job) SetJobIdBlocked(v int32)`

SetJobIdBlocked sets JobIdBlocked field to given value.

### HasJobIdBlocked

`func (o *Job) HasJobIdBlocked() bool`

HasJobIdBlocked returns a boolean if a field has been set.

### GetJobIdBlockedBy

`func (o *Job) GetJobIdBlockedBy() int32`

GetJobIdBlockedBy returns the JobIdBlockedBy field if non-nil, zero value otherwise.

### GetJobIdBlockedByOk

`func (o *Job) GetJobIdBlockedByOk() (*int32, bool)`

GetJobIdBlockedByOk returns a tuple with the JobIdBlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdBlockedBy

`func (o *Job) SetJobIdBlockedBy(v int32)`

SetJobIdBlockedBy sets JobIdBlockedBy field to given value.

### HasJobIdBlockedBy

`func (o *Job) HasJobIdBlockedBy() bool`

HasJobIdBlockedBy returns a boolean if a field has been set.

### GetJobGroupId

`func (o *Job) GetJobGroupId() int32`

GetJobGroupId returns the JobGroupId field if non-nil, zero value otherwise.

### GetJobGroupIdOk

`func (o *Job) GetJobGroupIdOk() (*int32, bool)`

GetJobGroupIdOk returns a tuple with the JobGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobGroupId

`func (o *Job) SetJobGroupId(v int32)`

SetJobGroupId sets JobGroupId field to given value.

### HasJobGroupId

`func (o *Job) HasJobGroupId() bool`

HasJobGroupId returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFunctionName

`func (o *Job) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *Job) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *Job) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetParams

`func (o *Job) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Job) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Job) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *Job) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetResponse

`func (o *Job) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Job) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Job) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Job) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCallCount

`func (o *Job) GetCallCount() int32`

GetCallCount returns the CallCount field if non-nil, zero value otherwise.

### GetCallCountOk

`func (o *Job) GetCallCountOk() (*int32, bool)`

GetCallCountOk returns a tuple with the CallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallCount

`func (o *Job) SetCallCount(v int32)`

SetCallCount sets CallCount field to given value.


### GetRetryMax

`func (o *Job) GetRetryMax() int32`

GetRetryMax returns the RetryMax field if non-nil, zero value otherwise.

### GetRetryMaxOk

`func (o *Job) GetRetryMaxOk() (*int32, bool)`

GetRetryMaxOk returns a tuple with the RetryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMax

`func (o *Job) SetRetryMax(v int32)`

SetRetryMax sets RetryMax field to given value.


### GetRetryCount

`func (o *Job) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Job) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Job) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.


### GetRetryMinSeconds

`func (o *Job) GetRetryMinSeconds() int32`

GetRetryMinSeconds returns the RetryMinSeconds field if non-nil, zero value otherwise.

### GetRetryMinSecondsOk

`func (o *Job) GetRetryMinSecondsOk() (*int32, bool)`

GetRetryMinSecondsOk returns a tuple with the RetryMinSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMinSeconds

`func (o *Job) SetRetryMinSeconds(v int32)`

SetRetryMinSeconds sets RetryMinSeconds field to given value.


### GetException

`func (o *Job) GetException() map[string]interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *Job) GetExceptionOk() (*map[string]interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *Job) SetException(v map[string]interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *Job) HasException() bool`

HasException returns a boolean if a field has been set.

### GetExtraInfo

`func (o *Job) GetExtraInfo() map[string]interface{}`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *Job) GetExtraInfoOk() (*map[string]interface{}, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *Job) SetExtraInfo(v map[string]interface{})`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *Job) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetDurationMilliseconds

`func (o *Job) GetDurationMilliseconds() int32`

GetDurationMilliseconds returns the DurationMilliseconds field if non-nil, zero value otherwise.

### GetDurationMillisecondsOk

`func (o *Job) GetDurationMillisecondsOk() (*int32, bool)`

GetDurationMillisecondsOk returns a tuple with the DurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMilliseconds

`func (o *Job) SetDurationMilliseconds(v int32)`

SetDurationMilliseconds sets DurationMilliseconds field to given value.

### HasDurationMilliseconds

`func (o *Job) HasDurationMilliseconds() bool`

HasDurationMilliseconds returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Job) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Job) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Job) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Job) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Job) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Job) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStartTimestamp

`func (o *Job) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Job) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Job) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *Job) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetMarkForDeath

`func (o *Job) GetMarkForDeath() string`

GetMarkForDeath returns the MarkForDeath field if non-nil, zero value otherwise.

### GetMarkForDeathOk

`func (o *Job) GetMarkForDeathOk() (*string, bool)`

GetMarkForDeathOk returns a tuple with the MarkForDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkForDeath

`func (o *Job) SetMarkForDeath(v string)`

SetMarkForDeath sets MarkForDeath field to given value.

### HasMarkForDeath

`func (o *Job) HasMarkForDeath() bool`

HasMarkForDeath returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *Job) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *Job) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *Job) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *Job) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.

### GetLinks

`func (o *Job) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Job) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Job) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



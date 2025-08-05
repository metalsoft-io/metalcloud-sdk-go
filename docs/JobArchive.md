# JobArchive

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
**RequiresConfirmation** | **bool** | Whether the job requires confirmation before execution | [default to false]
**Options** | [**JobOptionsDto**](JobOptionsDto.md) | Options for the job | 
**CreatedTimestamp** | **string** | The timestamp when the job was created | 
**UpdatedTimestamp** | **string** | The timestamp when the job was last updated | 
**StartTimestamp** | Pointer to **string** | The timestamp when the job was started | [optional] 
**MarkForDeath** | Pointer to **string** | Mark the job for death | [optional] 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewJobArchive

`func NewJobArchive(jobId int32, type_ string, status string, functionName string, callCount int32, retryMax int32, retryCount int32, retryMinSeconds int32, requiresConfirmation bool, options JobOptionsDto, createdTimestamp string, updatedTimestamp string, links map[string]interface{}, ) *JobArchive`

NewJobArchive instantiates a new JobArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobArchiveWithDefaults

`func NewJobArchiveWithDefaults() *JobArchive`

NewJobArchiveWithDefaults instantiates a new JobArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobArchive) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobArchive) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobArchive) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetSiteId

`func (o *JobArchive) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *JobArchive) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *JobArchive) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *JobArchive) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetInstanceId

`func (o *JobArchive) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *JobArchive) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *JobArchive) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *JobArchive) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetServerId

`func (o *JobArchive) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *JobArchive) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *JobArchive) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *JobArchive) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetVmPoolId

`func (o *JobArchive) GetVmPoolId() int32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *JobArchive) GetVmPoolIdOk() (*int32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *JobArchive) SetVmPoolId(v int32)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *JobArchive) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetStorageId

`func (o *JobArchive) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *JobArchive) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *JobArchive) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *JobArchive) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetNetworkDeviceId

`func (o *JobArchive) GetNetworkDeviceId() int32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *JobArchive) GetNetworkDeviceIdOk() (*int32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *JobArchive) SetNetworkDeviceId(v int32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *JobArchive) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *JobArchive) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *JobArchive) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *JobArchive) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *JobArchive) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetJobIdBlocked

`func (o *JobArchive) GetJobIdBlocked() int32`

GetJobIdBlocked returns the JobIdBlocked field if non-nil, zero value otherwise.

### GetJobIdBlockedOk

`func (o *JobArchive) GetJobIdBlockedOk() (*int32, bool)`

GetJobIdBlockedOk returns a tuple with the JobIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdBlocked

`func (o *JobArchive) SetJobIdBlocked(v int32)`

SetJobIdBlocked sets JobIdBlocked field to given value.

### HasJobIdBlocked

`func (o *JobArchive) HasJobIdBlocked() bool`

HasJobIdBlocked returns a boolean if a field has been set.

### GetJobIdBlockedBy

`func (o *JobArchive) GetJobIdBlockedBy() int32`

GetJobIdBlockedBy returns the JobIdBlockedBy field if non-nil, zero value otherwise.

### GetJobIdBlockedByOk

`func (o *JobArchive) GetJobIdBlockedByOk() (*int32, bool)`

GetJobIdBlockedByOk returns a tuple with the JobIdBlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdBlockedBy

`func (o *JobArchive) SetJobIdBlockedBy(v int32)`

SetJobIdBlockedBy sets JobIdBlockedBy field to given value.

### HasJobIdBlockedBy

`func (o *JobArchive) HasJobIdBlockedBy() bool`

HasJobIdBlockedBy returns a boolean if a field has been set.

### GetJobGroupId

`func (o *JobArchive) GetJobGroupId() int32`

GetJobGroupId returns the JobGroupId field if non-nil, zero value otherwise.

### GetJobGroupIdOk

`func (o *JobArchive) GetJobGroupIdOk() (*int32, bool)`

GetJobGroupIdOk returns a tuple with the JobGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobGroupId

`func (o *JobArchive) SetJobGroupId(v int32)`

SetJobGroupId sets JobGroupId field to given value.

### HasJobGroupId

`func (o *JobArchive) HasJobGroupId() bool`

HasJobGroupId returns a boolean if a field has been set.

### GetType

`func (o *JobArchive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobArchive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobArchive) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *JobArchive) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobArchive) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobArchive) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFunctionName

`func (o *JobArchive) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *JobArchive) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *JobArchive) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetParams

`func (o *JobArchive) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *JobArchive) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *JobArchive) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *JobArchive) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetResponse

`func (o *JobArchive) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *JobArchive) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *JobArchive) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *JobArchive) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCallCount

`func (o *JobArchive) GetCallCount() int32`

GetCallCount returns the CallCount field if non-nil, zero value otherwise.

### GetCallCountOk

`func (o *JobArchive) GetCallCountOk() (*int32, bool)`

GetCallCountOk returns a tuple with the CallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallCount

`func (o *JobArchive) SetCallCount(v int32)`

SetCallCount sets CallCount field to given value.


### GetRetryMax

`func (o *JobArchive) GetRetryMax() int32`

GetRetryMax returns the RetryMax field if non-nil, zero value otherwise.

### GetRetryMaxOk

`func (o *JobArchive) GetRetryMaxOk() (*int32, bool)`

GetRetryMaxOk returns a tuple with the RetryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMax

`func (o *JobArchive) SetRetryMax(v int32)`

SetRetryMax sets RetryMax field to given value.


### GetRetryCount

`func (o *JobArchive) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *JobArchive) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *JobArchive) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.


### GetRetryMinSeconds

`func (o *JobArchive) GetRetryMinSeconds() int32`

GetRetryMinSeconds returns the RetryMinSeconds field if non-nil, zero value otherwise.

### GetRetryMinSecondsOk

`func (o *JobArchive) GetRetryMinSecondsOk() (*int32, bool)`

GetRetryMinSecondsOk returns a tuple with the RetryMinSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMinSeconds

`func (o *JobArchive) SetRetryMinSeconds(v int32)`

SetRetryMinSeconds sets RetryMinSeconds field to given value.


### GetException

`func (o *JobArchive) GetException() map[string]interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *JobArchive) GetExceptionOk() (*map[string]interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *JobArchive) SetException(v map[string]interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *JobArchive) HasException() bool`

HasException returns a boolean if a field has been set.

### GetExtraInfo

`func (o *JobArchive) GetExtraInfo() map[string]interface{}`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *JobArchive) GetExtraInfoOk() (*map[string]interface{}, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *JobArchive) SetExtraInfo(v map[string]interface{})`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *JobArchive) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetDurationMilliseconds

`func (o *JobArchive) GetDurationMilliseconds() int32`

GetDurationMilliseconds returns the DurationMilliseconds field if non-nil, zero value otherwise.

### GetDurationMillisecondsOk

`func (o *JobArchive) GetDurationMillisecondsOk() (*int32, bool)`

GetDurationMillisecondsOk returns a tuple with the DurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMilliseconds

`func (o *JobArchive) SetDurationMilliseconds(v int32)`

SetDurationMilliseconds sets DurationMilliseconds field to given value.

### HasDurationMilliseconds

`func (o *JobArchive) HasDurationMilliseconds() bool`

HasDurationMilliseconds returns a boolean if a field has been set.

### GetRequiresConfirmation

`func (o *JobArchive) GetRequiresConfirmation() bool`

GetRequiresConfirmation returns the RequiresConfirmation field if non-nil, zero value otherwise.

### GetRequiresConfirmationOk

`func (o *JobArchive) GetRequiresConfirmationOk() (*bool, bool)`

GetRequiresConfirmationOk returns a tuple with the RequiresConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresConfirmation

`func (o *JobArchive) SetRequiresConfirmation(v bool)`

SetRequiresConfirmation sets RequiresConfirmation field to given value.


### GetOptions

`func (o *JobArchive) GetOptions() JobOptionsDto`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JobArchive) GetOptionsOk() (*JobOptionsDto, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JobArchive) SetOptions(v JobOptionsDto)`

SetOptions sets Options field to given value.


### GetCreatedTimestamp

`func (o *JobArchive) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *JobArchive) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *JobArchive) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *JobArchive) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *JobArchive) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *JobArchive) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStartTimestamp

`func (o *JobArchive) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *JobArchive) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *JobArchive) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *JobArchive) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetMarkForDeath

`func (o *JobArchive) GetMarkForDeath() string`

GetMarkForDeath returns the MarkForDeath field if non-nil, zero value otherwise.

### GetMarkForDeathOk

`func (o *JobArchive) GetMarkForDeathOk() (*string, bool)`

GetMarkForDeathOk returns a tuple with the MarkForDeath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkForDeath

`func (o *JobArchive) SetMarkForDeath(v string)`

SetMarkForDeath sets MarkForDeath field to given value.

### HasMarkForDeath

`func (o *JobArchive) HasMarkForDeath() bool`

HasMarkForDeath returns a boolean if a field has been set.

### GetLinks

`func (o *JobArchive) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *JobArchive) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *JobArchive) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



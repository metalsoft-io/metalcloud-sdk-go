# JobGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Group Id | 
**Type** | **string** | Group type | 
**Context** | Pointer to **map[string]interface{}** | Group context parameters | [optional] 
**Description** | **string** | Group description | 
**CreatedTimestamp** | **string** | Group created timestamp | 
**FinishedTimestamp** | Pointer to **string** | Group finished timestamp | [optional] 
**Params** | Pointer to **map[string]interface{}** | Group parameters | [optional] 
**Archived** | Pointer to **int32** | Group archived status | [optional] 
**InfrastructureId** | Pointer to **int64** | Infrastructure Id | [optional] 
**DriveId** | Pointer to **int64** | Drive Id | [optional] 
**ServerId** | Pointer to **int64** | Server Id | [optional] 
**NetworkDeviceId** | Pointer to **int64** | Network device Id | [optional] 
**VmPoolId** | Pointer to **int64** | VM Pool Id | [optional] 
**StorageId** | Pointer to **int64** | Storage Pool Id | [optional] 
**CancelReason** | Pointer to **string** | Cancel reason of the group | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewJobGroup

`func NewJobGroup(id int64, type_ string, description string, createdTimestamp string, ) *JobGroup`

NewJobGroup instantiates a new JobGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobGroupWithDefaults

`func NewJobGroupWithDefaults() *JobGroup`

NewJobGroupWithDefaults instantiates a new JobGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobGroup) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *JobGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobGroup) SetType(v string)`

SetType sets Type field to given value.


### GetContext

`func (o *JobGroup) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *JobGroup) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *JobGroup) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *JobGroup) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDescription

`func (o *JobGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobGroup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTimestamp

`func (o *JobGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *JobGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *JobGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetFinishedTimestamp

`func (o *JobGroup) GetFinishedTimestamp() string`

GetFinishedTimestamp returns the FinishedTimestamp field if non-nil, zero value otherwise.

### GetFinishedTimestampOk

`func (o *JobGroup) GetFinishedTimestampOk() (*string, bool)`

GetFinishedTimestampOk returns a tuple with the FinishedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTimestamp

`func (o *JobGroup) SetFinishedTimestamp(v string)`

SetFinishedTimestamp sets FinishedTimestamp field to given value.

### HasFinishedTimestamp

`func (o *JobGroup) HasFinishedTimestamp() bool`

HasFinishedTimestamp returns a boolean if a field has been set.

### GetParams

`func (o *JobGroup) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *JobGroup) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *JobGroup) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *JobGroup) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetArchived

`func (o *JobGroup) GetArchived() int32`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *JobGroup) GetArchivedOk() (*int32, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *JobGroup) SetArchived(v int32)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *JobGroup) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *JobGroup) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *JobGroup) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *JobGroup) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *JobGroup) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetDriveId

`func (o *JobGroup) GetDriveId() int64`

GetDriveId returns the DriveId field if non-nil, zero value otherwise.

### GetDriveIdOk

`func (o *JobGroup) GetDriveIdOk() (*int64, bool)`

GetDriveIdOk returns a tuple with the DriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveId

`func (o *JobGroup) SetDriveId(v int64)`

SetDriveId sets DriveId field to given value.

### HasDriveId

`func (o *JobGroup) HasDriveId() bool`

HasDriveId returns a boolean if a field has been set.

### GetServerId

`func (o *JobGroup) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *JobGroup) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *JobGroup) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *JobGroup) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetNetworkDeviceId

`func (o *JobGroup) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *JobGroup) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *JobGroup) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *JobGroup) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### GetVmPoolId

`func (o *JobGroup) GetVmPoolId() int64`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *JobGroup) GetVmPoolIdOk() (*int64, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *JobGroup) SetVmPoolId(v int64)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *JobGroup) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetStorageId

`func (o *JobGroup) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *JobGroup) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *JobGroup) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *JobGroup) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetCancelReason

`func (o *JobGroup) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *JobGroup) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *JobGroup) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *JobGroup) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetLinks

`func (o *JobGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *JobGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *JobGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *JobGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



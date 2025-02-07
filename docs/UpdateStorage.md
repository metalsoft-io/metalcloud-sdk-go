# UpdateStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **map[string]interface{}** | Options for the storage | [optional] 
**InMaintenance** | Pointer to **float32** | Specifies if the storage is in maintenance | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**PortGroupAllocationOrder** | Pointer to **map[string]interface{}** | Port group allocation order | [optional] 
**PortGroupPhysicalPorts** | Pointer to **map[string]interface{}** | Port group physical ports | [optional] 
**DefaultIoLimitPolicy** | Pointer to **string** | Default IO limit policy | [optional] 
**ArrayId** | Pointer to **string** | Array id | [optional] 
**DirectorId** | Pointer to **string** | Director id | [optional] 
**S3Hostname** | Pointer to **string** | S3 hostname | [optional] 
**S3Port** | Pointer to **string** | S3 port | [optional] 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Username** | Pointer to **string** | Username | [optional] 
**Password** | Pointer to **string** | The password to use. | [optional] 

## Methods

### NewUpdateStorage

`func NewUpdateStorage() *UpdateStorage`

NewUpdateStorage instantiates a new UpdateStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageWithDefaults

`func NewUpdateStorageWithDefaults() *UpdateStorage`

NewUpdateStorageWithDefaults instantiates a new UpdateStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *UpdateStorage) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateStorage) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateStorage) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateStorage) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetInMaintenance

`func (o *UpdateStorage) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *UpdateStorage) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *UpdateStorage) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *UpdateStorage) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateStorage) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateStorage) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateStorage) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateStorage) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetDrivePriority

`func (o *UpdateStorage) GetDrivePriority() float32`

GetDrivePriority returns the DrivePriority field if non-nil, zero value otherwise.

### GetDrivePriorityOk

`func (o *UpdateStorage) GetDrivePriorityOk() (*float32, bool)`

GetDrivePriorityOk returns a tuple with the DrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePriority

`func (o *UpdateStorage) SetDrivePriority(v float32)`

SetDrivePriority sets DrivePriority field to given value.

### HasDrivePriority

`func (o *UpdateStorage) HasDrivePriority() bool`

HasDrivePriority returns a boolean if a field has been set.

### GetSharedDrivePriority

`func (o *UpdateStorage) GetSharedDrivePriority() float32`

GetSharedDrivePriority returns the SharedDrivePriority field if non-nil, zero value otherwise.

### GetSharedDrivePriorityOk

`func (o *UpdateStorage) GetSharedDrivePriorityOk() (*float32, bool)`

GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrivePriority

`func (o *UpdateStorage) SetSharedDrivePriority(v float32)`

SetSharedDrivePriority sets SharedDrivePriority field to given value.

### HasSharedDrivePriority

`func (o *UpdateStorage) HasSharedDrivePriority() bool`

HasSharedDrivePriority returns a boolean if a field has been set.

### GetTags

`func (o *UpdateStorage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateStorage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateStorage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateStorage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPortGroupAllocationOrder

`func (o *UpdateStorage) GetPortGroupAllocationOrder() map[string]interface{}`

GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field if non-nil, zero value otherwise.

### GetPortGroupAllocationOrderOk

`func (o *UpdateStorage) GetPortGroupAllocationOrderOk() (*map[string]interface{}, bool)`

GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupAllocationOrder

`func (o *UpdateStorage) SetPortGroupAllocationOrder(v map[string]interface{})`

SetPortGroupAllocationOrder sets PortGroupAllocationOrder field to given value.

### HasPortGroupAllocationOrder

`func (o *UpdateStorage) HasPortGroupAllocationOrder() bool`

HasPortGroupAllocationOrder returns a boolean if a field has been set.

### GetPortGroupPhysicalPorts

`func (o *UpdateStorage) GetPortGroupPhysicalPorts() map[string]interface{}`

GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field if non-nil, zero value otherwise.

### GetPortGroupPhysicalPortsOk

`func (o *UpdateStorage) GetPortGroupPhysicalPortsOk() (*map[string]interface{}, bool)`

GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupPhysicalPorts

`func (o *UpdateStorage) SetPortGroupPhysicalPorts(v map[string]interface{})`

SetPortGroupPhysicalPorts sets PortGroupPhysicalPorts field to given value.

### HasPortGroupPhysicalPorts

`func (o *UpdateStorage) HasPortGroupPhysicalPorts() bool`

HasPortGroupPhysicalPorts returns a boolean if a field has been set.

### GetDefaultIoLimitPolicy

`func (o *UpdateStorage) GetDefaultIoLimitPolicy() string`

GetDefaultIoLimitPolicy returns the DefaultIoLimitPolicy field if non-nil, zero value otherwise.

### GetDefaultIoLimitPolicyOk

`func (o *UpdateStorage) GetDefaultIoLimitPolicyOk() (*string, bool)`

GetDefaultIoLimitPolicyOk returns a tuple with the DefaultIoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIoLimitPolicy

`func (o *UpdateStorage) SetDefaultIoLimitPolicy(v string)`

SetDefaultIoLimitPolicy sets DefaultIoLimitPolicy field to given value.

### HasDefaultIoLimitPolicy

`func (o *UpdateStorage) HasDefaultIoLimitPolicy() bool`

HasDefaultIoLimitPolicy returns a boolean if a field has been set.

### GetArrayId

`func (o *UpdateStorage) GetArrayId() string`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *UpdateStorage) GetArrayIdOk() (*string, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *UpdateStorage) SetArrayId(v string)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *UpdateStorage) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetDirectorId

`func (o *UpdateStorage) GetDirectorId() string`

GetDirectorId returns the DirectorId field if non-nil, zero value otherwise.

### GetDirectorIdOk

`func (o *UpdateStorage) GetDirectorIdOk() (*string, bool)`

GetDirectorIdOk returns a tuple with the DirectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorId

`func (o *UpdateStorage) SetDirectorId(v string)`

SetDirectorId sets DirectorId field to given value.

### HasDirectorId

`func (o *UpdateStorage) HasDirectorId() bool`

HasDirectorId returns a boolean if a field has been set.

### GetS3Hostname

`func (o *UpdateStorage) GetS3Hostname() string`

GetS3Hostname returns the S3Hostname field if non-nil, zero value otherwise.

### GetS3HostnameOk

`func (o *UpdateStorage) GetS3HostnameOk() (*string, bool)`

GetS3HostnameOk returns a tuple with the S3Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Hostname

`func (o *UpdateStorage) SetS3Hostname(v string)`

SetS3Hostname sets S3Hostname field to given value.

### HasS3Hostname

`func (o *UpdateStorage) HasS3Hostname() bool`

HasS3Hostname returns a boolean if a field has been set.

### GetS3Port

`func (o *UpdateStorage) GetS3Port() string`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *UpdateStorage) GetS3PortOk() (*string, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *UpdateStorage) SetS3Port(v string)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *UpdateStorage) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetJobInfo

`func (o *UpdateStorage) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *UpdateStorage) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *UpdateStorage) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *UpdateStorage) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### GetLinks

`func (o *UpdateStorage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateStorage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateStorage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateStorage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateStorage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateStorage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateStorage) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateStorage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateStorage) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateStorage) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateStorage) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateStorage) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



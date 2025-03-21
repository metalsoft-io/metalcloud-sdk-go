# Storage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageId** | **float32** | Id of the Storage | 
**UserId** | Pointer to **float32** | Id of the owner | [optional] 
**SiteId** | **float32** | Id of the site | 
**StorageDriver** | **string** | Storage driver | 
**StorageTechnology** | **string** | Storage technology | 
**StorageType** | **string** | Storage type | 
**Name** | **string** | Name of the storage | 
**IscsiHost** | Pointer to **string** | ISCSI host | [optional] 
**IscsiPort** | Pointer to **float32** | ISCSI port | [optional] 
**ManagementHost** | **string** | Management host | 
**Username** | **string** | Username | 
**PasswordEncrypted** | **string** | Password encrypted | 
**Options** | Pointer to [**StorageOptions**](StorageOptions.md) | Options for the storage | [optional] 
**InMaintenance** | **float32** | Specifies if the storage is in maintenance | 
**TargetIQN** | Pointer to **string** | Target IQN | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**AlternateSanIPs** | Pointer to **[]string** | Alternate SAN IPs | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**PortGroupAllocationOrder** | Pointer to **map[string]interface{}** | Port group allocation order | [optional] 
**PortGroupPhysicalPorts** | Pointer to **map[string]interface{}** | Port group physical ports | [optional] 
**DefaultIoLimitPolicy** | Pointer to **string** | Default IO limit policy | [optional] 
**SubnetType** | **string** | Subnet type | 
**ArrayId** | Pointer to **string** | Array id | [optional] 
**DirectorId** | Pointer to **string** | Director id | [optional] 
**S3Hostname** | Pointer to **string** | S3 hostname | [optional] 
**S3Port** | Pointer to **string** | S3 port | [optional] 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewStorage

`func NewStorage(storageId float32, siteId float32, storageDriver string, storageTechnology string, storageType string, name string, managementHost string, username string, passwordEncrypted string, inMaintenance float32, subnetType string, ) *Storage`

NewStorage instantiates a new Storage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageWithDefaults

`func NewStorageWithDefaults() *Storage`

NewStorageWithDefaults instantiates a new Storage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageId

`func (o *Storage) GetStorageId() float32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *Storage) GetStorageIdOk() (*float32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *Storage) SetStorageId(v float32)`

SetStorageId sets StorageId field to given value.


### GetUserId

`func (o *Storage) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Storage) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Storage) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Storage) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *Storage) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Storage) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Storage) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetStorageDriver

`func (o *Storage) GetStorageDriver() string`

GetStorageDriver returns the StorageDriver field if non-nil, zero value otherwise.

### GetStorageDriverOk

`func (o *Storage) GetStorageDriverOk() (*string, bool)`

GetStorageDriverOk returns a tuple with the StorageDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDriver

`func (o *Storage) SetStorageDriver(v string)`

SetStorageDriver sets StorageDriver field to given value.


### GetStorageTechnology

`func (o *Storage) GetStorageTechnology() string`

GetStorageTechnology returns the StorageTechnology field if non-nil, zero value otherwise.

### GetStorageTechnologyOk

`func (o *Storage) GetStorageTechnologyOk() (*string, bool)`

GetStorageTechnologyOk returns a tuple with the StorageTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTechnology

`func (o *Storage) SetStorageTechnology(v string)`

SetStorageTechnology sets StorageTechnology field to given value.


### GetStorageType

`func (o *Storage) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *Storage) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *Storage) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetName

`func (o *Storage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Storage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Storage) SetName(v string)`

SetName sets Name field to given value.


### GetIscsiHost

`func (o *Storage) GetIscsiHost() string`

GetIscsiHost returns the IscsiHost field if non-nil, zero value otherwise.

### GetIscsiHostOk

`func (o *Storage) GetIscsiHostOk() (*string, bool)`

GetIscsiHostOk returns a tuple with the IscsiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiHost

`func (o *Storage) SetIscsiHost(v string)`

SetIscsiHost sets IscsiHost field to given value.

### HasIscsiHost

`func (o *Storage) HasIscsiHost() bool`

HasIscsiHost returns a boolean if a field has been set.

### GetIscsiPort

`func (o *Storage) GetIscsiPort() float32`

GetIscsiPort returns the IscsiPort field if non-nil, zero value otherwise.

### GetIscsiPortOk

`func (o *Storage) GetIscsiPortOk() (*float32, bool)`

GetIscsiPortOk returns a tuple with the IscsiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiPort

`func (o *Storage) SetIscsiPort(v float32)`

SetIscsiPort sets IscsiPort field to given value.

### HasIscsiPort

`func (o *Storage) HasIscsiPort() bool`

HasIscsiPort returns a boolean if a field has been set.

### GetManagementHost

`func (o *Storage) GetManagementHost() string`

GetManagementHost returns the ManagementHost field if non-nil, zero value otherwise.

### GetManagementHostOk

`func (o *Storage) GetManagementHostOk() (*string, bool)`

GetManagementHostOk returns a tuple with the ManagementHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHost

`func (o *Storage) SetManagementHost(v string)`

SetManagementHost sets ManagementHost field to given value.


### GetUsername

`func (o *Storage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Storage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Storage) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPasswordEncrypted

`func (o *Storage) GetPasswordEncrypted() string`

GetPasswordEncrypted returns the PasswordEncrypted field if non-nil, zero value otherwise.

### GetPasswordEncryptedOk

`func (o *Storage) GetPasswordEncryptedOk() (*string, bool)`

GetPasswordEncryptedOk returns a tuple with the PasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncrypted

`func (o *Storage) SetPasswordEncrypted(v string)`

SetPasswordEncrypted sets PasswordEncrypted field to given value.


### GetOptions

`func (o *Storage) GetOptions() StorageOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Storage) GetOptionsOk() (*StorageOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Storage) SetOptions(v StorageOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Storage) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetInMaintenance

`func (o *Storage) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *Storage) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *Storage) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.


### GetTargetIQN

`func (o *Storage) GetTargetIQN() string`

GetTargetIQN returns the TargetIQN field if non-nil, zero value otherwise.

### GetTargetIQNOk

`func (o *Storage) GetTargetIQNOk() (*string, bool)`

GetTargetIQNOk returns a tuple with the TargetIQN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIQN

`func (o *Storage) SetTargetIQN(v string)`

SetTargetIQN sets TargetIQN field to given value.

### HasTargetIQN

`func (o *Storage) HasTargetIQN() bool`

HasTargetIQN returns a boolean if a field has been set.

### GetIsExperimental

`func (o *Storage) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *Storage) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *Storage) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *Storage) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetDrivePriority

`func (o *Storage) GetDrivePriority() float32`

GetDrivePriority returns the DrivePriority field if non-nil, zero value otherwise.

### GetDrivePriorityOk

`func (o *Storage) GetDrivePriorityOk() (*float32, bool)`

GetDrivePriorityOk returns a tuple with the DrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePriority

`func (o *Storage) SetDrivePriority(v float32)`

SetDrivePriority sets DrivePriority field to given value.

### HasDrivePriority

`func (o *Storage) HasDrivePriority() bool`

HasDrivePriority returns a boolean if a field has been set.

### GetSharedDrivePriority

`func (o *Storage) GetSharedDrivePriority() float32`

GetSharedDrivePriority returns the SharedDrivePriority field if non-nil, zero value otherwise.

### GetSharedDrivePriorityOk

`func (o *Storage) GetSharedDrivePriorityOk() (*float32, bool)`

GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrivePriority

`func (o *Storage) SetSharedDrivePriority(v float32)`

SetSharedDrivePriority sets SharedDrivePriority field to given value.

### HasSharedDrivePriority

`func (o *Storage) HasSharedDrivePriority() bool`

HasSharedDrivePriority returns a boolean if a field has been set.

### GetAlternateSanIPs

`func (o *Storage) GetAlternateSanIPs() []string`

GetAlternateSanIPs returns the AlternateSanIPs field if non-nil, zero value otherwise.

### GetAlternateSanIPsOk

`func (o *Storage) GetAlternateSanIPsOk() (*[]string, bool)`

GetAlternateSanIPsOk returns a tuple with the AlternateSanIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateSanIPs

`func (o *Storage) SetAlternateSanIPs(v []string)`

SetAlternateSanIPs sets AlternateSanIPs field to given value.

### HasAlternateSanIPs

`func (o *Storage) HasAlternateSanIPs() bool`

HasAlternateSanIPs returns a boolean if a field has been set.

### GetTags

`func (o *Storage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Storage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Storage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Storage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPortGroupAllocationOrder

`func (o *Storage) GetPortGroupAllocationOrder() map[string]interface{}`

GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field if non-nil, zero value otherwise.

### GetPortGroupAllocationOrderOk

`func (o *Storage) GetPortGroupAllocationOrderOk() (*map[string]interface{}, bool)`

GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupAllocationOrder

`func (o *Storage) SetPortGroupAllocationOrder(v map[string]interface{})`

SetPortGroupAllocationOrder sets PortGroupAllocationOrder field to given value.

### HasPortGroupAllocationOrder

`func (o *Storage) HasPortGroupAllocationOrder() bool`

HasPortGroupAllocationOrder returns a boolean if a field has been set.

### GetPortGroupPhysicalPorts

`func (o *Storage) GetPortGroupPhysicalPorts() map[string]interface{}`

GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field if non-nil, zero value otherwise.

### GetPortGroupPhysicalPortsOk

`func (o *Storage) GetPortGroupPhysicalPortsOk() (*map[string]interface{}, bool)`

GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortGroupPhysicalPorts

`func (o *Storage) SetPortGroupPhysicalPorts(v map[string]interface{})`

SetPortGroupPhysicalPorts sets PortGroupPhysicalPorts field to given value.

### HasPortGroupPhysicalPorts

`func (o *Storage) HasPortGroupPhysicalPorts() bool`

HasPortGroupPhysicalPorts returns a boolean if a field has been set.

### GetDefaultIoLimitPolicy

`func (o *Storage) GetDefaultIoLimitPolicy() string`

GetDefaultIoLimitPolicy returns the DefaultIoLimitPolicy field if non-nil, zero value otherwise.

### GetDefaultIoLimitPolicyOk

`func (o *Storage) GetDefaultIoLimitPolicyOk() (*string, bool)`

GetDefaultIoLimitPolicyOk returns a tuple with the DefaultIoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIoLimitPolicy

`func (o *Storage) SetDefaultIoLimitPolicy(v string)`

SetDefaultIoLimitPolicy sets DefaultIoLimitPolicy field to given value.

### HasDefaultIoLimitPolicy

`func (o *Storage) HasDefaultIoLimitPolicy() bool`

HasDefaultIoLimitPolicy returns a boolean if a field has been set.

### GetSubnetType

`func (o *Storage) GetSubnetType() string`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *Storage) GetSubnetTypeOk() (*string, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *Storage) SetSubnetType(v string)`

SetSubnetType sets SubnetType field to given value.


### GetArrayId

`func (o *Storage) GetArrayId() string`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *Storage) GetArrayIdOk() (*string, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *Storage) SetArrayId(v string)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *Storage) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetDirectorId

`func (o *Storage) GetDirectorId() string`

GetDirectorId returns the DirectorId field if non-nil, zero value otherwise.

### GetDirectorIdOk

`func (o *Storage) GetDirectorIdOk() (*string, bool)`

GetDirectorIdOk returns a tuple with the DirectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorId

`func (o *Storage) SetDirectorId(v string)`

SetDirectorId sets DirectorId field to given value.

### HasDirectorId

`func (o *Storage) HasDirectorId() bool`

HasDirectorId returns a boolean if a field has been set.

### GetS3Hostname

`func (o *Storage) GetS3Hostname() string`

GetS3Hostname returns the S3Hostname field if non-nil, zero value otherwise.

### GetS3HostnameOk

`func (o *Storage) GetS3HostnameOk() (*string, bool)`

GetS3HostnameOk returns a tuple with the S3Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Hostname

`func (o *Storage) SetS3Hostname(v string)`

SetS3Hostname sets S3Hostname field to given value.

### HasS3Hostname

`func (o *Storage) HasS3Hostname() bool`

HasS3Hostname returns a boolean if a field has been set.

### GetS3Port

`func (o *Storage) GetS3Port() string`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *Storage) GetS3PortOk() (*string, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *Storage) SetS3Port(v string)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *Storage) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetJobInfo

`func (o *Storage) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *Storage) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *Storage) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *Storage) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### GetLinks

`func (o *Storage) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Storage) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Storage) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Storage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



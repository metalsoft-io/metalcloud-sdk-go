# DriveConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision of the Drive Configuration | 
**Label** | **string** | Label of the Drive. | 
**GroupId** | **float32** | Drive Array Id | 
**ContainerId** | Pointer to **float32** |  | [optional] 
**InstanceId** | Pointer to **float32** |  | [optional] 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the Drive is assigned to | [optional] 
**SizeMb** | **float32** | Disk size in MB for Drive | 
**StorageImageName** | Pointer to **string** | The name of the storage image used by the Drive. | [optional] 
**IscsiIndexHex** | Pointer to **string** | The iSCSI Index in hex format of the Drive. | [optional] 
**TemplateIdOrigin** | Pointer to **float32** | Template Id | [optional] 
**OsAdminUsername** | Pointer to **string** | The OS Admin Username the Drive will use. | [optional] 
**OsAdminPasswordEncrypted** | Pointer to **string** | The OS Admin Password the Drive will use. | [optional] 
**StorageType** | **string** | Service status of the Drive | [default to "iscsi_ssd"]
**Subdomain** | Pointer to **string** | Subdomain of the Drive. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Drive last update. | 
**SshPort** | Pointer to **float32** | SSH port used by the Drive. | [optional] 
**OperatingSystemInfo** | Pointer to **map[string]interface{}** | Operating system information of the Drive. | [optional] 
**FilesystemInfo** | Pointer to **map[string]interface{}** | Filesystem information of the Drive. | [optional] 
**DnsSubdomainChangeId** | Pointer to **float32** | Id of the DNS subdomain for the Drive. | [optional] 
**DeployType** | **string** | Deploy type of the Drive | [default to "create"]
**DeployStatus** | **string** | Deploy status of the Drive | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the Drive. | [optional] 

## Methods

### NewDriveConfiguration

`func NewDriveConfiguration(revision float32, label string, groupId float32, sizeMb float32, storageType string, updatedTimestamp string, deployType string, deployStatus string, ) *DriveConfiguration`

NewDriveConfiguration instantiates a new DriveConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveConfigurationWithDefaults

`func NewDriveConfigurationWithDefaults() *DriveConfiguration`

NewDriveConfigurationWithDefaults instantiates a new DriveConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *DriveConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DriveConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DriveConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *DriveConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DriveConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DriveConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGroupId

`func (o *DriveConfiguration) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DriveConfiguration) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DriveConfiguration) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetContainerId

`func (o *DriveConfiguration) GetContainerId() float32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *DriveConfiguration) GetContainerIdOk() (*float32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *DriveConfiguration) SetContainerId(v float32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *DriveConfiguration) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetInstanceId

`func (o *DriveConfiguration) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *DriveConfiguration) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *DriveConfiguration) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *DriveConfiguration) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStoragePoolId

`func (o *DriveConfiguration) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *DriveConfiguration) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *DriveConfiguration) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *DriveConfiguration) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.

### GetSizeMb

`func (o *DriveConfiguration) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *DriveConfiguration) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *DriveConfiguration) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetStorageImageName

`func (o *DriveConfiguration) GetStorageImageName() string`

GetStorageImageName returns the StorageImageName field if non-nil, zero value otherwise.

### GetStorageImageNameOk

`func (o *DriveConfiguration) GetStorageImageNameOk() (*string, bool)`

GetStorageImageNameOk returns a tuple with the StorageImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageImageName

`func (o *DriveConfiguration) SetStorageImageName(v string)`

SetStorageImageName sets StorageImageName field to given value.

### HasStorageImageName

`func (o *DriveConfiguration) HasStorageImageName() bool`

HasStorageImageName returns a boolean if a field has been set.

### GetIscsiIndexHex

`func (o *DriveConfiguration) GetIscsiIndexHex() string`

GetIscsiIndexHex returns the IscsiIndexHex field if non-nil, zero value otherwise.

### GetIscsiIndexHexOk

`func (o *DriveConfiguration) GetIscsiIndexHexOk() (*string, bool)`

GetIscsiIndexHexOk returns a tuple with the IscsiIndexHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiIndexHex

`func (o *DriveConfiguration) SetIscsiIndexHex(v string)`

SetIscsiIndexHex sets IscsiIndexHex field to given value.

### HasIscsiIndexHex

`func (o *DriveConfiguration) HasIscsiIndexHex() bool`

HasIscsiIndexHex returns a boolean if a field has been set.

### GetTemplateIdOrigin

`func (o *DriveConfiguration) GetTemplateIdOrigin() float32`

GetTemplateIdOrigin returns the TemplateIdOrigin field if non-nil, zero value otherwise.

### GetTemplateIdOriginOk

`func (o *DriveConfiguration) GetTemplateIdOriginOk() (*float32, bool)`

GetTemplateIdOriginOk returns a tuple with the TemplateIdOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIdOrigin

`func (o *DriveConfiguration) SetTemplateIdOrigin(v float32)`

SetTemplateIdOrigin sets TemplateIdOrigin field to given value.

### HasTemplateIdOrigin

`func (o *DriveConfiguration) HasTemplateIdOrigin() bool`

HasTemplateIdOrigin returns a boolean if a field has been set.

### GetOsAdminUsername

`func (o *DriveConfiguration) GetOsAdminUsername() string`

GetOsAdminUsername returns the OsAdminUsername field if non-nil, zero value otherwise.

### GetOsAdminUsernameOk

`func (o *DriveConfiguration) GetOsAdminUsernameOk() (*string, bool)`

GetOsAdminUsernameOk returns a tuple with the OsAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdminUsername

`func (o *DriveConfiguration) SetOsAdminUsername(v string)`

SetOsAdminUsername sets OsAdminUsername field to given value.

### HasOsAdminUsername

`func (o *DriveConfiguration) HasOsAdminUsername() bool`

HasOsAdminUsername returns a boolean if a field has been set.

### GetOsAdminPasswordEncrypted

`func (o *DriveConfiguration) GetOsAdminPasswordEncrypted() string`

GetOsAdminPasswordEncrypted returns the OsAdminPasswordEncrypted field if non-nil, zero value otherwise.

### GetOsAdminPasswordEncryptedOk

`func (o *DriveConfiguration) GetOsAdminPasswordEncryptedOk() (*string, bool)`

GetOsAdminPasswordEncryptedOk returns a tuple with the OsAdminPasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsAdminPasswordEncrypted

`func (o *DriveConfiguration) SetOsAdminPasswordEncrypted(v string)`

SetOsAdminPasswordEncrypted sets OsAdminPasswordEncrypted field to given value.

### HasOsAdminPasswordEncrypted

`func (o *DriveConfiguration) HasOsAdminPasswordEncrypted() bool`

HasOsAdminPasswordEncrypted returns a boolean if a field has been set.

### GetStorageType

`func (o *DriveConfiguration) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DriveConfiguration) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DriveConfiguration) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetSubdomain

`func (o *DriveConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DriveConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DriveConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DriveConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DriveConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DriveConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DriveConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSshPort

`func (o *DriveConfiguration) GetSshPort() float32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *DriveConfiguration) GetSshPortOk() (*float32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *DriveConfiguration) SetSshPort(v float32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *DriveConfiguration) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetOperatingSystemInfo

`func (o *DriveConfiguration) GetOperatingSystemInfo() map[string]interface{}`

GetOperatingSystemInfo returns the OperatingSystemInfo field if non-nil, zero value otherwise.

### GetOperatingSystemInfoOk

`func (o *DriveConfiguration) GetOperatingSystemInfoOk() (*map[string]interface{}, bool)`

GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemInfo

`func (o *DriveConfiguration) SetOperatingSystemInfo(v map[string]interface{})`

SetOperatingSystemInfo sets OperatingSystemInfo field to given value.

### HasOperatingSystemInfo

`func (o *DriveConfiguration) HasOperatingSystemInfo() bool`

HasOperatingSystemInfo returns a boolean if a field has been set.

### GetFilesystemInfo

`func (o *DriveConfiguration) GetFilesystemInfo() map[string]interface{}`

GetFilesystemInfo returns the FilesystemInfo field if non-nil, zero value otherwise.

### GetFilesystemInfoOk

`func (o *DriveConfiguration) GetFilesystemInfoOk() (*map[string]interface{}, bool)`

GetFilesystemInfoOk returns a tuple with the FilesystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemInfo

`func (o *DriveConfiguration) SetFilesystemInfo(v map[string]interface{})`

SetFilesystemInfo sets FilesystemInfo field to given value.

### HasFilesystemInfo

`func (o *DriveConfiguration) HasFilesystemInfo() bool`

HasFilesystemInfo returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *DriveConfiguration) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *DriveConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *DriveConfiguration) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *DriveConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDeployType

`func (o *DriveConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DriveConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DriveConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *DriveConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *DriveConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *DriveConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *DriveConfiguration) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *DriveConfiguration) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *DriveConfiguration) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *DriveConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



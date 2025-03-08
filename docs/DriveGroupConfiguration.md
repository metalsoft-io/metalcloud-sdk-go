# DriveGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision of the Drive Group Configuration | 
**Label** | **string** | Label of the Drive Group. | 
**InfrastructureId** | **float32** | Infrastructure id of the Drive Group | 
**TemplateId** | Pointer to **float32** | Template Id | [optional] 
**DriveCount** | **float32** | Number of drives in the Drive Group | 
**DriveSizeMbDefault** | **float32** | Default disk size in MB for new Drives in the Drive Group | 
**ServerInstanceGroupId** | Pointer to **float32** |  | [optional] 
**ContainerArrayId** | Pointer to **float32** |  | [optional] 
**ExpandWithServerInstanceGroup** | **float32** | Flag to determine whether the Drive Group should be expanded with a Server Instance Group by adding one drive for each instance | 
**IoLimitPolicy** | Pointer to **string** | The IO limit policy of the Drive Group. | [optional] 
**StorageType** | **string** | Service status of the Drive Group | [default to "iscsi_ssd"]
**FilesystemInfo** | Pointer to **map[string]interface{}** | Filesystem information of the Drive Group. | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the Drive Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Drive Group last update. | 
**DnsSubdomainChangeId** | Pointer to **float32** | Id of the DNS subdomain for the Drive Group. | [optional] 
**DeployType** | **string** | Deploy type of the Drive Group | [default to "create"]
**DeployStatus** | **string** | Deploy status of the Drive Group | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the Drive Group. | [optional] 

## Methods

### NewDriveGroupConfiguration

`func NewDriveGroupConfiguration(revision float32, label string, infrastructureId float32, driveCount float32, driveSizeMbDefault float32, expandWithServerInstanceGroup float32, storageType string, updatedTimestamp string, deployType string, deployStatus string, ) *DriveGroupConfiguration`

NewDriveGroupConfiguration instantiates a new DriveGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveGroupConfigurationWithDefaults

`func NewDriveGroupConfigurationWithDefaults() *DriveGroupConfiguration`

NewDriveGroupConfigurationWithDefaults instantiates a new DriveGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *DriveGroupConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DriveGroupConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DriveGroupConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *DriveGroupConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DriveGroupConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DriveGroupConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInfrastructureId

`func (o *DriveGroupConfiguration) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *DriveGroupConfiguration) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *DriveGroupConfiguration) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetTemplateId

`func (o *DriveGroupConfiguration) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DriveGroupConfiguration) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DriveGroupConfiguration) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DriveGroupConfiguration) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDriveCount

`func (o *DriveGroupConfiguration) GetDriveCount() float32`

GetDriveCount returns the DriveCount field if non-nil, zero value otherwise.

### GetDriveCountOk

`func (o *DriveGroupConfiguration) GetDriveCountOk() (*float32, bool)`

GetDriveCountOk returns a tuple with the DriveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveCount

`func (o *DriveGroupConfiguration) SetDriveCount(v float32)`

SetDriveCount sets DriveCount field to given value.


### GetDriveSizeMbDefault

`func (o *DriveGroupConfiguration) GetDriveSizeMbDefault() float32`

GetDriveSizeMbDefault returns the DriveSizeMbDefault field if non-nil, zero value otherwise.

### GetDriveSizeMbDefaultOk

`func (o *DriveGroupConfiguration) GetDriveSizeMbDefaultOk() (*float32, bool)`

GetDriveSizeMbDefaultOk returns a tuple with the DriveSizeMbDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSizeMbDefault

`func (o *DriveGroupConfiguration) SetDriveSizeMbDefault(v float32)`

SetDriveSizeMbDefault sets DriveSizeMbDefault field to given value.


### GetServerInstanceGroupId

`func (o *DriveGroupConfiguration) GetServerInstanceGroupId() float32`

GetServerInstanceGroupId returns the ServerInstanceGroupId field if non-nil, zero value otherwise.

### GetServerInstanceGroupIdOk

`func (o *DriveGroupConfiguration) GetServerInstanceGroupIdOk() (*float32, bool)`

GetServerInstanceGroupIdOk returns a tuple with the ServerInstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroupId

`func (o *DriveGroupConfiguration) SetServerInstanceGroupId(v float32)`

SetServerInstanceGroupId sets ServerInstanceGroupId field to given value.

### HasServerInstanceGroupId

`func (o *DriveGroupConfiguration) HasServerInstanceGroupId() bool`

HasServerInstanceGroupId returns a boolean if a field has been set.

### GetContainerArrayId

`func (o *DriveGroupConfiguration) GetContainerArrayId() float32`

GetContainerArrayId returns the ContainerArrayId field if non-nil, zero value otherwise.

### GetContainerArrayIdOk

`func (o *DriveGroupConfiguration) GetContainerArrayIdOk() (*float32, bool)`

GetContainerArrayIdOk returns a tuple with the ContainerArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerArrayId

`func (o *DriveGroupConfiguration) SetContainerArrayId(v float32)`

SetContainerArrayId sets ContainerArrayId field to given value.

### HasContainerArrayId

`func (o *DriveGroupConfiguration) HasContainerArrayId() bool`

HasContainerArrayId returns a boolean if a field has been set.

### GetExpandWithServerInstanceGroup

`func (o *DriveGroupConfiguration) GetExpandWithServerInstanceGroup() float32`

GetExpandWithServerInstanceGroup returns the ExpandWithServerInstanceGroup field if non-nil, zero value otherwise.

### GetExpandWithServerInstanceGroupOk

`func (o *DriveGroupConfiguration) GetExpandWithServerInstanceGroupOk() (*float32, bool)`

GetExpandWithServerInstanceGroupOk returns a tuple with the ExpandWithServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandWithServerInstanceGroup

`func (o *DriveGroupConfiguration) SetExpandWithServerInstanceGroup(v float32)`

SetExpandWithServerInstanceGroup sets ExpandWithServerInstanceGroup field to given value.


### GetIoLimitPolicy

`func (o *DriveGroupConfiguration) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *DriveGroupConfiguration) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *DriveGroupConfiguration) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *DriveGroupConfiguration) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetStorageType

`func (o *DriveGroupConfiguration) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DriveGroupConfiguration) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DriveGroupConfiguration) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetFilesystemInfo

`func (o *DriveGroupConfiguration) GetFilesystemInfo() map[string]interface{}`

GetFilesystemInfo returns the FilesystemInfo field if non-nil, zero value otherwise.

### GetFilesystemInfoOk

`func (o *DriveGroupConfiguration) GetFilesystemInfoOk() (*map[string]interface{}, bool)`

GetFilesystemInfoOk returns a tuple with the FilesystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemInfo

`func (o *DriveGroupConfiguration) SetFilesystemInfo(v map[string]interface{})`

SetFilesystemInfo sets FilesystemInfo field to given value.

### HasFilesystemInfo

`func (o *DriveGroupConfiguration) HasFilesystemInfo() bool`

HasFilesystemInfo returns a boolean if a field has been set.

### GetSubdomain

`func (o *DriveGroupConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DriveGroupConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DriveGroupConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DriveGroupConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DriveGroupConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DriveGroupConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DriveGroupConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetDnsSubdomainChangeId

`func (o *DriveGroupConfiguration) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *DriveGroupConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *DriveGroupConfiguration) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *DriveGroupConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDeployType

`func (o *DriveGroupConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DriveGroupConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DriveGroupConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *DriveGroupConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *DriveGroupConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *DriveGroupConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *DriveGroupConfiguration) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *DriveGroupConfiguration) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *DriveGroupConfiguration) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *DriveGroupConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



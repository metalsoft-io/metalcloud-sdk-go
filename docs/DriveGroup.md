# DriveGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the Drive Group. | 
**InfrastructureId** | **float32** | Infrastructure id of the Drive Group | 
**TemplateId** | Pointer to **float32** | Template Id | [optional] 
**DriveSizeMbDefault** | **float32** | Default disk size in MB for new Drives in the Drive Group | 
**ServerInstanceGroupId** | Pointer to **float32** |  | [optional] 
**ExpandWithServerInstanceGroup** | **float32** | Flag to determine whether the Drive Group should be expanded with a Server Instance Group by adding one drive for each instance | 
**IoLimitPolicy** | Pointer to **string** | The IO limit policy of the Drive Group. | [optional] 
**StorageType** | **string** | Service status of the Drive Group | [default to "iscsi_ssd"]
**FilesystemInfo** | Pointer to **map[string]interface{}** | Filesystem information of the Drive Group. | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the Drive Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Drive Group last update. | 
**Id** | **float32** | Id of the Drive Group | 
**Revision** | **float32** | Revision of the Drive Group State | 
**ExtensionInstanceId** | Pointer to **float32** |  | [optional] 
**ContainerClusterId** | Pointer to **float32** |  | [optional] 
**ServiceStatus** | **string** | Service status of the Drive Group | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Drive Group. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the Drive Group. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the Drive Group. | [optional] 
**AllocationAffinity** | **string** | Allocation affinity of the Drive Group | 
**Config** | [**DriveGroupConfiguration**](DriveGroupConfiguration.md) | The current changes to be deployed for the Drive Group. | 
**CreatedTimestamp** | **string** | Timestamp of the Drive Group creation. | 
**Meta** | [**DriveGroupMeta**](DriveGroupMeta.md) | Meta information of the Drive Group. | 

## Methods

### NewDriveGroup

`func NewDriveGroup(label string, infrastructureId float32, driveSizeMbDefault float32, expandWithServerInstanceGroup float32, storageType string, updatedTimestamp string, id float32, revision float32, serviceStatus string, allocationAffinity string, config DriveGroupConfiguration, createdTimestamp string, meta DriveGroupMeta, ) *DriveGroup`

NewDriveGroup instantiates a new DriveGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveGroupWithDefaults

`func NewDriveGroupWithDefaults() *DriveGroup`

NewDriveGroupWithDefaults instantiates a new DriveGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *DriveGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DriveGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DriveGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInfrastructureId

`func (o *DriveGroup) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *DriveGroup) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *DriveGroup) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetTemplateId

`func (o *DriveGroup) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DriveGroup) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DriveGroup) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DriveGroup) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDriveSizeMbDefault

`func (o *DriveGroup) GetDriveSizeMbDefault() float32`

GetDriveSizeMbDefault returns the DriveSizeMbDefault field if non-nil, zero value otherwise.

### GetDriveSizeMbDefaultOk

`func (o *DriveGroup) GetDriveSizeMbDefaultOk() (*float32, bool)`

GetDriveSizeMbDefaultOk returns a tuple with the DriveSizeMbDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSizeMbDefault

`func (o *DriveGroup) SetDriveSizeMbDefault(v float32)`

SetDriveSizeMbDefault sets DriveSizeMbDefault field to given value.


### GetServerInstanceGroupId

`func (o *DriveGroup) GetServerInstanceGroupId() float32`

GetServerInstanceGroupId returns the ServerInstanceGroupId field if non-nil, zero value otherwise.

### GetServerInstanceGroupIdOk

`func (o *DriveGroup) GetServerInstanceGroupIdOk() (*float32, bool)`

GetServerInstanceGroupIdOk returns a tuple with the ServerInstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroupId

`func (o *DriveGroup) SetServerInstanceGroupId(v float32)`

SetServerInstanceGroupId sets ServerInstanceGroupId field to given value.

### HasServerInstanceGroupId

`func (o *DriveGroup) HasServerInstanceGroupId() bool`

HasServerInstanceGroupId returns a boolean if a field has been set.

### GetExpandWithServerInstanceGroup

`func (o *DriveGroup) GetExpandWithServerInstanceGroup() float32`

GetExpandWithServerInstanceGroup returns the ExpandWithServerInstanceGroup field if non-nil, zero value otherwise.

### GetExpandWithServerInstanceGroupOk

`func (o *DriveGroup) GetExpandWithServerInstanceGroupOk() (*float32, bool)`

GetExpandWithServerInstanceGroupOk returns a tuple with the ExpandWithServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandWithServerInstanceGroup

`func (o *DriveGroup) SetExpandWithServerInstanceGroup(v float32)`

SetExpandWithServerInstanceGroup sets ExpandWithServerInstanceGroup field to given value.


### GetIoLimitPolicy

`func (o *DriveGroup) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *DriveGroup) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *DriveGroup) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *DriveGroup) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetStorageType

`func (o *DriveGroup) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DriveGroup) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DriveGroup) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetFilesystemInfo

`func (o *DriveGroup) GetFilesystemInfo() map[string]interface{}`

GetFilesystemInfo returns the FilesystemInfo field if non-nil, zero value otherwise.

### GetFilesystemInfoOk

`func (o *DriveGroup) GetFilesystemInfoOk() (*map[string]interface{}, bool)`

GetFilesystemInfoOk returns a tuple with the FilesystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemInfo

`func (o *DriveGroup) SetFilesystemInfo(v map[string]interface{})`

SetFilesystemInfo sets FilesystemInfo field to given value.

### HasFilesystemInfo

`func (o *DriveGroup) HasFilesystemInfo() bool`

HasFilesystemInfo returns a boolean if a field has been set.

### GetSubdomain

`func (o *DriveGroup) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DriveGroup) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DriveGroup) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DriveGroup) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DriveGroup) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DriveGroup) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DriveGroup) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *DriveGroup) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveGroup) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveGroup) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *DriveGroup) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DriveGroup) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DriveGroup) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetExtensionInstanceId

`func (o *DriveGroup) GetExtensionInstanceId() float32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *DriveGroup) GetExtensionInstanceIdOk() (*float32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *DriveGroup) SetExtensionInstanceId(v float32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *DriveGroup) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetContainerClusterId

`func (o *DriveGroup) GetContainerClusterId() float32`

GetContainerClusterId returns the ContainerClusterId field if non-nil, zero value otherwise.

### GetContainerClusterIdOk

`func (o *DriveGroup) GetContainerClusterIdOk() (*float32, bool)`

GetContainerClusterIdOk returns a tuple with the ContainerClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerClusterId

`func (o *DriveGroup) SetContainerClusterId(v float32)`

SetContainerClusterId sets ContainerClusterId field to given value.

### HasContainerClusterId

`func (o *DriveGroup) HasContainerClusterId() bool`

HasContainerClusterId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *DriveGroup) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *DriveGroup) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *DriveGroup) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *DriveGroup) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *DriveGroup) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *DriveGroup) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *DriveGroup) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *DriveGroup) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *DriveGroup) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *DriveGroup) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *DriveGroup) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *DriveGroup) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *DriveGroup) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *DriveGroup) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *DriveGroup) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetAllocationAffinity

`func (o *DriveGroup) GetAllocationAffinity() string`

GetAllocationAffinity returns the AllocationAffinity field if non-nil, zero value otherwise.

### GetAllocationAffinityOk

`func (o *DriveGroup) GetAllocationAffinityOk() (*string, bool)`

GetAllocationAffinityOk returns a tuple with the AllocationAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAffinity

`func (o *DriveGroup) SetAllocationAffinity(v string)`

SetAllocationAffinity sets AllocationAffinity field to given value.


### GetConfig

`func (o *DriveGroup) GetConfig() DriveGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DriveGroup) GetConfigOk() (*DriveGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DriveGroup) SetConfig(v DriveGroupConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *DriveGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DriveGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DriveGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetMeta

`func (o *DriveGroup) GetMeta() DriveGroupMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DriveGroup) GetMetaOk() (*DriveGroupMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DriveGroup) SetMeta(v DriveGroupMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



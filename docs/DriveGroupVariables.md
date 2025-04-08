# DriveGroupVariables

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

## Methods

### NewDriveGroupVariables

`func NewDriveGroupVariables(label string, infrastructureId float32, driveSizeMbDefault float32, expandWithServerInstanceGroup float32, storageType string, updatedTimestamp string, id float32, revision float32, serviceStatus string, allocationAffinity string, config DriveGroupConfiguration, createdTimestamp string, ) *DriveGroupVariables`

NewDriveGroupVariables instantiates a new DriveGroupVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveGroupVariablesWithDefaults

`func NewDriveGroupVariablesWithDefaults() *DriveGroupVariables`

NewDriveGroupVariablesWithDefaults instantiates a new DriveGroupVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *DriveGroupVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DriveGroupVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DriveGroupVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInfrastructureId

`func (o *DriveGroupVariables) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *DriveGroupVariables) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *DriveGroupVariables) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetTemplateId

`func (o *DriveGroupVariables) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DriveGroupVariables) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DriveGroupVariables) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DriveGroupVariables) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDriveSizeMbDefault

`func (o *DriveGroupVariables) GetDriveSizeMbDefault() float32`

GetDriveSizeMbDefault returns the DriveSizeMbDefault field if non-nil, zero value otherwise.

### GetDriveSizeMbDefaultOk

`func (o *DriveGroupVariables) GetDriveSizeMbDefaultOk() (*float32, bool)`

GetDriveSizeMbDefaultOk returns a tuple with the DriveSizeMbDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSizeMbDefault

`func (o *DriveGroupVariables) SetDriveSizeMbDefault(v float32)`

SetDriveSizeMbDefault sets DriveSizeMbDefault field to given value.


### GetServerInstanceGroupId

`func (o *DriveGroupVariables) GetServerInstanceGroupId() float32`

GetServerInstanceGroupId returns the ServerInstanceGroupId field if non-nil, zero value otherwise.

### GetServerInstanceGroupIdOk

`func (o *DriveGroupVariables) GetServerInstanceGroupIdOk() (*float32, bool)`

GetServerInstanceGroupIdOk returns a tuple with the ServerInstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroupId

`func (o *DriveGroupVariables) SetServerInstanceGroupId(v float32)`

SetServerInstanceGroupId sets ServerInstanceGroupId field to given value.

### HasServerInstanceGroupId

`func (o *DriveGroupVariables) HasServerInstanceGroupId() bool`

HasServerInstanceGroupId returns a boolean if a field has been set.

### GetExpandWithServerInstanceGroup

`func (o *DriveGroupVariables) GetExpandWithServerInstanceGroup() float32`

GetExpandWithServerInstanceGroup returns the ExpandWithServerInstanceGroup field if non-nil, zero value otherwise.

### GetExpandWithServerInstanceGroupOk

`func (o *DriveGroupVariables) GetExpandWithServerInstanceGroupOk() (*float32, bool)`

GetExpandWithServerInstanceGroupOk returns a tuple with the ExpandWithServerInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandWithServerInstanceGroup

`func (o *DriveGroupVariables) SetExpandWithServerInstanceGroup(v float32)`

SetExpandWithServerInstanceGroup sets ExpandWithServerInstanceGroup field to given value.


### GetIoLimitPolicy

`func (o *DriveGroupVariables) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *DriveGroupVariables) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *DriveGroupVariables) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *DriveGroupVariables) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetStorageType

`func (o *DriveGroupVariables) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DriveGroupVariables) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DriveGroupVariables) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetFilesystemInfo

`func (o *DriveGroupVariables) GetFilesystemInfo() map[string]interface{}`

GetFilesystemInfo returns the FilesystemInfo field if non-nil, zero value otherwise.

### GetFilesystemInfoOk

`func (o *DriveGroupVariables) GetFilesystemInfoOk() (*map[string]interface{}, bool)`

GetFilesystemInfoOk returns a tuple with the FilesystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemInfo

`func (o *DriveGroupVariables) SetFilesystemInfo(v map[string]interface{})`

SetFilesystemInfo sets FilesystemInfo field to given value.

### HasFilesystemInfo

`func (o *DriveGroupVariables) HasFilesystemInfo() bool`

HasFilesystemInfo returns a boolean if a field has been set.

### GetSubdomain

`func (o *DriveGroupVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DriveGroupVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DriveGroupVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DriveGroupVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *DriveGroupVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DriveGroupVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DriveGroupVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *DriveGroupVariables) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveGroupVariables) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveGroupVariables) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *DriveGroupVariables) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DriveGroupVariables) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DriveGroupVariables) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetExtensionInstanceId

`func (o *DriveGroupVariables) GetExtensionInstanceId() float32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *DriveGroupVariables) GetExtensionInstanceIdOk() (*float32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *DriveGroupVariables) SetExtensionInstanceId(v float32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *DriveGroupVariables) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetContainerClusterId

`func (o *DriveGroupVariables) GetContainerClusterId() float32`

GetContainerClusterId returns the ContainerClusterId field if non-nil, zero value otherwise.

### GetContainerClusterIdOk

`func (o *DriveGroupVariables) GetContainerClusterIdOk() (*float32, bool)`

GetContainerClusterIdOk returns a tuple with the ContainerClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerClusterId

`func (o *DriveGroupVariables) SetContainerClusterId(v float32)`

SetContainerClusterId sets ContainerClusterId field to given value.

### HasContainerClusterId

`func (o *DriveGroupVariables) HasContainerClusterId() bool`

HasContainerClusterId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *DriveGroupVariables) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *DriveGroupVariables) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *DriveGroupVariables) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *DriveGroupVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *DriveGroupVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *DriveGroupVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *DriveGroupVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *DriveGroupVariables) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *DriveGroupVariables) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *DriveGroupVariables) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *DriveGroupVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *DriveGroupVariables) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *DriveGroupVariables) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *DriveGroupVariables) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *DriveGroupVariables) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetAllocationAffinity

`func (o *DriveGroupVariables) GetAllocationAffinity() string`

GetAllocationAffinity returns the AllocationAffinity field if non-nil, zero value otherwise.

### GetAllocationAffinityOk

`func (o *DriveGroupVariables) GetAllocationAffinityOk() (*string, bool)`

GetAllocationAffinityOk returns a tuple with the AllocationAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationAffinity

`func (o *DriveGroupVariables) SetAllocationAffinity(v string)`

SetAllocationAffinity sets AllocationAffinity field to given value.


### GetConfig

`func (o *DriveGroupVariables) GetConfig() DriveGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DriveGroupVariables) GetConfigOk() (*DriveGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DriveGroupVariables) SetConfig(v DriveGroupConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *DriveGroupVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DriveGroupVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DriveGroupVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



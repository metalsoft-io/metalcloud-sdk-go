# ContainerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Name of the Container Instance. | 
**ContainerId** | Pointer to **int64** | Id of the Container. | [optional] 
**TypeId** | **int64** | Id of the Container Type. | 
**Subdomain** | Pointer to **string** | Subdomain of the Container Instance. | [optional] 
**VmPoolId** | Pointer to **int64** | Id of the VM Pool. | [optional] 
**DiskSizeGB** | **float32** | Disk size in GB of the Container Instance. | 
**RamGB** | **float32** | RAM size in GB of the Container Instance. | 
**CpuCores** | **float32** | Number of CPU cores for the Container Instance. | 
**OsTemplateId** | Pointer to **int64** | Id of the template used by the Container Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the Container Instance. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Container Instance last update. | 
**Id** | **int64** | Container Instance ID | 
**Revision** | **int64** | Revision of the Container Instance | 
**GroupId** | **int64** | Id of the Container Instance Group. | 
**InfrastructureId** | **int64** | Id of the Infrastructure. | 
**Infrastructure** | [**ParentInfrastructure**](ParentInfrastructure.md) | Infrastructure information | 
**ServiceStatus** | **string** | Service status of the Container Instance. | 
**InstanceType** | **string** | Type of the Instance. | 
**GpuInfo** | Pointer to [**[]VMTypeGPUInfo**](VMTypeGPUInfo.md) | Information about GPUs available for this Container Instance | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Container Instance. | [optional] 
**DnsSubdomainId** | Pointer to **int64** | Id of the DNS subdomain for the Container Instance. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int64** | Id of the permanent DNS subdomain for the Container Instance. | [optional] 
**UniqueIndex** | Pointer to **float32** | Unique index of the Container Instance. | [optional] 
**CreatedTimestamp** | **string** | Timestamp of the Container Instance creation. | 
**Config** | [**ContainerInstanceConfiguration**](ContainerInstanceConfiguration.md) | The current changes to be deployed for the Container Instance. | 
**Meta** | [**ContainerInstanceMeta**](ContainerInstanceMeta.md) | Meta information of the Container Instance. | 
**AllocatedVMPoolInfo** | Pointer to [**ContainerInstanceAllocatedVMPoolInfo**](ContainerInstanceAllocatedVMPoolInfo.md) | Information about the allocated VM Pool. | [optional] 
**AllocatedVMPoolGpusInfo** | Pointer to [**[]VMPoolGPU**](VMPoolGPU.md) | Information about the allocated GPUs from the VM Pool. | [optional] 
**AllocatedVMPoolGpuPartitionsInfo** | Pointer to [**[]VMPoolGPUPartition**](VMPoolGPUPartition.md) | Information about the allocated GPU partitions from the VM Pool. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewContainerInstance

`func NewContainerInstance(label string, typeId int64, diskSizeGB float32, ramGB float32, cpuCores float32, updatedTimestamp string, id int64, revision int64, groupId int64, infrastructureId int64, infrastructure ParentInfrastructure, serviceStatus string, instanceType string, createdTimestamp string, config ContainerInstanceConfiguration, meta ContainerInstanceMeta, ) *ContainerInstance`

NewContainerInstance instantiates a new ContainerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceWithDefaults

`func NewContainerInstanceWithDefaults() *ContainerInstance`

NewContainerInstanceWithDefaults instantiates a new ContainerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ContainerInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContainerInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContainerInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetContainerId

`func (o *ContainerInstance) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ContainerInstance) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ContainerInstance) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ContainerInstance) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetTypeId

`func (o *ContainerInstance) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ContainerInstance) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ContainerInstance) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetSubdomain

`func (o *ContainerInstance) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ContainerInstance) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ContainerInstance) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ContainerInstance) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetVmPoolId

`func (o *ContainerInstance) GetVmPoolId() int64`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *ContainerInstance) GetVmPoolIdOk() (*int64, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *ContainerInstance) SetVmPoolId(v int64)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *ContainerInstance) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *ContainerInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ContainerInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ContainerInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetRamGB

`func (o *ContainerInstance) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *ContainerInstance) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *ContainerInstance) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetCpuCores

`func (o *ContainerInstance) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ContainerInstance) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ContainerInstance) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetOsTemplateId

`func (o *ContainerInstance) GetOsTemplateId() int64`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ContainerInstance) GetOsTemplateIdOk() (*int64, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ContainerInstance) SetOsTemplateId(v int64)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ContainerInstance) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ContainerInstance) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ContainerInstance) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ContainerInstance) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ContainerInstance) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ContainerInstance) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ContainerInstance) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ContainerInstance) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *ContainerInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerInstance) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *ContainerInstance) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ContainerInstance) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ContainerInstance) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetGroupId

`func (o *ContainerInstance) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ContainerInstance) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ContainerInstance) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetInfrastructureId

`func (o *ContainerInstance) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ContainerInstance) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ContainerInstance) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructure

`func (o *ContainerInstance) GetInfrastructure() ParentInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ContainerInstance) GetInfrastructureOk() (*ParentInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ContainerInstance) SetInfrastructure(v ParentInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.


### GetServiceStatus

`func (o *ContainerInstance) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ContainerInstance) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ContainerInstance) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetInstanceType

`func (o *ContainerInstance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ContainerInstance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ContainerInstance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetGpuInfo

`func (o *ContainerInstance) GetGpuInfo() []VMTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *ContainerInstance) GetGpuInfoOk() (*[]VMTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *ContainerInstance) SetGpuInfo(v []VMTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *ContainerInstance) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ContainerInstance) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ContainerInstance) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ContainerInstance) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ContainerInstance) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ContainerInstance) GetDnsSubdomainId() int64`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ContainerInstance) GetDnsSubdomainIdOk() (*int64, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ContainerInstance) SetDnsSubdomainId(v int64)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ContainerInstance) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ContainerInstance) GetDnsSubdomainPermanentId() int64`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ContainerInstance) GetDnsSubdomainPermanentIdOk() (*int64, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ContainerInstance) SetDnsSubdomainPermanentId(v int64)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ContainerInstance) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetUniqueIndex

`func (o *ContainerInstance) GetUniqueIndex() float32`

GetUniqueIndex returns the UniqueIndex field if non-nil, zero value otherwise.

### GetUniqueIndexOk

`func (o *ContainerInstance) GetUniqueIndexOk() (*float32, bool)`

GetUniqueIndexOk returns a tuple with the UniqueIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIndex

`func (o *ContainerInstance) SetUniqueIndex(v float32)`

SetUniqueIndex sets UniqueIndex field to given value.

### HasUniqueIndex

`func (o *ContainerInstance) HasUniqueIndex() bool`

HasUniqueIndex returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ContainerInstance) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ContainerInstance) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ContainerInstance) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetConfig

`func (o *ContainerInstance) GetConfig() ContainerInstanceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerInstance) GetConfigOk() (*ContainerInstanceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerInstance) SetConfig(v ContainerInstanceConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *ContainerInstance) GetMeta() ContainerInstanceMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContainerInstance) GetMetaOk() (*ContainerInstanceMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContainerInstance) SetMeta(v ContainerInstanceMeta)`

SetMeta sets Meta field to given value.


### GetAllocatedVMPoolInfo

`func (o *ContainerInstance) GetAllocatedVMPoolInfo() ContainerInstanceAllocatedVMPoolInfo`

GetAllocatedVMPoolInfo returns the AllocatedVMPoolInfo field if non-nil, zero value otherwise.

### GetAllocatedVMPoolInfoOk

`func (o *ContainerInstance) GetAllocatedVMPoolInfoOk() (*ContainerInstanceAllocatedVMPoolInfo, bool)`

GetAllocatedVMPoolInfoOk returns a tuple with the AllocatedVMPoolInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVMPoolInfo

`func (o *ContainerInstance) SetAllocatedVMPoolInfo(v ContainerInstanceAllocatedVMPoolInfo)`

SetAllocatedVMPoolInfo sets AllocatedVMPoolInfo field to given value.

### HasAllocatedVMPoolInfo

`func (o *ContainerInstance) HasAllocatedVMPoolInfo() bool`

HasAllocatedVMPoolInfo returns a boolean if a field has been set.

### GetAllocatedVMPoolGpusInfo

`func (o *ContainerInstance) GetAllocatedVMPoolGpusInfo() []VMPoolGPU`

GetAllocatedVMPoolGpusInfo returns the AllocatedVMPoolGpusInfo field if non-nil, zero value otherwise.

### GetAllocatedVMPoolGpusInfoOk

`func (o *ContainerInstance) GetAllocatedVMPoolGpusInfoOk() (*[]VMPoolGPU, bool)`

GetAllocatedVMPoolGpusInfoOk returns a tuple with the AllocatedVMPoolGpusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVMPoolGpusInfo

`func (o *ContainerInstance) SetAllocatedVMPoolGpusInfo(v []VMPoolGPU)`

SetAllocatedVMPoolGpusInfo sets AllocatedVMPoolGpusInfo field to given value.

### HasAllocatedVMPoolGpusInfo

`func (o *ContainerInstance) HasAllocatedVMPoolGpusInfo() bool`

HasAllocatedVMPoolGpusInfo returns a boolean if a field has been set.

### GetAllocatedVMPoolGpuPartitionsInfo

`func (o *ContainerInstance) GetAllocatedVMPoolGpuPartitionsInfo() []VMPoolGPUPartition`

GetAllocatedVMPoolGpuPartitionsInfo returns the AllocatedVMPoolGpuPartitionsInfo field if non-nil, zero value otherwise.

### GetAllocatedVMPoolGpuPartitionsInfoOk

`func (o *ContainerInstance) GetAllocatedVMPoolGpuPartitionsInfoOk() (*[]VMPoolGPUPartition, bool)`

GetAllocatedVMPoolGpuPartitionsInfoOk returns a tuple with the AllocatedVMPoolGpuPartitionsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedVMPoolGpuPartitionsInfo

`func (o *ContainerInstance) SetAllocatedVMPoolGpuPartitionsInfo(v []VMPoolGPUPartition)`

SetAllocatedVMPoolGpuPartitionsInfo sets AllocatedVMPoolGpuPartitionsInfo field to given value.

### HasAllocatedVMPoolGpuPartitionsInfo

`func (o *ContainerInstance) HasAllocatedVMPoolGpuPartitionsInfo() bool`

HasAllocatedVMPoolGpuPartitionsInfo returns a boolean if a field has been set.

### GetLinks

`func (o *ContainerInstance) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContainerInstance) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContainerInstance) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContainerInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



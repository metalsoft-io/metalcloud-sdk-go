# ContainerInstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int64** | Revision of the Container Instance Configuration | 
**Label** | **string** | Name of the Container Instance. | 
**ContainerId** | Pointer to **int64** | Id of the Container. | [optional] 
**TypeId** | **int64** | Id of the Container Type. | 
**Subdomain** | Pointer to **string** | Subdomain of the Container Instance. | [optional] 
**DnsSubdomainChangeId** | Pointer to **int64** | Id of the DNS subdomain for the Container Instance. | [optional] 
**DeployType** | **string** | Deploy type of the Container Instance | [default to "create"]
**DeployStatus** | **string** | Deploy status of the Container Instance | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **int64** | Id of the deployment for the Container Instance. | [optional] 
**VmPoolId** | Pointer to **int64** | Id of the VM Pool. | [optional] 
**DiskSizeGB** | **float32** | Disk size in GB of the Container Instance. | 
**RamGB** | **float32** | RAM size in GB of the Container Instance. | 
**CpuCores** | **float32** | Number of CPU cores for the Container Instance. | 
**GpuInfo** | Pointer to [**[]VMTypeGPUInfo**](VMTypeGPUInfo.md) | Information about GPUs available for this Container Instance | [optional] 
**OsTemplateId** | Pointer to **int64** | Id of the template used by the Container Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the Container Instance. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Container Instance last update. | 

## Methods

### NewContainerInstanceConfiguration

`func NewContainerInstanceConfiguration(revision int64, label string, typeId int64, deployType string, deployStatus string, diskSizeGB float32, ramGB float32, cpuCores float32, updatedTimestamp string, ) *ContainerInstanceConfiguration`

NewContainerInstanceConfiguration instantiates a new ContainerInstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceConfigurationWithDefaults

`func NewContainerInstanceConfigurationWithDefaults() *ContainerInstanceConfiguration`

NewContainerInstanceConfigurationWithDefaults instantiates a new ContainerInstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ContainerInstanceConfiguration) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ContainerInstanceConfiguration) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ContainerInstanceConfiguration) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ContainerInstanceConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContainerInstanceConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContainerInstanceConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetContainerId

`func (o *ContainerInstanceConfiguration) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ContainerInstanceConfiguration) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ContainerInstanceConfiguration) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ContainerInstanceConfiguration) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetTypeId

`func (o *ContainerInstanceConfiguration) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *ContainerInstanceConfiguration) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *ContainerInstanceConfiguration) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetSubdomain

`func (o *ContainerInstanceConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ContainerInstanceConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ContainerInstanceConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ContainerInstanceConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *ContainerInstanceConfiguration) GetDnsSubdomainChangeId() int64`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ContainerInstanceConfiguration) GetDnsSubdomainChangeIdOk() (*int64, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ContainerInstanceConfiguration) SetDnsSubdomainChangeId(v int64)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ContainerInstanceConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDeployType

`func (o *ContainerInstanceConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ContainerInstanceConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ContainerInstanceConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ContainerInstanceConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ContainerInstanceConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ContainerInstanceConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *ContainerInstanceConfiguration) GetInfrastructureDeployId() int64`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ContainerInstanceConfiguration) GetInfrastructureDeployIdOk() (*int64, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ContainerInstanceConfiguration) SetInfrastructureDeployId(v int64)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ContainerInstanceConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetVmPoolId

`func (o *ContainerInstanceConfiguration) GetVmPoolId() int64`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *ContainerInstanceConfiguration) GetVmPoolIdOk() (*int64, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *ContainerInstanceConfiguration) SetVmPoolId(v int64)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *ContainerInstanceConfiguration) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *ContainerInstanceConfiguration) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ContainerInstanceConfiguration) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ContainerInstanceConfiguration) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetRamGB

`func (o *ContainerInstanceConfiguration) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *ContainerInstanceConfiguration) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *ContainerInstanceConfiguration) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetCpuCores

`func (o *ContainerInstanceConfiguration) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ContainerInstanceConfiguration) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ContainerInstanceConfiguration) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetGpuInfo

`func (o *ContainerInstanceConfiguration) GetGpuInfo() []VMTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *ContainerInstanceConfiguration) GetGpuInfoOk() (*[]VMTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *ContainerInstanceConfiguration) SetGpuInfo(v []VMTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *ContainerInstanceConfiguration) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ContainerInstanceConfiguration) GetOsTemplateId() int64`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ContainerInstanceConfiguration) GetOsTemplateIdOk() (*int64, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ContainerInstanceConfiguration) SetOsTemplateId(v int64)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ContainerInstanceConfiguration) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ContainerInstanceConfiguration) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ContainerInstanceConfiguration) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ContainerInstanceConfiguration) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ContainerInstanceConfiguration) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ContainerInstanceConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ContainerInstanceConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ContainerInstanceConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Container ID | 
**Name** | **string** | Name of the Container | 
**SiteId** | **int64** | Id of the site for the Container | 
**InfrastructureId** | **int64** | ID of the infrastructure where this Container is deployed | 
**UserId** | **int64** | ID of the user that owns this Container | 
**InstanceId** | **int64** | ID of the instance where this Container is deployed | 
**ContainerInstanceId** | **int64** | The id of the Container Instance. This is a number. | 
**Host** | **string** | Name of the host | 
**Hosts** | **[]string** | List of hosts | 
**CpuCores** | **float32** | Number of CPU cores for the Container | 
**RamGB** | **float32** | RAM in GB for the Container | 
**DiskSizeGB** | **float32** | Disk size in GB for the Container | 
**GpuInfo** | Pointer to [**[]VMGpuInfoDto**](VMGpuInfoDto.md) | Information about GPUs assigned to this Container | [optional] 
**TypeId** | **int64** | The id of the Container Type. This is a number. | 
**PoolId** | **int64** | The id of the VM Pool. This is a number. | 
**AdministrationState** | **string** | The administration state of the Container. | 
**Comments** | Pointer to **float32** | Container comments. | [optional] 
**NumaNodes** | Pointer to **[]float32** | NUMA nodes the Container is assigned to. Only applicable for Containers with GPU / DPU requirements. | [optional] 
**PowerState** | **string** | The power state of the Container. | 
**PowerStateLastUpdatedTimestamp** | **string** | Timestamp when the Container power state was last updated. | 
**CreatedTimestamp** | **string** | Timestamp when the Container was created | 
**AllocationTimestamp** | **string** | Timestamp when the Container was allocated | 
**Tags** | Pointer to **[]string** | Tags for the Container. This is a JSON object. | [optional] 
**Disks** | [**[]ContainerDisk**](ContainerDisk.md) | The disks of the Container. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewContainer

`func NewContainer(id int64, name string, siteId int64, infrastructureId int64, userId int64, instanceId int64, containerInstanceId int64, host string, hosts []string, cpuCores float32, ramGB float32, diskSizeGB float32, typeId int64, poolId int64, administrationState string, powerState string, powerStateLastUpdatedTimestamp string, createdTimestamp string, allocationTimestamp string, disks []ContainerDisk, ) *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Container) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Container) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Container) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Container) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *Container) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Container) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Container) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetInfrastructureId

`func (o *Container) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Container) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Container) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetUserId

`func (o *Container) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Container) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Container) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetInstanceId

`func (o *Container) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Container) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Container) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetContainerInstanceId

`func (o *Container) GetContainerInstanceId() int64`

GetContainerInstanceId returns the ContainerInstanceId field if non-nil, zero value otherwise.

### GetContainerInstanceIdOk

`func (o *Container) GetContainerInstanceIdOk() (*int64, bool)`

GetContainerInstanceIdOk returns a tuple with the ContainerInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerInstanceId

`func (o *Container) SetContainerInstanceId(v int64)`

SetContainerInstanceId sets ContainerInstanceId field to given value.


### GetHost

`func (o *Container) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Container) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Container) SetHost(v string)`

SetHost sets Host field to given value.


### GetHosts

`func (o *Container) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *Container) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *Container) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetCpuCores

`func (o *Container) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *Container) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *Container) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetRamGB

`func (o *Container) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *Container) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *Container) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetDiskSizeGB

`func (o *Container) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *Container) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *Container) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetGpuInfo

`func (o *Container) GetGpuInfo() []VMGpuInfoDto`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *Container) GetGpuInfoOk() (*[]VMGpuInfoDto, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *Container) SetGpuInfo(v []VMGpuInfoDto)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *Container) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetTypeId

`func (o *Container) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Container) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Container) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetPoolId

`func (o *Container) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Container) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Container) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetAdministrationState

`func (o *Container) GetAdministrationState() string`

GetAdministrationState returns the AdministrationState field if non-nil, zero value otherwise.

### GetAdministrationStateOk

`func (o *Container) GetAdministrationStateOk() (*string, bool)`

GetAdministrationStateOk returns a tuple with the AdministrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrationState

`func (o *Container) SetAdministrationState(v string)`

SetAdministrationState sets AdministrationState field to given value.


### GetComments

`func (o *Container) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Container) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Container) SetComments(v float32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Container) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetNumaNodes

`func (o *Container) GetNumaNodes() []float32`

GetNumaNodes returns the NumaNodes field if non-nil, zero value otherwise.

### GetNumaNodesOk

`func (o *Container) GetNumaNodesOk() (*[]float32, bool)`

GetNumaNodesOk returns a tuple with the NumaNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNodes

`func (o *Container) SetNumaNodes(v []float32)`

SetNumaNodes sets NumaNodes field to given value.

### HasNumaNodes

`func (o *Container) HasNumaNodes() bool`

HasNumaNodes returns a boolean if a field has been set.

### GetPowerState

`func (o *Container) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *Container) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *Container) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.


### GetPowerStateLastUpdatedTimestamp

`func (o *Container) GetPowerStateLastUpdatedTimestamp() string`

GetPowerStateLastUpdatedTimestamp returns the PowerStateLastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetPowerStateLastUpdatedTimestampOk

`func (o *Container) GetPowerStateLastUpdatedTimestampOk() (*string, bool)`

GetPowerStateLastUpdatedTimestampOk returns a tuple with the PowerStateLastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateLastUpdatedTimestamp

`func (o *Container) SetPowerStateLastUpdatedTimestamp(v string)`

SetPowerStateLastUpdatedTimestamp sets PowerStateLastUpdatedTimestamp field to given value.


### GetCreatedTimestamp

`func (o *Container) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Container) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Container) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetAllocationTimestamp

`func (o *Container) GetAllocationTimestamp() string`

GetAllocationTimestamp returns the AllocationTimestamp field if non-nil, zero value otherwise.

### GetAllocationTimestampOk

`func (o *Container) GetAllocationTimestampOk() (*string, bool)`

GetAllocationTimestampOk returns a tuple with the AllocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimestamp

`func (o *Container) SetAllocationTimestamp(v string)`

SetAllocationTimestamp sets AllocationTimestamp field to given value.


### GetTags

`func (o *Container) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Container) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Container) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Container) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDisks

`func (o *Container) GetDisks() []ContainerDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *Container) GetDisksOk() (*[]ContainerDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *Container) SetDisks(v []ContainerDisk)`

SetDisks sets Disks field to given value.


### GetLinks

`func (o *Container) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Container) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Container) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Container) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



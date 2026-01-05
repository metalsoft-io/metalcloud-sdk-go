# VM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM ID | 
**Name** | **string** | Name of the VM | 
**SiteId** | **float32** | Id of the site for the VM | 
**DatacenterName** | **string** | Datacenter of the VM | 
**InfrastructureId** | **float32** | ID of the infrastructure where this VM is deployed | 
**UserId** | **float32** | ID of the user that owns this VM | 
**UserEmail** | **string** | Email of the user that owns this VM | 
**InstanceId** | **float32** | ID of the instance where this VM is deployed | 
**VmInstanceId** | **float32** | The id of the VM Instance. This is a number. | 
**Host** | **string** | Name of the host | 
**Hosts** | **[]string** | List of hosts | 
**CpuCores** | **float32** | Number of CPU cores for the VM | 
**RamGB** | **float32** | RAM in GB for the VM | 
**DiskSizeGB** | **float32** | Disk size in GB for the VM | 
**TypeId** | **float32** | The id of the VM Type. This is a number. | 
**PoolId** | **float32** | The id of the VM Pool. This is a number. | 
**AdministrationState** | **string** | The administration state of the VM. | 
**Comments** | Pointer to **float32** | VM comments. | [optional] 
**PowerState** | **string** | The power state of the VM. | 
**PowerStateLastUpdatedTimestamp** | **string** | Timestamp when the VM power state was last updated. | 
**CreatedTimestamp** | **string** | Timestamp when the VM was created | 
**AllocationTimestamp** | **string** | Timestamp when the VM was allocated | 
**Tags** | Pointer to **[]string** | Tags for the VM. This is a JSON object. | [optional] 
**Disks** | [**[]VMDisk**](VMDisk.md) | The disks of the VM. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewVM

`func NewVM(id float32, name string, siteId float32, datacenterName string, infrastructureId float32, userId float32, userEmail string, instanceId float32, vmInstanceId float32, host string, hosts []string, cpuCores float32, ramGB float32, diskSizeGB float32, typeId float32, poolId float32, administrationState string, powerState string, powerStateLastUpdatedTimestamp string, createdTimestamp string, allocationTimestamp string, disks []VMDisk, ) *VM`

NewVM instantiates a new VM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMWithDefaults

`func NewVMWithDefaults() *VM`

NewVMWithDefaults instantiates a new VM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VM) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VM) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VM) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *VM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VM) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *VM) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VM) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VM) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetDatacenterName

`func (o *VM) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *VM) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *VM) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetInfrastructureId

`func (o *VM) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VM) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VM) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetUserId

`func (o *VM) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VM) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VM) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetUserEmail

`func (o *VM) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *VM) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *VM) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetInstanceId

`func (o *VM) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *VM) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *VM) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### GetVmInstanceId

`func (o *VM) GetVmInstanceId() float32`

GetVmInstanceId returns the VmInstanceId field if non-nil, zero value otherwise.

### GetVmInstanceIdOk

`func (o *VM) GetVmInstanceIdOk() (*float32, bool)`

GetVmInstanceIdOk returns a tuple with the VmInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceId

`func (o *VM) SetVmInstanceId(v float32)`

SetVmInstanceId sets VmInstanceId field to given value.


### GetHost

`func (o *VM) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VM) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VM) SetHost(v string)`

SetHost sets Host field to given value.


### GetHosts

`func (o *VM) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VM) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VM) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetCpuCores

`func (o *VM) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *VM) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *VM) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetRamGB

`func (o *VM) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *VM) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *VM) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetDiskSizeGB

`func (o *VM) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *VM) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *VM) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetTypeId

`func (o *VM) GetTypeId() float32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VM) GetTypeIdOk() (*float32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VM) SetTypeId(v float32)`

SetTypeId sets TypeId field to given value.


### GetPoolId

`func (o *VM) GetPoolId() float32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *VM) GetPoolIdOk() (*float32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *VM) SetPoolId(v float32)`

SetPoolId sets PoolId field to given value.


### GetAdministrationState

`func (o *VM) GetAdministrationState() string`

GetAdministrationState returns the AdministrationState field if non-nil, zero value otherwise.

### GetAdministrationStateOk

`func (o *VM) GetAdministrationStateOk() (*string, bool)`

GetAdministrationStateOk returns a tuple with the AdministrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrationState

`func (o *VM) SetAdministrationState(v string)`

SetAdministrationState sets AdministrationState field to given value.


### GetComments

`func (o *VM) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VM) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VM) SetComments(v float32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VM) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPowerState

`func (o *VM) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VM) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VM) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.


### GetPowerStateLastUpdatedTimestamp

`func (o *VM) GetPowerStateLastUpdatedTimestamp() string`

GetPowerStateLastUpdatedTimestamp returns the PowerStateLastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetPowerStateLastUpdatedTimestampOk

`func (o *VM) GetPowerStateLastUpdatedTimestampOk() (*string, bool)`

GetPowerStateLastUpdatedTimestampOk returns a tuple with the PowerStateLastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateLastUpdatedTimestamp

`func (o *VM) SetPowerStateLastUpdatedTimestamp(v string)`

SetPowerStateLastUpdatedTimestamp sets PowerStateLastUpdatedTimestamp field to given value.


### GetCreatedTimestamp

`func (o *VM) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VM) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VM) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetAllocationTimestamp

`func (o *VM) GetAllocationTimestamp() string`

GetAllocationTimestamp returns the AllocationTimestamp field if non-nil, zero value otherwise.

### GetAllocationTimestampOk

`func (o *VM) GetAllocationTimestampOk() (*string, bool)`

GetAllocationTimestampOk returns a tuple with the AllocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimestamp

`func (o *VM) SetAllocationTimestamp(v string)`

SetAllocationTimestamp sets AllocationTimestamp field to given value.


### GetTags

`func (o *VM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDisks

`func (o *VM) GetDisks() []VMDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VM) GetDisksOk() (*[]VMDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VM) SetDisks(v []VMDisk)`

SetDisks sets Disks field to given value.


### GetLinks

`func (o *VM) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VM) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VM) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VM) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



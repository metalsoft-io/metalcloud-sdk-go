# VMPoolGPUPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the GPU partition | 
**VmPoolId** | **int64** | ID of the VM Pool linked to the GPU partition | 
**HostId** | Pointer to **int64** | ID of the host linked to the GPU partition | [optional] 
**Label** | **float32** | Label of the GPU partition | 
**GpuCount** | **float32** | Number of GPUs in the partition | 
**VmInstanceId** | Pointer to **int64** | ID of the VM Instance the GPU partition is assigned to | [optional] 
**DeploymentStatus** | **string** | Deployment status of the GPU partition | 
**Gpus** | Pointer to [**[]VMPoolGPU**](VMPoolGPU.md) | List of GPUs in that partition | [optional] 

## Methods

### NewVMPoolGPUPartition

`func NewVMPoolGPUPartition(id int64, vmPoolId int64, label float32, gpuCount float32, deploymentStatus string, ) *VMPoolGPUPartition`

NewVMPoolGPUPartition instantiates a new VMPoolGPUPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolGPUPartitionWithDefaults

`func NewVMPoolGPUPartitionWithDefaults() *VMPoolGPUPartition`

NewVMPoolGPUPartitionWithDefaults instantiates a new VMPoolGPUPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPoolGPUPartition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPoolGPUPartition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPoolGPUPartition) SetId(v int64)`

SetId sets Id field to given value.


### GetVmPoolId

`func (o *VMPoolGPUPartition) GetVmPoolId() int64`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMPoolGPUPartition) GetVmPoolIdOk() (*int64, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMPoolGPUPartition) SetVmPoolId(v int64)`

SetVmPoolId sets VmPoolId field to given value.


### GetHostId

`func (o *VMPoolGPUPartition) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *VMPoolGPUPartition) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *VMPoolGPUPartition) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *VMPoolGPUPartition) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetLabel

`func (o *VMPoolGPUPartition) GetLabel() float32`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMPoolGPUPartition) GetLabelOk() (*float32, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMPoolGPUPartition) SetLabel(v float32)`

SetLabel sets Label field to given value.


### GetGpuCount

`func (o *VMPoolGPUPartition) GetGpuCount() float32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *VMPoolGPUPartition) GetGpuCountOk() (*float32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *VMPoolGPUPartition) SetGpuCount(v float32)`

SetGpuCount sets GpuCount field to given value.


### GetVmInstanceId

`func (o *VMPoolGPUPartition) GetVmInstanceId() int64`

GetVmInstanceId returns the VmInstanceId field if non-nil, zero value otherwise.

### GetVmInstanceIdOk

`func (o *VMPoolGPUPartition) GetVmInstanceIdOk() (*int64, bool)`

GetVmInstanceIdOk returns a tuple with the VmInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceId

`func (o *VMPoolGPUPartition) SetVmInstanceId(v int64)`

SetVmInstanceId sets VmInstanceId field to given value.

### HasVmInstanceId

`func (o *VMPoolGPUPartition) HasVmInstanceId() bool`

HasVmInstanceId returns a boolean if a field has been set.

### GetDeploymentStatus

`func (o *VMPoolGPUPartition) GetDeploymentStatus() string`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *VMPoolGPUPartition) GetDeploymentStatusOk() (*string, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *VMPoolGPUPartition) SetDeploymentStatus(v string)`

SetDeploymentStatus sets DeploymentStatus field to given value.


### GetGpus

`func (o *VMPoolGPUPartition) GetGpus() []VMPoolGPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *VMPoolGPUPartition) GetGpusOk() (*[]VMPoolGPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *VMPoolGPUPartition) SetGpus(v []VMPoolGPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *VMPoolGPUPartition) HasGpus() bool`

HasGpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



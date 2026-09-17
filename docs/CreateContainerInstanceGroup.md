# CreateContainerInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | Pointer to **float32** |  | [optional] [default to 1]
**DiskSizeGB** | **float32** | Disk size in GB for each VM Instance in the VM Instance Group. | 
**TypeId** | **int64** | Id of the VM Type. | 
**OsTemplateId** | **int64** | Id of the template used by the VM Instance Group. | 
**Tags** | Pointer to **[]string** | Tags for the VM Instance Group. | [optional] 
**VmPoolId** | **int64** | Id of the VM Pool on which the VM Instance Group will be provisioned. | 

## Methods

### NewCreateContainerInstanceGroup

`func NewCreateContainerInstanceGroup(diskSizeGB float32, typeId int64, osTemplateId int64, vmPoolId int64, ) *CreateContainerInstanceGroup`

NewCreateContainerInstanceGroup instantiates a new CreateContainerInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerInstanceGroupWithDefaults

`func NewCreateContainerInstanceGroupWithDefaults() *CreateContainerInstanceGroup`

NewCreateContainerInstanceGroupWithDefaults instantiates a new CreateContainerInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *CreateContainerInstanceGroup) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *CreateContainerInstanceGroup) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *CreateContainerInstanceGroup) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *CreateContainerInstanceGroup) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *CreateContainerInstanceGroup) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *CreateContainerInstanceGroup) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *CreateContainerInstanceGroup) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetTypeId

`func (o *CreateContainerInstanceGroup) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CreateContainerInstanceGroup) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CreateContainerInstanceGroup) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetOsTemplateId

`func (o *CreateContainerInstanceGroup) GetOsTemplateId() int64`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *CreateContainerInstanceGroup) GetOsTemplateIdOk() (*int64, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *CreateContainerInstanceGroup) SetOsTemplateId(v int64)`

SetOsTemplateId sets OsTemplateId field to given value.


### GetTags

`func (o *CreateContainerInstanceGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContainerInstanceGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContainerInstanceGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContainerInstanceGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVmPoolId

`func (o *CreateContainerInstanceGroup) GetVmPoolId() int64`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *CreateContainerInstanceGroup) GetVmPoolIdOk() (*int64, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *CreateContainerInstanceGroup) SetVmPoolId(v int64)`

SetVmPoolId sets VmPoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



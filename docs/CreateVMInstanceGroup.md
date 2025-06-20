# CreateVMInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | Pointer to **float32** |  | [optional] [default to 1]
**DiskSizeGB** | **float32** | Disk size in GB for each VM Instance in the VM Instance Group. | 
**TypeId** | **float32** | Id of the VM Type. | 
**OsTemplateId** | **float32** | Id of the template used by the VM Instance Group. | 
**Tags** | Pointer to **[]string** | Tags for the VM Instance Group. | [optional] 

## Methods

### NewCreateVMInstanceGroup

`func NewCreateVMInstanceGroup(diskSizeGB float32, typeId float32, osTemplateId float32, ) *CreateVMInstanceGroup`

NewCreateVMInstanceGroup instantiates a new CreateVMInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMInstanceGroupWithDefaults

`func NewCreateVMInstanceGroupWithDefaults() *CreateVMInstanceGroup`

NewCreateVMInstanceGroupWithDefaults instantiates a new CreateVMInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *CreateVMInstanceGroup) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *CreateVMInstanceGroup) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *CreateVMInstanceGroup) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *CreateVMInstanceGroup) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *CreateVMInstanceGroup) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *CreateVMInstanceGroup) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *CreateVMInstanceGroup) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetTypeId

`func (o *CreateVMInstanceGroup) GetTypeId() float32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CreateVMInstanceGroup) GetTypeIdOk() (*float32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CreateVMInstanceGroup) SetTypeId(v float32)`

SetTypeId sets TypeId field to given value.


### GetOsTemplateId

`func (o *CreateVMInstanceGroup) GetOsTemplateId() float32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *CreateVMInstanceGroup) GetOsTemplateIdOk() (*float32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *CreateVMInstanceGroup) SetOsTemplateId(v float32)`

SetOsTemplateId sets OsTemplateId field to given value.


### GetTags

`func (o *CreateVMInstanceGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVMInstanceGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVMInstanceGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVMInstanceGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



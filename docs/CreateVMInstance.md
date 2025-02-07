# CreateVMInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **float32** | Id of the VM Instance Group. | 
**TypeId** | **float32** | Id of the VM Type. | 
**Tags** | Pointer to **[]string** | Tags for the VM Instance. | [optional] 
**DiskSizeGB** | Pointer to **float32** | Disk size in GB of the VM Instance. If not passed, the default disk size from the group will be used | [optional] 

## Methods

### NewCreateVMInstance

`func NewCreateVMInstance(groupId float32, typeId float32, ) *CreateVMInstance`

NewCreateVMInstance instantiates a new CreateVMInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMInstanceWithDefaults

`func NewCreateVMInstanceWithDefaults() *CreateVMInstance`

NewCreateVMInstanceWithDefaults instantiates a new CreateVMInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CreateVMInstance) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateVMInstance) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateVMInstance) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetTypeId

`func (o *CreateVMInstance) GetTypeId() float32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CreateVMInstance) GetTypeIdOk() (*float32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CreateVMInstance) SetTypeId(v float32)`

SetTypeId sets TypeId field to given value.


### GetTags

`func (o *CreateVMInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVMInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVMInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVMInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *CreateVMInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *CreateVMInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *CreateVMInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *CreateVMInstance) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



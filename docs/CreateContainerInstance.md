# CreateContainerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeId** | **int64** | Id of the VM Type. | 
**GroupId** | **int64** | Id of the VM Instance Group. | 
**DiskSizeGB** | Pointer to **float32** | Disk size in GB of the Instance. If not passed, the default disk size from the group will be used | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Instance. | [optional] 

## Methods

### NewCreateContainerInstance

`func NewCreateContainerInstance(typeId int64, groupId int64, ) *CreateContainerInstance`

NewCreateContainerInstance instantiates a new CreateContainerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerInstanceWithDefaults

`func NewCreateContainerInstanceWithDefaults() *CreateContainerInstance`

NewCreateContainerInstanceWithDefaults instantiates a new CreateContainerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeId

`func (o *CreateContainerInstance) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CreateContainerInstance) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CreateContainerInstance) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetGroupId

`func (o *CreateContainerInstance) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateContainerInstance) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateContainerInstance) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetDiskSizeGB

`func (o *CreateContainerInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *CreateContainerInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *CreateContainerInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.

### HasDiskSizeGB

`func (o *CreateContainerInstance) HasDiskSizeGB() bool`

HasDiskSizeGB returns a boolean if a field has been set.

### GetTags

`func (o *CreateContainerInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContainerInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContainerInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContainerInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



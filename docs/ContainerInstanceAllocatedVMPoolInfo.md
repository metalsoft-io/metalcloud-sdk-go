# ContainerInstanceAllocatedVMPoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the allocated VM Pool. | 
**Type** | **string** | Driver type of the allocated VM Pool. | 

## Methods

### NewContainerInstanceAllocatedVMPoolInfo

`func NewContainerInstanceAllocatedVMPoolInfo(id int64, type_ string, ) *ContainerInstanceAllocatedVMPoolInfo`

NewContainerInstanceAllocatedVMPoolInfo instantiates a new ContainerInstanceAllocatedVMPoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceAllocatedVMPoolInfoWithDefaults

`func NewContainerInstanceAllocatedVMPoolInfoWithDefaults() *ContainerInstanceAllocatedVMPoolInfo`

NewContainerInstanceAllocatedVMPoolInfoWithDefaults instantiates a new ContainerInstanceAllocatedVMPoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerInstanceAllocatedVMPoolInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerInstanceAllocatedVMPoolInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerInstanceAllocatedVMPoolInfo) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *ContainerInstanceAllocatedVMPoolInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerInstanceAllocatedVMPoolInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerInstanceAllocatedVMPoolInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ContainerDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Container ID | 
**ContainerId** | **int64** | Container ID | 
**Datastore** | **string** | Name of the Container Disk datastore | 
**SizeGB** | **float32** | Size of the Container Disk in GB | 

## Methods

### NewContainerDisk

`func NewContainerDisk(id int64, containerId int64, datastore string, sizeGB float32, ) *ContainerDisk`

NewContainerDisk instantiates a new ContainerDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerDiskWithDefaults

`func NewContainerDiskWithDefaults() *ContainerDisk`

NewContainerDiskWithDefaults instantiates a new ContainerDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerDisk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerDisk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerDisk) SetId(v int64)`

SetId sets Id field to given value.


### GetContainerId

`func (o *ContainerDisk) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ContainerDisk) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ContainerDisk) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.


### GetDatastore

`func (o *ContainerDisk) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ContainerDisk) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ContainerDisk) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.


### GetSizeGB

`func (o *ContainerDisk) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *ContainerDisk) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *ContainerDisk) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



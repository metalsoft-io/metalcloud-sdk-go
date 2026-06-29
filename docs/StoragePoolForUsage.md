# StoragePoolForUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the Storage Pool | 
**Name** | **string** | Name of the Storage Pool | 
**Technologies** | **[]string** | Storage technology | 
**Driver** | **string** | Driver of the Storage Pool | 
**FabricId** | **int64** | Fabric ID of the Storage Pool | 

## Methods

### NewStoragePoolForUsage

`func NewStoragePoolForUsage(id int64, name string, technologies []string, driver string, fabricId int64, ) *StoragePoolForUsage`

NewStoragePoolForUsage instantiates a new StoragePoolForUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePoolForUsageWithDefaults

`func NewStoragePoolForUsageWithDefaults() *StoragePoolForUsage`

NewStoragePoolForUsageWithDefaults instantiates a new StoragePoolForUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoragePoolForUsage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoragePoolForUsage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoragePoolForUsage) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *StoragePoolForUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePoolForUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePoolForUsage) SetName(v string)`

SetName sets Name field to given value.


### GetTechnologies

`func (o *StoragePoolForUsage) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *StoragePoolForUsage) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *StoragePoolForUsage) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.


### GetDriver

`func (o *StoragePoolForUsage) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *StoragePoolForUsage) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *StoragePoolForUsage) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetFabricId

`func (o *StoragePoolForUsage) GetFabricId() int64`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *StoragePoolForUsage) GetFabricIdOk() (*int64, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *StoragePoolForUsage) SetFabricId(v int64)`

SetFabricId sets FabricId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



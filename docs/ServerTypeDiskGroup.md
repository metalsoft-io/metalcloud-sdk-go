# ServerTypeDiskGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageControllers** | **map[string]interface{}** | The storage controllers for the server type. Key is the controller name and value is the controller information. | 

## Methods

### NewServerTypeDiskGroup

`func NewServerTypeDiskGroup(storageControllers map[string]interface{}, ) *ServerTypeDiskGroup`

NewServerTypeDiskGroup instantiates a new ServerTypeDiskGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeDiskGroupWithDefaults

`func NewServerTypeDiskGroupWithDefaults() *ServerTypeDiskGroup`

NewServerTypeDiskGroupWithDefaults instantiates a new ServerTypeDiskGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageControllers

`func (o *ServerTypeDiskGroup) GetStorageControllers() map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ServerTypeDiskGroup) GetStorageControllersOk() (*map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ServerTypeDiskGroup) SetStorageControllers(v map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



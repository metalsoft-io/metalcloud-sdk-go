# ServerTypeStorageController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageControllerOptions** | [**ServerTypeStorageControllerOptions**](ServerTypeStorageControllerOptions.md) | The options for the storage controller. | 
**Disks** | [**[]ServerTypeDisks**](ServerTypeDisks.md) | The disks information for the server type. | 

## Methods

### NewServerTypeStorageController

`func NewServerTypeStorageController(storageControllerOptions ServerTypeStorageControllerOptions, disks []ServerTypeDisks, ) *ServerTypeStorageController`

NewServerTypeStorageController instantiates a new ServerTypeStorageController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeStorageControllerWithDefaults

`func NewServerTypeStorageControllerWithDefaults() *ServerTypeStorageController`

NewServerTypeStorageControllerWithDefaults instantiates a new ServerTypeStorageController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageControllerOptions

`func (o *ServerTypeStorageController) GetStorageControllerOptions() ServerTypeStorageControllerOptions`

GetStorageControllerOptions returns the StorageControllerOptions field if non-nil, zero value otherwise.

### GetStorageControllerOptionsOk

`func (o *ServerTypeStorageController) GetStorageControllerOptionsOk() (*ServerTypeStorageControllerOptions, bool)`

GetStorageControllerOptionsOk returns a tuple with the StorageControllerOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllerOptions

`func (o *ServerTypeStorageController) SetStorageControllerOptions(v ServerTypeStorageControllerOptions)`

SetStorageControllerOptions sets StorageControllerOptions field to given value.


### GetDisks

`func (o *ServerTypeStorageController) GetDisks() []ServerTypeDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *ServerTypeStorageController) GetDisksOk() (*[]ServerTypeDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *ServerTypeStorageController) SetDisks(v []ServerTypeDisks)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ServerInstanceStorageProfileController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Storage Controller. | 
**Mode** | **string** | The mode of the Storage Controller. | 
**Volumes** | [**[]ServerInstanceStorageProfileControllerVolume**](ServerInstanceStorageProfileControllerVolume.md) | The information for each Volume of the Storage Controller. | 

## Methods

### NewServerInstanceStorageProfileController

`func NewServerInstanceStorageProfileController(id string, mode string, volumes []ServerInstanceStorageProfileControllerVolume, ) *ServerInstanceStorageProfileController`

NewServerInstanceStorageProfileController instantiates a new ServerInstanceStorageProfileController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceStorageProfileControllerWithDefaults

`func NewServerInstanceStorageProfileControllerWithDefaults() *ServerInstanceStorageProfileController`

NewServerInstanceStorageProfileControllerWithDefaults instantiates a new ServerInstanceStorageProfileController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceStorageProfileController) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceStorageProfileController) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceStorageProfileController) SetId(v string)`

SetId sets Id field to given value.


### GetMode

`func (o *ServerInstanceStorageProfileController) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ServerInstanceStorageProfileController) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ServerInstanceStorageProfileController) SetMode(v string)`

SetMode sets Mode field to given value.


### GetVolumes

`func (o *ServerInstanceStorageProfileController) GetVolumes() []ServerInstanceStorageProfileControllerVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ServerInstanceStorageProfileController) GetVolumesOk() (*[]ServerInstanceStorageProfileControllerVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ServerInstanceStorageProfileController) SetVolumes(v []ServerInstanceStorageProfileControllerVolume)`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



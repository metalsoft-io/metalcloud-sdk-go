# ServerStorageController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The id of the storage controller. | 
**ServerId** | **float32** | The id of the server. | 
**Name** | **string** | The name of the storage controller. | 
**Label** | **string** | The label of the storage controller. | 
**Description** | **string** | The description of the storage controller. | 
**Options** | [**ServerTypeStorageControllerOptions**](ServerTypeStorageControllerOptions.md) | The options of the storage controller. | 
**Mode** | **string** | The mode of the storage controller. | 

## Methods

### NewServerStorageController

`func NewServerStorageController(id float32, serverId float32, name string, label string, description string, options ServerTypeStorageControllerOptions, mode string, ) *ServerStorageController`

NewServerStorageController instantiates a new ServerStorageController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStorageControllerWithDefaults

`func NewServerStorageControllerWithDefaults() *ServerStorageController`

NewServerStorageControllerWithDefaults instantiates a new ServerStorageController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerStorageController) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerStorageController) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerStorageController) SetId(v float32)`

SetId sets Id field to given value.


### GetServerId

`func (o *ServerStorageController) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerStorageController) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerStorageController) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetName

`func (o *ServerStorageController) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerStorageController) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerStorageController) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ServerStorageController) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerStorageController) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerStorageController) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ServerStorageController) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerStorageController) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerStorageController) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOptions

`func (o *ServerStorageController) GetOptions() ServerTypeStorageControllerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ServerStorageController) GetOptionsOk() (*ServerTypeStorageControllerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ServerStorageController) SetOptions(v ServerTypeStorageControllerOptions)`

SetOptions sets Options field to given value.


### GetMode

`func (o *ServerStorageController) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ServerStorageController) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ServerStorageController) SetMode(v string)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateNetworkFabricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The network fabric name | 
**Description** | Pointer to **string** | Network fabric description | [optional] 
**FabricType** | **string** | The type of network fabric | 
**FabricConfiguration** | [**NetworkFabricFabricConfiguration**](NetworkFabricFabricConfiguration.md) |  | 
**Revision** | **float32** | Revision number of the entity | 
**Id** | **string** | The network fabric ID. | 

## Methods

### NewUpdateNetworkFabricDto

`func NewUpdateNetworkFabricDto(name string, fabricType string, fabricConfiguration NetworkFabricFabricConfiguration, revision float32, id string, ) *UpdateNetworkFabricDto`

NewUpdateNetworkFabricDto instantiates a new UpdateNetworkFabricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFabricDtoWithDefaults

`func NewUpdateNetworkFabricDtoWithDefaults() *UpdateNetworkFabricDto`

NewUpdateNetworkFabricDtoWithDefaults instantiates a new UpdateNetworkFabricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkFabricDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkFabricDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkFabricDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateNetworkFabricDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkFabricDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkFabricDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkFabricDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFabricType

`func (o *UpdateNetworkFabricDto) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *UpdateNetworkFabricDto) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *UpdateNetworkFabricDto) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.


### GetFabricConfiguration

`func (o *UpdateNetworkFabricDto) GetFabricConfiguration() NetworkFabricFabricConfiguration`

GetFabricConfiguration returns the FabricConfiguration field if non-nil, zero value otherwise.

### GetFabricConfigurationOk

`func (o *UpdateNetworkFabricDto) GetFabricConfigurationOk() (*NetworkFabricFabricConfiguration, bool)`

GetFabricConfigurationOk returns a tuple with the FabricConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricConfiguration

`func (o *UpdateNetworkFabricDto) SetFabricConfiguration(v NetworkFabricFabricConfiguration)`

SetFabricConfiguration sets FabricConfiguration field to given value.


### GetRevision

`func (o *UpdateNetworkFabricDto) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UpdateNetworkFabricDto) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UpdateNetworkFabricDto) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetId

`func (o *UpdateNetworkFabricDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkFabricDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkFabricDto) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



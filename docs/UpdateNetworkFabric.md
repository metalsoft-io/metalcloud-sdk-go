# UpdateNetworkFabric

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

### NewUpdateNetworkFabric

`func NewUpdateNetworkFabric(name string, fabricType string, fabricConfiguration NetworkFabricFabricConfiguration, revision float32, id string, ) *UpdateNetworkFabric`

NewUpdateNetworkFabric instantiates a new UpdateNetworkFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFabricWithDefaults

`func NewUpdateNetworkFabricWithDefaults() *UpdateNetworkFabric`

NewUpdateNetworkFabricWithDefaults instantiates a new UpdateNetworkFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkFabric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkFabric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkFabric) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateNetworkFabric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkFabric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkFabric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkFabric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFabricType

`func (o *UpdateNetworkFabric) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *UpdateNetworkFabric) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *UpdateNetworkFabric) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.


### GetFabricConfiguration

`func (o *UpdateNetworkFabric) GetFabricConfiguration() NetworkFabricFabricConfiguration`

GetFabricConfiguration returns the FabricConfiguration field if non-nil, zero value otherwise.

### GetFabricConfigurationOk

`func (o *UpdateNetworkFabric) GetFabricConfigurationOk() (*NetworkFabricFabricConfiguration, bool)`

GetFabricConfigurationOk returns a tuple with the FabricConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricConfiguration

`func (o *UpdateNetworkFabric) SetFabricConfiguration(v NetworkFabricFabricConfiguration)`

SetFabricConfiguration sets FabricConfiguration field to given value.


### GetRevision

`func (o *UpdateNetworkFabric) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UpdateNetworkFabric) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UpdateNetworkFabric) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetId

`func (o *UpdateNetworkFabric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkFabric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkFabric) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



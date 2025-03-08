# CreateNetworkFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The network fabric name | 
**Description** | Pointer to **string** | Network fabric description | [optional] 
**FabricConfiguration** | [**NetworkFabricFabricConfiguration**](NetworkFabricFabricConfiguration.md) |  | 

## Methods

### NewCreateNetworkFabric

`func NewCreateNetworkFabric(name string, fabricConfiguration NetworkFabricFabricConfiguration, ) *CreateNetworkFabric`

NewCreateNetworkFabric instantiates a new CreateNetworkFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFabricWithDefaults

`func NewCreateNetworkFabricWithDefaults() *CreateNetworkFabric`

NewCreateNetworkFabricWithDefaults instantiates a new CreateNetworkFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkFabric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkFabric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkFabric) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNetworkFabric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkFabric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkFabric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkFabric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFabricConfiguration

`func (o *CreateNetworkFabric) GetFabricConfiguration() NetworkFabricFabricConfiguration`

GetFabricConfiguration returns the FabricConfiguration field if non-nil, zero value otherwise.

### GetFabricConfigurationOk

`func (o *CreateNetworkFabric) GetFabricConfigurationOk() (*NetworkFabricFabricConfiguration, bool)`

GetFabricConfigurationOk returns a tuple with the FabricConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricConfiguration

`func (o *CreateNetworkFabric) SetFabricConfiguration(v NetworkFabricFabricConfiguration)`

SetFabricConfiguration sets FabricConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



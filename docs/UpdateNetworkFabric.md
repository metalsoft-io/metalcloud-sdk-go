# UpdateNetworkFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The network fabric name | [optional] 
**Description** | Pointer to **string** | Network fabric description | [optional] 
**FabricConfiguration** | Pointer to [**NetworkFabricFabricConfiguration**](NetworkFabricFabricConfiguration.md) |  | [optional] 

## Methods

### NewUpdateNetworkFabric

`func NewUpdateNetworkFabric() *UpdateNetworkFabric`

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

### HasName

`func (o *UpdateNetworkFabric) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasFabricConfiguration

`func (o *UpdateNetworkFabric) HasFabricConfiguration() bool`

HasFabricConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



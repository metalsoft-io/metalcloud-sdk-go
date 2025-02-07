# UpdateExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The extension name | 
**Label** | Pointer to **string** | The extension unique label | [optional] 
**Description** | **string** | The extension description | 
**Definition** | [**ExtensionDefinition**](ExtensionDefinition.md) |  | 

## Methods

### NewUpdateExtension

`func NewUpdateExtension(name string, description string, definition ExtensionDefinition, ) *UpdateExtension`

NewUpdateExtension instantiates a new UpdateExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExtensionWithDefaults

`func NewUpdateExtensionWithDefaults() *UpdateExtension`

NewUpdateExtensionWithDefaults instantiates a new UpdateExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateExtension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExtension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExtension) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *UpdateExtension) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateExtension) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateExtension) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateExtension) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateExtension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateExtension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateExtension) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefinition

`func (o *UpdateExtension) GetDefinition() ExtensionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *UpdateExtension) GetDefinitionOk() (*ExtensionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *UpdateExtension) SetDefinition(v ExtensionDefinition)`

SetDefinition sets Definition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



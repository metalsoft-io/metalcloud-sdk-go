# CreateExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The extension name | 
**Label** | Pointer to **string** | The extension unique label | [optional] 
**Description** | **string** | The extension description | 
**Definition** | [**ExtensionDefinition**](ExtensionDefinition.md) |  | 

## Methods

### NewCreateExtension

`func NewCreateExtension(name string, description string, definition ExtensionDefinition, ) *CreateExtension`

NewCreateExtension instantiates a new CreateExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExtensionWithDefaults

`func NewCreateExtensionWithDefaults() *CreateExtension`

NewCreateExtensionWithDefaults instantiates a new CreateExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateExtension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExtension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExtension) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CreateExtension) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateExtension) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateExtension) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateExtension) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *CreateExtension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateExtension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateExtension) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefinition

`func (o *CreateExtension) GetDefinition() ExtensionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateExtension) GetDefinitionOk() (*ExtensionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateExtension) SetDefinition(v ExtensionDefinition)`

SetDefinition sets Definition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



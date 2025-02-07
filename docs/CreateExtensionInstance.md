# CreateExtensionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | Pointer to **float32** | The extension ID. | [optional] 
**Label** | Pointer to **string** | The extension instance label. Will be automatically generated if not provided. | [optional] 
**InputVariables** | Pointer to [**[]ExtensionVariable**](ExtensionVariable.md) | Input variables values. | [optional] 

## Methods

### NewCreateExtensionInstance

`func NewCreateExtensionInstance() *CreateExtensionInstance`

NewCreateExtensionInstance instantiates a new CreateExtensionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExtensionInstanceWithDefaults

`func NewCreateExtensionInstanceWithDefaults() *CreateExtensionInstance`

NewCreateExtensionInstanceWithDefaults instantiates a new CreateExtensionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionId

`func (o *CreateExtensionInstance) GetExtensionId() float32`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *CreateExtensionInstance) GetExtensionIdOk() (*float32, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *CreateExtensionInstance) SetExtensionId(v float32)`

SetExtensionId sets ExtensionId field to given value.

### HasExtensionId

`func (o *CreateExtensionInstance) HasExtensionId() bool`

HasExtensionId returns a boolean if a field has been set.

### GetLabel

`func (o *CreateExtensionInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateExtensionInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateExtensionInstance) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateExtensionInstance) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetInputVariables

`func (o *CreateExtensionInstance) GetInputVariables() []ExtensionVariable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *CreateExtensionInstance) GetInputVariablesOk() (*[]ExtensionVariable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *CreateExtensionInstance) SetInputVariables(v []ExtensionVariable)`

SetInputVariables sets InputVariables field to given value.

### HasInputVariables

`func (o *CreateExtensionInstance) HasInputVariables() bool`

HasInputVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExtensionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The extension instance ID. | 
**Revision** | **float32** | Revision number | 
**InfrastructureId** | **float32** | The infrastructure ID. | 
**ExtensionId** | **float32** | The extension ID. | 
**Label** | **string** | The extension instance label. Will be automatically generated if not provided. | 
**InputVariables** | [**[]ExtensionVariable**](ExtensionVariable.md) | Input variables values. | 
**OutputVariables** | [**[]ExtensionVariable**](ExtensionVariable.md) | Output variables values. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExtensionInstance

`func NewExtensionInstance(id float32, revision float32, infrastructureId float32, extensionId float32, label string, inputVariables []ExtensionVariable, outputVariables []ExtensionVariable, ) *ExtensionInstance`

NewExtensionInstance instantiates a new ExtensionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstanceWithDefaults

`func NewExtensionInstanceWithDefaults() *ExtensionInstance`

NewExtensionInstanceWithDefaults instantiates a new ExtensionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtensionInstance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtensionInstance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtensionInstance) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ExtensionInstance) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ExtensionInstance) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ExtensionInstance) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *ExtensionInstance) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ExtensionInstance) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ExtensionInstance) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetExtensionId

`func (o *ExtensionInstance) GetExtensionId() float32`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *ExtensionInstance) GetExtensionIdOk() (*float32, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *ExtensionInstance) SetExtensionId(v float32)`

SetExtensionId sets ExtensionId field to given value.


### GetLabel

`func (o *ExtensionInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInputVariables

`func (o *ExtensionInstance) GetInputVariables() []ExtensionVariable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *ExtensionInstance) GetInputVariablesOk() (*[]ExtensionVariable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *ExtensionInstance) SetInputVariables(v []ExtensionVariable)`

SetInputVariables sets InputVariables field to given value.


### GetOutputVariables

`func (o *ExtensionInstance) GetOutputVariables() []ExtensionVariable`

GetOutputVariables returns the OutputVariables field if non-nil, zero value otherwise.

### GetOutputVariablesOk

`func (o *ExtensionInstance) GetOutputVariablesOk() (*[]ExtensionVariable, bool)`

GetOutputVariablesOk returns a tuple with the OutputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputVariables

`func (o *ExtensionInstance) SetOutputVariables(v []ExtensionVariable)`

SetOutputVariables sets OutputVariables field to given value.


### GetLinks

`func (o *ExtensionInstance) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtensionInstance) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtensionInstance) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExtensionInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



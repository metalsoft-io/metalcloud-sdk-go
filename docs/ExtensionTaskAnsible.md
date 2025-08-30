# ExtensionTaskAnsible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the task. | 
**TaskType** | [**ExtensionTaskType**](ExtensionTaskType.md) |  | 
**Options** | [**ExtensionTaskOptionAnsible**](ExtensionTaskOptionAnsible.md) |  | 

## Methods

### NewExtensionTaskAnsible

`func NewExtensionTaskAnsible(label string, taskType ExtensionTaskType, options ExtensionTaskOptionAnsible, ) *ExtensionTaskAnsible`

NewExtensionTaskAnsible instantiates a new ExtensionTaskAnsible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskAnsibleWithDefaults

`func NewExtensionTaskAnsibleWithDefaults() *ExtensionTaskAnsible`

NewExtensionTaskAnsibleWithDefaults instantiates a new ExtensionTaskAnsible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionTaskAnsible) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionTaskAnsible) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionTaskAnsible) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTaskType

`func (o *ExtensionTaskAnsible) GetTaskType() ExtensionTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ExtensionTaskAnsible) GetTaskTypeOk() (*ExtensionTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ExtensionTaskAnsible) SetTaskType(v ExtensionTaskType)`

SetTaskType sets TaskType field to given value.


### GetOptions

`func (o *ExtensionTaskAnsible) GetOptions() ExtensionTaskOptionAnsible`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionTaskAnsible) GetOptionsOk() (*ExtensionTaskOptionAnsible, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionTaskAnsible) SetOptions(v ExtensionTaskOptionAnsible)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



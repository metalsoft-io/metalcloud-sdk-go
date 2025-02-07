# ExtensionTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the task. | 
**TaskType** | **string** | Type of the task. | 
**Options** | [**ExtensionTaskOptions**](ExtensionTaskOptions.md) |  | 

## Methods

### NewExtensionTask

`func NewExtensionTask(label string, taskType string, options ExtensionTaskOptions, ) *ExtensionTask`

NewExtensionTask instantiates a new ExtensionTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskWithDefaults

`func NewExtensionTaskWithDefaults() *ExtensionTask`

NewExtensionTaskWithDefaults instantiates a new ExtensionTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionTask) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionTask) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionTask) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTaskType

`func (o *ExtensionTask) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ExtensionTask) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ExtensionTask) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.


### GetOptions

`func (o *ExtensionTask) GetOptions() ExtensionTaskOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionTask) GetOptionsOk() (*ExtensionTaskOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionTask) SetOptions(v ExtensionTaskOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



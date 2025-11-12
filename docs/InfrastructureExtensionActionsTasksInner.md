# InfrastructureExtensionActionsTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the task. | 
**TaskType** | [**ExtensionTaskType**](ExtensionTaskType.md) |  | 
**Options** | [**ExtensionTaskOptionSsh**](ExtensionTaskOptionSsh.md) |  | 

## Methods

### NewInfrastructureExtensionActionsTasksInner

`func NewInfrastructureExtensionActionsTasksInner(label string, taskType ExtensionTaskType, options ExtensionTaskOptionSsh, ) *InfrastructureExtensionActionsTasksInner`

NewInfrastructureExtensionActionsTasksInner instantiates a new InfrastructureExtensionActionsTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureExtensionActionsTasksInnerWithDefaults

`func NewInfrastructureExtensionActionsTasksInnerWithDefaults() *InfrastructureExtensionActionsTasksInner`

NewInfrastructureExtensionActionsTasksInnerWithDefaults instantiates a new InfrastructureExtensionActionsTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InfrastructureExtensionActionsTasksInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfrastructureExtensionActionsTasksInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfrastructureExtensionActionsTasksInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTaskType

`func (o *InfrastructureExtensionActionsTasksInner) GetTaskType() ExtensionTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *InfrastructureExtensionActionsTasksInner) GetTaskTypeOk() (*ExtensionTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *InfrastructureExtensionActionsTasksInner) SetTaskType(v ExtensionTaskType)`

SetTaskType sets TaskType field to given value.


### GetOptions

`func (o *InfrastructureExtensionActionsTasksInner) GetOptions() ExtensionTaskOptionSsh`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InfrastructureExtensionActionsTasksInner) GetOptionsOk() (*ExtensionTaskOptionSsh, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InfrastructureExtensionActionsTasksInner) SetOptions(v ExtensionTaskOptionSsh)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



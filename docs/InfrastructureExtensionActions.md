# InfrastructureExtensionActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | [**InfrastructureExtensionActionStage**](InfrastructureExtensionActionStage.md) |  | 
**Tasks** | Pointer to [**[]InfrastructureExtensionActionsTasksDataItem**](InfrastructureExtensionActionsTasksDataItem.md) | Tasks. | [optional] 

## Methods

### NewInfrastructureExtensionActions

`func NewInfrastructureExtensionActions(stage InfrastructureExtensionActionStage, ) *InfrastructureExtensionActions`

NewInfrastructureExtensionActions instantiates a new InfrastructureExtensionActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureExtensionActionsWithDefaults

`func NewInfrastructureExtensionActionsWithDefaults() *InfrastructureExtensionActions`

NewInfrastructureExtensionActionsWithDefaults instantiates a new InfrastructureExtensionActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *InfrastructureExtensionActions) GetStage() InfrastructureExtensionActionStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *InfrastructureExtensionActions) GetStageOk() (*InfrastructureExtensionActionStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *InfrastructureExtensionActions) SetStage(v InfrastructureExtensionActionStage)`

SetStage sets Stage field to given value.


### GetTasks

`func (o *InfrastructureExtensionActions) GetTasks() []InfrastructureExtensionActionsTasksDataItem`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *InfrastructureExtensionActions) GetTasksOk() (*[]InfrastructureExtensionActionsTasksDataItem, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *InfrastructureExtensionActions) SetTasks(v []InfrastructureExtensionActionsTasksDataItem)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *InfrastructureExtensionActions) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



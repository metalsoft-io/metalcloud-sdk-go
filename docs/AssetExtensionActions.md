# AssetExtensionActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | [**AssetExtensionActionStage**](AssetExtensionActionStage.md) |  | 
**Tasks** | Pointer to [**[]InfrastructureExtensionActionsTasksInner**](InfrastructureExtensionActionsTasksInner.md) | Tasks. | [optional] 

## Methods

### NewAssetExtensionActions

`func NewAssetExtensionActions(stage AssetExtensionActionStage, ) *AssetExtensionActions`

NewAssetExtensionActions instantiates a new AssetExtensionActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetExtensionActionsWithDefaults

`func NewAssetExtensionActionsWithDefaults() *AssetExtensionActions`

NewAssetExtensionActionsWithDefaults instantiates a new AssetExtensionActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *AssetExtensionActions) GetStage() AssetExtensionActionStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *AssetExtensionActions) GetStageOk() (*AssetExtensionActionStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *AssetExtensionActions) SetStage(v AssetExtensionActionStage)`

SetStage sets Stage field to given value.


### GetTasks

`func (o *AssetExtensionActions) GetTasks() []InfrastructureExtensionActionsTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AssetExtensionActions) GetTasksOk() (*[]InfrastructureExtensionActionsTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AssetExtensionActions) SetTasks(v []InfrastructureExtensionActionsTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AssetExtensionActions) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



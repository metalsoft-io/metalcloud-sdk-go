# ExtensionTaskWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the task. | 
**TaskType** | [**ExtensionTaskType**](ExtensionTaskType.md) |  | 
**Options** | [**ExtensionTaskOptionWebhook**](ExtensionTaskOptionWebhook.md) |  | 

## Methods

### NewExtensionTaskWebhook

`func NewExtensionTaskWebhook(label string, taskType ExtensionTaskType, options ExtensionTaskOptionWebhook, ) *ExtensionTaskWebhook`

NewExtensionTaskWebhook instantiates a new ExtensionTaskWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskWebhookWithDefaults

`func NewExtensionTaskWebhookWithDefaults() *ExtensionTaskWebhook`

NewExtensionTaskWebhookWithDefaults instantiates a new ExtensionTaskWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionTaskWebhook) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionTaskWebhook) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionTaskWebhook) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTaskType

`func (o *ExtensionTaskWebhook) GetTaskType() ExtensionTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ExtensionTaskWebhook) GetTaskTypeOk() (*ExtensionTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ExtensionTaskWebhook) SetTaskType(v ExtensionTaskType)`

SetTaskType sets TaskType field to given value.


### GetOptions

`func (o *ExtensionTaskWebhook) GetOptions() ExtensionTaskOptionWebhook`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtensionTaskWebhook) GetOptionsOk() (*ExtensionTaskOptionWebhook, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtensionTaskWebhook) SetOptions(v ExtensionTaskOptionWebhook)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExtensionTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** | Asset for the Ansible task. | [optional] 
**Playbook** | Pointer to **string** | Playbook for the Ansible task. | [optional] 
**ExecutionTimeout** | Pointer to **int32** | Execution Timeout. | [optional] 
**ExecutionTimeoutTick** | Pointer to **int32** | Execution Timeout Tick. | [optional] 
**Version** | Pointer to **string** | Version. | [optional] 
**Endpoint** | Pointer to **string** | Webhook task endpoint. | [optional] 
**RequestTemplate** | Pointer to **string** | Webhook task request body template. | [optional] 

## Methods

### NewExtensionTaskOptions

`func NewExtensionTaskOptions() *ExtensionTaskOptions`

NewExtensionTaskOptions instantiates a new ExtensionTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskOptionsWithDefaults

`func NewExtensionTaskOptionsWithDefaults() *ExtensionTaskOptions`

NewExtensionTaskOptionsWithDefaults instantiates a new ExtensionTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *ExtensionTaskOptions) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *ExtensionTaskOptions) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *ExtensionTaskOptions) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *ExtensionTaskOptions) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetPlaybook

`func (o *ExtensionTaskOptions) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *ExtensionTaskOptions) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *ExtensionTaskOptions) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *ExtensionTaskOptions) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetExecutionTimeout

`func (o *ExtensionTaskOptions) GetExecutionTimeout() int32`

GetExecutionTimeout returns the ExecutionTimeout field if non-nil, zero value otherwise.

### GetExecutionTimeoutOk

`func (o *ExtensionTaskOptions) GetExecutionTimeoutOk() (*int32, bool)`

GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeout

`func (o *ExtensionTaskOptions) SetExecutionTimeout(v int32)`

SetExecutionTimeout sets ExecutionTimeout field to given value.

### HasExecutionTimeout

`func (o *ExtensionTaskOptions) HasExecutionTimeout() bool`

HasExecutionTimeout returns a boolean if a field has been set.

### GetExecutionTimeoutTick

`func (o *ExtensionTaskOptions) GetExecutionTimeoutTick() int32`

GetExecutionTimeoutTick returns the ExecutionTimeoutTick field if non-nil, zero value otherwise.

### GetExecutionTimeoutTickOk

`func (o *ExtensionTaskOptions) GetExecutionTimeoutTickOk() (*int32, bool)`

GetExecutionTimeoutTickOk returns a tuple with the ExecutionTimeoutTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeoutTick

`func (o *ExtensionTaskOptions) SetExecutionTimeoutTick(v int32)`

SetExecutionTimeoutTick sets ExecutionTimeoutTick field to given value.

### HasExecutionTimeoutTick

`func (o *ExtensionTaskOptions) HasExecutionTimeoutTick() bool`

HasExecutionTimeoutTick returns a boolean if a field has been set.

### GetVersion

`func (o *ExtensionTaskOptions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionTaskOptions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionTaskOptions) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionTaskOptions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEndpoint

`func (o *ExtensionTaskOptions) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ExtensionTaskOptions) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ExtensionTaskOptions) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ExtensionTaskOptions) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetRequestTemplate

`func (o *ExtensionTaskOptions) GetRequestTemplate() string`

GetRequestTemplate returns the RequestTemplate field if non-nil, zero value otherwise.

### GetRequestTemplateOk

`func (o *ExtensionTaskOptions) GetRequestTemplateOk() (*string, bool)`

GetRequestTemplateOk returns a tuple with the RequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplate

`func (o *ExtensionTaskOptions) SetRequestTemplate(v string)`

SetRequestTemplate sets RequestTemplate field to given value.

### HasRequestTemplate

`func (o *ExtensionTaskOptions) HasRequestTemplate() bool`

HasRequestTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



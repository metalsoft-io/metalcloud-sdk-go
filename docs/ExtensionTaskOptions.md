# ExtensionTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** | Asset for the Ansible task. | [optional] 
**Playbook** | Pointer to **string** | Playbook for the Ansible task. | [optional] 
**ExecutionTimeout** | Pointer to **int32** | Execution Timeout. | [optional] 
**ExecutionTimeoutTick** | Pointer to **int32** | Execution Timeout Tick. | [optional] 
**Version** | Pointer to **string** | Version. | [optional] 
**Endpoint** | **string** | Webhook task endpoint. | 
**Method** | **string** | Webhook task request method. | 
**Headers** | Pointer to **map[string]interface{}** | Request headers for the webhook task. | [optional] 
**RequestTemplate** | **string** | Request template for the webhook task. | 
**ExpectedResponseStatuses** | Pointer to **[]float32** | Expected response statuses for the webhook task. | [optional] 
**Host** | **string** | Host to execute the SSH command on. | 
**Port** | **int32** | Port to connect to the host via SSH. | 
**Username** | Pointer to **string** | Username for SSH connection. | [optional] 
**Password** | Pointer to **string** | Password for SSH connection. | [optional] 
**Timeout** | **int32** | Timeout for the SSH command execution in seconds. | 
**CommandTemplate** | **string** | Command template to execute via SSH. | 

## Methods

### NewExtensionTaskOptions

`func NewExtensionTaskOptions(endpoint string, method string, requestTemplate string, host string, port int32, timeout int32, commandTemplate string, ) *ExtensionTaskOptions`

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


### GetMethod

`func (o *ExtensionTaskOptions) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ExtensionTaskOptions) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ExtensionTaskOptions) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetHeaders

`func (o *ExtensionTaskOptions) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ExtensionTaskOptions) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ExtensionTaskOptions) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ExtensionTaskOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

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


### GetExpectedResponseStatuses

`func (o *ExtensionTaskOptions) GetExpectedResponseStatuses() []float32`

GetExpectedResponseStatuses returns the ExpectedResponseStatuses field if non-nil, zero value otherwise.

### GetExpectedResponseStatusesOk

`func (o *ExtensionTaskOptions) GetExpectedResponseStatusesOk() (*[]float32, bool)`

GetExpectedResponseStatusesOk returns a tuple with the ExpectedResponseStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponseStatuses

`func (o *ExtensionTaskOptions) SetExpectedResponseStatuses(v []float32)`

SetExpectedResponseStatuses sets ExpectedResponseStatuses field to given value.

### HasExpectedResponseStatuses

`func (o *ExtensionTaskOptions) HasExpectedResponseStatuses() bool`

HasExpectedResponseStatuses returns a boolean if a field has been set.

### GetHost

`func (o *ExtensionTaskOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ExtensionTaskOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ExtensionTaskOptions) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ExtensionTaskOptions) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ExtensionTaskOptions) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ExtensionTaskOptions) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *ExtensionTaskOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ExtensionTaskOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ExtensionTaskOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ExtensionTaskOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ExtensionTaskOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExtensionTaskOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExtensionTaskOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ExtensionTaskOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTimeout

`func (o *ExtensionTaskOptions) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ExtensionTaskOptions) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ExtensionTaskOptions) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetCommandTemplate

`func (o *ExtensionTaskOptions) GetCommandTemplate() string`

GetCommandTemplate returns the CommandTemplate field if non-nil, zero value otherwise.

### GetCommandTemplateOk

`func (o *ExtensionTaskOptions) GetCommandTemplateOk() (*string, bool)`

GetCommandTemplateOk returns a tuple with the CommandTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandTemplate

`func (o *ExtensionTaskOptions) SetCommandTemplate(v string)`

SetCommandTemplate sets CommandTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



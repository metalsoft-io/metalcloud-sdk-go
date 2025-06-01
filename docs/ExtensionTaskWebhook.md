# ExtensionTaskWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Webhook task endpoint. | 
**Method** | **string** | Webhook task request method. | 
**Headers** | Pointer to **map[string]interface{}** | Request headers for the webhook task. | [optional] 
**RequestTemplate** | **string** | Request template for the webhook task. | 
**ExpectedResponseStatuses** | Pointer to **[]float32** | Expected response statuses for the webhook task. | [optional] 

## Methods

### NewExtensionTaskWebhook

`func NewExtensionTaskWebhook(endpoint string, method string, requestTemplate string, ) *ExtensionTaskWebhook`

NewExtensionTaskWebhook instantiates a new ExtensionTaskWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskWebhookWithDefaults

`func NewExtensionTaskWebhookWithDefaults() *ExtensionTaskWebhook`

NewExtensionTaskWebhookWithDefaults instantiates a new ExtensionTaskWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ExtensionTaskWebhook) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ExtensionTaskWebhook) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ExtensionTaskWebhook) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetMethod

`func (o *ExtensionTaskWebhook) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ExtensionTaskWebhook) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ExtensionTaskWebhook) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetHeaders

`func (o *ExtensionTaskWebhook) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ExtensionTaskWebhook) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ExtensionTaskWebhook) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ExtensionTaskWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRequestTemplate

`func (o *ExtensionTaskWebhook) GetRequestTemplate() string`

GetRequestTemplate returns the RequestTemplate field if non-nil, zero value otherwise.

### GetRequestTemplateOk

`func (o *ExtensionTaskWebhook) GetRequestTemplateOk() (*string, bool)`

GetRequestTemplateOk returns a tuple with the RequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplate

`func (o *ExtensionTaskWebhook) SetRequestTemplate(v string)`

SetRequestTemplate sets RequestTemplate field to given value.


### GetExpectedResponseStatuses

`func (o *ExtensionTaskWebhook) GetExpectedResponseStatuses() []float32`

GetExpectedResponseStatuses returns the ExpectedResponseStatuses field if non-nil, zero value otherwise.

### GetExpectedResponseStatusesOk

`func (o *ExtensionTaskWebhook) GetExpectedResponseStatusesOk() (*[]float32, bool)`

GetExpectedResponseStatusesOk returns a tuple with the ExpectedResponseStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponseStatuses

`func (o *ExtensionTaskWebhook) SetExpectedResponseStatuses(v []float32)`

SetExpectedResponseStatuses sets ExpectedResponseStatuses field to given value.

### HasExpectedResponseStatuses

`func (o *ExtensionTaskWebhook) HasExpectedResponseStatuses() bool`

HasExpectedResponseStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



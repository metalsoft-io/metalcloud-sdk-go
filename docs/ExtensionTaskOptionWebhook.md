# ExtensionTaskOptionWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Webhook task endpoint. | 
**Method** | **string** | Webhook task request method. | 
**Headers** | Pointer to **map[string]interface{}** | Request headers for the webhook task. | [optional] 
**RequestTemplate** | **string** | Request template for the webhook task. | 
**ExpectedResponseStatuses** | Pointer to **[]float32** | Expected response statuses for the webhook task. | [optional] 
**Timeout** | Pointer to **int32** | Timeout for the webhook task in seconds. | [optional] [default to 30]
**InsecureSkipVerify** | Pointer to **bool** | Flag to indicate if certificates should be verified for the webhook task. | [optional] [default to true]

## Methods

### NewExtensionTaskOptionWebhook

`func NewExtensionTaskOptionWebhook(endpoint string, method string, requestTemplate string, ) *ExtensionTaskOptionWebhook`

NewExtensionTaskOptionWebhook instantiates a new ExtensionTaskOptionWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionTaskOptionWebhookWithDefaults

`func NewExtensionTaskOptionWebhookWithDefaults() *ExtensionTaskOptionWebhook`

NewExtensionTaskOptionWebhookWithDefaults instantiates a new ExtensionTaskOptionWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ExtensionTaskOptionWebhook) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ExtensionTaskOptionWebhook) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ExtensionTaskOptionWebhook) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetMethod

`func (o *ExtensionTaskOptionWebhook) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ExtensionTaskOptionWebhook) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ExtensionTaskOptionWebhook) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetHeaders

`func (o *ExtensionTaskOptionWebhook) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ExtensionTaskOptionWebhook) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ExtensionTaskOptionWebhook) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ExtensionTaskOptionWebhook) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRequestTemplate

`func (o *ExtensionTaskOptionWebhook) GetRequestTemplate() string`

GetRequestTemplate returns the RequestTemplate field if non-nil, zero value otherwise.

### GetRequestTemplateOk

`func (o *ExtensionTaskOptionWebhook) GetRequestTemplateOk() (*string, bool)`

GetRequestTemplateOk returns a tuple with the RequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplate

`func (o *ExtensionTaskOptionWebhook) SetRequestTemplate(v string)`

SetRequestTemplate sets RequestTemplate field to given value.


### GetExpectedResponseStatuses

`func (o *ExtensionTaskOptionWebhook) GetExpectedResponseStatuses() []float32`

GetExpectedResponseStatuses returns the ExpectedResponseStatuses field if non-nil, zero value otherwise.

### GetExpectedResponseStatusesOk

`func (o *ExtensionTaskOptionWebhook) GetExpectedResponseStatusesOk() (*[]float32, bool)`

GetExpectedResponseStatusesOk returns a tuple with the ExpectedResponseStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponseStatuses

`func (o *ExtensionTaskOptionWebhook) SetExpectedResponseStatuses(v []float32)`

SetExpectedResponseStatuses sets ExpectedResponseStatuses field to given value.

### HasExpectedResponseStatuses

`func (o *ExtensionTaskOptionWebhook) HasExpectedResponseStatuses() bool`

HasExpectedResponseStatuses returns a boolean if a field has been set.

### GetTimeout

`func (o *ExtensionTaskOptionWebhook) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ExtensionTaskOptionWebhook) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ExtensionTaskOptionWebhook) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ExtensionTaskOptionWebhook) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetInsecureSkipVerify

`func (o *ExtensionTaskOptionWebhook) GetInsecureSkipVerify() bool`

GetInsecureSkipVerify returns the InsecureSkipVerify field if non-nil, zero value otherwise.

### GetInsecureSkipVerifyOk

`func (o *ExtensionTaskOptionWebhook) GetInsecureSkipVerifyOk() (*bool, bool)`

GetInsecureSkipVerifyOk returns a tuple with the InsecureSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipVerify

`func (o *ExtensionTaskOptionWebhook) SetInsecureSkipVerify(v bool)`

SetInsecureSkipVerify sets InsecureSkipVerify field to given value.

### HasInsecureSkipVerify

`func (o *ExtensionTaskOptionWebhook) HasInsecureSkipVerify() bool`

HasInsecureSkipVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



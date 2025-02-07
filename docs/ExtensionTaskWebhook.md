# ExtensionTaskWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | Webhook task endpoint. | [optional] 
**RequestTemplate** | Pointer to **string** | Webhook task request body template. | [optional] 

## Methods

### NewExtensionTaskWebhook

`func NewExtensionTaskWebhook() *ExtensionTaskWebhook`

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

### HasEndpoint

`func (o *ExtensionTaskWebhook) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

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

### HasRequestTemplate

`func (o *ExtensionTaskWebhook) HasRequestTemplate() bool`

HasRequestTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



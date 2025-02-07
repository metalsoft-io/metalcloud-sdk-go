# AIGenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacenter** | **string** | Datacenter that is being filtered | 
**Prompt** | **string** | String input sent to the AI | 

## Methods

### NewAIGenerateRequest

`func NewAIGenerateRequest(datacenter string, prompt string, ) *AIGenerateRequest`

NewAIGenerateRequest instantiates a new AIGenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIGenerateRequestWithDefaults

`func NewAIGenerateRequestWithDefaults() *AIGenerateRequest`

NewAIGenerateRequestWithDefaults instantiates a new AIGenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenter

`func (o *AIGenerateRequest) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *AIGenerateRequest) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *AIGenerateRequest) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetPrompt

`func (o *AIGenerateRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AIGenerateRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AIGenerateRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



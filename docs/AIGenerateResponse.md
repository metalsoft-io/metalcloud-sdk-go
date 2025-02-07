# AIGenerateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Response to the question asked by the user | 
**Steps** | **string** | Steps to be taken to achieve the result | 

## Methods

### NewAIGenerateResponse

`func NewAIGenerateResponse(result string, steps string, ) *AIGenerateResponse`

NewAIGenerateResponse instantiates a new AIGenerateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIGenerateResponseWithDefaults

`func NewAIGenerateResponseWithDefaults() *AIGenerateResponse`

NewAIGenerateResponseWithDefaults instantiates a new AIGenerateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AIGenerateResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AIGenerateResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AIGenerateResponse) SetResult(v string)`

SetResult sets Result field to given value.


### GetSteps

`func (o *AIGenerateResponse) GetSteps() string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AIGenerateResponse) GetStepsOk() (*string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AIGenerateResponse) SetSteps(v string)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



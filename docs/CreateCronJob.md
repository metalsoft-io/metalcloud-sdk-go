# CreateCronJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FunctionName** | **string** |  | 
**Params** | **[]map[string]interface{}** |  | 
**Schedule** | **string** |  | 
**WaitForCompletion** | **float32** |  | 
**LifetimeSeconds** | **float32** |  | 
**Disabled** | **float32** |  | 

## Methods

### NewCreateCronJob

`func NewCreateCronJob(label string, functionName string, params []map[string]interface{}, schedule string, waitForCompletion float32, lifetimeSeconds float32, disabled float32, ) *CreateCronJob`

NewCreateCronJob instantiates a new CreateCronJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCronJobWithDefaults

`func NewCreateCronJobWithDefaults() *CreateCronJob`

NewCreateCronJobWithDefaults instantiates a new CreateCronJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateCronJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateCronJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateCronJob) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CreateCronJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCronJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCronJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCronJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionName

`func (o *CreateCronJob) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *CreateCronJob) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *CreateCronJob) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetParams

`func (o *CreateCronJob) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateCronJob) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateCronJob) SetParams(v []map[string]interface{})`

SetParams sets Params field to given value.


### GetSchedule

`func (o *CreateCronJob) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateCronJob) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateCronJob) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetWaitForCompletion

`func (o *CreateCronJob) GetWaitForCompletion() float32`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *CreateCronJob) GetWaitForCompletionOk() (*float32, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *CreateCronJob) SetWaitForCompletion(v float32)`

SetWaitForCompletion sets WaitForCompletion field to given value.


### GetLifetimeSeconds

`func (o *CreateCronJob) GetLifetimeSeconds() float32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *CreateCronJob) GetLifetimeSecondsOk() (*float32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *CreateCronJob) SetLifetimeSeconds(v float32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.


### GetDisabled

`func (o *CreateCronJob) GetDisabled() float32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateCronJob) GetDisabledOk() (*float32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateCronJob) SetDisabled(v float32)`

SetDisabled sets Disabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



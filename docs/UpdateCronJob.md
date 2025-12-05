# UpdateCronJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FunctionName** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**WaitForCompletion** | Pointer to **float32** |  | [optional] 
**LifetimeSeconds** | Pointer to **float32** |  | [optional] 
**Disabled** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateCronJob

`func NewUpdateCronJob() *UpdateCronJob`

NewUpdateCronJob instantiates a new UpdateCronJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCronJobWithDefaults

`func NewUpdateCronJobWithDefaults() *UpdateCronJob`

NewUpdateCronJobWithDefaults instantiates a new UpdateCronJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateCronJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateCronJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateCronJob) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateCronJob) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCronJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCronJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCronJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCronJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionName

`func (o *UpdateCronJob) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *UpdateCronJob) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *UpdateCronJob) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *UpdateCronJob) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetParams

`func (o *UpdateCronJob) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UpdateCronJob) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *UpdateCronJob) SetParams(v []map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *UpdateCronJob) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateCronJob) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateCronJob) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateCronJob) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateCronJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetWaitForCompletion

`func (o *UpdateCronJob) GetWaitForCompletion() float32`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *UpdateCronJob) GetWaitForCompletionOk() (*float32, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *UpdateCronJob) SetWaitForCompletion(v float32)`

SetWaitForCompletion sets WaitForCompletion field to given value.

### HasWaitForCompletion

`func (o *UpdateCronJob) HasWaitForCompletion() bool`

HasWaitForCompletion returns a boolean if a field has been set.

### GetLifetimeSeconds

`func (o *UpdateCronJob) GetLifetimeSeconds() float32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *UpdateCronJob) GetLifetimeSecondsOk() (*float32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *UpdateCronJob) SetLifetimeSeconds(v float32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.

### HasLifetimeSeconds

`func (o *UpdateCronJob) HasLifetimeSeconds() bool`

HasLifetimeSeconds returns a boolean if a field has been set.

### GetDisabled

`func (o *UpdateCronJob) GetDisabled() float32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UpdateCronJob) GetDisabledOk() (*float32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UpdateCronJob) SetDisabled(v float32)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UpdateCronJob) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



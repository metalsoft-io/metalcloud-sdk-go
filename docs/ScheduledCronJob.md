# ScheduledCronJob

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

### NewScheduledCronJob

`func NewScheduledCronJob(label string, functionName string, params []map[string]interface{}, schedule string, waitForCompletion float32, lifetimeSeconds float32, disabled float32, ) *ScheduledCronJob`

NewScheduledCronJob instantiates a new ScheduledCronJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledCronJobWithDefaults

`func NewScheduledCronJobWithDefaults() *ScheduledCronJob`

NewScheduledCronJobWithDefaults instantiates a new ScheduledCronJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ScheduledCronJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScheduledCronJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScheduledCronJob) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ScheduledCronJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledCronJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledCronJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledCronJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionName

`func (o *ScheduledCronJob) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ScheduledCronJob) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ScheduledCronJob) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetParams

`func (o *ScheduledCronJob) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScheduledCronJob) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScheduledCronJob) SetParams(v []map[string]interface{})`

SetParams sets Params field to given value.


### GetSchedule

`func (o *ScheduledCronJob) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduledCronJob) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduledCronJob) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetWaitForCompletion

`func (o *ScheduledCronJob) GetWaitForCompletion() float32`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *ScheduledCronJob) GetWaitForCompletionOk() (*float32, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *ScheduledCronJob) SetWaitForCompletion(v float32)`

SetWaitForCompletion sets WaitForCompletion field to given value.


### GetLifetimeSeconds

`func (o *ScheduledCronJob) GetLifetimeSeconds() float32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *ScheduledCronJob) GetLifetimeSecondsOk() (*float32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *ScheduledCronJob) SetLifetimeSeconds(v float32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.


### GetDisabled

`func (o *ScheduledCronJob) GetDisabled() float32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ScheduledCronJob) GetDisabledOk() (*float32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ScheduledCronJob) SetDisabled(v float32)`

SetDisabled sets Disabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



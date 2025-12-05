# CronJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FunctionName** | **string** |  | 
**Params** | **[]map[string]interface{}** |  | 
**Schedule** | **string** |  | 
**WaitForCompletion** | **float32** |  | 
**LifetimeSeconds** | **float32** |  | 
**Disabled** | **float32** |  | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewCronJob

`func NewCronJob(id float32, label string, functionName string, params []map[string]interface{}, schedule string, waitForCompletion float32, lifetimeSeconds float32, disabled float32, links map[string]interface{}, ) *CronJob`

NewCronJob instantiates a new CronJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobWithDefaults

`func NewCronJobWithDefaults() *CronJob`

NewCronJobWithDefaults instantiates a new CronJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CronJob) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CronJob) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CronJob) SetId(v float32)`

SetId sets Id field to given value.


### GetLabel

`func (o *CronJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CronJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CronJob) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CronJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CronJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CronJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CronJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFunctionName

`func (o *CronJob) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *CronJob) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *CronJob) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetParams

`func (o *CronJob) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CronJob) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CronJob) SetParams(v []map[string]interface{})`

SetParams sets Params field to given value.


### GetSchedule

`func (o *CronJob) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJob) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJob) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetWaitForCompletion

`func (o *CronJob) GetWaitForCompletion() float32`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *CronJob) GetWaitForCompletionOk() (*float32, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *CronJob) SetWaitForCompletion(v float32)`

SetWaitForCompletion sets WaitForCompletion field to given value.


### GetLifetimeSeconds

`func (o *CronJob) GetLifetimeSeconds() float32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *CronJob) GetLifetimeSecondsOk() (*float32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *CronJob) SetLifetimeSeconds(v float32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.


### GetDisabled

`func (o *CronJob) GetDisabled() float32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CronJob) GetDisabledOk() (*float32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CronJob) SetDisabled(v float32)`

SetDisabled sets Disabled field to given value.


### GetLinks

`func (o *CronJob) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CronJob) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CronJob) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



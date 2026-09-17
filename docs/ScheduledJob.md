# ScheduledJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SystemDefined** | **bool** | Wether this cron is system defined and cannot be deleted | 
**FunctionName** | **string** |  | 
**Params** | **[]map[string]interface{}** |  | 
**Schedule** | **string** |  | 
**WaitForCompletion** | **float32** |  | 
**LifetimeSeconds** | **float32** |  | 
**Disabled** | **float32** |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewScheduledJob

`func NewScheduledJob(id int64, label string, systemDefined bool, functionName string, params []map[string]interface{}, schedule string, waitForCompletion float32, lifetimeSeconds float32, disabled float32, ) *ScheduledJob`

NewScheduledJob instantiates a new ScheduledJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledJobWithDefaults

`func NewScheduledJobWithDefaults() *ScheduledJob`

NewScheduledJobWithDefaults instantiates a new ScheduledJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledJob) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *ScheduledJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScheduledJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScheduledJob) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *ScheduledJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSystemDefined

`func (o *ScheduledJob) GetSystemDefined() bool`

GetSystemDefined returns the SystemDefined field if non-nil, zero value otherwise.

### GetSystemDefinedOk

`func (o *ScheduledJob) GetSystemDefinedOk() (*bool, bool)`

GetSystemDefinedOk returns a tuple with the SystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefined

`func (o *ScheduledJob) SetSystemDefined(v bool)`

SetSystemDefined sets SystemDefined field to given value.


### GetFunctionName

`func (o *ScheduledJob) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *ScheduledJob) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *ScheduledJob) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetParams

`func (o *ScheduledJob) GetParams() []map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScheduledJob) GetParamsOk() (*[]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScheduledJob) SetParams(v []map[string]interface{})`

SetParams sets Params field to given value.


### GetSchedule

`func (o *ScheduledJob) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduledJob) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduledJob) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetWaitForCompletion

`func (o *ScheduledJob) GetWaitForCompletion() float32`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *ScheduledJob) GetWaitForCompletionOk() (*float32, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *ScheduledJob) SetWaitForCompletion(v float32)`

SetWaitForCompletion sets WaitForCompletion field to given value.


### GetLifetimeSeconds

`func (o *ScheduledJob) GetLifetimeSeconds() float32`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *ScheduledJob) GetLifetimeSecondsOk() (*float32, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *ScheduledJob) SetLifetimeSeconds(v float32)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.


### GetDisabled

`func (o *ScheduledJob) GetDisabled() float32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ScheduledJob) GetDisabledOk() (*float32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ScheduledJob) SetDisabled(v float32)`

SetDisabled sets Disabled field to given value.


### GetLinks

`func (o *ScheduledJob) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ScheduledJob) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ScheduledJob) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ScheduledJob) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



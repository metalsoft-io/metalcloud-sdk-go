# SyslogRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]SyslogCondition**](SyslogCondition.md) | Conditions that must all be met to trigger this rule | 
**Description** | **string** | Human-readable description of the rule | 
**EmailsList** | **[]string** | Email addresses to notify when the rule is triggered | 
**SendEvent** | Pointer to **map[string]interface{}** | Whether to emit a platform event when the rule is triggered | [optional] 
**SendToMonitoringAgent** | Pointer to **map[string]interface{}** | Whether to forward the notification to the monitoring agent | [optional] 

## Methods

### NewSyslogRule

`func NewSyslogRule(conditions []SyslogCondition, description string, emailsList []string, ) *SyslogRule`

NewSyslogRule instantiates a new SyslogRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogRuleWithDefaults

`func NewSyslogRuleWithDefaults() *SyslogRule`

NewSyslogRuleWithDefaults instantiates a new SyslogRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *SyslogRule) GetConditions() []SyslogCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SyslogRule) GetConditionsOk() (*[]SyslogCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SyslogRule) SetConditions(v []SyslogCondition)`

SetConditions sets Conditions field to given value.


### GetDescription

`func (o *SyslogRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyslogRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyslogRule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailsList

`func (o *SyslogRule) GetEmailsList() []string`

GetEmailsList returns the EmailsList field if non-nil, zero value otherwise.

### GetEmailsListOk

`func (o *SyslogRule) GetEmailsListOk() (*[]string, bool)`

GetEmailsListOk returns a tuple with the EmailsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsList

`func (o *SyslogRule) SetEmailsList(v []string)`

SetEmailsList sets EmailsList field to given value.


### GetSendEvent

`func (o *SyslogRule) GetSendEvent() map[string]interface{}`

GetSendEvent returns the SendEvent field if non-nil, zero value otherwise.

### GetSendEventOk

`func (o *SyslogRule) GetSendEventOk() (*map[string]interface{}, bool)`

GetSendEventOk returns a tuple with the SendEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEvent

`func (o *SyslogRule) SetSendEvent(v map[string]interface{})`

SetSendEvent sets SendEvent field to given value.

### HasSendEvent

`func (o *SyslogRule) HasSendEvent() bool`

HasSendEvent returns a boolean if a field has been set.

### GetSendToMonitoringAgent

`func (o *SyslogRule) GetSendToMonitoringAgent() map[string]interface{}`

GetSendToMonitoringAgent returns the SendToMonitoringAgent field if non-nil, zero value otherwise.

### GetSendToMonitoringAgentOk

`func (o *SyslogRule) GetSendToMonitoringAgentOk() (*map[string]interface{}, bool)`

GetSendToMonitoringAgentOk returns a tuple with the SendToMonitoringAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToMonitoringAgent

`func (o *SyslogRule) SetSendToMonitoringAgent(v map[string]interface{})`

SetSendToMonitoringAgent sets SendToMonitoringAgent field to given value.

### HasSendToMonitoringAgent

`func (o *SyslogRule) HasSendToMonitoringAgent() bool`

HasSendToMonitoringAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



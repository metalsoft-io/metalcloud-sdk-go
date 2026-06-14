# SyslogRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]SyslogConditionDto**](SyslogConditionDto.md) | Conditions that must all be met to trigger this rule | 
**Description** | **string** | Human-readable description of the rule | 
**EmailsList** | **[]string** | Email addresses to notify when the rule is triggered | 
**SendEvent** | Pointer to **map[string]interface{}** | Whether to emit a platform event when the rule is triggered | [optional] 
**SendToMonitoringAgent** | Pointer to **map[string]interface{}** | Whether to forward the notification to the monitoring agent | [optional] 

## Methods

### NewSyslogRuleDto

`func NewSyslogRuleDto(conditions []SyslogConditionDto, description string, emailsList []string, ) *SyslogRuleDto`

NewSyslogRuleDto instantiates a new SyslogRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogRuleDtoWithDefaults

`func NewSyslogRuleDtoWithDefaults() *SyslogRuleDto`

NewSyslogRuleDtoWithDefaults instantiates a new SyslogRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *SyslogRuleDto) GetConditions() []SyslogConditionDto`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SyslogRuleDto) GetConditionsOk() (*[]SyslogConditionDto, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SyslogRuleDto) SetConditions(v []SyslogConditionDto)`

SetConditions sets Conditions field to given value.


### GetDescription

`func (o *SyslogRuleDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyslogRuleDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyslogRuleDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailsList

`func (o *SyslogRuleDto) GetEmailsList() []string`

GetEmailsList returns the EmailsList field if non-nil, zero value otherwise.

### GetEmailsListOk

`func (o *SyslogRuleDto) GetEmailsListOk() (*[]string, bool)`

GetEmailsListOk returns a tuple with the EmailsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsList

`func (o *SyslogRuleDto) SetEmailsList(v []string)`

SetEmailsList sets EmailsList field to given value.


### GetSendEvent

`func (o *SyslogRuleDto) GetSendEvent() map[string]interface{}`

GetSendEvent returns the SendEvent field if non-nil, zero value otherwise.

### GetSendEventOk

`func (o *SyslogRuleDto) GetSendEventOk() (*map[string]interface{}, bool)`

GetSendEventOk returns a tuple with the SendEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEvent

`func (o *SyslogRuleDto) SetSendEvent(v map[string]interface{})`

SetSendEvent sets SendEvent field to given value.

### HasSendEvent

`func (o *SyslogRuleDto) HasSendEvent() bool`

HasSendEvent returns a boolean if a field has been set.

### GetSendToMonitoringAgent

`func (o *SyslogRuleDto) GetSendToMonitoringAgent() map[string]interface{}`

GetSendToMonitoringAgent returns the SendToMonitoringAgent field if non-nil, zero value otherwise.

### GetSendToMonitoringAgentOk

`func (o *SyslogRuleDto) GetSendToMonitoringAgentOk() (*map[string]interface{}, bool)`

GetSendToMonitoringAgentOk returns a tuple with the SendToMonitoringAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToMonitoringAgent

`func (o *SyslogRuleDto) SetSendToMonitoringAgent(v map[string]interface{})`

SetSendToMonitoringAgent sets SendToMonitoringAgent field to given value.

### HasSendToMonitoringAgent

`func (o *SyslogRuleDto) HasSendToMonitoringAgent() bool`

HasSendToMonitoringAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



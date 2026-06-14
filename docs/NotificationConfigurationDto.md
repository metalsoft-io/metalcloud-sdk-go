# NotificationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyslogNotificationRules** | Pointer to [**[]SyslogRuleDto**](SyslogRuleDto.md) | Syslog notification rules for email alerting | [optional] 
**ServerHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about server health events | [optional] 
**NetworkDeviceHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device health events | [optional] 
**NetworkDeviceConfigurationDriftNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device configuration drift | [optional] 
**VmPoolHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about VM pool health events | [optional] 
**EventRetention** | Pointer to **map[string]interface{}** | Event retention policy configuration | [optional] 
**SyslogBufferTtlMs** | Pointer to **float32** | Time-to-live for buffered syslog messages in milliseconds | [optional] 
**SyslogBufferMaxUnique** | Pointer to **float32** | Maximum number of unique syslog messages to buffer before flushing | [optional] 

## Methods

### NewNotificationConfigurationDto

`func NewNotificationConfigurationDto() *NotificationConfigurationDto`

NewNotificationConfigurationDto instantiates a new NotificationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigurationDtoWithDefaults

`func NewNotificationConfigurationDtoWithDefaults() *NotificationConfigurationDto`

NewNotificationConfigurationDtoWithDefaults instantiates a new NotificationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyslogNotificationRules

`func (o *NotificationConfigurationDto) GetSyslogNotificationRules() []SyslogRuleDto`

GetSyslogNotificationRules returns the SyslogNotificationRules field if non-nil, zero value otherwise.

### GetSyslogNotificationRulesOk

`func (o *NotificationConfigurationDto) GetSyslogNotificationRulesOk() (*[]SyslogRuleDto, bool)`

GetSyslogNotificationRulesOk returns a tuple with the SyslogNotificationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNotificationRules

`func (o *NotificationConfigurationDto) SetSyslogNotificationRules(v []SyslogRuleDto)`

SetSyslogNotificationRules sets SyslogNotificationRules field to given value.

### HasSyslogNotificationRules

`func (o *NotificationConfigurationDto) HasSyslogNotificationRules() bool`

HasSyslogNotificationRules returns a boolean if a field has been set.

### GetServerHealthNotificationEmails

`func (o *NotificationConfigurationDto) GetServerHealthNotificationEmails() []string`

GetServerHealthNotificationEmails returns the ServerHealthNotificationEmails field if non-nil, zero value otherwise.

### GetServerHealthNotificationEmailsOk

`func (o *NotificationConfigurationDto) GetServerHealthNotificationEmailsOk() (*[]string, bool)`

GetServerHealthNotificationEmailsOk returns a tuple with the ServerHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHealthNotificationEmails

`func (o *NotificationConfigurationDto) SetServerHealthNotificationEmails(v []string)`

SetServerHealthNotificationEmails sets ServerHealthNotificationEmails field to given value.

### HasServerHealthNotificationEmails

`func (o *NotificationConfigurationDto) HasServerHealthNotificationEmails() bool`

HasServerHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceHealthNotificationEmails

`func (o *NotificationConfigurationDto) GetNetworkDeviceHealthNotificationEmails() []string`

GetNetworkDeviceHealthNotificationEmails returns the NetworkDeviceHealthNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceHealthNotificationEmailsOk

`func (o *NotificationConfigurationDto) GetNetworkDeviceHealthNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceHealthNotificationEmailsOk returns a tuple with the NetworkDeviceHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceHealthNotificationEmails

`func (o *NotificationConfigurationDto) SetNetworkDeviceHealthNotificationEmails(v []string)`

SetNetworkDeviceHealthNotificationEmails sets NetworkDeviceHealthNotificationEmails field to given value.

### HasNetworkDeviceHealthNotificationEmails

`func (o *NotificationConfigurationDto) HasNetworkDeviceHealthNotificationEmails() bool`

HasNetworkDeviceHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *NotificationConfigurationDto) GetNetworkDeviceConfigurationDriftNotificationEmails() []string`

GetNetworkDeviceConfigurationDriftNotificationEmails returns the NetworkDeviceConfigurationDriftNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceConfigurationDriftNotificationEmailsOk

`func (o *NotificationConfigurationDto) GetNetworkDeviceConfigurationDriftNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceConfigurationDriftNotificationEmailsOk returns a tuple with the NetworkDeviceConfigurationDriftNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *NotificationConfigurationDto) SetNetworkDeviceConfigurationDriftNotificationEmails(v []string)`

SetNetworkDeviceConfigurationDriftNotificationEmails sets NetworkDeviceConfigurationDriftNotificationEmails field to given value.

### HasNetworkDeviceConfigurationDriftNotificationEmails

`func (o *NotificationConfigurationDto) HasNetworkDeviceConfigurationDriftNotificationEmails() bool`

HasNetworkDeviceConfigurationDriftNotificationEmails returns a boolean if a field has been set.

### GetVmPoolHealthNotificationEmails

`func (o *NotificationConfigurationDto) GetVmPoolHealthNotificationEmails() []string`

GetVmPoolHealthNotificationEmails returns the VmPoolHealthNotificationEmails field if non-nil, zero value otherwise.

### GetVmPoolHealthNotificationEmailsOk

`func (o *NotificationConfigurationDto) GetVmPoolHealthNotificationEmailsOk() (*[]string, bool)`

GetVmPoolHealthNotificationEmailsOk returns a tuple with the VmPoolHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolHealthNotificationEmails

`func (o *NotificationConfigurationDto) SetVmPoolHealthNotificationEmails(v []string)`

SetVmPoolHealthNotificationEmails sets VmPoolHealthNotificationEmails field to given value.

### HasVmPoolHealthNotificationEmails

`func (o *NotificationConfigurationDto) HasVmPoolHealthNotificationEmails() bool`

HasVmPoolHealthNotificationEmails returns a boolean if a field has been set.

### GetEventRetention

`func (o *NotificationConfigurationDto) GetEventRetention() map[string]interface{}`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *NotificationConfigurationDto) GetEventRetentionOk() (*map[string]interface{}, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *NotificationConfigurationDto) SetEventRetention(v map[string]interface{})`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *NotificationConfigurationDto) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetSyslogBufferTtlMs

`func (o *NotificationConfigurationDto) GetSyslogBufferTtlMs() float32`

GetSyslogBufferTtlMs returns the SyslogBufferTtlMs field if non-nil, zero value otherwise.

### GetSyslogBufferTtlMsOk

`func (o *NotificationConfigurationDto) GetSyslogBufferTtlMsOk() (*float32, bool)`

GetSyslogBufferTtlMsOk returns a tuple with the SyslogBufferTtlMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferTtlMs

`func (o *NotificationConfigurationDto) SetSyslogBufferTtlMs(v float32)`

SetSyslogBufferTtlMs sets SyslogBufferTtlMs field to given value.

### HasSyslogBufferTtlMs

`func (o *NotificationConfigurationDto) HasSyslogBufferTtlMs() bool`

HasSyslogBufferTtlMs returns a boolean if a field has been set.

### GetSyslogBufferMaxUnique

`func (o *NotificationConfigurationDto) GetSyslogBufferMaxUnique() float32`

GetSyslogBufferMaxUnique returns the SyslogBufferMaxUnique field if non-nil, zero value otherwise.

### GetSyslogBufferMaxUniqueOk

`func (o *NotificationConfigurationDto) GetSyslogBufferMaxUniqueOk() (*float32, bool)`

GetSyslogBufferMaxUniqueOk returns a tuple with the SyslogBufferMaxUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferMaxUnique

`func (o *NotificationConfigurationDto) SetSyslogBufferMaxUnique(v float32)`

SetSyslogBufferMaxUnique sets SyslogBufferMaxUnique field to given value.

### HasSyslogBufferMaxUnique

`func (o *NotificationConfigurationDto) HasSyslogBufferMaxUnique() bool`

HasSyslogBufferMaxUnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



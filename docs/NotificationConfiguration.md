# NotificationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyslogNotificationRules** | Pointer to [**[]SyslogRule**](SyslogRule.md) | Syslog notification rules for email alerting | [optional] 
**ServerHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about server health events | [optional] 
**NetworkDeviceHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device health events | [optional] 
**NetworkDeviceConfigurationDriftNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device configuration drift | [optional] 
**ServerFirmwareBaselineDriftNotificationEmails** | Pointer to **[]string** | Email addresses to notify when a server falls below the minimum versions of its firmware baseline | [optional] 
**ServerConfigurationDriftNotificationEmails** | Pointer to **[]string** | Email addresses to notify when a server’s configuration changes outside of a planned change. | [optional] 
**VmPoolHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about VM pool health events | [optional] 
**EventRetention** | Pointer to **map[string]interface{}** | Event retention policy configuration | [optional] 
**SyslogBufferTtlMs** | Pointer to **float32** | Time-to-live for buffered syslog messages in milliseconds | [optional] 
**SyslogBufferMaxUnique** | Pointer to **float32** | Maximum number of unique syslog messages to buffer before flushing | [optional] 
**ServerMonitoringAgentDiagnosticsEnabled** | Pointer to **bool** | Whether to invoke the monitoring agent for a diagnostic when a server health check degrades. When disabled the alert email is still sent, using the plain template. | [optional] 
**NetworkDeviceMonitoringAgentDiagnosticsEnabled** | Pointer to **bool** | Whether to invoke the monitoring agent for a diagnostic when a network device health check degrades. When disabled the alert email is still sent, using the plain template. | [optional] 
**MonitoringAgentHealthDiagnosticsTimeoutMs** | Pointer to **float32** | How long to wait for the monitoring agent diagnostic, in milliseconds. Once it elapses the alert email is sent without the monitoring agent section rather than being held back. | [optional] 

## Methods

### NewNotificationConfiguration

`func NewNotificationConfiguration() *NotificationConfiguration`

NewNotificationConfiguration instantiates a new NotificationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigurationWithDefaults

`func NewNotificationConfigurationWithDefaults() *NotificationConfiguration`

NewNotificationConfigurationWithDefaults instantiates a new NotificationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyslogNotificationRules

`func (o *NotificationConfiguration) GetSyslogNotificationRules() []SyslogRule`

GetSyslogNotificationRules returns the SyslogNotificationRules field if non-nil, zero value otherwise.

### GetSyslogNotificationRulesOk

`func (o *NotificationConfiguration) GetSyslogNotificationRulesOk() (*[]SyslogRule, bool)`

GetSyslogNotificationRulesOk returns a tuple with the SyslogNotificationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNotificationRules

`func (o *NotificationConfiguration) SetSyslogNotificationRules(v []SyslogRule)`

SetSyslogNotificationRules sets SyslogNotificationRules field to given value.

### HasSyslogNotificationRules

`func (o *NotificationConfiguration) HasSyslogNotificationRules() bool`

HasSyslogNotificationRules returns a boolean if a field has been set.

### GetServerHealthNotificationEmails

`func (o *NotificationConfiguration) GetServerHealthNotificationEmails() []string`

GetServerHealthNotificationEmails returns the ServerHealthNotificationEmails field if non-nil, zero value otherwise.

### GetServerHealthNotificationEmailsOk

`func (o *NotificationConfiguration) GetServerHealthNotificationEmailsOk() (*[]string, bool)`

GetServerHealthNotificationEmailsOk returns a tuple with the ServerHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHealthNotificationEmails

`func (o *NotificationConfiguration) SetServerHealthNotificationEmails(v []string)`

SetServerHealthNotificationEmails sets ServerHealthNotificationEmails field to given value.

### HasServerHealthNotificationEmails

`func (o *NotificationConfiguration) HasServerHealthNotificationEmails() bool`

HasServerHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceHealthNotificationEmails

`func (o *NotificationConfiguration) GetNetworkDeviceHealthNotificationEmails() []string`

GetNetworkDeviceHealthNotificationEmails returns the NetworkDeviceHealthNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceHealthNotificationEmailsOk

`func (o *NotificationConfiguration) GetNetworkDeviceHealthNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceHealthNotificationEmailsOk returns a tuple with the NetworkDeviceHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceHealthNotificationEmails

`func (o *NotificationConfiguration) SetNetworkDeviceHealthNotificationEmails(v []string)`

SetNetworkDeviceHealthNotificationEmails sets NetworkDeviceHealthNotificationEmails field to given value.

### HasNetworkDeviceHealthNotificationEmails

`func (o *NotificationConfiguration) HasNetworkDeviceHealthNotificationEmails() bool`

HasNetworkDeviceHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *NotificationConfiguration) GetNetworkDeviceConfigurationDriftNotificationEmails() []string`

GetNetworkDeviceConfigurationDriftNotificationEmails returns the NetworkDeviceConfigurationDriftNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceConfigurationDriftNotificationEmailsOk

`func (o *NotificationConfiguration) GetNetworkDeviceConfigurationDriftNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceConfigurationDriftNotificationEmailsOk returns a tuple with the NetworkDeviceConfigurationDriftNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *NotificationConfiguration) SetNetworkDeviceConfigurationDriftNotificationEmails(v []string)`

SetNetworkDeviceConfigurationDriftNotificationEmails sets NetworkDeviceConfigurationDriftNotificationEmails field to given value.

### HasNetworkDeviceConfigurationDriftNotificationEmails

`func (o *NotificationConfiguration) HasNetworkDeviceConfigurationDriftNotificationEmails() bool`

HasNetworkDeviceConfigurationDriftNotificationEmails returns a boolean if a field has been set.

### GetServerFirmwareBaselineDriftNotificationEmails

`func (o *NotificationConfiguration) GetServerFirmwareBaselineDriftNotificationEmails() []string`

GetServerFirmwareBaselineDriftNotificationEmails returns the ServerFirmwareBaselineDriftNotificationEmails field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineDriftNotificationEmailsOk

`func (o *NotificationConfiguration) GetServerFirmwareBaselineDriftNotificationEmailsOk() (*[]string, bool)`

GetServerFirmwareBaselineDriftNotificationEmailsOk returns a tuple with the ServerFirmwareBaselineDriftNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineDriftNotificationEmails

`func (o *NotificationConfiguration) SetServerFirmwareBaselineDriftNotificationEmails(v []string)`

SetServerFirmwareBaselineDriftNotificationEmails sets ServerFirmwareBaselineDriftNotificationEmails field to given value.

### HasServerFirmwareBaselineDriftNotificationEmails

`func (o *NotificationConfiguration) HasServerFirmwareBaselineDriftNotificationEmails() bool`

HasServerFirmwareBaselineDriftNotificationEmails returns a boolean if a field has been set.

### GetServerConfigurationDriftNotificationEmails

`func (o *NotificationConfiguration) GetServerConfigurationDriftNotificationEmails() []string`

GetServerConfigurationDriftNotificationEmails returns the ServerConfigurationDriftNotificationEmails field if non-nil, zero value otherwise.

### GetServerConfigurationDriftNotificationEmailsOk

`func (o *NotificationConfiguration) GetServerConfigurationDriftNotificationEmailsOk() (*[]string, bool)`

GetServerConfigurationDriftNotificationEmailsOk returns a tuple with the ServerConfigurationDriftNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerConfigurationDriftNotificationEmails

`func (o *NotificationConfiguration) SetServerConfigurationDriftNotificationEmails(v []string)`

SetServerConfigurationDriftNotificationEmails sets ServerConfigurationDriftNotificationEmails field to given value.

### HasServerConfigurationDriftNotificationEmails

`func (o *NotificationConfiguration) HasServerConfigurationDriftNotificationEmails() bool`

HasServerConfigurationDriftNotificationEmails returns a boolean if a field has been set.

### GetVmPoolHealthNotificationEmails

`func (o *NotificationConfiguration) GetVmPoolHealthNotificationEmails() []string`

GetVmPoolHealthNotificationEmails returns the VmPoolHealthNotificationEmails field if non-nil, zero value otherwise.

### GetVmPoolHealthNotificationEmailsOk

`func (o *NotificationConfiguration) GetVmPoolHealthNotificationEmailsOk() (*[]string, bool)`

GetVmPoolHealthNotificationEmailsOk returns a tuple with the VmPoolHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolHealthNotificationEmails

`func (o *NotificationConfiguration) SetVmPoolHealthNotificationEmails(v []string)`

SetVmPoolHealthNotificationEmails sets VmPoolHealthNotificationEmails field to given value.

### HasVmPoolHealthNotificationEmails

`func (o *NotificationConfiguration) HasVmPoolHealthNotificationEmails() bool`

HasVmPoolHealthNotificationEmails returns a boolean if a field has been set.

### GetEventRetention

`func (o *NotificationConfiguration) GetEventRetention() map[string]interface{}`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *NotificationConfiguration) GetEventRetentionOk() (*map[string]interface{}, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *NotificationConfiguration) SetEventRetention(v map[string]interface{})`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *NotificationConfiguration) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetSyslogBufferTtlMs

`func (o *NotificationConfiguration) GetSyslogBufferTtlMs() float32`

GetSyslogBufferTtlMs returns the SyslogBufferTtlMs field if non-nil, zero value otherwise.

### GetSyslogBufferTtlMsOk

`func (o *NotificationConfiguration) GetSyslogBufferTtlMsOk() (*float32, bool)`

GetSyslogBufferTtlMsOk returns a tuple with the SyslogBufferTtlMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferTtlMs

`func (o *NotificationConfiguration) SetSyslogBufferTtlMs(v float32)`

SetSyslogBufferTtlMs sets SyslogBufferTtlMs field to given value.

### HasSyslogBufferTtlMs

`func (o *NotificationConfiguration) HasSyslogBufferTtlMs() bool`

HasSyslogBufferTtlMs returns a boolean if a field has been set.

### GetSyslogBufferMaxUnique

`func (o *NotificationConfiguration) GetSyslogBufferMaxUnique() float32`

GetSyslogBufferMaxUnique returns the SyslogBufferMaxUnique field if non-nil, zero value otherwise.

### GetSyslogBufferMaxUniqueOk

`func (o *NotificationConfiguration) GetSyslogBufferMaxUniqueOk() (*float32, bool)`

GetSyslogBufferMaxUniqueOk returns a tuple with the SyslogBufferMaxUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferMaxUnique

`func (o *NotificationConfiguration) SetSyslogBufferMaxUnique(v float32)`

SetSyslogBufferMaxUnique sets SyslogBufferMaxUnique field to given value.

### HasSyslogBufferMaxUnique

`func (o *NotificationConfiguration) HasSyslogBufferMaxUnique() bool`

HasSyslogBufferMaxUnique returns a boolean if a field has been set.

### GetServerMonitoringAgentDiagnosticsEnabled

`func (o *NotificationConfiguration) GetServerMonitoringAgentDiagnosticsEnabled() bool`

GetServerMonitoringAgentDiagnosticsEnabled returns the ServerMonitoringAgentDiagnosticsEnabled field if non-nil, zero value otherwise.

### GetServerMonitoringAgentDiagnosticsEnabledOk

`func (o *NotificationConfiguration) GetServerMonitoringAgentDiagnosticsEnabledOk() (*bool, bool)`

GetServerMonitoringAgentDiagnosticsEnabledOk returns a tuple with the ServerMonitoringAgentDiagnosticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMonitoringAgentDiagnosticsEnabled

`func (o *NotificationConfiguration) SetServerMonitoringAgentDiagnosticsEnabled(v bool)`

SetServerMonitoringAgentDiagnosticsEnabled sets ServerMonitoringAgentDiagnosticsEnabled field to given value.

### HasServerMonitoringAgentDiagnosticsEnabled

`func (o *NotificationConfiguration) HasServerMonitoringAgentDiagnosticsEnabled() bool`

HasServerMonitoringAgentDiagnosticsEnabled returns a boolean if a field has been set.

### GetNetworkDeviceMonitoringAgentDiagnosticsEnabled

`func (o *NotificationConfiguration) GetNetworkDeviceMonitoringAgentDiagnosticsEnabled() bool`

GetNetworkDeviceMonitoringAgentDiagnosticsEnabled returns the NetworkDeviceMonitoringAgentDiagnosticsEnabled field if non-nil, zero value otherwise.

### GetNetworkDeviceMonitoringAgentDiagnosticsEnabledOk

`func (o *NotificationConfiguration) GetNetworkDeviceMonitoringAgentDiagnosticsEnabledOk() (*bool, bool)`

GetNetworkDeviceMonitoringAgentDiagnosticsEnabledOk returns a tuple with the NetworkDeviceMonitoringAgentDiagnosticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceMonitoringAgentDiagnosticsEnabled

`func (o *NotificationConfiguration) SetNetworkDeviceMonitoringAgentDiagnosticsEnabled(v bool)`

SetNetworkDeviceMonitoringAgentDiagnosticsEnabled sets NetworkDeviceMonitoringAgentDiagnosticsEnabled field to given value.

### HasNetworkDeviceMonitoringAgentDiagnosticsEnabled

`func (o *NotificationConfiguration) HasNetworkDeviceMonitoringAgentDiagnosticsEnabled() bool`

HasNetworkDeviceMonitoringAgentDiagnosticsEnabled returns a boolean if a field has been set.

### GetMonitoringAgentHealthDiagnosticsTimeoutMs

`func (o *NotificationConfiguration) GetMonitoringAgentHealthDiagnosticsTimeoutMs() float32`

GetMonitoringAgentHealthDiagnosticsTimeoutMs returns the MonitoringAgentHealthDiagnosticsTimeoutMs field if non-nil, zero value otherwise.

### GetMonitoringAgentHealthDiagnosticsTimeoutMsOk

`func (o *NotificationConfiguration) GetMonitoringAgentHealthDiagnosticsTimeoutMsOk() (*float32, bool)`

GetMonitoringAgentHealthDiagnosticsTimeoutMsOk returns a tuple with the MonitoringAgentHealthDiagnosticsTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringAgentHealthDiagnosticsTimeoutMs

`func (o *NotificationConfiguration) SetMonitoringAgentHealthDiagnosticsTimeoutMs(v float32)`

SetMonitoringAgentHealthDiagnosticsTimeoutMs sets MonitoringAgentHealthDiagnosticsTimeoutMs field to given value.

### HasMonitoringAgentHealthDiagnosticsTimeoutMs

`func (o *NotificationConfiguration) HasMonitoringAgentHealthDiagnosticsTimeoutMs() bool`

HasMonitoringAgentHealthDiagnosticsTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



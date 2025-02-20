# AgentCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxyEnabled** | **bool** | HTTP proxy capability status | 
**InbandHttpProxyEnabled** | **bool** | Inband HTTP proxy capability status | 
**FileTransferEnabled** | **bool** | File transfer capability status | 
**InbandFileTransferEnabled** | **bool** | Inband file transfer capability status | 
**SwitchSubscriptionEnabled** | **bool** | Switch subscription capability status | 
**CommandExecutionEnabled** | **bool** | Command execution capability status | 
**VncEnabled** | **bool** | VNC capability status | 
**SpiceEnabled** | **bool** | SPICE capability status | 
**SyslogEnabled** | **bool** | Syslog capability status | 
**DhcpOobEnabled** | **bool** | DHCP OOB capability status | 
**NetconfEnabled** | **bool** | NETCONF capability status | 
**AnsibleRunnerEnabled** | **bool** | Ansible runner capability status | 

## Methods

### NewAgentCapabilities

`func NewAgentCapabilities(httpProxyEnabled bool, inbandHttpProxyEnabled bool, fileTransferEnabled bool, inbandFileTransferEnabled bool, switchSubscriptionEnabled bool, commandExecutionEnabled bool, vncEnabled bool, spiceEnabled bool, syslogEnabled bool, dhcpOobEnabled bool, netconfEnabled bool, ansibleRunnerEnabled bool, ) *AgentCapabilities`

NewAgentCapabilities instantiates a new AgentCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCapabilitiesWithDefaults

`func NewAgentCapabilitiesWithDefaults() *AgentCapabilities`

NewAgentCapabilitiesWithDefaults instantiates a new AgentCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxyEnabled

`func (o *AgentCapabilities) GetHttpProxyEnabled() bool`

GetHttpProxyEnabled returns the HttpProxyEnabled field if non-nil, zero value otherwise.

### GetHttpProxyEnabledOk

`func (o *AgentCapabilities) GetHttpProxyEnabledOk() (*bool, bool)`

GetHttpProxyEnabledOk returns a tuple with the HttpProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyEnabled

`func (o *AgentCapabilities) SetHttpProxyEnabled(v bool)`

SetHttpProxyEnabled sets HttpProxyEnabled field to given value.


### GetInbandHttpProxyEnabled

`func (o *AgentCapabilities) GetInbandHttpProxyEnabled() bool`

GetInbandHttpProxyEnabled returns the InbandHttpProxyEnabled field if non-nil, zero value otherwise.

### GetInbandHttpProxyEnabledOk

`func (o *AgentCapabilities) GetInbandHttpProxyEnabledOk() (*bool, bool)`

GetInbandHttpProxyEnabledOk returns a tuple with the InbandHttpProxyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandHttpProxyEnabled

`func (o *AgentCapabilities) SetInbandHttpProxyEnabled(v bool)`

SetInbandHttpProxyEnabled sets InbandHttpProxyEnabled field to given value.


### GetFileTransferEnabled

`func (o *AgentCapabilities) GetFileTransferEnabled() bool`

GetFileTransferEnabled returns the FileTransferEnabled field if non-nil, zero value otherwise.

### GetFileTransferEnabledOk

`func (o *AgentCapabilities) GetFileTransferEnabledOk() (*bool, bool)`

GetFileTransferEnabledOk returns a tuple with the FileTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransferEnabled

`func (o *AgentCapabilities) SetFileTransferEnabled(v bool)`

SetFileTransferEnabled sets FileTransferEnabled field to given value.


### GetInbandFileTransferEnabled

`func (o *AgentCapabilities) GetInbandFileTransferEnabled() bool`

GetInbandFileTransferEnabled returns the InbandFileTransferEnabled field if non-nil, zero value otherwise.

### GetInbandFileTransferEnabledOk

`func (o *AgentCapabilities) GetInbandFileTransferEnabledOk() (*bool, bool)`

GetInbandFileTransferEnabledOk returns a tuple with the InbandFileTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandFileTransferEnabled

`func (o *AgentCapabilities) SetInbandFileTransferEnabled(v bool)`

SetInbandFileTransferEnabled sets InbandFileTransferEnabled field to given value.


### GetSwitchSubscriptionEnabled

`func (o *AgentCapabilities) GetSwitchSubscriptionEnabled() bool`

GetSwitchSubscriptionEnabled returns the SwitchSubscriptionEnabled field if non-nil, zero value otherwise.

### GetSwitchSubscriptionEnabledOk

`func (o *AgentCapabilities) GetSwitchSubscriptionEnabledOk() (*bool, bool)`

GetSwitchSubscriptionEnabledOk returns a tuple with the SwitchSubscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSubscriptionEnabled

`func (o *AgentCapabilities) SetSwitchSubscriptionEnabled(v bool)`

SetSwitchSubscriptionEnabled sets SwitchSubscriptionEnabled field to given value.


### GetCommandExecutionEnabled

`func (o *AgentCapabilities) GetCommandExecutionEnabled() bool`

GetCommandExecutionEnabled returns the CommandExecutionEnabled field if non-nil, zero value otherwise.

### GetCommandExecutionEnabledOk

`func (o *AgentCapabilities) GetCommandExecutionEnabledOk() (*bool, bool)`

GetCommandExecutionEnabledOk returns a tuple with the CommandExecutionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandExecutionEnabled

`func (o *AgentCapabilities) SetCommandExecutionEnabled(v bool)`

SetCommandExecutionEnabled sets CommandExecutionEnabled field to given value.


### GetVncEnabled

`func (o *AgentCapabilities) GetVncEnabled() bool`

GetVncEnabled returns the VncEnabled field if non-nil, zero value otherwise.

### GetVncEnabledOk

`func (o *AgentCapabilities) GetVncEnabledOk() (*bool, bool)`

GetVncEnabledOk returns a tuple with the VncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncEnabled

`func (o *AgentCapabilities) SetVncEnabled(v bool)`

SetVncEnabled sets VncEnabled field to given value.


### GetSpiceEnabled

`func (o *AgentCapabilities) GetSpiceEnabled() bool`

GetSpiceEnabled returns the SpiceEnabled field if non-nil, zero value otherwise.

### GetSpiceEnabledOk

`func (o *AgentCapabilities) GetSpiceEnabledOk() (*bool, bool)`

GetSpiceEnabledOk returns a tuple with the SpiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiceEnabled

`func (o *AgentCapabilities) SetSpiceEnabled(v bool)`

SetSpiceEnabled sets SpiceEnabled field to given value.


### GetSyslogEnabled

`func (o *AgentCapabilities) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *AgentCapabilities) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *AgentCapabilities) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.


### GetDhcpOobEnabled

`func (o *AgentCapabilities) GetDhcpOobEnabled() bool`

GetDhcpOobEnabled returns the DhcpOobEnabled field if non-nil, zero value otherwise.

### GetDhcpOobEnabledOk

`func (o *AgentCapabilities) GetDhcpOobEnabledOk() (*bool, bool)`

GetDhcpOobEnabledOk returns a tuple with the DhcpOobEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOobEnabled

`func (o *AgentCapabilities) SetDhcpOobEnabled(v bool)`

SetDhcpOobEnabled sets DhcpOobEnabled field to given value.


### GetNetconfEnabled

`func (o *AgentCapabilities) GetNetconfEnabled() bool`

GetNetconfEnabled returns the NetconfEnabled field if non-nil, zero value otherwise.

### GetNetconfEnabledOk

`func (o *AgentCapabilities) GetNetconfEnabledOk() (*bool, bool)`

GetNetconfEnabledOk returns a tuple with the NetconfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetconfEnabled

`func (o *AgentCapabilities) SetNetconfEnabled(v bool)`

SetNetconfEnabled sets NetconfEnabled field to given value.


### GetAnsibleRunnerEnabled

`func (o *AgentCapabilities) GetAnsibleRunnerEnabled() bool`

GetAnsibleRunnerEnabled returns the AnsibleRunnerEnabled field if non-nil, zero value otherwise.

### GetAnsibleRunnerEnabledOk

`func (o *AgentCapabilities) GetAnsibleRunnerEnabledOk() (*bool, bool)`

GetAnsibleRunnerEnabledOk returns a tuple with the AnsibleRunnerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleRunnerEnabled

`func (o *AgentCapabilities) SetAnsibleRunnerEnabled(v bool)`

SetAnsibleRunnerEnabled sets AnsibleRunnerEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



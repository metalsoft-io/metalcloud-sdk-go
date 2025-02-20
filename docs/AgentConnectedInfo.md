# AgentConnectedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** | Agent identifier | 
**Hostname** | **string** | Host name | 
**ApplicationName** | **string** | Application name | 
**Timestamp** | **string** | Connection timestamp | 
**Capabilities** | [**AgentCapabilities**](AgentCapabilities.md) | Agent capabilities | 
**IpInfo** | [**AgentIpInfo**](AgentIpInfo.md) | IP information | 

## Methods

### NewAgentConnectedInfo

`func NewAgentConnectedInfo(agentId string, hostname string, applicationName string, timestamp string, capabilities AgentCapabilities, ipInfo AgentIpInfo, ) *AgentConnectedInfo`

NewAgentConnectedInfo instantiates a new AgentConnectedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConnectedInfoWithDefaults

`func NewAgentConnectedInfoWithDefaults() *AgentConnectedInfo`

NewAgentConnectedInfoWithDefaults instantiates a new AgentConnectedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AgentConnectedInfo) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentConnectedInfo) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentConnectedInfo) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetHostname

`func (o *AgentConnectedInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AgentConnectedInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AgentConnectedInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetApplicationName

`func (o *AgentConnectedInfo) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AgentConnectedInfo) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AgentConnectedInfo) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetTimestamp

`func (o *AgentConnectedInfo) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AgentConnectedInfo) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AgentConnectedInfo) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetCapabilities

`func (o *AgentConnectedInfo) GetCapabilities() AgentCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AgentConnectedInfo) GetCapabilitiesOk() (*AgentCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AgentConnectedInfo) SetCapabilities(v AgentCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetIpInfo

`func (o *AgentConnectedInfo) GetIpInfo() AgentIpInfo`

GetIpInfo returns the IpInfo field if non-nil, zero value otherwise.

### GetIpInfoOk

`func (o *AgentConnectedInfo) GetIpInfoOk() (*AgentIpInfo, bool)`

GetIpInfoOk returns a tuple with the IpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInfo

`func (o *AgentConnectedInfo) SetIpInfo(v AgentIpInfo)`

SetIpInfo sets IpInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



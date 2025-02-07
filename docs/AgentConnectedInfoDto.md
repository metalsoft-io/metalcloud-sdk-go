# AgentConnectedInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** | Agent identifier | 
**Hostname** | **string** | Host name | 
**ApplicationName** | **string** | Application name | 
**Timestamp** | **string** | Connection timestamp | 
**Capabilities** | [**AgentCapabilitiesDto**](AgentCapabilitiesDto.md) | Agent capabilities | 
**IpInfo** | [**AgentIpInfoDto**](AgentIpInfoDto.md) | IP information | 

## Methods

### NewAgentConnectedInfoDto

`func NewAgentConnectedInfoDto(agentId string, hostname string, applicationName string, timestamp string, capabilities AgentCapabilitiesDto, ipInfo AgentIpInfoDto, ) *AgentConnectedInfoDto`

NewAgentConnectedInfoDto instantiates a new AgentConnectedInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConnectedInfoDtoWithDefaults

`func NewAgentConnectedInfoDtoWithDefaults() *AgentConnectedInfoDto`

NewAgentConnectedInfoDtoWithDefaults instantiates a new AgentConnectedInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AgentConnectedInfoDto) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentConnectedInfoDto) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentConnectedInfoDto) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetHostname

`func (o *AgentConnectedInfoDto) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AgentConnectedInfoDto) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AgentConnectedInfoDto) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetApplicationName

`func (o *AgentConnectedInfoDto) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AgentConnectedInfoDto) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AgentConnectedInfoDto) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetTimestamp

`func (o *AgentConnectedInfoDto) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AgentConnectedInfoDto) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AgentConnectedInfoDto) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetCapabilities

`func (o *AgentConnectedInfoDto) GetCapabilities() AgentCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AgentConnectedInfoDto) GetCapabilitiesOk() (*AgentCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AgentConnectedInfoDto) SetCapabilities(v AgentCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.


### GetIpInfo

`func (o *AgentConnectedInfoDto) GetIpInfo() AgentIpInfoDto`

GetIpInfo returns the IpInfo field if non-nil, zero value otherwise.

### GetIpInfoOk

`func (o *AgentConnectedInfoDto) GetIpInfoOk() (*AgentIpInfoDto, bool)`

GetIpInfoOk returns a tuple with the IpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInfo

`func (o *AgentConnectedInfoDto) SetIpInfo(v AgentIpInfoDto)`

SetIpInfo sets IpInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



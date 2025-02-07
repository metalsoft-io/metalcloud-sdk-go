# AgentInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteName** | **string** | Site name | 
**AgentSeenIpAddress** | **string** | Agent seen IP address with port | 
**AgentType** | **string** | Agent type | 
**AgentVersion** | **string** | Agent version | 
**AgentCreationTimestamp** | **string** | Agent creation timestamp | 
**AgentUpdatedTimestamp** | **string** | Agent update timestamp | 
**AgentSeenTimestamp** | **string** | Agent last seen timestamp | 
**AgentConnectedInfo** | [**AgentConnectedInfoDto**](AgentConnectedInfoDto.md) | Agent connection information | 

## Methods

### NewAgentInfoDto

`func NewAgentInfoDto(siteName string, agentSeenIpAddress string, agentType string, agentVersion string, agentCreationTimestamp string, agentUpdatedTimestamp string, agentSeenTimestamp string, agentConnectedInfo AgentConnectedInfoDto, ) *AgentInfoDto`

NewAgentInfoDto instantiates a new AgentInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInfoDtoWithDefaults

`func NewAgentInfoDtoWithDefaults() *AgentInfoDto`

NewAgentInfoDtoWithDefaults instantiates a new AgentInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteName

`func (o *AgentInfoDto) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *AgentInfoDto) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *AgentInfoDto) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetAgentSeenIpAddress

`func (o *AgentInfoDto) GetAgentSeenIpAddress() string`

GetAgentSeenIpAddress returns the AgentSeenIpAddress field if non-nil, zero value otherwise.

### GetAgentSeenIpAddressOk

`func (o *AgentInfoDto) GetAgentSeenIpAddressOk() (*string, bool)`

GetAgentSeenIpAddressOk returns a tuple with the AgentSeenIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSeenIpAddress

`func (o *AgentInfoDto) SetAgentSeenIpAddress(v string)`

SetAgentSeenIpAddress sets AgentSeenIpAddress field to given value.


### GetAgentType

`func (o *AgentInfoDto) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentInfoDto) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentInfoDto) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetAgentVersion

`func (o *AgentInfoDto) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AgentInfoDto) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AgentInfoDto) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### GetAgentCreationTimestamp

`func (o *AgentInfoDto) GetAgentCreationTimestamp() string`

GetAgentCreationTimestamp returns the AgentCreationTimestamp field if non-nil, zero value otherwise.

### GetAgentCreationTimestampOk

`func (o *AgentInfoDto) GetAgentCreationTimestampOk() (*string, bool)`

GetAgentCreationTimestampOk returns a tuple with the AgentCreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCreationTimestamp

`func (o *AgentInfoDto) SetAgentCreationTimestamp(v string)`

SetAgentCreationTimestamp sets AgentCreationTimestamp field to given value.


### GetAgentUpdatedTimestamp

`func (o *AgentInfoDto) GetAgentUpdatedTimestamp() string`

GetAgentUpdatedTimestamp returns the AgentUpdatedTimestamp field if non-nil, zero value otherwise.

### GetAgentUpdatedTimestampOk

`func (o *AgentInfoDto) GetAgentUpdatedTimestampOk() (*string, bool)`

GetAgentUpdatedTimestampOk returns a tuple with the AgentUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUpdatedTimestamp

`func (o *AgentInfoDto) SetAgentUpdatedTimestamp(v string)`

SetAgentUpdatedTimestamp sets AgentUpdatedTimestamp field to given value.


### GetAgentSeenTimestamp

`func (o *AgentInfoDto) GetAgentSeenTimestamp() string`

GetAgentSeenTimestamp returns the AgentSeenTimestamp field if non-nil, zero value otherwise.

### GetAgentSeenTimestampOk

`func (o *AgentInfoDto) GetAgentSeenTimestampOk() (*string, bool)`

GetAgentSeenTimestampOk returns a tuple with the AgentSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSeenTimestamp

`func (o *AgentInfoDto) SetAgentSeenTimestamp(v string)`

SetAgentSeenTimestamp sets AgentSeenTimestamp field to given value.


### GetAgentConnectedInfo

`func (o *AgentInfoDto) GetAgentConnectedInfo() AgentConnectedInfoDto`

GetAgentConnectedInfo returns the AgentConnectedInfo field if non-nil, zero value otherwise.

### GetAgentConnectedInfoOk

`func (o *AgentInfoDto) GetAgentConnectedInfoOk() (*AgentConnectedInfoDto, bool)`

GetAgentConnectedInfoOk returns a tuple with the AgentConnectedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectedInfo

`func (o *AgentInfoDto) SetAgentConnectedInfo(v AgentConnectedInfoDto)`

SetAgentConnectedInfo sets AgentConnectedInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



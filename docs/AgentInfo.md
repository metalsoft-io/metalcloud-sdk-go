# AgentInfo

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
**AgentConnectedInfo** | [**AgentConnectedInfo**](AgentConnectedInfo.md) | Agent connection information | 

## Methods

### NewAgentInfo

`func NewAgentInfo(siteName string, agentSeenIpAddress string, agentType string, agentVersion string, agentCreationTimestamp string, agentUpdatedTimestamp string, agentSeenTimestamp string, agentConnectedInfo AgentConnectedInfo, ) *AgentInfo`

NewAgentInfo instantiates a new AgentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInfoWithDefaults

`func NewAgentInfoWithDefaults() *AgentInfo`

NewAgentInfoWithDefaults instantiates a new AgentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteName

`func (o *AgentInfo) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *AgentInfo) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *AgentInfo) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### GetAgentSeenIpAddress

`func (o *AgentInfo) GetAgentSeenIpAddress() string`

GetAgentSeenIpAddress returns the AgentSeenIpAddress field if non-nil, zero value otherwise.

### GetAgentSeenIpAddressOk

`func (o *AgentInfo) GetAgentSeenIpAddressOk() (*string, bool)`

GetAgentSeenIpAddressOk returns a tuple with the AgentSeenIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSeenIpAddress

`func (o *AgentInfo) SetAgentSeenIpAddress(v string)`

SetAgentSeenIpAddress sets AgentSeenIpAddress field to given value.


### GetAgentType

`func (o *AgentInfo) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentInfo) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentInfo) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetAgentVersion

`func (o *AgentInfo) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AgentInfo) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AgentInfo) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### GetAgentCreationTimestamp

`func (o *AgentInfo) GetAgentCreationTimestamp() string`

GetAgentCreationTimestamp returns the AgentCreationTimestamp field if non-nil, zero value otherwise.

### GetAgentCreationTimestampOk

`func (o *AgentInfo) GetAgentCreationTimestampOk() (*string, bool)`

GetAgentCreationTimestampOk returns a tuple with the AgentCreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCreationTimestamp

`func (o *AgentInfo) SetAgentCreationTimestamp(v string)`

SetAgentCreationTimestamp sets AgentCreationTimestamp field to given value.


### GetAgentUpdatedTimestamp

`func (o *AgentInfo) GetAgentUpdatedTimestamp() string`

GetAgentUpdatedTimestamp returns the AgentUpdatedTimestamp field if non-nil, zero value otherwise.

### GetAgentUpdatedTimestampOk

`func (o *AgentInfo) GetAgentUpdatedTimestampOk() (*string, bool)`

GetAgentUpdatedTimestampOk returns a tuple with the AgentUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentUpdatedTimestamp

`func (o *AgentInfo) SetAgentUpdatedTimestamp(v string)`

SetAgentUpdatedTimestamp sets AgentUpdatedTimestamp field to given value.


### GetAgentSeenTimestamp

`func (o *AgentInfo) GetAgentSeenTimestamp() string`

GetAgentSeenTimestamp returns the AgentSeenTimestamp field if non-nil, zero value otherwise.

### GetAgentSeenTimestampOk

`func (o *AgentInfo) GetAgentSeenTimestampOk() (*string, bool)`

GetAgentSeenTimestampOk returns a tuple with the AgentSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSeenTimestamp

`func (o *AgentInfo) SetAgentSeenTimestamp(v string)`

SetAgentSeenTimestamp sets AgentSeenTimestamp field to given value.


### GetAgentConnectedInfo

`func (o *AgentInfo) GetAgentConnectedInfo() AgentConnectedInfo`

GetAgentConnectedInfo returns the AgentConnectedInfo field if non-nil, zero value otherwise.

### GetAgentConnectedInfoOk

`func (o *AgentInfo) GetAgentConnectedInfoOk() (*AgentConnectedInfo, bool)`

GetAgentConnectedInfoOk returns a tuple with the AgentConnectedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectedInfo

`func (o *AgentInfo) SetAgentConnectedInfo(v AgentConnectedInfo)`

SetAgentConnectedInfo sets AgentConnectedInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



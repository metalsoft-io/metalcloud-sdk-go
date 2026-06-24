# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the event | 
**UserIdAuthenticated** | Pointer to **string** | The id of the user who triggered the event | [optional] 
**Type** | [**EventTypes**](EventTypes.md) | The type of the event | 
**Level** | [**EventLevel**](EventLevel.md) | The level of the event | 
**Visibility** | [**EventVisibility**](EventVisibility.md) | The visibility of the event | 
**InfrastructureId** | Pointer to **string** | The id of the infrastructure linked to the event | [optional] 
**UserId** | Pointer to **string** | The id of the user linked to the event | [optional] 
**ServerId** | Pointer to **string** | The id of the server linked to the event | [optional] 
**NetworkDeviceId** | Pointer to **string** | The id of the network device linked to the event | [optional] 
**NetworkDeviceControllerId** | Pointer to **string** | The id of the network device controller linked to the event | [optional] 
**StorageId** | Pointer to **string** | The id of the storage linked to the event | [optional] 
**VmPoolId** | Pointer to **string** | The id of the vm pool linked to the event | [optional] 
**JobId** | Pointer to **string** | The id of the job linked to the event | [optional] 
**SiteId** | Pointer to **string** | The id of the site linked to the event | [optional] 
**Title** | **string** | The title of the event | 
**Message** | **string** | The message of the event | 
**UserEmailAuthenticated** | Pointer to **string** | The email of the user who triggered the event | [optional] 
**Context** | Pointer to **map[string]interface{}** | The context of the event | [optional] 
**OccurredTimestamp** | **string** | The timestamp the event occurred | 
**HttpUserAgent** | Pointer to **string** | The http user agent linked to the event | [optional] 
**RemoteIpAddress** | Pointer to **string** | The remote ip address of the user who triggered the event | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewEvent

`func NewEvent(id string, type_ EventTypes, level EventLevel, visibility EventVisibility, title string, message string, occurredTimestamp string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetUserIdAuthenticated

`func (o *Event) GetUserIdAuthenticated() string`

GetUserIdAuthenticated returns the UserIdAuthenticated field if non-nil, zero value otherwise.

### GetUserIdAuthenticatedOk

`func (o *Event) GetUserIdAuthenticatedOk() (*string, bool)`

GetUserIdAuthenticatedOk returns a tuple with the UserIdAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAuthenticated

`func (o *Event) SetUserIdAuthenticated(v string)`

SetUserIdAuthenticated sets UserIdAuthenticated field to given value.

### HasUserIdAuthenticated

`func (o *Event) HasUserIdAuthenticated() bool`

HasUserIdAuthenticated returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() EventTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*EventTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v EventTypes)`

SetType sets Type field to given value.


### GetLevel

`func (o *Event) GetLevel() EventLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Event) GetLevelOk() (*EventLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Event) SetLevel(v EventLevel)`

SetLevel sets Level field to given value.


### GetVisibility

`func (o *Event) GetVisibility() EventVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Event) GetVisibilityOk() (*EventVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Event) SetVisibility(v EventVisibility)`

SetVisibility sets Visibility field to given value.


### GetInfrastructureId

`func (o *Event) GetInfrastructureId() string`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Event) GetInfrastructureIdOk() (*string, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Event) SetInfrastructureId(v string)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *Event) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetUserId

`func (o *Event) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Event) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Event) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Event) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetServerId

`func (o *Event) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Event) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Event) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *Event) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetNetworkDeviceId

`func (o *Event) GetNetworkDeviceId() string`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *Event) GetNetworkDeviceIdOk() (*string, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *Event) SetNetworkDeviceId(v string)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *Event) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### GetNetworkDeviceControllerId

`func (o *Event) GetNetworkDeviceControllerId() string`

GetNetworkDeviceControllerId returns the NetworkDeviceControllerId field if non-nil, zero value otherwise.

### GetNetworkDeviceControllerIdOk

`func (o *Event) GetNetworkDeviceControllerIdOk() (*string, bool)`

GetNetworkDeviceControllerIdOk returns a tuple with the NetworkDeviceControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceControllerId

`func (o *Event) SetNetworkDeviceControllerId(v string)`

SetNetworkDeviceControllerId sets NetworkDeviceControllerId field to given value.

### HasNetworkDeviceControllerId

`func (o *Event) HasNetworkDeviceControllerId() bool`

HasNetworkDeviceControllerId returns a boolean if a field has been set.

### GetStorageId

`func (o *Event) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *Event) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *Event) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *Event) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetVmPoolId

`func (o *Event) GetVmPoolId() string`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *Event) GetVmPoolIdOk() (*string, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *Event) SetVmPoolId(v string)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *Event) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetJobId

`func (o *Event) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Event) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Event) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *Event) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSiteId

`func (o *Event) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Event) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Event) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Event) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTitle

`func (o *Event) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Event) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Event) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMessage

`func (o *Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Event) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUserEmailAuthenticated

`func (o *Event) GetUserEmailAuthenticated() string`

GetUserEmailAuthenticated returns the UserEmailAuthenticated field if non-nil, zero value otherwise.

### GetUserEmailAuthenticatedOk

`func (o *Event) GetUserEmailAuthenticatedOk() (*string, bool)`

GetUserEmailAuthenticatedOk returns a tuple with the UserEmailAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailAuthenticated

`func (o *Event) SetUserEmailAuthenticated(v string)`

SetUserEmailAuthenticated sets UserEmailAuthenticated field to given value.

### HasUserEmailAuthenticated

`func (o *Event) HasUserEmailAuthenticated() bool`

HasUserEmailAuthenticated returns a boolean if a field has been set.

### GetContext

`func (o *Event) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Event) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Event) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Event) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetOccurredTimestamp

`func (o *Event) GetOccurredTimestamp() string`

GetOccurredTimestamp returns the OccurredTimestamp field if non-nil, zero value otherwise.

### GetOccurredTimestampOk

`func (o *Event) GetOccurredTimestampOk() (*string, bool)`

GetOccurredTimestampOk returns a tuple with the OccurredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTimestamp

`func (o *Event) SetOccurredTimestamp(v string)`

SetOccurredTimestamp sets OccurredTimestamp field to given value.


### GetHttpUserAgent

`func (o *Event) GetHttpUserAgent() string`

GetHttpUserAgent returns the HttpUserAgent field if non-nil, zero value otherwise.

### GetHttpUserAgentOk

`func (o *Event) GetHttpUserAgentOk() (*string, bool)`

GetHttpUserAgentOk returns a tuple with the HttpUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpUserAgent

`func (o *Event) SetHttpUserAgent(v string)`

SetHttpUserAgent sets HttpUserAgent field to given value.

### HasHttpUserAgent

`func (o *Event) HasHttpUserAgent() bool`

HasHttpUserAgent returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *Event) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *Event) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *Event) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *Event) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.

### GetLinks

`func (o *Event) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Event) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Event) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Event) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



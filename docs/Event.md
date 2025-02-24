# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The id of the event | 
**UserIdAuthenticated** | Pointer to **float32** | The id of the user who triggered the event | [optional] 
**Type** | **float32** | The type of the event | 
**Severity** | **string** | The severity of the event | 
**Visibility** | **string** | The visibility of the event | 
**InfrastructureId** | Pointer to **float32** | The id of the infrastructure linked to the event | [optional] 
**UserId** | Pointer to **float32** | The id of the user linked to the event | [optional] 
**ServerId** | Pointer to **float32** | The id of the server linked to the event | [optional] 
**JobId** | Pointer to **string** | The id of the job linked to the event | [optional] 
**SiteId** | Pointer to **float32** | The id of the site linked to the event | [optional] 
**Title** | **string** | The title of the event | 
**Message** | **string** | The message of the event | 
**UserEmailAuthenticated** | Pointer to **string** | The email of the user who triggered the event | [optional] 
**Context** | Pointer to **map[string]interface{}** | The context of the event | [optional] 
**OccurredTimestamp** | **string** | The timestamp the event occurred | 
**HttpUserAgent** | Pointer to **string** | The http user agent linked to the event | [optional] 
**RemoteIpAddress** | Pointer to **string** | The remote ip address of the user who triggered the event | [optional] 

## Methods

### NewEvent

`func NewEvent(id float32, type_ float32, severity string, visibility string, title string, message string, occurredTimestamp string, ) *Event`

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

`func (o *Event) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v float32)`

SetId sets Id field to given value.


### GetUserIdAuthenticated

`func (o *Event) GetUserIdAuthenticated() float32`

GetUserIdAuthenticated returns the UserIdAuthenticated field if non-nil, zero value otherwise.

### GetUserIdAuthenticatedOk

`func (o *Event) GetUserIdAuthenticatedOk() (*float32, bool)`

GetUserIdAuthenticatedOk returns a tuple with the UserIdAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAuthenticated

`func (o *Event) SetUserIdAuthenticated(v float32)`

SetUserIdAuthenticated sets UserIdAuthenticated field to given value.

### HasUserIdAuthenticated

`func (o *Event) HasUserIdAuthenticated() bool`

HasUserIdAuthenticated returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v float32)`

SetType sets Type field to given value.


### GetSeverity

`func (o *Event) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Event) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Event) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetVisibility

`func (o *Event) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Event) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Event) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetInfrastructureId

`func (o *Event) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *Event) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *Event) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *Event) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetUserId

`func (o *Event) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Event) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Event) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Event) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetServerId

`func (o *Event) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Event) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Event) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *Event) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

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

`func (o *Event) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Event) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Event) SetSiteId(v float32)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



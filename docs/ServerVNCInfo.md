# ServerVNCInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSessions** | **float32** | Active VNC sessions | 
**MaxSessions** | **float32** | Max active VNC sessions | 
**Timeout** | **float32** | VNC timeout | 
**Enable** | **string** | VNC enabled | 

## Methods

### NewServerVNCInfo

`func NewServerVNCInfo(activeSessions float32, maxSessions float32, timeout float32, enable string, ) *ServerVNCInfo`

NewServerVNCInfo instantiates a new ServerVNCInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVNCInfoWithDefaults

`func NewServerVNCInfoWithDefaults() *ServerVNCInfo`

NewServerVNCInfoWithDefaults instantiates a new ServerVNCInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSessions

`func (o *ServerVNCInfo) GetActiveSessions() float32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *ServerVNCInfo) GetActiveSessionsOk() (*float32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *ServerVNCInfo) SetActiveSessions(v float32)`

SetActiveSessions sets ActiveSessions field to given value.


### GetMaxSessions

`func (o *ServerVNCInfo) GetMaxSessions() float32`

GetMaxSessions returns the MaxSessions field if non-nil, zero value otherwise.

### GetMaxSessionsOk

`func (o *ServerVNCInfo) GetMaxSessionsOk() (*float32, bool)`

GetMaxSessionsOk returns a tuple with the MaxSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessions

`func (o *ServerVNCInfo) SetMaxSessions(v float32)`

SetMaxSessions sets MaxSessions field to given value.


### GetTimeout

`func (o *ServerVNCInfo) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ServerVNCInfo) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ServerVNCInfo) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.


### GetEnable

`func (o *ServerVNCInfo) GetEnable() string`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ServerVNCInfo) GetEnableOk() (*string, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ServerVNCInfo) SetEnable(v string)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



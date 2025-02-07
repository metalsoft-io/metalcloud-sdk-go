# IscsiBootServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **float32** | Id of the Server | 
**InstanceId** | **float32** | Id of the Instance | 
**InstanceSubdomain** | **string** | Instance subdomain | 
**InstanceSubdomainPermanent** | **string** | Instance permanent subdomain | 
**InfrastructureLabel** | **string** | Infrastructure label | 
**InfrastructureId** | **float32** | Id of the Infrastructure | 
**UserId** | **float32** | Id of the User | 
**UserEmail** | **string** | User email | 
**UserDisplayName** | **string** | User display name | 
**ServerDiskCount** | **float32** | Server disk count | 
**ServerPowerStatus** | **string** | Server power status | 

## Methods

### NewIscsiBootServerInfo

`func NewIscsiBootServerInfo(serverId float32, instanceId float32, instanceSubdomain string, instanceSubdomainPermanent string, infrastructureLabel string, infrastructureId float32, userId float32, userEmail string, userDisplayName string, serverDiskCount float32, serverPowerStatus string, ) *IscsiBootServerInfo`

NewIscsiBootServerInfo instantiates a new IscsiBootServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiBootServerInfoWithDefaults

`func NewIscsiBootServerInfoWithDefaults() *IscsiBootServerInfo`

NewIscsiBootServerInfoWithDefaults instantiates a new IscsiBootServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *IscsiBootServerInfo) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *IscsiBootServerInfo) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *IscsiBootServerInfo) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetInstanceId

`func (o *IscsiBootServerInfo) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *IscsiBootServerInfo) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *IscsiBootServerInfo) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceSubdomain

`func (o *IscsiBootServerInfo) GetInstanceSubdomain() string`

GetInstanceSubdomain returns the InstanceSubdomain field if non-nil, zero value otherwise.

### GetInstanceSubdomainOk

`func (o *IscsiBootServerInfo) GetInstanceSubdomainOk() (*string, bool)`

GetInstanceSubdomainOk returns a tuple with the InstanceSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSubdomain

`func (o *IscsiBootServerInfo) SetInstanceSubdomain(v string)`

SetInstanceSubdomain sets InstanceSubdomain field to given value.


### GetInstanceSubdomainPermanent

`func (o *IscsiBootServerInfo) GetInstanceSubdomainPermanent() string`

GetInstanceSubdomainPermanent returns the InstanceSubdomainPermanent field if non-nil, zero value otherwise.

### GetInstanceSubdomainPermanentOk

`func (o *IscsiBootServerInfo) GetInstanceSubdomainPermanentOk() (*string, bool)`

GetInstanceSubdomainPermanentOk returns a tuple with the InstanceSubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSubdomainPermanent

`func (o *IscsiBootServerInfo) SetInstanceSubdomainPermanent(v string)`

SetInstanceSubdomainPermanent sets InstanceSubdomainPermanent field to given value.


### GetInfrastructureLabel

`func (o *IscsiBootServerInfo) GetInfrastructureLabel() string`

GetInfrastructureLabel returns the InfrastructureLabel field if non-nil, zero value otherwise.

### GetInfrastructureLabelOk

`func (o *IscsiBootServerInfo) GetInfrastructureLabelOk() (*string, bool)`

GetInfrastructureLabelOk returns a tuple with the InfrastructureLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureLabel

`func (o *IscsiBootServerInfo) SetInfrastructureLabel(v string)`

SetInfrastructureLabel sets InfrastructureLabel field to given value.


### GetInfrastructureId

`func (o *IscsiBootServerInfo) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *IscsiBootServerInfo) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *IscsiBootServerInfo) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetUserId

`func (o *IscsiBootServerInfo) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IscsiBootServerInfo) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *IscsiBootServerInfo) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetUserEmail

`func (o *IscsiBootServerInfo) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *IscsiBootServerInfo) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *IscsiBootServerInfo) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserDisplayName

`func (o *IscsiBootServerInfo) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *IscsiBootServerInfo) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *IscsiBootServerInfo) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.


### GetServerDiskCount

`func (o *IscsiBootServerInfo) GetServerDiskCount() float32`

GetServerDiskCount returns the ServerDiskCount field if non-nil, zero value otherwise.

### GetServerDiskCountOk

`func (o *IscsiBootServerInfo) GetServerDiskCountOk() (*float32, bool)`

GetServerDiskCountOk returns a tuple with the ServerDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDiskCount

`func (o *IscsiBootServerInfo) SetServerDiskCount(v float32)`

SetServerDiskCount sets ServerDiskCount field to given value.


### GetServerPowerStatus

`func (o *IscsiBootServerInfo) GetServerPowerStatus() string`

GetServerPowerStatus returns the ServerPowerStatus field if non-nil, zero value otherwise.

### GetServerPowerStatusOk

`func (o *IscsiBootServerInfo) GetServerPowerStatusOk() (*string, bool)`

GetServerPowerStatusOk returns a tuple with the ServerPowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPowerStatus

`func (o *IscsiBootServerInfo) SetServerPowerStatus(v string)`

SetServerPowerStatus sets ServerPowerStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



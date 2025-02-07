# RegisterServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **float32** | The id of the server. | 
**Revision** | **float32** | Revision number | 
**ServerUUID** | Pointer to **string** | The UUID of the server. | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewRegisterServerResponse

`func NewRegisterServerResponse(serverId float32, revision float32, ) *RegisterServerResponse`

NewRegisterServerResponse instantiates a new RegisterServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterServerResponseWithDefaults

`func NewRegisterServerResponseWithDefaults() *RegisterServerResponse`

NewRegisterServerResponseWithDefaults instantiates a new RegisterServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *RegisterServerResponse) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *RegisterServerResponse) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *RegisterServerResponse) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetRevision

`func (o *RegisterServerResponse) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RegisterServerResponse) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RegisterServerResponse) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetServerUUID

`func (o *RegisterServerResponse) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *RegisterServerResponse) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *RegisterServerResponse) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *RegisterServerResponse) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *RegisterServerResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *RegisterServerResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *RegisterServerResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *RegisterServerResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetJobInfo

`func (o *RegisterServerResponse) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *RegisterServerResponse) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *RegisterServerResponse) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *RegisterServerResponse) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### GetLinks

`func (o *RegisterServerResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RegisterServerResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RegisterServerResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RegisterServerResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



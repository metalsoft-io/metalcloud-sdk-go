# ReRegisterServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **int64** | The id of the server. | 
**Revision** | **int64** | Revision number | 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewReRegisterServerResponse

`func NewReRegisterServerResponse(serverId int64, revision int64, ) *ReRegisterServerResponse`

NewReRegisterServerResponse instantiates a new ReRegisterServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReRegisterServerResponseWithDefaults

`func NewReRegisterServerResponseWithDefaults() *ReRegisterServerResponse`

NewReRegisterServerResponseWithDefaults instantiates a new ReRegisterServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ReRegisterServerResponse) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ReRegisterServerResponse) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ReRegisterServerResponse) SetServerId(v int64)`

SetServerId sets ServerId field to given value.


### GetRevision

`func (o *ReRegisterServerResponse) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ReRegisterServerResponse) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ReRegisterServerResponse) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetJobInfo

`func (o *ReRegisterServerResponse) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *ReRegisterServerResponse) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *ReRegisterServerResponse) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *ReRegisterServerResponse) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### GetLinks

`func (o *ReRegisterServerResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ReRegisterServerResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ReRegisterServerResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ReRegisterServerResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



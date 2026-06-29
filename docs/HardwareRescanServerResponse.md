# HardwareRescanServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **int64** | The id of the server. | 
**Revision** | **int64** | Revision number | 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewHardwareRescanServerResponse

`func NewHardwareRescanServerResponse(serverId int64, revision int64, ) *HardwareRescanServerResponse`

NewHardwareRescanServerResponse instantiates a new HardwareRescanServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareRescanServerResponseWithDefaults

`func NewHardwareRescanServerResponseWithDefaults() *HardwareRescanServerResponse`

NewHardwareRescanServerResponseWithDefaults instantiates a new HardwareRescanServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *HardwareRescanServerResponse) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *HardwareRescanServerResponse) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *HardwareRescanServerResponse) SetServerId(v int64)`

SetServerId sets ServerId field to given value.


### GetRevision

`func (o *HardwareRescanServerResponse) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HardwareRescanServerResponse) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HardwareRescanServerResponse) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetJobInfo

`func (o *HardwareRescanServerResponse) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *HardwareRescanServerResponse) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *HardwareRescanServerResponse) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *HardwareRescanServerResponse) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### GetLinks

`func (o *HardwareRescanServerResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HardwareRescanServerResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HardwareRescanServerResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HardwareRescanServerResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



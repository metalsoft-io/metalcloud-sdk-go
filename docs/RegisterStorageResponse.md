# RegisterStorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id of the Storage | 
**JobInfo** | Pointer to [**JobInfo**](JobInfo.md) | Job info | [optional] 

## Methods

### NewRegisterStorageResponse

`func NewRegisterStorageResponse(id float32, ) *RegisterStorageResponse`

NewRegisterStorageResponse instantiates a new RegisterStorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterStorageResponseWithDefaults

`func NewRegisterStorageResponseWithDefaults() *RegisterStorageResponse`

NewRegisterStorageResponseWithDefaults instantiates a new RegisterStorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterStorageResponse) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterStorageResponse) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterStorageResponse) SetId(v float32)`

SetId sets Id field to given value.


### GetJobInfo

`func (o *RegisterStorageResponse) GetJobInfo() JobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *RegisterStorageResponse) GetJobInfoOk() (*JobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *RegisterStorageResponse) SetJobInfo(v JobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *RegisterStorageResponse) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



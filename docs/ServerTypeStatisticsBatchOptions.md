# ServerTypeStatisticsBatchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIdOwner** | Pointer to **float32** | The id of user owning the resources. Used for filtering | [optional] 
**SiteId** | **float32** | The id of the site where the resources are located. Used for filtering | 
**MaximumResultsPerServerType** | Pointer to **float32** | The maximum returned results per server type. | [optional] 
**ServerTypeIds** | Pointer to **[]float32** | The id of the server types to get statistics for. | [optional] 
**InstanceArrayId** | Pointer to **float32** | If specified, treats only the active Instances of the Instance Array as available, instead of all active instances of userIdOwner. | [optional] 

## Methods

### NewServerTypeStatisticsBatchOptions

`func NewServerTypeStatisticsBatchOptions(siteId float32, ) *ServerTypeStatisticsBatchOptions`

NewServerTypeStatisticsBatchOptions instantiates a new ServerTypeStatisticsBatchOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeStatisticsBatchOptionsWithDefaults

`func NewServerTypeStatisticsBatchOptionsWithDefaults() *ServerTypeStatisticsBatchOptions`

NewServerTypeStatisticsBatchOptionsWithDefaults instantiates a new ServerTypeStatisticsBatchOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIdOwner

`func (o *ServerTypeStatisticsBatchOptions) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *ServerTypeStatisticsBatchOptions) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *ServerTypeStatisticsBatchOptions) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *ServerTypeStatisticsBatchOptions) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetSiteId

`func (o *ServerTypeStatisticsBatchOptions) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ServerTypeStatisticsBatchOptions) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ServerTypeStatisticsBatchOptions) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetMaximumResultsPerServerType

`func (o *ServerTypeStatisticsBatchOptions) GetMaximumResultsPerServerType() float32`

GetMaximumResultsPerServerType returns the MaximumResultsPerServerType field if non-nil, zero value otherwise.

### GetMaximumResultsPerServerTypeOk

`func (o *ServerTypeStatisticsBatchOptions) GetMaximumResultsPerServerTypeOk() (*float32, bool)`

GetMaximumResultsPerServerTypeOk returns a tuple with the MaximumResultsPerServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResultsPerServerType

`func (o *ServerTypeStatisticsBatchOptions) SetMaximumResultsPerServerType(v float32)`

SetMaximumResultsPerServerType sets MaximumResultsPerServerType field to given value.

### HasMaximumResultsPerServerType

`func (o *ServerTypeStatisticsBatchOptions) HasMaximumResultsPerServerType() bool`

HasMaximumResultsPerServerType returns a boolean if a field has been set.

### GetServerTypeIds

`func (o *ServerTypeStatisticsBatchOptions) GetServerTypeIds() []float32`

GetServerTypeIds returns the ServerTypeIds field if non-nil, zero value otherwise.

### GetServerTypeIdsOk

`func (o *ServerTypeStatisticsBatchOptions) GetServerTypeIdsOk() (*[]float32, bool)`

GetServerTypeIdsOk returns a tuple with the ServerTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeIds

`func (o *ServerTypeStatisticsBatchOptions) SetServerTypeIds(v []float32)`

SetServerTypeIds sets ServerTypeIds field to given value.

### HasServerTypeIds

`func (o *ServerTypeStatisticsBatchOptions) HasServerTypeIds() bool`

HasServerTypeIds returns a boolean if a field has been set.

### GetInstanceArrayId

`func (o *ServerTypeStatisticsBatchOptions) GetInstanceArrayId() float32`

GetInstanceArrayId returns the InstanceArrayId field if non-nil, zero value otherwise.

### GetInstanceArrayIdOk

`func (o *ServerTypeStatisticsBatchOptions) GetInstanceArrayIdOk() (*float32, bool)`

GetInstanceArrayIdOk returns a tuple with the InstanceArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrayId

`func (o *ServerTypeStatisticsBatchOptions) SetInstanceArrayId(v float32)`

SetInstanceArrayId sets InstanceArrayId field to given value.

### HasInstanceArrayId

`func (o *ServerTypeStatisticsBatchOptions) HasInstanceArrayId() bool`

HasInstanceArrayId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



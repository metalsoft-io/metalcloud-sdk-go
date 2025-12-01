# DiscoveryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discover** | **[]string** | Types of discovery to perform | 
**ReturnData** | Pointer to **bool** | Whether to return the discovered data in the response | [optional] 
**PersistData** | **bool** | Whether to persist the discovered data | 

## Methods

### NewDiscoveryQuery

`func NewDiscoveryQuery(discover []string, persistData bool, ) *DiscoveryQuery`

NewDiscoveryQuery instantiates a new DiscoveryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryQueryWithDefaults

`func NewDiscoveryQueryWithDefaults() *DiscoveryQuery`

NewDiscoveryQueryWithDefaults instantiates a new DiscoveryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscover

`func (o *DiscoveryQuery) GetDiscover() []string`

GetDiscover returns the Discover field if non-nil, zero value otherwise.

### GetDiscoverOk

`func (o *DiscoveryQuery) GetDiscoverOk() (*[]string, bool)`

GetDiscoverOk returns a tuple with the Discover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscover

`func (o *DiscoveryQuery) SetDiscover(v []string)`

SetDiscover sets Discover field to given value.


### GetReturnData

`func (o *DiscoveryQuery) GetReturnData() bool`

GetReturnData returns the ReturnData field if non-nil, zero value otherwise.

### GetReturnDataOk

`func (o *DiscoveryQuery) GetReturnDataOk() (*bool, bool)`

GetReturnDataOk returns a tuple with the ReturnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnData

`func (o *DiscoveryQuery) SetReturnData(v bool)`

SetReturnData sets ReturnData field to given value.

### HasReturnData

`func (o *DiscoveryQuery) HasReturnData() bool`

HasReturnData returns a boolean if a field has been set.

### GetPersistData

`func (o *DiscoveryQuery) GetPersistData() bool`

GetPersistData returns the PersistData field if non-nil, zero value otherwise.

### GetPersistDataOk

`func (o *DiscoveryQuery) GetPersistDataOk() (*bool, bool)`

GetPersistDataOk returns a tuple with the PersistData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistData

`func (o *DiscoveryQuery) SetPersistData(v bool)`

SetPersistData sets PersistData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



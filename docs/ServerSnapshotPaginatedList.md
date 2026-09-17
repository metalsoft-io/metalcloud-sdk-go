# ServerSnapshotPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServerSnapshot**](ServerSnapshot.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewServerSnapshotPaginatedList

`func NewServerSnapshotPaginatedList(data []ServerSnapshot, meta PaginatedResponseMeta, ) *ServerSnapshotPaginatedList`

NewServerSnapshotPaginatedList instantiates a new ServerSnapshotPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSnapshotPaginatedListWithDefaults

`func NewServerSnapshotPaginatedListWithDefaults() *ServerSnapshotPaginatedList`

NewServerSnapshotPaginatedListWithDefaults instantiates a new ServerSnapshotPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServerSnapshotPaginatedList) GetData() []ServerSnapshot`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServerSnapshotPaginatedList) GetDataOk() (*[]ServerSnapshot, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServerSnapshotPaginatedList) SetData(v []ServerSnapshot)`

SetData sets Data field to given value.


### GetMeta

`func (o *ServerSnapshotPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerSnapshotPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerSnapshotPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ServerSnapshotPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerSnapshotPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerSnapshotPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerSnapshotPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



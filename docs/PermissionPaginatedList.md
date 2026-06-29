# PermissionPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Permission**](Permission.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewPermissionPaginatedList

`func NewPermissionPaginatedList(data []Permission, meta PaginatedResponseMeta, ) *PermissionPaginatedList`

NewPermissionPaginatedList instantiates a new PermissionPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionPaginatedListWithDefaults

`func NewPermissionPaginatedListWithDefaults() *PermissionPaginatedList`

NewPermissionPaginatedListWithDefaults instantiates a new PermissionPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PermissionPaginatedList) GetData() []Permission`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PermissionPaginatedList) GetDataOk() (*[]Permission, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PermissionPaginatedList) SetData(v []Permission)`

SetData sets Data field to given value.


### GetMeta

`func (o *PermissionPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PermissionPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PermissionPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PermissionPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PermissionPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PermissionPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PermissionPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



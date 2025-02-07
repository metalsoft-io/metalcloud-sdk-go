# ServerInstanceProfilePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServerInstanceProfile**](ServerInstanceProfile.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewServerInstanceProfilePaginatedList

`func NewServerInstanceProfilePaginatedList(data []ServerInstanceProfile, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *ServerInstanceProfilePaginatedList`

NewServerInstanceProfilePaginatedList instantiates a new ServerInstanceProfilePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceProfilePaginatedListWithDefaults

`func NewServerInstanceProfilePaginatedListWithDefaults() *ServerInstanceProfilePaginatedList`

NewServerInstanceProfilePaginatedListWithDefaults instantiates a new ServerInstanceProfilePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServerInstanceProfilePaginatedList) GetData() []ServerInstanceProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServerInstanceProfilePaginatedList) GetDataOk() (*[]ServerInstanceProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServerInstanceProfilePaginatedList) SetData(v []ServerInstanceProfile)`

SetData sets Data field to given value.


### GetMeta

`func (o *ServerInstanceProfilePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerInstanceProfilePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerInstanceProfilePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ServerInstanceProfilePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceProfilePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceProfilePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



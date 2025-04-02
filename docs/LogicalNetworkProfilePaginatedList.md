# LogicalNetworkProfilePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]LogicalNetworkProfile**](LogicalNetworkProfile.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewLogicalNetworkProfilePaginatedList

`func NewLogicalNetworkProfilePaginatedList(data []LogicalNetworkProfile, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *LogicalNetworkProfilePaginatedList`

NewLogicalNetworkProfilePaginatedList instantiates a new LogicalNetworkProfilePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkProfilePaginatedListWithDefaults

`func NewLogicalNetworkProfilePaginatedListWithDefaults() *LogicalNetworkProfilePaginatedList`

NewLogicalNetworkProfilePaginatedListWithDefaults instantiates a new LogicalNetworkProfilePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LogicalNetworkProfilePaginatedList) GetData() []LogicalNetworkProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogicalNetworkProfilePaginatedList) GetDataOk() (*[]LogicalNetworkProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogicalNetworkProfilePaginatedList) SetData(v []LogicalNetworkProfile)`

SetData sets Data field to given value.


### GetMeta

`func (o *LogicalNetworkProfilePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LogicalNetworkProfilePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LogicalNetworkProfilePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *LogicalNetworkProfilePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogicalNetworkProfilePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogicalNetworkProfilePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



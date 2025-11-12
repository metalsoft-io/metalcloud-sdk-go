# PaginatedLogicalNetworkProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]LogicalNetworkProfile1Inner**](LogicalNetworkProfile1Inner.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewPaginatedLogicalNetworkProfile

`func NewPaginatedLogicalNetworkProfile(data []LogicalNetworkProfile1Inner, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *PaginatedLogicalNetworkProfile`

NewPaginatedLogicalNetworkProfile instantiates a new PaginatedLogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedLogicalNetworkProfileWithDefaults

`func NewPaginatedLogicalNetworkProfileWithDefaults() *PaginatedLogicalNetworkProfile`

NewPaginatedLogicalNetworkProfileWithDefaults instantiates a new PaginatedLogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginatedLogicalNetworkProfile) GetData() []LogicalNetworkProfile1Inner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedLogicalNetworkProfile) GetDataOk() (*[]LogicalNetworkProfile1Inner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedLogicalNetworkProfile) SetData(v []LogicalNetworkProfile1Inner)`

SetData sets Data field to given value.


### GetMeta

`func (o *PaginatedLogicalNetworkProfile) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedLogicalNetworkProfile) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedLogicalNetworkProfile) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PaginatedLogicalNetworkProfile) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedLogicalNetworkProfile) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedLogicalNetworkProfile) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



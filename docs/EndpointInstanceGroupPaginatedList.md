# EndpointInstanceGroupPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EndpointInstanceGroup**](EndpointInstanceGroup.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewEndpointInstanceGroupPaginatedList

`func NewEndpointInstanceGroupPaginatedList(data []EndpointInstanceGroup, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *EndpointInstanceGroupPaginatedList`

NewEndpointInstanceGroupPaginatedList instantiates a new EndpointInstanceGroupPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceGroupPaginatedListWithDefaults

`func NewEndpointInstanceGroupPaginatedListWithDefaults() *EndpointInstanceGroupPaginatedList`

NewEndpointInstanceGroupPaginatedListWithDefaults instantiates a new EndpointInstanceGroupPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EndpointInstanceGroupPaginatedList) GetData() []EndpointInstanceGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EndpointInstanceGroupPaginatedList) GetDataOk() (*[]EndpointInstanceGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EndpointInstanceGroupPaginatedList) SetData(v []EndpointInstanceGroup)`

SetData sets Data field to given value.


### GetMeta

`func (o *EndpointInstanceGroupPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EndpointInstanceGroupPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EndpointInstanceGroupPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *EndpointInstanceGroupPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointInstanceGroupPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointInstanceGroupPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



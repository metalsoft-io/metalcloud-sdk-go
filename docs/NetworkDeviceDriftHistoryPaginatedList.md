# NetworkDeviceDriftHistoryPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkDeviceDriftHistory**](NetworkDeviceDriftHistory.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewNetworkDeviceDriftHistoryPaginatedList

`func NewNetworkDeviceDriftHistoryPaginatedList(data []NetworkDeviceDriftHistory, meta PaginatedResponseMeta, ) *NetworkDeviceDriftHistoryPaginatedList`

NewNetworkDeviceDriftHistoryPaginatedList instantiates a new NetworkDeviceDriftHistoryPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDriftHistoryPaginatedListWithDefaults

`func NewNetworkDeviceDriftHistoryPaginatedListWithDefaults() *NetworkDeviceDriftHistoryPaginatedList`

NewNetworkDeviceDriftHistoryPaginatedListWithDefaults instantiates a new NetworkDeviceDriftHistoryPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkDeviceDriftHistoryPaginatedList) GetData() []NetworkDeviceDriftHistory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkDeviceDriftHistoryPaginatedList) GetDataOk() (*[]NetworkDeviceDriftHistory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkDeviceDriftHistoryPaginatedList) SetData(v []NetworkDeviceDriftHistory)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkDeviceDriftHistoryPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkDeviceDriftHistoryPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkDeviceDriftHistoryPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *NetworkDeviceDriftHistoryPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceDriftHistoryPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceDriftHistoryPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceDriftHistoryPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



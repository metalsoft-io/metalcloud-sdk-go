# VMInstanceGroupPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VMInstanceGroup**](VMInstanceGroup.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewVMInstanceGroupPaginatedList

`func NewVMInstanceGroupPaginatedList(data []VMInstanceGroup, meta PaginatedResponseMeta, ) *VMInstanceGroupPaginatedList`

NewVMInstanceGroupPaginatedList instantiates a new VMInstanceGroupPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupPaginatedListWithDefaults

`func NewVMInstanceGroupPaginatedListWithDefaults() *VMInstanceGroupPaginatedList`

NewVMInstanceGroupPaginatedListWithDefaults instantiates a new VMInstanceGroupPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VMInstanceGroupPaginatedList) GetData() []VMInstanceGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VMInstanceGroupPaginatedList) GetDataOk() (*[]VMInstanceGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VMInstanceGroupPaginatedList) SetData(v []VMInstanceGroup)`

SetData sets Data field to given value.


### GetMeta

`func (o *VMInstanceGroupPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VMInstanceGroupPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VMInstanceGroupPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *VMInstanceGroupPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMInstanceGroupPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMInstanceGroupPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VMInstanceGroupPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



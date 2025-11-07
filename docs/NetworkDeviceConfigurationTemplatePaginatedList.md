# NetworkDeviceConfigurationTemplatePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkDeviceConfigurationTemplate**](NetworkDeviceConfigurationTemplate.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewNetworkDeviceConfigurationTemplatePaginatedList

`func NewNetworkDeviceConfigurationTemplatePaginatedList(data []NetworkDeviceConfigurationTemplate, meta PaginatedResponseMeta, ) *NetworkDeviceConfigurationTemplatePaginatedList`

NewNetworkDeviceConfigurationTemplatePaginatedList instantiates a new NetworkDeviceConfigurationTemplatePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceConfigurationTemplatePaginatedListWithDefaults

`func NewNetworkDeviceConfigurationTemplatePaginatedListWithDefaults() *NetworkDeviceConfigurationTemplatePaginatedList`

NewNetworkDeviceConfigurationTemplatePaginatedListWithDefaults instantiates a new NetworkDeviceConfigurationTemplatePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) GetData() []NetworkDeviceConfigurationTemplate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) GetDataOk() (*[]NetworkDeviceConfigurationTemplate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) SetData(v []NetworkDeviceConfigurationTemplate)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceConfigurationTemplatePaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



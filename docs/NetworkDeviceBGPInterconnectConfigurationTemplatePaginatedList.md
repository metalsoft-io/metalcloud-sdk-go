# NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkDeviceBGPInterconnectConfigurationTemplate**](NetworkDeviceBGPInterconnectConfigurationTemplate.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewNetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList

`func NewNetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList(data []NetworkDeviceBGPInterconnectConfigurationTemplate, meta PaginatedResponseMeta, ) *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList`

NewNetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList instantiates a new NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceBGPInterconnectConfigurationTemplatePaginatedListWithDefaults

`func NewNetworkDeviceBGPInterconnectConfigurationTemplatePaginatedListWithDefaults() *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList`

NewNetworkDeviceBGPInterconnectConfigurationTemplatePaginatedListWithDefaults instantiates a new NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) GetData() []NetworkDeviceBGPInterconnectConfigurationTemplate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) GetDataOk() (*[]NetworkDeviceBGPInterconnectConfigurationTemplate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) SetData(v []NetworkDeviceBGPInterconnectConfigurationTemplate)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceBGPInterconnectConfigurationTemplatePaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



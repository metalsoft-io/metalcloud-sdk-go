# NetworkDeviceDefaultSecretsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkDeviceDefaultSecrets**](NetworkDeviceDefaultSecrets.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewNetworkDeviceDefaultSecretsPaginatedList

`func NewNetworkDeviceDefaultSecretsPaginatedList(data []NetworkDeviceDefaultSecrets, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *NetworkDeviceDefaultSecretsPaginatedList`

NewNetworkDeviceDefaultSecretsPaginatedList instantiates a new NetworkDeviceDefaultSecretsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDefaultSecretsPaginatedListWithDefaults

`func NewNetworkDeviceDefaultSecretsPaginatedListWithDefaults() *NetworkDeviceDefaultSecretsPaginatedList`

NewNetworkDeviceDefaultSecretsPaginatedListWithDefaults instantiates a new NetworkDeviceDefaultSecretsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkDeviceDefaultSecretsPaginatedList) GetData() []NetworkDeviceDefaultSecrets`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkDeviceDefaultSecretsPaginatedList) GetDataOk() (*[]NetworkDeviceDefaultSecrets, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkDeviceDefaultSecretsPaginatedList) SetData(v []NetworkDeviceDefaultSecrets)`

SetData sets Data field to given value.


### GetMeta

`func (o *NetworkDeviceDefaultSecretsPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkDeviceDefaultSecretsPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkDeviceDefaultSecretsPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *NetworkDeviceDefaultSecretsPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceDefaultSecretsPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceDefaultSecretsPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



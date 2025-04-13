# SharedDrivePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SharedDrive**](SharedDrive.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewSharedDrivePaginatedList

`func NewSharedDrivePaginatedList(data []SharedDrive, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *SharedDrivePaginatedList`

NewSharedDrivePaginatedList instantiates a new SharedDrivePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDrivePaginatedListWithDefaults

`func NewSharedDrivePaginatedListWithDefaults() *SharedDrivePaginatedList`

NewSharedDrivePaginatedListWithDefaults instantiates a new SharedDrivePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SharedDrivePaginatedList) GetData() []SharedDrive`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SharedDrivePaginatedList) GetDataOk() (*[]SharedDrive, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SharedDrivePaginatedList) SetData(v []SharedDrive)`

SetData sets Data field to given value.


### GetMeta

`func (o *SharedDrivePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SharedDrivePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SharedDrivePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *SharedDrivePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SharedDrivePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SharedDrivePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



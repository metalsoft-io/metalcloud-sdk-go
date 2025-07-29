# ServerDefaultCredentialsPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServerDefaultCredentials**](ServerDefaultCredentials.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewServerDefaultCredentialsPaginatedList

`func NewServerDefaultCredentialsPaginatedList(data []ServerDefaultCredentials, meta PaginatedResponseMeta, ) *ServerDefaultCredentialsPaginatedList`

NewServerDefaultCredentialsPaginatedList instantiates a new ServerDefaultCredentialsPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDefaultCredentialsPaginatedListWithDefaults

`func NewServerDefaultCredentialsPaginatedListWithDefaults() *ServerDefaultCredentialsPaginatedList`

NewServerDefaultCredentialsPaginatedListWithDefaults instantiates a new ServerDefaultCredentialsPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServerDefaultCredentialsPaginatedList) GetData() []ServerDefaultCredentials`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServerDefaultCredentialsPaginatedList) GetDataOk() (*[]ServerDefaultCredentials, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServerDefaultCredentialsPaginatedList) SetData(v []ServerDefaultCredentials)`

SetData sets Data field to given value.


### GetMeta

`func (o *ServerDefaultCredentialsPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerDefaultCredentialsPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerDefaultCredentialsPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ServerDefaultCredentialsPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerDefaultCredentialsPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerDefaultCredentialsPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerDefaultCredentialsPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



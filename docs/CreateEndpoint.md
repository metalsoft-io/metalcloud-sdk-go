# CreateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **float32** | The ID of the site where the entity is located. | 
**Name** | **string** | The endpoint name | 
**Label** | Pointer to **string** | The endpoint unique label | [optional] 
**ExternalId** | Pointer to **string** | The external ID of the endpoint, should be unique across the system. Usually either an ethernet MAC address or a UUID. | [optional] 

## Methods

### NewCreateEndpoint

`func NewCreateEndpoint(siteId float32, name string, ) *CreateEndpoint`

NewCreateEndpoint instantiates a new CreateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointWithDefaults

`func NewCreateEndpointWithDefaults() *CreateEndpoint`

NewCreateEndpointWithDefaults instantiates a new CreateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *CreateEndpoint) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateEndpoint) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateEndpoint) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetName

`func (o *CreateEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CreateEndpoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateEndpoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateEndpoint) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateEndpoint) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateEndpoint) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateEndpoint) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateEndpoint) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateEndpoint) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



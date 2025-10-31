# CreateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the endpoint, should be unique across the system. Usually either an ethernet MAC address or a UUID. | [optional] 
**SiteId** | **int32** | The ID of the site where the entity is located. | 
**Name** | **string** | The endpoint name | 
**Label** | **string** | The endpoint unique label | 
**EndpointInterfaces** | Pointer to [**[]CreateEndpointInterface**](CreateEndpointInterface.md) | The endpoint interfaces associated with this endpoint | [optional] 

## Methods

### NewCreateEndpoint

`func NewCreateEndpoint(siteId int32, name string, label string, ) *CreateEndpoint`

NewCreateEndpoint instantiates a new CreateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointWithDefaults

`func NewCreateEndpointWithDefaults() *CreateEndpoint`

NewCreateEndpointWithDefaults instantiates a new CreateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetSiteId

`func (o *CreateEndpoint) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateEndpoint) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateEndpoint) SetSiteId(v int32)`

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


### GetEndpointInterfaces

`func (o *CreateEndpoint) GetEndpointInterfaces() []CreateEndpointInterface`

GetEndpointInterfaces returns the EndpointInterfaces field if non-nil, zero value otherwise.

### GetEndpointInterfacesOk

`func (o *CreateEndpoint) GetEndpointInterfacesOk() (*[]CreateEndpointInterface, bool)`

GetEndpointInterfacesOk returns a tuple with the EndpointInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointInterfaces

`func (o *CreateEndpoint) SetEndpointInterfaces(v []CreateEndpointInterface)`

SetEndpointInterfaces sets EndpointInterfaces field to given value.

### HasEndpointInterfaces

`func (o *CreateEndpoint) HasEndpointInterfaces() bool`

HasEndpointInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The endpoint name | 
**Label** | Pointer to **string** | The endpoint unique label | [optional] 
**ExternalId** | Pointer to **string** | The external ID of the endpoint, should be unique across the system. Usually either an ethernet MAC address or a UUID. | [optional] 

## Methods

### NewUpdateEndpoint

`func NewUpdateEndpoint(name string, ) *UpdateEndpoint`

NewUpdateEndpoint instantiates a new UpdateEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEndpointWithDefaults

`func NewUpdateEndpointWithDefaults() *UpdateEndpoint`

NewUpdateEndpointWithDefaults instantiates a new UpdateEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *UpdateEndpoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateEndpoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateEndpoint) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateEndpoint) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateEndpoint) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateEndpoint) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateEndpoint) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateEndpoint) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



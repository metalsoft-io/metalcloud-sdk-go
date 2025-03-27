# CreateNetworkEndpointGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **float32** | The ID of the site where the entity is located. | [optional] 
**Name** | **string** | The name of the network endpoint group | 

## Methods

### NewCreateNetworkEndpointGroup

`func NewCreateNetworkEndpointGroup(name string, ) *CreateNetworkEndpointGroup`

NewCreateNetworkEndpointGroup instantiates a new CreateNetworkEndpointGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkEndpointGroupWithDefaults

`func NewCreateNetworkEndpointGroupWithDefaults() *CreateNetworkEndpointGroup`

NewCreateNetworkEndpointGroupWithDefaults instantiates a new CreateNetworkEndpointGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *CreateNetworkEndpointGroup) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateNetworkEndpointGroup) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateNetworkEndpointGroup) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CreateNetworkEndpointGroup) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkEndpointGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkEndpointGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkEndpointGroup) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



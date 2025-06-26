# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **float32** | The ID of the site where the entity is located. | 
**Revision** | **string** | Revision number of the entity | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | The network endpoint id. | 
**Name** | **string** | The endpoint name | 
**Label** | **string** | The endpoint unique label | 
**ExternalId** | Pointer to **string** | The external ID of the endpoint, should be unique across the system. Usually either an ethernet MAC address or a UUID. | [optional] 
**EndpointInterfaces** | Pointer to [**[]EndpointInterface**](EndpointInterface.md) | The endpoint interfaces associated with this endpoint | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint(siteId float32, revision string, createdTimestamp time.Time, updatedTimestamp time.Time, id string, name string, label string, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *Endpoint) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Endpoint) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Endpoint) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetRevision

`func (o *Endpoint) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Endpoint) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Endpoint) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetCreatedTimestamp

`func (o *Endpoint) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Endpoint) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Endpoint) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *Endpoint) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Endpoint) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Endpoint) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *Endpoint) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Endpoint) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Endpoint) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Endpoint) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Endpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Endpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Endpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Endpoint) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Endpoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Endpoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Endpoint) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetExternalId

`func (o *Endpoint) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Endpoint) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Endpoint) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Endpoint) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEndpointInterfaces

`func (o *Endpoint) GetEndpointInterfaces() []EndpointInterface`

GetEndpointInterfaces returns the EndpointInterfaces field if non-nil, zero value otherwise.

### GetEndpointInterfacesOk

`func (o *Endpoint) GetEndpointInterfacesOk() (*[]EndpointInterface, bool)`

GetEndpointInterfacesOk returns a tuple with the EndpointInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointInterfaces

`func (o *Endpoint) SetEndpointInterfaces(v []EndpointInterface)`

SetEndpointInterfaces sets EndpointInterfaces field to given value.

### HasEndpointInterfaces

`func (o *Endpoint) HasEndpointInterfaces() bool`

HasEndpointInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



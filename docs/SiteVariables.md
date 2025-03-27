# SiteVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The site ID | 
**Revision** | **float32** | Revision number | 
**Slug** | Pointer to **string** | The site unique slug | [optional] 
**Name** | **string** | The site name | 
**Location** | Pointer to [**Location**](Location.md) | Location details | [optional] 
**OwnerId** | Pointer to **int32** | ID of the site owner | [optional] 

## Methods

### NewSiteVariables

`func NewSiteVariables(id int32, revision float32, name string, ) *SiteVariables`

NewSiteVariables instantiates a new SiteVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteVariablesWithDefaults

`func NewSiteVariablesWithDefaults() *SiteVariables`

NewSiteVariablesWithDefaults instantiates a new SiteVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteVariables) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteVariables) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteVariables) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *SiteVariables) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SiteVariables) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SiteVariables) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetSlug

`func (o *SiteVariables) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SiteVariables) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SiteVariables) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *SiteVariables) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *SiteVariables) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteVariables) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteVariables) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *SiteVariables) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteVariables) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteVariables) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteVariables) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOwnerId

`func (o *SiteVariables) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SiteVariables) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SiteVariables) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SiteVariables) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



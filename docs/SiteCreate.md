# SiteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | The site unique slug | 
**Name** | **string** | The site name | 
**Location** | Pointer to [**Location**](Location.md) | Location details | [optional] 
**IsHidden** | Pointer to **bool** | True if the site is hidden | [optional] 
**IsInMaintenance** | Pointer to **bool** | True if the site is in maintenance mode | [optional] 

## Methods

### NewSiteCreate

`func NewSiteCreate(slug string, name string, ) *SiteCreate`

NewSiteCreate instantiates a new SiteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteCreateWithDefaults

`func NewSiteCreateWithDefaults() *SiteCreate`

NewSiteCreateWithDefaults instantiates a new SiteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *SiteCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SiteCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SiteCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *SiteCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *SiteCreate) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteCreate) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteCreate) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteCreate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIsHidden

`func (o *SiteCreate) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *SiteCreate) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *SiteCreate) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *SiteCreate) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsInMaintenance

`func (o *SiteCreate) GetIsInMaintenance() bool`

GetIsInMaintenance returns the IsInMaintenance field if non-nil, zero value otherwise.

### GetIsInMaintenanceOk

`func (o *SiteCreate) GetIsInMaintenanceOk() (*bool, bool)`

GetIsInMaintenanceOk returns a tuple with the IsInMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInMaintenance

`func (o *SiteCreate) SetIsInMaintenance(v bool)`

SetIsInMaintenance sets IsInMaintenance field to given value.

### HasIsInMaintenance

`func (o *SiteCreate) HasIsInMaintenance() bool`

HasIsInMaintenance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



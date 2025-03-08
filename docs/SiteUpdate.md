# SiteUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | Pointer to **string** | The site unique slug | [optional] 
**Name** | Pointer to **string** | The site name | [optional] 
**Location** | Pointer to [**Location**](Location.md) | Location details | [optional] 
**IsHidden** | Pointer to **bool** | True if the site is hidden | [optional] 
**IsInMaintenance** | Pointer to **bool** | True if the site is in maintenance mode | [optional] 

## Methods

### NewSiteUpdate

`func NewSiteUpdate() *SiteUpdate`

NewSiteUpdate instantiates a new SiteUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteUpdateWithDefaults

`func NewSiteUpdateWithDefaults() *SiteUpdate`

NewSiteUpdateWithDefaults instantiates a new SiteUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *SiteUpdate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SiteUpdate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SiteUpdate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *SiteUpdate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *SiteUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *SiteUpdate) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteUpdate) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteUpdate) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteUpdate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIsHidden

`func (o *SiteUpdate) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *SiteUpdate) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *SiteUpdate) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *SiteUpdate) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsInMaintenance

`func (o *SiteUpdate) GetIsInMaintenance() bool`

GetIsInMaintenance returns the IsInMaintenance field if non-nil, zero value otherwise.

### GetIsInMaintenanceOk

`func (o *SiteUpdate) GetIsInMaintenanceOk() (*bool, bool)`

GetIsInMaintenanceOk returns a tuple with the IsInMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInMaintenance

`func (o *SiteUpdate) SetIsInMaintenance(v bool)`

SetIsInMaintenance sets IsInMaintenance field to given value.

### HasIsInMaintenance

`func (o *SiteUpdate) HasIsInMaintenance() bool`

HasIsInMaintenance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



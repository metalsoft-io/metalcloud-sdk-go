# SiteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | The site unique slug | 
**Name** | **string** | The site name | 
**IsHidden** | Pointer to **bool** | True if the site is hidden | [optional] 
**IsInMaintenance** | Pointer to **bool** | True if the site is in maintenance mode | [optional] 
**Config** | Pointer to [**SiteConfig**](SiteConfig.md) | Site configuration | [optional] 

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

### GetConfig

`func (o *SiteCreate) GetConfig() SiteConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SiteCreate) GetConfigOk() (*SiteConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SiteCreate) SetConfig(v SiteConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SiteCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



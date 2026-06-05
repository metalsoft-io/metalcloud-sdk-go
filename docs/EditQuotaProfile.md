# EditQuotaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Updated name for the quota profile | [optional] 
**Description** | Pointer to **string** | Updated description | [optional] 
**Limits** | Pointer to [**PatchQuotaProfileLimitsDto**](PatchQuotaProfileLimitsDto.md) | Partial set of limits to update. Only provided fields are changed, the rest remain unchanged. | [optional] 

## Methods

### NewEditQuotaProfile

`func NewEditQuotaProfile() *EditQuotaProfile`

NewEditQuotaProfile instantiates a new EditQuotaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditQuotaProfileWithDefaults

`func NewEditQuotaProfileWithDefaults() *EditQuotaProfile`

NewEditQuotaProfileWithDefaults instantiates a new EditQuotaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditQuotaProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditQuotaProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditQuotaProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditQuotaProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditQuotaProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditQuotaProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditQuotaProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditQuotaProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimits

`func (o *EditQuotaProfile) GetLimits() PatchQuotaProfileLimitsDto`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *EditQuotaProfile) GetLimitsOk() (*PatchQuotaProfileLimitsDto, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *EditQuotaProfile) SetLimits(v PatchQuotaProfileLimitsDto)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *EditQuotaProfile) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



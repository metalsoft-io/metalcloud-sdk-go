# CreateQuotaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Optional custom ID. Must contain only lowercase letters, digits, hyphens, and underscores. If omitted, derived from the name. | [optional] 
**Name** | **string** | Human-readable name for the new quota profile | 
**Description** | Pointer to **string** | Optional description | [optional] 
**Limits** | Pointer to [**PatchQuotaProfileLimitsDto**](PatchQuotaProfileLimitsDto.md) | Partial set of limits for the new quota profile. Omitted fields default to the values from the default profile. | [optional] 

## Methods

### NewCreateQuotaProfile

`func NewCreateQuotaProfile(name string, ) *CreateQuotaProfile`

NewCreateQuotaProfile instantiates a new CreateQuotaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuotaProfileWithDefaults

`func NewCreateQuotaProfileWithDefaults() *CreateQuotaProfile`

NewCreateQuotaProfileWithDefaults instantiates a new CreateQuotaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateQuotaProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateQuotaProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateQuotaProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateQuotaProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateQuotaProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateQuotaProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateQuotaProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateQuotaProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateQuotaProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateQuotaProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateQuotaProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimits

`func (o *CreateQuotaProfile) GetLimits() PatchQuotaProfileLimitsDto`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CreateQuotaProfile) GetLimitsOk() (*PatchQuotaProfileLimitsDto, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CreateQuotaProfile) SetLimits(v PatchQuotaProfileLimitsDto)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CreateQuotaProfile) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



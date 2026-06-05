# QuotaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the quota profile | 
**Name** | **string** | Human-readable name of the quota profile | 
**Description** | Pointer to **string** | Optional description | [optional] 
**Limits** | [**QuotaProfileLimits**](QuotaProfileLimits.md) | Resource limits defined by this quota profile | 

## Methods

### NewQuotaProfile

`func NewQuotaProfile(id string, name string, limits QuotaProfileLimits, ) *QuotaProfile`

NewQuotaProfile instantiates a new QuotaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaProfileWithDefaults

`func NewQuotaProfileWithDefaults() *QuotaProfile`

NewQuotaProfileWithDefaults instantiates a new QuotaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuotaProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *QuotaProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuotaProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuotaProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *QuotaProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotaProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotaProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotaProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimits

`func (o *QuotaProfile) GetLimits() QuotaProfileLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *QuotaProfile) GetLimitsOk() (*QuotaProfileLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *QuotaProfile) SetLimits(v QuotaProfileLimits)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



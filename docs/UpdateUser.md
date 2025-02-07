# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUrl** | Pointer to **string** | The redirect URL for the user | [optional] 
**DisplayName** | Pointer to **string** | The new display name of the user | [optional] 
**Email** | Pointer to **string** | The new login e-mail of the user | [optional] 
**EmailStatus** | Pointer to **string** | The new e-mail status of the user | [optional] 
**Language** | Pointer to **string** | The new language of the user | [optional] 
**Brand** | Pointer to **string** | The new brand of the user | [optional] 
**IsBrandManager** | Pointer to **bool** | The new brand manager status of the user | [optional] 
**Delegates** | Pointer to [**[]UserDelegate**](UserDelegate.md) | The new user delegates of the user. | [optional] 
**IsBlocked** | Pointer to **bool** | The new blocked status of the user | [optional] 
**Suspension** | Pointer to [**UserSuspend**](UserSuspend.md) |  | [optional] 
**AccessLevel** | Pointer to **string** | The new access level of the user | [optional] 
**Promotions** | Pointer to [**[]UserPromotion**](UserPromotion.md) | The new promotions of the user. | [optional] 
**ExperimentalTags** | Pointer to [**[]UserExperimentalTag**](UserExperimentalTag.md) | The new experimental tags of the user. | [optional] 
**Permissions** | Pointer to [**UserResourcePermissions**](UserResourcePermissions.md) |  | [optional] 
**IsTestAccount** | Pointer to **bool** | Whether the user account is a test one. | [optional] 
**IsBillable** | Pointer to **bool** | Whether the user account is billable. | [optional] 

## Methods

### NewUpdateUser

`func NewUpdateUser() *UpdateUser`

NewUpdateUser instantiates a new UpdateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserWithDefaults

`func NewUpdateUserWithDefaults() *UpdateUser`

NewUpdateUserWithDefaults instantiates a new UpdateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUrl

`func (o *UpdateUser) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateUser) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateUser) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateUser) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailStatus

`func (o *UpdateUser) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *UpdateUser) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *UpdateUser) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.

### HasEmailStatus

`func (o *UpdateUser) HasEmailStatus() bool`

HasEmailStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *UpdateUser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateUser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateUser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdateUser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetBrand

`func (o *UpdateUser) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UpdateUser) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UpdateUser) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *UpdateUser) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetIsBrandManager

`func (o *UpdateUser) GetIsBrandManager() bool`

GetIsBrandManager returns the IsBrandManager field if non-nil, zero value otherwise.

### GetIsBrandManagerOk

`func (o *UpdateUser) GetIsBrandManagerOk() (*bool, bool)`

GetIsBrandManagerOk returns a tuple with the IsBrandManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandManager

`func (o *UpdateUser) SetIsBrandManager(v bool)`

SetIsBrandManager sets IsBrandManager field to given value.

### HasIsBrandManager

`func (o *UpdateUser) HasIsBrandManager() bool`

HasIsBrandManager returns a boolean if a field has been set.

### GetDelegates

`func (o *UpdateUser) GetDelegates() []UserDelegate`

GetDelegates returns the Delegates field if non-nil, zero value otherwise.

### GetDelegatesOk

`func (o *UpdateUser) GetDelegatesOk() (*[]UserDelegate, bool)`

GetDelegatesOk returns a tuple with the Delegates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegates

`func (o *UpdateUser) SetDelegates(v []UserDelegate)`

SetDelegates sets Delegates field to given value.

### HasDelegates

`func (o *UpdateUser) HasDelegates() bool`

HasDelegates returns a boolean if a field has been set.

### GetIsBlocked

`func (o *UpdateUser) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *UpdateUser) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *UpdateUser) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *UpdateUser) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetSuspension

`func (o *UpdateUser) GetSuspension() UserSuspend`

GetSuspension returns the Suspension field if non-nil, zero value otherwise.

### GetSuspensionOk

`func (o *UpdateUser) GetSuspensionOk() (*UserSuspend, bool)`

GetSuspensionOk returns a tuple with the Suspension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspension

`func (o *UpdateUser) SetSuspension(v UserSuspend)`

SetSuspension sets Suspension field to given value.

### HasSuspension

`func (o *UpdateUser) HasSuspension() bool`

HasSuspension returns a boolean if a field has been set.

### GetAccessLevel

`func (o *UpdateUser) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UpdateUser) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UpdateUser) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *UpdateUser) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetPromotions

`func (o *UpdateUser) GetPromotions() []UserPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *UpdateUser) GetPromotionsOk() (*[]UserPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *UpdateUser) SetPromotions(v []UserPromotion)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *UpdateUser) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.

### GetExperimentalTags

`func (o *UpdateUser) GetExperimentalTags() []UserExperimentalTag`

GetExperimentalTags returns the ExperimentalTags field if non-nil, zero value otherwise.

### GetExperimentalTagsOk

`func (o *UpdateUser) GetExperimentalTagsOk() (*[]UserExperimentalTag, bool)`

GetExperimentalTagsOk returns a tuple with the ExperimentalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalTags

`func (o *UpdateUser) SetExperimentalTags(v []UserExperimentalTag)`

SetExperimentalTags sets ExperimentalTags field to given value.

### HasExperimentalTags

`func (o *UpdateUser) HasExperimentalTags() bool`

HasExperimentalTags returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateUser) GetPermissions() UserResourcePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateUser) GetPermissionsOk() (*UserResourcePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateUser) SetPermissions(v UserResourcePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateUser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetIsTestAccount

`func (o *UpdateUser) GetIsTestAccount() bool`

GetIsTestAccount returns the IsTestAccount field if non-nil, zero value otherwise.

### GetIsTestAccountOk

`func (o *UpdateUser) GetIsTestAccountOk() (*bool, bool)`

GetIsTestAccountOk returns a tuple with the IsTestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestAccount

`func (o *UpdateUser) SetIsTestAccount(v bool)`

SetIsTestAccount sets IsTestAccount field to given value.

### HasIsTestAccount

`func (o *UpdateUser) HasIsTestAccount() bool`

HasIsTestAccount returns a boolean if a field has been set.

### GetIsBillable

`func (o *UpdateUser) GetIsBillable() bool`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *UpdateUser) GetIsBillableOk() (*bool, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *UpdateUser) SetIsBillable(v bool)`

SetIsBillable sets IsBillable field to given value.

### HasIsBillable

`func (o *UpdateUser) HasIsBillable() bool`

HasIsBillable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the user | [optional] 
**EmailStatus** | Pointer to **string** | The email status of the user | [optional] 
**Language** | Pointer to **string** | The language of the user | [optional] 
**Brand** | Pointer to **string** | The brand of the user | [optional] [default to "default"]
**IsBrandManager** | Pointer to **bool** | Whether the user is a brand manager | [optional] 
**LastLoginTimestamp** | Pointer to **string** | The timestamp when the user logged in last | [optional] 
**LastLoginType** | Pointer to **string** | The last login type of the user | [optional] [default to "md5"]
**IsBlocked** | Pointer to **bool** | Whether the user is blocked | [optional] 
**PasswordChangeRequired** | Pointer to **bool** | Whether the user must change password | [optional] 
**AccessLevel** | Pointer to **string** | The access level of the user | [optional] 
**IsBillable** | Pointer to **bool** | Whether the user is billable | [optional] 
**IsTestingMode** | Pointer to **bool** | Whether the user is in testing mode | [optional] 
**InfrastructureIdDefault** | Pointer to **float32** | The default infrastructure ID of the user | [optional] 
**AuthenticatorMustChange** | Pointer to **bool** | Whether the user must change authenticator | [optional] 
**AuthenticatorCreatedTimestamp** | Pointer to **string** | The timestamp when the authenticator was created | [optional] 
**AuthenticatorEnabled** | Pointer to **bool** | Whether the user has an authenticator | [optional] 
**PromotionTags** | Pointer to **[]string** | The promotion tags of the user | [optional] 
**ExperimentalTags** | Pointer to **[]string** | The experimental tags of the user | [optional] 
**ExternalIds** | Pointer to **map[string]interface{}** | The external IDs of the user | [optional] 
**ExcludeFromReports** | Pointer to **bool** | Whether the user is excluded from reports | [optional] 
**IsTestAccount** | Pointer to **bool** | Whether the user is a test account | [optional] 
**IsArchived** | Pointer to **bool** | Whether the user is a archived | [optional] 
**IsDatastorePublisher** | Pointer to **bool** | Whether the user is a datastore publisher | [optional] 
**AccountId** | Pointer to **float32** | The account ID of the user | [optional] 

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

### GetLastLoginTimestamp

`func (o *UpdateUser) GetLastLoginTimestamp() string`

GetLastLoginTimestamp returns the LastLoginTimestamp field if non-nil, zero value otherwise.

### GetLastLoginTimestampOk

`func (o *UpdateUser) GetLastLoginTimestampOk() (*string, bool)`

GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimestamp

`func (o *UpdateUser) SetLastLoginTimestamp(v string)`

SetLastLoginTimestamp sets LastLoginTimestamp field to given value.

### HasLastLoginTimestamp

`func (o *UpdateUser) HasLastLoginTimestamp() bool`

HasLastLoginTimestamp returns a boolean if a field has been set.

### GetLastLoginType

`func (o *UpdateUser) GetLastLoginType() string`

GetLastLoginType returns the LastLoginType field if non-nil, zero value otherwise.

### GetLastLoginTypeOk

`func (o *UpdateUser) GetLastLoginTypeOk() (*string, bool)`

GetLastLoginTypeOk returns a tuple with the LastLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginType

`func (o *UpdateUser) SetLastLoginType(v string)`

SetLastLoginType sets LastLoginType field to given value.

### HasLastLoginType

`func (o *UpdateUser) HasLastLoginType() bool`

HasLastLoginType returns a boolean if a field has been set.

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

### GetPasswordChangeRequired

`func (o *UpdateUser) GetPasswordChangeRequired() bool`

GetPasswordChangeRequired returns the PasswordChangeRequired field if non-nil, zero value otherwise.

### GetPasswordChangeRequiredOk

`func (o *UpdateUser) GetPasswordChangeRequiredOk() (*bool, bool)`

GetPasswordChangeRequiredOk returns a tuple with the PasswordChangeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeRequired

`func (o *UpdateUser) SetPasswordChangeRequired(v bool)`

SetPasswordChangeRequired sets PasswordChangeRequired field to given value.

### HasPasswordChangeRequired

`func (o *UpdateUser) HasPasswordChangeRequired() bool`

HasPasswordChangeRequired returns a boolean if a field has been set.

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

### GetIsTestingMode

`func (o *UpdateUser) GetIsTestingMode() bool`

GetIsTestingMode returns the IsTestingMode field if non-nil, zero value otherwise.

### GetIsTestingModeOk

`func (o *UpdateUser) GetIsTestingModeOk() (*bool, bool)`

GetIsTestingModeOk returns a tuple with the IsTestingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestingMode

`func (o *UpdateUser) SetIsTestingMode(v bool)`

SetIsTestingMode sets IsTestingMode field to given value.

### HasIsTestingMode

`func (o *UpdateUser) HasIsTestingMode() bool`

HasIsTestingMode returns a boolean if a field has been set.

### GetInfrastructureIdDefault

`func (o *UpdateUser) GetInfrastructureIdDefault() float32`

GetInfrastructureIdDefault returns the InfrastructureIdDefault field if non-nil, zero value otherwise.

### GetInfrastructureIdDefaultOk

`func (o *UpdateUser) GetInfrastructureIdDefaultOk() (*float32, bool)`

GetInfrastructureIdDefaultOk returns a tuple with the InfrastructureIdDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureIdDefault

`func (o *UpdateUser) SetInfrastructureIdDefault(v float32)`

SetInfrastructureIdDefault sets InfrastructureIdDefault field to given value.

### HasInfrastructureIdDefault

`func (o *UpdateUser) HasInfrastructureIdDefault() bool`

HasInfrastructureIdDefault returns a boolean if a field has been set.

### GetAuthenticatorMustChange

`func (o *UpdateUser) GetAuthenticatorMustChange() bool`

GetAuthenticatorMustChange returns the AuthenticatorMustChange field if non-nil, zero value otherwise.

### GetAuthenticatorMustChangeOk

`func (o *UpdateUser) GetAuthenticatorMustChangeOk() (*bool, bool)`

GetAuthenticatorMustChangeOk returns a tuple with the AuthenticatorMustChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorMustChange

`func (o *UpdateUser) SetAuthenticatorMustChange(v bool)`

SetAuthenticatorMustChange sets AuthenticatorMustChange field to given value.

### HasAuthenticatorMustChange

`func (o *UpdateUser) HasAuthenticatorMustChange() bool`

HasAuthenticatorMustChange returns a boolean if a field has been set.

### GetAuthenticatorCreatedTimestamp

`func (o *UpdateUser) GetAuthenticatorCreatedTimestamp() string`

GetAuthenticatorCreatedTimestamp returns the AuthenticatorCreatedTimestamp field if non-nil, zero value otherwise.

### GetAuthenticatorCreatedTimestampOk

`func (o *UpdateUser) GetAuthenticatorCreatedTimestampOk() (*string, bool)`

GetAuthenticatorCreatedTimestampOk returns a tuple with the AuthenticatorCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCreatedTimestamp

`func (o *UpdateUser) SetAuthenticatorCreatedTimestamp(v string)`

SetAuthenticatorCreatedTimestamp sets AuthenticatorCreatedTimestamp field to given value.

### HasAuthenticatorCreatedTimestamp

`func (o *UpdateUser) HasAuthenticatorCreatedTimestamp() bool`

HasAuthenticatorCreatedTimestamp returns a boolean if a field has been set.

### GetAuthenticatorEnabled

`func (o *UpdateUser) GetAuthenticatorEnabled() bool`

GetAuthenticatorEnabled returns the AuthenticatorEnabled field if non-nil, zero value otherwise.

### GetAuthenticatorEnabledOk

`func (o *UpdateUser) GetAuthenticatorEnabledOk() (*bool, bool)`

GetAuthenticatorEnabledOk returns a tuple with the AuthenticatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnabled

`func (o *UpdateUser) SetAuthenticatorEnabled(v bool)`

SetAuthenticatorEnabled sets AuthenticatorEnabled field to given value.

### HasAuthenticatorEnabled

`func (o *UpdateUser) HasAuthenticatorEnabled() bool`

HasAuthenticatorEnabled returns a boolean if a field has been set.

### GetPromotionTags

`func (o *UpdateUser) GetPromotionTags() []string`

GetPromotionTags returns the PromotionTags field if non-nil, zero value otherwise.

### GetPromotionTagsOk

`func (o *UpdateUser) GetPromotionTagsOk() (*[]string, bool)`

GetPromotionTagsOk returns a tuple with the PromotionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionTags

`func (o *UpdateUser) SetPromotionTags(v []string)`

SetPromotionTags sets PromotionTags field to given value.

### HasPromotionTags

`func (o *UpdateUser) HasPromotionTags() bool`

HasPromotionTags returns a boolean if a field has been set.

### GetExperimentalTags

`func (o *UpdateUser) GetExperimentalTags() []string`

GetExperimentalTags returns the ExperimentalTags field if non-nil, zero value otherwise.

### GetExperimentalTagsOk

`func (o *UpdateUser) GetExperimentalTagsOk() (*[]string, bool)`

GetExperimentalTagsOk returns a tuple with the ExperimentalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalTags

`func (o *UpdateUser) SetExperimentalTags(v []string)`

SetExperimentalTags sets ExperimentalTags field to given value.

### HasExperimentalTags

`func (o *UpdateUser) HasExperimentalTags() bool`

HasExperimentalTags returns a boolean if a field has been set.

### GetExternalIds

`func (o *UpdateUser) GetExternalIds() map[string]interface{}`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *UpdateUser) GetExternalIdsOk() (*map[string]interface{}, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *UpdateUser) SetExternalIds(v map[string]interface{})`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *UpdateUser) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExcludeFromReports

`func (o *UpdateUser) GetExcludeFromReports() bool`

GetExcludeFromReports returns the ExcludeFromReports field if non-nil, zero value otherwise.

### GetExcludeFromReportsOk

`func (o *UpdateUser) GetExcludeFromReportsOk() (*bool, bool)`

GetExcludeFromReportsOk returns a tuple with the ExcludeFromReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromReports

`func (o *UpdateUser) SetExcludeFromReports(v bool)`

SetExcludeFromReports sets ExcludeFromReports field to given value.

### HasExcludeFromReports

`func (o *UpdateUser) HasExcludeFromReports() bool`

HasExcludeFromReports returns a boolean if a field has been set.

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

### GetIsArchived

`func (o *UpdateUser) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UpdateUser) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UpdateUser) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *UpdateUser) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsDatastorePublisher

`func (o *UpdateUser) GetIsDatastorePublisher() bool`

GetIsDatastorePublisher returns the IsDatastorePublisher field if non-nil, zero value otherwise.

### GetIsDatastorePublisherOk

`func (o *UpdateUser) GetIsDatastorePublisherOk() (*bool, bool)`

GetIsDatastorePublisherOk returns a tuple with the IsDatastorePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDatastorePublisher

`func (o *UpdateUser) SetIsDatastorePublisher(v bool)`

SetIsDatastorePublisher sets IsDatastorePublisher field to given value.

### HasIsDatastorePublisher

`func (o *UpdateUser) HasIsDatastorePublisher() bool`

HasIsDatastorePublisher returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateUser) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateUser) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateUser) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



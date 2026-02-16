# UserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision of the user configuration | 
**DisplayName** | **string** | The display name of the user | 
**EmailStatus** | **string** | The email status of the user | 
**Language** | **string** | The language of the user | 
**Brand** | **string** | The brand of the user | [default to "default"]
**IsBrandManager** | **bool** | Whether the user is a brand manager | 
**LastLoginTimestamp** | **string** | The timestamp when the user logged in last | 
**LastLoginType** | **string** | The last login type of the user | [default to "md5"]
**IsBlocked** | **bool** | Whether the user is blocked | 
**PasswordChangeRequired** | **bool** | Whether the user must change password | 
**AccessLevel** | **string** | The access level of the user | 
**IsBillable** | **bool** | Whether the user is billable | 
**IsTestingMode** | **bool** | Whether the user is in testing mode | 
**InfrastructureIdDefault** | Pointer to **float32** | The default infrastructure ID of the user | [optional] 
**AuthenticatorMustChange** | **bool** | Whether the user must change authenticator | 
**AuthenticatorCreatedTimestamp** | **string** | The timestamp when the authenticator was created | 
**PromotionTags** | Pointer to **[]string** | The promotion tags of the user | [optional] 
**ExperimentalTags** | Pointer to **[]string** | The experimental tags of the user | [optional] 
**ExternalIds** | Pointer to **map[string]interface{}** | The external IDs of the user | [optional] 
**ExcludeFromReports** | **bool** | Whether the user is excluded from reports | 
**IsTestAccount** | **bool** | Whether the user is a test account | 
**IsArchived** | **bool** | Whether the user is a archived | 
**IsDatastorePublisher** | **bool** | Whether the user is a datastore publisher | 
**AccountId** | Pointer to **float32** | The account ID of the user | [optional] 
**Provider** | **string** | The provider of the user | [default to "mysql"]
**PasswordLastChangedTimestamp** | **string** | The timestamp when the user last changed their password | 

## Methods

### NewUserConfiguration

`func NewUserConfiguration(revision float32, displayName string, emailStatus string, language string, brand string, isBrandManager bool, lastLoginTimestamp string, lastLoginType string, isBlocked bool, passwordChangeRequired bool, accessLevel string, isBillable bool, isTestingMode bool, authenticatorMustChange bool, authenticatorCreatedTimestamp string, excludeFromReports bool, isTestAccount bool, isArchived bool, isDatastorePublisher bool, provider string, passwordLastChangedTimestamp string, ) *UserConfiguration`

NewUserConfiguration instantiates a new UserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigurationWithDefaults

`func NewUserConfigurationWithDefaults() *UserConfiguration`

NewUserConfigurationWithDefaults instantiates a new UserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *UserConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UserConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UserConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetDisplayName

`func (o *UserConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmailStatus

`func (o *UserConfiguration) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *UserConfiguration) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *UserConfiguration) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.


### GetLanguage

`func (o *UserConfiguration) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserConfiguration) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserConfiguration) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetBrand

`func (o *UserConfiguration) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UserConfiguration) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UserConfiguration) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsBrandManager

`func (o *UserConfiguration) GetIsBrandManager() bool`

GetIsBrandManager returns the IsBrandManager field if non-nil, zero value otherwise.

### GetIsBrandManagerOk

`func (o *UserConfiguration) GetIsBrandManagerOk() (*bool, bool)`

GetIsBrandManagerOk returns a tuple with the IsBrandManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandManager

`func (o *UserConfiguration) SetIsBrandManager(v bool)`

SetIsBrandManager sets IsBrandManager field to given value.


### GetLastLoginTimestamp

`func (o *UserConfiguration) GetLastLoginTimestamp() string`

GetLastLoginTimestamp returns the LastLoginTimestamp field if non-nil, zero value otherwise.

### GetLastLoginTimestampOk

`func (o *UserConfiguration) GetLastLoginTimestampOk() (*string, bool)`

GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimestamp

`func (o *UserConfiguration) SetLastLoginTimestamp(v string)`

SetLastLoginTimestamp sets LastLoginTimestamp field to given value.


### GetLastLoginType

`func (o *UserConfiguration) GetLastLoginType() string`

GetLastLoginType returns the LastLoginType field if non-nil, zero value otherwise.

### GetLastLoginTypeOk

`func (o *UserConfiguration) GetLastLoginTypeOk() (*string, bool)`

GetLastLoginTypeOk returns a tuple with the LastLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginType

`func (o *UserConfiguration) SetLastLoginType(v string)`

SetLastLoginType sets LastLoginType field to given value.


### GetIsBlocked

`func (o *UserConfiguration) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *UserConfiguration) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *UserConfiguration) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.


### GetPasswordChangeRequired

`func (o *UserConfiguration) GetPasswordChangeRequired() bool`

GetPasswordChangeRequired returns the PasswordChangeRequired field if non-nil, zero value otherwise.

### GetPasswordChangeRequiredOk

`func (o *UserConfiguration) GetPasswordChangeRequiredOk() (*bool, bool)`

GetPasswordChangeRequiredOk returns a tuple with the PasswordChangeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeRequired

`func (o *UserConfiguration) SetPasswordChangeRequired(v bool)`

SetPasswordChangeRequired sets PasswordChangeRequired field to given value.


### GetAccessLevel

`func (o *UserConfiguration) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UserConfiguration) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UserConfiguration) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetIsBillable

`func (o *UserConfiguration) GetIsBillable() bool`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *UserConfiguration) GetIsBillableOk() (*bool, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *UserConfiguration) SetIsBillable(v bool)`

SetIsBillable sets IsBillable field to given value.


### GetIsTestingMode

`func (o *UserConfiguration) GetIsTestingMode() bool`

GetIsTestingMode returns the IsTestingMode field if non-nil, zero value otherwise.

### GetIsTestingModeOk

`func (o *UserConfiguration) GetIsTestingModeOk() (*bool, bool)`

GetIsTestingModeOk returns a tuple with the IsTestingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestingMode

`func (o *UserConfiguration) SetIsTestingMode(v bool)`

SetIsTestingMode sets IsTestingMode field to given value.


### GetInfrastructureIdDefault

`func (o *UserConfiguration) GetInfrastructureIdDefault() float32`

GetInfrastructureIdDefault returns the InfrastructureIdDefault field if non-nil, zero value otherwise.

### GetInfrastructureIdDefaultOk

`func (o *UserConfiguration) GetInfrastructureIdDefaultOk() (*float32, bool)`

GetInfrastructureIdDefaultOk returns a tuple with the InfrastructureIdDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureIdDefault

`func (o *UserConfiguration) SetInfrastructureIdDefault(v float32)`

SetInfrastructureIdDefault sets InfrastructureIdDefault field to given value.

### HasInfrastructureIdDefault

`func (o *UserConfiguration) HasInfrastructureIdDefault() bool`

HasInfrastructureIdDefault returns a boolean if a field has been set.

### GetAuthenticatorMustChange

`func (o *UserConfiguration) GetAuthenticatorMustChange() bool`

GetAuthenticatorMustChange returns the AuthenticatorMustChange field if non-nil, zero value otherwise.

### GetAuthenticatorMustChangeOk

`func (o *UserConfiguration) GetAuthenticatorMustChangeOk() (*bool, bool)`

GetAuthenticatorMustChangeOk returns a tuple with the AuthenticatorMustChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorMustChange

`func (o *UserConfiguration) SetAuthenticatorMustChange(v bool)`

SetAuthenticatorMustChange sets AuthenticatorMustChange field to given value.


### GetAuthenticatorCreatedTimestamp

`func (o *UserConfiguration) GetAuthenticatorCreatedTimestamp() string`

GetAuthenticatorCreatedTimestamp returns the AuthenticatorCreatedTimestamp field if non-nil, zero value otherwise.

### GetAuthenticatorCreatedTimestampOk

`func (o *UserConfiguration) GetAuthenticatorCreatedTimestampOk() (*string, bool)`

GetAuthenticatorCreatedTimestampOk returns a tuple with the AuthenticatorCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCreatedTimestamp

`func (o *UserConfiguration) SetAuthenticatorCreatedTimestamp(v string)`

SetAuthenticatorCreatedTimestamp sets AuthenticatorCreatedTimestamp field to given value.


### GetPromotionTags

`func (o *UserConfiguration) GetPromotionTags() []string`

GetPromotionTags returns the PromotionTags field if non-nil, zero value otherwise.

### GetPromotionTagsOk

`func (o *UserConfiguration) GetPromotionTagsOk() (*[]string, bool)`

GetPromotionTagsOk returns a tuple with the PromotionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionTags

`func (o *UserConfiguration) SetPromotionTags(v []string)`

SetPromotionTags sets PromotionTags field to given value.

### HasPromotionTags

`func (o *UserConfiguration) HasPromotionTags() bool`

HasPromotionTags returns a boolean if a field has been set.

### GetExperimentalTags

`func (o *UserConfiguration) GetExperimentalTags() []string`

GetExperimentalTags returns the ExperimentalTags field if non-nil, zero value otherwise.

### GetExperimentalTagsOk

`func (o *UserConfiguration) GetExperimentalTagsOk() (*[]string, bool)`

GetExperimentalTagsOk returns a tuple with the ExperimentalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalTags

`func (o *UserConfiguration) SetExperimentalTags(v []string)`

SetExperimentalTags sets ExperimentalTags field to given value.

### HasExperimentalTags

`func (o *UserConfiguration) HasExperimentalTags() bool`

HasExperimentalTags returns a boolean if a field has been set.

### GetExternalIds

`func (o *UserConfiguration) GetExternalIds() map[string]interface{}`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *UserConfiguration) GetExternalIdsOk() (*map[string]interface{}, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *UserConfiguration) SetExternalIds(v map[string]interface{})`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *UserConfiguration) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExcludeFromReports

`func (o *UserConfiguration) GetExcludeFromReports() bool`

GetExcludeFromReports returns the ExcludeFromReports field if non-nil, zero value otherwise.

### GetExcludeFromReportsOk

`func (o *UserConfiguration) GetExcludeFromReportsOk() (*bool, bool)`

GetExcludeFromReportsOk returns a tuple with the ExcludeFromReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromReports

`func (o *UserConfiguration) SetExcludeFromReports(v bool)`

SetExcludeFromReports sets ExcludeFromReports field to given value.


### GetIsTestAccount

`func (o *UserConfiguration) GetIsTestAccount() bool`

GetIsTestAccount returns the IsTestAccount field if non-nil, zero value otherwise.

### GetIsTestAccountOk

`func (o *UserConfiguration) GetIsTestAccountOk() (*bool, bool)`

GetIsTestAccountOk returns a tuple with the IsTestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestAccount

`func (o *UserConfiguration) SetIsTestAccount(v bool)`

SetIsTestAccount sets IsTestAccount field to given value.


### GetIsArchived

`func (o *UserConfiguration) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UserConfiguration) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UserConfiguration) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetIsDatastorePublisher

`func (o *UserConfiguration) GetIsDatastorePublisher() bool`

GetIsDatastorePublisher returns the IsDatastorePublisher field if non-nil, zero value otherwise.

### GetIsDatastorePublisherOk

`func (o *UserConfiguration) GetIsDatastorePublisherOk() (*bool, bool)`

GetIsDatastorePublisherOk returns a tuple with the IsDatastorePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDatastorePublisher

`func (o *UserConfiguration) SetIsDatastorePublisher(v bool)`

SetIsDatastorePublisher sets IsDatastorePublisher field to given value.


### GetAccountId

`func (o *UserConfiguration) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserConfiguration) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserConfiguration) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserConfiguration) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProvider

`func (o *UserConfiguration) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserConfiguration) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserConfiguration) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPasswordLastChangedTimestamp

`func (o *UserConfiguration) GetPasswordLastChangedTimestamp() string`

GetPasswordLastChangedTimestamp returns the PasswordLastChangedTimestamp field if non-nil, zero value otherwise.

### GetPasswordLastChangedTimestampOk

`func (o *UserConfiguration) GetPasswordLastChangedTimestampOk() (*string, bool)`

GetPasswordLastChangedTimestampOk returns a tuple with the PasswordLastChangedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChangedTimestamp

`func (o *UserConfiguration) SetPasswordLastChangedTimestamp(v string)`

SetPasswordLastChangedTimestamp sets PasswordLastChangedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



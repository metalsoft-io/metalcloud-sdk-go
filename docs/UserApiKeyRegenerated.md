# UserApiKeyRegenerated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**AccountId** | Pointer to **int64** | The account ID of the user | [optional] 
**Provider** | **string** | The provider of the user | [default to "built-in"]
**PasswordLastChangedTimestamp** | **string** | The timestamp when the user last changed their password | 
**ApiKeyLastChangedTimestamp** | Pointer to **string** | The timestamp when the user API key was last regenerated | [optional] 
**ApiKeyValidityDuration** | Pointer to **int32** | Per-user override for the API key validity duration in seconds; 0 means the API key never expires; unset means the account or global setting applies (MS-10501). Accepted range: 0 - 4294967295. | [optional] 
**Id** | **int64** | User ID | 
**Revision** | **int64** | Revision of the user | 
**Email** | **string** | The email address of the user | 
**Franchise** | **string** | The franchise of the user | 
**CreatedTimestamp** | **string** | The timestamp when the user was created | 
**PlanType** | **string** | The plan type of the user | [default to "vanilla"]
**ParentDelegates** | Pointer to **[]string** | The email addresses of the parent delegate users | [optional] 
**ChildDelegates** | Pointer to **[]string** | The email addresses of the child delegate users | [optional] 
**IsSuspended** | **bool** | Whether the user is suspended | 
**AuthenticatorEnabled** | **bool** | Whether the user has an authenticator | 
**AccessibleResources** | Pointer to **map[string][]string** | The resources accessible by the user | [optional] 
**RoleAccessibleResources** | Pointer to **map[string][]string** | The resources granted by the user&#39;s role/permissions, excluding the baseline policies granted to every user | [optional] 
**Config** | [**UserConfiguration**](UserConfiguration.md) | The new configuration of the user. | 
**Meta** | [**UserMeta**](UserMeta.md) | Meta information of the user. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**ApiKey** | Pointer to **string** | The newly generated API key. Under the &#x60;once&#x60; visibility this response is the only place it is ever served. | [optional] 

## Methods

### NewUserApiKeyRegenerated

`func NewUserApiKeyRegenerated(displayName string, emailStatus string, language string, brand string, isBrandManager bool, lastLoginTimestamp string, lastLoginType string, isBlocked bool, passwordChangeRequired bool, accessLevel string, isBillable bool, isTestingMode bool, authenticatorMustChange bool, authenticatorCreatedTimestamp string, excludeFromReports bool, isTestAccount bool, isArchived bool, isDatastorePublisher bool, provider string, passwordLastChangedTimestamp string, id int64, revision int64, email string, franchise string, createdTimestamp string, planType string, isSuspended bool, authenticatorEnabled bool, config UserConfiguration, meta UserMeta, ) *UserApiKeyRegenerated`

NewUserApiKeyRegenerated instantiates a new UserApiKeyRegenerated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApiKeyRegeneratedWithDefaults

`func NewUserApiKeyRegeneratedWithDefaults() *UserApiKeyRegenerated`

NewUserApiKeyRegeneratedWithDefaults instantiates a new UserApiKeyRegenerated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UserApiKeyRegenerated) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserApiKeyRegenerated) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserApiKeyRegenerated) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmailStatus

`func (o *UserApiKeyRegenerated) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *UserApiKeyRegenerated) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *UserApiKeyRegenerated) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.


### GetLanguage

`func (o *UserApiKeyRegenerated) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserApiKeyRegenerated) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserApiKeyRegenerated) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetBrand

`func (o *UserApiKeyRegenerated) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UserApiKeyRegenerated) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UserApiKeyRegenerated) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsBrandManager

`func (o *UserApiKeyRegenerated) GetIsBrandManager() bool`

GetIsBrandManager returns the IsBrandManager field if non-nil, zero value otherwise.

### GetIsBrandManagerOk

`func (o *UserApiKeyRegenerated) GetIsBrandManagerOk() (*bool, bool)`

GetIsBrandManagerOk returns a tuple with the IsBrandManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandManager

`func (o *UserApiKeyRegenerated) SetIsBrandManager(v bool)`

SetIsBrandManager sets IsBrandManager field to given value.


### GetLastLoginTimestamp

`func (o *UserApiKeyRegenerated) GetLastLoginTimestamp() string`

GetLastLoginTimestamp returns the LastLoginTimestamp field if non-nil, zero value otherwise.

### GetLastLoginTimestampOk

`func (o *UserApiKeyRegenerated) GetLastLoginTimestampOk() (*string, bool)`

GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimestamp

`func (o *UserApiKeyRegenerated) SetLastLoginTimestamp(v string)`

SetLastLoginTimestamp sets LastLoginTimestamp field to given value.


### GetLastLoginType

`func (o *UserApiKeyRegenerated) GetLastLoginType() string`

GetLastLoginType returns the LastLoginType field if non-nil, zero value otherwise.

### GetLastLoginTypeOk

`func (o *UserApiKeyRegenerated) GetLastLoginTypeOk() (*string, bool)`

GetLastLoginTypeOk returns a tuple with the LastLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginType

`func (o *UserApiKeyRegenerated) SetLastLoginType(v string)`

SetLastLoginType sets LastLoginType field to given value.


### GetIsBlocked

`func (o *UserApiKeyRegenerated) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *UserApiKeyRegenerated) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *UserApiKeyRegenerated) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.


### GetPasswordChangeRequired

`func (o *UserApiKeyRegenerated) GetPasswordChangeRequired() bool`

GetPasswordChangeRequired returns the PasswordChangeRequired field if non-nil, zero value otherwise.

### GetPasswordChangeRequiredOk

`func (o *UserApiKeyRegenerated) GetPasswordChangeRequiredOk() (*bool, bool)`

GetPasswordChangeRequiredOk returns a tuple with the PasswordChangeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeRequired

`func (o *UserApiKeyRegenerated) SetPasswordChangeRequired(v bool)`

SetPasswordChangeRequired sets PasswordChangeRequired field to given value.


### GetAccessLevel

`func (o *UserApiKeyRegenerated) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UserApiKeyRegenerated) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UserApiKeyRegenerated) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetIsBillable

`func (o *UserApiKeyRegenerated) GetIsBillable() bool`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *UserApiKeyRegenerated) GetIsBillableOk() (*bool, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *UserApiKeyRegenerated) SetIsBillable(v bool)`

SetIsBillable sets IsBillable field to given value.


### GetIsTestingMode

`func (o *UserApiKeyRegenerated) GetIsTestingMode() bool`

GetIsTestingMode returns the IsTestingMode field if non-nil, zero value otherwise.

### GetIsTestingModeOk

`func (o *UserApiKeyRegenerated) GetIsTestingModeOk() (*bool, bool)`

GetIsTestingModeOk returns a tuple with the IsTestingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestingMode

`func (o *UserApiKeyRegenerated) SetIsTestingMode(v bool)`

SetIsTestingMode sets IsTestingMode field to given value.


### GetInfrastructureIdDefault

`func (o *UserApiKeyRegenerated) GetInfrastructureIdDefault() float32`

GetInfrastructureIdDefault returns the InfrastructureIdDefault field if non-nil, zero value otherwise.

### GetInfrastructureIdDefaultOk

`func (o *UserApiKeyRegenerated) GetInfrastructureIdDefaultOk() (*float32, bool)`

GetInfrastructureIdDefaultOk returns a tuple with the InfrastructureIdDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureIdDefault

`func (o *UserApiKeyRegenerated) SetInfrastructureIdDefault(v float32)`

SetInfrastructureIdDefault sets InfrastructureIdDefault field to given value.

### HasInfrastructureIdDefault

`func (o *UserApiKeyRegenerated) HasInfrastructureIdDefault() bool`

HasInfrastructureIdDefault returns a boolean if a field has been set.

### GetAuthenticatorMustChange

`func (o *UserApiKeyRegenerated) GetAuthenticatorMustChange() bool`

GetAuthenticatorMustChange returns the AuthenticatorMustChange field if non-nil, zero value otherwise.

### GetAuthenticatorMustChangeOk

`func (o *UserApiKeyRegenerated) GetAuthenticatorMustChangeOk() (*bool, bool)`

GetAuthenticatorMustChangeOk returns a tuple with the AuthenticatorMustChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorMustChange

`func (o *UserApiKeyRegenerated) SetAuthenticatorMustChange(v bool)`

SetAuthenticatorMustChange sets AuthenticatorMustChange field to given value.


### GetAuthenticatorCreatedTimestamp

`func (o *UserApiKeyRegenerated) GetAuthenticatorCreatedTimestamp() string`

GetAuthenticatorCreatedTimestamp returns the AuthenticatorCreatedTimestamp field if non-nil, zero value otherwise.

### GetAuthenticatorCreatedTimestampOk

`func (o *UserApiKeyRegenerated) GetAuthenticatorCreatedTimestampOk() (*string, bool)`

GetAuthenticatorCreatedTimestampOk returns a tuple with the AuthenticatorCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCreatedTimestamp

`func (o *UserApiKeyRegenerated) SetAuthenticatorCreatedTimestamp(v string)`

SetAuthenticatorCreatedTimestamp sets AuthenticatorCreatedTimestamp field to given value.


### GetPromotionTags

`func (o *UserApiKeyRegenerated) GetPromotionTags() []string`

GetPromotionTags returns the PromotionTags field if non-nil, zero value otherwise.

### GetPromotionTagsOk

`func (o *UserApiKeyRegenerated) GetPromotionTagsOk() (*[]string, bool)`

GetPromotionTagsOk returns a tuple with the PromotionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionTags

`func (o *UserApiKeyRegenerated) SetPromotionTags(v []string)`

SetPromotionTags sets PromotionTags field to given value.

### HasPromotionTags

`func (o *UserApiKeyRegenerated) HasPromotionTags() bool`

HasPromotionTags returns a boolean if a field has been set.

### GetExperimentalTags

`func (o *UserApiKeyRegenerated) GetExperimentalTags() []string`

GetExperimentalTags returns the ExperimentalTags field if non-nil, zero value otherwise.

### GetExperimentalTagsOk

`func (o *UserApiKeyRegenerated) GetExperimentalTagsOk() (*[]string, bool)`

GetExperimentalTagsOk returns a tuple with the ExperimentalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalTags

`func (o *UserApiKeyRegenerated) SetExperimentalTags(v []string)`

SetExperimentalTags sets ExperimentalTags field to given value.

### HasExperimentalTags

`func (o *UserApiKeyRegenerated) HasExperimentalTags() bool`

HasExperimentalTags returns a boolean if a field has been set.

### GetExternalIds

`func (o *UserApiKeyRegenerated) GetExternalIds() map[string]interface{}`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *UserApiKeyRegenerated) GetExternalIdsOk() (*map[string]interface{}, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *UserApiKeyRegenerated) SetExternalIds(v map[string]interface{})`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *UserApiKeyRegenerated) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExcludeFromReports

`func (o *UserApiKeyRegenerated) GetExcludeFromReports() bool`

GetExcludeFromReports returns the ExcludeFromReports field if non-nil, zero value otherwise.

### GetExcludeFromReportsOk

`func (o *UserApiKeyRegenerated) GetExcludeFromReportsOk() (*bool, bool)`

GetExcludeFromReportsOk returns a tuple with the ExcludeFromReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromReports

`func (o *UserApiKeyRegenerated) SetExcludeFromReports(v bool)`

SetExcludeFromReports sets ExcludeFromReports field to given value.


### GetIsTestAccount

`func (o *UserApiKeyRegenerated) GetIsTestAccount() bool`

GetIsTestAccount returns the IsTestAccount field if non-nil, zero value otherwise.

### GetIsTestAccountOk

`func (o *UserApiKeyRegenerated) GetIsTestAccountOk() (*bool, bool)`

GetIsTestAccountOk returns a tuple with the IsTestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestAccount

`func (o *UserApiKeyRegenerated) SetIsTestAccount(v bool)`

SetIsTestAccount sets IsTestAccount field to given value.


### GetIsArchived

`func (o *UserApiKeyRegenerated) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UserApiKeyRegenerated) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UserApiKeyRegenerated) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetIsDatastorePublisher

`func (o *UserApiKeyRegenerated) GetIsDatastorePublisher() bool`

GetIsDatastorePublisher returns the IsDatastorePublisher field if non-nil, zero value otherwise.

### GetIsDatastorePublisherOk

`func (o *UserApiKeyRegenerated) GetIsDatastorePublisherOk() (*bool, bool)`

GetIsDatastorePublisherOk returns a tuple with the IsDatastorePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDatastorePublisher

`func (o *UserApiKeyRegenerated) SetIsDatastorePublisher(v bool)`

SetIsDatastorePublisher sets IsDatastorePublisher field to given value.


### GetAccountId

`func (o *UserApiKeyRegenerated) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserApiKeyRegenerated) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserApiKeyRegenerated) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserApiKeyRegenerated) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProvider

`func (o *UserApiKeyRegenerated) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserApiKeyRegenerated) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserApiKeyRegenerated) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPasswordLastChangedTimestamp

`func (o *UserApiKeyRegenerated) GetPasswordLastChangedTimestamp() string`

GetPasswordLastChangedTimestamp returns the PasswordLastChangedTimestamp field if non-nil, zero value otherwise.

### GetPasswordLastChangedTimestampOk

`func (o *UserApiKeyRegenerated) GetPasswordLastChangedTimestampOk() (*string, bool)`

GetPasswordLastChangedTimestampOk returns a tuple with the PasswordLastChangedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastChangedTimestamp

`func (o *UserApiKeyRegenerated) SetPasswordLastChangedTimestamp(v string)`

SetPasswordLastChangedTimestamp sets PasswordLastChangedTimestamp field to given value.


### GetApiKeyLastChangedTimestamp

`func (o *UserApiKeyRegenerated) GetApiKeyLastChangedTimestamp() string`

GetApiKeyLastChangedTimestamp returns the ApiKeyLastChangedTimestamp field if non-nil, zero value otherwise.

### GetApiKeyLastChangedTimestampOk

`func (o *UserApiKeyRegenerated) GetApiKeyLastChangedTimestampOk() (*string, bool)`

GetApiKeyLastChangedTimestampOk returns a tuple with the ApiKeyLastChangedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyLastChangedTimestamp

`func (o *UserApiKeyRegenerated) SetApiKeyLastChangedTimestamp(v string)`

SetApiKeyLastChangedTimestamp sets ApiKeyLastChangedTimestamp field to given value.

### HasApiKeyLastChangedTimestamp

`func (o *UserApiKeyRegenerated) HasApiKeyLastChangedTimestamp() bool`

HasApiKeyLastChangedTimestamp returns a boolean if a field has been set.

### GetApiKeyValidityDuration

`func (o *UserApiKeyRegenerated) GetApiKeyValidityDuration() int32`

GetApiKeyValidityDuration returns the ApiKeyValidityDuration field if non-nil, zero value otherwise.

### GetApiKeyValidityDurationOk

`func (o *UserApiKeyRegenerated) GetApiKeyValidityDurationOk() (*int32, bool)`

GetApiKeyValidityDurationOk returns a tuple with the ApiKeyValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyValidityDuration

`func (o *UserApiKeyRegenerated) SetApiKeyValidityDuration(v int32)`

SetApiKeyValidityDuration sets ApiKeyValidityDuration field to given value.

### HasApiKeyValidityDuration

`func (o *UserApiKeyRegenerated) HasApiKeyValidityDuration() bool`

HasApiKeyValidityDuration returns a boolean if a field has been set.

### GetId

`func (o *UserApiKeyRegenerated) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserApiKeyRegenerated) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserApiKeyRegenerated) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *UserApiKeyRegenerated) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UserApiKeyRegenerated) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UserApiKeyRegenerated) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetEmail

`func (o *UserApiKeyRegenerated) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserApiKeyRegenerated) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserApiKeyRegenerated) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFranchise

`func (o *UserApiKeyRegenerated) GetFranchise() string`

GetFranchise returns the Franchise field if non-nil, zero value otherwise.

### GetFranchiseOk

`func (o *UserApiKeyRegenerated) GetFranchiseOk() (*string, bool)`

GetFranchiseOk returns a tuple with the Franchise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFranchise

`func (o *UserApiKeyRegenerated) SetFranchise(v string)`

SetFranchise sets Franchise field to given value.


### GetCreatedTimestamp

`func (o *UserApiKeyRegenerated) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *UserApiKeyRegenerated) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *UserApiKeyRegenerated) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetPlanType

`func (o *UserApiKeyRegenerated) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *UserApiKeyRegenerated) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *UserApiKeyRegenerated) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetParentDelegates

`func (o *UserApiKeyRegenerated) GetParentDelegates() []string`

GetParentDelegates returns the ParentDelegates field if non-nil, zero value otherwise.

### GetParentDelegatesOk

`func (o *UserApiKeyRegenerated) GetParentDelegatesOk() (*[]string, bool)`

GetParentDelegatesOk returns a tuple with the ParentDelegates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDelegates

`func (o *UserApiKeyRegenerated) SetParentDelegates(v []string)`

SetParentDelegates sets ParentDelegates field to given value.

### HasParentDelegates

`func (o *UserApiKeyRegenerated) HasParentDelegates() bool`

HasParentDelegates returns a boolean if a field has been set.

### GetChildDelegates

`func (o *UserApiKeyRegenerated) GetChildDelegates() []string`

GetChildDelegates returns the ChildDelegates field if non-nil, zero value otherwise.

### GetChildDelegatesOk

`func (o *UserApiKeyRegenerated) GetChildDelegatesOk() (*[]string, bool)`

GetChildDelegatesOk returns a tuple with the ChildDelegates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildDelegates

`func (o *UserApiKeyRegenerated) SetChildDelegates(v []string)`

SetChildDelegates sets ChildDelegates field to given value.

### HasChildDelegates

`func (o *UserApiKeyRegenerated) HasChildDelegates() bool`

HasChildDelegates returns a boolean if a field has been set.

### GetIsSuspended

`func (o *UserApiKeyRegenerated) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *UserApiKeyRegenerated) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *UserApiKeyRegenerated) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.


### GetAuthenticatorEnabled

`func (o *UserApiKeyRegenerated) GetAuthenticatorEnabled() bool`

GetAuthenticatorEnabled returns the AuthenticatorEnabled field if non-nil, zero value otherwise.

### GetAuthenticatorEnabledOk

`func (o *UserApiKeyRegenerated) GetAuthenticatorEnabledOk() (*bool, bool)`

GetAuthenticatorEnabledOk returns a tuple with the AuthenticatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnabled

`func (o *UserApiKeyRegenerated) SetAuthenticatorEnabled(v bool)`

SetAuthenticatorEnabled sets AuthenticatorEnabled field to given value.


### GetAccessibleResources

`func (o *UserApiKeyRegenerated) GetAccessibleResources() map[string][]string`

GetAccessibleResources returns the AccessibleResources field if non-nil, zero value otherwise.

### GetAccessibleResourcesOk

`func (o *UserApiKeyRegenerated) GetAccessibleResourcesOk() (*map[string][]string, bool)`

GetAccessibleResourcesOk returns a tuple with the AccessibleResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleResources

`func (o *UserApiKeyRegenerated) SetAccessibleResources(v map[string][]string)`

SetAccessibleResources sets AccessibleResources field to given value.

### HasAccessibleResources

`func (o *UserApiKeyRegenerated) HasAccessibleResources() bool`

HasAccessibleResources returns a boolean if a field has been set.

### GetRoleAccessibleResources

`func (o *UserApiKeyRegenerated) GetRoleAccessibleResources() map[string][]string`

GetRoleAccessibleResources returns the RoleAccessibleResources field if non-nil, zero value otherwise.

### GetRoleAccessibleResourcesOk

`func (o *UserApiKeyRegenerated) GetRoleAccessibleResourcesOk() (*map[string][]string, bool)`

GetRoleAccessibleResourcesOk returns a tuple with the RoleAccessibleResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAccessibleResources

`func (o *UserApiKeyRegenerated) SetRoleAccessibleResources(v map[string][]string)`

SetRoleAccessibleResources sets RoleAccessibleResources field to given value.

### HasRoleAccessibleResources

`func (o *UserApiKeyRegenerated) HasRoleAccessibleResources() bool`

HasRoleAccessibleResources returns a boolean if a field has been set.

### GetConfig

`func (o *UserApiKeyRegenerated) GetConfig() UserConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UserApiKeyRegenerated) GetConfigOk() (*UserConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UserApiKeyRegenerated) SetConfig(v UserConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *UserApiKeyRegenerated) GetMeta() UserMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UserApiKeyRegenerated) GetMetaOk() (*UserMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UserApiKeyRegenerated) SetMeta(v UserMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *UserApiKeyRegenerated) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserApiKeyRegenerated) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserApiKeyRegenerated) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserApiKeyRegenerated) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApiKey

`func (o *UserApiKeyRegenerated) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserApiKeyRegenerated) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserApiKeyRegenerated) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UserApiKeyRegenerated) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



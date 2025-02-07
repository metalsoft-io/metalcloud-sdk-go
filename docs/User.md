# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID | 
**Revision** | **string** | Revision number | 
**Franchise** | **string** | The franchise of the user | 
**DisplayName** | **string** | The display name of the user | 
**Email** | **string** | The email address of the user | 
**AuthenticatorMustChange** | **float32** | Whether the user must change authenticator | 
**AuthenticatorCreatedTimestamp** | **string** | The timestamp when the authenticator was created | 
**AuthenticatorEnabled** | **bool** | Whether the user has an authenticator | 
**CreatedTimestamp** | **string** | The timestamp when the user was created | 
**LastLoginTimestamp** | **string** | The timestamp when the user logged in last | 
**LastLoginType** | **string** | The last login type of the user | 
**Blocked** | **float32** | Whether the user is blocked | 
**EmailStatus** | **string** | The email status of the user | 
**PasswordChangeRequired** | **float32** | Whether the user must change password | 
**AccessLevel** | **string** | The access level of the user | 
**IsBillable** | **float32** | Whether the user is billable | 
**IsTestingMode** | **float32** | Whether the user is in testing mode | 
**IsSuspended** | **float32** | Whether the user is suspended | 
**Language** | **string** | The language of the user | 
**InfrastructureIdDefault** | Pointer to **float32** | The default infrastructure ID of the user | [optional] 
**GuiSettings** | Pointer to **map[string]interface{}** | The GUI settings of the user | [optional] 
**PromotionTags** | Pointer to **map[string]interface{}** | The promotion tags of the user | [optional] 
**ExperimentalTags** | Pointer to **map[string]interface{}** | The experimental tags of the user | [optional] 
**Permissions** | Pointer to [**UserPermissions**](UserPermissions.md) | The permissions of the user | [optional] 
**ExternalIds** | Pointer to **map[string]interface{}** | The external IDs of the user | [optional] 
**Limits** | Pointer to [**UserLimits**](UserLimits.md) |  | [optional] 
**ExcludeFromReports** | **float32** | Whether the user is excluded from reports | 
**IsTestAccount** | **float32** | Whether the user is a test account | 
**PlanType** | **string** | The plan type of the user | [default to "vanilla"]
**Brand** | **string** | The brand of the user | [default to "default"]
**IsBrandManager** | **float32** | Whether the user is a brand manager | 
**IsDatastorePublisher** | **float32** | Whether the user is a datastore publisher | 
**Provider** | **string** | The provider of the user | [default to "default"]
**Archived** | **float32** | Whether the user is a archived | 
**AccountId** | **float32** | The account ID of the user | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewUser

`func NewUser(id string, revision string, franchise string, displayName string, email string, authenticatorMustChange float32, authenticatorCreatedTimestamp string, authenticatorEnabled bool, createdTimestamp string, lastLoginTimestamp string, lastLoginType string, blocked float32, emailStatus string, passwordChangeRequired float32, accessLevel string, isBillable float32, isTestingMode float32, isSuspended float32, language string, excludeFromReports float32, isTestAccount float32, planType string, brand string, isBrandManager float32, isDatastorePublisher float32, provider string, archived float32, accountId float32, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetRevision

`func (o *User) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *User) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *User) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetFranchise

`func (o *User) GetFranchise() string`

GetFranchise returns the Franchise field if non-nil, zero value otherwise.

### GetFranchiseOk

`func (o *User) GetFranchiseOk() (*string, bool)`

GetFranchiseOk returns a tuple with the Franchise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFranchise

`func (o *User) SetFranchise(v string)`

SetFranchise sets Franchise field to given value.


### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAuthenticatorMustChange

`func (o *User) GetAuthenticatorMustChange() float32`

GetAuthenticatorMustChange returns the AuthenticatorMustChange field if non-nil, zero value otherwise.

### GetAuthenticatorMustChangeOk

`func (o *User) GetAuthenticatorMustChangeOk() (*float32, bool)`

GetAuthenticatorMustChangeOk returns a tuple with the AuthenticatorMustChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorMustChange

`func (o *User) SetAuthenticatorMustChange(v float32)`

SetAuthenticatorMustChange sets AuthenticatorMustChange field to given value.


### GetAuthenticatorCreatedTimestamp

`func (o *User) GetAuthenticatorCreatedTimestamp() string`

GetAuthenticatorCreatedTimestamp returns the AuthenticatorCreatedTimestamp field if non-nil, zero value otherwise.

### GetAuthenticatorCreatedTimestampOk

`func (o *User) GetAuthenticatorCreatedTimestampOk() (*string, bool)`

GetAuthenticatorCreatedTimestampOk returns a tuple with the AuthenticatorCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCreatedTimestamp

`func (o *User) SetAuthenticatorCreatedTimestamp(v string)`

SetAuthenticatorCreatedTimestamp sets AuthenticatorCreatedTimestamp field to given value.


### GetAuthenticatorEnabled

`func (o *User) GetAuthenticatorEnabled() bool`

GetAuthenticatorEnabled returns the AuthenticatorEnabled field if non-nil, zero value otherwise.

### GetAuthenticatorEnabledOk

`func (o *User) GetAuthenticatorEnabledOk() (*bool, bool)`

GetAuthenticatorEnabledOk returns a tuple with the AuthenticatorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnabled

`func (o *User) SetAuthenticatorEnabled(v bool)`

SetAuthenticatorEnabled sets AuthenticatorEnabled field to given value.


### GetCreatedTimestamp

`func (o *User) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *User) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *User) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLastLoginTimestamp

`func (o *User) GetLastLoginTimestamp() string`

GetLastLoginTimestamp returns the LastLoginTimestamp field if non-nil, zero value otherwise.

### GetLastLoginTimestampOk

`func (o *User) GetLastLoginTimestampOk() (*string, bool)`

GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimestamp

`func (o *User) SetLastLoginTimestamp(v string)`

SetLastLoginTimestamp sets LastLoginTimestamp field to given value.


### GetLastLoginType

`func (o *User) GetLastLoginType() string`

GetLastLoginType returns the LastLoginType field if non-nil, zero value otherwise.

### GetLastLoginTypeOk

`func (o *User) GetLastLoginTypeOk() (*string, bool)`

GetLastLoginTypeOk returns a tuple with the LastLoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginType

`func (o *User) SetLastLoginType(v string)`

SetLastLoginType sets LastLoginType field to given value.


### GetBlocked

`func (o *User) GetBlocked() float32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *User) GetBlockedOk() (*float32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *User) SetBlocked(v float32)`

SetBlocked sets Blocked field to given value.


### GetEmailStatus

`func (o *User) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *User) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *User) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.


### GetPasswordChangeRequired

`func (o *User) GetPasswordChangeRequired() float32`

GetPasswordChangeRequired returns the PasswordChangeRequired field if non-nil, zero value otherwise.

### GetPasswordChangeRequiredOk

`func (o *User) GetPasswordChangeRequiredOk() (*float32, bool)`

GetPasswordChangeRequiredOk returns a tuple with the PasswordChangeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeRequired

`func (o *User) SetPasswordChangeRequired(v float32)`

SetPasswordChangeRequired sets PasswordChangeRequired field to given value.


### GetAccessLevel

`func (o *User) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *User) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *User) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetIsBillable

`func (o *User) GetIsBillable() float32`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *User) GetIsBillableOk() (*float32, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *User) SetIsBillable(v float32)`

SetIsBillable sets IsBillable field to given value.


### GetIsTestingMode

`func (o *User) GetIsTestingMode() float32`

GetIsTestingMode returns the IsTestingMode field if non-nil, zero value otherwise.

### GetIsTestingModeOk

`func (o *User) GetIsTestingModeOk() (*float32, bool)`

GetIsTestingModeOk returns a tuple with the IsTestingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestingMode

`func (o *User) SetIsTestingMode(v float32)`

SetIsTestingMode sets IsTestingMode field to given value.


### GetIsSuspended

`func (o *User) GetIsSuspended() float32`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *User) GetIsSuspendedOk() (*float32, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *User) SetIsSuspended(v float32)`

SetIsSuspended sets IsSuspended field to given value.


### GetLanguage

`func (o *User) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *User) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *User) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetInfrastructureIdDefault

`func (o *User) GetInfrastructureIdDefault() float32`

GetInfrastructureIdDefault returns the InfrastructureIdDefault field if non-nil, zero value otherwise.

### GetInfrastructureIdDefaultOk

`func (o *User) GetInfrastructureIdDefaultOk() (*float32, bool)`

GetInfrastructureIdDefaultOk returns a tuple with the InfrastructureIdDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureIdDefault

`func (o *User) SetInfrastructureIdDefault(v float32)`

SetInfrastructureIdDefault sets InfrastructureIdDefault field to given value.

### HasInfrastructureIdDefault

`func (o *User) HasInfrastructureIdDefault() bool`

HasInfrastructureIdDefault returns a boolean if a field has been set.

### GetGuiSettings

`func (o *User) GetGuiSettings() map[string]interface{}`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *User) GetGuiSettingsOk() (*map[string]interface{}, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *User) SetGuiSettings(v map[string]interface{})`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *User) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetPromotionTags

`func (o *User) GetPromotionTags() map[string]interface{}`

GetPromotionTags returns the PromotionTags field if non-nil, zero value otherwise.

### GetPromotionTagsOk

`func (o *User) GetPromotionTagsOk() (*map[string]interface{}, bool)`

GetPromotionTagsOk returns a tuple with the PromotionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionTags

`func (o *User) SetPromotionTags(v map[string]interface{})`

SetPromotionTags sets PromotionTags field to given value.

### HasPromotionTags

`func (o *User) HasPromotionTags() bool`

HasPromotionTags returns a boolean if a field has been set.

### GetExperimentalTags

`func (o *User) GetExperimentalTags() map[string]interface{}`

GetExperimentalTags returns the ExperimentalTags field if non-nil, zero value otherwise.

### GetExperimentalTagsOk

`func (o *User) GetExperimentalTagsOk() (*map[string]interface{}, bool)`

GetExperimentalTagsOk returns a tuple with the ExperimentalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalTags

`func (o *User) SetExperimentalTags(v map[string]interface{})`

SetExperimentalTags sets ExperimentalTags field to given value.

### HasExperimentalTags

`func (o *User) HasExperimentalTags() bool`

HasExperimentalTags returns a boolean if a field has been set.

### GetPermissions

`func (o *User) GetPermissions() UserPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *User) GetPermissionsOk() (*UserPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *User) SetPermissions(v UserPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *User) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetExternalIds

`func (o *User) GetExternalIds() map[string]interface{}`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *User) GetExternalIdsOk() (*map[string]interface{}, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *User) SetExternalIds(v map[string]interface{})`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *User) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetLimits

`func (o *User) GetLimits() UserLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *User) GetLimitsOk() (*UserLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *User) SetLimits(v UserLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *User) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetExcludeFromReports

`func (o *User) GetExcludeFromReports() float32`

GetExcludeFromReports returns the ExcludeFromReports field if non-nil, zero value otherwise.

### GetExcludeFromReportsOk

`func (o *User) GetExcludeFromReportsOk() (*float32, bool)`

GetExcludeFromReportsOk returns a tuple with the ExcludeFromReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromReports

`func (o *User) SetExcludeFromReports(v float32)`

SetExcludeFromReports sets ExcludeFromReports field to given value.


### GetIsTestAccount

`func (o *User) GetIsTestAccount() float32`

GetIsTestAccount returns the IsTestAccount field if non-nil, zero value otherwise.

### GetIsTestAccountOk

`func (o *User) GetIsTestAccountOk() (*float32, bool)`

GetIsTestAccountOk returns a tuple with the IsTestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestAccount

`func (o *User) SetIsTestAccount(v float32)`

SetIsTestAccount sets IsTestAccount field to given value.


### GetPlanType

`func (o *User) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *User) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *User) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetBrand

`func (o *User) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *User) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *User) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetIsBrandManager

`func (o *User) GetIsBrandManager() float32`

GetIsBrandManager returns the IsBrandManager field if non-nil, zero value otherwise.

### GetIsBrandManagerOk

`func (o *User) GetIsBrandManagerOk() (*float32, bool)`

GetIsBrandManagerOk returns a tuple with the IsBrandManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandManager

`func (o *User) SetIsBrandManager(v float32)`

SetIsBrandManager sets IsBrandManager field to given value.


### GetIsDatastorePublisher

`func (o *User) GetIsDatastorePublisher() float32`

GetIsDatastorePublisher returns the IsDatastorePublisher field if non-nil, zero value otherwise.

### GetIsDatastorePublisherOk

`func (o *User) GetIsDatastorePublisherOk() (*float32, bool)`

GetIsDatastorePublisherOk returns a tuple with the IsDatastorePublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDatastorePublisher

`func (o *User) SetIsDatastorePublisher(v float32)`

SetIsDatastorePublisher sets IsDatastorePublisher field to given value.


### GetProvider

`func (o *User) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *User) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *User) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetArchived

`func (o *User) GetArchived() float32`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *User) GetArchivedOk() (*float32, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *User) SetArchived(v float32)`

SetArchived sets Archived field to given value.


### GetAccountId

`func (o *User) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *User) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetLinks

`func (o *User) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *User) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *User) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *User) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



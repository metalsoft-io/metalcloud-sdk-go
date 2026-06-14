# UsersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsersProvider** | [**UsersProvider**](UsersProvider.md) | Enabled authentication provider flags | 
**DefaultUserProvider** | **string** | Default authentication provider for new users | 
**UserProviders** | [**[]UserProviderDto**](UserProviderDto.md) | Per-provider authentication configuration | 
**MaxFailedLoginAttempts** | **map[string]interface{}** | Maximum consecutive failed login attempts before temporary lockout | 
**FailedLoginAttemptsTTL** | **map[string]interface{}** | Duration in seconds for the failed login attempt counter to reset | 
**UserCacheTTL** | **map[string]interface{}** | Duration in seconds to cache user data in memory | 
**MaxUserSessionDuration** | **map[string]interface{}** | Maximum session duration in seconds before forced re-authentication | 
**SessionInactivityTimeoutDuration** | **map[string]interface{}** | Idle timeout in seconds after which an inactive session expires | 
**PasswordRegex** | **map[string]interface{}** | Regular expression pattern that passwords must match | 
**PasswordAllowedSpecialCharacters** | **map[string]interface{}** | Characters allowed as special characters in passwords | 
**PasswordMinLength** | **map[string]interface{}** | Minimum password length | 
**PasswordMaxLength** | **map[string]interface{}** | Maximum password length | 
**PasswordMinSpecialCharacters** | **map[string]interface{}** | Minimum number of special characters required in passwords | 
**PasswordMinUppercaseCharacters** | **map[string]interface{}** | Minimum number of uppercase characters required in passwords | 
**PasswordMinLowercaseCharacters** | **map[string]interface{}** | Minimum number of lowercase characters required in passwords | 
**PasswordMinDigits** | **map[string]interface{}** | Minimum number of digits required in passwords | 
**PasswordReuseLockoutDuration** | **map[string]interface{}** | Duration in seconds during which a previously used password cannot be reused | 
**PasswordValidityDuration** | **map[string]interface{}** | Duration in seconds after which passwords expire; 0 means passwords never expire | 
**CookieExpirationBufferDuration** | **map[string]interface{}** | Seconds added to the JWT exp beyond the session expiration to absorb client/server clock skew; kept small because session invalidation is enforced server-side | 
**CaptchaEnabled** | **map[string]interface{}** | Whether CAPTCHA verification is required on the login page | 
**CaptchaVerifyUrl** | **map[string]interface{}** | CAPTCHA server-side verification endpoint URL | 
**CaptchaWidgets** | Pointer to **map[string]interface{}** | CAPTCHA widget configurations keyed by widget identifier | [optional] 

## Methods

### NewUsersDto

`func NewUsersDto(usersProvider UsersProvider, defaultUserProvider string, userProviders []UserProviderDto, maxFailedLoginAttempts map[string]interface{}, failedLoginAttemptsTTL map[string]interface{}, userCacheTTL map[string]interface{}, maxUserSessionDuration map[string]interface{}, sessionInactivityTimeoutDuration map[string]interface{}, passwordRegex map[string]interface{}, passwordAllowedSpecialCharacters map[string]interface{}, passwordMinLength map[string]interface{}, passwordMaxLength map[string]interface{}, passwordMinSpecialCharacters map[string]interface{}, passwordMinUppercaseCharacters map[string]interface{}, passwordMinLowercaseCharacters map[string]interface{}, passwordMinDigits map[string]interface{}, passwordReuseLockoutDuration map[string]interface{}, passwordValidityDuration map[string]interface{}, cookieExpirationBufferDuration map[string]interface{}, captchaEnabled map[string]interface{}, captchaVerifyUrl map[string]interface{}, ) *UsersDto`

NewUsersDto instantiates a new UsersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersDtoWithDefaults

`func NewUsersDtoWithDefaults() *UsersDto`

NewUsersDtoWithDefaults instantiates a new UsersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsersProvider

`func (o *UsersDto) GetUsersProvider() UsersProvider`

GetUsersProvider returns the UsersProvider field if non-nil, zero value otherwise.

### GetUsersProviderOk

`func (o *UsersDto) GetUsersProviderOk() (*UsersProvider, bool)`

GetUsersProviderOk returns a tuple with the UsersProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersProvider

`func (o *UsersDto) SetUsersProvider(v UsersProvider)`

SetUsersProvider sets UsersProvider field to given value.


### GetDefaultUserProvider

`func (o *UsersDto) GetDefaultUserProvider() string`

GetDefaultUserProvider returns the DefaultUserProvider field if non-nil, zero value otherwise.

### GetDefaultUserProviderOk

`func (o *UsersDto) GetDefaultUserProviderOk() (*string, bool)`

GetDefaultUserProviderOk returns a tuple with the DefaultUserProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserProvider

`func (o *UsersDto) SetDefaultUserProvider(v string)`

SetDefaultUserProvider sets DefaultUserProvider field to given value.


### GetUserProviders

`func (o *UsersDto) GetUserProviders() []UserProviderDto`

GetUserProviders returns the UserProviders field if non-nil, zero value otherwise.

### GetUserProvidersOk

`func (o *UsersDto) GetUserProvidersOk() (*[]UserProviderDto, bool)`

GetUserProvidersOk returns a tuple with the UserProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProviders

`func (o *UsersDto) SetUserProviders(v []UserProviderDto)`

SetUserProviders sets UserProviders field to given value.


### GetMaxFailedLoginAttempts

`func (o *UsersDto) GetMaxFailedLoginAttempts() map[string]interface{}`

GetMaxFailedLoginAttempts returns the MaxFailedLoginAttempts field if non-nil, zero value otherwise.

### GetMaxFailedLoginAttemptsOk

`func (o *UsersDto) GetMaxFailedLoginAttemptsOk() (*map[string]interface{}, bool)`

GetMaxFailedLoginAttemptsOk returns a tuple with the MaxFailedLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailedLoginAttempts

`func (o *UsersDto) SetMaxFailedLoginAttempts(v map[string]interface{})`

SetMaxFailedLoginAttempts sets MaxFailedLoginAttempts field to given value.


### GetFailedLoginAttemptsTTL

`func (o *UsersDto) GetFailedLoginAttemptsTTL() map[string]interface{}`

GetFailedLoginAttemptsTTL returns the FailedLoginAttemptsTTL field if non-nil, zero value otherwise.

### GetFailedLoginAttemptsTTLOk

`func (o *UsersDto) GetFailedLoginAttemptsTTLOk() (*map[string]interface{}, bool)`

GetFailedLoginAttemptsTTLOk returns a tuple with the FailedLoginAttemptsTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginAttemptsTTL

`func (o *UsersDto) SetFailedLoginAttemptsTTL(v map[string]interface{})`

SetFailedLoginAttemptsTTL sets FailedLoginAttemptsTTL field to given value.


### GetUserCacheTTL

`func (o *UsersDto) GetUserCacheTTL() map[string]interface{}`

GetUserCacheTTL returns the UserCacheTTL field if non-nil, zero value otherwise.

### GetUserCacheTTLOk

`func (o *UsersDto) GetUserCacheTTLOk() (*map[string]interface{}, bool)`

GetUserCacheTTLOk returns a tuple with the UserCacheTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCacheTTL

`func (o *UsersDto) SetUserCacheTTL(v map[string]interface{})`

SetUserCacheTTL sets UserCacheTTL field to given value.


### GetMaxUserSessionDuration

`func (o *UsersDto) GetMaxUserSessionDuration() map[string]interface{}`

GetMaxUserSessionDuration returns the MaxUserSessionDuration field if non-nil, zero value otherwise.

### GetMaxUserSessionDurationOk

`func (o *UsersDto) GetMaxUserSessionDurationOk() (*map[string]interface{}, bool)`

GetMaxUserSessionDurationOk returns a tuple with the MaxUserSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserSessionDuration

`func (o *UsersDto) SetMaxUserSessionDuration(v map[string]interface{})`

SetMaxUserSessionDuration sets MaxUserSessionDuration field to given value.


### GetSessionInactivityTimeoutDuration

`func (o *UsersDto) GetSessionInactivityTimeoutDuration() map[string]interface{}`

GetSessionInactivityTimeoutDuration returns the SessionInactivityTimeoutDuration field if non-nil, zero value otherwise.

### GetSessionInactivityTimeoutDurationOk

`func (o *UsersDto) GetSessionInactivityTimeoutDurationOk() (*map[string]interface{}, bool)`

GetSessionInactivityTimeoutDurationOk returns a tuple with the SessionInactivityTimeoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInactivityTimeoutDuration

`func (o *UsersDto) SetSessionInactivityTimeoutDuration(v map[string]interface{})`

SetSessionInactivityTimeoutDuration sets SessionInactivityTimeoutDuration field to given value.


### GetPasswordRegex

`func (o *UsersDto) GetPasswordRegex() map[string]interface{}`

GetPasswordRegex returns the PasswordRegex field if non-nil, zero value otherwise.

### GetPasswordRegexOk

`func (o *UsersDto) GetPasswordRegexOk() (*map[string]interface{}, bool)`

GetPasswordRegexOk returns a tuple with the PasswordRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRegex

`func (o *UsersDto) SetPasswordRegex(v map[string]interface{})`

SetPasswordRegex sets PasswordRegex field to given value.


### GetPasswordAllowedSpecialCharacters

`func (o *UsersDto) GetPasswordAllowedSpecialCharacters() map[string]interface{}`

GetPasswordAllowedSpecialCharacters returns the PasswordAllowedSpecialCharacters field if non-nil, zero value otherwise.

### GetPasswordAllowedSpecialCharactersOk

`func (o *UsersDto) GetPasswordAllowedSpecialCharactersOk() (*map[string]interface{}, bool)`

GetPasswordAllowedSpecialCharactersOk returns a tuple with the PasswordAllowedSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAllowedSpecialCharacters

`func (o *UsersDto) SetPasswordAllowedSpecialCharacters(v map[string]interface{})`

SetPasswordAllowedSpecialCharacters sets PasswordAllowedSpecialCharacters field to given value.


### GetPasswordMinLength

`func (o *UsersDto) GetPasswordMinLength() map[string]interface{}`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *UsersDto) GetPasswordMinLengthOk() (*map[string]interface{}, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *UsersDto) SetPasswordMinLength(v map[string]interface{})`

SetPasswordMinLength sets PasswordMinLength field to given value.


### GetPasswordMaxLength

`func (o *UsersDto) GetPasswordMaxLength() map[string]interface{}`

GetPasswordMaxLength returns the PasswordMaxLength field if non-nil, zero value otherwise.

### GetPasswordMaxLengthOk

`func (o *UsersDto) GetPasswordMaxLengthOk() (*map[string]interface{}, bool)`

GetPasswordMaxLengthOk returns a tuple with the PasswordMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMaxLength

`func (o *UsersDto) SetPasswordMaxLength(v map[string]interface{})`

SetPasswordMaxLength sets PasswordMaxLength field to given value.


### GetPasswordMinSpecialCharacters

`func (o *UsersDto) GetPasswordMinSpecialCharacters() map[string]interface{}`

GetPasswordMinSpecialCharacters returns the PasswordMinSpecialCharacters field if non-nil, zero value otherwise.

### GetPasswordMinSpecialCharactersOk

`func (o *UsersDto) GetPasswordMinSpecialCharactersOk() (*map[string]interface{}, bool)`

GetPasswordMinSpecialCharactersOk returns a tuple with the PasswordMinSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinSpecialCharacters

`func (o *UsersDto) SetPasswordMinSpecialCharacters(v map[string]interface{})`

SetPasswordMinSpecialCharacters sets PasswordMinSpecialCharacters field to given value.


### GetPasswordMinUppercaseCharacters

`func (o *UsersDto) GetPasswordMinUppercaseCharacters() map[string]interface{}`

GetPasswordMinUppercaseCharacters returns the PasswordMinUppercaseCharacters field if non-nil, zero value otherwise.

### GetPasswordMinUppercaseCharactersOk

`func (o *UsersDto) GetPasswordMinUppercaseCharactersOk() (*map[string]interface{}, bool)`

GetPasswordMinUppercaseCharactersOk returns a tuple with the PasswordMinUppercaseCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinUppercaseCharacters

`func (o *UsersDto) SetPasswordMinUppercaseCharacters(v map[string]interface{})`

SetPasswordMinUppercaseCharacters sets PasswordMinUppercaseCharacters field to given value.


### GetPasswordMinLowercaseCharacters

`func (o *UsersDto) GetPasswordMinLowercaseCharacters() map[string]interface{}`

GetPasswordMinLowercaseCharacters returns the PasswordMinLowercaseCharacters field if non-nil, zero value otherwise.

### GetPasswordMinLowercaseCharactersOk

`func (o *UsersDto) GetPasswordMinLowercaseCharactersOk() (*map[string]interface{}, bool)`

GetPasswordMinLowercaseCharactersOk returns a tuple with the PasswordMinLowercaseCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLowercaseCharacters

`func (o *UsersDto) SetPasswordMinLowercaseCharacters(v map[string]interface{})`

SetPasswordMinLowercaseCharacters sets PasswordMinLowercaseCharacters field to given value.


### GetPasswordMinDigits

`func (o *UsersDto) GetPasswordMinDigits() map[string]interface{}`

GetPasswordMinDigits returns the PasswordMinDigits field if non-nil, zero value otherwise.

### GetPasswordMinDigitsOk

`func (o *UsersDto) GetPasswordMinDigitsOk() (*map[string]interface{}, bool)`

GetPasswordMinDigitsOk returns a tuple with the PasswordMinDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinDigits

`func (o *UsersDto) SetPasswordMinDigits(v map[string]interface{})`

SetPasswordMinDigits sets PasswordMinDigits field to given value.


### GetPasswordReuseLockoutDuration

`func (o *UsersDto) GetPasswordReuseLockoutDuration() map[string]interface{}`

GetPasswordReuseLockoutDuration returns the PasswordReuseLockoutDuration field if non-nil, zero value otherwise.

### GetPasswordReuseLockoutDurationOk

`func (o *UsersDto) GetPasswordReuseLockoutDurationOk() (*map[string]interface{}, bool)`

GetPasswordReuseLockoutDurationOk returns a tuple with the PasswordReuseLockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReuseLockoutDuration

`func (o *UsersDto) SetPasswordReuseLockoutDuration(v map[string]interface{})`

SetPasswordReuseLockoutDuration sets PasswordReuseLockoutDuration field to given value.


### GetPasswordValidityDuration

`func (o *UsersDto) GetPasswordValidityDuration() map[string]interface{}`

GetPasswordValidityDuration returns the PasswordValidityDuration field if non-nil, zero value otherwise.

### GetPasswordValidityDurationOk

`func (o *UsersDto) GetPasswordValidityDurationOk() (*map[string]interface{}, bool)`

GetPasswordValidityDurationOk returns a tuple with the PasswordValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordValidityDuration

`func (o *UsersDto) SetPasswordValidityDuration(v map[string]interface{})`

SetPasswordValidityDuration sets PasswordValidityDuration field to given value.


### GetCookieExpirationBufferDuration

`func (o *UsersDto) GetCookieExpirationBufferDuration() map[string]interface{}`

GetCookieExpirationBufferDuration returns the CookieExpirationBufferDuration field if non-nil, zero value otherwise.

### GetCookieExpirationBufferDurationOk

`func (o *UsersDto) GetCookieExpirationBufferDurationOk() (*map[string]interface{}, bool)`

GetCookieExpirationBufferDurationOk returns a tuple with the CookieExpirationBufferDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieExpirationBufferDuration

`func (o *UsersDto) SetCookieExpirationBufferDuration(v map[string]interface{})`

SetCookieExpirationBufferDuration sets CookieExpirationBufferDuration field to given value.


### GetCaptchaEnabled

`func (o *UsersDto) GetCaptchaEnabled() map[string]interface{}`

GetCaptchaEnabled returns the CaptchaEnabled field if non-nil, zero value otherwise.

### GetCaptchaEnabledOk

`func (o *UsersDto) GetCaptchaEnabledOk() (*map[string]interface{}, bool)`

GetCaptchaEnabledOk returns a tuple with the CaptchaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaEnabled

`func (o *UsersDto) SetCaptchaEnabled(v map[string]interface{})`

SetCaptchaEnabled sets CaptchaEnabled field to given value.


### GetCaptchaVerifyUrl

`func (o *UsersDto) GetCaptchaVerifyUrl() map[string]interface{}`

GetCaptchaVerifyUrl returns the CaptchaVerifyUrl field if non-nil, zero value otherwise.

### GetCaptchaVerifyUrlOk

`func (o *UsersDto) GetCaptchaVerifyUrlOk() (*map[string]interface{}, bool)`

GetCaptchaVerifyUrlOk returns a tuple with the CaptchaVerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaVerifyUrl

`func (o *UsersDto) SetCaptchaVerifyUrl(v map[string]interface{})`

SetCaptchaVerifyUrl sets CaptchaVerifyUrl field to given value.


### GetCaptchaWidgets

`func (o *UsersDto) GetCaptchaWidgets() map[string]interface{}`

GetCaptchaWidgets returns the CaptchaWidgets field if non-nil, zero value otherwise.

### GetCaptchaWidgetsOk

`func (o *UsersDto) GetCaptchaWidgetsOk() (*map[string]interface{}, bool)`

GetCaptchaWidgetsOk returns a tuple with the CaptchaWidgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaWidgets

`func (o *UsersDto) SetCaptchaWidgets(v map[string]interface{})`

SetCaptchaWidgets sets CaptchaWidgets field to given value.

### HasCaptchaWidgets

`func (o *UsersDto) HasCaptchaWidgets() bool`

HasCaptchaWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



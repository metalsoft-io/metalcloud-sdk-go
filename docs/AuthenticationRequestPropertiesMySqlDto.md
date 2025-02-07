# AuthenticationRequestPropertiesMySqlDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneTimePassword** | Pointer to **string** | The one-time password. | [optional] 
**RememberLogin** | Pointer to **bool** | Whether to remember the login | [optional] [default to true]
**TestCredentials** | Pointer to **bool** | Whether to test the credentials | [optional] [default to false]
**Email** | **string** | The email of the user | 
**Password** | **string** | The password of the user | 

## Methods

### NewAuthenticationRequestPropertiesMySqlDto

`func NewAuthenticationRequestPropertiesMySqlDto(email string, password string, ) *AuthenticationRequestPropertiesMySqlDto`

NewAuthenticationRequestPropertiesMySqlDto instantiates a new AuthenticationRequestPropertiesMySqlDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationRequestPropertiesMySqlDtoWithDefaults

`func NewAuthenticationRequestPropertiesMySqlDtoWithDefaults() *AuthenticationRequestPropertiesMySqlDto`

NewAuthenticationRequestPropertiesMySqlDtoWithDefaults instantiates a new AuthenticationRequestPropertiesMySqlDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneTimePassword

`func (o *AuthenticationRequestPropertiesMySqlDto) GetOneTimePassword() string`

GetOneTimePassword returns the OneTimePassword field if non-nil, zero value otherwise.

### GetOneTimePasswordOk

`func (o *AuthenticationRequestPropertiesMySqlDto) GetOneTimePasswordOk() (*string, bool)`

GetOneTimePasswordOk returns a tuple with the OneTimePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePassword

`func (o *AuthenticationRequestPropertiesMySqlDto) SetOneTimePassword(v string)`

SetOneTimePassword sets OneTimePassword field to given value.

### HasOneTimePassword

`func (o *AuthenticationRequestPropertiesMySqlDto) HasOneTimePassword() bool`

HasOneTimePassword returns a boolean if a field has been set.

### GetRememberLogin

`func (o *AuthenticationRequestPropertiesMySqlDto) GetRememberLogin() bool`

GetRememberLogin returns the RememberLogin field if non-nil, zero value otherwise.

### GetRememberLoginOk

`func (o *AuthenticationRequestPropertiesMySqlDto) GetRememberLoginOk() (*bool, bool)`

GetRememberLoginOk returns a tuple with the RememberLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberLogin

`func (o *AuthenticationRequestPropertiesMySqlDto) SetRememberLogin(v bool)`

SetRememberLogin sets RememberLogin field to given value.

### HasRememberLogin

`func (o *AuthenticationRequestPropertiesMySqlDto) HasRememberLogin() bool`

HasRememberLogin returns a boolean if a field has been set.

### GetTestCredentials

`func (o *AuthenticationRequestPropertiesMySqlDto) GetTestCredentials() bool`

GetTestCredentials returns the TestCredentials field if non-nil, zero value otherwise.

### GetTestCredentialsOk

`func (o *AuthenticationRequestPropertiesMySqlDto) GetTestCredentialsOk() (*bool, bool)`

GetTestCredentialsOk returns a tuple with the TestCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCredentials

`func (o *AuthenticationRequestPropertiesMySqlDto) SetTestCredentials(v bool)`

SetTestCredentials sets TestCredentials field to given value.

### HasTestCredentials

`func (o *AuthenticationRequestPropertiesMySqlDto) HasTestCredentials() bool`

HasTestCredentials returns a boolean if a field has been set.

### GetEmail

`func (o *AuthenticationRequestPropertiesMySqlDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthenticationRequestPropertiesMySqlDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthenticationRequestPropertiesMySqlDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AuthenticationRequestPropertiesMySqlDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthenticationRequestPropertiesMySqlDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthenticationRequestPropertiesMySqlDto) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AuthenticationRequestPropertiesLdapDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneTimePassword** | Pointer to **string** | The one-time password. | [optional] 
**RememberLogin** | Pointer to **bool** | Whether to remember the login | [optional] [default to true]
**TestCredentials** | Pointer to **bool** | Whether to test the credentials | [optional] [default to false]
**Data** | **map[string]interface{}** | The data for the LDAP authentication request. | 

## Methods

### NewAuthenticationRequestPropertiesLdapDto

`func NewAuthenticationRequestPropertiesLdapDto(data map[string]interface{}, ) *AuthenticationRequestPropertiesLdapDto`

NewAuthenticationRequestPropertiesLdapDto instantiates a new AuthenticationRequestPropertiesLdapDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationRequestPropertiesLdapDtoWithDefaults

`func NewAuthenticationRequestPropertiesLdapDtoWithDefaults() *AuthenticationRequestPropertiesLdapDto`

NewAuthenticationRequestPropertiesLdapDtoWithDefaults instantiates a new AuthenticationRequestPropertiesLdapDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneTimePassword

`func (o *AuthenticationRequestPropertiesLdapDto) GetOneTimePassword() string`

GetOneTimePassword returns the OneTimePassword field if non-nil, zero value otherwise.

### GetOneTimePasswordOk

`func (o *AuthenticationRequestPropertiesLdapDto) GetOneTimePasswordOk() (*string, bool)`

GetOneTimePasswordOk returns a tuple with the OneTimePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePassword

`func (o *AuthenticationRequestPropertiesLdapDto) SetOneTimePassword(v string)`

SetOneTimePassword sets OneTimePassword field to given value.

### HasOneTimePassword

`func (o *AuthenticationRequestPropertiesLdapDto) HasOneTimePassword() bool`

HasOneTimePassword returns a boolean if a field has been set.

### GetRememberLogin

`func (o *AuthenticationRequestPropertiesLdapDto) GetRememberLogin() bool`

GetRememberLogin returns the RememberLogin field if non-nil, zero value otherwise.

### GetRememberLoginOk

`func (o *AuthenticationRequestPropertiesLdapDto) GetRememberLoginOk() (*bool, bool)`

GetRememberLoginOk returns a tuple with the RememberLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberLogin

`func (o *AuthenticationRequestPropertiesLdapDto) SetRememberLogin(v bool)`

SetRememberLogin sets RememberLogin field to given value.

### HasRememberLogin

`func (o *AuthenticationRequestPropertiesLdapDto) HasRememberLogin() bool`

HasRememberLogin returns a boolean if a field has been set.

### GetTestCredentials

`func (o *AuthenticationRequestPropertiesLdapDto) GetTestCredentials() bool`

GetTestCredentials returns the TestCredentials field if non-nil, zero value otherwise.

### GetTestCredentialsOk

`func (o *AuthenticationRequestPropertiesLdapDto) GetTestCredentialsOk() (*bool, bool)`

GetTestCredentialsOk returns a tuple with the TestCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCredentials

`func (o *AuthenticationRequestPropertiesLdapDto) SetTestCredentials(v bool)`

SetTestCredentials sets TestCredentials field to given value.

### HasTestCredentials

`func (o *AuthenticationRequestPropertiesLdapDto) HasTestCredentials() bool`

HasTestCredentials returns a boolean if a field has been set.

### GetData

`func (o *AuthenticationRequestPropertiesLdapDto) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthenticationRequestPropertiesLdapDto) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthenticationRequestPropertiesLdapDto) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



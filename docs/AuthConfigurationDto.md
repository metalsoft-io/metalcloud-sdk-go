# AuthConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**UsersDto**](UsersDto.md) | User management and authentication provider settings | [optional] 
**Ldap** | Pointer to [**LdapDto**](LdapDto.md) | LDAP directory integration settings | [optional] 
**Saml** | Pointer to [**SamlDto**](SamlDto.md) | SAML SSO integration settings | [optional] 
**Oauth** | Pointer to [**OauthDto**](OauthDto.md) | OAuth/OIDC token verification settings | [optional] 
**Microservices** | Pointer to [**AuthMicroservicesDto**](AuthMicroservicesDto.md) | Connected microservices configuration | [optional] 
**OauthWhitelistedPermissions** | Pointer to **[]string** | Permissions automatically granted to OAuth token holders regardless of their token scopes | [optional] 

## Methods

### NewAuthConfigurationDto

`func NewAuthConfigurationDto() *AuthConfigurationDto`

NewAuthConfigurationDto instantiates a new AuthConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigurationDtoWithDefaults

`func NewAuthConfigurationDtoWithDefaults() *AuthConfigurationDto`

NewAuthConfigurationDtoWithDefaults instantiates a new AuthConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *AuthConfigurationDto) GetUsers() UsersDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AuthConfigurationDto) GetUsersOk() (*UsersDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AuthConfigurationDto) SetUsers(v UsersDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AuthConfigurationDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetLdap

`func (o *AuthConfigurationDto) GetLdap() LdapDto`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *AuthConfigurationDto) GetLdapOk() (*LdapDto, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *AuthConfigurationDto) SetLdap(v LdapDto)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *AuthConfigurationDto) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSaml

`func (o *AuthConfigurationDto) GetSaml() SamlDto`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *AuthConfigurationDto) GetSamlOk() (*SamlDto, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *AuthConfigurationDto) SetSaml(v SamlDto)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *AuthConfigurationDto) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOauth

`func (o *AuthConfigurationDto) GetOauth() OauthDto`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *AuthConfigurationDto) GetOauthOk() (*OauthDto, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *AuthConfigurationDto) SetOauth(v OauthDto)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *AuthConfigurationDto) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetMicroservices

`func (o *AuthConfigurationDto) GetMicroservices() AuthMicroservicesDto`

GetMicroservices returns the Microservices field if non-nil, zero value otherwise.

### GetMicroservicesOk

`func (o *AuthConfigurationDto) GetMicroservicesOk() (*AuthMicroservicesDto, bool)`

GetMicroservicesOk returns a tuple with the Microservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroservices

`func (o *AuthConfigurationDto) SetMicroservices(v AuthMicroservicesDto)`

SetMicroservices sets Microservices field to given value.

### HasMicroservices

`func (o *AuthConfigurationDto) HasMicroservices() bool`

HasMicroservices returns a boolean if a field has been set.

### GetOauthWhitelistedPermissions

`func (o *AuthConfigurationDto) GetOauthWhitelistedPermissions() []string`

GetOauthWhitelistedPermissions returns the OauthWhitelistedPermissions field if non-nil, zero value otherwise.

### GetOauthWhitelistedPermissionsOk

`func (o *AuthConfigurationDto) GetOauthWhitelistedPermissionsOk() (*[]string, bool)`

GetOauthWhitelistedPermissionsOk returns a tuple with the OauthWhitelistedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthWhitelistedPermissions

`func (o *AuthConfigurationDto) SetOauthWhitelistedPermissions(v []string)`

SetOauthWhitelistedPermissions sets OauthWhitelistedPermissions field to given value.

### HasOauthWhitelistedPermissions

`func (o *AuthConfigurationDto) HasOauthWhitelistedPermissions() bool`

HasOauthWhitelistedPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



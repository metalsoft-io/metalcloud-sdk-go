# AuthConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**Users**](Users.md) | User management and authentication provider settings | [optional] 
**Ldap** | Pointer to [**Ldap**](Ldap.md) | LDAP directory integration settings | [optional] 
**Saml** | Pointer to [**Saml**](Saml.md) | SAML SSO integration settings | [optional] 

## Methods

### NewAuthConfiguration

`func NewAuthConfiguration() *AuthConfiguration`

NewAuthConfiguration instantiates a new AuthConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigurationWithDefaults

`func NewAuthConfigurationWithDefaults() *AuthConfiguration`

NewAuthConfigurationWithDefaults instantiates a new AuthConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *AuthConfiguration) GetUsers() Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AuthConfiguration) GetUsersOk() (*Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AuthConfiguration) SetUsers(v Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AuthConfiguration) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetLdap

`func (o *AuthConfiguration) GetLdap() Ldap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *AuthConfiguration) GetLdapOk() (*Ldap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *AuthConfiguration) SetLdap(v Ldap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *AuthConfiguration) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSaml

`func (o *AuthConfiguration) GetSaml() Saml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *AuthConfiguration) GetSamlOk() (*Saml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *AuthConfiguration) SetSaml(v Saml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *AuthConfiguration) HasSaml() bool`

HasSaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



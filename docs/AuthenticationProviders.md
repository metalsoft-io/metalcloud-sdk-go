# AuthenticationProviders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ldap** | **bool** | Whether or not the application allows LDAP authentication | 
**Mysql** | **bool** | Whether or not the application allows MYSQL authentication | 
**Saml** | **bool** | Whether or not the application allows SAML authentication | 
**SsoEntryPoint** | **string** | The entry point for SAML authentication | 

## Methods

### NewAuthenticationProviders

`func NewAuthenticationProviders(ldap bool, mysql bool, saml bool, ssoEntryPoint string, ) *AuthenticationProviders`

NewAuthenticationProviders instantiates a new AuthenticationProviders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProvidersWithDefaults

`func NewAuthenticationProvidersWithDefaults() *AuthenticationProviders`

NewAuthenticationProvidersWithDefaults instantiates a new AuthenticationProviders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdap

`func (o *AuthenticationProviders) GetLdap() bool`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *AuthenticationProviders) GetLdapOk() (*bool, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *AuthenticationProviders) SetLdap(v bool)`

SetLdap sets Ldap field to given value.


### GetMysql

`func (o *AuthenticationProviders) GetMysql() bool`

GetMysql returns the Mysql field if non-nil, zero value otherwise.

### GetMysqlOk

`func (o *AuthenticationProviders) GetMysqlOk() (*bool, bool)`

GetMysqlOk returns a tuple with the Mysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql

`func (o *AuthenticationProviders) SetMysql(v bool)`

SetMysql sets Mysql field to given value.


### GetSaml

`func (o *AuthenticationProviders) GetSaml() bool`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *AuthenticationProviders) GetSamlOk() (*bool, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *AuthenticationProviders) SetSaml(v bool)`

SetSaml sets Saml field to given value.


### GetSsoEntryPoint

`func (o *AuthenticationProviders) GetSsoEntryPoint() string`

GetSsoEntryPoint returns the SsoEntryPoint field if non-nil, zero value otherwise.

### GetSsoEntryPointOk

`func (o *AuthenticationProviders) GetSsoEntryPointOk() (*string, bool)`

GetSsoEntryPointOk returns a tuple with the SsoEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEntryPoint

`func (o *AuthenticationProviders) SetSsoEntryPoint(v string)`

SetSsoEntryPoint sets SsoEntryPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



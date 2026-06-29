# LdapServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **map[string]interface{}** | LDAP server URL (e.g. ldap://host:389 or ldaps://host:636) | [optional] 
**SearchBase** | Pointer to **map[string]interface{}** | Base DN for user searches | [optional] 
**SearchFilter** | Pointer to **map[string]interface{}** | LDAP filter for locating user entries | [optional] 
**GroupSearchBase** | Pointer to **map[string]interface{}** | Base DN for group searches | [optional] 
**GroupSearchFilter** | Pointer to **map[string]interface{}** | LDAP filter for locating group entries | [optional] 
**BindDN** | Pointer to **map[string]interface{}** | DN of the service account used to bind to the LDAP directory | [optional] 
**BindCredentials** | Pointer to **map[string]interface{}** | Password for the LDAP bind account | [optional] 
**Certificate** | Pointer to **map[string]interface{}** | PEM-encoded TLS certificate for LDAPS connections | [optional] 

## Methods

### NewLdapServer

`func NewLdapServer() *LdapServer`

NewLdapServer instantiates a new LdapServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapServerWithDefaults

`func NewLdapServerWithDefaults() *LdapServer`

NewLdapServerWithDefaults instantiates a new LdapServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LdapServer) GetUrl() map[string]interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LdapServer) GetUrlOk() (*map[string]interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LdapServer) SetUrl(v map[string]interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LdapServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSearchBase

`func (o *LdapServer) GetSearchBase() map[string]interface{}`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *LdapServer) GetSearchBaseOk() (*map[string]interface{}, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *LdapServer) SetSearchBase(v map[string]interface{})`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *LdapServer) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetSearchFilter

`func (o *LdapServer) GetSearchFilter() map[string]interface{}`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *LdapServer) GetSearchFilterOk() (*map[string]interface{}, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *LdapServer) SetSearchFilter(v map[string]interface{})`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *LdapServer) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### GetGroupSearchBase

`func (o *LdapServer) GetGroupSearchBase() map[string]interface{}`

GetGroupSearchBase returns the GroupSearchBase field if non-nil, zero value otherwise.

### GetGroupSearchBaseOk

`func (o *LdapServer) GetGroupSearchBaseOk() (*map[string]interface{}, bool)`

GetGroupSearchBaseOk returns a tuple with the GroupSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBase

`func (o *LdapServer) SetGroupSearchBase(v map[string]interface{})`

SetGroupSearchBase sets GroupSearchBase field to given value.

### HasGroupSearchBase

`func (o *LdapServer) HasGroupSearchBase() bool`

HasGroupSearchBase returns a boolean if a field has been set.

### GetGroupSearchFilter

`func (o *LdapServer) GetGroupSearchFilter() map[string]interface{}`

GetGroupSearchFilter returns the GroupSearchFilter field if non-nil, zero value otherwise.

### GetGroupSearchFilterOk

`func (o *LdapServer) GetGroupSearchFilterOk() (*map[string]interface{}, bool)`

GetGroupSearchFilterOk returns a tuple with the GroupSearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchFilter

`func (o *LdapServer) SetGroupSearchFilter(v map[string]interface{})`

SetGroupSearchFilter sets GroupSearchFilter field to given value.

### HasGroupSearchFilter

`func (o *LdapServer) HasGroupSearchFilter() bool`

HasGroupSearchFilter returns a boolean if a field has been set.

### GetBindDN

`func (o *LdapServer) GetBindDN() map[string]interface{}`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *LdapServer) GetBindDNOk() (*map[string]interface{}, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *LdapServer) SetBindDN(v map[string]interface{})`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *LdapServer) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetBindCredentials

`func (o *LdapServer) GetBindCredentials() map[string]interface{}`

GetBindCredentials returns the BindCredentials field if non-nil, zero value otherwise.

### GetBindCredentialsOk

`func (o *LdapServer) GetBindCredentialsOk() (*map[string]interface{}, bool)`

GetBindCredentialsOk returns a tuple with the BindCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCredentials

`func (o *LdapServer) SetBindCredentials(v map[string]interface{})`

SetBindCredentials sets BindCredentials field to given value.

### HasBindCredentials

`func (o *LdapServer) HasBindCredentials() bool`

HasBindCredentials returns a boolean if a field has been set.

### GetCertificate

`func (o *LdapServer) GetCertificate() map[string]interface{}`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LdapServer) GetCertificateOk() (*map[string]interface{}, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LdapServer) SetCertificate(v map[string]interface{})`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LdapServer) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LdapServerDto

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

### NewLdapServerDto

`func NewLdapServerDto() *LdapServerDto`

NewLdapServerDto instantiates a new LdapServerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapServerDtoWithDefaults

`func NewLdapServerDtoWithDefaults() *LdapServerDto`

NewLdapServerDtoWithDefaults instantiates a new LdapServerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LdapServerDto) GetUrl() map[string]interface{}`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LdapServerDto) GetUrlOk() (*map[string]interface{}, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LdapServerDto) SetUrl(v map[string]interface{})`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LdapServerDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSearchBase

`func (o *LdapServerDto) GetSearchBase() map[string]interface{}`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *LdapServerDto) GetSearchBaseOk() (*map[string]interface{}, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *LdapServerDto) SetSearchBase(v map[string]interface{})`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *LdapServerDto) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetSearchFilter

`func (o *LdapServerDto) GetSearchFilter() map[string]interface{}`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *LdapServerDto) GetSearchFilterOk() (*map[string]interface{}, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *LdapServerDto) SetSearchFilter(v map[string]interface{})`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *LdapServerDto) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### GetGroupSearchBase

`func (o *LdapServerDto) GetGroupSearchBase() map[string]interface{}`

GetGroupSearchBase returns the GroupSearchBase field if non-nil, zero value otherwise.

### GetGroupSearchBaseOk

`func (o *LdapServerDto) GetGroupSearchBaseOk() (*map[string]interface{}, bool)`

GetGroupSearchBaseOk returns a tuple with the GroupSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBase

`func (o *LdapServerDto) SetGroupSearchBase(v map[string]interface{})`

SetGroupSearchBase sets GroupSearchBase field to given value.

### HasGroupSearchBase

`func (o *LdapServerDto) HasGroupSearchBase() bool`

HasGroupSearchBase returns a boolean if a field has been set.

### GetGroupSearchFilter

`func (o *LdapServerDto) GetGroupSearchFilter() map[string]interface{}`

GetGroupSearchFilter returns the GroupSearchFilter field if non-nil, zero value otherwise.

### GetGroupSearchFilterOk

`func (o *LdapServerDto) GetGroupSearchFilterOk() (*map[string]interface{}, bool)`

GetGroupSearchFilterOk returns a tuple with the GroupSearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchFilter

`func (o *LdapServerDto) SetGroupSearchFilter(v map[string]interface{})`

SetGroupSearchFilter sets GroupSearchFilter field to given value.

### HasGroupSearchFilter

`func (o *LdapServerDto) HasGroupSearchFilter() bool`

HasGroupSearchFilter returns a boolean if a field has been set.

### GetBindDN

`func (o *LdapServerDto) GetBindDN() map[string]interface{}`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *LdapServerDto) GetBindDNOk() (*map[string]interface{}, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *LdapServerDto) SetBindDN(v map[string]interface{})`

SetBindDN sets BindDN field to given value.

### HasBindDN

`func (o *LdapServerDto) HasBindDN() bool`

HasBindDN returns a boolean if a field has been set.

### GetBindCredentials

`func (o *LdapServerDto) GetBindCredentials() map[string]interface{}`

GetBindCredentials returns the BindCredentials field if non-nil, zero value otherwise.

### GetBindCredentialsOk

`func (o *LdapServerDto) GetBindCredentialsOk() (*map[string]interface{}, bool)`

GetBindCredentialsOk returns a tuple with the BindCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCredentials

`func (o *LdapServerDto) SetBindCredentials(v map[string]interface{})`

SetBindCredentials sets BindCredentials field to given value.

### HasBindCredentials

`func (o *LdapServerDto) HasBindCredentials() bool`

HasBindCredentials returns a boolean if a field has been set.

### GetCertificate

`func (o *LdapServerDto) GetCertificate() map[string]interface{}`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LdapServerDto) GetCertificateOk() (*map[string]interface{}, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LdapServerDto) SetCertificate(v map[string]interface{})`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LdapServerDto) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



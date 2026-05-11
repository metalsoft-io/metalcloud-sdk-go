# AuthenticationProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | True if the provider is enabled | 
**Domains** | Pointer to **[]string** | Permitted email domains (valid email format required) | [optional] 
**Subtrees** | Pointer to **[]string** | Non-email domain identifiers for LDAP subtrees (e.g. \&quot;adm\&quot;, \&quot;people.group\&quot;). LDAP provider only. | [optional] 
**StoreMail** | Pointer to **bool** | When true (default), the mail attribute from the LDAP response is stored as the user email. When false, the username is always used instead. LDAP provider only. | [optional] 

## Methods

### NewAuthenticationProviderUpdate

`func NewAuthenticationProviderUpdate(enabled bool, ) *AuthenticationProviderUpdate`

NewAuthenticationProviderUpdate instantiates a new AuthenticationProviderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderUpdateWithDefaults

`func NewAuthenticationProviderUpdateWithDefaults() *AuthenticationProviderUpdate`

NewAuthenticationProviderUpdateWithDefaults instantiates a new AuthenticationProviderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AuthenticationProviderUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthenticationProviderUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthenticationProviderUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDomains

`func (o *AuthenticationProviderUpdate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AuthenticationProviderUpdate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AuthenticationProviderUpdate) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AuthenticationProviderUpdate) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetSubtrees

`func (o *AuthenticationProviderUpdate) GetSubtrees() []string`

GetSubtrees returns the Subtrees field if non-nil, zero value otherwise.

### GetSubtreesOk

`func (o *AuthenticationProviderUpdate) GetSubtreesOk() (*[]string, bool)`

GetSubtreesOk returns a tuple with the Subtrees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtrees

`func (o *AuthenticationProviderUpdate) SetSubtrees(v []string)`

SetSubtrees sets Subtrees field to given value.

### HasSubtrees

`func (o *AuthenticationProviderUpdate) HasSubtrees() bool`

HasSubtrees returns a boolean if a field has been set.

### GetStoreMail

`func (o *AuthenticationProviderUpdate) GetStoreMail() bool`

GetStoreMail returns the StoreMail field if non-nil, zero value otherwise.

### GetStoreMailOk

`func (o *AuthenticationProviderUpdate) GetStoreMailOk() (*bool, bool)`

GetStoreMailOk returns a tuple with the StoreMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreMail

`func (o *AuthenticationProviderUpdate) SetStoreMail(v bool)`

SetStoreMail sets StoreMail field to given value.

### HasStoreMail

`func (o *AuthenticationProviderUpdate) HasStoreMail() bool`

HasStoreMail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



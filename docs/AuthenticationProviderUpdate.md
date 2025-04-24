# AuthenticationProviderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | True if the provider is enabled | 
**Domains** | Pointer to **[]string** | Permitted domains | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AuthenticationProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Authentication provider name | 
**Enabled** | **bool** | True if the provider is enabled | 
**Domains** | **[]string** | Permitted domains | 

## Methods

### NewAuthenticationProvider

`func NewAuthenticationProvider(name string, enabled bool, domains []string, ) *AuthenticationProvider`

NewAuthenticationProvider instantiates a new AuthenticationProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderWithDefaults

`func NewAuthenticationProviderWithDefaults() *AuthenticationProvider`

NewAuthenticationProviderWithDefaults instantiates a new AuthenticationProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticationProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationProvider) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *AuthenticationProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthenticationProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthenticationProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDomains

`func (o *AuthenticationProvider) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AuthenticationProvider) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AuthenticationProvider) SetDomains(v []string)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



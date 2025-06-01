# AuthenticationUserProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The provider name | 
**SsoEntryPoint** | Pointer to **string** | The SSO entry point URL for SAML authentication | [optional] 

## Methods

### NewAuthenticationUserProvider

`func NewAuthenticationUserProvider(provider string, ) *AuthenticationUserProvider`

NewAuthenticationUserProvider instantiates a new AuthenticationUserProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationUserProviderWithDefaults

`func NewAuthenticationUserProviderWithDefaults() *AuthenticationUserProvider`

NewAuthenticationUserProviderWithDefaults instantiates a new AuthenticationUserProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AuthenticationUserProvider) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticationUserProvider) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticationUserProvider) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSsoEntryPoint

`func (o *AuthenticationUserProvider) GetSsoEntryPoint() string`

GetSsoEntryPoint returns the SsoEntryPoint field if non-nil, zero value otherwise.

### GetSsoEntryPointOk

`func (o *AuthenticationUserProvider) GetSsoEntryPointOk() (*string, bool)`

GetSsoEntryPointOk returns a tuple with the SsoEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEntryPoint

`func (o *AuthenticationUserProvider) SetSsoEntryPoint(v string)`

SetSsoEntryPoint sets SsoEntryPoint field to given value.

### HasSsoEntryPoint

`func (o *AuthenticationUserProvider) HasSsoEntryPoint() bool`

HasSsoEntryPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



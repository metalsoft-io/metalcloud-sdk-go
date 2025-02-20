# AuthenticationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Type of the authentication request. | [default to "mysql"]
**Properties** | [**AuthenticationRequestProperties**](AuthenticationRequestProperties.md) |  | 

## Methods

### NewAuthenticationRequest

`func NewAuthenticationRequest(provider string, properties AuthenticationRequestProperties, ) *AuthenticationRequest`

NewAuthenticationRequest instantiates a new AuthenticationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationRequestWithDefaults

`func NewAuthenticationRequestWithDefaults() *AuthenticationRequest`

NewAuthenticationRequestWithDefaults instantiates a new AuthenticationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AuthenticationRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticationRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticationRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProperties

`func (o *AuthenticationRequest) GetProperties() AuthenticationRequestProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AuthenticationRequest) GetPropertiesOk() (*AuthenticationRequestProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AuthenticationRequest) SetProperties(v AuthenticationRequestProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



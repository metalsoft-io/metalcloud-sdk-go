# AuthenticationRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Type of the authentication request. | [default to "mysql"]
**Properties** | [**AuthenticationRequestDtoProperties**](AuthenticationRequestDtoProperties.md) |  | 

## Methods

### NewAuthenticationRequestDto

`func NewAuthenticationRequestDto(provider string, properties AuthenticationRequestDtoProperties, ) *AuthenticationRequestDto`

NewAuthenticationRequestDto instantiates a new AuthenticationRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationRequestDtoWithDefaults

`func NewAuthenticationRequestDtoWithDefaults() *AuthenticationRequestDto`

NewAuthenticationRequestDtoWithDefaults instantiates a new AuthenticationRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AuthenticationRequestDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticationRequestDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticationRequestDto) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProperties

`func (o *AuthenticationRequestDto) GetProperties() AuthenticationRequestDtoProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AuthenticationRequestDto) GetPropertiesOk() (*AuthenticationRequestDtoProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AuthenticationRequestDto) SetProperties(v AuthenticationRequestDtoProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



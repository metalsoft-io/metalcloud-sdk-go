# AuthenticationRequestPropertiesSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SAMLResponse** | **string** | Base64-encoded string SAML Response from the Identity Provider | 
**RelayState** | Pointer to **string** | A value that helps maintain state between the authentication request and response. It is typically used to prevent cross-site request forgery (CSRF) attacks. | [optional] 
**Signature** | Pointer to **string** | Signature of the SAML Response, used to verify the authenticity and integrity of the response. | [optional] 

## Methods

### NewAuthenticationRequestPropertiesSaml

`func NewAuthenticationRequestPropertiesSaml(sAMLResponse string, ) *AuthenticationRequestPropertiesSaml`

NewAuthenticationRequestPropertiesSaml instantiates a new AuthenticationRequestPropertiesSaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationRequestPropertiesSamlWithDefaults

`func NewAuthenticationRequestPropertiesSamlWithDefaults() *AuthenticationRequestPropertiesSaml`

NewAuthenticationRequestPropertiesSamlWithDefaults instantiates a new AuthenticationRequestPropertiesSaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSAMLResponse

`func (o *AuthenticationRequestPropertiesSaml) GetSAMLResponse() string`

GetSAMLResponse returns the SAMLResponse field if non-nil, zero value otherwise.

### GetSAMLResponseOk

`func (o *AuthenticationRequestPropertiesSaml) GetSAMLResponseOk() (*string, bool)`

GetSAMLResponseOk returns a tuple with the SAMLResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLResponse

`func (o *AuthenticationRequestPropertiesSaml) SetSAMLResponse(v string)`

SetSAMLResponse sets SAMLResponse field to given value.


### GetRelayState

`func (o *AuthenticationRequestPropertiesSaml) GetRelayState() string`

GetRelayState returns the RelayState field if non-nil, zero value otherwise.

### GetRelayStateOk

`func (o *AuthenticationRequestPropertiesSaml) GetRelayStateOk() (*string, bool)`

GetRelayStateOk returns a tuple with the RelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayState

`func (o *AuthenticationRequestPropertiesSaml) SetRelayState(v string)`

SetRelayState sets RelayState field to given value.

### HasRelayState

`func (o *AuthenticationRequestPropertiesSaml) HasRelayState() bool`

HasRelayState returns a boolean if a field has been set.

### GetSignature

`func (o *AuthenticationRequestPropertiesSaml) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AuthenticationRequestPropertiesSaml) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AuthenticationRequestPropertiesSaml) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AuthenticationRequestPropertiesSaml) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



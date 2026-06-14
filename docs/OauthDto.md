# OauthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwksUri** | **map[string]interface{}** | JWKS endpoint URL used for verifying OAuth/OIDC token signatures | 

## Methods

### NewOauthDto

`func NewOauthDto(jwksUri map[string]interface{}, ) *OauthDto`

NewOauthDto instantiates a new OauthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthDtoWithDefaults

`func NewOauthDtoWithDefaults() *OauthDto`

NewOauthDtoWithDefaults instantiates a new OauthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwksUri

`func (o *OauthDto) GetJwksUri() map[string]interface{}`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OauthDto) GetJwksUriOk() (*map[string]interface{}, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OauthDto) SetJwksUri(v map[string]interface{})`

SetJwksUri sets JwksUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



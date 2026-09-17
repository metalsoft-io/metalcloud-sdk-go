# UserApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **NullableString** | The user API key; null when the configured visibility is &#x60;once&#x60;, which withholds it after creation | [optional] 
**ApiKeyExpirationTimestamp** | Pointer to **NullableString** | ISO-8601 timestamp (UTC) when this API key expires; null when the key never expires. | [optional] 
**ApiKeyVisibility** | [**ApiKeyVisibility**](ApiKeyVisibility.md) | The effective visibility policy for this API key | 

## Methods

### NewUserApiKey

`func NewUserApiKey(apiKeyVisibility ApiKeyVisibility, ) *UserApiKey`

NewUserApiKey instantiates a new UserApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApiKeyWithDefaults

`func NewUserApiKeyWithDefaults() *UserApiKey`

NewUserApiKeyWithDefaults instantiates a new UserApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *UserApiKey) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserApiKey) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserApiKey) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UserApiKey) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *UserApiKey) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *UserApiKey) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetApiKeyExpirationTimestamp

`func (o *UserApiKey) GetApiKeyExpirationTimestamp() string`

GetApiKeyExpirationTimestamp returns the ApiKeyExpirationTimestamp field if non-nil, zero value otherwise.

### GetApiKeyExpirationTimestampOk

`func (o *UserApiKey) GetApiKeyExpirationTimestampOk() (*string, bool)`

GetApiKeyExpirationTimestampOk returns a tuple with the ApiKeyExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyExpirationTimestamp

`func (o *UserApiKey) SetApiKeyExpirationTimestamp(v string)`

SetApiKeyExpirationTimestamp sets ApiKeyExpirationTimestamp field to given value.

### HasApiKeyExpirationTimestamp

`func (o *UserApiKey) HasApiKeyExpirationTimestamp() bool`

HasApiKeyExpirationTimestamp returns a boolean if a field has been set.

### SetApiKeyExpirationTimestampNil

`func (o *UserApiKey) SetApiKeyExpirationTimestampNil(b bool)`

 SetApiKeyExpirationTimestampNil sets the value for ApiKeyExpirationTimestamp to be an explicit nil

### UnsetApiKeyExpirationTimestamp
`func (o *UserApiKey) UnsetApiKeyExpirationTimestamp()`

UnsetApiKeyExpirationTimestamp ensures that no value is present for ApiKeyExpirationTimestamp, not even an explicit nil
### GetApiKeyVisibility

`func (o *UserApiKey) GetApiKeyVisibility() ApiKeyVisibility`

GetApiKeyVisibility returns the ApiKeyVisibility field if non-nil, zero value otherwise.

### GetApiKeyVisibilityOk

`func (o *UserApiKey) GetApiKeyVisibilityOk() (*ApiKeyVisibility, bool)`

GetApiKeyVisibilityOk returns a tuple with the ApiKeyVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyVisibility

`func (o *UserApiKey) SetApiKeyVisibility(v ApiKeyVisibility)`

SetApiKeyVisibility sets ApiKeyVisibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateDeviceAuthProviderSharedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | **string** | The new shared secret used to encrypt communication with the TACACS+ server. | 

## Methods

### NewUpdateDeviceAuthProviderSharedSecret

`func NewUpdateDeviceAuthProviderSharedSecret(sharedSecret string, ) *UpdateDeviceAuthProviderSharedSecret`

NewUpdateDeviceAuthProviderSharedSecret instantiates a new UpdateDeviceAuthProviderSharedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceAuthProviderSharedSecretWithDefaults

`func NewUpdateDeviceAuthProviderSharedSecretWithDefaults() *UpdateDeviceAuthProviderSharedSecret`

NewUpdateDeviceAuthProviderSharedSecretWithDefaults instantiates a new UpdateDeviceAuthProviderSharedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *UpdateDeviceAuthProviderSharedSecret) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *UpdateDeviceAuthProviderSharedSecret) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *UpdateDeviceAuthProviderSharedSecret) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeviceAuthProviderCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username used to authenticate against the provider. | 
**Password** | Pointer to **string** | The password used to authenticate against the provider. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret used to encrypt communication with the TACACS+ server. | [optional] 

## Methods

### NewDeviceAuthProviderCredentials

`func NewDeviceAuthProviderCredentials(username string, ) *DeviceAuthProviderCredentials`

NewDeviceAuthProviderCredentials instantiates a new DeviceAuthProviderCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthProviderCredentialsWithDefaults

`func NewDeviceAuthProviderCredentialsWithDefaults() *DeviceAuthProviderCredentials`

NewDeviceAuthProviderCredentialsWithDefaults instantiates a new DeviceAuthProviderCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *DeviceAuthProviderCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeviceAuthProviderCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeviceAuthProviderCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *DeviceAuthProviderCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DeviceAuthProviderCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DeviceAuthProviderCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DeviceAuthProviderCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSharedSecret

`func (o *DeviceAuthProviderCredentials) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *DeviceAuthProviderCredentials) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *DeviceAuthProviderCredentials) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *DeviceAuthProviderCredentials) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



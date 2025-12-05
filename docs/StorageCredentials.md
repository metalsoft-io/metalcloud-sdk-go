# StorageCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username to use. | [optional] 
**Password** | Pointer to **string** | The password to use. | [optional] 
**ClientId** | Pointer to **string** | The client ID to use (for certain storage drivers) | [optional] 
**KeyId** | Pointer to **string** | The key ID to use (for certain storage drivers) | [optional] 
**Issuer** | Pointer to **string** | The application issuer to use (for certain storage drivers) | [optional] 
**PrivateKey** | Pointer to **string** | The private key to use (for certain storage drivers) | [optional] 

## Methods

### NewStorageCredentials

`func NewStorageCredentials() *StorageCredentials`

NewStorageCredentials instantiates a new StorageCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageCredentialsWithDefaults

`func NewStorageCredentialsWithDefaults() *StorageCredentials`

NewStorageCredentialsWithDefaults instantiates a new StorageCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *StorageCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *StorageCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *StorageCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetClientId

`func (o *StorageCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StorageCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StorageCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *StorageCredentials) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetKeyId

`func (o *StorageCredentials) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *StorageCredentials) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *StorageCredentials) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *StorageCredentials) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetIssuer

`func (o *StorageCredentials) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *StorageCredentials) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *StorageCredentials) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *StorageCredentials) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetPrivateKey

`func (o *StorageCredentials) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *StorageCredentials) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *StorageCredentials) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *StorageCredentials) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



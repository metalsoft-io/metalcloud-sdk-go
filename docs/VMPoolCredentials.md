# VMPoolCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Certificate of the VM Pool | [optional] 
**PrivateKey** | Pointer to **string** | Private key of the VM Pool | [optional] 
**Username** | Pointer to **string** | Username of the VM Pool | [optional] 
**Password** | Pointer to **string** | Password of the VM Pool | [optional] 

## Methods

### NewVMPoolCredentials

`func NewVMPoolCredentials() *VMPoolCredentials`

NewVMPoolCredentials instantiates a new VMPoolCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolCredentialsWithDefaults

`func NewVMPoolCredentialsWithDefaults() *VMPoolCredentials`

NewVMPoolCredentialsWithDefaults instantiates a new VMPoolCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *VMPoolCredentials) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *VMPoolCredentials) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *VMPoolCredentials) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *VMPoolCredentials) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VMPoolCredentials) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VMPoolCredentials) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VMPoolCredentials) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VMPoolCredentials) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetUsername

`func (o *VMPoolCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VMPoolCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VMPoolCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VMPoolCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *VMPoolCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VMPoolCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VMPoolCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VMPoolCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



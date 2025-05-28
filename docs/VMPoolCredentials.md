# VMPoolCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Certificate of the VM Pool | 
**PrivateKey** | **string** | Private key of the VM Pool | 

## Methods

### NewVMPoolCredentials

`func NewVMPoolCredentials(certificate string, privateKey string, ) *VMPoolCredentials`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



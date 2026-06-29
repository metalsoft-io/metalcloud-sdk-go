# LicenseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The hostname the license is bound to | 
**Valid** | **bool** | Whether the license is currently valid | 
**Expiration** | **string** | The license expiration date | 
**Signature** | **string** | The cryptographic signature of the license status | 

## Methods

### NewLicenseStatus

`func NewLicenseStatus(hostname string, valid bool, expiration string, signature string, ) *LicenseStatus`

NewLicenseStatus instantiates a new LicenseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseStatusWithDefaults

`func NewLicenseStatusWithDefaults() *LicenseStatus`

NewLicenseStatusWithDefaults instantiates a new LicenseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *LicenseStatus) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LicenseStatus) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LicenseStatus) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetValid

`func (o *LicenseStatus) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *LicenseStatus) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *LicenseStatus) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetExpiration

`func (o *LicenseStatus) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LicenseStatus) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LicenseStatus) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetSignature

`func (o *LicenseStatus) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *LicenseStatus) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *LicenseStatus) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



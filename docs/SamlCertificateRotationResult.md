# SamlCertificateRotationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpCert** | **string** | The newly generated service provider certificate in PEM format. Hand this to the identity provider administrator; the private key never leaves the deployment. | 
**CommonName** | **string** | Common name the certificate was issued for, taken from the deployment host | 
**ValidFrom** | **string** | Not-before date of the new certificate, in ISO 8601 format | 
**ValidTo** | **string** | Not-after date of the new certificate, in ISO 8601 format | 

## Methods

### NewSamlCertificateRotationResult

`func NewSamlCertificateRotationResult(spCert string, commonName string, validFrom string, validTo string, ) *SamlCertificateRotationResult`

NewSamlCertificateRotationResult instantiates a new SamlCertificateRotationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlCertificateRotationResultWithDefaults

`func NewSamlCertificateRotationResultWithDefaults() *SamlCertificateRotationResult`

NewSamlCertificateRotationResultWithDefaults instantiates a new SamlCertificateRotationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpCert

`func (o *SamlCertificateRotationResult) GetSpCert() string`

GetSpCert returns the SpCert field if non-nil, zero value otherwise.

### GetSpCertOk

`func (o *SamlCertificateRotationResult) GetSpCertOk() (*string, bool)`

GetSpCertOk returns a tuple with the SpCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCert

`func (o *SamlCertificateRotationResult) SetSpCert(v string)`

SetSpCert sets SpCert field to given value.


### GetCommonName

`func (o *SamlCertificateRotationResult) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *SamlCertificateRotationResult) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *SamlCertificateRotationResult) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetValidFrom

`func (o *SamlCertificateRotationResult) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SamlCertificateRotationResult) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *SamlCertificateRotationResult) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *SamlCertificateRotationResult) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SamlCertificateRotationResult) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *SamlCertificateRotationResult) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



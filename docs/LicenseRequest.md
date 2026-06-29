# LicenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseRequest** | **string** | The Base64-encoded license request document. It carries the exact original signed bytes returned by the license service (no re-serialization in the gateway). | 

## Methods

### NewLicenseRequest

`func NewLicenseRequest(licenseRequest string, ) *LicenseRequest`

NewLicenseRequest instantiates a new LicenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseRequestWithDefaults

`func NewLicenseRequestWithDefaults() *LicenseRequest`

NewLicenseRequestWithDefaults instantiates a new LicenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseRequest

`func (o *LicenseRequest) GetLicenseRequest() string`

GetLicenseRequest returns the LicenseRequest field if non-nil, zero value otherwise.

### GetLicenseRequestOk

`func (o *LicenseRequest) GetLicenseRequestOk() (*string, bool)`

GetLicenseRequestOk returns a tuple with the LicenseRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRequest

`func (o *LicenseRequest) SetLicenseRequest(v string)`

SetLicenseRequest sets LicenseRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



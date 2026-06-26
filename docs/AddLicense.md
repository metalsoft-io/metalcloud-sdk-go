# AddLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | **string** | The Base64-encoded signed license document. It is decoded and forwarded verbatim so the original signed bytes are preserved (no re-serialization in the gateway). | 

## Methods

### NewAddLicense

`func NewAddLicense(license string, ) *AddLicense`

NewAddLicense instantiates a new AddLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLicenseWithDefaults

`func NewAddLicenseWithDefaults() *AddLicense`

NewAddLicenseWithDefaults instantiates a new AddLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *AddLicense) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *AddLicense) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *AddLicense) SetLicense(v string)`

SetLicense sets License field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



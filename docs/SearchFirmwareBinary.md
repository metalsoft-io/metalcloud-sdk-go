# SearchFirmwareBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | [**FirmwareVendorType**](FirmwareVendorType.md) |  | 
**BaselineFilter** | [**BaselineFilter**](BaselineFilter.md) |  | 
**ServerComponentFilter** | Pointer to [**SearchFirmwareBinaryServerComponentFilter**](SearchFirmwareBinaryServerComponentFilter.md) |  | [optional] 

## Methods

### NewSearchFirmwareBinary

`func NewSearchFirmwareBinary(vendor FirmwareVendorType, baselineFilter BaselineFilter, ) *SearchFirmwareBinary`

NewSearchFirmwareBinary instantiates a new SearchFirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFirmwareBinaryWithDefaults

`func NewSearchFirmwareBinaryWithDefaults() *SearchFirmwareBinary`

NewSearchFirmwareBinaryWithDefaults instantiates a new SearchFirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *SearchFirmwareBinary) GetVendor() FirmwareVendorType`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SearchFirmwareBinary) GetVendorOk() (*FirmwareVendorType, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SearchFirmwareBinary) SetVendor(v FirmwareVendorType)`

SetVendor sets Vendor field to given value.


### GetBaselineFilter

`func (o *SearchFirmwareBinary) GetBaselineFilter() BaselineFilter`

GetBaselineFilter returns the BaselineFilter field if non-nil, zero value otherwise.

### GetBaselineFilterOk

`func (o *SearchFirmwareBinary) GetBaselineFilterOk() (*BaselineFilter, bool)`

GetBaselineFilterOk returns a tuple with the BaselineFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineFilter

`func (o *SearchFirmwareBinary) SetBaselineFilter(v BaselineFilter)`

SetBaselineFilter sets BaselineFilter field to given value.


### GetServerComponentFilter

`func (o *SearchFirmwareBinary) GetServerComponentFilter() SearchFirmwareBinaryServerComponentFilter`

GetServerComponentFilter returns the ServerComponentFilter field if non-nil, zero value otherwise.

### GetServerComponentFilterOk

`func (o *SearchFirmwareBinary) GetServerComponentFilterOk() (*SearchFirmwareBinaryServerComponentFilter, bool)`

GetServerComponentFilterOk returns a tuple with the ServerComponentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComponentFilter

`func (o *SearchFirmwareBinary) SetServerComponentFilter(v SearchFirmwareBinaryServerComponentFilter)`

SetServerComponentFilter sets ServerComponentFilter field to given value.

### HasServerComponentFilter

`func (o *SearchFirmwareBinary) HasServerComponentFilter() bool`

HasServerComponentFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



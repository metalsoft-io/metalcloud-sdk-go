# SearchFirmwareBinaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | [**FirmwareVendorType**](FirmwareVendorType.md) |  | 
**BaselineFilter** | [**BaselineFilter**](BaselineFilter.md) |  | 
**ServerComponentFilter** | Pointer to [**[]ServerComponentFilterInner**](ServerComponentFilterInner.md) | Array of vendor component filters, structure is specific per vendor | [optional] 

## Methods

### NewSearchFirmwareBinaryDto

`func NewSearchFirmwareBinaryDto(vendor FirmwareVendorType, baselineFilter BaselineFilter, ) *SearchFirmwareBinaryDto`

NewSearchFirmwareBinaryDto instantiates a new SearchFirmwareBinaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFirmwareBinaryDtoWithDefaults

`func NewSearchFirmwareBinaryDtoWithDefaults() *SearchFirmwareBinaryDto`

NewSearchFirmwareBinaryDtoWithDefaults instantiates a new SearchFirmwareBinaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *SearchFirmwareBinaryDto) GetVendor() FirmwareVendorType`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SearchFirmwareBinaryDto) GetVendorOk() (*FirmwareVendorType, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SearchFirmwareBinaryDto) SetVendor(v FirmwareVendorType)`

SetVendor sets Vendor field to given value.


### GetBaselineFilter

`func (o *SearchFirmwareBinaryDto) GetBaselineFilter() BaselineFilter`

GetBaselineFilter returns the BaselineFilter field if non-nil, zero value otherwise.

### GetBaselineFilterOk

`func (o *SearchFirmwareBinaryDto) GetBaselineFilterOk() (*BaselineFilter, bool)`

GetBaselineFilterOk returns a tuple with the BaselineFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineFilter

`func (o *SearchFirmwareBinaryDto) SetBaselineFilter(v BaselineFilter)`

SetBaselineFilter sets BaselineFilter field to given value.


### GetServerComponentFilter

`func (o *SearchFirmwareBinaryDto) GetServerComponentFilter() []ServerComponentFilterInner`

GetServerComponentFilter returns the ServerComponentFilter field if non-nil, zero value otherwise.

### GetServerComponentFilterOk

`func (o *SearchFirmwareBinaryDto) GetServerComponentFilterOk() (*[]ServerComponentFilterInner, bool)`

GetServerComponentFilterOk returns a tuple with the ServerComponentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComponentFilter

`func (o *SearchFirmwareBinaryDto) SetServerComponentFilter(v []ServerComponentFilterInner)`

SetServerComponentFilter sets ServerComponentFilter field to given value.

### HasServerComponentFilter

`func (o *SearchFirmwareBinaryDto) HasServerComponentFilter() bool`

HasServerComponentFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



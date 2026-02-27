# SearchFirmwareBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | [**ServerFirmwareCatalogVendor**](ServerFirmwareCatalogVendor.md) |  | 
**BaselineIds** | Pointer to **[]string** |  | [optional] 
**ServerComponentFilter** | Pointer to [**SearchFirmwareBinaryServerComponentFilter**](SearchFirmwareBinaryServerComponentFilter.md) |  | [optional] 

## Methods

### NewSearchFirmwareBinary

`func NewSearchFirmwareBinary(vendor ServerFirmwareCatalogVendor, ) *SearchFirmwareBinary`

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

`func (o *SearchFirmwareBinary) GetVendor() ServerFirmwareCatalogVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SearchFirmwareBinary) GetVendorOk() (*ServerFirmwareCatalogVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SearchFirmwareBinary) SetVendor(v ServerFirmwareCatalogVendor)`

SetVendor sets Vendor field to given value.


### GetBaselineIds

`func (o *SearchFirmwareBinary) GetBaselineIds() []string`

GetBaselineIds returns the BaselineIds field if non-nil, zero value otherwise.

### GetBaselineIdsOk

`func (o *SearchFirmwareBinary) GetBaselineIdsOk() (*[]string, bool)`

GetBaselineIdsOk returns a tuple with the BaselineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineIds

`func (o *SearchFirmwareBinary) SetBaselineIds(v []string)`

SetBaselineIds sets BaselineIds field to given value.

### HasBaselineIds

`func (o *SearchFirmwareBinary) HasBaselineIds() bool`

HasBaselineIds returns a boolean if a field has been set.

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



# UpdateFirmwareCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Vendor** | [**ServerFirmwareCatalogVendor**](ServerFirmwareCatalogVendor.md) |  | 
**UpdateType** | [**CatalogUpdateType**](CatalogUpdateType.md) |  | 
**VendorId** | Pointer to **string** |  | [optional] 
**VendorUrl** | Pointer to **string** |  | [optional] 
**VendorReleaseTimestamp** | Pointer to **string** |  | [optional] 
**MetalsoftServerTypesSupported** | Pointer to **[]string** |  | [optional] 
**VendorServerTypesSupported** | Pointer to **[]string** | Array of the server types supported by the vendor for this catalog | [optional] 
**VendorConfiguration** | Pointer to **map[string]interface{}** | Record of the vendor configuration for this catalog | [optional] 

## Methods

### NewUpdateFirmwareCatalog

`func NewUpdateFirmwareCatalog(name string, vendor ServerFirmwareCatalogVendor, updateType CatalogUpdateType, ) *UpdateFirmwareCatalog`

NewUpdateFirmwareCatalog instantiates a new UpdateFirmwareCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareCatalogWithDefaults

`func NewUpdateFirmwareCatalogWithDefaults() *UpdateFirmwareCatalog`

NewUpdateFirmwareCatalogWithDefaults instantiates a new UpdateFirmwareCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFirmwareCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFirmwareCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFirmwareCatalog) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateFirmwareCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFirmwareCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFirmwareCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFirmwareCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVendor

`func (o *UpdateFirmwareCatalog) GetVendor() ServerFirmwareCatalogVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *UpdateFirmwareCatalog) GetVendorOk() (*ServerFirmwareCatalogVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *UpdateFirmwareCatalog) SetVendor(v ServerFirmwareCatalogVendor)`

SetVendor sets Vendor field to given value.


### GetUpdateType

`func (o *UpdateFirmwareCatalog) GetUpdateType() CatalogUpdateType`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *UpdateFirmwareCatalog) GetUpdateTypeOk() (*CatalogUpdateType, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *UpdateFirmwareCatalog) SetUpdateType(v CatalogUpdateType)`

SetUpdateType sets UpdateType field to given value.


### GetVendorId

`func (o *UpdateFirmwareCatalog) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *UpdateFirmwareCatalog) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *UpdateFirmwareCatalog) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *UpdateFirmwareCatalog) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorUrl

`func (o *UpdateFirmwareCatalog) GetVendorUrl() string`

GetVendorUrl returns the VendorUrl field if non-nil, zero value otherwise.

### GetVendorUrlOk

`func (o *UpdateFirmwareCatalog) GetVendorUrlOk() (*string, bool)`

GetVendorUrlOk returns a tuple with the VendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUrl

`func (o *UpdateFirmwareCatalog) SetVendorUrl(v string)`

SetVendorUrl sets VendorUrl field to given value.

### HasVendorUrl

`func (o *UpdateFirmwareCatalog) HasVendorUrl() bool`

HasVendorUrl returns a boolean if a field has been set.

### GetVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalog) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *UpdateFirmwareCatalog) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalog) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalog) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetMetalsoftServerTypesSupported

`func (o *UpdateFirmwareCatalog) GetMetalsoftServerTypesSupported() []string`

GetMetalsoftServerTypesSupported returns the MetalsoftServerTypesSupported field if non-nil, zero value otherwise.

### GetMetalsoftServerTypesSupportedOk

`func (o *UpdateFirmwareCatalog) GetMetalsoftServerTypesSupportedOk() (*[]string, bool)`

GetMetalsoftServerTypesSupportedOk returns a tuple with the MetalsoftServerTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalsoftServerTypesSupported

`func (o *UpdateFirmwareCatalog) SetMetalsoftServerTypesSupported(v []string)`

SetMetalsoftServerTypesSupported sets MetalsoftServerTypesSupported field to given value.

### HasMetalsoftServerTypesSupported

`func (o *UpdateFirmwareCatalog) HasMetalsoftServerTypesSupported() bool`

HasMetalsoftServerTypesSupported returns a boolean if a field has been set.

### GetVendorServerTypesSupported

`func (o *UpdateFirmwareCatalog) GetVendorServerTypesSupported() []string`

GetVendorServerTypesSupported returns the VendorServerTypesSupported field if non-nil, zero value otherwise.

### GetVendorServerTypesSupportedOk

`func (o *UpdateFirmwareCatalog) GetVendorServerTypesSupportedOk() (*[]string, bool)`

GetVendorServerTypesSupportedOk returns a tuple with the VendorServerTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorServerTypesSupported

`func (o *UpdateFirmwareCatalog) SetVendorServerTypesSupported(v []string)`

SetVendorServerTypesSupported sets VendorServerTypesSupported field to given value.

### HasVendorServerTypesSupported

`func (o *UpdateFirmwareCatalog) HasVendorServerTypesSupported() bool`

HasVendorServerTypesSupported returns a boolean if a field has been set.

### GetVendorConfiguration

`func (o *UpdateFirmwareCatalog) GetVendorConfiguration() map[string]interface{}`

GetVendorConfiguration returns the VendorConfiguration field if non-nil, zero value otherwise.

### GetVendorConfigurationOk

`func (o *UpdateFirmwareCatalog) GetVendorConfigurationOk() (*map[string]interface{}, bool)`

GetVendorConfigurationOk returns a tuple with the VendorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfiguration

`func (o *UpdateFirmwareCatalog) SetVendorConfiguration(v map[string]interface{})`

SetVendorConfiguration sets VendorConfiguration field to given value.

### HasVendorConfiguration

`func (o *UpdateFirmwareCatalog) HasVendorConfiguration() bool`

HasVendorConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



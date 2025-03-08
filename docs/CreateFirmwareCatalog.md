# CreateFirmwareCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Vendor** | [**FirmwareVendorType**](FirmwareVendorType.md) |  | 
**UpdateType** | [**CatalogUpdateType**](CatalogUpdateType.md) |  | 
**VendorId** | Pointer to **string** |  | [optional] 
**VendorUrl** | Pointer to **string** |  | [optional] 
**VendorReleaseTimestamp** | Pointer to **time.Time** |  | [optional] 
**MetalsoftServerTypesSupported** | Pointer to **[]string** |  | [optional] 
**VendorServerTypesSupported** | Pointer to **[]string** | Array of the server types supported by the vendor for this catalog | [optional] 
**VendorConfiguration** | Pointer to **map[string]interface{}** | Record of the vendor configuration for this catalog | [optional] 

## Methods

### NewCreateFirmwareCatalog

`func NewCreateFirmwareCatalog(name string, vendor FirmwareVendorType, updateType CatalogUpdateType, ) *CreateFirmwareCatalog`

NewCreateFirmwareCatalog instantiates a new CreateFirmwareCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareCatalogWithDefaults

`func NewCreateFirmwareCatalogWithDefaults() *CreateFirmwareCatalog`

NewCreateFirmwareCatalogWithDefaults instantiates a new CreateFirmwareCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFirmwareCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirmwareCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirmwareCatalog) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateFirmwareCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirmwareCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirmwareCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirmwareCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVendor

`func (o *CreateFirmwareCatalog) GetVendor() FirmwareVendorType`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CreateFirmwareCatalog) GetVendorOk() (*FirmwareVendorType, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CreateFirmwareCatalog) SetVendor(v FirmwareVendorType)`

SetVendor sets Vendor field to given value.


### GetUpdateType

`func (o *CreateFirmwareCatalog) GetUpdateType() CatalogUpdateType`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *CreateFirmwareCatalog) GetUpdateTypeOk() (*CatalogUpdateType, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *CreateFirmwareCatalog) SetUpdateType(v CatalogUpdateType)`

SetUpdateType sets UpdateType field to given value.


### GetVendorId

`func (o *CreateFirmwareCatalog) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *CreateFirmwareCatalog) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *CreateFirmwareCatalog) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *CreateFirmwareCatalog) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorUrl

`func (o *CreateFirmwareCatalog) GetVendorUrl() string`

GetVendorUrl returns the VendorUrl field if non-nil, zero value otherwise.

### GetVendorUrlOk

`func (o *CreateFirmwareCatalog) GetVendorUrlOk() (*string, bool)`

GetVendorUrlOk returns a tuple with the VendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUrl

`func (o *CreateFirmwareCatalog) SetVendorUrl(v string)`

SetVendorUrl sets VendorUrl field to given value.

### HasVendorUrl

`func (o *CreateFirmwareCatalog) HasVendorUrl() bool`

HasVendorUrl returns a boolean if a field has been set.

### GetVendorReleaseTimestamp

`func (o *CreateFirmwareCatalog) GetVendorReleaseTimestamp() time.Time`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *CreateFirmwareCatalog) GetVendorReleaseTimestampOk() (*time.Time, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *CreateFirmwareCatalog) SetVendorReleaseTimestamp(v time.Time)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *CreateFirmwareCatalog) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetMetalsoftServerTypesSupported

`func (o *CreateFirmwareCatalog) GetMetalsoftServerTypesSupported() []string`

GetMetalsoftServerTypesSupported returns the MetalsoftServerTypesSupported field if non-nil, zero value otherwise.

### GetMetalsoftServerTypesSupportedOk

`func (o *CreateFirmwareCatalog) GetMetalsoftServerTypesSupportedOk() (*[]string, bool)`

GetMetalsoftServerTypesSupportedOk returns a tuple with the MetalsoftServerTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalsoftServerTypesSupported

`func (o *CreateFirmwareCatalog) SetMetalsoftServerTypesSupported(v []string)`

SetMetalsoftServerTypesSupported sets MetalsoftServerTypesSupported field to given value.

### HasMetalsoftServerTypesSupported

`func (o *CreateFirmwareCatalog) HasMetalsoftServerTypesSupported() bool`

HasMetalsoftServerTypesSupported returns a boolean if a field has been set.

### GetVendorServerTypesSupported

`func (o *CreateFirmwareCatalog) GetVendorServerTypesSupported() []string`

GetVendorServerTypesSupported returns the VendorServerTypesSupported field if non-nil, zero value otherwise.

### GetVendorServerTypesSupportedOk

`func (o *CreateFirmwareCatalog) GetVendorServerTypesSupportedOk() (*[]string, bool)`

GetVendorServerTypesSupportedOk returns a tuple with the VendorServerTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorServerTypesSupported

`func (o *CreateFirmwareCatalog) SetVendorServerTypesSupported(v []string)`

SetVendorServerTypesSupported sets VendorServerTypesSupported field to given value.

### HasVendorServerTypesSupported

`func (o *CreateFirmwareCatalog) HasVendorServerTypesSupported() bool`

HasVendorServerTypesSupported returns a boolean if a field has been set.

### GetVendorConfiguration

`func (o *CreateFirmwareCatalog) GetVendorConfiguration() map[string]interface{}`

GetVendorConfiguration returns the VendorConfiguration field if non-nil, zero value otherwise.

### GetVendorConfigurationOk

`func (o *CreateFirmwareCatalog) GetVendorConfigurationOk() (*map[string]interface{}, bool)`

GetVendorConfigurationOk returns a tuple with the VendorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfiguration

`func (o *CreateFirmwareCatalog) SetVendorConfiguration(v map[string]interface{})`

SetVendorConfiguration sets VendorConfiguration field to given value.

### HasVendorConfiguration

`func (o *CreateFirmwareCatalog) HasVendorConfiguration() bool`

HasVendorConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



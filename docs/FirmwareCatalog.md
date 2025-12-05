# FirmwareCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Unique identifier for the firmware catalog | 
**Name** | **string** | Name of the catalog, must be unique | 
**Description** | Pointer to **string** | User description of the catalog | [optional] 
**Vendor** | **string** | Firmware catalog vendor: dell, lenovo, hp, supermicro | 
**VendorId** | Pointer to **string** | Vendor identifier for the catalog | [optional] 
**VendorUrl** | Pointer to **string** | Vendor URL for the firmware catalog | [optional] 
**VendorReleaseTimestamp** | Pointer to **string** | Vendor release timestamp for the catalog | [optional] 
**UpdateType** | **string** | Catalog update type: online or offline | [default to "online"]
**MetalsoftServerTypesSupported** | Pointer to **[]string** | List of supported Metalsoft server types for which firmware binaries are available | [optional] 
**VendorServerTypesSupported** | Pointer to **[]string** | List of supported server types for which firmware binaries are available | [optional] 
**VendorConfiguration** | Pointer to **map[string]interface{}** | Vendor configuration | [optional] 
**CreatedTimestamp** | **string** | Timestamp when the catalog was created | 
**Links** | [**[]Link**](Link.md) | Links to other resources | 

## Methods

### NewFirmwareCatalog

`func NewFirmwareCatalog(id float32, name string, vendor string, updateType string, createdTimestamp string, links []Link, ) *FirmwareCatalog`

NewFirmwareCatalog instantiates a new FirmwareCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareCatalogWithDefaults

`func NewFirmwareCatalogWithDefaults() *FirmwareCatalog`

NewFirmwareCatalogWithDefaults instantiates a new FirmwareCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirmwareCatalog) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareCatalog) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareCatalog) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *FirmwareCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareCatalog) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FirmwareCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirmwareCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirmwareCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirmwareCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVendor

`func (o *FirmwareCatalog) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareCatalog) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareCatalog) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetVendorId

`func (o *FirmwareCatalog) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *FirmwareCatalog) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *FirmwareCatalog) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *FirmwareCatalog) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorUrl

`func (o *FirmwareCatalog) GetVendorUrl() string`

GetVendorUrl returns the VendorUrl field if non-nil, zero value otherwise.

### GetVendorUrlOk

`func (o *FirmwareCatalog) GetVendorUrlOk() (*string, bool)`

GetVendorUrlOk returns a tuple with the VendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUrl

`func (o *FirmwareCatalog) SetVendorUrl(v string)`

SetVendorUrl sets VendorUrl field to given value.

### HasVendorUrl

`func (o *FirmwareCatalog) HasVendorUrl() bool`

HasVendorUrl returns a boolean if a field has been set.

### GetVendorReleaseTimestamp

`func (o *FirmwareCatalog) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *FirmwareCatalog) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *FirmwareCatalog) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *FirmwareCatalog) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetUpdateType

`func (o *FirmwareCatalog) GetUpdateType() string`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *FirmwareCatalog) GetUpdateTypeOk() (*string, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *FirmwareCatalog) SetUpdateType(v string)`

SetUpdateType sets UpdateType field to given value.


### GetMetalsoftServerTypesSupported

`func (o *FirmwareCatalog) GetMetalsoftServerTypesSupported() []string`

GetMetalsoftServerTypesSupported returns the MetalsoftServerTypesSupported field if non-nil, zero value otherwise.

### GetMetalsoftServerTypesSupportedOk

`func (o *FirmwareCatalog) GetMetalsoftServerTypesSupportedOk() (*[]string, bool)`

GetMetalsoftServerTypesSupportedOk returns a tuple with the MetalsoftServerTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalsoftServerTypesSupported

`func (o *FirmwareCatalog) SetMetalsoftServerTypesSupported(v []string)`

SetMetalsoftServerTypesSupported sets MetalsoftServerTypesSupported field to given value.

### HasMetalsoftServerTypesSupported

`func (o *FirmwareCatalog) HasMetalsoftServerTypesSupported() bool`

HasMetalsoftServerTypesSupported returns a boolean if a field has been set.

### GetVendorServerTypesSupported

`func (o *FirmwareCatalog) GetVendorServerTypesSupported() []string`

GetVendorServerTypesSupported returns the VendorServerTypesSupported field if non-nil, zero value otherwise.

### GetVendorServerTypesSupportedOk

`func (o *FirmwareCatalog) GetVendorServerTypesSupportedOk() (*[]string, bool)`

GetVendorServerTypesSupportedOk returns a tuple with the VendorServerTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorServerTypesSupported

`func (o *FirmwareCatalog) SetVendorServerTypesSupported(v []string)`

SetVendorServerTypesSupported sets VendorServerTypesSupported field to given value.

### HasVendorServerTypesSupported

`func (o *FirmwareCatalog) HasVendorServerTypesSupported() bool`

HasVendorServerTypesSupported returns a boolean if a field has been set.

### GetVendorConfiguration

`func (o *FirmwareCatalog) GetVendorConfiguration() map[string]interface{}`

GetVendorConfiguration returns the VendorConfiguration field if non-nil, zero value otherwise.

### GetVendorConfigurationOk

`func (o *FirmwareCatalog) GetVendorConfigurationOk() (*map[string]interface{}, bool)`

GetVendorConfigurationOk returns a tuple with the VendorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorConfiguration

`func (o *FirmwareCatalog) SetVendorConfiguration(v map[string]interface{})`

SetVendorConfiguration sets VendorConfiguration field to given value.

### HasVendorConfiguration

`func (o *FirmwareCatalog) HasVendorConfiguration() bool`

HasVendorConfiguration returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *FirmwareCatalog) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FirmwareCatalog) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FirmwareCatalog) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLinks

`func (o *FirmwareCatalog) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirmwareCatalog) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirmwareCatalog) SetLinks(v []Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



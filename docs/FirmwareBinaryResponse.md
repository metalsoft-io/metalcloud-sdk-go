# FirmwareBinaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Unique identifier for the firmware binary | 
**CatalogId** | **float32** | Unique identifier for the firmware catalog | 
**ExternalId** | Pointer to **string** | External/vendor identifier for the firmware binary | [optional] 
**VendorInfoUrl** | Pointer to **string** | Vendor URL with information about the firmware binary | [optional] 
**VendorDownloadUrl** | **string** | Vendor download URL for the firmware binary | 
**CacheDownloadUrl** | Pointer to **string** | Local cache download URL for the firmware binary | [optional] 
**Name** | **string** | Name of the firmware binary as specified by the vendor | 
**PackageId** | Pointer to **string** | Vendor package identifier | [optional] 
**PackageVersion** | **string** | Vendor package version | 
**RebootRequired** | **bool** | Indicates if a reboot is required for this binary | 
**UpdateSeverity** | **string** | Severity of the firmware update | 
**VendorSupportedDevices** | **map[string]interface{}** | Vendor specific record for supported devices | 
**VendorSupportedSystems** | **[]map[string]interface{}** | Vendor specific records for supported systems | 
**Vendor** | **map[string]interface{}** | Vendor specific record containing all other firmware binary information | 
**VendorReleaseTimestamp** | Pointer to **string** | Vendor release timestamp for the firmware binary | [optional] 
**CreatedTimestamp** | **string** | Timestamp when the firmware binary was created | 
**Catalog** | Pointer to [**FirmwareCatalog**](FirmwareCatalog.md) | Associated firmware catalog | [optional] 

## Methods

### NewFirmwareBinaryResponse

`func NewFirmwareBinaryResponse(id float32, catalogId float32, vendorDownloadUrl string, name string, packageVersion string, rebootRequired bool, updateSeverity string, vendorSupportedDevices map[string]interface{}, vendorSupportedSystems []map[string]interface{}, vendor map[string]interface{}, createdTimestamp string, ) *FirmwareBinaryResponse`

NewFirmwareBinaryResponse instantiates a new FirmwareBinaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBinaryResponseWithDefaults

`func NewFirmwareBinaryResponseWithDefaults() *FirmwareBinaryResponse`

NewFirmwareBinaryResponseWithDefaults instantiates a new FirmwareBinaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirmwareBinaryResponse) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareBinaryResponse) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareBinaryResponse) SetId(v float32)`

SetId sets Id field to given value.


### GetCatalogId

`func (o *FirmwareBinaryResponse) GetCatalogId() float32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *FirmwareBinaryResponse) GetCatalogIdOk() (*float32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *FirmwareBinaryResponse) SetCatalogId(v float32)`

SetCatalogId sets CatalogId field to given value.


### GetExternalId

`func (o *FirmwareBinaryResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FirmwareBinaryResponse) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FirmwareBinaryResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FirmwareBinaryResponse) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVendorInfoUrl

`func (o *FirmwareBinaryResponse) GetVendorInfoUrl() string`

GetVendorInfoUrl returns the VendorInfoUrl field if non-nil, zero value otherwise.

### GetVendorInfoUrlOk

`func (o *FirmwareBinaryResponse) GetVendorInfoUrlOk() (*string, bool)`

GetVendorInfoUrlOk returns a tuple with the VendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfoUrl

`func (o *FirmwareBinaryResponse) SetVendorInfoUrl(v string)`

SetVendorInfoUrl sets VendorInfoUrl field to given value.

### HasVendorInfoUrl

`func (o *FirmwareBinaryResponse) HasVendorInfoUrl() bool`

HasVendorInfoUrl returns a boolean if a field has been set.

### GetVendorDownloadUrl

`func (o *FirmwareBinaryResponse) GetVendorDownloadUrl() string`

GetVendorDownloadUrl returns the VendorDownloadUrl field if non-nil, zero value otherwise.

### GetVendorDownloadUrlOk

`func (o *FirmwareBinaryResponse) GetVendorDownloadUrlOk() (*string, bool)`

GetVendorDownloadUrlOk returns a tuple with the VendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorDownloadUrl

`func (o *FirmwareBinaryResponse) SetVendorDownloadUrl(v string)`

SetVendorDownloadUrl sets VendorDownloadUrl field to given value.


### GetCacheDownloadUrl

`func (o *FirmwareBinaryResponse) GetCacheDownloadUrl() string`

GetCacheDownloadUrl returns the CacheDownloadUrl field if non-nil, zero value otherwise.

### GetCacheDownloadUrlOk

`func (o *FirmwareBinaryResponse) GetCacheDownloadUrlOk() (*string, bool)`

GetCacheDownloadUrlOk returns a tuple with the CacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDownloadUrl

`func (o *FirmwareBinaryResponse) SetCacheDownloadUrl(v string)`

SetCacheDownloadUrl sets CacheDownloadUrl field to given value.

### HasCacheDownloadUrl

`func (o *FirmwareBinaryResponse) HasCacheDownloadUrl() bool`

HasCacheDownloadUrl returns a boolean if a field has been set.

### GetName

`func (o *FirmwareBinaryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareBinaryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareBinaryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *FirmwareBinaryResponse) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *FirmwareBinaryResponse) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *FirmwareBinaryResponse) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *FirmwareBinaryResponse) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPackageVersion

`func (o *FirmwareBinaryResponse) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *FirmwareBinaryResponse) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *FirmwareBinaryResponse) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetRebootRequired

`func (o *FirmwareBinaryResponse) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *FirmwareBinaryResponse) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *FirmwareBinaryResponse) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetUpdateSeverity

`func (o *FirmwareBinaryResponse) GetUpdateSeverity() string`

GetUpdateSeverity returns the UpdateSeverity field if non-nil, zero value otherwise.

### GetUpdateSeverityOk

`func (o *FirmwareBinaryResponse) GetUpdateSeverityOk() (*string, bool)`

GetUpdateSeverityOk returns a tuple with the UpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeverity

`func (o *FirmwareBinaryResponse) SetUpdateSeverity(v string)`

SetUpdateSeverity sets UpdateSeverity field to given value.


### GetVendorSupportedDevices

`func (o *FirmwareBinaryResponse) GetVendorSupportedDevices() map[string]interface{}`

GetVendorSupportedDevices returns the VendorSupportedDevices field if non-nil, zero value otherwise.

### GetVendorSupportedDevicesOk

`func (o *FirmwareBinaryResponse) GetVendorSupportedDevicesOk() (*map[string]interface{}, bool)`

GetVendorSupportedDevicesOk returns a tuple with the VendorSupportedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedDevices

`func (o *FirmwareBinaryResponse) SetVendorSupportedDevices(v map[string]interface{})`

SetVendorSupportedDevices sets VendorSupportedDevices field to given value.


### GetVendorSupportedSystems

`func (o *FirmwareBinaryResponse) GetVendorSupportedSystems() []map[string]interface{}`

GetVendorSupportedSystems returns the VendorSupportedSystems field if non-nil, zero value otherwise.

### GetVendorSupportedSystemsOk

`func (o *FirmwareBinaryResponse) GetVendorSupportedSystemsOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedSystemsOk returns a tuple with the VendorSupportedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedSystems

`func (o *FirmwareBinaryResponse) SetVendorSupportedSystems(v []map[string]interface{})`

SetVendorSupportedSystems sets VendorSupportedSystems field to given value.


### GetVendor

`func (o *FirmwareBinaryResponse) GetVendor() map[string]interface{}`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareBinaryResponse) GetVendorOk() (*map[string]interface{}, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareBinaryResponse) SetVendor(v map[string]interface{})`

SetVendor sets Vendor field to given value.


### GetVendorReleaseTimestamp

`func (o *FirmwareBinaryResponse) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *FirmwareBinaryResponse) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *FirmwareBinaryResponse) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *FirmwareBinaryResponse) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *FirmwareBinaryResponse) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FirmwareBinaryResponse) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FirmwareBinaryResponse) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetCatalog

`func (o *FirmwareBinaryResponse) GetCatalog() FirmwareCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareBinaryResponse) GetCatalogOk() (*FirmwareCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareBinaryResponse) SetCatalog(v FirmwareCatalog)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareBinaryResponse) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



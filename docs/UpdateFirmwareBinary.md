# UpdateFirmwareBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogId** | **float32** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**VendorInfoUrl** | Pointer to **string** |  | [optional] 
**VendorDownloadUrl** | **string** |  | 
**CacheDownloadUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**PackageId** | Pointer to **string** |  | [optional] 
**PackageVersion** | Pointer to **string** |  | [optional] 
**RebootRequired** | **bool** |  | 
**UpdateSeverity** | [**FirmwareBinaryUpdateSeverity**](FirmwareBinaryUpdateSeverity.md) |  | 
**VendorSupportedDevices** | **[]map[string]interface{}** |  | 
**VendorSupportedSystems** | **[]map[string]interface{}** |  | 
**VendorReleaseTimestamp** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateFirmwareBinary

`func NewUpdateFirmwareBinary(catalogId float32, vendorDownloadUrl string, name string, rebootRequired bool, updateSeverity FirmwareBinaryUpdateSeverity, vendorSupportedDevices []map[string]interface{}, vendorSupportedSystems []map[string]interface{}, ) *UpdateFirmwareBinary`

NewUpdateFirmwareBinary instantiates a new UpdateFirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareBinaryWithDefaults

`func NewUpdateFirmwareBinaryWithDefaults() *UpdateFirmwareBinary`

NewUpdateFirmwareBinaryWithDefaults instantiates a new UpdateFirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *UpdateFirmwareBinary) GetCatalogId() float32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *UpdateFirmwareBinary) GetCatalogIdOk() (*float32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *UpdateFirmwareBinary) SetCatalogId(v float32)`

SetCatalogId sets CatalogId field to given value.


### GetExternalId

`func (o *UpdateFirmwareBinary) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateFirmwareBinary) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateFirmwareBinary) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateFirmwareBinary) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVendorInfoUrl

`func (o *UpdateFirmwareBinary) GetVendorInfoUrl() string`

GetVendorInfoUrl returns the VendorInfoUrl field if non-nil, zero value otherwise.

### GetVendorInfoUrlOk

`func (o *UpdateFirmwareBinary) GetVendorInfoUrlOk() (*string, bool)`

GetVendorInfoUrlOk returns a tuple with the VendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfoUrl

`func (o *UpdateFirmwareBinary) SetVendorInfoUrl(v string)`

SetVendorInfoUrl sets VendorInfoUrl field to given value.

### HasVendorInfoUrl

`func (o *UpdateFirmwareBinary) HasVendorInfoUrl() bool`

HasVendorInfoUrl returns a boolean if a field has been set.

### GetVendorDownloadUrl

`func (o *UpdateFirmwareBinary) GetVendorDownloadUrl() string`

GetVendorDownloadUrl returns the VendorDownloadUrl field if non-nil, zero value otherwise.

### GetVendorDownloadUrlOk

`func (o *UpdateFirmwareBinary) GetVendorDownloadUrlOk() (*string, bool)`

GetVendorDownloadUrlOk returns a tuple with the VendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorDownloadUrl

`func (o *UpdateFirmwareBinary) SetVendorDownloadUrl(v string)`

SetVendorDownloadUrl sets VendorDownloadUrl field to given value.


### GetCacheDownloadUrl

`func (o *UpdateFirmwareBinary) GetCacheDownloadUrl() string`

GetCacheDownloadUrl returns the CacheDownloadUrl field if non-nil, zero value otherwise.

### GetCacheDownloadUrlOk

`func (o *UpdateFirmwareBinary) GetCacheDownloadUrlOk() (*string, bool)`

GetCacheDownloadUrlOk returns a tuple with the CacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDownloadUrl

`func (o *UpdateFirmwareBinary) SetCacheDownloadUrl(v string)`

SetCacheDownloadUrl sets CacheDownloadUrl field to given value.

### HasCacheDownloadUrl

`func (o *UpdateFirmwareBinary) HasCacheDownloadUrl() bool`

HasCacheDownloadUrl returns a boolean if a field has been set.

### GetName

`func (o *UpdateFirmwareBinary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFirmwareBinary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFirmwareBinary) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *UpdateFirmwareBinary) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *UpdateFirmwareBinary) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *UpdateFirmwareBinary) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *UpdateFirmwareBinary) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPackageVersion

`func (o *UpdateFirmwareBinary) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *UpdateFirmwareBinary) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *UpdateFirmwareBinary) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *UpdateFirmwareBinary) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetRebootRequired

`func (o *UpdateFirmwareBinary) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *UpdateFirmwareBinary) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *UpdateFirmwareBinary) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetUpdateSeverity

`func (o *UpdateFirmwareBinary) GetUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetUpdateSeverity returns the UpdateSeverity field if non-nil, zero value otherwise.

### GetUpdateSeverityOk

`func (o *UpdateFirmwareBinary) GetUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetUpdateSeverityOk returns a tuple with the UpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeverity

`func (o *UpdateFirmwareBinary) SetUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetUpdateSeverity sets UpdateSeverity field to given value.


### GetVendorSupportedDevices

`func (o *UpdateFirmwareBinary) GetVendorSupportedDevices() []map[string]interface{}`

GetVendorSupportedDevices returns the VendorSupportedDevices field if non-nil, zero value otherwise.

### GetVendorSupportedDevicesOk

`func (o *UpdateFirmwareBinary) GetVendorSupportedDevicesOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedDevicesOk returns a tuple with the VendorSupportedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedDevices

`func (o *UpdateFirmwareBinary) SetVendorSupportedDevices(v []map[string]interface{})`

SetVendorSupportedDevices sets VendorSupportedDevices field to given value.


### GetVendorSupportedSystems

`func (o *UpdateFirmwareBinary) GetVendorSupportedSystems() []map[string]interface{}`

GetVendorSupportedSystems returns the VendorSupportedSystems field if non-nil, zero value otherwise.

### GetVendorSupportedSystemsOk

`func (o *UpdateFirmwareBinary) GetVendorSupportedSystemsOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedSystemsOk returns a tuple with the VendorSupportedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedSystems

`func (o *UpdateFirmwareBinary) SetVendorSupportedSystems(v []map[string]interface{})`

SetVendorSupportedSystems sets VendorSupportedSystems field to given value.


### GetVendorReleaseTimestamp

`func (o *UpdateFirmwareBinary) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *UpdateFirmwareBinary) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *UpdateFirmwareBinary) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *UpdateFirmwareBinary) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetVendor

`func (o *UpdateFirmwareBinary) GetVendor() map[string]interface{}`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *UpdateFirmwareBinary) GetVendorOk() (*map[string]interface{}, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *UpdateFirmwareBinary) SetVendor(v map[string]interface{})`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *UpdateFirmwareBinary) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateFirmwareBinary

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

### NewCreateFirmwareBinary

`func NewCreateFirmwareBinary(catalogId float32, vendorDownloadUrl string, name string, rebootRequired bool, updateSeverity FirmwareBinaryUpdateSeverity, vendorSupportedDevices []map[string]interface{}, vendorSupportedSystems []map[string]interface{}, ) *CreateFirmwareBinary`

NewCreateFirmwareBinary instantiates a new CreateFirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareBinaryWithDefaults

`func NewCreateFirmwareBinaryWithDefaults() *CreateFirmwareBinary`

NewCreateFirmwareBinaryWithDefaults instantiates a new CreateFirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogId

`func (o *CreateFirmwareBinary) GetCatalogId() float32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CreateFirmwareBinary) GetCatalogIdOk() (*float32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CreateFirmwareBinary) SetCatalogId(v float32)`

SetCatalogId sets CatalogId field to given value.


### GetExternalId

`func (o *CreateFirmwareBinary) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateFirmwareBinary) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateFirmwareBinary) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateFirmwareBinary) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVendorInfoUrl

`func (o *CreateFirmwareBinary) GetVendorInfoUrl() string`

GetVendorInfoUrl returns the VendorInfoUrl field if non-nil, zero value otherwise.

### GetVendorInfoUrlOk

`func (o *CreateFirmwareBinary) GetVendorInfoUrlOk() (*string, bool)`

GetVendorInfoUrlOk returns a tuple with the VendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfoUrl

`func (o *CreateFirmwareBinary) SetVendorInfoUrl(v string)`

SetVendorInfoUrl sets VendorInfoUrl field to given value.

### HasVendorInfoUrl

`func (o *CreateFirmwareBinary) HasVendorInfoUrl() bool`

HasVendorInfoUrl returns a boolean if a field has been set.

### GetVendorDownloadUrl

`func (o *CreateFirmwareBinary) GetVendorDownloadUrl() string`

GetVendorDownloadUrl returns the VendorDownloadUrl field if non-nil, zero value otherwise.

### GetVendorDownloadUrlOk

`func (o *CreateFirmwareBinary) GetVendorDownloadUrlOk() (*string, bool)`

GetVendorDownloadUrlOk returns a tuple with the VendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorDownloadUrl

`func (o *CreateFirmwareBinary) SetVendorDownloadUrl(v string)`

SetVendorDownloadUrl sets VendorDownloadUrl field to given value.


### GetCacheDownloadUrl

`func (o *CreateFirmwareBinary) GetCacheDownloadUrl() string`

GetCacheDownloadUrl returns the CacheDownloadUrl field if non-nil, zero value otherwise.

### GetCacheDownloadUrlOk

`func (o *CreateFirmwareBinary) GetCacheDownloadUrlOk() (*string, bool)`

GetCacheDownloadUrlOk returns a tuple with the CacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDownloadUrl

`func (o *CreateFirmwareBinary) SetCacheDownloadUrl(v string)`

SetCacheDownloadUrl sets CacheDownloadUrl field to given value.

### HasCacheDownloadUrl

`func (o *CreateFirmwareBinary) HasCacheDownloadUrl() bool`

HasCacheDownloadUrl returns a boolean if a field has been set.

### GetName

`func (o *CreateFirmwareBinary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirmwareBinary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirmwareBinary) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *CreateFirmwareBinary) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CreateFirmwareBinary) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CreateFirmwareBinary) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CreateFirmwareBinary) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPackageVersion

`func (o *CreateFirmwareBinary) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *CreateFirmwareBinary) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *CreateFirmwareBinary) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *CreateFirmwareBinary) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetRebootRequired

`func (o *CreateFirmwareBinary) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *CreateFirmwareBinary) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *CreateFirmwareBinary) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetUpdateSeverity

`func (o *CreateFirmwareBinary) GetUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetUpdateSeverity returns the UpdateSeverity field if non-nil, zero value otherwise.

### GetUpdateSeverityOk

`func (o *CreateFirmwareBinary) GetUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetUpdateSeverityOk returns a tuple with the UpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeverity

`func (o *CreateFirmwareBinary) SetUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetUpdateSeverity sets UpdateSeverity field to given value.


### GetVendorSupportedDevices

`func (o *CreateFirmwareBinary) GetVendorSupportedDevices() []map[string]interface{}`

GetVendorSupportedDevices returns the VendorSupportedDevices field if non-nil, zero value otherwise.

### GetVendorSupportedDevicesOk

`func (o *CreateFirmwareBinary) GetVendorSupportedDevicesOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedDevicesOk returns a tuple with the VendorSupportedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedDevices

`func (o *CreateFirmwareBinary) SetVendorSupportedDevices(v []map[string]interface{})`

SetVendorSupportedDevices sets VendorSupportedDevices field to given value.


### GetVendorSupportedSystems

`func (o *CreateFirmwareBinary) GetVendorSupportedSystems() []map[string]interface{}`

GetVendorSupportedSystems returns the VendorSupportedSystems field if non-nil, zero value otherwise.

### GetVendorSupportedSystemsOk

`func (o *CreateFirmwareBinary) GetVendorSupportedSystemsOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedSystemsOk returns a tuple with the VendorSupportedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedSystems

`func (o *CreateFirmwareBinary) SetVendorSupportedSystems(v []map[string]interface{})`

SetVendorSupportedSystems sets VendorSupportedSystems field to given value.


### GetVendorReleaseTimestamp

`func (o *CreateFirmwareBinary) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *CreateFirmwareBinary) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *CreateFirmwareBinary) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *CreateFirmwareBinary) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetVendor

`func (o *CreateFirmwareBinary) GetVendor() map[string]interface{}`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CreateFirmwareBinary) GetVendorOk() (*map[string]interface{}, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CreateFirmwareBinary) SetVendor(v map[string]interface{})`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CreateFirmwareBinary) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FirmwareBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
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
**Vendor** | **map[string]interface{}** |  | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewFirmwareBinary

`func NewFirmwareBinary(id float32, catalogId float32, vendorDownloadUrl string, name string, rebootRequired bool, updateSeverity FirmwareBinaryUpdateSeverity, vendorSupportedDevices []map[string]interface{}, vendorSupportedSystems []map[string]interface{}, vendor map[string]interface{}, links map[string]interface{}, ) *FirmwareBinary`

NewFirmwareBinary instantiates a new FirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBinaryWithDefaults

`func NewFirmwareBinaryWithDefaults() *FirmwareBinary`

NewFirmwareBinaryWithDefaults instantiates a new FirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirmwareBinary) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareBinary) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareBinary) SetId(v float32)`

SetId sets Id field to given value.


### GetCatalogId

`func (o *FirmwareBinary) GetCatalogId() float32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *FirmwareBinary) GetCatalogIdOk() (*float32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *FirmwareBinary) SetCatalogId(v float32)`

SetCatalogId sets CatalogId field to given value.


### GetExternalId

`func (o *FirmwareBinary) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *FirmwareBinary) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *FirmwareBinary) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *FirmwareBinary) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVendorInfoUrl

`func (o *FirmwareBinary) GetVendorInfoUrl() string`

GetVendorInfoUrl returns the VendorInfoUrl field if non-nil, zero value otherwise.

### GetVendorInfoUrlOk

`func (o *FirmwareBinary) GetVendorInfoUrlOk() (*string, bool)`

GetVendorInfoUrlOk returns a tuple with the VendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfoUrl

`func (o *FirmwareBinary) SetVendorInfoUrl(v string)`

SetVendorInfoUrl sets VendorInfoUrl field to given value.

### HasVendorInfoUrl

`func (o *FirmwareBinary) HasVendorInfoUrl() bool`

HasVendorInfoUrl returns a boolean if a field has been set.

### GetVendorDownloadUrl

`func (o *FirmwareBinary) GetVendorDownloadUrl() string`

GetVendorDownloadUrl returns the VendorDownloadUrl field if non-nil, zero value otherwise.

### GetVendorDownloadUrlOk

`func (o *FirmwareBinary) GetVendorDownloadUrlOk() (*string, bool)`

GetVendorDownloadUrlOk returns a tuple with the VendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorDownloadUrl

`func (o *FirmwareBinary) SetVendorDownloadUrl(v string)`

SetVendorDownloadUrl sets VendorDownloadUrl field to given value.


### GetCacheDownloadUrl

`func (o *FirmwareBinary) GetCacheDownloadUrl() string`

GetCacheDownloadUrl returns the CacheDownloadUrl field if non-nil, zero value otherwise.

### GetCacheDownloadUrlOk

`func (o *FirmwareBinary) GetCacheDownloadUrlOk() (*string, bool)`

GetCacheDownloadUrlOk returns a tuple with the CacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDownloadUrl

`func (o *FirmwareBinary) SetCacheDownloadUrl(v string)`

SetCacheDownloadUrl sets CacheDownloadUrl field to given value.

### HasCacheDownloadUrl

`func (o *FirmwareBinary) HasCacheDownloadUrl() bool`

HasCacheDownloadUrl returns a boolean if a field has been set.

### GetName

`func (o *FirmwareBinary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareBinary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareBinary) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *FirmwareBinary) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *FirmwareBinary) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *FirmwareBinary) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *FirmwareBinary) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPackageVersion

`func (o *FirmwareBinary) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *FirmwareBinary) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *FirmwareBinary) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *FirmwareBinary) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetRebootRequired

`func (o *FirmwareBinary) GetRebootRequired() bool`

GetRebootRequired returns the RebootRequired field if non-nil, zero value otherwise.

### GetRebootRequiredOk

`func (o *FirmwareBinary) GetRebootRequiredOk() (*bool, bool)`

GetRebootRequiredOk returns a tuple with the RebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequired

`func (o *FirmwareBinary) SetRebootRequired(v bool)`

SetRebootRequired sets RebootRequired field to given value.


### GetUpdateSeverity

`func (o *FirmwareBinary) GetUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetUpdateSeverity returns the UpdateSeverity field if non-nil, zero value otherwise.

### GetUpdateSeverityOk

`func (o *FirmwareBinary) GetUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetUpdateSeverityOk returns a tuple with the UpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSeverity

`func (o *FirmwareBinary) SetUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetUpdateSeverity sets UpdateSeverity field to given value.


### GetVendorSupportedDevices

`func (o *FirmwareBinary) GetVendorSupportedDevices() []map[string]interface{}`

GetVendorSupportedDevices returns the VendorSupportedDevices field if non-nil, zero value otherwise.

### GetVendorSupportedDevicesOk

`func (o *FirmwareBinary) GetVendorSupportedDevicesOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedDevicesOk returns a tuple with the VendorSupportedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedDevices

`func (o *FirmwareBinary) SetVendorSupportedDevices(v []map[string]interface{})`

SetVendorSupportedDevices sets VendorSupportedDevices field to given value.


### GetVendorSupportedSystems

`func (o *FirmwareBinary) GetVendorSupportedSystems() []map[string]interface{}`

GetVendorSupportedSystems returns the VendorSupportedSystems field if non-nil, zero value otherwise.

### GetVendorSupportedSystemsOk

`func (o *FirmwareBinary) GetVendorSupportedSystemsOk() (*[]map[string]interface{}, bool)`

GetVendorSupportedSystemsOk returns a tuple with the VendorSupportedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSupportedSystems

`func (o *FirmwareBinary) SetVendorSupportedSystems(v []map[string]interface{})`

SetVendorSupportedSystems sets VendorSupportedSystems field to given value.


### GetVendorReleaseTimestamp

`func (o *FirmwareBinary) GetVendorReleaseTimestamp() string`

GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetVendorReleaseTimestampOk

`func (o *FirmwareBinary) GetVendorReleaseTimestampOk() (*string, bool)`

GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReleaseTimestamp

`func (o *FirmwareBinary) SetVendorReleaseTimestamp(v string)`

SetVendorReleaseTimestamp sets VendorReleaseTimestamp field to given value.

### HasVendorReleaseTimestamp

`func (o *FirmwareBinary) HasVendorReleaseTimestamp() bool`

HasVendorReleaseTimestamp returns a boolean if a field has been set.

### GetVendor

`func (o *FirmwareBinary) GetVendor() map[string]interface{}`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FirmwareBinary) GetVendorOk() (*map[string]interface{}, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FirmwareBinary) SetVendor(v map[string]interface{})`

SetVendor sets Vendor field to given value.


### GetLinks

`func (o *FirmwareBinary) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirmwareBinary) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirmwareBinary) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



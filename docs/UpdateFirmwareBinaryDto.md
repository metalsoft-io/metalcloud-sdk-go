# UpdateFirmwareBinaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareBinaryCatalogId** | **float32** |  | 
**ServerFirmwareBinaryExternalId** | Pointer to **string** |  | [optional] 
**ServerFirmwareBinaryVendorInfoUrl** | Pointer to **string** |  | [optional] 
**ServerFirmwareBinaryVendorDownloadUrl** | **string** |  | 
**ServerFirmwareBinaryCacheDownloadUrl** | Pointer to **string** |  | [optional] 
**ServerFirmwareBinaryName** | **string** |  | 
**ServerFirmwareBinaryPackageId** | Pointer to **string** |  | [optional] 
**ServerFirmwareBinaryPackageVersion** | Pointer to **string** |  | [optional] 
**ServerFirmwareBinaryRebootRequired** | **bool** |  | 
**ServerFirmwareBinaryUpdateSeverity** | [**FirmwareBinaryUpdateSeverity**](FirmwareBinaryUpdateSeverity.md) |  | 
**ServerFirmwareBinaryVendorSupportedDevicesJson** | **string** |  | 
**ServerFirmwareBinaryVendorSupportedSystemsJson** | **string** |  | 
**ServerFirmwareBinaryVendorReleaseTimestamp** | Pointer to **string** |  | [optional] 
**ServerFirmwareBinaryVendorJson** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateFirmwareBinaryDto

`func NewUpdateFirmwareBinaryDto(serverFirmwareBinaryCatalogId float32, serverFirmwareBinaryVendorDownloadUrl string, serverFirmwareBinaryName string, serverFirmwareBinaryRebootRequired bool, serverFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity, serverFirmwareBinaryVendorSupportedDevicesJson string, serverFirmwareBinaryVendorSupportedSystemsJson string, ) *UpdateFirmwareBinaryDto`

NewUpdateFirmwareBinaryDto instantiates a new UpdateFirmwareBinaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareBinaryDtoWithDefaults

`func NewUpdateFirmwareBinaryDtoWithDefaults() *UpdateFirmwareBinaryDto`

NewUpdateFirmwareBinaryDtoWithDefaults instantiates a new UpdateFirmwareBinaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBinaryCatalogId

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryCatalogId() float32`

GetServerFirmwareBinaryCatalogId returns the ServerFirmwareBinaryCatalogId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCatalogIdOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryCatalogIdOk() (*float32, bool)`

GetServerFirmwareBinaryCatalogIdOk returns a tuple with the ServerFirmwareBinaryCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCatalogId

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryCatalogId(v float32)`

SetServerFirmwareBinaryCatalogId sets ServerFirmwareBinaryCatalogId field to given value.


### GetServerFirmwareBinaryExternalId

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryExternalId() string`

GetServerFirmwareBinaryExternalId returns the ServerFirmwareBinaryExternalId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryExternalIdOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryExternalIdOk() (*string, bool)`

GetServerFirmwareBinaryExternalIdOk returns a tuple with the ServerFirmwareBinaryExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryExternalId

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryExternalId(v string)`

SetServerFirmwareBinaryExternalId sets ServerFirmwareBinaryExternalId field to given value.

### HasServerFirmwareBinaryExternalId

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryExternalId() bool`

HasServerFirmwareBinaryExternalId returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorInfoUrl

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorInfoUrl() string`

GetServerFirmwareBinaryVendorInfoUrl returns the ServerFirmwareBinaryVendorInfoUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorInfoUrlOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorInfoUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorInfoUrlOk returns a tuple with the ServerFirmwareBinaryVendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorInfoUrl

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryVendorInfoUrl(v string)`

SetServerFirmwareBinaryVendorInfoUrl sets ServerFirmwareBinaryVendorInfoUrl field to given value.

### HasServerFirmwareBinaryVendorInfoUrl

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryVendorInfoUrl() bool`

HasServerFirmwareBinaryVendorInfoUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorDownloadUrl

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorDownloadUrl() string`

GetServerFirmwareBinaryVendorDownloadUrl returns the ServerFirmwareBinaryVendorDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorDownloadUrlOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorDownloadUrlOk returns a tuple with the ServerFirmwareBinaryVendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorDownloadUrl

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryVendorDownloadUrl(v string)`

SetServerFirmwareBinaryVendorDownloadUrl sets ServerFirmwareBinaryVendorDownloadUrl field to given value.


### GetServerFirmwareBinaryCacheDownloadUrl

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryCacheDownloadUrl() string`

GetServerFirmwareBinaryCacheDownloadUrl returns the ServerFirmwareBinaryCacheDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCacheDownloadUrlOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryCacheDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryCacheDownloadUrlOk returns a tuple with the ServerFirmwareBinaryCacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCacheDownloadUrl

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryCacheDownloadUrl(v string)`

SetServerFirmwareBinaryCacheDownloadUrl sets ServerFirmwareBinaryCacheDownloadUrl field to given value.

### HasServerFirmwareBinaryCacheDownloadUrl

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryCacheDownloadUrl() bool`

HasServerFirmwareBinaryCacheDownloadUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryName

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryName() string`

GetServerFirmwareBinaryName returns the ServerFirmwareBinaryName field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryNameOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryNameOk() (*string, bool)`

GetServerFirmwareBinaryNameOk returns a tuple with the ServerFirmwareBinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryName

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryName(v string)`

SetServerFirmwareBinaryName sets ServerFirmwareBinaryName field to given value.


### GetServerFirmwareBinaryPackageId

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryPackageId() string`

GetServerFirmwareBinaryPackageId returns the ServerFirmwareBinaryPackageId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageIdOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryPackageIdOk() (*string, bool)`

GetServerFirmwareBinaryPackageIdOk returns a tuple with the ServerFirmwareBinaryPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageId

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryPackageId(v string)`

SetServerFirmwareBinaryPackageId sets ServerFirmwareBinaryPackageId field to given value.

### HasServerFirmwareBinaryPackageId

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryPackageId() bool`

HasServerFirmwareBinaryPackageId returns a boolean if a field has been set.

### GetServerFirmwareBinaryPackageVersion

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryPackageVersion() string`

GetServerFirmwareBinaryPackageVersion returns the ServerFirmwareBinaryPackageVersion field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageVersionOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryPackageVersionOk() (*string, bool)`

GetServerFirmwareBinaryPackageVersionOk returns a tuple with the ServerFirmwareBinaryPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageVersion

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryPackageVersion(v string)`

SetServerFirmwareBinaryPackageVersion sets ServerFirmwareBinaryPackageVersion field to given value.

### HasServerFirmwareBinaryPackageVersion

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryPackageVersion() bool`

HasServerFirmwareBinaryPackageVersion returns a boolean if a field has been set.

### GetServerFirmwareBinaryRebootRequired

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryRebootRequired() bool`

GetServerFirmwareBinaryRebootRequired returns the ServerFirmwareBinaryRebootRequired field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryRebootRequiredOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryRebootRequiredOk() (*bool, bool)`

GetServerFirmwareBinaryRebootRequiredOk returns a tuple with the ServerFirmwareBinaryRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryRebootRequired

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryRebootRequired(v bool)`

SetServerFirmwareBinaryRebootRequired sets ServerFirmwareBinaryRebootRequired field to given value.


### GetServerFirmwareBinaryUpdateSeverity

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetServerFirmwareBinaryUpdateSeverity returns the ServerFirmwareBinaryUpdateSeverity field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryUpdateSeverityOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetServerFirmwareBinaryUpdateSeverityOk returns a tuple with the ServerFirmwareBinaryUpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryUpdateSeverity

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetServerFirmwareBinaryUpdateSeverity sets ServerFirmwareBinaryUpdateSeverity field to given value.


### GetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedDevicesJson() string`

GetServerFirmwareBinaryVendorSupportedDevicesJson returns the ServerFirmwareBinaryVendorSupportedDevicesJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedDevicesJsonOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedDevicesJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedDevicesJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedDevicesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryVendorSupportedDevicesJson(v string)`

SetServerFirmwareBinaryVendorSupportedDevicesJson sets ServerFirmwareBinaryVendorSupportedDevicesJson field to given value.


### GetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedSystemsJson() string`

GetServerFirmwareBinaryVendorSupportedSystemsJson returns the ServerFirmwareBinaryVendorSupportedSystemsJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedSystemsJsonOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedSystemsJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedSystemsJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedSystemsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryVendorSupportedSystemsJson(v string)`

SetServerFirmwareBinaryVendorSupportedSystemsJson sets ServerFirmwareBinaryVendorSupportedSystemsJson field to given value.


### GetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorReleaseTimestamp() string`

GetServerFirmwareBinaryVendorReleaseTimestamp returns the ServerFirmwareBinaryVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorReleaseTimestampOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorReleaseTimestampOk() (*string, bool)`

GetServerFirmwareBinaryVendorReleaseTimestampOk returns a tuple with the ServerFirmwareBinaryVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryVendorReleaseTimestamp(v string)`

SetServerFirmwareBinaryVendorReleaseTimestamp sets ServerFirmwareBinaryVendorReleaseTimestamp field to given value.

### HasServerFirmwareBinaryVendorReleaseTimestamp

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryVendorReleaseTimestamp() bool`

HasServerFirmwareBinaryVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorJson

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorJson() string`

GetServerFirmwareBinaryVendorJson returns the ServerFirmwareBinaryVendorJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorJsonOk

`func (o *UpdateFirmwareBinaryDto) GetServerFirmwareBinaryVendorJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorJsonOk returns a tuple with the ServerFirmwareBinaryVendorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorJson

`func (o *UpdateFirmwareBinaryDto) SetServerFirmwareBinaryVendorJson(v string)`

SetServerFirmwareBinaryVendorJson sets ServerFirmwareBinaryVendorJson field to given value.

### HasServerFirmwareBinaryVendorJson

`func (o *UpdateFirmwareBinaryDto) HasServerFirmwareBinaryVendorJson() bool`

HasServerFirmwareBinaryVendorJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



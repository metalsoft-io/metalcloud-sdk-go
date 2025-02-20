# UpdateFirmwareBinary

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

### NewUpdateFirmwareBinary

`func NewUpdateFirmwareBinary(serverFirmwareBinaryCatalogId float32, serverFirmwareBinaryVendorDownloadUrl string, serverFirmwareBinaryName string, serverFirmwareBinaryRebootRequired bool, serverFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity, serverFirmwareBinaryVendorSupportedDevicesJson string, serverFirmwareBinaryVendorSupportedSystemsJson string, ) *UpdateFirmwareBinary`

NewUpdateFirmwareBinary instantiates a new UpdateFirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareBinaryWithDefaults

`func NewUpdateFirmwareBinaryWithDefaults() *UpdateFirmwareBinary`

NewUpdateFirmwareBinaryWithDefaults instantiates a new UpdateFirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBinaryCatalogId

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryCatalogId() float32`

GetServerFirmwareBinaryCatalogId returns the ServerFirmwareBinaryCatalogId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCatalogIdOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryCatalogIdOk() (*float32, bool)`

GetServerFirmwareBinaryCatalogIdOk returns a tuple with the ServerFirmwareBinaryCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCatalogId

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryCatalogId(v float32)`

SetServerFirmwareBinaryCatalogId sets ServerFirmwareBinaryCatalogId field to given value.


### GetServerFirmwareBinaryExternalId

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryExternalId() string`

GetServerFirmwareBinaryExternalId returns the ServerFirmwareBinaryExternalId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryExternalIdOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryExternalIdOk() (*string, bool)`

GetServerFirmwareBinaryExternalIdOk returns a tuple with the ServerFirmwareBinaryExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryExternalId

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryExternalId(v string)`

SetServerFirmwareBinaryExternalId sets ServerFirmwareBinaryExternalId field to given value.

### HasServerFirmwareBinaryExternalId

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryExternalId() bool`

HasServerFirmwareBinaryExternalId returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorInfoUrl

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorInfoUrl() string`

GetServerFirmwareBinaryVendorInfoUrl returns the ServerFirmwareBinaryVendorInfoUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorInfoUrlOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorInfoUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorInfoUrlOk returns a tuple with the ServerFirmwareBinaryVendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorInfoUrl

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryVendorInfoUrl(v string)`

SetServerFirmwareBinaryVendorInfoUrl sets ServerFirmwareBinaryVendorInfoUrl field to given value.

### HasServerFirmwareBinaryVendorInfoUrl

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryVendorInfoUrl() bool`

HasServerFirmwareBinaryVendorInfoUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorDownloadUrl

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrl() string`

GetServerFirmwareBinaryVendorDownloadUrl returns the ServerFirmwareBinaryVendorDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorDownloadUrlOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorDownloadUrlOk returns a tuple with the ServerFirmwareBinaryVendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorDownloadUrl

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryVendorDownloadUrl(v string)`

SetServerFirmwareBinaryVendorDownloadUrl sets ServerFirmwareBinaryVendorDownloadUrl field to given value.


### GetServerFirmwareBinaryCacheDownloadUrl

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrl() string`

GetServerFirmwareBinaryCacheDownloadUrl returns the ServerFirmwareBinaryCacheDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCacheDownloadUrlOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryCacheDownloadUrlOk returns a tuple with the ServerFirmwareBinaryCacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCacheDownloadUrl

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryCacheDownloadUrl(v string)`

SetServerFirmwareBinaryCacheDownloadUrl sets ServerFirmwareBinaryCacheDownloadUrl field to given value.

### HasServerFirmwareBinaryCacheDownloadUrl

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryCacheDownloadUrl() bool`

HasServerFirmwareBinaryCacheDownloadUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryName

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryName() string`

GetServerFirmwareBinaryName returns the ServerFirmwareBinaryName field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryNameOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryNameOk() (*string, bool)`

GetServerFirmwareBinaryNameOk returns a tuple with the ServerFirmwareBinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryName

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryName(v string)`

SetServerFirmwareBinaryName sets ServerFirmwareBinaryName field to given value.


### GetServerFirmwareBinaryPackageId

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryPackageId() string`

GetServerFirmwareBinaryPackageId returns the ServerFirmwareBinaryPackageId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageIdOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryPackageIdOk() (*string, bool)`

GetServerFirmwareBinaryPackageIdOk returns a tuple with the ServerFirmwareBinaryPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageId

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryPackageId(v string)`

SetServerFirmwareBinaryPackageId sets ServerFirmwareBinaryPackageId field to given value.

### HasServerFirmwareBinaryPackageId

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryPackageId() bool`

HasServerFirmwareBinaryPackageId returns a boolean if a field has been set.

### GetServerFirmwareBinaryPackageVersion

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryPackageVersion() string`

GetServerFirmwareBinaryPackageVersion returns the ServerFirmwareBinaryPackageVersion field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageVersionOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryPackageVersionOk() (*string, bool)`

GetServerFirmwareBinaryPackageVersionOk returns a tuple with the ServerFirmwareBinaryPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageVersion

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryPackageVersion(v string)`

SetServerFirmwareBinaryPackageVersion sets ServerFirmwareBinaryPackageVersion field to given value.

### HasServerFirmwareBinaryPackageVersion

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryPackageVersion() bool`

HasServerFirmwareBinaryPackageVersion returns a boolean if a field has been set.

### GetServerFirmwareBinaryRebootRequired

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryRebootRequired() bool`

GetServerFirmwareBinaryRebootRequired returns the ServerFirmwareBinaryRebootRequired field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryRebootRequiredOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryRebootRequiredOk() (*bool, bool)`

GetServerFirmwareBinaryRebootRequiredOk returns a tuple with the ServerFirmwareBinaryRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryRebootRequired

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryRebootRequired(v bool)`

SetServerFirmwareBinaryRebootRequired sets ServerFirmwareBinaryRebootRequired field to given value.


### GetServerFirmwareBinaryUpdateSeverity

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetServerFirmwareBinaryUpdateSeverity returns the ServerFirmwareBinaryUpdateSeverity field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryUpdateSeverityOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetServerFirmwareBinaryUpdateSeverityOk returns a tuple with the ServerFirmwareBinaryUpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryUpdateSeverity

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetServerFirmwareBinaryUpdateSeverity sets ServerFirmwareBinaryUpdateSeverity field to given value.


### GetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJson() string`

GetServerFirmwareBinaryVendorSupportedDevicesJson returns the ServerFirmwareBinaryVendorSupportedDevicesJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedDevicesJsonOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedDevicesJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedDevicesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryVendorSupportedDevicesJson(v string)`

SetServerFirmwareBinaryVendorSupportedDevicesJson sets ServerFirmwareBinaryVendorSupportedDevicesJson field to given value.


### GetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJson() string`

GetServerFirmwareBinaryVendorSupportedSystemsJson returns the ServerFirmwareBinaryVendorSupportedSystemsJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedSystemsJsonOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedSystemsJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedSystemsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryVendorSupportedSystemsJson(v string)`

SetServerFirmwareBinaryVendorSupportedSystemsJson sets ServerFirmwareBinaryVendorSupportedSystemsJson field to given value.


### GetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestamp() string`

GetServerFirmwareBinaryVendorReleaseTimestamp returns the ServerFirmwareBinaryVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorReleaseTimestampOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestampOk() (*string, bool)`

GetServerFirmwareBinaryVendorReleaseTimestampOk returns a tuple with the ServerFirmwareBinaryVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryVendorReleaseTimestamp(v string)`

SetServerFirmwareBinaryVendorReleaseTimestamp sets ServerFirmwareBinaryVendorReleaseTimestamp field to given value.

### HasServerFirmwareBinaryVendorReleaseTimestamp

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryVendorReleaseTimestamp() bool`

HasServerFirmwareBinaryVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorJson

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorJson() string`

GetServerFirmwareBinaryVendorJson returns the ServerFirmwareBinaryVendorJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorJsonOk

`func (o *UpdateFirmwareBinary) GetServerFirmwareBinaryVendorJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorJsonOk returns a tuple with the ServerFirmwareBinaryVendorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorJson

`func (o *UpdateFirmwareBinary) SetServerFirmwareBinaryVendorJson(v string)`

SetServerFirmwareBinaryVendorJson sets ServerFirmwareBinaryVendorJson field to given value.

### HasServerFirmwareBinaryVendorJson

`func (o *UpdateFirmwareBinary) HasServerFirmwareBinaryVendorJson() bool`

HasServerFirmwareBinaryVendorJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



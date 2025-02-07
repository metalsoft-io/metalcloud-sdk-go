# CreateFirmwareBinaryDto

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

### NewCreateFirmwareBinaryDto

`func NewCreateFirmwareBinaryDto(serverFirmwareBinaryCatalogId float32, serverFirmwareBinaryVendorDownloadUrl string, serverFirmwareBinaryName string, serverFirmwareBinaryRebootRequired bool, serverFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity, serverFirmwareBinaryVendorSupportedDevicesJson string, serverFirmwareBinaryVendorSupportedSystemsJson string, ) *CreateFirmwareBinaryDto`

NewCreateFirmwareBinaryDto instantiates a new CreateFirmwareBinaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareBinaryDtoWithDefaults

`func NewCreateFirmwareBinaryDtoWithDefaults() *CreateFirmwareBinaryDto`

NewCreateFirmwareBinaryDtoWithDefaults instantiates a new CreateFirmwareBinaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBinaryCatalogId

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryCatalogId() float32`

GetServerFirmwareBinaryCatalogId returns the ServerFirmwareBinaryCatalogId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCatalogIdOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryCatalogIdOk() (*float32, bool)`

GetServerFirmwareBinaryCatalogIdOk returns a tuple with the ServerFirmwareBinaryCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCatalogId

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryCatalogId(v float32)`

SetServerFirmwareBinaryCatalogId sets ServerFirmwareBinaryCatalogId field to given value.


### GetServerFirmwareBinaryExternalId

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryExternalId() string`

GetServerFirmwareBinaryExternalId returns the ServerFirmwareBinaryExternalId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryExternalIdOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryExternalIdOk() (*string, bool)`

GetServerFirmwareBinaryExternalIdOk returns a tuple with the ServerFirmwareBinaryExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryExternalId

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryExternalId(v string)`

SetServerFirmwareBinaryExternalId sets ServerFirmwareBinaryExternalId field to given value.

### HasServerFirmwareBinaryExternalId

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryExternalId() bool`

HasServerFirmwareBinaryExternalId returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorInfoUrl

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorInfoUrl() string`

GetServerFirmwareBinaryVendorInfoUrl returns the ServerFirmwareBinaryVendorInfoUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorInfoUrlOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorInfoUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorInfoUrlOk returns a tuple with the ServerFirmwareBinaryVendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorInfoUrl

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryVendorInfoUrl(v string)`

SetServerFirmwareBinaryVendorInfoUrl sets ServerFirmwareBinaryVendorInfoUrl field to given value.

### HasServerFirmwareBinaryVendorInfoUrl

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryVendorInfoUrl() bool`

HasServerFirmwareBinaryVendorInfoUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorDownloadUrl

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorDownloadUrl() string`

GetServerFirmwareBinaryVendorDownloadUrl returns the ServerFirmwareBinaryVendorDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorDownloadUrlOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorDownloadUrlOk returns a tuple with the ServerFirmwareBinaryVendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorDownloadUrl

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryVendorDownloadUrl(v string)`

SetServerFirmwareBinaryVendorDownloadUrl sets ServerFirmwareBinaryVendorDownloadUrl field to given value.


### GetServerFirmwareBinaryCacheDownloadUrl

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryCacheDownloadUrl() string`

GetServerFirmwareBinaryCacheDownloadUrl returns the ServerFirmwareBinaryCacheDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCacheDownloadUrlOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryCacheDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryCacheDownloadUrlOk returns a tuple with the ServerFirmwareBinaryCacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCacheDownloadUrl

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryCacheDownloadUrl(v string)`

SetServerFirmwareBinaryCacheDownloadUrl sets ServerFirmwareBinaryCacheDownloadUrl field to given value.

### HasServerFirmwareBinaryCacheDownloadUrl

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryCacheDownloadUrl() bool`

HasServerFirmwareBinaryCacheDownloadUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryName

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryName() string`

GetServerFirmwareBinaryName returns the ServerFirmwareBinaryName field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryNameOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryNameOk() (*string, bool)`

GetServerFirmwareBinaryNameOk returns a tuple with the ServerFirmwareBinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryName

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryName(v string)`

SetServerFirmwareBinaryName sets ServerFirmwareBinaryName field to given value.


### GetServerFirmwareBinaryPackageId

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryPackageId() string`

GetServerFirmwareBinaryPackageId returns the ServerFirmwareBinaryPackageId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageIdOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryPackageIdOk() (*string, bool)`

GetServerFirmwareBinaryPackageIdOk returns a tuple with the ServerFirmwareBinaryPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageId

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryPackageId(v string)`

SetServerFirmwareBinaryPackageId sets ServerFirmwareBinaryPackageId field to given value.

### HasServerFirmwareBinaryPackageId

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryPackageId() bool`

HasServerFirmwareBinaryPackageId returns a boolean if a field has been set.

### GetServerFirmwareBinaryPackageVersion

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryPackageVersion() string`

GetServerFirmwareBinaryPackageVersion returns the ServerFirmwareBinaryPackageVersion field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageVersionOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryPackageVersionOk() (*string, bool)`

GetServerFirmwareBinaryPackageVersionOk returns a tuple with the ServerFirmwareBinaryPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageVersion

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryPackageVersion(v string)`

SetServerFirmwareBinaryPackageVersion sets ServerFirmwareBinaryPackageVersion field to given value.

### HasServerFirmwareBinaryPackageVersion

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryPackageVersion() bool`

HasServerFirmwareBinaryPackageVersion returns a boolean if a field has been set.

### GetServerFirmwareBinaryRebootRequired

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryRebootRequired() bool`

GetServerFirmwareBinaryRebootRequired returns the ServerFirmwareBinaryRebootRequired field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryRebootRequiredOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryRebootRequiredOk() (*bool, bool)`

GetServerFirmwareBinaryRebootRequiredOk returns a tuple with the ServerFirmwareBinaryRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryRebootRequired

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryRebootRequired(v bool)`

SetServerFirmwareBinaryRebootRequired sets ServerFirmwareBinaryRebootRequired field to given value.


### GetServerFirmwareBinaryUpdateSeverity

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetServerFirmwareBinaryUpdateSeverity returns the ServerFirmwareBinaryUpdateSeverity field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryUpdateSeverityOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetServerFirmwareBinaryUpdateSeverityOk returns a tuple with the ServerFirmwareBinaryUpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryUpdateSeverity

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetServerFirmwareBinaryUpdateSeverity sets ServerFirmwareBinaryUpdateSeverity field to given value.


### GetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedDevicesJson() string`

GetServerFirmwareBinaryVendorSupportedDevicesJson returns the ServerFirmwareBinaryVendorSupportedDevicesJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedDevicesJsonOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedDevicesJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedDevicesJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedDevicesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryVendorSupportedDevicesJson(v string)`

SetServerFirmwareBinaryVendorSupportedDevicesJson sets ServerFirmwareBinaryVendorSupportedDevicesJson field to given value.


### GetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedSystemsJson() string`

GetServerFirmwareBinaryVendorSupportedSystemsJson returns the ServerFirmwareBinaryVendorSupportedSystemsJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedSystemsJsonOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorSupportedSystemsJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedSystemsJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedSystemsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryVendorSupportedSystemsJson(v string)`

SetServerFirmwareBinaryVendorSupportedSystemsJson sets ServerFirmwareBinaryVendorSupportedSystemsJson field to given value.


### GetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorReleaseTimestamp() string`

GetServerFirmwareBinaryVendorReleaseTimestamp returns the ServerFirmwareBinaryVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorReleaseTimestampOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorReleaseTimestampOk() (*string, bool)`

GetServerFirmwareBinaryVendorReleaseTimestampOk returns a tuple with the ServerFirmwareBinaryVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryVendorReleaseTimestamp(v string)`

SetServerFirmwareBinaryVendorReleaseTimestamp sets ServerFirmwareBinaryVendorReleaseTimestamp field to given value.

### HasServerFirmwareBinaryVendorReleaseTimestamp

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryVendorReleaseTimestamp() bool`

HasServerFirmwareBinaryVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorJson

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorJson() string`

GetServerFirmwareBinaryVendorJson returns the ServerFirmwareBinaryVendorJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorJsonOk

`func (o *CreateFirmwareBinaryDto) GetServerFirmwareBinaryVendorJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorJsonOk returns a tuple with the ServerFirmwareBinaryVendorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorJson

`func (o *CreateFirmwareBinaryDto) SetServerFirmwareBinaryVendorJson(v string)`

SetServerFirmwareBinaryVendorJson sets ServerFirmwareBinaryVendorJson field to given value.

### HasServerFirmwareBinaryVendorJson

`func (o *CreateFirmwareBinaryDto) HasServerFirmwareBinaryVendorJson() bool`

HasServerFirmwareBinaryVendorJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



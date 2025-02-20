# CreateFirmwareBinary

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

### NewCreateFirmwareBinary

`func NewCreateFirmwareBinary(serverFirmwareBinaryCatalogId float32, serverFirmwareBinaryVendorDownloadUrl string, serverFirmwareBinaryName string, serverFirmwareBinaryRebootRequired bool, serverFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity, serverFirmwareBinaryVendorSupportedDevicesJson string, serverFirmwareBinaryVendorSupportedSystemsJson string, ) *CreateFirmwareBinary`

NewCreateFirmwareBinary instantiates a new CreateFirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareBinaryWithDefaults

`func NewCreateFirmwareBinaryWithDefaults() *CreateFirmwareBinary`

NewCreateFirmwareBinaryWithDefaults instantiates a new CreateFirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBinaryCatalogId

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCatalogId() float32`

GetServerFirmwareBinaryCatalogId returns the ServerFirmwareBinaryCatalogId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCatalogIdOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCatalogIdOk() (*float32, bool)`

GetServerFirmwareBinaryCatalogIdOk returns a tuple with the ServerFirmwareBinaryCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCatalogId

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryCatalogId(v float32)`

SetServerFirmwareBinaryCatalogId sets ServerFirmwareBinaryCatalogId field to given value.


### GetServerFirmwareBinaryExternalId

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryExternalId() string`

GetServerFirmwareBinaryExternalId returns the ServerFirmwareBinaryExternalId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryExternalIdOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryExternalIdOk() (*string, bool)`

GetServerFirmwareBinaryExternalIdOk returns a tuple with the ServerFirmwareBinaryExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryExternalId

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryExternalId(v string)`

SetServerFirmwareBinaryExternalId sets ServerFirmwareBinaryExternalId field to given value.

### HasServerFirmwareBinaryExternalId

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryExternalId() bool`

HasServerFirmwareBinaryExternalId returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorInfoUrl

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorInfoUrl() string`

GetServerFirmwareBinaryVendorInfoUrl returns the ServerFirmwareBinaryVendorInfoUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorInfoUrlOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorInfoUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorInfoUrlOk returns a tuple with the ServerFirmwareBinaryVendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorInfoUrl

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorInfoUrl(v string)`

SetServerFirmwareBinaryVendorInfoUrl sets ServerFirmwareBinaryVendorInfoUrl field to given value.

### HasServerFirmwareBinaryVendorInfoUrl

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryVendorInfoUrl() bool`

HasServerFirmwareBinaryVendorInfoUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorDownloadUrl

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrl() string`

GetServerFirmwareBinaryVendorDownloadUrl returns the ServerFirmwareBinaryVendorDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorDownloadUrlOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorDownloadUrlOk returns a tuple with the ServerFirmwareBinaryVendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorDownloadUrl

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorDownloadUrl(v string)`

SetServerFirmwareBinaryVendorDownloadUrl sets ServerFirmwareBinaryVendorDownloadUrl field to given value.


### GetServerFirmwareBinaryCacheDownloadUrl

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrl() string`

GetServerFirmwareBinaryCacheDownloadUrl returns the ServerFirmwareBinaryCacheDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCacheDownloadUrlOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryCacheDownloadUrlOk returns a tuple with the ServerFirmwareBinaryCacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCacheDownloadUrl

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryCacheDownloadUrl(v string)`

SetServerFirmwareBinaryCacheDownloadUrl sets ServerFirmwareBinaryCacheDownloadUrl field to given value.

### HasServerFirmwareBinaryCacheDownloadUrl

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryCacheDownloadUrl() bool`

HasServerFirmwareBinaryCacheDownloadUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryName

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryName() string`

GetServerFirmwareBinaryName returns the ServerFirmwareBinaryName field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryNameOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryNameOk() (*string, bool)`

GetServerFirmwareBinaryNameOk returns a tuple with the ServerFirmwareBinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryName

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryName(v string)`

SetServerFirmwareBinaryName sets ServerFirmwareBinaryName field to given value.


### GetServerFirmwareBinaryPackageId

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageId() string`

GetServerFirmwareBinaryPackageId returns the ServerFirmwareBinaryPackageId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageIdOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageIdOk() (*string, bool)`

GetServerFirmwareBinaryPackageIdOk returns a tuple with the ServerFirmwareBinaryPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageId

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryPackageId(v string)`

SetServerFirmwareBinaryPackageId sets ServerFirmwareBinaryPackageId field to given value.

### HasServerFirmwareBinaryPackageId

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryPackageId() bool`

HasServerFirmwareBinaryPackageId returns a boolean if a field has been set.

### GetServerFirmwareBinaryPackageVersion

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageVersion() string`

GetServerFirmwareBinaryPackageVersion returns the ServerFirmwareBinaryPackageVersion field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageVersionOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageVersionOk() (*string, bool)`

GetServerFirmwareBinaryPackageVersionOk returns a tuple with the ServerFirmwareBinaryPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageVersion

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryPackageVersion(v string)`

SetServerFirmwareBinaryPackageVersion sets ServerFirmwareBinaryPackageVersion field to given value.

### HasServerFirmwareBinaryPackageVersion

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryPackageVersion() bool`

HasServerFirmwareBinaryPackageVersion returns a boolean if a field has been set.

### GetServerFirmwareBinaryRebootRequired

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryRebootRequired() bool`

GetServerFirmwareBinaryRebootRequired returns the ServerFirmwareBinaryRebootRequired field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryRebootRequiredOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryRebootRequiredOk() (*bool, bool)`

GetServerFirmwareBinaryRebootRequiredOk returns a tuple with the ServerFirmwareBinaryRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryRebootRequired

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryRebootRequired(v bool)`

SetServerFirmwareBinaryRebootRequired sets ServerFirmwareBinaryRebootRequired field to given value.


### GetServerFirmwareBinaryUpdateSeverity

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetServerFirmwareBinaryUpdateSeverity returns the ServerFirmwareBinaryUpdateSeverity field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryUpdateSeverityOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetServerFirmwareBinaryUpdateSeverityOk returns a tuple with the ServerFirmwareBinaryUpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryUpdateSeverity

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetServerFirmwareBinaryUpdateSeverity sets ServerFirmwareBinaryUpdateSeverity field to given value.


### GetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJson() string`

GetServerFirmwareBinaryVendorSupportedDevicesJson returns the ServerFirmwareBinaryVendorSupportedDevicesJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedDevicesJsonOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedDevicesJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedDevicesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorSupportedDevicesJson(v string)`

SetServerFirmwareBinaryVendorSupportedDevicesJson sets ServerFirmwareBinaryVendorSupportedDevicesJson field to given value.


### GetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJson() string`

GetServerFirmwareBinaryVendorSupportedSystemsJson returns the ServerFirmwareBinaryVendorSupportedSystemsJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedSystemsJsonOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedSystemsJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedSystemsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorSupportedSystemsJson(v string)`

SetServerFirmwareBinaryVendorSupportedSystemsJson sets ServerFirmwareBinaryVendorSupportedSystemsJson field to given value.


### GetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestamp() string`

GetServerFirmwareBinaryVendorReleaseTimestamp returns the ServerFirmwareBinaryVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorReleaseTimestampOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestampOk() (*string, bool)`

GetServerFirmwareBinaryVendorReleaseTimestampOk returns a tuple with the ServerFirmwareBinaryVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorReleaseTimestamp(v string)`

SetServerFirmwareBinaryVendorReleaseTimestamp sets ServerFirmwareBinaryVendorReleaseTimestamp field to given value.

### HasServerFirmwareBinaryVendorReleaseTimestamp

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryVendorReleaseTimestamp() bool`

HasServerFirmwareBinaryVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorJson

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorJson() string`

GetServerFirmwareBinaryVendorJson returns the ServerFirmwareBinaryVendorJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorJsonOk

`func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorJsonOk returns a tuple with the ServerFirmwareBinaryVendorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorJson

`func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorJson(v string)`

SetServerFirmwareBinaryVendorJson sets ServerFirmwareBinaryVendorJson field to given value.

### HasServerFirmwareBinaryVendorJson

`func (o *CreateFirmwareBinary) HasServerFirmwareBinaryVendorJson() bool`

HasServerFirmwareBinaryVendorJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



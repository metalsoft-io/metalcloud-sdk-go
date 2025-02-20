# FirmwareBinary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareBinaryId** | **float32** |  | 
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
**ServerFirmwareBinaryVendorJson** | **string** |  | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewFirmwareBinary

`func NewFirmwareBinary(serverFirmwareBinaryId float32, serverFirmwareBinaryCatalogId float32, serverFirmwareBinaryVendorDownloadUrl string, serverFirmwareBinaryName string, serverFirmwareBinaryRebootRequired bool, serverFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity, serverFirmwareBinaryVendorSupportedDevicesJson string, serverFirmwareBinaryVendorSupportedSystemsJson string, serverFirmwareBinaryVendorJson string, links map[string]interface{}, ) *FirmwareBinary`

NewFirmwareBinary instantiates a new FirmwareBinary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBinaryWithDefaults

`func NewFirmwareBinaryWithDefaults() *FirmwareBinary`

NewFirmwareBinaryWithDefaults instantiates a new FirmwareBinary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBinaryId

`func (o *FirmwareBinary) GetServerFirmwareBinaryId() float32`

GetServerFirmwareBinaryId returns the ServerFirmwareBinaryId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryIdOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryIdOk() (*float32, bool)`

GetServerFirmwareBinaryIdOk returns a tuple with the ServerFirmwareBinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryId

`func (o *FirmwareBinary) SetServerFirmwareBinaryId(v float32)`

SetServerFirmwareBinaryId sets ServerFirmwareBinaryId field to given value.


### GetServerFirmwareBinaryCatalogId

`func (o *FirmwareBinary) GetServerFirmwareBinaryCatalogId() float32`

GetServerFirmwareBinaryCatalogId returns the ServerFirmwareBinaryCatalogId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCatalogIdOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryCatalogIdOk() (*float32, bool)`

GetServerFirmwareBinaryCatalogIdOk returns a tuple with the ServerFirmwareBinaryCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCatalogId

`func (o *FirmwareBinary) SetServerFirmwareBinaryCatalogId(v float32)`

SetServerFirmwareBinaryCatalogId sets ServerFirmwareBinaryCatalogId field to given value.


### GetServerFirmwareBinaryExternalId

`func (o *FirmwareBinary) GetServerFirmwareBinaryExternalId() string`

GetServerFirmwareBinaryExternalId returns the ServerFirmwareBinaryExternalId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryExternalIdOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryExternalIdOk() (*string, bool)`

GetServerFirmwareBinaryExternalIdOk returns a tuple with the ServerFirmwareBinaryExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryExternalId

`func (o *FirmwareBinary) SetServerFirmwareBinaryExternalId(v string)`

SetServerFirmwareBinaryExternalId sets ServerFirmwareBinaryExternalId field to given value.

### HasServerFirmwareBinaryExternalId

`func (o *FirmwareBinary) HasServerFirmwareBinaryExternalId() bool`

HasServerFirmwareBinaryExternalId returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorInfoUrl

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorInfoUrl() string`

GetServerFirmwareBinaryVendorInfoUrl returns the ServerFirmwareBinaryVendorInfoUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorInfoUrlOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorInfoUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorInfoUrlOk returns a tuple with the ServerFirmwareBinaryVendorInfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorInfoUrl

`func (o *FirmwareBinary) SetServerFirmwareBinaryVendorInfoUrl(v string)`

SetServerFirmwareBinaryVendorInfoUrl sets ServerFirmwareBinaryVendorInfoUrl field to given value.

### HasServerFirmwareBinaryVendorInfoUrl

`func (o *FirmwareBinary) HasServerFirmwareBinaryVendorInfoUrl() bool`

HasServerFirmwareBinaryVendorInfoUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorDownloadUrl

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrl() string`

GetServerFirmwareBinaryVendorDownloadUrl returns the ServerFirmwareBinaryVendorDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorDownloadUrlOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryVendorDownloadUrlOk returns a tuple with the ServerFirmwareBinaryVendorDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorDownloadUrl

`func (o *FirmwareBinary) SetServerFirmwareBinaryVendorDownloadUrl(v string)`

SetServerFirmwareBinaryVendorDownloadUrl sets ServerFirmwareBinaryVendorDownloadUrl field to given value.


### GetServerFirmwareBinaryCacheDownloadUrl

`func (o *FirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrl() string`

GetServerFirmwareBinaryCacheDownloadUrl returns the ServerFirmwareBinaryCacheDownloadUrl field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryCacheDownloadUrlOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrlOk() (*string, bool)`

GetServerFirmwareBinaryCacheDownloadUrlOk returns a tuple with the ServerFirmwareBinaryCacheDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryCacheDownloadUrl

`func (o *FirmwareBinary) SetServerFirmwareBinaryCacheDownloadUrl(v string)`

SetServerFirmwareBinaryCacheDownloadUrl sets ServerFirmwareBinaryCacheDownloadUrl field to given value.

### HasServerFirmwareBinaryCacheDownloadUrl

`func (o *FirmwareBinary) HasServerFirmwareBinaryCacheDownloadUrl() bool`

HasServerFirmwareBinaryCacheDownloadUrl returns a boolean if a field has been set.

### GetServerFirmwareBinaryName

`func (o *FirmwareBinary) GetServerFirmwareBinaryName() string`

GetServerFirmwareBinaryName returns the ServerFirmwareBinaryName field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryNameOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryNameOk() (*string, bool)`

GetServerFirmwareBinaryNameOk returns a tuple with the ServerFirmwareBinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryName

`func (o *FirmwareBinary) SetServerFirmwareBinaryName(v string)`

SetServerFirmwareBinaryName sets ServerFirmwareBinaryName field to given value.


### GetServerFirmwareBinaryPackageId

`func (o *FirmwareBinary) GetServerFirmwareBinaryPackageId() string`

GetServerFirmwareBinaryPackageId returns the ServerFirmwareBinaryPackageId field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageIdOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryPackageIdOk() (*string, bool)`

GetServerFirmwareBinaryPackageIdOk returns a tuple with the ServerFirmwareBinaryPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageId

`func (o *FirmwareBinary) SetServerFirmwareBinaryPackageId(v string)`

SetServerFirmwareBinaryPackageId sets ServerFirmwareBinaryPackageId field to given value.

### HasServerFirmwareBinaryPackageId

`func (o *FirmwareBinary) HasServerFirmwareBinaryPackageId() bool`

HasServerFirmwareBinaryPackageId returns a boolean if a field has been set.

### GetServerFirmwareBinaryPackageVersion

`func (o *FirmwareBinary) GetServerFirmwareBinaryPackageVersion() string`

GetServerFirmwareBinaryPackageVersion returns the ServerFirmwareBinaryPackageVersion field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryPackageVersionOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryPackageVersionOk() (*string, bool)`

GetServerFirmwareBinaryPackageVersionOk returns a tuple with the ServerFirmwareBinaryPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryPackageVersion

`func (o *FirmwareBinary) SetServerFirmwareBinaryPackageVersion(v string)`

SetServerFirmwareBinaryPackageVersion sets ServerFirmwareBinaryPackageVersion field to given value.

### HasServerFirmwareBinaryPackageVersion

`func (o *FirmwareBinary) HasServerFirmwareBinaryPackageVersion() bool`

HasServerFirmwareBinaryPackageVersion returns a boolean if a field has been set.

### GetServerFirmwareBinaryRebootRequired

`func (o *FirmwareBinary) GetServerFirmwareBinaryRebootRequired() bool`

GetServerFirmwareBinaryRebootRequired returns the ServerFirmwareBinaryRebootRequired field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryRebootRequiredOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryRebootRequiredOk() (*bool, bool)`

GetServerFirmwareBinaryRebootRequiredOk returns a tuple with the ServerFirmwareBinaryRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryRebootRequired

`func (o *FirmwareBinary) SetServerFirmwareBinaryRebootRequired(v bool)`

SetServerFirmwareBinaryRebootRequired sets ServerFirmwareBinaryRebootRequired field to given value.


### GetServerFirmwareBinaryUpdateSeverity

`func (o *FirmwareBinary) GetServerFirmwareBinaryUpdateSeverity() FirmwareBinaryUpdateSeverity`

GetServerFirmwareBinaryUpdateSeverity returns the ServerFirmwareBinaryUpdateSeverity field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryUpdateSeverityOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool)`

GetServerFirmwareBinaryUpdateSeverityOk returns a tuple with the ServerFirmwareBinaryUpdateSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryUpdateSeverity

`func (o *FirmwareBinary) SetServerFirmwareBinaryUpdateSeverity(v FirmwareBinaryUpdateSeverity)`

SetServerFirmwareBinaryUpdateSeverity sets ServerFirmwareBinaryUpdateSeverity field to given value.


### GetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJson() string`

GetServerFirmwareBinaryVendorSupportedDevicesJson returns the ServerFirmwareBinaryVendorSupportedDevicesJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedDevicesJsonOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedDevicesJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedDevicesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedDevicesJson

`func (o *FirmwareBinary) SetServerFirmwareBinaryVendorSupportedDevicesJson(v string)`

SetServerFirmwareBinaryVendorSupportedDevicesJson sets ServerFirmwareBinaryVendorSupportedDevicesJson field to given value.


### GetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJson() string`

GetServerFirmwareBinaryVendorSupportedSystemsJson returns the ServerFirmwareBinaryVendorSupportedSystemsJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorSupportedSystemsJsonOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorSupportedSystemsJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedSystemsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorSupportedSystemsJson

`func (o *FirmwareBinary) SetServerFirmwareBinaryVendorSupportedSystemsJson(v string)`

SetServerFirmwareBinaryVendorSupportedSystemsJson sets ServerFirmwareBinaryVendorSupportedSystemsJson field to given value.


### GetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestamp() string`

GetServerFirmwareBinaryVendorReleaseTimestamp returns the ServerFirmwareBinaryVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorReleaseTimestampOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestampOk() (*string, bool)`

GetServerFirmwareBinaryVendorReleaseTimestampOk returns a tuple with the ServerFirmwareBinaryVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorReleaseTimestamp

`func (o *FirmwareBinary) SetServerFirmwareBinaryVendorReleaseTimestamp(v string)`

SetServerFirmwareBinaryVendorReleaseTimestamp sets ServerFirmwareBinaryVendorReleaseTimestamp field to given value.

### HasServerFirmwareBinaryVendorReleaseTimestamp

`func (o *FirmwareBinary) HasServerFirmwareBinaryVendorReleaseTimestamp() bool`

HasServerFirmwareBinaryVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareBinaryVendorJson

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorJson() string`

GetServerFirmwareBinaryVendorJson returns the ServerFirmwareBinaryVendorJson field if non-nil, zero value otherwise.

### GetServerFirmwareBinaryVendorJsonOk

`func (o *FirmwareBinary) GetServerFirmwareBinaryVendorJsonOk() (*string, bool)`

GetServerFirmwareBinaryVendorJsonOk returns a tuple with the ServerFirmwareBinaryVendorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBinaryVendorJson

`func (o *FirmwareBinary) SetServerFirmwareBinaryVendorJson(v string)`

SetServerFirmwareBinaryVendorJson sets ServerFirmwareBinaryVendorJson field to given value.


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



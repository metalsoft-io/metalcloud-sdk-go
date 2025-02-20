# UpdateFirmwareCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareCatalogName** | **string** |  | 
**ServerFirmwareCatalogDescription** | Pointer to **string** |  | [optional] 
**ServerFirmwareCatalogVendor** | [**FirmwareVendorType**](FirmwareVendorType.md) |  | 
**ServerFirmwareCatalogUpdateType** | [**CatalogUpdateType**](CatalogUpdateType.md) |  | 
**ServerFirmwareCatalogVendorId** | Pointer to **string** |  | [optional] 
**ServerFirmwareCatalogVendorUrl** | Pointer to **string** |  | [optional] 
**ServerFirmwareCatalogVendorReleaseTimestamp** | Pointer to **time.Time** |  | [optional] 
**ServerFirmwareCatalogMetalsoftServerTypesSupportedJson** | Pointer to **string** |  | [optional] 
**ServerFirmwareCatalogVendorServerTypesSupportedJson** | Pointer to **string** | Serialized JSON object of the server types supported by the vendor for this catalog | [optional] 
**ServerFirmwareCatalogVendorConfigurationJson** | Pointer to **string** | Serialized JSON object of the vendor configuration for this catalog | [optional] 

## Methods

### NewUpdateFirmwareCatalog

`func NewUpdateFirmwareCatalog(serverFirmwareCatalogName string, serverFirmwareCatalogVendor FirmwareVendorType, serverFirmwareCatalogUpdateType CatalogUpdateType, ) *UpdateFirmwareCatalog`

NewUpdateFirmwareCatalog instantiates a new UpdateFirmwareCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareCatalogWithDefaults

`func NewUpdateFirmwareCatalogWithDefaults() *UpdateFirmwareCatalog`

NewUpdateFirmwareCatalogWithDefaults instantiates a new UpdateFirmwareCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareCatalogName

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogName() string`

GetServerFirmwareCatalogName returns the ServerFirmwareCatalogName field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogNameOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogNameOk() (*string, bool)`

GetServerFirmwareCatalogNameOk returns a tuple with the ServerFirmwareCatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogName

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogName(v string)`

SetServerFirmwareCatalogName sets ServerFirmwareCatalogName field to given value.


### GetServerFirmwareCatalogDescription

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogDescription() string`

GetServerFirmwareCatalogDescription returns the ServerFirmwareCatalogDescription field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogDescriptionOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogDescriptionOk() (*string, bool)`

GetServerFirmwareCatalogDescriptionOk returns a tuple with the ServerFirmwareCatalogDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogDescription

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogDescription(v string)`

SetServerFirmwareCatalogDescription sets ServerFirmwareCatalogDescription field to given value.

### HasServerFirmwareCatalogDescription

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogDescription() bool`

HasServerFirmwareCatalogDescription returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendor

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendor() FirmwareVendorType`

GetServerFirmwareCatalogVendor returns the ServerFirmwareCatalogVendor field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorOk() (*FirmwareVendorType, bool)`

GetServerFirmwareCatalogVendorOk returns a tuple with the ServerFirmwareCatalogVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendor

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogVendor(v FirmwareVendorType)`

SetServerFirmwareCatalogVendor sets ServerFirmwareCatalogVendor field to given value.


### GetServerFirmwareCatalogUpdateType

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogUpdateType() CatalogUpdateType`

GetServerFirmwareCatalogUpdateType returns the ServerFirmwareCatalogUpdateType field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogUpdateTypeOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogUpdateTypeOk() (*CatalogUpdateType, bool)`

GetServerFirmwareCatalogUpdateTypeOk returns a tuple with the ServerFirmwareCatalogUpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogUpdateType

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogUpdateType(v CatalogUpdateType)`

SetServerFirmwareCatalogUpdateType sets ServerFirmwareCatalogUpdateType field to given value.


### GetServerFirmwareCatalogVendorId

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorId() string`

GetServerFirmwareCatalogVendorId returns the ServerFirmwareCatalogVendorId field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorIdOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorIdOk() (*string, bool)`

GetServerFirmwareCatalogVendorIdOk returns a tuple with the ServerFirmwareCatalogVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorId

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogVendorId(v string)`

SetServerFirmwareCatalogVendorId sets ServerFirmwareCatalogVendorId field to given value.

### HasServerFirmwareCatalogVendorId

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogVendorId() bool`

HasServerFirmwareCatalogVendorId returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorUrl

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorUrl() string`

GetServerFirmwareCatalogVendorUrl returns the ServerFirmwareCatalogVendorUrl field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorUrlOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorUrlOk() (*string, bool)`

GetServerFirmwareCatalogVendorUrlOk returns a tuple with the ServerFirmwareCatalogVendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorUrl

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogVendorUrl(v string)`

SetServerFirmwareCatalogVendorUrl sets ServerFirmwareCatalogVendorUrl field to given value.

### HasServerFirmwareCatalogVendorUrl

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogVendorUrl() bool`

HasServerFirmwareCatalogVendorUrl returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestamp() time.Time`

GetServerFirmwareCatalogVendorReleaseTimestamp returns the ServerFirmwareCatalogVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorReleaseTimestampOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestampOk() (*time.Time, bool)`

GetServerFirmwareCatalogVendorReleaseTimestampOk returns a tuple with the ServerFirmwareCatalogVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogVendorReleaseTimestamp(v time.Time)`

SetServerFirmwareCatalogVendorReleaseTimestamp sets ServerFirmwareCatalogVendorReleaseTimestamp field to given value.

### HasServerFirmwareCatalogVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogVendorReleaseTimestamp() bool`

HasServerFirmwareCatalogVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson() string`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson sets ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson() bool`

HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJson() string`

GetServerFirmwareCatalogVendorServerTypesSupportedJson returns the ServerFirmwareCatalogVendorServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogVendorServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogVendorServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogVendorServerTypesSupportedJson sets ServerFirmwareCatalogVendorServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogVendorServerTypesSupportedJson() bool`

HasServerFirmwareCatalogVendorServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorConfigurationJson

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJson() string`

GetServerFirmwareCatalogVendorConfigurationJson returns the ServerFirmwareCatalogVendorConfigurationJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorConfigurationJsonOk

`func (o *UpdateFirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorConfigurationJsonOk returns a tuple with the ServerFirmwareCatalogVendorConfigurationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorConfigurationJson

`func (o *UpdateFirmwareCatalog) SetServerFirmwareCatalogVendorConfigurationJson(v string)`

SetServerFirmwareCatalogVendorConfigurationJson sets ServerFirmwareCatalogVendorConfigurationJson field to given value.

### HasServerFirmwareCatalogVendorConfigurationJson

`func (o *UpdateFirmwareCatalog) HasServerFirmwareCatalogVendorConfigurationJson() bool`

HasServerFirmwareCatalogVendorConfigurationJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



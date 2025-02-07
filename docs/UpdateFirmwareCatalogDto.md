# UpdateFirmwareCatalogDto

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

### NewUpdateFirmwareCatalogDto

`func NewUpdateFirmwareCatalogDto(serverFirmwareCatalogName string, serverFirmwareCatalogVendor FirmwareVendorType, serverFirmwareCatalogUpdateType CatalogUpdateType, ) *UpdateFirmwareCatalogDto`

NewUpdateFirmwareCatalogDto instantiates a new UpdateFirmwareCatalogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareCatalogDtoWithDefaults

`func NewUpdateFirmwareCatalogDtoWithDefaults() *UpdateFirmwareCatalogDto`

NewUpdateFirmwareCatalogDtoWithDefaults instantiates a new UpdateFirmwareCatalogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareCatalogName

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogName() string`

GetServerFirmwareCatalogName returns the ServerFirmwareCatalogName field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogNameOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogNameOk() (*string, bool)`

GetServerFirmwareCatalogNameOk returns a tuple with the ServerFirmwareCatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogName

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogName(v string)`

SetServerFirmwareCatalogName sets ServerFirmwareCatalogName field to given value.


### GetServerFirmwareCatalogDescription

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogDescription() string`

GetServerFirmwareCatalogDescription returns the ServerFirmwareCatalogDescription field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogDescriptionOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogDescriptionOk() (*string, bool)`

GetServerFirmwareCatalogDescriptionOk returns a tuple with the ServerFirmwareCatalogDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogDescription

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogDescription(v string)`

SetServerFirmwareCatalogDescription sets ServerFirmwareCatalogDescription field to given value.

### HasServerFirmwareCatalogDescription

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogDescription() bool`

HasServerFirmwareCatalogDescription returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendor

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendor() FirmwareVendorType`

GetServerFirmwareCatalogVendor returns the ServerFirmwareCatalogVendor field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorOk() (*FirmwareVendorType, bool)`

GetServerFirmwareCatalogVendorOk returns a tuple with the ServerFirmwareCatalogVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendor

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogVendor(v FirmwareVendorType)`

SetServerFirmwareCatalogVendor sets ServerFirmwareCatalogVendor field to given value.


### GetServerFirmwareCatalogUpdateType

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogUpdateType() CatalogUpdateType`

GetServerFirmwareCatalogUpdateType returns the ServerFirmwareCatalogUpdateType field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogUpdateTypeOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogUpdateTypeOk() (*CatalogUpdateType, bool)`

GetServerFirmwareCatalogUpdateTypeOk returns a tuple with the ServerFirmwareCatalogUpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogUpdateType

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogUpdateType(v CatalogUpdateType)`

SetServerFirmwareCatalogUpdateType sets ServerFirmwareCatalogUpdateType field to given value.


### GetServerFirmwareCatalogVendorId

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorId() string`

GetServerFirmwareCatalogVendorId returns the ServerFirmwareCatalogVendorId field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorIdOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorIdOk() (*string, bool)`

GetServerFirmwareCatalogVendorIdOk returns a tuple with the ServerFirmwareCatalogVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorId

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogVendorId(v string)`

SetServerFirmwareCatalogVendorId sets ServerFirmwareCatalogVendorId field to given value.

### HasServerFirmwareCatalogVendorId

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogVendorId() bool`

HasServerFirmwareCatalogVendorId returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorUrl

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorUrl() string`

GetServerFirmwareCatalogVendorUrl returns the ServerFirmwareCatalogVendorUrl field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorUrlOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorUrlOk() (*string, bool)`

GetServerFirmwareCatalogVendorUrlOk returns a tuple with the ServerFirmwareCatalogVendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorUrl

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogVendorUrl(v string)`

SetServerFirmwareCatalogVendorUrl sets ServerFirmwareCatalogVendorUrl field to given value.

### HasServerFirmwareCatalogVendorUrl

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogVendorUrl() bool`

HasServerFirmwareCatalogVendorUrl returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorReleaseTimestamp() time.Time`

GetServerFirmwareCatalogVendorReleaseTimestamp returns the ServerFirmwareCatalogVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorReleaseTimestampOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorReleaseTimestampOk() (*time.Time, bool)`

GetServerFirmwareCatalogVendorReleaseTimestampOk returns a tuple with the ServerFirmwareCatalogVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogVendorReleaseTimestamp(v time.Time)`

SetServerFirmwareCatalogVendorReleaseTimestamp sets ServerFirmwareCatalogVendorReleaseTimestamp field to given value.

### HasServerFirmwareCatalogVendorReleaseTimestamp

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogVendorReleaseTimestamp() bool`

HasServerFirmwareCatalogVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson() string`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson sets ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson() bool`

HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorServerTypesSupportedJson() string`

GetServerFirmwareCatalogVendorServerTypesSupportedJson returns the ServerFirmwareCatalogVendorServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogVendorServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogVendorServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogVendorServerTypesSupportedJson sets ServerFirmwareCatalogVendorServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogVendorServerTypesSupportedJson() bool`

HasServerFirmwareCatalogVendorServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorConfigurationJson

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorConfigurationJson() string`

GetServerFirmwareCatalogVendorConfigurationJson returns the ServerFirmwareCatalogVendorConfigurationJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorConfigurationJsonOk

`func (o *UpdateFirmwareCatalogDto) GetServerFirmwareCatalogVendorConfigurationJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorConfigurationJsonOk returns a tuple with the ServerFirmwareCatalogVendorConfigurationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorConfigurationJson

`func (o *UpdateFirmwareCatalogDto) SetServerFirmwareCatalogVendorConfigurationJson(v string)`

SetServerFirmwareCatalogVendorConfigurationJson sets ServerFirmwareCatalogVendorConfigurationJson field to given value.

### HasServerFirmwareCatalogVendorConfigurationJson

`func (o *UpdateFirmwareCatalogDto) HasServerFirmwareCatalogVendorConfigurationJson() bool`

HasServerFirmwareCatalogVendorConfigurationJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



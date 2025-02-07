# CreateFirmwareCatalogDto

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

### NewCreateFirmwareCatalogDto

`func NewCreateFirmwareCatalogDto(serverFirmwareCatalogName string, serverFirmwareCatalogVendor FirmwareVendorType, serverFirmwareCatalogUpdateType CatalogUpdateType, ) *CreateFirmwareCatalogDto`

NewCreateFirmwareCatalogDto instantiates a new CreateFirmwareCatalogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareCatalogDtoWithDefaults

`func NewCreateFirmwareCatalogDtoWithDefaults() *CreateFirmwareCatalogDto`

NewCreateFirmwareCatalogDtoWithDefaults instantiates a new CreateFirmwareCatalogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareCatalogName

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogName() string`

GetServerFirmwareCatalogName returns the ServerFirmwareCatalogName field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogNameOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogNameOk() (*string, bool)`

GetServerFirmwareCatalogNameOk returns a tuple with the ServerFirmwareCatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogName

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogName(v string)`

SetServerFirmwareCatalogName sets ServerFirmwareCatalogName field to given value.


### GetServerFirmwareCatalogDescription

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogDescription() string`

GetServerFirmwareCatalogDescription returns the ServerFirmwareCatalogDescription field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogDescriptionOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogDescriptionOk() (*string, bool)`

GetServerFirmwareCatalogDescriptionOk returns a tuple with the ServerFirmwareCatalogDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogDescription

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogDescription(v string)`

SetServerFirmwareCatalogDescription sets ServerFirmwareCatalogDescription field to given value.

### HasServerFirmwareCatalogDescription

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogDescription() bool`

HasServerFirmwareCatalogDescription returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendor

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendor() FirmwareVendorType`

GetServerFirmwareCatalogVendor returns the ServerFirmwareCatalogVendor field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorOk() (*FirmwareVendorType, bool)`

GetServerFirmwareCatalogVendorOk returns a tuple with the ServerFirmwareCatalogVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendor

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogVendor(v FirmwareVendorType)`

SetServerFirmwareCatalogVendor sets ServerFirmwareCatalogVendor field to given value.


### GetServerFirmwareCatalogUpdateType

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogUpdateType() CatalogUpdateType`

GetServerFirmwareCatalogUpdateType returns the ServerFirmwareCatalogUpdateType field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogUpdateTypeOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogUpdateTypeOk() (*CatalogUpdateType, bool)`

GetServerFirmwareCatalogUpdateTypeOk returns a tuple with the ServerFirmwareCatalogUpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogUpdateType

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogUpdateType(v CatalogUpdateType)`

SetServerFirmwareCatalogUpdateType sets ServerFirmwareCatalogUpdateType field to given value.


### GetServerFirmwareCatalogVendorId

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorId() string`

GetServerFirmwareCatalogVendorId returns the ServerFirmwareCatalogVendorId field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorIdOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorIdOk() (*string, bool)`

GetServerFirmwareCatalogVendorIdOk returns a tuple with the ServerFirmwareCatalogVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorId

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogVendorId(v string)`

SetServerFirmwareCatalogVendorId sets ServerFirmwareCatalogVendorId field to given value.

### HasServerFirmwareCatalogVendorId

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogVendorId() bool`

HasServerFirmwareCatalogVendorId returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorUrl

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorUrl() string`

GetServerFirmwareCatalogVendorUrl returns the ServerFirmwareCatalogVendorUrl field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorUrlOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorUrlOk() (*string, bool)`

GetServerFirmwareCatalogVendorUrlOk returns a tuple with the ServerFirmwareCatalogVendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorUrl

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogVendorUrl(v string)`

SetServerFirmwareCatalogVendorUrl sets ServerFirmwareCatalogVendorUrl field to given value.

### HasServerFirmwareCatalogVendorUrl

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogVendorUrl() bool`

HasServerFirmwareCatalogVendorUrl returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorReleaseTimestamp() time.Time`

GetServerFirmwareCatalogVendorReleaseTimestamp returns the ServerFirmwareCatalogVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorReleaseTimestampOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorReleaseTimestampOk() (*time.Time, bool)`

GetServerFirmwareCatalogVendorReleaseTimestampOk returns a tuple with the ServerFirmwareCatalogVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogVendorReleaseTimestamp(v time.Time)`

SetServerFirmwareCatalogVendorReleaseTimestamp sets ServerFirmwareCatalogVendorReleaseTimestamp field to given value.

### HasServerFirmwareCatalogVendorReleaseTimestamp

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogVendorReleaseTimestamp() bool`

HasServerFirmwareCatalogVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson() string`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson sets ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson() bool`

HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorServerTypesSupportedJson() string`

GetServerFirmwareCatalogVendorServerTypesSupportedJson returns the ServerFirmwareCatalogVendorServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogVendorServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogVendorServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogVendorServerTypesSupportedJson sets ServerFirmwareCatalogVendorServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogVendorServerTypesSupportedJson() bool`

HasServerFirmwareCatalogVendorServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorConfigurationJson

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorConfigurationJson() string`

GetServerFirmwareCatalogVendorConfigurationJson returns the ServerFirmwareCatalogVendorConfigurationJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorConfigurationJsonOk

`func (o *CreateFirmwareCatalogDto) GetServerFirmwareCatalogVendorConfigurationJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorConfigurationJsonOk returns a tuple with the ServerFirmwareCatalogVendorConfigurationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorConfigurationJson

`func (o *CreateFirmwareCatalogDto) SetServerFirmwareCatalogVendorConfigurationJson(v string)`

SetServerFirmwareCatalogVendorConfigurationJson sets ServerFirmwareCatalogVendorConfigurationJson field to given value.

### HasServerFirmwareCatalogVendorConfigurationJson

`func (o *CreateFirmwareCatalogDto) HasServerFirmwareCatalogVendorConfigurationJson() bool`

HasServerFirmwareCatalogVendorConfigurationJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



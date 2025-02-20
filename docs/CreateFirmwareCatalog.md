# CreateFirmwareCatalog

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

### NewCreateFirmwareCatalog

`func NewCreateFirmwareCatalog(serverFirmwareCatalogName string, serverFirmwareCatalogVendor FirmwareVendorType, serverFirmwareCatalogUpdateType CatalogUpdateType, ) *CreateFirmwareCatalog`

NewCreateFirmwareCatalog instantiates a new CreateFirmwareCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareCatalogWithDefaults

`func NewCreateFirmwareCatalogWithDefaults() *CreateFirmwareCatalog`

NewCreateFirmwareCatalogWithDefaults instantiates a new CreateFirmwareCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareCatalogName

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogName() string`

GetServerFirmwareCatalogName returns the ServerFirmwareCatalogName field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogNameOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogNameOk() (*string, bool)`

GetServerFirmwareCatalogNameOk returns a tuple with the ServerFirmwareCatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogName

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogName(v string)`

SetServerFirmwareCatalogName sets ServerFirmwareCatalogName field to given value.


### GetServerFirmwareCatalogDescription

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogDescription() string`

GetServerFirmwareCatalogDescription returns the ServerFirmwareCatalogDescription field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogDescriptionOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogDescriptionOk() (*string, bool)`

GetServerFirmwareCatalogDescriptionOk returns a tuple with the ServerFirmwareCatalogDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogDescription

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogDescription(v string)`

SetServerFirmwareCatalogDescription sets ServerFirmwareCatalogDescription field to given value.

### HasServerFirmwareCatalogDescription

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogDescription() bool`

HasServerFirmwareCatalogDescription returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendor

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendor() FirmwareVendorType`

GetServerFirmwareCatalogVendor returns the ServerFirmwareCatalogVendor field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorOk() (*FirmwareVendorType, bool)`

GetServerFirmwareCatalogVendorOk returns a tuple with the ServerFirmwareCatalogVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendor

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogVendor(v FirmwareVendorType)`

SetServerFirmwareCatalogVendor sets ServerFirmwareCatalogVendor field to given value.


### GetServerFirmwareCatalogUpdateType

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogUpdateType() CatalogUpdateType`

GetServerFirmwareCatalogUpdateType returns the ServerFirmwareCatalogUpdateType field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogUpdateTypeOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogUpdateTypeOk() (*CatalogUpdateType, bool)`

GetServerFirmwareCatalogUpdateTypeOk returns a tuple with the ServerFirmwareCatalogUpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogUpdateType

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogUpdateType(v CatalogUpdateType)`

SetServerFirmwareCatalogUpdateType sets ServerFirmwareCatalogUpdateType field to given value.


### GetServerFirmwareCatalogVendorId

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorId() string`

GetServerFirmwareCatalogVendorId returns the ServerFirmwareCatalogVendorId field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorIdOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorIdOk() (*string, bool)`

GetServerFirmwareCatalogVendorIdOk returns a tuple with the ServerFirmwareCatalogVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorId

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogVendorId(v string)`

SetServerFirmwareCatalogVendorId sets ServerFirmwareCatalogVendorId field to given value.

### HasServerFirmwareCatalogVendorId

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogVendorId() bool`

HasServerFirmwareCatalogVendorId returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorUrl

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorUrl() string`

GetServerFirmwareCatalogVendorUrl returns the ServerFirmwareCatalogVendorUrl field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorUrlOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorUrlOk() (*string, bool)`

GetServerFirmwareCatalogVendorUrlOk returns a tuple with the ServerFirmwareCatalogVendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorUrl

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogVendorUrl(v string)`

SetServerFirmwareCatalogVendorUrl sets ServerFirmwareCatalogVendorUrl field to given value.

### HasServerFirmwareCatalogVendorUrl

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogVendorUrl() bool`

HasServerFirmwareCatalogVendorUrl returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestamp() time.Time`

GetServerFirmwareCatalogVendorReleaseTimestamp returns the ServerFirmwareCatalogVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorReleaseTimestampOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestampOk() (*time.Time, bool)`

GetServerFirmwareCatalogVendorReleaseTimestampOk returns a tuple with the ServerFirmwareCatalogVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogVendorReleaseTimestamp(v time.Time)`

SetServerFirmwareCatalogVendorReleaseTimestamp sets ServerFirmwareCatalogVendorReleaseTimestamp field to given value.

### HasServerFirmwareCatalogVendorReleaseTimestamp

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogVendorReleaseTimestamp() bool`

HasServerFirmwareCatalogVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson() string`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson sets ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson() bool`

HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJson() string`

GetServerFirmwareCatalogVendorServerTypesSupportedJson returns the ServerFirmwareCatalogVendorServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogVendorServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogVendorServerTypesSupportedJson(v string)`

SetServerFirmwareCatalogVendorServerTypesSupportedJson sets ServerFirmwareCatalogVendorServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogVendorServerTypesSupportedJson() bool`

HasServerFirmwareCatalogVendorServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorConfigurationJson

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJson() string`

GetServerFirmwareCatalogVendorConfigurationJson returns the ServerFirmwareCatalogVendorConfigurationJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorConfigurationJsonOk

`func (o *CreateFirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJsonOk() (*string, bool)`

GetServerFirmwareCatalogVendorConfigurationJsonOk returns a tuple with the ServerFirmwareCatalogVendorConfigurationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorConfigurationJson

`func (o *CreateFirmwareCatalog) SetServerFirmwareCatalogVendorConfigurationJson(v string)`

SetServerFirmwareCatalogVendorConfigurationJson sets ServerFirmwareCatalogVendorConfigurationJson field to given value.

### HasServerFirmwareCatalogVendorConfigurationJson

`func (o *CreateFirmwareCatalog) HasServerFirmwareCatalogVendorConfigurationJson() bool`

HasServerFirmwareCatalogVendorConfigurationJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



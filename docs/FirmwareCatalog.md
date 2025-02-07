# FirmwareCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareCatalogId** | **float32** | Unique identifier for the firmware catalog | 
**ServerFirmwareCatalogName** | **string** | Name of the catalog, must be unique | 
**ServerFirmwareCatalogDescription** | Pointer to **string** | User description of the catalog | [optional] 
**ServerFirmwareCatalogVendor** | **string** | Firmware catalog vendor: dell, hp, lenovo | 
**ServerFirmwareCatalogVendorId** | Pointer to **string** | Vendor identifier for the catalog | [optional] 
**ServerFirmwareCatalogVendorUrl** | Pointer to **string** | Vendor URL for the firmware catalog | [optional] 
**ServerFirmwareCatalogVendorReleaseTimestamp** | Pointer to **string** | Vendor release timestamp for the catalog | [optional] 
**ServerFirmwareCatalogUpdateType** | **string** | Catalog update type: online or offline | [default to "online"]
**ServerFirmwareCatalogMetalsoftServerTypesSupportedJson** | Pointer to **map[string]interface{}** | List of supported Metalsoft server types for which firmware binaries are available | [optional] 
**ServerFirmwareCatalogVendorServerTypesSupportedJson** | Pointer to **map[string]interface{}** | List of supported server types for which firmware binaries are available | [optional] 
**ServerFirmwareCatalogVendorConfigurationJson** | Pointer to **map[string]interface{}** | Vendor configuration in JSON format | [optional] 
**ServerFirmwareCatalogCreatedTimestamp** | **string** | Timestamp when the catalog was created | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewFirmwareCatalog

`func NewFirmwareCatalog(serverFirmwareCatalogId float32, serverFirmwareCatalogName string, serverFirmwareCatalogVendor string, serverFirmwareCatalogUpdateType string, serverFirmwareCatalogCreatedTimestamp string, links map[string]interface{}, ) *FirmwareCatalog`

NewFirmwareCatalog instantiates a new FirmwareCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareCatalogWithDefaults

`func NewFirmwareCatalogWithDefaults() *FirmwareCatalog`

NewFirmwareCatalogWithDefaults instantiates a new FirmwareCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareCatalogId

`func (o *FirmwareCatalog) GetServerFirmwareCatalogId() float32`

GetServerFirmwareCatalogId returns the ServerFirmwareCatalogId field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogIdOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogIdOk() (*float32, bool)`

GetServerFirmwareCatalogIdOk returns a tuple with the ServerFirmwareCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogId

`func (o *FirmwareCatalog) SetServerFirmwareCatalogId(v float32)`

SetServerFirmwareCatalogId sets ServerFirmwareCatalogId field to given value.


### GetServerFirmwareCatalogName

`func (o *FirmwareCatalog) GetServerFirmwareCatalogName() string`

GetServerFirmwareCatalogName returns the ServerFirmwareCatalogName field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogNameOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogNameOk() (*string, bool)`

GetServerFirmwareCatalogNameOk returns a tuple with the ServerFirmwareCatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogName

`func (o *FirmwareCatalog) SetServerFirmwareCatalogName(v string)`

SetServerFirmwareCatalogName sets ServerFirmwareCatalogName field to given value.


### GetServerFirmwareCatalogDescription

`func (o *FirmwareCatalog) GetServerFirmwareCatalogDescription() string`

GetServerFirmwareCatalogDescription returns the ServerFirmwareCatalogDescription field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogDescriptionOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogDescriptionOk() (*string, bool)`

GetServerFirmwareCatalogDescriptionOk returns a tuple with the ServerFirmwareCatalogDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogDescription

`func (o *FirmwareCatalog) SetServerFirmwareCatalogDescription(v string)`

SetServerFirmwareCatalogDescription sets ServerFirmwareCatalogDescription field to given value.

### HasServerFirmwareCatalogDescription

`func (o *FirmwareCatalog) HasServerFirmwareCatalogDescription() bool`

HasServerFirmwareCatalogDescription returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendor

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendor() string`

GetServerFirmwareCatalogVendor returns the ServerFirmwareCatalogVendor field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorOk() (*string, bool)`

GetServerFirmwareCatalogVendorOk returns a tuple with the ServerFirmwareCatalogVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendor

`func (o *FirmwareCatalog) SetServerFirmwareCatalogVendor(v string)`

SetServerFirmwareCatalogVendor sets ServerFirmwareCatalogVendor field to given value.


### GetServerFirmwareCatalogVendorId

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorId() string`

GetServerFirmwareCatalogVendorId returns the ServerFirmwareCatalogVendorId field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorIdOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorIdOk() (*string, bool)`

GetServerFirmwareCatalogVendorIdOk returns a tuple with the ServerFirmwareCatalogVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorId

`func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorId(v string)`

SetServerFirmwareCatalogVendorId sets ServerFirmwareCatalogVendorId field to given value.

### HasServerFirmwareCatalogVendorId

`func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorId() bool`

HasServerFirmwareCatalogVendorId returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorUrl

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorUrl() string`

GetServerFirmwareCatalogVendorUrl returns the ServerFirmwareCatalogVendorUrl field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorUrlOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorUrlOk() (*string, bool)`

GetServerFirmwareCatalogVendorUrlOk returns a tuple with the ServerFirmwareCatalogVendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorUrl

`func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorUrl(v string)`

SetServerFirmwareCatalogVendorUrl sets ServerFirmwareCatalogVendorUrl field to given value.

### HasServerFirmwareCatalogVendorUrl

`func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorUrl() bool`

HasServerFirmwareCatalogVendorUrl returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestamp() string`

GetServerFirmwareCatalogVendorReleaseTimestamp returns the ServerFirmwareCatalogVendorReleaseTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorReleaseTimestampOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestampOk() (*string, bool)`

GetServerFirmwareCatalogVendorReleaseTimestampOk returns a tuple with the ServerFirmwareCatalogVendorReleaseTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorReleaseTimestamp

`func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorReleaseTimestamp(v string)`

SetServerFirmwareCatalogVendorReleaseTimestamp sets ServerFirmwareCatalogVendorReleaseTimestamp field to given value.

### HasServerFirmwareCatalogVendorReleaseTimestamp

`func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorReleaseTimestamp() bool`

HasServerFirmwareCatalogVendorReleaseTimestamp returns a boolean if a field has been set.

### GetServerFirmwareCatalogUpdateType

`func (o *FirmwareCatalog) GetServerFirmwareCatalogUpdateType() string`

GetServerFirmwareCatalogUpdateType returns the ServerFirmwareCatalogUpdateType field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogUpdateTypeOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogUpdateTypeOk() (*string, bool)`

GetServerFirmwareCatalogUpdateTypeOk returns a tuple with the ServerFirmwareCatalogUpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogUpdateType

`func (o *FirmwareCatalog) SetServerFirmwareCatalogUpdateType(v string)`

SetServerFirmwareCatalogUpdateType sets ServerFirmwareCatalogUpdateType field to given value.


### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *FirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson() map[string]interface{}`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk() (*map[string]interface{}, bool)`

GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *FirmwareCatalog) SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson(v map[string]interface{})`

SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson sets ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson

`func (o *FirmwareCatalog) HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson() bool`

HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJson() map[string]interface{}`

GetServerFirmwareCatalogVendorServerTypesSupportedJson returns the ServerFirmwareCatalogVendorServerTypesSupportedJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk() (*map[string]interface{}, bool)`

GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogVendorServerTypesSupportedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorServerTypesSupportedJson(v map[string]interface{})`

SetServerFirmwareCatalogVendorServerTypesSupportedJson sets ServerFirmwareCatalogVendorServerTypesSupportedJson field to given value.

### HasServerFirmwareCatalogVendorServerTypesSupportedJson

`func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorServerTypesSupportedJson() bool`

HasServerFirmwareCatalogVendorServerTypesSupportedJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogVendorConfigurationJson

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJson() map[string]interface{}`

GetServerFirmwareCatalogVendorConfigurationJson returns the ServerFirmwareCatalogVendorConfigurationJson field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogVendorConfigurationJsonOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJsonOk() (*map[string]interface{}, bool)`

GetServerFirmwareCatalogVendorConfigurationJsonOk returns a tuple with the ServerFirmwareCatalogVendorConfigurationJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogVendorConfigurationJson

`func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorConfigurationJson(v map[string]interface{})`

SetServerFirmwareCatalogVendorConfigurationJson sets ServerFirmwareCatalogVendorConfigurationJson field to given value.

### HasServerFirmwareCatalogVendorConfigurationJson

`func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorConfigurationJson() bool`

HasServerFirmwareCatalogVendorConfigurationJson returns a boolean if a field has been set.

### GetServerFirmwareCatalogCreatedTimestamp

`func (o *FirmwareCatalog) GetServerFirmwareCatalogCreatedTimestamp() string`

GetServerFirmwareCatalogCreatedTimestamp returns the ServerFirmwareCatalogCreatedTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogCreatedTimestampOk

`func (o *FirmwareCatalog) GetServerFirmwareCatalogCreatedTimestampOk() (*string, bool)`

GetServerFirmwareCatalogCreatedTimestampOk returns a tuple with the ServerFirmwareCatalogCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogCreatedTimestamp

`func (o *FirmwareCatalog) SetServerFirmwareCatalogCreatedTimestamp(v string)`

SetServerFirmwareCatalogCreatedTimestamp sets ServerFirmwareCatalogCreatedTimestamp field to given value.


### GetLinks

`func (o *FirmwareCatalog) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirmwareCatalog) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirmwareCatalog) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



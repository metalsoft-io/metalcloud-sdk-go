# FirmwareBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerFirmwareBaselineId** | **float32** | Unique identifier for the firmware baseline | 
**ServerFirmwareBaselineName** | **string** |  | 
**ServerFirmwareBaselineDescription** | Pointer to **string** |  | [optional] 
**ServerFirmwareBaselineCatalogJson** | Pointer to **string** |  | [optional] 
**ServerFirmwareBaselineLevel** | [**BaselineLevelType**](BaselineLevelType.md) |  | 
**ServerFirmwareBaselineLevelFilterJson** | **string** |  | 
**ServerFirmwareCatalogCreatedTimestamp** | **string** | Timestamp when the baseline was created | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewFirmwareBaseline

`func NewFirmwareBaseline(serverFirmwareBaselineId float32, serverFirmwareBaselineName string, serverFirmwareBaselineLevel BaselineLevelType, serverFirmwareBaselineLevelFilterJson string, serverFirmwareCatalogCreatedTimestamp string, links map[string]interface{}, ) *FirmwareBaseline`

NewFirmwareBaseline instantiates a new FirmwareBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaselineWithDefaults

`func NewFirmwareBaselineWithDefaults() *FirmwareBaseline`

NewFirmwareBaselineWithDefaults instantiates a new FirmwareBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBaselineId

`func (o *FirmwareBaseline) GetServerFirmwareBaselineId() float32`

GetServerFirmwareBaselineId returns the ServerFirmwareBaselineId field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineIdOk

`func (o *FirmwareBaseline) GetServerFirmwareBaselineIdOk() (*float32, bool)`

GetServerFirmwareBaselineIdOk returns a tuple with the ServerFirmwareBaselineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineId

`func (o *FirmwareBaseline) SetServerFirmwareBaselineId(v float32)`

SetServerFirmwareBaselineId sets ServerFirmwareBaselineId field to given value.


### GetServerFirmwareBaselineName

`func (o *FirmwareBaseline) GetServerFirmwareBaselineName() string`

GetServerFirmwareBaselineName returns the ServerFirmwareBaselineName field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineNameOk

`func (o *FirmwareBaseline) GetServerFirmwareBaselineNameOk() (*string, bool)`

GetServerFirmwareBaselineNameOk returns a tuple with the ServerFirmwareBaselineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineName

`func (o *FirmwareBaseline) SetServerFirmwareBaselineName(v string)`

SetServerFirmwareBaselineName sets ServerFirmwareBaselineName field to given value.


### GetServerFirmwareBaselineDescription

`func (o *FirmwareBaseline) GetServerFirmwareBaselineDescription() string`

GetServerFirmwareBaselineDescription returns the ServerFirmwareBaselineDescription field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineDescriptionOk

`func (o *FirmwareBaseline) GetServerFirmwareBaselineDescriptionOk() (*string, bool)`

GetServerFirmwareBaselineDescriptionOk returns a tuple with the ServerFirmwareBaselineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineDescription

`func (o *FirmwareBaseline) SetServerFirmwareBaselineDescription(v string)`

SetServerFirmwareBaselineDescription sets ServerFirmwareBaselineDescription field to given value.

### HasServerFirmwareBaselineDescription

`func (o *FirmwareBaseline) HasServerFirmwareBaselineDescription() bool`

HasServerFirmwareBaselineDescription returns a boolean if a field has been set.

### GetServerFirmwareBaselineCatalogJson

`func (o *FirmwareBaseline) GetServerFirmwareBaselineCatalogJson() string`

GetServerFirmwareBaselineCatalogJson returns the ServerFirmwareBaselineCatalogJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineCatalogJsonOk

`func (o *FirmwareBaseline) GetServerFirmwareBaselineCatalogJsonOk() (*string, bool)`

GetServerFirmwareBaselineCatalogJsonOk returns a tuple with the ServerFirmwareBaselineCatalogJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineCatalogJson

`func (o *FirmwareBaseline) SetServerFirmwareBaselineCatalogJson(v string)`

SetServerFirmwareBaselineCatalogJson sets ServerFirmwareBaselineCatalogJson field to given value.

### HasServerFirmwareBaselineCatalogJson

`func (o *FirmwareBaseline) HasServerFirmwareBaselineCatalogJson() bool`

HasServerFirmwareBaselineCatalogJson returns a boolean if a field has been set.

### GetServerFirmwareBaselineLevel

`func (o *FirmwareBaseline) GetServerFirmwareBaselineLevel() BaselineLevelType`

GetServerFirmwareBaselineLevel returns the ServerFirmwareBaselineLevel field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelOk

`func (o *FirmwareBaseline) GetServerFirmwareBaselineLevelOk() (*BaselineLevelType, bool)`

GetServerFirmwareBaselineLevelOk returns a tuple with the ServerFirmwareBaselineLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevel

`func (o *FirmwareBaseline) SetServerFirmwareBaselineLevel(v BaselineLevelType)`

SetServerFirmwareBaselineLevel sets ServerFirmwareBaselineLevel field to given value.


### GetServerFirmwareBaselineLevelFilterJson

`func (o *FirmwareBaseline) GetServerFirmwareBaselineLevelFilterJson() string`

GetServerFirmwareBaselineLevelFilterJson returns the ServerFirmwareBaselineLevelFilterJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelFilterJsonOk

`func (o *FirmwareBaseline) GetServerFirmwareBaselineLevelFilterJsonOk() (*string, bool)`

GetServerFirmwareBaselineLevelFilterJsonOk returns a tuple with the ServerFirmwareBaselineLevelFilterJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevelFilterJson

`func (o *FirmwareBaseline) SetServerFirmwareBaselineLevelFilterJson(v string)`

SetServerFirmwareBaselineLevelFilterJson sets ServerFirmwareBaselineLevelFilterJson field to given value.


### GetServerFirmwareCatalogCreatedTimestamp

`func (o *FirmwareBaseline) GetServerFirmwareCatalogCreatedTimestamp() string`

GetServerFirmwareCatalogCreatedTimestamp returns the ServerFirmwareCatalogCreatedTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogCreatedTimestampOk

`func (o *FirmwareBaseline) GetServerFirmwareCatalogCreatedTimestampOk() (*string, bool)`

GetServerFirmwareCatalogCreatedTimestampOk returns a tuple with the ServerFirmwareCatalogCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogCreatedTimestamp

`func (o *FirmwareBaseline) SetServerFirmwareCatalogCreatedTimestamp(v string)`

SetServerFirmwareCatalogCreatedTimestamp sets ServerFirmwareCatalogCreatedTimestamp field to given value.


### GetLinks

`func (o *FirmwareBaseline) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirmwareBaseline) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirmwareBaseline) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



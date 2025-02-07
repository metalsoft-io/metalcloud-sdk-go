# FirmwareBaselineDto

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

### NewFirmwareBaselineDto

`func NewFirmwareBaselineDto(serverFirmwareBaselineId float32, serverFirmwareBaselineName string, serverFirmwareBaselineLevel BaselineLevelType, serverFirmwareBaselineLevelFilterJson string, serverFirmwareCatalogCreatedTimestamp string, links map[string]interface{}, ) *FirmwareBaselineDto`

NewFirmwareBaselineDto instantiates a new FirmwareBaselineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaselineDtoWithDefaults

`func NewFirmwareBaselineDtoWithDefaults() *FirmwareBaselineDto`

NewFirmwareBaselineDtoWithDefaults instantiates a new FirmwareBaselineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerFirmwareBaselineId

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineId() float32`

GetServerFirmwareBaselineId returns the ServerFirmwareBaselineId field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineIdOk

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineIdOk() (*float32, bool)`

GetServerFirmwareBaselineIdOk returns a tuple with the ServerFirmwareBaselineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineId

`func (o *FirmwareBaselineDto) SetServerFirmwareBaselineId(v float32)`

SetServerFirmwareBaselineId sets ServerFirmwareBaselineId field to given value.


### GetServerFirmwareBaselineName

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineName() string`

GetServerFirmwareBaselineName returns the ServerFirmwareBaselineName field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineNameOk

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineNameOk() (*string, bool)`

GetServerFirmwareBaselineNameOk returns a tuple with the ServerFirmwareBaselineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineName

`func (o *FirmwareBaselineDto) SetServerFirmwareBaselineName(v string)`

SetServerFirmwareBaselineName sets ServerFirmwareBaselineName field to given value.


### GetServerFirmwareBaselineDescription

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineDescription() string`

GetServerFirmwareBaselineDescription returns the ServerFirmwareBaselineDescription field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineDescriptionOk

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineDescriptionOk() (*string, bool)`

GetServerFirmwareBaselineDescriptionOk returns a tuple with the ServerFirmwareBaselineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineDescription

`func (o *FirmwareBaselineDto) SetServerFirmwareBaselineDescription(v string)`

SetServerFirmwareBaselineDescription sets ServerFirmwareBaselineDescription field to given value.

### HasServerFirmwareBaselineDescription

`func (o *FirmwareBaselineDto) HasServerFirmwareBaselineDescription() bool`

HasServerFirmwareBaselineDescription returns a boolean if a field has been set.

### GetServerFirmwareBaselineCatalogJson

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineCatalogJson() string`

GetServerFirmwareBaselineCatalogJson returns the ServerFirmwareBaselineCatalogJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineCatalogJsonOk

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineCatalogJsonOk() (*string, bool)`

GetServerFirmwareBaselineCatalogJsonOk returns a tuple with the ServerFirmwareBaselineCatalogJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineCatalogJson

`func (o *FirmwareBaselineDto) SetServerFirmwareBaselineCatalogJson(v string)`

SetServerFirmwareBaselineCatalogJson sets ServerFirmwareBaselineCatalogJson field to given value.

### HasServerFirmwareBaselineCatalogJson

`func (o *FirmwareBaselineDto) HasServerFirmwareBaselineCatalogJson() bool`

HasServerFirmwareBaselineCatalogJson returns a boolean if a field has been set.

### GetServerFirmwareBaselineLevel

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineLevel() BaselineLevelType`

GetServerFirmwareBaselineLevel returns the ServerFirmwareBaselineLevel field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelOk

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineLevelOk() (*BaselineLevelType, bool)`

GetServerFirmwareBaselineLevelOk returns a tuple with the ServerFirmwareBaselineLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevel

`func (o *FirmwareBaselineDto) SetServerFirmwareBaselineLevel(v BaselineLevelType)`

SetServerFirmwareBaselineLevel sets ServerFirmwareBaselineLevel field to given value.


### GetServerFirmwareBaselineLevelFilterJson

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineLevelFilterJson() string`

GetServerFirmwareBaselineLevelFilterJson returns the ServerFirmwareBaselineLevelFilterJson field if non-nil, zero value otherwise.

### GetServerFirmwareBaselineLevelFilterJsonOk

`func (o *FirmwareBaselineDto) GetServerFirmwareBaselineLevelFilterJsonOk() (*string, bool)`

GetServerFirmwareBaselineLevelFilterJsonOk returns a tuple with the ServerFirmwareBaselineLevelFilterJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareBaselineLevelFilterJson

`func (o *FirmwareBaselineDto) SetServerFirmwareBaselineLevelFilterJson(v string)`

SetServerFirmwareBaselineLevelFilterJson sets ServerFirmwareBaselineLevelFilterJson field to given value.


### GetServerFirmwareCatalogCreatedTimestamp

`func (o *FirmwareBaselineDto) GetServerFirmwareCatalogCreatedTimestamp() string`

GetServerFirmwareCatalogCreatedTimestamp returns the ServerFirmwareCatalogCreatedTimestamp field if non-nil, zero value otherwise.

### GetServerFirmwareCatalogCreatedTimestampOk

`func (o *FirmwareBaselineDto) GetServerFirmwareCatalogCreatedTimestampOk() (*string, bool)`

GetServerFirmwareCatalogCreatedTimestampOk returns a tuple with the ServerFirmwareCatalogCreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareCatalogCreatedTimestamp

`func (o *FirmwareBaselineDto) SetServerFirmwareCatalogCreatedTimestamp(v string)`

SetServerFirmwareCatalogCreatedTimestamp sets ServerFirmwareCatalogCreatedTimestamp field to given value.


### GetLinks

`func (o *FirmwareBaselineDto) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirmwareBaselineDto) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirmwareBaselineDto) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



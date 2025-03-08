# FirmwareBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Unique identifier for the firmware baseline | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Catalog** | Pointer to **[]string** |  | [optional] 
**Level** | [**BaselineLevelType**](BaselineLevelType.md) |  | 
**LevelFilter** | **[]string** |  | 
**CreatedTimestamp** | **string** | Timestamp when the baseline was created | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewFirmwareBaseline

`func NewFirmwareBaseline(id float32, name string, level BaselineLevelType, levelFilter []string, createdTimestamp string, links map[string]interface{}, ) *FirmwareBaseline`

NewFirmwareBaseline instantiates a new FirmwareBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaselineWithDefaults

`func NewFirmwareBaselineWithDefaults() *FirmwareBaseline`

NewFirmwareBaselineWithDefaults instantiates a new FirmwareBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirmwareBaseline) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareBaseline) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareBaseline) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *FirmwareBaseline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareBaseline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareBaseline) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FirmwareBaseline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirmwareBaseline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirmwareBaseline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirmwareBaseline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCatalog

`func (o *FirmwareBaseline) GetCatalog() []string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareBaseline) GetCatalogOk() (*[]string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareBaseline) SetCatalog(v []string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareBaseline) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetLevel

`func (o *FirmwareBaseline) GetLevel() BaselineLevelType`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *FirmwareBaseline) GetLevelOk() (*BaselineLevelType, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *FirmwareBaseline) SetLevel(v BaselineLevelType)`

SetLevel sets Level field to given value.


### GetLevelFilter

`func (o *FirmwareBaseline) GetLevelFilter() []string`

GetLevelFilter returns the LevelFilter field if non-nil, zero value otherwise.

### GetLevelFilterOk

`func (o *FirmwareBaseline) GetLevelFilterOk() (*[]string, bool)`

GetLevelFilterOk returns a tuple with the LevelFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFilter

`func (o *FirmwareBaseline) SetLevelFilter(v []string)`

SetLevelFilter sets LevelFilter field to given value.


### GetCreatedTimestamp

`func (o *FirmwareBaseline) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *FirmwareBaseline) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *FirmwareBaseline) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


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



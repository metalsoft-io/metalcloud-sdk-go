# CreateFirmwareBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Catalog** | Pointer to **[]string** |  | [optional] 
**Level** | [**BaselineLevelType**](BaselineLevelType.md) |  | 
**LevelFilter** | **[]string** |  | 

## Methods

### NewCreateFirmwareBaseline

`func NewCreateFirmwareBaseline(name string, level BaselineLevelType, levelFilter []string, ) *CreateFirmwareBaseline`

NewCreateFirmwareBaseline instantiates a new CreateFirmwareBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirmwareBaselineWithDefaults

`func NewCreateFirmwareBaselineWithDefaults() *CreateFirmwareBaseline`

NewCreateFirmwareBaselineWithDefaults instantiates a new CreateFirmwareBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFirmwareBaseline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirmwareBaseline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirmwareBaseline) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateFirmwareBaseline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirmwareBaseline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirmwareBaseline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirmwareBaseline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCatalog

`func (o *CreateFirmwareBaseline) GetCatalog() []string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *CreateFirmwareBaseline) GetCatalogOk() (*[]string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *CreateFirmwareBaseline) SetCatalog(v []string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *CreateFirmwareBaseline) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetLevel

`func (o *CreateFirmwareBaseline) GetLevel() BaselineLevelType`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CreateFirmwareBaseline) GetLevelOk() (*BaselineLevelType, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CreateFirmwareBaseline) SetLevel(v BaselineLevelType)`

SetLevel sets Level field to given value.


### GetLevelFilter

`func (o *CreateFirmwareBaseline) GetLevelFilter() []string`

GetLevelFilter returns the LevelFilter field if non-nil, zero value otherwise.

### GetLevelFilterOk

`func (o *CreateFirmwareBaseline) GetLevelFilterOk() (*[]string, bool)`

GetLevelFilterOk returns a tuple with the LevelFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFilter

`func (o *CreateFirmwareBaseline) SetLevelFilter(v []string)`

SetLevelFilter sets LevelFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



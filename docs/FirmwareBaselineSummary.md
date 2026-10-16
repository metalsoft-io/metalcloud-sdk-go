# FirmwareBaselineSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the firmware baseline | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Catalog** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirmwareBaselineSummary

`func NewFirmwareBaselineSummary(id int64, name string, ) *FirmwareBaselineSummary`

NewFirmwareBaselineSummary instantiates a new FirmwareBaselineSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareBaselineSummaryWithDefaults

`func NewFirmwareBaselineSummaryWithDefaults() *FirmwareBaselineSummary`

NewFirmwareBaselineSummaryWithDefaults instantiates a new FirmwareBaselineSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirmwareBaselineSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareBaselineSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareBaselineSummary) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *FirmwareBaselineSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareBaselineSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareBaselineSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FirmwareBaselineSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirmwareBaselineSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirmwareBaselineSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirmwareBaselineSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCatalog

`func (o *FirmwareBaselineSummary) GetCatalog() []string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *FirmwareBaselineSummary) GetCatalogOk() (*[]string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *FirmwareBaselineSummary) SetCatalog(v []string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *FirmwareBaselineSummary) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



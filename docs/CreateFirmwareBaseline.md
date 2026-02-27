# CreateFirmwareBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Catalog** | Pointer to **[]string** |  | [optional] 
**MinimumVersions** | Pointer to [**[]FirmwareMinimumVersion**](FirmwareMinimumVersion.md) | Array of minimum firmware version requirements for this baseline | [optional] 

## Methods

### NewCreateFirmwareBaseline

`func NewCreateFirmwareBaseline(name string, ) *CreateFirmwareBaseline`

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

### GetMinimumVersions

`func (o *CreateFirmwareBaseline) GetMinimumVersions() []FirmwareMinimumVersion`

GetMinimumVersions returns the MinimumVersions field if non-nil, zero value otherwise.

### GetMinimumVersionsOk

`func (o *CreateFirmwareBaseline) GetMinimumVersionsOk() (*[]FirmwareMinimumVersion, bool)`

GetMinimumVersionsOk returns a tuple with the MinimumVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersions

`func (o *CreateFirmwareBaseline) SetMinimumVersions(v []FirmwareMinimumVersion)`

SetMinimumVersions sets MinimumVersions field to given value.

### HasMinimumVersions

`func (o *CreateFirmwareBaseline) HasMinimumVersions() bool`

HasMinimumVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateFirmwareBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Catalog** | Pointer to **[]string** |  | [optional] 
**Level** | [**BaselineLevelType**](BaselineLevelType.md) | Level type: datacenter, serverType, osTemplate for infrastructure-based baselines, or vendorModel for hardware-based baselines | 
**LevelFilter** | **[]string** | Array of filter values. For datacenter/serverType/osTemplate: array of strings (e.g., [\&quot;datacenter1\&quot;, \&quot;datacenter2\&quot;]). For vendorModel: array of objects with vendor and model properties (e.g., [{ vendor: \&quot;Dell\&quot;, model: \&quot;R740\&quot; }, { vendor: \&quot;HPE\&quot;, model: \&quot;*\&quot; }]). Use \&quot;*\&quot; for wildcards. | 
**MinimumVersions** | Pointer to [**[]FirmwareMinimumVersion**](FirmwareMinimumVersion.md) | Array of minimum firmware version requirements for this baseline | [optional] 
**ApplyOnServerRegistration** | Pointer to **bool** | Whether to apply this baseline during server registration process | [optional] 

## Methods

### NewUpdateFirmwareBaseline

`func NewUpdateFirmwareBaseline(name string, level BaselineLevelType, levelFilter []string, ) *UpdateFirmwareBaseline`

NewUpdateFirmwareBaseline instantiates a new UpdateFirmwareBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFirmwareBaselineWithDefaults

`func NewUpdateFirmwareBaselineWithDefaults() *UpdateFirmwareBaseline`

NewUpdateFirmwareBaselineWithDefaults instantiates a new UpdateFirmwareBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFirmwareBaseline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFirmwareBaseline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFirmwareBaseline) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateFirmwareBaseline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFirmwareBaseline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFirmwareBaseline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFirmwareBaseline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCatalog

`func (o *UpdateFirmwareBaseline) GetCatalog() []string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *UpdateFirmwareBaseline) GetCatalogOk() (*[]string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *UpdateFirmwareBaseline) SetCatalog(v []string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *UpdateFirmwareBaseline) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetLevel

`func (o *UpdateFirmwareBaseline) GetLevel() BaselineLevelType`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UpdateFirmwareBaseline) GetLevelOk() (*BaselineLevelType, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UpdateFirmwareBaseline) SetLevel(v BaselineLevelType)`

SetLevel sets Level field to given value.


### GetLevelFilter

`func (o *UpdateFirmwareBaseline) GetLevelFilter() []string`

GetLevelFilter returns the LevelFilter field if non-nil, zero value otherwise.

### GetLevelFilterOk

`func (o *UpdateFirmwareBaseline) GetLevelFilterOk() (*[]string, bool)`

GetLevelFilterOk returns a tuple with the LevelFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFilter

`func (o *UpdateFirmwareBaseline) SetLevelFilter(v []string)`

SetLevelFilter sets LevelFilter field to given value.


### GetMinimumVersions

`func (o *UpdateFirmwareBaseline) GetMinimumVersions() []FirmwareMinimumVersion`

GetMinimumVersions returns the MinimumVersions field if non-nil, zero value otherwise.

### GetMinimumVersionsOk

`func (o *UpdateFirmwareBaseline) GetMinimumVersionsOk() (*[]FirmwareMinimumVersion, bool)`

GetMinimumVersionsOk returns a tuple with the MinimumVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersions

`func (o *UpdateFirmwareBaseline) SetMinimumVersions(v []FirmwareMinimumVersion)`

SetMinimumVersions sets MinimumVersions field to given value.

### HasMinimumVersions

`func (o *UpdateFirmwareBaseline) HasMinimumVersions() bool`

HasMinimumVersions returns a boolean if a field has been set.

### GetApplyOnServerRegistration

`func (o *UpdateFirmwareBaseline) GetApplyOnServerRegistration() bool`

GetApplyOnServerRegistration returns the ApplyOnServerRegistration field if non-nil, zero value otherwise.

### GetApplyOnServerRegistrationOk

`func (o *UpdateFirmwareBaseline) GetApplyOnServerRegistrationOk() (*bool, bool)`

GetApplyOnServerRegistrationOk returns a tuple with the ApplyOnServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyOnServerRegistration

`func (o *UpdateFirmwareBaseline) SetApplyOnServerRegistration(v bool)`

SetApplyOnServerRegistration sets ApplyOnServerRegistration field to given value.

### HasApplyOnServerRegistration

`func (o *UpdateFirmwareBaseline) HasApplyOnServerRegistration() bool`

HasApplyOnServerRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



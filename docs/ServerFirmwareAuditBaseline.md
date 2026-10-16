# ServerFirmwareAuditBaseline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The id of the firmware baseline used for the audit. | 
**Name** | **string** | The name of the firmware baseline. | 
**Description** | Pointer to **string** | The description of the firmware baseline. | [optional] 
**Catalog** | Pointer to **[]string** | The ids of the firmware catalogs the baseline draws its binaries from. | [optional] 
**MinimumVersions** | Pointer to [**[]FirmwareMinimumVersion**](FirmwareMinimumVersion.md) | The minimum component versions the baseline requires. | [optional] 

## Methods

### NewServerFirmwareAuditBaseline

`func NewServerFirmwareAuditBaseline(id int64, name string, ) *ServerFirmwareAuditBaseline`

NewServerFirmwareAuditBaseline instantiates a new ServerFirmwareAuditBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareAuditBaselineWithDefaults

`func NewServerFirmwareAuditBaselineWithDefaults() *ServerFirmwareAuditBaseline`

NewServerFirmwareAuditBaselineWithDefaults instantiates a new ServerFirmwareAuditBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerFirmwareAuditBaseline) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerFirmwareAuditBaseline) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerFirmwareAuditBaseline) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ServerFirmwareAuditBaseline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerFirmwareAuditBaseline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerFirmwareAuditBaseline) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServerFirmwareAuditBaseline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerFirmwareAuditBaseline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerFirmwareAuditBaseline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerFirmwareAuditBaseline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCatalog

`func (o *ServerFirmwareAuditBaseline) GetCatalog() []string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ServerFirmwareAuditBaseline) GetCatalogOk() (*[]string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ServerFirmwareAuditBaseline) SetCatalog(v []string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ServerFirmwareAuditBaseline) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetMinimumVersions

`func (o *ServerFirmwareAuditBaseline) GetMinimumVersions() []FirmwareMinimumVersion`

GetMinimumVersions returns the MinimumVersions field if non-nil, zero value otherwise.

### GetMinimumVersionsOk

`func (o *ServerFirmwareAuditBaseline) GetMinimumVersionsOk() (*[]FirmwareMinimumVersion, bool)`

GetMinimumVersionsOk returns a tuple with the MinimumVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersions

`func (o *ServerFirmwareAuditBaseline) SetMinimumVersions(v []FirmwareMinimumVersion)`

SetMinimumVersions sets MinimumVersions field to given value.

### HasMinimumVersions

`func (o *ServerFirmwareAuditBaseline) HasMinimumVersions() bool`

HasMinimumVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



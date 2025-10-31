# VMPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM Pool ID | 
**SiteId** | **float32** | Id of the site for the VM | 
**DatacenterName** | **string** | Datacenter of the VM | 
**ManagementHost** | **string** | Host of the VM Pool | 
**ManagementPort** | **float32** | Port of the VM Pool | 
**Name** | **string** | Name of the VM Pool | 
**Description** | Pointer to **string** | Description of the VM Pool | [optional] 
**Type** | **string** | Type of the VM Pool | 
**Username** | Pointer to **string** | Username of the VM Pool | [optional] 
**Status** | **string** | Status of the VM Pool | 
**ExternalIdentifier** | Pointer to **string** | External identifier of the VM Pool | [optional] 
**InMaintenance** | Pointer to **float32** | Flag to indicate if the VM Pool is in maintenance mode. 1 for true, 0 for false. Default is 0. | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**CreatedTimestamp** | **string** | Timestamp when the VM Pool was created | 
**UpdatedTimestamp** | **string** | Timestamp when the VM Pool was updated | 
**Tags** | Pointer to **[]string** | Tags for the VM Pool. | [optional] 
**Options** | Pointer to [**VMPoolOptions**](VMPoolOptions.md) | Options for the vm pool | [optional] 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMPool

`func NewVMPool(id float32, siteId float32, datacenterName string, managementHost string, managementPort float32, name string, type_ string, status string, createdTimestamp string, updatedTimestamp string, links map[string]interface{}, ) *VMPool`

NewVMPool instantiates a new VMPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolWithDefaults

`func NewVMPoolWithDefaults() *VMPool`

NewVMPoolWithDefaults instantiates a new VMPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPool) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPool) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPool) SetId(v float32)`

SetId sets Id field to given value.


### GetSiteId

`func (o *VMPool) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *VMPool) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *VMPool) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetDatacenterName

`func (o *VMPool) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *VMPool) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *VMPool) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetManagementHost

`func (o *VMPool) GetManagementHost() string`

GetManagementHost returns the ManagementHost field if non-nil, zero value otherwise.

### GetManagementHostOk

`func (o *VMPool) GetManagementHostOk() (*string, bool)`

GetManagementHostOk returns a tuple with the ManagementHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHost

`func (o *VMPool) SetManagementHost(v string)`

SetManagementHost sets ManagementHost field to given value.


### GetManagementPort

`func (o *VMPool) GetManagementPort() float32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *VMPool) GetManagementPortOk() (*float32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *VMPool) SetManagementPort(v float32)`

SetManagementPort sets ManagementPort field to given value.


### GetName

`func (o *VMPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMPool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VMPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VMPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VMPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VMPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *VMPool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VMPool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VMPool) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *VMPool) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VMPool) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VMPool) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VMPool) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetStatus

`func (o *VMPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VMPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VMPool) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExternalIdentifier

`func (o *VMPool) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *VMPool) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *VMPool) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *VMPool) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### GetInMaintenance

`func (o *VMPool) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *VMPool) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *VMPool) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *VMPool) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetIsExperimental

`func (o *VMPool) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *VMPool) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *VMPool) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *VMPool) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *VMPool) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VMPool) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VMPool) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *VMPool) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMPool) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMPool) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetTags

`func (o *VMPool) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VMPool) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VMPool) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VMPool) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOptions

`func (o *VMPool) GetOptions() VMPoolOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *VMPool) GetOptionsOk() (*VMPoolOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *VMPool) SetOptions(v VMPoolOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *VMPool) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetLinks

`func (o *VMPool) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMPool) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMPool) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



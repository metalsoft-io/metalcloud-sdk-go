# ServerComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The unique identifier of the server component. | 
**ExternalId** | Pointer to **string** | The external identifier of the server component. | [optional] 
**ServerId** | Pointer to **float32** | The unique identifier of the server. | [optional] 
**Name** | Pointer to **string** | The name of the server component. | [optional] 
**FirmwareTargetVersion** | Pointer to **string** | The target firmware version of the server component. | [optional] 
**FirmwareUpdateAvailableVersions** | Pointer to **map[string]interface{}** | The available firmware versions of the server component. | [optional] 
**FirmwareUpdateRequiredFixes** | Pointer to **map[string]interface{}** | The required fixes for the firmware upgrade of the server component. | [optional] 
**FirmwareUpdateable** | **float32** | The flag indicating if the server component has updateable firmware. | 
**FirmwareVersion** | Pointer to **string** | The current firmware version of the server component. | [optional] 
**FirmwareInfo** | Pointer to **map[string]interface{}** | The firmware information of the server component. | [optional] 
**FirmwareUpgradeNeedsConfirmation** | **float32** | The flag indicating if the firmware upgrade of the server component requires confirmation. | 
**FirmwareUpdateTimestamp** | Pointer to **string** | The timestamp of the last firmware upgrade of the server component. | [optional] 
**FirmwareScheduledTimestamp** | Pointer to **string** | The timestamp of the scheduled firmware upgrade of the server component. | [optional] 
**FirmwareStatus** | **string** | The status of the firmware upgrade of the server component. | 
**Type** | Pointer to **string** | The type of the server component. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerComponent

`func NewServerComponent(id float32, firmwareUpdateable float32, firmwareUpgradeNeedsConfirmation float32, firmwareStatus string, ) *ServerComponent`

NewServerComponent instantiates a new ServerComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerComponentWithDefaults

`func NewServerComponentWithDefaults() *ServerComponent`

NewServerComponentWithDefaults instantiates a new ServerComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerComponent) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerComponent) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerComponent) SetId(v float32)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ServerComponent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ServerComponent) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ServerComponent) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ServerComponent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetServerId

`func (o *ServerComponent) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerComponent) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerComponent) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ServerComponent) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetName

`func (o *ServerComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFirmwareTargetVersion

`func (o *ServerComponent) GetFirmwareTargetVersion() string`

GetFirmwareTargetVersion returns the FirmwareTargetVersion field if non-nil, zero value otherwise.

### GetFirmwareTargetVersionOk

`func (o *ServerComponent) GetFirmwareTargetVersionOk() (*string, bool)`

GetFirmwareTargetVersionOk returns a tuple with the FirmwareTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareTargetVersion

`func (o *ServerComponent) SetFirmwareTargetVersion(v string)`

SetFirmwareTargetVersion sets FirmwareTargetVersion field to given value.

### HasFirmwareTargetVersion

`func (o *ServerComponent) HasFirmwareTargetVersion() bool`

HasFirmwareTargetVersion returns a boolean if a field has been set.

### GetFirmwareUpdateAvailableVersions

`func (o *ServerComponent) GetFirmwareUpdateAvailableVersions() map[string]interface{}`

GetFirmwareUpdateAvailableVersions returns the FirmwareUpdateAvailableVersions field if non-nil, zero value otherwise.

### GetFirmwareUpdateAvailableVersionsOk

`func (o *ServerComponent) GetFirmwareUpdateAvailableVersionsOk() (*map[string]interface{}, bool)`

GetFirmwareUpdateAvailableVersionsOk returns a tuple with the FirmwareUpdateAvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpdateAvailableVersions

`func (o *ServerComponent) SetFirmwareUpdateAvailableVersions(v map[string]interface{})`

SetFirmwareUpdateAvailableVersions sets FirmwareUpdateAvailableVersions field to given value.

### HasFirmwareUpdateAvailableVersions

`func (o *ServerComponent) HasFirmwareUpdateAvailableVersions() bool`

HasFirmwareUpdateAvailableVersions returns a boolean if a field has been set.

### GetFirmwareUpdateRequiredFixes

`func (o *ServerComponent) GetFirmwareUpdateRequiredFixes() map[string]interface{}`

GetFirmwareUpdateRequiredFixes returns the FirmwareUpdateRequiredFixes field if non-nil, zero value otherwise.

### GetFirmwareUpdateRequiredFixesOk

`func (o *ServerComponent) GetFirmwareUpdateRequiredFixesOk() (*map[string]interface{}, bool)`

GetFirmwareUpdateRequiredFixesOk returns a tuple with the FirmwareUpdateRequiredFixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpdateRequiredFixes

`func (o *ServerComponent) SetFirmwareUpdateRequiredFixes(v map[string]interface{})`

SetFirmwareUpdateRequiredFixes sets FirmwareUpdateRequiredFixes field to given value.

### HasFirmwareUpdateRequiredFixes

`func (o *ServerComponent) HasFirmwareUpdateRequiredFixes() bool`

HasFirmwareUpdateRequiredFixes returns a boolean if a field has been set.

### GetFirmwareUpdateable

`func (o *ServerComponent) GetFirmwareUpdateable() float32`

GetFirmwareUpdateable returns the FirmwareUpdateable field if non-nil, zero value otherwise.

### GetFirmwareUpdateableOk

`func (o *ServerComponent) GetFirmwareUpdateableOk() (*float32, bool)`

GetFirmwareUpdateableOk returns a tuple with the FirmwareUpdateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpdateable

`func (o *ServerComponent) SetFirmwareUpdateable(v float32)`

SetFirmwareUpdateable sets FirmwareUpdateable field to given value.


### GetFirmwareVersion

`func (o *ServerComponent) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *ServerComponent) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *ServerComponent) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *ServerComponent) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.

### GetFirmwareInfo

`func (o *ServerComponent) GetFirmwareInfo() map[string]interface{}`

GetFirmwareInfo returns the FirmwareInfo field if non-nil, zero value otherwise.

### GetFirmwareInfoOk

`func (o *ServerComponent) GetFirmwareInfoOk() (*map[string]interface{}, bool)`

GetFirmwareInfoOk returns a tuple with the FirmwareInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareInfo

`func (o *ServerComponent) SetFirmwareInfo(v map[string]interface{})`

SetFirmwareInfo sets FirmwareInfo field to given value.

### HasFirmwareInfo

`func (o *ServerComponent) HasFirmwareInfo() bool`

HasFirmwareInfo returns a boolean if a field has been set.

### GetFirmwareUpgradeNeedsConfirmation

`func (o *ServerComponent) GetFirmwareUpgradeNeedsConfirmation() float32`

GetFirmwareUpgradeNeedsConfirmation returns the FirmwareUpgradeNeedsConfirmation field if non-nil, zero value otherwise.

### GetFirmwareUpgradeNeedsConfirmationOk

`func (o *ServerComponent) GetFirmwareUpgradeNeedsConfirmationOk() (*float32, bool)`

GetFirmwareUpgradeNeedsConfirmationOk returns a tuple with the FirmwareUpgradeNeedsConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpgradeNeedsConfirmation

`func (o *ServerComponent) SetFirmwareUpgradeNeedsConfirmation(v float32)`

SetFirmwareUpgradeNeedsConfirmation sets FirmwareUpgradeNeedsConfirmation field to given value.


### GetFirmwareUpdateTimestamp

`func (o *ServerComponent) GetFirmwareUpdateTimestamp() string`

GetFirmwareUpdateTimestamp returns the FirmwareUpdateTimestamp field if non-nil, zero value otherwise.

### GetFirmwareUpdateTimestampOk

`func (o *ServerComponent) GetFirmwareUpdateTimestampOk() (*string, bool)`

GetFirmwareUpdateTimestampOk returns a tuple with the FirmwareUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpdateTimestamp

`func (o *ServerComponent) SetFirmwareUpdateTimestamp(v string)`

SetFirmwareUpdateTimestamp sets FirmwareUpdateTimestamp field to given value.

### HasFirmwareUpdateTimestamp

`func (o *ServerComponent) HasFirmwareUpdateTimestamp() bool`

HasFirmwareUpdateTimestamp returns a boolean if a field has been set.

### GetFirmwareScheduledTimestamp

`func (o *ServerComponent) GetFirmwareScheduledTimestamp() string`

GetFirmwareScheduledTimestamp returns the FirmwareScheduledTimestamp field if non-nil, zero value otherwise.

### GetFirmwareScheduledTimestampOk

`func (o *ServerComponent) GetFirmwareScheduledTimestampOk() (*string, bool)`

GetFirmwareScheduledTimestampOk returns a tuple with the FirmwareScheduledTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareScheduledTimestamp

`func (o *ServerComponent) SetFirmwareScheduledTimestamp(v string)`

SetFirmwareScheduledTimestamp sets FirmwareScheduledTimestamp field to given value.

### HasFirmwareScheduledTimestamp

`func (o *ServerComponent) HasFirmwareScheduledTimestamp() bool`

HasFirmwareScheduledTimestamp returns a boolean if a field has been set.

### GetFirmwareStatus

`func (o *ServerComponent) GetFirmwareStatus() string`

GetFirmwareStatus returns the FirmwareStatus field if non-nil, zero value otherwise.

### GetFirmwareStatusOk

`func (o *ServerComponent) GetFirmwareStatusOk() (*string, bool)`

GetFirmwareStatusOk returns a tuple with the FirmwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareStatus

`func (o *ServerComponent) SetFirmwareStatus(v string)`

SetFirmwareStatus sets FirmwareStatus field to given value.


### GetType

`func (o *ServerComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerComponent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ServerComponent) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerComponent) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerComponent) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerComponent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



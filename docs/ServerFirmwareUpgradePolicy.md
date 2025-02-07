# ServerFirmwareUpgradePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The unique identifier of the firmware upgrade policy. | 
**Label** | **string** | The label of the firmware upgrade policy. | 
**UserIdOwner** | Pointer to **float32** | The unique identifier of the user who owns the firmware upgrade policy. | [optional] 
**CreatedTimestamp** | **string** | Timestamp of the Server Firmware Upgrade Policy creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the Server Firmware Upgrade Policy last update. | 
**Status** | **string** | The status of the firmware upgrade policy. | 
**Action** | **string** | The action of the firmware upgrade policy. | 
**Rules** | Pointer to [**[]ServerFirmwareUpgradePolicyRule**](ServerFirmwareUpgradePolicyRule.md) | The rules of the firmware upgrade policy. | [optional] 
**InstanceArrayIds** | Pointer to **[]float32** | The unique identifiers of the instance arrays associated with the firmware upgrade policy. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerFirmwareUpgradePolicy

`func NewServerFirmwareUpgradePolicy(id float32, label string, createdTimestamp string, updatedTimestamp string, status string, action string, ) *ServerFirmwareUpgradePolicy`

NewServerFirmwareUpgradePolicy instantiates a new ServerFirmwareUpgradePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareUpgradePolicyWithDefaults

`func NewServerFirmwareUpgradePolicyWithDefaults() *ServerFirmwareUpgradePolicy`

NewServerFirmwareUpgradePolicyWithDefaults instantiates a new ServerFirmwareUpgradePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerFirmwareUpgradePolicy) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerFirmwareUpgradePolicy) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerFirmwareUpgradePolicy) SetId(v float32)`

SetId sets Id field to given value.


### GetLabel

`func (o *ServerFirmwareUpgradePolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerFirmwareUpgradePolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerFirmwareUpgradePolicy) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUserIdOwner

`func (o *ServerFirmwareUpgradePolicy) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *ServerFirmwareUpgradePolicy) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *ServerFirmwareUpgradePolicy) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *ServerFirmwareUpgradePolicy) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ServerFirmwareUpgradePolicy) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerFirmwareUpgradePolicy) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerFirmwareUpgradePolicy) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerFirmwareUpgradePolicy) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerFirmwareUpgradePolicy) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerFirmwareUpgradePolicy) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetStatus

`func (o *ServerFirmwareUpgradePolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerFirmwareUpgradePolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerFirmwareUpgradePolicy) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAction

`func (o *ServerFirmwareUpgradePolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServerFirmwareUpgradePolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServerFirmwareUpgradePolicy) SetAction(v string)`

SetAction sets Action field to given value.


### GetRules

`func (o *ServerFirmwareUpgradePolicy) GetRules() []ServerFirmwareUpgradePolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ServerFirmwareUpgradePolicy) GetRulesOk() (*[]ServerFirmwareUpgradePolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ServerFirmwareUpgradePolicy) SetRules(v []ServerFirmwareUpgradePolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ServerFirmwareUpgradePolicy) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetInstanceArrayIds

`func (o *ServerFirmwareUpgradePolicy) GetInstanceArrayIds() []float32`

GetInstanceArrayIds returns the InstanceArrayIds field if non-nil, zero value otherwise.

### GetInstanceArrayIdsOk

`func (o *ServerFirmwareUpgradePolicy) GetInstanceArrayIdsOk() (*[]float32, bool)`

GetInstanceArrayIdsOk returns a tuple with the InstanceArrayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrayIds

`func (o *ServerFirmwareUpgradePolicy) SetInstanceArrayIds(v []float32)`

SetInstanceArrayIds sets InstanceArrayIds field to given value.

### HasInstanceArrayIds

`func (o *ServerFirmwareUpgradePolicy) HasInstanceArrayIds() bool`

HasInstanceArrayIds returns a boolean if a field has been set.

### GetLinks

`func (o *ServerFirmwareUpgradePolicy) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerFirmwareUpgradePolicy) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerFirmwareUpgradePolicy) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerFirmwareUpgradePolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



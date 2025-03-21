# UpdateServerFirmwareUpgradePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]ServerFirmwareUpgradePolicyRule**](ServerFirmwareUpgradePolicyRule.md) | The rules of the firmware upgrade policy. | [optional] 
**ServerInstanceGroupIds** | Pointer to **[]float32** | The unique identifiers of the server instance groups associated with the firmware upgrade policy. | [optional] 
**Label** | Pointer to **string** | The label of the firmware upgrade policy. | [optional] 
**Status** | Pointer to **string** | The status of the firmware upgrade policy. | [optional] 
**Action** | Pointer to **string** | The action of the firmware upgrade policy. | [optional] 

## Methods

### NewUpdateServerFirmwareUpgradePolicy

`func NewUpdateServerFirmwareUpgradePolicy() *UpdateServerFirmwareUpgradePolicy`

NewUpdateServerFirmwareUpgradePolicy instantiates a new UpdateServerFirmwareUpgradePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerFirmwareUpgradePolicyWithDefaults

`func NewUpdateServerFirmwareUpgradePolicyWithDefaults() *UpdateServerFirmwareUpgradePolicy`

NewUpdateServerFirmwareUpgradePolicyWithDefaults instantiates a new UpdateServerFirmwareUpgradePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *UpdateServerFirmwareUpgradePolicy) GetRules() []ServerFirmwareUpgradePolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateServerFirmwareUpgradePolicy) GetRulesOk() (*[]ServerFirmwareUpgradePolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateServerFirmwareUpgradePolicy) SetRules(v []ServerFirmwareUpgradePolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateServerFirmwareUpgradePolicy) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetServerInstanceGroupIds

`func (o *UpdateServerFirmwareUpgradePolicy) GetServerInstanceGroupIds() []float32`

GetServerInstanceGroupIds returns the ServerInstanceGroupIds field if non-nil, zero value otherwise.

### GetServerInstanceGroupIdsOk

`func (o *UpdateServerFirmwareUpgradePolicy) GetServerInstanceGroupIdsOk() (*[]float32, bool)`

GetServerInstanceGroupIdsOk returns a tuple with the ServerInstanceGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroupIds

`func (o *UpdateServerFirmwareUpgradePolicy) SetServerInstanceGroupIds(v []float32)`

SetServerInstanceGroupIds sets ServerInstanceGroupIds field to given value.

### HasServerInstanceGroupIds

`func (o *UpdateServerFirmwareUpgradePolicy) HasServerInstanceGroupIds() bool`

HasServerInstanceGroupIds returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateServerFirmwareUpgradePolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateServerFirmwareUpgradePolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateServerFirmwareUpgradePolicy) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateServerFirmwareUpgradePolicy) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateServerFirmwareUpgradePolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateServerFirmwareUpgradePolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateServerFirmwareUpgradePolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateServerFirmwareUpgradePolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAction

`func (o *UpdateServerFirmwareUpgradePolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateServerFirmwareUpgradePolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateServerFirmwareUpgradePolicy) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UpdateServerFirmwareUpgradePolicy) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



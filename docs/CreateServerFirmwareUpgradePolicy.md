# CreateServerFirmwareUpgradePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The label of the firmware upgrade policy. | 
**UserIdOwner** | Pointer to **float32** | The unique identifier of the user who owns the firmware upgrade policy. | [optional] 
**Action** | **string** | The action of the firmware upgrade policy. | 
**Rules** | Pointer to [**[]ServerFirmwareUpgradePolicyRule**](ServerFirmwareUpgradePolicyRule.md) | The rules of the firmware upgrade policy. | [optional] 
**InstanceArrayIds** | Pointer to **[]float32** | The unique identifiers of the instance arrays associated with the firmware upgrade policy. | [optional] 

## Methods

### NewCreateServerFirmwareUpgradePolicy

`func NewCreateServerFirmwareUpgradePolicy(label string, action string, ) *CreateServerFirmwareUpgradePolicy`

NewCreateServerFirmwareUpgradePolicy instantiates a new CreateServerFirmwareUpgradePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerFirmwareUpgradePolicyWithDefaults

`func NewCreateServerFirmwareUpgradePolicyWithDefaults() *CreateServerFirmwareUpgradePolicy`

NewCreateServerFirmwareUpgradePolicyWithDefaults instantiates a new CreateServerFirmwareUpgradePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateServerFirmwareUpgradePolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateServerFirmwareUpgradePolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateServerFirmwareUpgradePolicy) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUserIdOwner

`func (o *CreateServerFirmwareUpgradePolicy) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *CreateServerFirmwareUpgradePolicy) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *CreateServerFirmwareUpgradePolicy) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *CreateServerFirmwareUpgradePolicy) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetAction

`func (o *CreateServerFirmwareUpgradePolicy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateServerFirmwareUpgradePolicy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateServerFirmwareUpgradePolicy) SetAction(v string)`

SetAction sets Action field to given value.


### GetRules

`func (o *CreateServerFirmwareUpgradePolicy) GetRules() []ServerFirmwareUpgradePolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateServerFirmwareUpgradePolicy) GetRulesOk() (*[]ServerFirmwareUpgradePolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateServerFirmwareUpgradePolicy) SetRules(v []ServerFirmwareUpgradePolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateServerFirmwareUpgradePolicy) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetInstanceArrayIds

`func (o *CreateServerFirmwareUpgradePolicy) GetInstanceArrayIds() []float32`

GetInstanceArrayIds returns the InstanceArrayIds field if non-nil, zero value otherwise.

### GetInstanceArrayIdsOk

`func (o *CreateServerFirmwareUpgradePolicy) GetInstanceArrayIdsOk() (*[]float32, bool)`

GetInstanceArrayIdsOk returns a tuple with the InstanceArrayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrayIds

`func (o *CreateServerFirmwareUpgradePolicy) SetInstanceArrayIds(v []float32)`

SetInstanceArrayIds sets InstanceArrayIds field to given value.

### HasInstanceArrayIds

`func (o *CreateServerFirmwareUpgradePolicy) HasInstanceArrayIds() bool`

HasInstanceArrayIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



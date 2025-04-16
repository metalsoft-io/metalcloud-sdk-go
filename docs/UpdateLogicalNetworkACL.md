# UpdateLogicalNetworkACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | Pointer to [**ACLType**](ACLType.md) | The type of the rule | [optional] 
**Direction** | Pointer to [**ACLDirection**](ACLDirection.md) | The direction of the rule | [optional] 
**Sequence** | Pointer to **int32** | The sequence number of the rule | [optional] 
**ForwardingAction** | Pointer to [**ACLForwardingAction**](ACLForwardingAction.md) | The forwarding action of the rule | [optional] 
**NetworkProtocol** | Pointer to **string** | The network protocol of the ACL rule | [optional] 
**SourceAddress** | Pointer to **string** | The source address of the rule if the rule type is IPv4, IPv6 | [optional] 
**DestinationAddress** | Pointer to **string** | The destination address of the rule if the rule type is IPv4, IPv6 | [optional] 
**SourcePort** | Pointer to **string** | The source port of the rule if the rule type is IPv4, IPv6, TCP or UDP | [optional] 
**DestinationPort** | Pointer to **string** | The destination port of the rule if the rule type is IPv4, IPv6 | [optional] 
**SourceMac** | Pointer to **string** | The source MAC address of the rule if the rule type is MAC | [optional] 
**DestinationMac** | Pointer to **string** | The destination MAC address of the rule if the rule type is MAC | [optional] 
**EnforcementPoint** | Pointer to [**ACLEnforcementPoint**](ACLEnforcementPoint.md) | The enforcement point of the rule | [optional] 

## Methods

### NewUpdateLogicalNetworkACL

`func NewUpdateLogicalNetworkACL() *UpdateLogicalNetworkACL`

NewUpdateLogicalNetworkACL instantiates a new UpdateLogicalNetworkACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLogicalNetworkACLWithDefaults

`func NewUpdateLogicalNetworkACLWithDefaults() *UpdateLogicalNetworkACL`

NewUpdateLogicalNetworkACLWithDefaults instantiates a new UpdateLogicalNetworkACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *UpdateLogicalNetworkACL) GetRuleType() ACLType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *UpdateLogicalNetworkACL) GetRuleTypeOk() (*ACLType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *UpdateLogicalNetworkACL) SetRuleType(v ACLType)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *UpdateLogicalNetworkACL) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetDirection

`func (o *UpdateLogicalNetworkACL) GetDirection() ACLDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *UpdateLogicalNetworkACL) GetDirectionOk() (*ACLDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *UpdateLogicalNetworkACL) SetDirection(v ACLDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *UpdateLogicalNetworkACL) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSequence

`func (o *UpdateLogicalNetworkACL) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *UpdateLogicalNetworkACL) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *UpdateLogicalNetworkACL) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *UpdateLogicalNetworkACL) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetForwardingAction

`func (o *UpdateLogicalNetworkACL) GetForwardingAction() ACLForwardingAction`

GetForwardingAction returns the ForwardingAction field if non-nil, zero value otherwise.

### GetForwardingActionOk

`func (o *UpdateLogicalNetworkACL) GetForwardingActionOk() (*ACLForwardingAction, bool)`

GetForwardingActionOk returns a tuple with the ForwardingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAction

`func (o *UpdateLogicalNetworkACL) SetForwardingAction(v ACLForwardingAction)`

SetForwardingAction sets ForwardingAction field to given value.

### HasForwardingAction

`func (o *UpdateLogicalNetworkACL) HasForwardingAction() bool`

HasForwardingAction returns a boolean if a field has been set.

### GetNetworkProtocol

`func (o *UpdateLogicalNetworkACL) GetNetworkProtocol() string`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *UpdateLogicalNetworkACL) GetNetworkProtocolOk() (*string, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *UpdateLogicalNetworkACL) SetNetworkProtocol(v string)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *UpdateLogicalNetworkACL) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetSourceAddress

`func (o *UpdateLogicalNetworkACL) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *UpdateLogicalNetworkACL) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *UpdateLogicalNetworkACL) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *UpdateLogicalNetworkACL) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *UpdateLogicalNetworkACL) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *UpdateLogicalNetworkACL) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *UpdateLogicalNetworkACL) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *UpdateLogicalNetworkACL) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetSourcePort

`func (o *UpdateLogicalNetworkACL) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *UpdateLogicalNetworkACL) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *UpdateLogicalNetworkACL) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *UpdateLogicalNetworkACL) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *UpdateLogicalNetworkACL) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *UpdateLogicalNetworkACL) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *UpdateLogicalNetworkACL) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *UpdateLogicalNetworkACL) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetSourceMac

`func (o *UpdateLogicalNetworkACL) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *UpdateLogicalNetworkACL) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *UpdateLogicalNetworkACL) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *UpdateLogicalNetworkACL) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetDestinationMac

`func (o *UpdateLogicalNetworkACL) GetDestinationMac() string`

GetDestinationMac returns the DestinationMac field if non-nil, zero value otherwise.

### GetDestinationMacOk

`func (o *UpdateLogicalNetworkACL) GetDestinationMacOk() (*string, bool)`

GetDestinationMacOk returns a tuple with the DestinationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMac

`func (o *UpdateLogicalNetworkACL) SetDestinationMac(v string)`

SetDestinationMac sets DestinationMac field to given value.

### HasDestinationMac

`func (o *UpdateLogicalNetworkACL) HasDestinationMac() bool`

HasDestinationMac returns a boolean if a field has been set.

### GetEnforcementPoint

`func (o *UpdateLogicalNetworkACL) GetEnforcementPoint() ACLEnforcementPoint`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *UpdateLogicalNetworkACL) GetEnforcementPointOk() (*ACLEnforcementPoint, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *UpdateLogicalNetworkACL) SetEnforcementPoint(v ACLEnforcementPoint)`

SetEnforcementPoint sets EnforcementPoint field to given value.

### HasEnforcementPoint

`func (o *UpdateLogicalNetworkACL) HasEnforcementPoint() bool`

HasEnforcementPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateLogicalNetworkACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | [**ACLType**](ACLType.md) | The type of the rule such as ipv4, ipv6, mac | 
**Direction** | [**ACLDirection**](ACLDirection.md) | The direction of the rule such as in, out | 
**Sequence** | **int32** | The sequence number of the rule important for the order evaluation of the rules | 
**ForwardingAction** | [**ACLForwardingAction**](ACLForwardingAction.md) | The forwarding action of the rule | 
**NetworkProtocol** | Pointer to **string** | The network protocol of the ACL rule | [optional] 
**SourceAddress** | Pointer to **string** | The source address of the rule if the rule type is IPv4, IPv6 | [optional] 
**DestinationAddress** | Pointer to **string** | The destination address of the rule if the rule type is IPv4, IPv6 | [optional] 
**SourcePort** | Pointer to **string** | The source port of the rule if the rule type is IPv4, IPv6, TCP or UDP | [optional] 
**DestinationPort** | Pointer to **string** | The destination port of the rule if the rule type is IPv4, IPv6 | [optional] 
**SourceMac** | Pointer to **string** | The source MAC address of the rule if the rule type is MAC | [optional] 
**DestinationMac** | Pointer to **string** | The destination MAC address of the rule if the rule type is MAC | [optional] 
**EnforcementPoint** | [**ACLEnforcementPoint**](ACLEnforcementPoint.md) | The enforcement point of the rule | 

## Methods

### NewCreateLogicalNetworkACL

`func NewCreateLogicalNetworkACL(ruleType ACLType, direction ACLDirection, sequence int32, forwardingAction ACLForwardingAction, enforcementPoint ACLEnforcementPoint, ) *CreateLogicalNetworkACL`

NewCreateLogicalNetworkACL instantiates a new CreateLogicalNetworkACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkACLWithDefaults

`func NewCreateLogicalNetworkACLWithDefaults() *CreateLogicalNetworkACL`

NewCreateLogicalNetworkACLWithDefaults instantiates a new CreateLogicalNetworkACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *CreateLogicalNetworkACL) GetRuleType() ACLType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *CreateLogicalNetworkACL) GetRuleTypeOk() (*ACLType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *CreateLogicalNetworkACL) SetRuleType(v ACLType)`

SetRuleType sets RuleType field to given value.


### GetDirection

`func (o *CreateLogicalNetworkACL) GetDirection() ACLDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateLogicalNetworkACL) GetDirectionOk() (*ACLDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateLogicalNetworkACL) SetDirection(v ACLDirection)`

SetDirection sets Direction field to given value.


### GetSequence

`func (o *CreateLogicalNetworkACL) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *CreateLogicalNetworkACL) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *CreateLogicalNetworkACL) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetForwardingAction

`func (o *CreateLogicalNetworkACL) GetForwardingAction() ACLForwardingAction`

GetForwardingAction returns the ForwardingAction field if non-nil, zero value otherwise.

### GetForwardingActionOk

`func (o *CreateLogicalNetworkACL) GetForwardingActionOk() (*ACLForwardingAction, bool)`

GetForwardingActionOk returns a tuple with the ForwardingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAction

`func (o *CreateLogicalNetworkACL) SetForwardingAction(v ACLForwardingAction)`

SetForwardingAction sets ForwardingAction field to given value.


### GetNetworkProtocol

`func (o *CreateLogicalNetworkACL) GetNetworkProtocol() string`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *CreateLogicalNetworkACL) GetNetworkProtocolOk() (*string, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *CreateLogicalNetworkACL) SetNetworkProtocol(v string)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *CreateLogicalNetworkACL) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetSourceAddress

`func (o *CreateLogicalNetworkACL) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *CreateLogicalNetworkACL) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *CreateLogicalNetworkACL) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *CreateLogicalNetworkACL) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreateLogicalNetworkACL) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreateLogicalNetworkACL) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreateLogicalNetworkACL) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *CreateLogicalNetworkACL) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetSourcePort

`func (o *CreateLogicalNetworkACL) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *CreateLogicalNetworkACL) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *CreateLogicalNetworkACL) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *CreateLogicalNetworkACL) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *CreateLogicalNetworkACL) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *CreateLogicalNetworkACL) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *CreateLogicalNetworkACL) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *CreateLogicalNetworkACL) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetSourceMac

`func (o *CreateLogicalNetworkACL) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *CreateLogicalNetworkACL) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *CreateLogicalNetworkACL) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *CreateLogicalNetworkACL) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetDestinationMac

`func (o *CreateLogicalNetworkACL) GetDestinationMac() string`

GetDestinationMac returns the DestinationMac field if non-nil, zero value otherwise.

### GetDestinationMacOk

`func (o *CreateLogicalNetworkACL) GetDestinationMacOk() (*string, bool)`

GetDestinationMacOk returns a tuple with the DestinationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMac

`func (o *CreateLogicalNetworkACL) SetDestinationMac(v string)`

SetDestinationMac sets DestinationMac field to given value.

### HasDestinationMac

`func (o *CreateLogicalNetworkACL) HasDestinationMac() bool`

HasDestinationMac returns a boolean if a field has been set.

### GetEnforcementPoint

`func (o *CreateLogicalNetworkACL) GetEnforcementPoint() ACLEnforcementPoint`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *CreateLogicalNetworkACL) GetEnforcementPointOk() (*ACLEnforcementPoint, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *CreateLogicalNetworkACL) SetEnforcementPoint(v ACLEnforcementPoint)`

SetEnforcementPoint sets EnforcementPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



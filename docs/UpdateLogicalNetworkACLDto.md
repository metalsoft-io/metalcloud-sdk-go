# UpdateLogicalNetworkACLDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | Pointer to **string** | The type of the rule | [optional] 
**Direction** | Pointer to **string** | The direction of the rule | [optional] 
**Sequence** | Pointer to **int32** | The sequence number of the rule | [optional] 
**ForwardingAction** | Pointer to **string** | The forwarding action of the rule | [optional] 
**NetworkProtocol** | Pointer to **string** | The network protocol of the ACL rule | [optional] 
**SourceAddress** | Pointer to **string** | The source address of the rule if the rule type is IPv4, IPv6 | [optional] 
**DestinationAddress** | Pointer to **string** | The destination address of the rule if the rule type is IPv4, IPv6 | [optional] 
**SourcePort** | Pointer to **string** | The source port of the rule if the rule type is IPv4, IPv6, TCP or UDP | [optional] 
**DestinationPort** | Pointer to **string** | The destination port of the rule if the rule type is IPv4, IPv6 | [optional] 
**SourceMac** | Pointer to **string** | The source MAC address of the rule if the rule type is MAC | [optional] 
**DestinationMac** | Pointer to **string** | The destination MAC address of the rule if the rule type is MAC | [optional] 
**EnforcementPoint** | Pointer to **string** | The enforcement point of the rule | [optional] 

## Methods

### NewUpdateLogicalNetworkACLDto

`func NewUpdateLogicalNetworkACLDto() *UpdateLogicalNetworkACLDto`

NewUpdateLogicalNetworkACLDto instantiates a new UpdateLogicalNetworkACLDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLogicalNetworkACLDtoWithDefaults

`func NewUpdateLogicalNetworkACLDtoWithDefaults() *UpdateLogicalNetworkACLDto`

NewUpdateLogicalNetworkACLDtoWithDefaults instantiates a new UpdateLogicalNetworkACLDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *UpdateLogicalNetworkACLDto) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *UpdateLogicalNetworkACLDto) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *UpdateLogicalNetworkACLDto) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *UpdateLogicalNetworkACLDto) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetDirection

`func (o *UpdateLogicalNetworkACLDto) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *UpdateLogicalNetworkACLDto) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *UpdateLogicalNetworkACLDto) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *UpdateLogicalNetworkACLDto) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSequence

`func (o *UpdateLogicalNetworkACLDto) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *UpdateLogicalNetworkACLDto) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *UpdateLogicalNetworkACLDto) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *UpdateLogicalNetworkACLDto) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetForwardingAction

`func (o *UpdateLogicalNetworkACLDto) GetForwardingAction() string`

GetForwardingAction returns the ForwardingAction field if non-nil, zero value otherwise.

### GetForwardingActionOk

`func (o *UpdateLogicalNetworkACLDto) GetForwardingActionOk() (*string, bool)`

GetForwardingActionOk returns a tuple with the ForwardingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAction

`func (o *UpdateLogicalNetworkACLDto) SetForwardingAction(v string)`

SetForwardingAction sets ForwardingAction field to given value.

### HasForwardingAction

`func (o *UpdateLogicalNetworkACLDto) HasForwardingAction() bool`

HasForwardingAction returns a boolean if a field has been set.

### GetNetworkProtocol

`func (o *UpdateLogicalNetworkACLDto) GetNetworkProtocol() string`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *UpdateLogicalNetworkACLDto) GetNetworkProtocolOk() (*string, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *UpdateLogicalNetworkACLDto) SetNetworkProtocol(v string)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *UpdateLogicalNetworkACLDto) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetSourceAddress

`func (o *UpdateLogicalNetworkACLDto) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *UpdateLogicalNetworkACLDto) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *UpdateLogicalNetworkACLDto) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *UpdateLogicalNetworkACLDto) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *UpdateLogicalNetworkACLDto) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *UpdateLogicalNetworkACLDto) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *UpdateLogicalNetworkACLDto) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *UpdateLogicalNetworkACLDto) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetSourcePort

`func (o *UpdateLogicalNetworkACLDto) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *UpdateLogicalNetworkACLDto) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *UpdateLogicalNetworkACLDto) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *UpdateLogicalNetworkACLDto) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *UpdateLogicalNetworkACLDto) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *UpdateLogicalNetworkACLDto) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *UpdateLogicalNetworkACLDto) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *UpdateLogicalNetworkACLDto) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetSourceMac

`func (o *UpdateLogicalNetworkACLDto) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *UpdateLogicalNetworkACLDto) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *UpdateLogicalNetworkACLDto) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *UpdateLogicalNetworkACLDto) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetDestinationMac

`func (o *UpdateLogicalNetworkACLDto) GetDestinationMac() string`

GetDestinationMac returns the DestinationMac field if non-nil, zero value otherwise.

### GetDestinationMacOk

`func (o *UpdateLogicalNetworkACLDto) GetDestinationMacOk() (*string, bool)`

GetDestinationMacOk returns a tuple with the DestinationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMac

`func (o *UpdateLogicalNetworkACLDto) SetDestinationMac(v string)`

SetDestinationMac sets DestinationMac field to given value.

### HasDestinationMac

`func (o *UpdateLogicalNetworkACLDto) HasDestinationMac() bool`

HasDestinationMac returns a boolean if a field has been set.

### GetEnforcementPoint

`func (o *UpdateLogicalNetworkACLDto) GetEnforcementPoint() string`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *UpdateLogicalNetworkACLDto) GetEnforcementPointOk() (*string, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *UpdateLogicalNetworkACLDto) SetEnforcementPoint(v string)`

SetEnforcementPoint sets EnforcementPoint field to given value.

### HasEnforcementPoint

`func (o *UpdateLogicalNetworkACLDto) HasEnforcementPoint() bool`

HasEnforcementPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



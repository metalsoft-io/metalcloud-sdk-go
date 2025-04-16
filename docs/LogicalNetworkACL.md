# LogicalNetworkACL

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
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | The ID of the ACL | 
**EndpointGroupId** | **string** | The ID of the endpoint group logical network entity | 
**LogicalNetworkId** | **string** | The ID of the logical network | 

## Methods

### NewLogicalNetworkACL

`func NewLogicalNetworkACL(ruleType ACLType, direction ACLDirection, sequence int32, forwardingAction ACLForwardingAction, enforcementPoint ACLEnforcementPoint, createdTimestamp time.Time, updatedTimestamp time.Time, id string, endpointGroupId string, logicalNetworkId string, ) *LogicalNetworkACL`

NewLogicalNetworkACL instantiates a new LogicalNetworkACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkACLWithDefaults

`func NewLogicalNetworkACLWithDefaults() *LogicalNetworkACL`

NewLogicalNetworkACLWithDefaults instantiates a new LogicalNetworkACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *LogicalNetworkACL) GetRuleType() ACLType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *LogicalNetworkACL) GetRuleTypeOk() (*ACLType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *LogicalNetworkACL) SetRuleType(v ACLType)`

SetRuleType sets RuleType field to given value.


### GetDirection

`func (o *LogicalNetworkACL) GetDirection() ACLDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LogicalNetworkACL) GetDirectionOk() (*ACLDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LogicalNetworkACL) SetDirection(v ACLDirection)`

SetDirection sets Direction field to given value.


### GetSequence

`func (o *LogicalNetworkACL) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *LogicalNetworkACL) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *LogicalNetworkACL) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetForwardingAction

`func (o *LogicalNetworkACL) GetForwardingAction() ACLForwardingAction`

GetForwardingAction returns the ForwardingAction field if non-nil, zero value otherwise.

### GetForwardingActionOk

`func (o *LogicalNetworkACL) GetForwardingActionOk() (*ACLForwardingAction, bool)`

GetForwardingActionOk returns a tuple with the ForwardingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingAction

`func (o *LogicalNetworkACL) SetForwardingAction(v ACLForwardingAction)`

SetForwardingAction sets ForwardingAction field to given value.


### GetNetworkProtocol

`func (o *LogicalNetworkACL) GetNetworkProtocol() string`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *LogicalNetworkACL) GetNetworkProtocolOk() (*string, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *LogicalNetworkACL) SetNetworkProtocol(v string)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *LogicalNetworkACL) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.

### GetSourceAddress

`func (o *LogicalNetworkACL) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *LogicalNetworkACL) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *LogicalNetworkACL) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *LogicalNetworkACL) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *LogicalNetworkACL) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *LogicalNetworkACL) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *LogicalNetworkACL) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *LogicalNetworkACL) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetSourcePort

`func (o *LogicalNetworkACL) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *LogicalNetworkACL) GetSourcePortOk() (*string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *LogicalNetworkACL) SetSourcePort(v string)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *LogicalNetworkACL) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *LogicalNetworkACL) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *LogicalNetworkACL) GetDestinationPortOk() (*string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *LogicalNetworkACL) SetDestinationPort(v string)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *LogicalNetworkACL) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetSourceMac

`func (o *LogicalNetworkACL) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *LogicalNetworkACL) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *LogicalNetworkACL) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *LogicalNetworkACL) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetDestinationMac

`func (o *LogicalNetworkACL) GetDestinationMac() string`

GetDestinationMac returns the DestinationMac field if non-nil, zero value otherwise.

### GetDestinationMacOk

`func (o *LogicalNetworkACL) GetDestinationMacOk() (*string, bool)`

GetDestinationMacOk returns a tuple with the DestinationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMac

`func (o *LogicalNetworkACL) SetDestinationMac(v string)`

SetDestinationMac sets DestinationMac field to given value.

### HasDestinationMac

`func (o *LogicalNetworkACL) HasDestinationMac() bool`

HasDestinationMac returns a boolean if a field has been set.

### GetEnforcementPoint

`func (o *LogicalNetworkACL) GetEnforcementPoint() ACLEnforcementPoint`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *LogicalNetworkACL) GetEnforcementPointOk() (*ACLEnforcementPoint, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *LogicalNetworkACL) SetEnforcementPoint(v ACLEnforcementPoint)`

SetEnforcementPoint sets EnforcementPoint field to given value.


### GetCreatedTimestamp

`func (o *LogicalNetworkACL) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *LogicalNetworkACL) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *LogicalNetworkACL) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *LogicalNetworkACL) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *LogicalNetworkACL) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *LogicalNetworkACL) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *LogicalNetworkACL) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogicalNetworkACL) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogicalNetworkACL) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LogicalNetworkACL) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *LogicalNetworkACL) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkACL) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkACL) SetId(v string)`

SetId sets Id field to given value.


### GetEndpointGroupId

`func (o *LogicalNetworkACL) GetEndpointGroupId() string`

GetEndpointGroupId returns the EndpointGroupId field if non-nil, zero value otherwise.

### GetEndpointGroupIdOk

`func (o *LogicalNetworkACL) GetEndpointGroupIdOk() (*string, bool)`

GetEndpointGroupIdOk returns a tuple with the EndpointGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupId

`func (o *LogicalNetworkACL) SetEndpointGroupId(v string)`

SetEndpointGroupId sets EndpointGroupId field to given value.


### GetLogicalNetworkId

`func (o *LogicalNetworkACL) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *LogicalNetworkACL) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *LogicalNetworkACL) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



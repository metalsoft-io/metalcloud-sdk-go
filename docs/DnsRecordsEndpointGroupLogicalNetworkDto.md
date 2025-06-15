# DnsRecordsEndpointGroupLogicalNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisionInstanceDnsRecords** | Pointer to **bool** | If true, DNS records (type A and PTR) will be created for each server instance that is part of this server instance group. The DNS record will be created with the default \&quot;hostname\&quot; of &lt;server-instance&gt; suffixed with the server instance ID (e.g., \&quot;server-instance-24\&quot;) plus the DNS zone name. If the \&quot;hostname\&quot; property is set on the server instance, it will use that value to construct the DNS record name. | [optional] [default to false]
**ProvisionLoadBalancingDnsRecord** | Pointer to **bool** | Flag to indicate if for the server instance group should create a DNS Load Balancing record. If true, a DNS Load Balancing record (type A) will be created that points to all server instances that are part of this server instance group. The DNS Load Balancing record will be created with the default \&quot;hostname\&quot; of &lt;server-instance-group&gt; suffixed with the server instance group ID (e.g., \&quot;server-instance-group-34\&quot;) plus the DNS zone name. If the \&quot;hostname\&quot; property is set, it will use that value to construct the DNS Load Balancing record name. | [optional] [default to false]

## Methods

### NewDnsRecordsEndpointGroupLogicalNetworkDto

`func NewDnsRecordsEndpointGroupLogicalNetworkDto() *DnsRecordsEndpointGroupLogicalNetworkDto`

NewDnsRecordsEndpointGroupLogicalNetworkDto instantiates a new DnsRecordsEndpointGroupLogicalNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordsEndpointGroupLogicalNetworkDtoWithDefaults

`func NewDnsRecordsEndpointGroupLogicalNetworkDtoWithDefaults() *DnsRecordsEndpointGroupLogicalNetworkDto`

NewDnsRecordsEndpointGroupLogicalNetworkDtoWithDefaults instantiates a new DnsRecordsEndpointGroupLogicalNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisionInstanceDnsRecords

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) GetProvisionInstanceDnsRecords() bool`

GetProvisionInstanceDnsRecords returns the ProvisionInstanceDnsRecords field if non-nil, zero value otherwise.

### GetProvisionInstanceDnsRecordsOk

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) GetProvisionInstanceDnsRecordsOk() (*bool, bool)`

GetProvisionInstanceDnsRecordsOk returns a tuple with the ProvisionInstanceDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionInstanceDnsRecords

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) SetProvisionInstanceDnsRecords(v bool)`

SetProvisionInstanceDnsRecords sets ProvisionInstanceDnsRecords field to given value.

### HasProvisionInstanceDnsRecords

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) HasProvisionInstanceDnsRecords() bool`

HasProvisionInstanceDnsRecords returns a boolean if a field has been set.

### GetProvisionLoadBalancingDnsRecord

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) GetProvisionLoadBalancingDnsRecord() bool`

GetProvisionLoadBalancingDnsRecord returns the ProvisionLoadBalancingDnsRecord field if non-nil, zero value otherwise.

### GetProvisionLoadBalancingDnsRecordOk

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) GetProvisionLoadBalancingDnsRecordOk() (*bool, bool)`

GetProvisionLoadBalancingDnsRecordOk returns a tuple with the ProvisionLoadBalancingDnsRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionLoadBalancingDnsRecord

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) SetProvisionLoadBalancingDnsRecord(v bool)`

SetProvisionLoadBalancingDnsRecord sets ProvisionLoadBalancingDnsRecord field to given value.

### HasProvisionLoadBalancingDnsRecord

`func (o *DnsRecordsEndpointGroupLogicalNetworkDto) HasProvisionLoadBalancingDnsRecord() bool`

HasProvisionLoadBalancingDnsRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



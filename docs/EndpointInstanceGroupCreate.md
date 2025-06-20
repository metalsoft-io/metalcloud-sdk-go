# EndpointInstanceGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Product Instance label. Will be automatically generated if not provided. | [optional] 
**EndpointGroupName** | Pointer to **string** |  | [optional] 
**ExtensionInstanceId** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS Load Balancing record. If set, this will be used as the DNS Load Balancing record name instead of the default \&quot;endpoint-instance-group\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;dnsLoadBalancingRecord\&quot; is true. It will be automatically suffixed with the endpoint instance group ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 
**ResourcePoolId** | Pointer to **int32** | The resource pool assigned to this instance array | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEndpointInstanceGroupCreate

`func NewEndpointInstanceGroupCreate() *EndpointInstanceGroupCreate`

NewEndpointInstanceGroupCreate instantiates a new EndpointInstanceGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceGroupCreateWithDefaults

`func NewEndpointInstanceGroupCreateWithDefaults() *EndpointInstanceGroupCreate`

NewEndpointInstanceGroupCreateWithDefaults instantiates a new EndpointInstanceGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EndpointInstanceGroupCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceGroupCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceGroupCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EndpointInstanceGroupCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *EndpointInstanceGroupCreate) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *EndpointInstanceGroupCreate) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *EndpointInstanceGroupCreate) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *EndpointInstanceGroupCreate) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetExtensionInstanceId

`func (o *EndpointInstanceGroupCreate) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *EndpointInstanceGroupCreate) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *EndpointInstanceGroupCreate) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *EndpointInstanceGroupCreate) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetHostname

`func (o *EndpointInstanceGroupCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EndpointInstanceGroupCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EndpointInstanceGroupCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EndpointInstanceGroupCreate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *EndpointInstanceGroupCreate) GetResourcePoolId() int32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *EndpointInstanceGroupCreate) GetResourcePoolIdOk() (*int32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *EndpointInstanceGroupCreate) SetResourcePoolId(v int32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *EndpointInstanceGroupCreate) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetTags

`func (o *EndpointInstanceGroupCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EndpointInstanceGroupCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EndpointInstanceGroupCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EndpointInstanceGroupCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



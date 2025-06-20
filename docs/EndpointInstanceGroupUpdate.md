# EndpointInstanceGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The endpoint instance group label. | [optional] 
**EndpointGroupName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS Load Balancing record. If set, this will be used as the DNS Load Balancing record name instead of the default \&quot;endpoint-instance-group\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;dnsLoadBalancingRecord\&quot; is true. It will be automatically suffixed with the endpoint instance group ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 

## Methods

### NewEndpointInstanceGroupUpdate

`func NewEndpointInstanceGroupUpdate() *EndpointInstanceGroupUpdate`

NewEndpointInstanceGroupUpdate instantiates a new EndpointInstanceGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceGroupUpdateWithDefaults

`func NewEndpointInstanceGroupUpdateWithDefaults() *EndpointInstanceGroupUpdate`

NewEndpointInstanceGroupUpdateWithDefaults instantiates a new EndpointInstanceGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EndpointInstanceGroupUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceGroupUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceGroupUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EndpointInstanceGroupUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *EndpointInstanceGroupUpdate) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *EndpointInstanceGroupUpdate) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *EndpointInstanceGroupUpdate) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *EndpointInstanceGroupUpdate) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetHostname

`func (o *EndpointInstanceGroupUpdate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EndpointInstanceGroupUpdate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EndpointInstanceGroupUpdate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EndpointInstanceGroupUpdate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



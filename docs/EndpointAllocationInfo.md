# EndpointAllocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **float32** | The id of the instance. | 
**InstanceLabel** | **string** | The label of the instance. | 
**InstanceGroupLabel** | **string** | The label of the instance group. | 
**ExtensionInstanceId** | Pointer to **float32** | The id of the extension instance. | [optional] 
**Infrastructure** | **map[string]interface{}** | The infrastructure of the instance. | 
**OsTemplateId** | Pointer to **float32** | The id of the os template used by the instance. | [optional] 

## Methods

### NewEndpointAllocationInfo

`func NewEndpointAllocationInfo(instanceId float32, instanceLabel string, instanceGroupLabel string, infrastructure map[string]interface{}, ) *EndpointAllocationInfo`

NewEndpointAllocationInfo instantiates a new EndpointAllocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointAllocationInfoWithDefaults

`func NewEndpointAllocationInfoWithDefaults() *EndpointAllocationInfo`

NewEndpointAllocationInfoWithDefaults instantiates a new EndpointAllocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *EndpointAllocationInfo) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *EndpointAllocationInfo) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *EndpointAllocationInfo) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceLabel

`func (o *EndpointAllocationInfo) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *EndpointAllocationInfo) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *EndpointAllocationInfo) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.


### GetInstanceGroupLabel

`func (o *EndpointAllocationInfo) GetInstanceGroupLabel() string`

GetInstanceGroupLabel returns the InstanceGroupLabel field if non-nil, zero value otherwise.

### GetInstanceGroupLabelOk

`func (o *EndpointAllocationInfo) GetInstanceGroupLabelOk() (*string, bool)`

GetInstanceGroupLabelOk returns a tuple with the InstanceGroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupLabel

`func (o *EndpointAllocationInfo) SetInstanceGroupLabel(v string)`

SetInstanceGroupLabel sets InstanceGroupLabel field to given value.


### GetExtensionInstanceId

`func (o *EndpointAllocationInfo) GetExtensionInstanceId() float32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *EndpointAllocationInfo) GetExtensionInstanceIdOk() (*float32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *EndpointAllocationInfo) SetExtensionInstanceId(v float32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *EndpointAllocationInfo) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetInfrastructure

`func (o *EndpointAllocationInfo) GetInfrastructure() map[string]interface{}`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *EndpointAllocationInfo) GetInfrastructureOk() (*map[string]interface{}, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *EndpointAllocationInfo) SetInfrastructure(v map[string]interface{})`

SetInfrastructure sets Infrastructure field to given value.


### GetOsTemplateId

`func (o *EndpointAllocationInfo) GetOsTemplateId() float32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *EndpointAllocationInfo) GetOsTemplateIdOk() (*float32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *EndpointAllocationInfo) SetOsTemplateId(v float32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *EndpointAllocationInfo) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



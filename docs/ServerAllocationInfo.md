# ServerAllocationInfo

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

### NewServerAllocationInfo

`func NewServerAllocationInfo(instanceId float32, instanceLabel string, instanceGroupLabel string, infrastructure map[string]interface{}, ) *ServerAllocationInfo`

NewServerAllocationInfo instantiates a new ServerAllocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerAllocationInfoWithDefaults

`func NewServerAllocationInfoWithDefaults() *ServerAllocationInfo`

NewServerAllocationInfoWithDefaults instantiates a new ServerAllocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *ServerAllocationInfo) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ServerAllocationInfo) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ServerAllocationInfo) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceLabel

`func (o *ServerAllocationInfo) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ServerAllocationInfo) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ServerAllocationInfo) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.


### GetInstanceGroupLabel

`func (o *ServerAllocationInfo) GetInstanceGroupLabel() string`

GetInstanceGroupLabel returns the InstanceGroupLabel field if non-nil, zero value otherwise.

### GetInstanceGroupLabelOk

`func (o *ServerAllocationInfo) GetInstanceGroupLabelOk() (*string, bool)`

GetInstanceGroupLabelOk returns a tuple with the InstanceGroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupLabel

`func (o *ServerAllocationInfo) SetInstanceGroupLabel(v string)`

SetInstanceGroupLabel sets InstanceGroupLabel field to given value.


### GetExtensionInstanceId

`func (o *ServerAllocationInfo) GetExtensionInstanceId() float32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *ServerAllocationInfo) GetExtensionInstanceIdOk() (*float32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *ServerAllocationInfo) SetExtensionInstanceId(v float32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *ServerAllocationInfo) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetInfrastructure

`func (o *ServerAllocationInfo) GetInfrastructure() map[string]interface{}`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ServerAllocationInfo) GetInfrastructureOk() (*map[string]interface{}, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ServerAllocationInfo) SetInfrastructure(v map[string]interface{})`

SetInfrastructure sets Infrastructure field to given value.


### GetOsTemplateId

`func (o *ServerAllocationInfo) GetOsTemplateId() float32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerAllocationInfo) GetOsTemplateIdOk() (*float32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerAllocationInfo) SetOsTemplateId(v float32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerAllocationInfo) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



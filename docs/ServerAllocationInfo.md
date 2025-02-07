# ServerAllocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **float32** | The id of the instance. | 
**Infrastructure** | **map[string]interface{}** | The infrastructure of the instance. | 
**VolumeTemplate** | Pointer to **map[string]interface{}** | The volume template of the instance. | [optional] 

## Methods

### NewServerAllocationInfo

`func NewServerAllocationInfo(instanceId float32, infrastructure map[string]interface{}, ) *ServerAllocationInfo`

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


### GetVolumeTemplate

`func (o *ServerAllocationInfo) GetVolumeTemplate() map[string]interface{}`

GetVolumeTemplate returns the VolumeTemplate field if non-nil, zero value otherwise.

### GetVolumeTemplateOk

`func (o *ServerAllocationInfo) GetVolumeTemplateOk() (*map[string]interface{}, bool)`

GetVolumeTemplateOk returns a tuple with the VolumeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplate

`func (o *ServerAllocationInfo) SetVolumeTemplate(v map[string]interface{})`

SetVolumeTemplate sets VolumeTemplate field to given value.

### HasVolumeTemplate

`func (o *ServerAllocationInfo) HasVolumeTemplate() bool`

HasVolumeTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



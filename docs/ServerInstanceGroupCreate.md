# ServerInstanceGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The server instance group label. Will be automatically generated if not provided. | [optional] 
**ExtensionInstanceId** | Pointer to **int32** |  | [optional] 
**InstanceCount** | Pointer to **int32** |  | [optional] [default to 1]
**VolumeTemplateId** | Pointer to **int32** |  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**DiskCount** | Pointer to **int32** |  | [optional] 
**DiskSizeMbytes** | Pointer to **int32** |  | [optional] 
**DiskTypes** | Pointer to **[]string** |  | [optional] 
**DefaultServerProfileID** | Pointer to **int32** | The group&#39;s default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied. | [optional] 

## Methods

### NewServerInstanceGroupCreate

`func NewServerInstanceGroupCreate() *ServerInstanceGroupCreate`

NewServerInstanceGroupCreate instantiates a new ServerInstanceGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupCreateWithDefaults

`func NewServerInstanceGroupCreateWithDefaults() *ServerInstanceGroupCreate`

NewServerInstanceGroupCreateWithDefaults instantiates a new ServerInstanceGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceGroupCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceGroupCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetExtensionInstanceId

`func (o *ServerInstanceGroupCreate) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *ServerInstanceGroupCreate) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *ServerInstanceGroupCreate) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *ServerInstanceGroupCreate) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ServerInstanceGroupCreate) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ServerInstanceGroupCreate) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ServerInstanceGroupCreate) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ServerInstanceGroupCreate) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetVolumeTemplateId

`func (o *ServerInstanceGroupCreate) GetVolumeTemplateId() int32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *ServerInstanceGroupCreate) GetVolumeTemplateIdOk() (*int32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *ServerInstanceGroupCreate) SetVolumeTemplateId(v int32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *ServerInstanceGroupCreate) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceGroupCreate) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroupCreate) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroupCreate) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroupCreate) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetDiskCount

`func (o *ServerInstanceGroupCreate) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceGroupCreate) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceGroupCreate) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *ServerInstanceGroupCreate) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetDiskSizeMbytes

`func (o *ServerInstanceGroupCreate) GetDiskSizeMbytes() int32`

GetDiskSizeMbytes returns the DiskSizeMbytes field if non-nil, zero value otherwise.

### GetDiskSizeMbytesOk

`func (o *ServerInstanceGroupCreate) GetDiskSizeMbytesOk() (*int32, bool)`

GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMbytes

`func (o *ServerInstanceGroupCreate) SetDiskSizeMbytes(v int32)`

SetDiskSizeMbytes sets DiskSizeMbytes field to given value.

### HasDiskSizeMbytes

`func (o *ServerInstanceGroupCreate) HasDiskSizeMbytes() bool`

HasDiskSizeMbytes returns a boolean if a field has been set.

### GetDiskTypes

`func (o *ServerInstanceGroupCreate) GetDiskTypes() []string`

GetDiskTypes returns the DiskTypes field if non-nil, zero value otherwise.

### GetDiskTypesOk

`func (o *ServerInstanceGroupCreate) GetDiskTypesOk() (*[]string, bool)`

GetDiskTypesOk returns a tuple with the DiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypes

`func (o *ServerInstanceGroupCreate) SetDiskTypes(v []string)`

SetDiskTypes sets DiskTypes field to given value.

### HasDiskTypes

`func (o *ServerInstanceGroupCreate) HasDiskTypes() bool`

HasDiskTypes returns a boolean if a field has been set.

### GetDefaultServerProfileID

`func (o *ServerInstanceGroupCreate) GetDefaultServerProfileID() int32`

GetDefaultServerProfileID returns the DefaultServerProfileID field if non-nil, zero value otherwise.

### GetDefaultServerProfileIDOk

`func (o *ServerInstanceGroupCreate) GetDefaultServerProfileIDOk() (*int32, bool)`

GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerProfileID

`func (o *ServerInstanceGroupCreate) SetDefaultServerProfileID(v int32)`

SetDefaultServerProfileID sets DefaultServerProfileID field to given value.

### HasDefaultServerProfileID

`func (o *ServerInstanceGroupCreate) HasDefaultServerProfileID() bool`

HasDefaultServerProfileID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



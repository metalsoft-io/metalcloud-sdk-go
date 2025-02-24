# ServerInstanceGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The server instance group label. Will be automatically generated if not provided. | [optional] 
**ServerGroupName** | Pointer to **string** |  | [optional] 
**InstanceCount** | Pointer to **int32** |  | [optional] [default to 1]
**VolumeTemplateId** | Pointer to **int32** |  | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**DiskCount** | Pointer to **int32** |  | [optional] 
**DiskSizeMbytes** | Pointer to **int32** |  | [optional] 
**DiskTypes** | Pointer to **[]string** |  | [optional] 
**DefaultServerProfileID** | Pointer to **int32** | The group&#39;s default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied. | [optional] 

## Methods

### NewServerInstanceGroupUpdate

`func NewServerInstanceGroupUpdate() *ServerInstanceGroupUpdate`

NewServerInstanceGroupUpdate instantiates a new ServerInstanceGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupUpdateWithDefaults

`func NewServerInstanceGroupUpdateWithDefaults() *ServerInstanceGroupUpdate`

NewServerInstanceGroupUpdateWithDefaults instantiates a new ServerInstanceGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceGroupUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceGroupUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetServerGroupName

`func (o *ServerInstanceGroupUpdate) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ServerInstanceGroupUpdate) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ServerInstanceGroupUpdate) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ServerInstanceGroupUpdate) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ServerInstanceGroupUpdate) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ServerInstanceGroupUpdate) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ServerInstanceGroupUpdate) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ServerInstanceGroupUpdate) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetVolumeTemplateId

`func (o *ServerInstanceGroupUpdate) GetVolumeTemplateId() int32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *ServerInstanceGroupUpdate) GetVolumeTemplateIdOk() (*int32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *ServerInstanceGroupUpdate) SetVolumeTemplateId(v int32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *ServerInstanceGroupUpdate) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceGroupUpdate) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroupUpdate) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroupUpdate) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroupUpdate) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetDiskCount

`func (o *ServerInstanceGroupUpdate) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceGroupUpdate) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceGroupUpdate) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *ServerInstanceGroupUpdate) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetDiskSizeMbytes

`func (o *ServerInstanceGroupUpdate) GetDiskSizeMbytes() int32`

GetDiskSizeMbytes returns the DiskSizeMbytes field if non-nil, zero value otherwise.

### GetDiskSizeMbytesOk

`func (o *ServerInstanceGroupUpdate) GetDiskSizeMbytesOk() (*int32, bool)`

GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMbytes

`func (o *ServerInstanceGroupUpdate) SetDiskSizeMbytes(v int32)`

SetDiskSizeMbytes sets DiskSizeMbytes field to given value.

### HasDiskSizeMbytes

`func (o *ServerInstanceGroupUpdate) HasDiskSizeMbytes() bool`

HasDiskSizeMbytes returns a boolean if a field has been set.

### GetDiskTypes

`func (o *ServerInstanceGroupUpdate) GetDiskTypes() []string`

GetDiskTypes returns the DiskTypes field if non-nil, zero value otherwise.

### GetDiskTypesOk

`func (o *ServerInstanceGroupUpdate) GetDiskTypesOk() (*[]string, bool)`

GetDiskTypesOk returns a tuple with the DiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypes

`func (o *ServerInstanceGroupUpdate) SetDiskTypes(v []string)`

SetDiskTypes sets DiskTypes field to given value.

### HasDiskTypes

`func (o *ServerInstanceGroupUpdate) HasDiskTypes() bool`

HasDiskTypes returns a boolean if a field has been set.

### GetDefaultServerProfileID

`func (o *ServerInstanceGroupUpdate) GetDefaultServerProfileID() int32`

GetDefaultServerProfileID returns the DefaultServerProfileID field if non-nil, zero value otherwise.

### GetDefaultServerProfileIDOk

`func (o *ServerInstanceGroupUpdate) GetDefaultServerProfileIDOk() (*int32, bool)`

GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerProfileID

`func (o *ServerInstanceGroupUpdate) SetDefaultServerProfileID(v int32)`

SetDefaultServerProfileID sets DefaultServerProfileID field to given value.

### HasDefaultServerProfileID

`func (o *ServerInstanceGroupUpdate) HasDefaultServerProfileID() bool`

HasDefaultServerProfileID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



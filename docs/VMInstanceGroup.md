# VMInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id of the VM Instance Group. | 
**InfrastructureId** | **float32** | Id of the Infrastructure. | 
**ServiceStatus** | **string** | Status of the VM Instance Group. | 
**ChangeId** | **float32** | Id of the VM Instance Group change object. | 
**Label** | **string** | Name of the VM Instance Group. | 
**InstanceCount** | **float32** | Number of VM instances in the VM Instance Group. | 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance Group creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group update. | 
**Operation** | **map[string]interface{}** | Operation object for the VM Instance Group. | 
**VmInstance** | [**[]VMInstance**](VMInstance.md) | Array of VM instances in the VM Instance Group. | 
**Tags** | Pointer to **[]string** | Tags for the VM Instance Group. | [optional] 
**DiskSizeGB** | **float32** | Disk size in GB for each VM Instance in the VM Instance Group. | 
**VolumeTemplateId** | Pointer to **float32** | Id of the template used by the VM Instance Group. | [optional] 
**VmInstanceGroupInterfaces** | Pointer to [**[]VMInstanceGroupInterface**](VMInstanceGroupInterface.md) | Interfaces for the VM Instance Group | [optional] 
**NetworkIdToNetworkProfileId** | Pointer to **map[string]interface{}** | Network Id to Network Profile Id for the VM Instance Group. This is a JSON object. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMInstanceGroup

`func NewVMInstanceGroup(id float32, infrastructureId float32, serviceStatus string, changeId float32, label string, instanceCount float32, createdTimestamp string, updatedTimestamp string, operation map[string]interface{}, vmInstance []VMInstance, diskSizeGB float32, links map[string]interface{}, ) *VMInstanceGroup`

NewVMInstanceGroup instantiates a new VMInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupWithDefaults

`func NewVMInstanceGroupWithDefaults() *VMInstanceGroup`

NewVMInstanceGroupWithDefaults instantiates a new VMInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMInstanceGroup) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInstanceGroup) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInstanceGroup) SetId(v float32)`

SetId sets Id field to given value.


### GetInfrastructureId

`func (o *VMInstanceGroup) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VMInstanceGroup) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VMInstanceGroup) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetServiceStatus

`func (o *VMInstanceGroup) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *VMInstanceGroup) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *VMInstanceGroup) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetChangeId

`func (o *VMInstanceGroup) GetChangeId() float32`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *VMInstanceGroup) GetChangeIdOk() (*float32, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *VMInstanceGroup) SetChangeId(v float32)`

SetChangeId sets ChangeId field to given value.


### GetLabel

`func (o *VMInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInstanceCount

`func (o *VMInstanceGroup) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *VMInstanceGroup) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *VMInstanceGroup) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.


### GetCreatedTimestamp

`func (o *VMInstanceGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VMInstanceGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VMInstanceGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *VMInstanceGroup) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstanceGroup) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstanceGroup) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetOperation

`func (o *VMInstanceGroup) GetOperation() map[string]interface{}`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VMInstanceGroup) GetOperationOk() (*map[string]interface{}, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VMInstanceGroup) SetOperation(v map[string]interface{})`

SetOperation sets Operation field to given value.


### GetVmInstance

`func (o *VMInstanceGroup) GetVmInstance() []VMInstance`

GetVmInstance returns the VmInstance field if non-nil, zero value otherwise.

### GetVmInstanceOk

`func (o *VMInstanceGroup) GetVmInstanceOk() (*[]VMInstance, bool)`

GetVmInstanceOk returns a tuple with the VmInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstance

`func (o *VMInstanceGroup) SetVmInstance(v []VMInstance)`

SetVmInstance sets VmInstance field to given value.


### GetTags

`func (o *VMInstanceGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VMInstanceGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VMInstanceGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VMInstanceGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *VMInstanceGroup) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *VMInstanceGroup) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *VMInstanceGroup) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetVolumeTemplateId

`func (o *VMInstanceGroup) GetVolumeTemplateId() float32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *VMInstanceGroup) GetVolumeTemplateIdOk() (*float32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *VMInstanceGroup) SetVolumeTemplateId(v float32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *VMInstanceGroup) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetVmInstanceGroupInterfaces

`func (o *VMInstanceGroup) GetVmInstanceGroupInterfaces() []VMInstanceGroupInterface`

GetVmInstanceGroupInterfaces returns the VmInstanceGroupInterfaces field if non-nil, zero value otherwise.

### GetVmInstanceGroupInterfacesOk

`func (o *VMInstanceGroup) GetVmInstanceGroupInterfacesOk() (*[]VMInstanceGroupInterface, bool)`

GetVmInstanceGroupInterfacesOk returns a tuple with the VmInstanceGroupInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceGroupInterfaces

`func (o *VMInstanceGroup) SetVmInstanceGroupInterfaces(v []VMInstanceGroupInterface)`

SetVmInstanceGroupInterfaces sets VmInstanceGroupInterfaces field to given value.

### HasVmInstanceGroupInterfaces

`func (o *VMInstanceGroup) HasVmInstanceGroupInterfaces() bool`

HasVmInstanceGroupInterfaces returns a boolean if a field has been set.

### GetNetworkIdToNetworkProfileId

`func (o *VMInstanceGroup) GetNetworkIdToNetworkProfileId() map[string]interface{}`

GetNetworkIdToNetworkProfileId returns the NetworkIdToNetworkProfileId field if non-nil, zero value otherwise.

### GetNetworkIdToNetworkProfileIdOk

`func (o *VMInstanceGroup) GetNetworkIdToNetworkProfileIdOk() (*map[string]interface{}, bool)`

GetNetworkIdToNetworkProfileIdOk returns a tuple with the NetworkIdToNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIdToNetworkProfileId

`func (o *VMInstanceGroup) SetNetworkIdToNetworkProfileId(v map[string]interface{})`

SetNetworkIdToNetworkProfileId sets NetworkIdToNetworkProfileId field to given value.

### HasNetworkIdToNetworkProfileId

`func (o *VMInstanceGroup) HasNetworkIdToNetworkProfileId() bool`

HasNetworkIdToNetworkProfileId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *VMInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *VMInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *VMInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *VMInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetGuiSettings

`func (o *VMInstanceGroup) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *VMInstanceGroup) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *VMInstanceGroup) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *VMInstanceGroup) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetLinks

`func (o *VMInstanceGroup) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMInstanceGroup) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMInstanceGroup) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



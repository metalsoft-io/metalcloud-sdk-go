# VMInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM Instance ID | 
**GroupId** | **float32** | Id of the VM Instance Group. | 
**InfrastructureId** | **float32** | Id of the Infrastructure. | 
**Label** | **string** | Name of the VM Instance. | 
**VmId** | **float32** | Id of the VM. | 
**TypeId** | **float32** | Id of the VM Type. | 
**Tags** | Pointer to **[]string** | Tags for the VM Instance. | [optional] 
**ServiceStatus** | **string** | Service status of the VM Instance. | 
**ChangeId** | **float32** | Id of the VM Instance change object. | 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance last update. | 
**Operation** | **map[string]interface{}** | VM Instance change object. | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 
**DiskSizeGB** | **float32** | Disk size in GB of the VM Instance. | 
**TemplateIdOrigin** | Pointer to **float32** | Id of the template used by the VM Instance. | [optional] 
**VmPoolId** | Pointer to **float32** | Id of the VM Pool. | [optional] 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMInstance

`func NewVMInstance(id float32, groupId float32, infrastructureId float32, label string, vmId float32, typeId float32, serviceStatus string, changeId float32, createdTimestamp string, updatedTimestamp string, operation map[string]interface{}, diskSizeGB float32, links map[string]interface{}, ) *VMInstance`

NewVMInstance instantiates a new VMInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceWithDefaults

`func NewVMInstanceWithDefaults() *VMInstance`

NewVMInstanceWithDefaults instantiates a new VMInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMInstance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInstance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInstance) SetId(v float32)`

SetId sets Id field to given value.


### GetGroupId

`func (o *VMInstance) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VMInstance) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VMInstance) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetInfrastructureId

`func (o *VMInstance) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VMInstance) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VMInstance) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetLabel

`func (o *VMInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVmId

`func (o *VMInstance) GetVmId() float32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VMInstance) GetVmIdOk() (*float32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VMInstance) SetVmId(v float32)`

SetVmId sets VmId field to given value.


### GetTypeId

`func (o *VMInstance) GetTypeId() float32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VMInstance) GetTypeIdOk() (*float32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VMInstance) SetTypeId(v float32)`

SetTypeId sets TypeId field to given value.


### GetTags

`func (o *VMInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VMInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VMInstance) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VMInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetServiceStatus

`func (o *VMInstance) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *VMInstance) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *VMInstance) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetChangeId

`func (o *VMInstance) GetChangeId() float32`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *VMInstance) GetChangeIdOk() (*float32, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *VMInstance) SetChangeId(v float32)`

SetChangeId sets ChangeId field to given value.


### GetCreatedTimestamp

`func (o *VMInstance) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VMInstance) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VMInstance) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *VMInstance) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstance) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstance) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetOperation

`func (o *VMInstance) GetOperation() map[string]interface{}`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VMInstance) GetOperationOk() (*map[string]interface{}, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VMInstance) SetOperation(v map[string]interface{})`

SetOperation sets Operation field to given value.


### GetCustomVariables

`func (o *VMInstance) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *VMInstance) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *VMInstance) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *VMInstance) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *VMInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *VMInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *VMInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetTemplateIdOrigin

`func (o *VMInstance) GetTemplateIdOrigin() float32`

GetTemplateIdOrigin returns the TemplateIdOrigin field if non-nil, zero value otherwise.

### GetTemplateIdOriginOk

`func (o *VMInstance) GetTemplateIdOriginOk() (*float32, bool)`

GetTemplateIdOriginOk returns a tuple with the TemplateIdOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIdOrigin

`func (o *VMInstance) SetTemplateIdOrigin(v float32)`

SetTemplateIdOrigin sets TemplateIdOrigin field to given value.

### HasTemplateIdOrigin

`func (o *VMInstance) HasTemplateIdOrigin() bool`

HasTemplateIdOrigin returns a boolean if a field has been set.

### GetVmPoolId

`func (o *VMInstance) GetVmPoolId() float32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMInstance) GetVmPoolIdOk() (*float32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMInstance) SetVmPoolId(v float32)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *VMInstance) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetLinks

`func (o *VMInstance) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMInstance) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMInstance) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



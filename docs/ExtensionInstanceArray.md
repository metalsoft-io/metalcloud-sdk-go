# ExtensionInstanceArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the instance array. | 
**InstanceCount** | **string** | Instance count value or reference to variable. | 
**ServerType** | **string** | Server type variable reference. | 
**OsTemplate** | Pointer to **string** | OS template variable reference. | [optional] 
**ConnectedSharedDrives** | Pointer to **[]string** | Connected shared drive arrays. | [optional] 
**CustomVariables** | Pointer to [**[]CustomVariable**](CustomVariable.md) | Custom variables. The value may be a reference to an input variable. | [optional] 
**Dependencies** | Pointer to **[]string** | Labels of instance arrays this one depends on. | [optional] 
**LogicalNetworks** | Pointer to [**[]ExtensionInstanceArrayLogicalNetworkDto**](ExtensionInstanceArrayLogicalNetworkDto.md) | Logical networks for the instance array. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Server Instance Group. | [optional] 

## Methods

### NewExtensionInstanceArray

`func NewExtensionInstanceArray(label string, instanceCount string, serverType string, ) *ExtensionInstanceArray`

NewExtensionInstanceArray instantiates a new ExtensionInstanceArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstanceArrayWithDefaults

`func NewExtensionInstanceArrayWithDefaults() *ExtensionInstanceArray`

NewExtensionInstanceArrayWithDefaults instantiates a new ExtensionInstanceArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionInstanceArray) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInstanceArray) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInstanceArray) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetInstanceCount

`func (o *ExtensionInstanceArray) GetInstanceCount() string`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ExtensionInstanceArray) GetInstanceCountOk() (*string, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ExtensionInstanceArray) SetInstanceCount(v string)`

SetInstanceCount sets InstanceCount field to given value.


### GetServerType

`func (o *ExtensionInstanceArray) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ExtensionInstanceArray) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ExtensionInstanceArray) SetServerType(v string)`

SetServerType sets ServerType field to given value.


### GetOsTemplate

`func (o *ExtensionInstanceArray) GetOsTemplate() string`

GetOsTemplate returns the OsTemplate field if non-nil, zero value otherwise.

### GetOsTemplateOk

`func (o *ExtensionInstanceArray) GetOsTemplateOk() (*string, bool)`

GetOsTemplateOk returns a tuple with the OsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplate

`func (o *ExtensionInstanceArray) SetOsTemplate(v string)`

SetOsTemplate sets OsTemplate field to given value.

### HasOsTemplate

`func (o *ExtensionInstanceArray) HasOsTemplate() bool`

HasOsTemplate returns a boolean if a field has been set.

### GetConnectedSharedDrives

`func (o *ExtensionInstanceArray) GetConnectedSharedDrives() []string`

GetConnectedSharedDrives returns the ConnectedSharedDrives field if non-nil, zero value otherwise.

### GetConnectedSharedDrivesOk

`func (o *ExtensionInstanceArray) GetConnectedSharedDrivesOk() (*[]string, bool)`

GetConnectedSharedDrivesOk returns a tuple with the ConnectedSharedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSharedDrives

`func (o *ExtensionInstanceArray) SetConnectedSharedDrives(v []string)`

SetConnectedSharedDrives sets ConnectedSharedDrives field to given value.

### HasConnectedSharedDrives

`func (o *ExtensionInstanceArray) HasConnectedSharedDrives() bool`

HasConnectedSharedDrives returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ExtensionInstanceArray) GetCustomVariables() []CustomVariable`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ExtensionInstanceArray) GetCustomVariablesOk() (*[]CustomVariable, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ExtensionInstanceArray) SetCustomVariables(v []CustomVariable)`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ExtensionInstanceArray) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetDependencies

`func (o *ExtensionInstanceArray) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ExtensionInstanceArray) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ExtensionInstanceArray) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *ExtensionInstanceArray) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetLogicalNetworks

`func (o *ExtensionInstanceArray) GetLogicalNetworks() []ExtensionInstanceArrayLogicalNetworkDto`

GetLogicalNetworks returns the LogicalNetworks field if non-nil, zero value otherwise.

### GetLogicalNetworksOk

`func (o *ExtensionInstanceArray) GetLogicalNetworksOk() (*[]ExtensionInstanceArrayLogicalNetworkDto, bool)`

GetLogicalNetworksOk returns a tuple with the LogicalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworks

`func (o *ExtensionInstanceArray) SetLogicalNetworks(v []ExtensionInstanceArrayLogicalNetworkDto)`

SetLogicalNetworks sets LogicalNetworks field to given value.

### HasLogicalNetworks

`func (o *ExtensionInstanceArray) HasLogicalNetworks() bool`

HasLogicalNetworks returns a boolean if a field has been set.

### GetTags

`func (o *ExtensionInstanceArray) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtensionInstanceArray) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtensionInstanceArray) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtensionInstanceArray) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



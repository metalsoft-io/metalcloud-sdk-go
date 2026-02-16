# NetworkDeviceLinkAggregationConfigurationTemplateRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationType** | **string** | The type of link aggregation | 
**LocalSwitchId** | **float32** | The ID of the local switch. | 
**LocalSwitchIdentifier** | **string** | The identifier of the local switch. | 
**LocalSwitchRole** | **string** | The role of the local switch. | 
**LocalSwitchInterfaceName** | **string** | The name of the local switch interface. | 
**LocalSwitchInterfaceLagId** | **float32** | The LAG ID of the local switch interface. | 
**RemoteSwitchId** | **float32** | The ID of the remote switch. | 
**RemoteSwitchIdentifier** | **string** | The identifier of the remote switch. | 
**RemoteSwitchRole** | **string** | The role of the remote switch. | 
**RemoteSwitchInterfaceName** | **string** | The name of the remote switch interface. | 
**RemoteSwitchInterfaceLagId** | **float32** | The LAG ID of the remote switch interface. | 
**FabricCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the fabric. | [optional] 
**LinkCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the link. | [optional] 
**LocalSwitchCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the local switch. | [optional] 
**RemoteSwitchCustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the remote switch. | [optional] 

## Methods

### NewNetworkDeviceLinkAggregationConfigurationTemplateRecordSet

`func NewNetworkDeviceLinkAggregationConfigurationTemplateRecordSet(aggregationType string, localSwitchId float32, localSwitchIdentifier string, localSwitchRole string, localSwitchInterfaceName string, localSwitchInterfaceLagId float32, remoteSwitchId float32, remoteSwitchIdentifier string, remoteSwitchRole string, remoteSwitchInterfaceName string, remoteSwitchInterfaceLagId float32, ) *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet`

NewNetworkDeviceLinkAggregationConfigurationTemplateRecordSet instantiates a new NetworkDeviceLinkAggregationConfigurationTemplateRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceLinkAggregationConfigurationTemplateRecordSetWithDefaults

`func NewNetworkDeviceLinkAggregationConfigurationTemplateRecordSetWithDefaults() *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet`

NewNetworkDeviceLinkAggregationConfigurationTemplateRecordSetWithDefaults instantiates a new NetworkDeviceLinkAggregationConfigurationTemplateRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationType

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.


### GetLocalSwitchId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchId() float32`

GetLocalSwitchId returns the LocalSwitchId field if non-nil, zero value otherwise.

### GetLocalSwitchIdOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchIdOk() (*float32, bool)`

GetLocalSwitchIdOk returns a tuple with the LocalSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLocalSwitchId(v float32)`

SetLocalSwitchId sets LocalSwitchId field to given value.


### GetLocalSwitchIdentifier

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchIdentifier() string`

GetLocalSwitchIdentifier returns the LocalSwitchIdentifier field if non-nil, zero value otherwise.

### GetLocalSwitchIdentifierOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchIdentifierOk() (*string, bool)`

GetLocalSwitchIdentifierOk returns a tuple with the LocalSwitchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchIdentifier

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLocalSwitchIdentifier(v string)`

SetLocalSwitchIdentifier sets LocalSwitchIdentifier field to given value.


### GetLocalSwitchRole

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchRole() string`

GetLocalSwitchRole returns the LocalSwitchRole field if non-nil, zero value otherwise.

### GetLocalSwitchRoleOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchRoleOk() (*string, bool)`

GetLocalSwitchRoleOk returns a tuple with the LocalSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchRole

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLocalSwitchRole(v string)`

SetLocalSwitchRole sets LocalSwitchRole field to given value.


### GetLocalSwitchInterfaceName

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchInterfaceName() string`

GetLocalSwitchInterfaceName returns the LocalSwitchInterfaceName field if non-nil, zero value otherwise.

### GetLocalSwitchInterfaceNameOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchInterfaceNameOk() (*string, bool)`

GetLocalSwitchInterfaceNameOk returns a tuple with the LocalSwitchInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchInterfaceName

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLocalSwitchInterfaceName(v string)`

SetLocalSwitchInterfaceName sets LocalSwitchInterfaceName field to given value.


### GetLocalSwitchInterfaceLagId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchInterfaceLagId() float32`

GetLocalSwitchInterfaceLagId returns the LocalSwitchInterfaceLagId field if non-nil, zero value otherwise.

### GetLocalSwitchInterfaceLagIdOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchInterfaceLagIdOk() (*float32, bool)`

GetLocalSwitchInterfaceLagIdOk returns a tuple with the LocalSwitchInterfaceLagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchInterfaceLagId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLocalSwitchInterfaceLagId(v float32)`

SetLocalSwitchInterfaceLagId sets LocalSwitchInterfaceLagId field to given value.


### GetRemoteSwitchId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchId() float32`

GetRemoteSwitchId returns the RemoteSwitchId field if non-nil, zero value otherwise.

### GetRemoteSwitchIdOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchIdOk() (*float32, bool)`

GetRemoteSwitchIdOk returns a tuple with the RemoteSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetRemoteSwitchId(v float32)`

SetRemoteSwitchId sets RemoteSwitchId field to given value.


### GetRemoteSwitchIdentifier

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchIdentifier() string`

GetRemoteSwitchIdentifier returns the RemoteSwitchIdentifier field if non-nil, zero value otherwise.

### GetRemoteSwitchIdentifierOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchIdentifierOk() (*string, bool)`

GetRemoteSwitchIdentifierOk returns a tuple with the RemoteSwitchIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchIdentifier

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetRemoteSwitchIdentifier(v string)`

SetRemoteSwitchIdentifier sets RemoteSwitchIdentifier field to given value.


### GetRemoteSwitchRole

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchRole() string`

GetRemoteSwitchRole returns the RemoteSwitchRole field if non-nil, zero value otherwise.

### GetRemoteSwitchRoleOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchRoleOk() (*string, bool)`

GetRemoteSwitchRoleOk returns a tuple with the RemoteSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchRole

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetRemoteSwitchRole(v string)`

SetRemoteSwitchRole sets RemoteSwitchRole field to given value.


### GetRemoteSwitchInterfaceName

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceName() string`

GetRemoteSwitchInterfaceName returns the RemoteSwitchInterfaceName field if non-nil, zero value otherwise.

### GetRemoteSwitchInterfaceNameOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceNameOk() (*string, bool)`

GetRemoteSwitchInterfaceNameOk returns a tuple with the RemoteSwitchInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchInterfaceName

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetRemoteSwitchInterfaceName(v string)`

SetRemoteSwitchInterfaceName sets RemoteSwitchInterfaceName field to given value.


### GetRemoteSwitchInterfaceLagId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceLagId() float32`

GetRemoteSwitchInterfaceLagId returns the RemoteSwitchInterfaceLagId field if non-nil, zero value otherwise.

### GetRemoteSwitchInterfaceLagIdOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchInterfaceLagIdOk() (*float32, bool)`

GetRemoteSwitchInterfaceLagIdOk returns a tuple with the RemoteSwitchInterfaceLagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchInterfaceLagId

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetRemoteSwitchInterfaceLagId(v float32)`

SetRemoteSwitchInterfaceLagId sets RemoteSwitchInterfaceLagId field to given value.


### GetFabricCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetFabricCustomVariables() map[string]interface{}`

GetFabricCustomVariables returns the FabricCustomVariables field if non-nil, zero value otherwise.

### GetFabricCustomVariablesOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetFabricCustomVariablesOk() (*map[string]interface{}, bool)`

GetFabricCustomVariablesOk returns a tuple with the FabricCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetFabricCustomVariables(v map[string]interface{})`

SetFabricCustomVariables sets FabricCustomVariables field to given value.

### HasFabricCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) HasFabricCustomVariables() bool`

HasFabricCustomVariables returns a boolean if a field has been set.

### GetLinkCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLinkCustomVariables() map[string]interface{}`

GetLinkCustomVariables returns the LinkCustomVariables field if non-nil, zero value otherwise.

### GetLinkCustomVariablesOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLinkCustomVariablesOk() (*map[string]interface{}, bool)`

GetLinkCustomVariablesOk returns a tuple with the LinkCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLinkCustomVariables(v map[string]interface{})`

SetLinkCustomVariables sets LinkCustomVariables field to given value.

### HasLinkCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) HasLinkCustomVariables() bool`

HasLinkCustomVariables returns a boolean if a field has been set.

### GetLocalSwitchCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchCustomVariables() map[string]interface{}`

GetLocalSwitchCustomVariables returns the LocalSwitchCustomVariables field if non-nil, zero value otherwise.

### GetLocalSwitchCustomVariablesOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetLocalSwitchCustomVariablesOk() (*map[string]interface{}, bool)`

GetLocalSwitchCustomVariablesOk returns a tuple with the LocalSwitchCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSwitchCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetLocalSwitchCustomVariables(v map[string]interface{})`

SetLocalSwitchCustomVariables sets LocalSwitchCustomVariables field to given value.

### HasLocalSwitchCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) HasLocalSwitchCustomVariables() bool`

HasLocalSwitchCustomVariables returns a boolean if a field has been set.

### GetRemoteSwitchCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchCustomVariables() map[string]interface{}`

GetRemoteSwitchCustomVariables returns the RemoteSwitchCustomVariables field if non-nil, zero value otherwise.

### GetRemoteSwitchCustomVariablesOk

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) GetRemoteSwitchCustomVariablesOk() (*map[string]interface{}, bool)`

GetRemoteSwitchCustomVariablesOk returns a tuple with the RemoteSwitchCustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSwitchCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) SetRemoteSwitchCustomVariables(v map[string]interface{})`

SetRemoteSwitchCustomVariables sets RemoteSwitchCustomVariables field to given value.

### HasRemoteSwitchCustomVariables

`func (o *NetworkDeviceLinkAggregationConfigurationTemplateRecordSet) HasRemoteSwitchCustomVariables() bool`

HasRemoteSwitchCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



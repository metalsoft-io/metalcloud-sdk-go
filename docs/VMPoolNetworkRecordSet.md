# VMPoolNetworkRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmPoolId** | **float32** | The ID of the VM Pool. | 
**Hostname** | **string** | The hostname of the VM Pool. | 
**Operation** | **string** | The operation of the VM Pool. | 
**SwitchName** | Pointer to **string** | The switch name on which to operate (for VMware VDS). Null if not applicable. | [optional] 
**VcenterDatacenter** | Pointer to **string** | The vCenter datacenter name (for VMware). Null if not applicable. | [optional] 
**NetworkConfigurations** | [**[]VMPoolNetworkConfigurationRecordSet**](VMPoolNetworkConfigurationRecordSet.md) | The network configuration of the VM Pool. | 

## Methods

### NewVMPoolNetworkRecordSet

`func NewVMPoolNetworkRecordSet(vmPoolId float32, hostname string, operation string, networkConfigurations []VMPoolNetworkConfigurationRecordSet, ) *VMPoolNetworkRecordSet`

NewVMPoolNetworkRecordSet instantiates a new VMPoolNetworkRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolNetworkRecordSetWithDefaults

`func NewVMPoolNetworkRecordSetWithDefaults() *VMPoolNetworkRecordSet`

NewVMPoolNetworkRecordSetWithDefaults instantiates a new VMPoolNetworkRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmPoolId

`func (o *VMPoolNetworkRecordSet) GetVmPoolId() float32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMPoolNetworkRecordSet) GetVmPoolIdOk() (*float32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMPoolNetworkRecordSet) SetVmPoolId(v float32)`

SetVmPoolId sets VmPoolId field to given value.


### GetHostname

`func (o *VMPoolNetworkRecordSet) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VMPoolNetworkRecordSet) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VMPoolNetworkRecordSet) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetOperation

`func (o *VMPoolNetworkRecordSet) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VMPoolNetworkRecordSet) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VMPoolNetworkRecordSet) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetSwitchName

`func (o *VMPoolNetworkRecordSet) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *VMPoolNetworkRecordSet) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *VMPoolNetworkRecordSet) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *VMPoolNetworkRecordSet) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetVcenterDatacenter

`func (o *VMPoolNetworkRecordSet) GetVcenterDatacenter() string`

GetVcenterDatacenter returns the VcenterDatacenter field if non-nil, zero value otherwise.

### GetVcenterDatacenterOk

`func (o *VMPoolNetworkRecordSet) GetVcenterDatacenterOk() (*string, bool)`

GetVcenterDatacenterOk returns a tuple with the VcenterDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDatacenter

`func (o *VMPoolNetworkRecordSet) SetVcenterDatacenter(v string)`

SetVcenterDatacenter sets VcenterDatacenter field to given value.

### HasVcenterDatacenter

`func (o *VMPoolNetworkRecordSet) HasVcenterDatacenter() bool`

HasVcenterDatacenter returns a boolean if a field has been set.

### GetNetworkConfigurations

`func (o *VMPoolNetworkRecordSet) GetNetworkConfigurations() []VMPoolNetworkConfigurationRecordSet`

GetNetworkConfigurations returns the NetworkConfigurations field if non-nil, zero value otherwise.

### GetNetworkConfigurationsOk

`func (o *VMPoolNetworkRecordSet) GetNetworkConfigurationsOk() (*[]VMPoolNetworkConfigurationRecordSet, bool)`

GetNetworkConfigurationsOk returns a tuple with the NetworkConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigurations

`func (o *VMPoolNetworkRecordSet) SetNetworkConfigurations(v []VMPoolNetworkConfigurationRecordSet)`

SetNetworkConfigurations sets NetworkConfigurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



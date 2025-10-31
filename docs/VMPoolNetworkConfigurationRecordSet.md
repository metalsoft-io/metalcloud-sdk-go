# VMPoolNetworkConfigurationRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmNetName** | **string** | The name of the network. | 
**VlanId** | **float32** | The VLAN ID of the network. | 

## Methods

### NewVMPoolNetworkConfigurationRecordSet

`func NewVMPoolNetworkConfigurationRecordSet(vmNetName string, vlanId float32, ) *VMPoolNetworkConfigurationRecordSet`

NewVMPoolNetworkConfigurationRecordSet instantiates a new VMPoolNetworkConfigurationRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolNetworkConfigurationRecordSetWithDefaults

`func NewVMPoolNetworkConfigurationRecordSetWithDefaults() *VMPoolNetworkConfigurationRecordSet`

NewVMPoolNetworkConfigurationRecordSetWithDefaults instantiates a new VMPoolNetworkConfigurationRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmNetName

`func (o *VMPoolNetworkConfigurationRecordSet) GetVmNetName() string`

GetVmNetName returns the VmNetName field if non-nil, zero value otherwise.

### GetVmNetNameOk

`func (o *VMPoolNetworkConfigurationRecordSet) GetVmNetNameOk() (*string, bool)`

GetVmNetNameOk returns a tuple with the VmNetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetName

`func (o *VMPoolNetworkConfigurationRecordSet) SetVmNetName(v string)`

SetVmNetName sets VmNetName field to given value.


### GetVlanId

`func (o *VMPoolNetworkConfigurationRecordSet) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VMPoolNetworkConfigurationRecordSet) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VMPoolNetworkConfigurationRecordSet) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VMInstanceNetworkConfigurationRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmNetName** | **string** | The network name of the VM Instance. | 
**VcenterNetNameInternal** | Pointer to **string** | The vCenter internal network name of the VM Instance. | [optional] 

## Methods

### NewVMInstanceNetworkConfigurationRecordSet

`func NewVMInstanceNetworkConfigurationRecordSet(vmNetName string, ) *VMInstanceNetworkConfigurationRecordSet`

NewVMInstanceNetworkConfigurationRecordSet instantiates a new VMInstanceNetworkConfigurationRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceNetworkConfigurationRecordSetWithDefaults

`func NewVMInstanceNetworkConfigurationRecordSetWithDefaults() *VMInstanceNetworkConfigurationRecordSet`

NewVMInstanceNetworkConfigurationRecordSetWithDefaults instantiates a new VMInstanceNetworkConfigurationRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmNetName

`func (o *VMInstanceNetworkConfigurationRecordSet) GetVmNetName() string`

GetVmNetName returns the VmNetName field if non-nil, zero value otherwise.

### GetVmNetNameOk

`func (o *VMInstanceNetworkConfigurationRecordSet) GetVmNetNameOk() (*string, bool)`

GetVmNetNameOk returns a tuple with the VmNetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetName

`func (o *VMInstanceNetworkConfigurationRecordSet) SetVmNetName(v string)`

SetVmNetName sets VmNetName field to given value.


### GetVcenterNetNameInternal

`func (o *VMInstanceNetworkConfigurationRecordSet) GetVcenterNetNameInternal() string`

GetVcenterNetNameInternal returns the VcenterNetNameInternal field if non-nil, zero value otherwise.

### GetVcenterNetNameInternalOk

`func (o *VMInstanceNetworkConfigurationRecordSet) GetVcenterNetNameInternalOk() (*string, bool)`

GetVcenterNetNameInternalOk returns a tuple with the VcenterNetNameInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterNetNameInternal

`func (o *VMInstanceNetworkConfigurationRecordSet) SetVcenterNetNameInternal(v string)`

SetVcenterNetNameInternal sets VcenterNetNameInternal field to given value.

### HasVcenterNetNameInternal

`func (o *VMInstanceNetworkConfigurationRecordSet) HasVcenterNetNameInternal() bool`

HasVcenterNetNameInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



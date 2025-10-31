# VMInstanceRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmPoolId** | **float32** | The ID of the VM Pool. | 
**Hostname** | **string** | The hostname of the VM Pool. | 
**Username** | Pointer to **string** | The username for the VM Instance. | [optional] 
**Password** | Pointer to **string** | The password for the VM Instance. | [optional] 
**Certificate** | Pointer to **string** | The certificate for the VM Instance. | [optional] 
**PrivateKey** | Pointer to **string** | The private key for the VM Instance. | [optional] 
**Operation** | **string** | The operation to be performed on the VM Instance. | 
**VmName** | **string** | The name of the VM Instance. | 
**VmDiskGb** | **float32** | The disk size of the VM Instance in GB. | 
**VmCpuCores** | **float32** | The number of CPU cores for the VM Instance. | 
**VmRamGb** | **float32** | The RAM size of the VM Instance in GB. | 
**VmTemplate** | **string** | The template used by the VM Instance. | 
**NetworkConfigurations** | Pointer to [**[]VMInstanceNetworkConfigurationRecordSet**](VMInstanceNetworkConfigurationRecordSet.md) | The network configuration of the VM Instance. | [optional] 
**VcenterDatacenter** | Pointer to **string** | The vCenter datacenter for the VM Instance. | [optional] 
**VcenterCluster** | Pointer to **string** | The vCenter cluster for the VM Instance. | [optional] 
**VcenterVmDestinationFolder** | Pointer to **string** | The vCenter VM destination folder for the VM Instance. | [optional] 
**VcenterDiskDatastore** | Pointer to **string** | The vCenter datastore name for the VM Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables from the VM instance group | [optional] 

## Methods

### NewVMInstanceRecordSet

`func NewVMInstanceRecordSet(vmPoolId float32, hostname string, operation string, vmName string, vmDiskGb float32, vmCpuCores float32, vmRamGb float32, vmTemplate string, ) *VMInstanceRecordSet`

NewVMInstanceRecordSet instantiates a new VMInstanceRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceRecordSetWithDefaults

`func NewVMInstanceRecordSetWithDefaults() *VMInstanceRecordSet`

NewVMInstanceRecordSetWithDefaults instantiates a new VMInstanceRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmPoolId

`func (o *VMInstanceRecordSet) GetVmPoolId() float32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMInstanceRecordSet) GetVmPoolIdOk() (*float32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMInstanceRecordSet) SetVmPoolId(v float32)`

SetVmPoolId sets VmPoolId field to given value.


### GetHostname

`func (o *VMInstanceRecordSet) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VMInstanceRecordSet) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VMInstanceRecordSet) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetUsername

`func (o *VMInstanceRecordSet) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VMInstanceRecordSet) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VMInstanceRecordSet) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VMInstanceRecordSet) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *VMInstanceRecordSet) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VMInstanceRecordSet) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VMInstanceRecordSet) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VMInstanceRecordSet) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCertificate

`func (o *VMInstanceRecordSet) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *VMInstanceRecordSet) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *VMInstanceRecordSet) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *VMInstanceRecordSet) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *VMInstanceRecordSet) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VMInstanceRecordSet) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VMInstanceRecordSet) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *VMInstanceRecordSet) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetOperation

`func (o *VMInstanceRecordSet) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VMInstanceRecordSet) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VMInstanceRecordSet) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetVmName

`func (o *VMInstanceRecordSet) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *VMInstanceRecordSet) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *VMInstanceRecordSet) SetVmName(v string)`

SetVmName sets VmName field to given value.


### GetVmDiskGb

`func (o *VMInstanceRecordSet) GetVmDiskGb() float32`

GetVmDiskGb returns the VmDiskGb field if non-nil, zero value otherwise.

### GetVmDiskGbOk

`func (o *VMInstanceRecordSet) GetVmDiskGbOk() (*float32, bool)`

GetVmDiskGbOk returns a tuple with the VmDiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDiskGb

`func (o *VMInstanceRecordSet) SetVmDiskGb(v float32)`

SetVmDiskGb sets VmDiskGb field to given value.


### GetVmCpuCores

`func (o *VMInstanceRecordSet) GetVmCpuCores() float32`

GetVmCpuCores returns the VmCpuCores field if non-nil, zero value otherwise.

### GetVmCpuCoresOk

`func (o *VMInstanceRecordSet) GetVmCpuCoresOk() (*float32, bool)`

GetVmCpuCoresOk returns a tuple with the VmCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpuCores

`func (o *VMInstanceRecordSet) SetVmCpuCores(v float32)`

SetVmCpuCores sets VmCpuCores field to given value.


### GetVmRamGb

`func (o *VMInstanceRecordSet) GetVmRamGb() float32`

GetVmRamGb returns the VmRamGb field if non-nil, zero value otherwise.

### GetVmRamGbOk

`func (o *VMInstanceRecordSet) GetVmRamGbOk() (*float32, bool)`

GetVmRamGbOk returns a tuple with the VmRamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRamGb

`func (o *VMInstanceRecordSet) SetVmRamGb(v float32)`

SetVmRamGb sets VmRamGb field to given value.


### GetVmTemplate

`func (o *VMInstanceRecordSet) GetVmTemplate() string`

GetVmTemplate returns the VmTemplate field if non-nil, zero value otherwise.

### GetVmTemplateOk

`func (o *VMInstanceRecordSet) GetVmTemplateOk() (*string, bool)`

GetVmTemplateOk returns a tuple with the VmTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplate

`func (o *VMInstanceRecordSet) SetVmTemplate(v string)`

SetVmTemplate sets VmTemplate field to given value.


### GetNetworkConfigurations

`func (o *VMInstanceRecordSet) GetNetworkConfigurations() []VMInstanceNetworkConfigurationRecordSet`

GetNetworkConfigurations returns the NetworkConfigurations field if non-nil, zero value otherwise.

### GetNetworkConfigurationsOk

`func (o *VMInstanceRecordSet) GetNetworkConfigurationsOk() (*[]VMInstanceNetworkConfigurationRecordSet, bool)`

GetNetworkConfigurationsOk returns a tuple with the NetworkConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfigurations

`func (o *VMInstanceRecordSet) SetNetworkConfigurations(v []VMInstanceNetworkConfigurationRecordSet)`

SetNetworkConfigurations sets NetworkConfigurations field to given value.

### HasNetworkConfigurations

`func (o *VMInstanceRecordSet) HasNetworkConfigurations() bool`

HasNetworkConfigurations returns a boolean if a field has been set.

### GetVcenterDatacenter

`func (o *VMInstanceRecordSet) GetVcenterDatacenter() string`

GetVcenterDatacenter returns the VcenterDatacenter field if non-nil, zero value otherwise.

### GetVcenterDatacenterOk

`func (o *VMInstanceRecordSet) GetVcenterDatacenterOk() (*string, bool)`

GetVcenterDatacenterOk returns a tuple with the VcenterDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDatacenter

`func (o *VMInstanceRecordSet) SetVcenterDatacenter(v string)`

SetVcenterDatacenter sets VcenterDatacenter field to given value.

### HasVcenterDatacenter

`func (o *VMInstanceRecordSet) HasVcenterDatacenter() bool`

HasVcenterDatacenter returns a boolean if a field has been set.

### GetVcenterCluster

`func (o *VMInstanceRecordSet) GetVcenterCluster() string`

GetVcenterCluster returns the VcenterCluster field if non-nil, zero value otherwise.

### GetVcenterClusterOk

`func (o *VMInstanceRecordSet) GetVcenterClusterOk() (*string, bool)`

GetVcenterClusterOk returns a tuple with the VcenterCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterCluster

`func (o *VMInstanceRecordSet) SetVcenterCluster(v string)`

SetVcenterCluster sets VcenterCluster field to given value.

### HasVcenterCluster

`func (o *VMInstanceRecordSet) HasVcenterCluster() bool`

HasVcenterCluster returns a boolean if a field has been set.

### GetVcenterVmDestinationFolder

`func (o *VMInstanceRecordSet) GetVcenterVmDestinationFolder() string`

GetVcenterVmDestinationFolder returns the VcenterVmDestinationFolder field if non-nil, zero value otherwise.

### GetVcenterVmDestinationFolderOk

`func (o *VMInstanceRecordSet) GetVcenterVmDestinationFolderOk() (*string, bool)`

GetVcenterVmDestinationFolderOk returns a tuple with the VcenterVmDestinationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterVmDestinationFolder

`func (o *VMInstanceRecordSet) SetVcenterVmDestinationFolder(v string)`

SetVcenterVmDestinationFolder sets VcenterVmDestinationFolder field to given value.

### HasVcenterVmDestinationFolder

`func (o *VMInstanceRecordSet) HasVcenterVmDestinationFolder() bool`

HasVcenterVmDestinationFolder returns a boolean if a field has been set.

### GetVcenterDiskDatastore

`func (o *VMInstanceRecordSet) GetVcenterDiskDatastore() string`

GetVcenterDiskDatastore returns the VcenterDiskDatastore field if non-nil, zero value otherwise.

### GetVcenterDiskDatastoreOk

`func (o *VMInstanceRecordSet) GetVcenterDiskDatastoreOk() (*string, bool)`

GetVcenterDiskDatastoreOk returns a tuple with the VcenterDiskDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDiskDatastore

`func (o *VMInstanceRecordSet) SetVcenterDiskDatastore(v string)`

SetVcenterDiskDatastore sets VcenterDiskDatastore field to given value.

### HasVcenterDiskDatastore

`func (o *VMInstanceRecordSet) HasVcenterDiskDatastore() bool`

HasVcenterDiskDatastore returns a boolean if a field has been set.

### GetCustomVariables

`func (o *VMInstanceRecordSet) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *VMInstanceRecordSet) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *VMInstanceRecordSet) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *VMInstanceRecordSet) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



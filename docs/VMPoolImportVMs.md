# VMPoolImportVMs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmNames** | **[]string** | List of VM names to import | 
**InfrastructureId** | Pointer to **float32** | If provided, the VMs will be added into this infrastructure | [optional] 

## Methods

### NewVMPoolImportVMs

`func NewVMPoolImportVMs(vmNames []string, ) *VMPoolImportVMs`

NewVMPoolImportVMs instantiates a new VMPoolImportVMs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolImportVMsWithDefaults

`func NewVMPoolImportVMsWithDefaults() *VMPoolImportVMs`

NewVMPoolImportVMsWithDefaults instantiates a new VMPoolImportVMs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmNames

`func (o *VMPoolImportVMs) GetVmNames() []string`

GetVmNames returns the VmNames field if non-nil, zero value otherwise.

### GetVmNamesOk

`func (o *VMPoolImportVMs) GetVmNamesOk() (*[]string, bool)`

GetVmNamesOk returns a tuple with the VmNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNames

`func (o *VMPoolImportVMs) SetVmNames(v []string)`

SetVmNames sets VmNames field to given value.


### GetInfrastructureId

`func (o *VMPoolImportVMs) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VMPoolImportVMs) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VMPoolImportVMs) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *VMPoolImportVMs) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VMPoolGPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID of the GPU | 
**VmPoolId** | **float32** | ID of the VM Pool linked to the GPU | 
**Model** | Pointer to **string** | Model of the GPU | [optional] 
**Vendor** | Pointer to **string** | Vendor of the GPU | [optional] 
**PciAddress** | Pointer to **string** | PCI address of the GPU | [optional] 

## Methods

### NewVMPoolGPU

`func NewVMPoolGPU(id float32, vmPoolId float32, ) *VMPoolGPU`

NewVMPoolGPU instantiates a new VMPoolGPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolGPUWithDefaults

`func NewVMPoolGPUWithDefaults() *VMPoolGPU`

NewVMPoolGPUWithDefaults instantiates a new VMPoolGPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPoolGPU) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPoolGPU) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPoolGPU) SetId(v float32)`

SetId sets Id field to given value.


### GetVmPoolId

`func (o *VMPoolGPU) GetVmPoolId() float32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMPoolGPU) GetVmPoolIdOk() (*float32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMPoolGPU) SetVmPoolId(v float32)`

SetVmPoolId sets VmPoolId field to given value.


### GetModel

`func (o *VMPoolGPU) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VMPoolGPU) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VMPoolGPU) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VMPoolGPU) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *VMPoolGPU) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VMPoolGPU) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VMPoolGPU) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VMPoolGPU) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetPciAddress

`func (o *VMPoolGPU) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *VMPoolGPU) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *VMPoolGPU) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *VMPoolGPU) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



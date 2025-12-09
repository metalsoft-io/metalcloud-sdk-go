# VMGpuInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the GPU | 
**Vendor** | **string** | The vendor of the GPU | 
**PciAddress** | **string** | The PCI address of the GPU | 

## Methods

### NewVMGpuInfoDto

`func NewVMGpuInfoDto(name string, vendor string, pciAddress string, ) *VMGpuInfoDto`

NewVMGpuInfoDto instantiates a new VMGpuInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMGpuInfoDtoWithDefaults

`func NewVMGpuInfoDtoWithDefaults() *VMGpuInfoDto`

NewVMGpuInfoDtoWithDefaults instantiates a new VMGpuInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VMGpuInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMGpuInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMGpuInfoDto) SetName(v string)`

SetName sets Name field to given value.


### GetVendor

`func (o *VMGpuInfoDto) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VMGpuInfoDto) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VMGpuInfoDto) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetPciAddress

`func (o *VMGpuInfoDto) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *VMGpuInfoDto) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *VMGpuInfoDto) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



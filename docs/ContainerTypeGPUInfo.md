# ContainerTypeGPUInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the GPU as it appears in the PCI device list | 
**Count** | **float32** | Number of GPU units available for this Container Type | 

## Methods

### NewContainerTypeGPUInfo

`func NewContainerTypeGPUInfo(name string, count float32, ) *ContainerTypeGPUInfo`

NewContainerTypeGPUInfo instantiates a new ContainerTypeGPUInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTypeGPUInfoWithDefaults

`func NewContainerTypeGPUInfoWithDefaults() *ContainerTypeGPUInfo`

NewContainerTypeGPUInfoWithDefaults instantiates a new ContainerTypeGPUInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerTypeGPUInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerTypeGPUInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerTypeGPUInfo) SetName(v string)`

SetName sets Name field to given value.


### GetCount

`func (o *ContainerTypeGPUInfo) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ContainerTypeGPUInfo) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ContainerTypeGPUInfo) SetCount(v float32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



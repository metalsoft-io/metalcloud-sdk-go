# ServerGpuInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | The model of the GPU | 
**Vendor** | **string** | The vendor of the GPU | 

## Methods

### NewServerGpuInfo

`func NewServerGpuInfo(model string, vendor string, ) *ServerGpuInfo`

NewServerGpuInfo instantiates a new ServerGpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGpuInfoWithDefaults

`func NewServerGpuInfoWithDefaults() *ServerGpuInfo`

NewServerGpuInfoWithDefaults instantiates a new ServerGpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ServerGpuInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServerGpuInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServerGpuInfo) SetModel(v string)`

SetModel sets Model field to given value.


### GetVendor

`func (o *ServerGpuInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerGpuInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerGpuInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



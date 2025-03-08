# ExtensionInputOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationRegEx** | Pointer to **string** | Validation regex pattern for the input. | [optional] 
**MinValue** | Pointer to **int32** | Minimum allowed value. | [optional] 
**MaxValue** | Pointer to **int32** | Maximum allowed value. | [optional] 
**DeniedValues** | Pointer to **[]int32** | Denied values. | [optional] 
**MinCpu** | Pointer to **int32** | Minimum CPU cores. | [optional] 
**MinRamGb** | Pointer to **int32** | Minimum RAM GBs. | [optional] 
**Vendor** | Pointer to **string** | Server vendor. | [optional] 
**OsFamily** | Pointer to **string** | OS family type. | [optional] 

## Methods

### NewExtensionInputOptions

`func NewExtensionInputOptions() *ExtensionInputOptions`

NewExtensionInputOptions instantiates a new ExtensionInputOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInputOptionsWithDefaults

`func NewExtensionInputOptionsWithDefaults() *ExtensionInputOptions`

NewExtensionInputOptionsWithDefaults instantiates a new ExtensionInputOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationRegEx

`func (o *ExtensionInputOptions) GetValidationRegEx() string`

GetValidationRegEx returns the ValidationRegEx field if non-nil, zero value otherwise.

### GetValidationRegExOk

`func (o *ExtensionInputOptions) GetValidationRegExOk() (*string, bool)`

GetValidationRegExOk returns a tuple with the ValidationRegEx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegEx

`func (o *ExtensionInputOptions) SetValidationRegEx(v string)`

SetValidationRegEx sets ValidationRegEx field to given value.

### HasValidationRegEx

`func (o *ExtensionInputOptions) HasValidationRegEx() bool`

HasValidationRegEx returns a boolean if a field has been set.

### GetMinValue

`func (o *ExtensionInputOptions) GetMinValue() int32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ExtensionInputOptions) GetMinValueOk() (*int32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ExtensionInputOptions) SetMinValue(v int32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ExtensionInputOptions) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *ExtensionInputOptions) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ExtensionInputOptions) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ExtensionInputOptions) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ExtensionInputOptions) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetDeniedValues

`func (o *ExtensionInputOptions) GetDeniedValues() []int32`

GetDeniedValues returns the DeniedValues field if non-nil, zero value otherwise.

### GetDeniedValuesOk

`func (o *ExtensionInputOptions) GetDeniedValuesOk() (*[]int32, bool)`

GetDeniedValuesOk returns a tuple with the DeniedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedValues

`func (o *ExtensionInputOptions) SetDeniedValues(v []int32)`

SetDeniedValues sets DeniedValues field to given value.

### HasDeniedValues

`func (o *ExtensionInputOptions) HasDeniedValues() bool`

HasDeniedValues returns a boolean if a field has been set.

### GetMinCpu

`func (o *ExtensionInputOptions) GetMinCpu() int32`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *ExtensionInputOptions) GetMinCpuOk() (*int32, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *ExtensionInputOptions) SetMinCpu(v int32)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *ExtensionInputOptions) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMinRamGb

`func (o *ExtensionInputOptions) GetMinRamGb() int32`

GetMinRamGb returns the MinRamGb field if non-nil, zero value otherwise.

### GetMinRamGbOk

`func (o *ExtensionInputOptions) GetMinRamGbOk() (*int32, bool)`

GetMinRamGbOk returns a tuple with the MinRamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamGb

`func (o *ExtensionInputOptions) SetMinRamGb(v int32)`

SetMinRamGb sets MinRamGb field to given value.

### HasMinRamGb

`func (o *ExtensionInputOptions) HasMinRamGb() bool`

HasMinRamGb returns a boolean if a field has been set.

### GetVendor

`func (o *ExtensionInputOptions) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ExtensionInputOptions) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ExtensionInputOptions) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ExtensionInputOptions) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetOsFamily

`func (o *ExtensionInputOptions) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ExtensionInputOptions) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ExtensionInputOptions) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *ExtensionInputOptions) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



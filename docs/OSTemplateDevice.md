# OSTemplateDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The OS template device type | 
**BootMode** | **string** | The hardware boot mode supported by the OS template | 
**Architecture** | **string** | The hardware architecture supported by the OS template | 

## Methods

### NewOSTemplateDevice

`func NewOSTemplateDevice(type_ string, bootMode string, architecture string, ) *OSTemplateDevice`

NewOSTemplateDevice instantiates a new OSTemplateDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateDeviceWithDefaults

`func NewOSTemplateDeviceWithDefaults() *OSTemplateDevice`

NewOSTemplateDeviceWithDefaults instantiates a new OSTemplateDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OSTemplateDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OSTemplateDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OSTemplateDevice) SetType(v string)`

SetType sets Type field to given value.


### GetBootMode

`func (o *OSTemplateDevice) GetBootMode() string`

GetBootMode returns the BootMode field if non-nil, zero value otherwise.

### GetBootModeOk

`func (o *OSTemplateDevice) GetBootModeOk() (*string, bool)`

GetBootModeOk returns a tuple with the BootMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootMode

`func (o *OSTemplateDevice) SetBootMode(v string)`

SetBootMode sets BootMode field to given value.


### GetArchitecture

`func (o *OSTemplateDevice) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *OSTemplateDevice) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *OSTemplateDevice) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



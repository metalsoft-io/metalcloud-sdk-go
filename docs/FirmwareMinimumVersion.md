# FirmwareMinimumVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | **string** | Name of the firmware component (e.g., BMC, RAID Controller, BIOS) | 
**MinimumVersion** | **string** | Minimum required version for this component | 

## Methods

### NewFirmwareMinimumVersion

`func NewFirmwareMinimumVersion(componentName string, minimumVersion string, ) *FirmwareMinimumVersion`

NewFirmwareMinimumVersion instantiates a new FirmwareMinimumVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareMinimumVersionWithDefaults

`func NewFirmwareMinimumVersionWithDefaults() *FirmwareMinimumVersion`

NewFirmwareMinimumVersionWithDefaults instantiates a new FirmwareMinimumVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *FirmwareMinimumVersion) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *FirmwareMinimumVersion) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *FirmwareMinimumVersion) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetMinimumVersion

`func (o *FirmwareMinimumVersion) GetMinimumVersion() string`

GetMinimumVersion returns the MinimumVersion field if non-nil, zero value otherwise.

### GetMinimumVersionOk

`func (o *FirmwareMinimumVersion) GetMinimumVersionOk() (*string, bool)`

GetMinimumVersionOk returns a tuple with the MinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumVersion

`func (o *FirmwareMinimumVersion) SetMinimumVersion(v string)`

SetMinimumVersion sets MinimumVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



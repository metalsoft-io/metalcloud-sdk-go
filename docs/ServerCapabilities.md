# ServerCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirmwareUpgradeSupported** | **bool** | Server supports firmware upgrades | 
**FirmwareUpgradeApplyOnRebootSupported** | **bool** | Server supports firmware upgrades on reboot | 
**VirtualMediaDeviceCount** | **float32** | The number of virtual media devices supported by the server | 
**VncEnabled** | **bool** | Server has VNC enabled | 

## Methods

### NewServerCapabilities

`func NewServerCapabilities(firmwareUpgradeSupported bool, firmwareUpgradeApplyOnRebootSupported bool, virtualMediaDeviceCount float32, vncEnabled bool, ) *ServerCapabilities`

NewServerCapabilities instantiates a new ServerCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCapabilitiesWithDefaults

`func NewServerCapabilitiesWithDefaults() *ServerCapabilities`

NewServerCapabilitiesWithDefaults instantiates a new ServerCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirmwareUpgradeSupported

`func (o *ServerCapabilities) GetFirmwareUpgradeSupported() bool`

GetFirmwareUpgradeSupported returns the FirmwareUpgradeSupported field if non-nil, zero value otherwise.

### GetFirmwareUpgradeSupportedOk

`func (o *ServerCapabilities) GetFirmwareUpgradeSupportedOk() (*bool, bool)`

GetFirmwareUpgradeSupportedOk returns a tuple with the FirmwareUpgradeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpgradeSupported

`func (o *ServerCapabilities) SetFirmwareUpgradeSupported(v bool)`

SetFirmwareUpgradeSupported sets FirmwareUpgradeSupported field to given value.


### GetFirmwareUpgradeApplyOnRebootSupported

`func (o *ServerCapabilities) GetFirmwareUpgradeApplyOnRebootSupported() bool`

GetFirmwareUpgradeApplyOnRebootSupported returns the FirmwareUpgradeApplyOnRebootSupported field if non-nil, zero value otherwise.

### GetFirmwareUpgradeApplyOnRebootSupportedOk

`func (o *ServerCapabilities) GetFirmwareUpgradeApplyOnRebootSupportedOk() (*bool, bool)`

GetFirmwareUpgradeApplyOnRebootSupportedOk returns a tuple with the FirmwareUpgradeApplyOnRebootSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpgradeApplyOnRebootSupported

`func (o *ServerCapabilities) SetFirmwareUpgradeApplyOnRebootSupported(v bool)`

SetFirmwareUpgradeApplyOnRebootSupported sets FirmwareUpgradeApplyOnRebootSupported field to given value.


### GetVirtualMediaDeviceCount

`func (o *ServerCapabilities) GetVirtualMediaDeviceCount() float32`

GetVirtualMediaDeviceCount returns the VirtualMediaDeviceCount field if non-nil, zero value otherwise.

### GetVirtualMediaDeviceCountOk

`func (o *ServerCapabilities) GetVirtualMediaDeviceCountOk() (*float32, bool)`

GetVirtualMediaDeviceCountOk returns a tuple with the VirtualMediaDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMediaDeviceCount

`func (o *ServerCapabilities) SetVirtualMediaDeviceCount(v float32)`

SetVirtualMediaDeviceCount sets VirtualMediaDeviceCount field to given value.


### GetVncEnabled

`func (o *ServerCapabilities) GetVncEnabled() bool`

GetVncEnabled returns the VncEnabled field if non-nil, zero value otherwise.

### GetVncEnabledOk

`func (o *ServerCapabilities) GetVncEnabledOk() (*bool, bool)`

GetVncEnabledOk returns a tuple with the VncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncEnabled

`func (o *ServerCapabilities) SetVncEnabled(v bool)`

SetVncEnabled sets VncEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



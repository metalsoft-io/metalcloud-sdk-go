# LicenseAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | **int32** | Number of devices allowed by the license | 
**Switches** | **int32** | Number of switches allowed by the license | 
**Servers** | **int32** | Number of servers allowed by the license | 
**StorageGB** | **int32** | Amount of storage in GB allowed by the license | 

## Methods

### NewLicenseAllowance

`func NewLicenseAllowance(devices int32, switches int32, servers int32, storageGB int32, ) *LicenseAllowance`

NewLicenseAllowance instantiates a new LicenseAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAllowanceWithDefaults

`func NewLicenseAllowanceWithDefaults() *LicenseAllowance`

NewLicenseAllowanceWithDefaults instantiates a new LicenseAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *LicenseAllowance) GetDevices() int32`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *LicenseAllowance) GetDevicesOk() (*int32, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *LicenseAllowance) SetDevices(v int32)`

SetDevices sets Devices field to given value.


### GetSwitches

`func (o *LicenseAllowance) GetSwitches() int32`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *LicenseAllowance) GetSwitchesOk() (*int32, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *LicenseAllowance) SetSwitches(v int32)`

SetSwitches sets Switches field to given value.


### GetServers

`func (o *LicenseAllowance) GetServers() int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *LicenseAllowance) GetServersOk() (*int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *LicenseAllowance) SetServers(v int32)`

SetServers sets Servers field to given value.


### GetStorageGB

`func (o *LicenseAllowance) GetStorageGB() int32`

GetStorageGB returns the StorageGB field if non-nil, zero value otherwise.

### GetStorageGBOk

`func (o *LicenseAllowance) GetStorageGBOk() (*int32, bool)`

GetStorageGBOk returns a tuple with the StorageGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGB

`func (o *LicenseAllowance) SetStorageGB(v int32)`

SetStorageGB sets StorageGB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



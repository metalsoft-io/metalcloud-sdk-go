# ServerInstanceStorageProfileControllerVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerName** | **string** | The name of the controller. | 
**VolumeName** | **string** | The name of the volume. | 
**DiskSizeGb** | **float32** | The size of the disk in GB. | 
**DiskType** | **string** | The type of the disk. | 
**DiskCount** | **float32** | The number of disks. | 
**RaidType** | **string** | The RAID type of the volume. | 

## Methods

### NewServerInstanceStorageProfileControllerVolume

`func NewServerInstanceStorageProfileControllerVolume(controllerName string, volumeName string, diskSizeGb float32, diskType string, diskCount float32, raidType string, ) *ServerInstanceStorageProfileControllerVolume`

NewServerInstanceStorageProfileControllerVolume instantiates a new ServerInstanceStorageProfileControllerVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceStorageProfileControllerVolumeWithDefaults

`func NewServerInstanceStorageProfileControllerVolumeWithDefaults() *ServerInstanceStorageProfileControllerVolume`

NewServerInstanceStorageProfileControllerVolumeWithDefaults instantiates a new ServerInstanceStorageProfileControllerVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerName

`func (o *ServerInstanceStorageProfileControllerVolume) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *ServerInstanceStorageProfileControllerVolume) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *ServerInstanceStorageProfileControllerVolume) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.


### GetVolumeName

`func (o *ServerInstanceStorageProfileControllerVolume) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *ServerInstanceStorageProfileControllerVolume) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *ServerInstanceStorageProfileControllerVolume) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.


### GetDiskSizeGb

`func (o *ServerInstanceStorageProfileControllerVolume) GetDiskSizeGb() float32`

GetDiskSizeGb returns the DiskSizeGb field if non-nil, zero value otherwise.

### GetDiskSizeGbOk

`func (o *ServerInstanceStorageProfileControllerVolume) GetDiskSizeGbOk() (*float32, bool)`

GetDiskSizeGbOk returns a tuple with the DiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGb

`func (o *ServerInstanceStorageProfileControllerVolume) SetDiskSizeGb(v float32)`

SetDiskSizeGb sets DiskSizeGb field to given value.


### GetDiskType

`func (o *ServerInstanceStorageProfileControllerVolume) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *ServerInstanceStorageProfileControllerVolume) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *ServerInstanceStorageProfileControllerVolume) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.


### GetDiskCount

`func (o *ServerInstanceStorageProfileControllerVolume) GetDiskCount() float32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceStorageProfileControllerVolume) GetDiskCountOk() (*float32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceStorageProfileControllerVolume) SetDiskCount(v float32)`

SetDiskCount sets DiskCount field to given value.


### GetRaidType

`func (o *ServerInstanceStorageProfileControllerVolume) GetRaidType() string`

GetRaidType returns the RaidType field if non-nil, zero value otherwise.

### GetRaidTypeOk

`func (o *ServerInstanceStorageProfileControllerVolume) GetRaidTypeOk() (*string, bool)`

GetRaidTypeOk returns a tuple with the RaidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidType

`func (o *ServerInstanceStorageProfileControllerVolume) SetRaidType(v string)`

SetRaidType sets RaidType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



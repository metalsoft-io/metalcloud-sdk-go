# UpdateStorageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDataReduction** | Pointer to **bool** | Enable data reduction | [optional] 
**EnableAdvancedDeduplication** | Pointer to **bool** | Enable advanced deduplication | [optional] 
**EnableDataReductionSharedVolume** | Pointer to **bool** | Enable data reduction on a newly created shared volume | [optional] 
**VolumeName** | Pointer to **string** | Volume name | [optional] 
**ArrayId** | Pointer to **string** | Array id to use (for certain storage drivers) | [optional] 
**DirectorId** | Pointer to **string** | Director id to use (for certain storage drivers) | [optional] 
**S3Hostname** | Pointer to **string** | S3 Hostname to use (for certain storage drivers) | [optional] 
**S3Port** | Pointer to **float32** | Enable advanced deduplication | [optional] 
**FibreChannelEnabled** | Pointer to **bool** | Fibre channel enabled | [optional] 

## Methods

### NewUpdateStorageOptions

`func NewUpdateStorageOptions() *UpdateStorageOptions`

NewUpdateStorageOptions instantiates a new UpdateStorageOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageOptionsWithDefaults

`func NewUpdateStorageOptionsWithDefaults() *UpdateStorageOptions`

NewUpdateStorageOptionsWithDefaults instantiates a new UpdateStorageOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableDataReduction

`func (o *UpdateStorageOptions) GetEnableDataReduction() bool`

GetEnableDataReduction returns the EnableDataReduction field if non-nil, zero value otherwise.

### GetEnableDataReductionOk

`func (o *UpdateStorageOptions) GetEnableDataReductionOk() (*bool, bool)`

GetEnableDataReductionOk returns a tuple with the EnableDataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDataReduction

`func (o *UpdateStorageOptions) SetEnableDataReduction(v bool)`

SetEnableDataReduction sets EnableDataReduction field to given value.

### HasEnableDataReduction

`func (o *UpdateStorageOptions) HasEnableDataReduction() bool`

HasEnableDataReduction returns a boolean if a field has been set.

### GetEnableAdvancedDeduplication

`func (o *UpdateStorageOptions) GetEnableAdvancedDeduplication() bool`

GetEnableAdvancedDeduplication returns the EnableAdvancedDeduplication field if non-nil, zero value otherwise.

### GetEnableAdvancedDeduplicationOk

`func (o *UpdateStorageOptions) GetEnableAdvancedDeduplicationOk() (*bool, bool)`

GetEnableAdvancedDeduplicationOk returns a tuple with the EnableAdvancedDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdvancedDeduplication

`func (o *UpdateStorageOptions) SetEnableAdvancedDeduplication(v bool)`

SetEnableAdvancedDeduplication sets EnableAdvancedDeduplication field to given value.

### HasEnableAdvancedDeduplication

`func (o *UpdateStorageOptions) HasEnableAdvancedDeduplication() bool`

HasEnableAdvancedDeduplication returns a boolean if a field has been set.

### GetEnableDataReductionSharedVolume

`func (o *UpdateStorageOptions) GetEnableDataReductionSharedVolume() bool`

GetEnableDataReductionSharedVolume returns the EnableDataReductionSharedVolume field if non-nil, zero value otherwise.

### GetEnableDataReductionSharedVolumeOk

`func (o *UpdateStorageOptions) GetEnableDataReductionSharedVolumeOk() (*bool, bool)`

GetEnableDataReductionSharedVolumeOk returns a tuple with the EnableDataReductionSharedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDataReductionSharedVolume

`func (o *UpdateStorageOptions) SetEnableDataReductionSharedVolume(v bool)`

SetEnableDataReductionSharedVolume sets EnableDataReductionSharedVolume field to given value.

### HasEnableDataReductionSharedVolume

`func (o *UpdateStorageOptions) HasEnableDataReductionSharedVolume() bool`

HasEnableDataReductionSharedVolume returns a boolean if a field has been set.

### GetVolumeName

`func (o *UpdateStorageOptions) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *UpdateStorageOptions) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *UpdateStorageOptions) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *UpdateStorageOptions) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetArrayId

`func (o *UpdateStorageOptions) GetArrayId() string`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *UpdateStorageOptions) GetArrayIdOk() (*string, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *UpdateStorageOptions) SetArrayId(v string)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *UpdateStorageOptions) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetDirectorId

`func (o *UpdateStorageOptions) GetDirectorId() string`

GetDirectorId returns the DirectorId field if non-nil, zero value otherwise.

### GetDirectorIdOk

`func (o *UpdateStorageOptions) GetDirectorIdOk() (*string, bool)`

GetDirectorIdOk returns a tuple with the DirectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorId

`func (o *UpdateStorageOptions) SetDirectorId(v string)`

SetDirectorId sets DirectorId field to given value.

### HasDirectorId

`func (o *UpdateStorageOptions) HasDirectorId() bool`

HasDirectorId returns a boolean if a field has been set.

### GetS3Hostname

`func (o *UpdateStorageOptions) GetS3Hostname() string`

GetS3Hostname returns the S3Hostname field if non-nil, zero value otherwise.

### GetS3HostnameOk

`func (o *UpdateStorageOptions) GetS3HostnameOk() (*string, bool)`

GetS3HostnameOk returns a tuple with the S3Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Hostname

`func (o *UpdateStorageOptions) SetS3Hostname(v string)`

SetS3Hostname sets S3Hostname field to given value.

### HasS3Hostname

`func (o *UpdateStorageOptions) HasS3Hostname() bool`

HasS3Hostname returns a boolean if a field has been set.

### GetS3Port

`func (o *UpdateStorageOptions) GetS3Port() float32`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *UpdateStorageOptions) GetS3PortOk() (*float32, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *UpdateStorageOptions) SetS3Port(v float32)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *UpdateStorageOptions) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetFibreChannelEnabled

`func (o *UpdateStorageOptions) GetFibreChannelEnabled() bool`

GetFibreChannelEnabled returns the FibreChannelEnabled field if non-nil, zero value otherwise.

### GetFibreChannelEnabledOk

`func (o *UpdateStorageOptions) GetFibreChannelEnabledOk() (*bool, bool)`

GetFibreChannelEnabledOk returns a tuple with the FibreChannelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelEnabled

`func (o *UpdateStorageOptions) SetFibreChannelEnabled(v bool)`

SetFibreChannelEnabled sets FibreChannelEnabled field to given value.

### HasFibreChannelEnabled

`func (o *UpdateStorageOptions) HasFibreChannelEnabled() bool`

HasFibreChannelEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PlatformStorageConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageOverprovisioningFactor** | Pointer to **float32** | Storage overprovisioning multiplier (e.g. 3000 &#x3D; 3x overprovisioning) | [optional] 

## Methods

### NewPlatformStorageConfigDto

`func NewPlatformStorageConfigDto() *PlatformStorageConfigDto`

NewPlatformStorageConfigDto instantiates a new PlatformStorageConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformStorageConfigDtoWithDefaults

`func NewPlatformStorageConfigDtoWithDefaults() *PlatformStorageConfigDto`

NewPlatformStorageConfigDtoWithDefaults instantiates a new PlatformStorageConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageOverprovisioningFactor

`func (o *PlatformStorageConfigDto) GetStorageOverprovisioningFactor() float32`

GetStorageOverprovisioningFactor returns the StorageOverprovisioningFactor field if non-nil, zero value otherwise.

### GetStorageOverprovisioningFactorOk

`func (o *PlatformStorageConfigDto) GetStorageOverprovisioningFactorOk() (*float32, bool)`

GetStorageOverprovisioningFactorOk returns a tuple with the StorageOverprovisioningFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOverprovisioningFactor

`func (o *PlatformStorageConfigDto) SetStorageOverprovisioningFactor(v float32)`

SetStorageOverprovisioningFactor sets StorageOverprovisioningFactor field to given value.

### HasStorageOverprovisioningFactor

`func (o *PlatformStorageConfigDto) HasStorageOverprovisioningFactor() bool`

HasStorageOverprovisioningFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



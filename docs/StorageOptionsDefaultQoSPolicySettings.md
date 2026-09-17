# StorageOptionsDefaultQoSPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxBandwidthMBpsLimit** | Pointer to **float32** | Max bandwidth limit in MBps for the resource. Set to 0 to leave it unlimited (default) | [optional] 
**MaxIopsLimit** | Pointer to **float32** | Max IOPs for the resource. Set to 0 to leave it unlimited (default) | [optional] 
**FloorBandwidthMBpsLimit** | Pointer to **float32** | Floor bandwidth limit in MBps for the resource (for certain storages). Set to 0 to leave it unlimited (default) | [optional] 
**FloorIopsLimit** | Pointer to **float32** | Floor IOPs for the resource (for certain storages). Set to 0 to leave it unlimited (default) | [optional] 

## Methods

### NewStorageOptionsDefaultQoSPolicySettings

`func NewStorageOptionsDefaultQoSPolicySettings() *StorageOptionsDefaultQoSPolicySettings`

NewStorageOptionsDefaultQoSPolicySettings instantiates a new StorageOptionsDefaultQoSPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsDefaultQoSPolicySettingsWithDefaults

`func NewStorageOptionsDefaultQoSPolicySettingsWithDefaults() *StorageOptionsDefaultQoSPolicySettings`

NewStorageOptionsDefaultQoSPolicySettingsWithDefaults instantiates a new StorageOptionsDefaultQoSPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxBandwidthMBpsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) GetMaxBandwidthMBpsLimit() float32`

GetMaxBandwidthMBpsLimit returns the MaxBandwidthMBpsLimit field if non-nil, zero value otherwise.

### GetMaxBandwidthMBpsLimitOk

`func (o *StorageOptionsDefaultQoSPolicySettings) GetMaxBandwidthMBpsLimitOk() (*float32, bool)`

GetMaxBandwidthMBpsLimitOk returns a tuple with the MaxBandwidthMBpsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBandwidthMBpsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) SetMaxBandwidthMBpsLimit(v float32)`

SetMaxBandwidthMBpsLimit sets MaxBandwidthMBpsLimit field to given value.

### HasMaxBandwidthMBpsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) HasMaxBandwidthMBpsLimit() bool`

HasMaxBandwidthMBpsLimit returns a boolean if a field has been set.

### GetMaxIopsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) GetMaxIopsLimit() float32`

GetMaxIopsLimit returns the MaxIopsLimit field if non-nil, zero value otherwise.

### GetMaxIopsLimitOk

`func (o *StorageOptionsDefaultQoSPolicySettings) GetMaxIopsLimitOk() (*float32, bool)`

GetMaxIopsLimitOk returns a tuple with the MaxIopsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIopsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) SetMaxIopsLimit(v float32)`

SetMaxIopsLimit sets MaxIopsLimit field to given value.

### HasMaxIopsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) HasMaxIopsLimit() bool`

HasMaxIopsLimit returns a boolean if a field has been set.

### GetFloorBandwidthMBpsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) GetFloorBandwidthMBpsLimit() float32`

GetFloorBandwidthMBpsLimit returns the FloorBandwidthMBpsLimit field if non-nil, zero value otherwise.

### GetFloorBandwidthMBpsLimitOk

`func (o *StorageOptionsDefaultQoSPolicySettings) GetFloorBandwidthMBpsLimitOk() (*float32, bool)`

GetFloorBandwidthMBpsLimitOk returns a tuple with the FloorBandwidthMBpsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorBandwidthMBpsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) SetFloorBandwidthMBpsLimit(v float32)`

SetFloorBandwidthMBpsLimit sets FloorBandwidthMBpsLimit field to given value.

### HasFloorBandwidthMBpsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) HasFloorBandwidthMBpsLimit() bool`

HasFloorBandwidthMBpsLimit returns a boolean if a field has been set.

### GetFloorIopsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) GetFloorIopsLimit() float32`

GetFloorIopsLimit returns the FloorIopsLimit field if non-nil, zero value otherwise.

### GetFloorIopsLimitOk

`func (o *StorageOptionsDefaultQoSPolicySettings) GetFloorIopsLimitOk() (*float32, bool)`

GetFloorIopsLimitOk returns a tuple with the FloorIopsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorIopsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) SetFloorIopsLimit(v float32)`

SetFloorIopsLimit sets FloorIopsLimit field to given value.

### HasFloorIopsLimit

`func (o *StorageOptionsDefaultQoSPolicySettings) HasFloorIopsLimit() bool`

HasFloorIopsLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



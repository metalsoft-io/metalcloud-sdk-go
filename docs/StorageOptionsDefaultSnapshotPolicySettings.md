# StorageOptionsDefaultSnapshotPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of an existing snapshot policy already defined on the storage system to use as-is. If set, &#x60;rules&#x60; below is ignored. If not set, a policy will be created from &#x60;rules&#x60;. | [optional] 
**Rules** | Pointer to [**[]StorageOptionsDefaultSnapshotPolicyRule**](StorageOptionsDefaultSnapshotPolicyRule.md) | Schedule rules to create a new snapshot policy from. Ignored if &#x60;name&#x60; is set. | [optional] 

## Methods

### NewStorageOptionsDefaultSnapshotPolicySettings

`func NewStorageOptionsDefaultSnapshotPolicySettings() *StorageOptionsDefaultSnapshotPolicySettings`

NewStorageOptionsDefaultSnapshotPolicySettings instantiates a new StorageOptionsDefaultSnapshotPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsDefaultSnapshotPolicySettingsWithDefaults

`func NewStorageOptionsDefaultSnapshotPolicySettingsWithDefaults() *StorageOptionsDefaultSnapshotPolicySettings`

NewStorageOptionsDefaultSnapshotPolicySettingsWithDefaults instantiates a new StorageOptionsDefaultSnapshotPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageOptionsDefaultSnapshotPolicySettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageOptionsDefaultSnapshotPolicySettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageOptionsDefaultSnapshotPolicySettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageOptionsDefaultSnapshotPolicySettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *StorageOptionsDefaultSnapshotPolicySettings) GetRules() []StorageOptionsDefaultSnapshotPolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *StorageOptionsDefaultSnapshotPolicySettings) GetRulesOk() (*[]StorageOptionsDefaultSnapshotPolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *StorageOptionsDefaultSnapshotPolicySettings) SetRules(v []StorageOptionsDefaultSnapshotPolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *StorageOptionsDefaultSnapshotPolicySettings) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StorageOptionsDefaultSnapshotPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **string** | How often a snapshot should be taken for this rule. | 
**RetentionCount** | Pointer to **float32** | Number of snapshots to retain for this rule. Once this count is reached, the oldest snapshot for the rule is deleted when a new one is created (applicable to NetApp ONTAP and Huawei OceanStor Dorado) | [optional] 
**RetentionPeriodDays** | Pointer to **float32** | Number of days to retain snapshots for this rule before they are eligible for deletion (for storages that expire snapshots by age rather than, or in addition to, a fixed count) | [optional] 

## Methods

### NewStorageOptionsDefaultSnapshotPolicyRule

`func NewStorageOptionsDefaultSnapshotPolicyRule(frequency string, ) *StorageOptionsDefaultSnapshotPolicyRule`

NewStorageOptionsDefaultSnapshotPolicyRule instantiates a new StorageOptionsDefaultSnapshotPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsDefaultSnapshotPolicyRuleWithDefaults

`func NewStorageOptionsDefaultSnapshotPolicyRuleWithDefaults() *StorageOptionsDefaultSnapshotPolicyRule`

NewStorageOptionsDefaultSnapshotPolicyRuleWithDefaults instantiates a new StorageOptionsDefaultSnapshotPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *StorageOptionsDefaultSnapshotPolicyRule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *StorageOptionsDefaultSnapshotPolicyRule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *StorageOptionsDefaultSnapshotPolicyRule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetRetentionCount

`func (o *StorageOptionsDefaultSnapshotPolicyRule) GetRetentionCount() float32`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *StorageOptionsDefaultSnapshotPolicyRule) GetRetentionCountOk() (*float32, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *StorageOptionsDefaultSnapshotPolicyRule) SetRetentionCount(v float32)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *StorageOptionsDefaultSnapshotPolicyRule) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetRetentionPeriodDays

`func (o *StorageOptionsDefaultSnapshotPolicyRule) GetRetentionPeriodDays() float32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *StorageOptionsDefaultSnapshotPolicyRule) GetRetentionPeriodDaysOk() (*float32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *StorageOptionsDefaultSnapshotPolicyRule) SetRetentionPeriodDays(v float32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.

### HasRetentionPeriodDays

`func (o *StorageOptionsDefaultSnapshotPolicyRule) HasRetentionPeriodDays() bool`

HasRetentionPeriodDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



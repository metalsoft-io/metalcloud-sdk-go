# BulkAssignDeviceConfigurationTemplateProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**[]DeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateProfile.md) | Newly created profile rows. | 
**Skipped** | [**[]BulkAssignSkippedItem**](BulkAssignSkippedItem.md) | Devices that were skipped because a matching profile already existed. | 
**TargetDeviceCount** | **float32** | Total target devices resolved (created + skipped). | 

## Methods

### NewBulkAssignDeviceConfigurationTemplateProfileResult

`func NewBulkAssignDeviceConfigurationTemplateProfileResult(created []DeviceConfigurationTemplateProfile, skipped []BulkAssignSkippedItem, targetDeviceCount float32, ) *BulkAssignDeviceConfigurationTemplateProfileResult`

NewBulkAssignDeviceConfigurationTemplateProfileResult instantiates a new BulkAssignDeviceConfigurationTemplateProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAssignDeviceConfigurationTemplateProfileResultWithDefaults

`func NewBulkAssignDeviceConfigurationTemplateProfileResultWithDefaults() *BulkAssignDeviceConfigurationTemplateProfileResult`

NewBulkAssignDeviceConfigurationTemplateProfileResultWithDefaults instantiates a new BulkAssignDeviceConfigurationTemplateProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) GetCreated() []DeviceConfigurationTemplateProfile`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) GetCreatedOk() (*[]DeviceConfigurationTemplateProfile, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) SetCreated(v []DeviceConfigurationTemplateProfile)`

SetCreated sets Created field to given value.


### GetSkipped

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) GetSkipped() []BulkAssignSkippedItem`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) GetSkippedOk() (*[]BulkAssignSkippedItem, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) SetSkipped(v []BulkAssignSkippedItem)`

SetSkipped sets Skipped field to given value.


### GetTargetDeviceCount

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) GetTargetDeviceCount() float32`

GetTargetDeviceCount returns the TargetDeviceCount field if non-nil, zero value otherwise.

### GetTargetDeviceCountOk

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) GetTargetDeviceCountOk() (*float32, bool)`

GetTargetDeviceCountOk returns a tuple with the TargetDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeviceCount

`func (o *BulkAssignDeviceConfigurationTemplateProfileResult) SetTargetDeviceCount(v float32)`

SetTargetDeviceCount sets TargetDeviceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



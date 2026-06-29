# BulkAssignSkippedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **int64** |  | 
**Reason** | **string** | Always \&quot;duplicate\&quot; for now; reserved for future skip reasons. | 
**ExistingProfileId** | **string** | Id of the existing profile that matches. | 

## Methods

### NewBulkAssignSkippedItem

`func NewBulkAssignSkippedItem(networkDeviceId int64, reason string, existingProfileId string, ) *BulkAssignSkippedItem`

NewBulkAssignSkippedItem instantiates a new BulkAssignSkippedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAssignSkippedItemWithDefaults

`func NewBulkAssignSkippedItemWithDefaults() *BulkAssignSkippedItem`

NewBulkAssignSkippedItemWithDefaults instantiates a new BulkAssignSkippedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *BulkAssignSkippedItem) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *BulkAssignSkippedItem) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *BulkAssignSkippedItem) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetReason

`func (o *BulkAssignSkippedItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BulkAssignSkippedItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BulkAssignSkippedItem) SetReason(v string)`

SetReason sets Reason field to given value.


### GetExistingProfileId

`func (o *BulkAssignSkippedItem) GetExistingProfileId() string`

GetExistingProfileId returns the ExistingProfileId field if non-nil, zero value otherwise.

### GetExistingProfileIdOk

`func (o *BulkAssignSkippedItem) GetExistingProfileIdOk() (*string, bool)`

GetExistingProfileIdOk returns a tuple with the ExistingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingProfileId

`func (o *BulkAssignSkippedItem) SetExistingProfileId(v string)`

SetExistingProfileId sets ExistingProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



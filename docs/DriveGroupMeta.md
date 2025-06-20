# DriveGroupMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Drive Group | 
**Tags** | Pointer to **[]string** | Tags for the Drive Group. | [optional] 

## Methods

### NewDriveGroupMeta

`func NewDriveGroupMeta(name string, ) *DriveGroupMeta`

NewDriveGroupMeta instantiates a new DriveGroupMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveGroupMetaWithDefaults

`func NewDriveGroupMetaWithDefaults() *DriveGroupMeta`

NewDriveGroupMetaWithDefaults instantiates a new DriveGroupMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DriveGroupMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DriveGroupMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DriveGroupMeta) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *DriveGroupMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DriveGroupMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DriveGroupMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DriveGroupMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



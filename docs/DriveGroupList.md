# DriveGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DriveGroup**](DriveGroup.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewDriveGroupList

`func NewDriveGroupList(data []DriveGroup, ) *DriveGroupList`

NewDriveGroupList instantiates a new DriveGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveGroupListWithDefaults

`func NewDriveGroupListWithDefaults() *DriveGroupList`

NewDriveGroupListWithDefaults instantiates a new DriveGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DriveGroupList) GetData() []DriveGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DriveGroupList) GetDataOk() (*[]DriveGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DriveGroupList) SetData(v []DriveGroup)`

SetData sets Data field to given value.


### GetLinks

`func (o *DriveGroupList) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DriveGroupList) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DriveGroupList) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DriveGroupList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



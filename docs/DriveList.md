# DriveList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Drive**](Drive.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewDriveList

`func NewDriveList(data []Drive, ) *DriveList`

NewDriveList instantiates a new DriveList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveListWithDefaults

`func NewDriveListWithDefaults() *DriveList`

NewDriveListWithDefaults instantiates a new DriveList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DriveList) GetData() []Drive`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DriveList) GetDataOk() (*[]Drive, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DriveList) SetData(v []Drive)`

SetData sets Data field to given value.


### GetLinks

`func (o *DriveList) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DriveList) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DriveList) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DriveList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



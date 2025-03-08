# CreateDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label of the Drive. | [optional] 
**GroupId** | **float32** | Id of the Drive Group | 
**SizeMb** | Pointer to **float32** | Disk size in MB for Drive | [optional] 

## Methods

### NewCreateDrive

`func NewCreateDrive(groupId float32, ) *CreateDrive`

NewCreateDrive instantiates a new CreateDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDriveWithDefaults

`func NewCreateDriveWithDefaults() *CreateDrive`

NewCreateDriveWithDefaults instantiates a new CreateDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateDrive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateDrive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateDrive) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateDrive) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupId

`func (o *CreateDrive) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateDrive) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateDrive) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetSizeMb

`func (o *CreateDrive) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *CreateDrive) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *CreateDrive) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.

### HasSizeMb

`func (o *CreateDrive) HasSizeMb() bool`

HasSizeMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



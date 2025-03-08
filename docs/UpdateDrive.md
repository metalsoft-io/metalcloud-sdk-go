# UpdateDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label of the Drive. | [optional] 
**SizeMb** | Pointer to **float32** | Disk size in MB for Drive | [optional] 

## Methods

### NewUpdateDrive

`func NewUpdateDrive() *UpdateDrive`

NewUpdateDrive instantiates a new UpdateDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDriveWithDefaults

`func NewUpdateDriveWithDefaults() *UpdateDrive`

NewUpdateDriveWithDefaults instantiates a new UpdateDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateDrive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateDrive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateDrive) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateDrive) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSizeMb

`func (o *UpdateDrive) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *UpdateDrive) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *UpdateDrive) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.

### HasSizeMb

`func (o *UpdateDrive) HasSizeMb() bool`

HasSizeMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



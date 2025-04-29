# CreateSharedDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeMb** | **float32** | Disk size in MiB for Drive | 
**LogicalNetworkId** | Pointer to **float32** | Id of the Logical Network for the Drive. | [optional] 
**Label** | Pointer to **string** | Label of the Drive. | [optional] 

## Methods

### NewCreateSharedDrive

`func NewCreateSharedDrive(sizeMb float32, ) *CreateSharedDrive`

NewCreateSharedDrive instantiates a new CreateSharedDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSharedDriveWithDefaults

`func NewCreateSharedDriveWithDefaults() *CreateSharedDrive`

NewCreateSharedDriveWithDefaults instantiates a new CreateSharedDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeMb

`func (o *CreateSharedDrive) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *CreateSharedDrive) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *CreateSharedDrive) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetLogicalNetworkId

`func (o *CreateSharedDrive) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateSharedDrive) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateSharedDrive) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *CreateSharedDrive) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetLabel

`func (o *CreateSharedDrive) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateSharedDrive) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateSharedDrive) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateSharedDrive) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



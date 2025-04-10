# UpdateDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label of the Drive. | [optional] 
**SizeMb** | Pointer to **float32** | Disk size in MB for Drive | [optional] 
**IoLimitPolicy** | Pointer to **string** | The IO limit policy of the Drive. | [optional] 
**LogicalNetworkId** | Pointer to **float32** | Id of the Logical Network for the Drive. | [optional] 

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

### GetIoLimitPolicy

`func (o *UpdateDrive) GetIoLimitPolicy() string`

GetIoLimitPolicy returns the IoLimitPolicy field if non-nil, zero value otherwise.

### GetIoLimitPolicyOk

`func (o *UpdateDrive) GetIoLimitPolicyOk() (*string, bool)`

GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoLimitPolicy

`func (o *UpdateDrive) SetIoLimitPolicy(v string)`

SetIoLimitPolicy sets IoLimitPolicy field to given value.

### HasIoLimitPolicy

`func (o *UpdateDrive) HasIoLimitPolicy() bool`

HasIoLimitPolicy returns a boolean if a field has been set.

### GetLogicalNetworkId

`func (o *UpdateDrive) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *UpdateDrive) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *UpdateDrive) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *UpdateDrive) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateFileShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for File Share | 
**LogicalNetworkId** | Pointer to **float32** | Id of the Logical Network for the File Share. | [optional] 
**Label** | Pointer to **string** | Display name of the File Share. | [optional] 
**Meta** | Pointer to [**FileShareMeta**](FileShareMeta.md) |  | [optional] 

## Methods

### NewCreateFileShare

`func NewCreateFileShare(sizeGB float32, ) *CreateFileShare`

NewCreateFileShare instantiates a new CreateFileShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileShareWithDefaults

`func NewCreateFileShareWithDefaults() *CreateFileShare`

NewCreateFileShareWithDefaults instantiates a new CreateFileShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *CreateFileShare) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *CreateFileShare) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *CreateFileShare) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.


### GetLogicalNetworkId

`func (o *CreateFileShare) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateFileShare) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateFileShare) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *CreateFileShare) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetLabel

`func (o *CreateFileShare) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateFileShare) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateFileShare) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateFileShare) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMeta

`func (o *CreateFileShare) GetMeta() FileShareMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateFileShare) GetMetaOk() (*FileShareMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateFileShare) SetMeta(v FileShareMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateFileShare) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



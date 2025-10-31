# CreateBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for Bucket | 
**Label** | Pointer to **string** | Label of the Bucket. | [optional] 
**Meta** | Pointer to [**BucketMeta**](BucketMeta.md) |  | [optional] 
**LogicalNetworkId** | **float32** | Id of the Logical Network for the Bucket. | 
**StoragePoolId** | Pointer to **float32** | Id of the storage pool the Bucket is assigned to | [optional] 

## Methods

### NewCreateBucket

`func NewCreateBucket(sizeGB float32, logicalNetworkId float32, ) *CreateBucket`

NewCreateBucket instantiates a new CreateBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketWithDefaults

`func NewCreateBucketWithDefaults() *CreateBucket`

NewCreateBucketWithDefaults instantiates a new CreateBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *CreateBucket) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *CreateBucket) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *CreateBucket) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.


### GetLabel

`func (o *CreateBucket) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateBucket) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateBucket) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateBucket) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMeta

`func (o *CreateBucket) GetMeta() BucketMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateBucket) GetMetaOk() (*BucketMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateBucket) SetMeta(v BucketMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateBucket) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLogicalNetworkId

`func (o *CreateBucket) GetLogicalNetworkId() float32`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateBucket) GetLogicalNetworkIdOk() (*float32, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateBucket) SetLogicalNetworkId(v float32)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetStoragePoolId

`func (o *CreateBucket) GetStoragePoolId() float32`

GetStoragePoolId returns the StoragePoolId field if non-nil, zero value otherwise.

### GetStoragePoolIdOk

`func (o *CreateBucket) GetStoragePoolIdOk() (*float32, bool)`

GetStoragePoolIdOk returns a tuple with the StoragePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolId

`func (o *CreateBucket) SetStoragePoolId(v float32)`

SetStoragePoolId sets StoragePoolId field to given value.

### HasStoragePoolId

`func (o *CreateBucket) HasStoragePoolId() bool`

HasStoragePoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



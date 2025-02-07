# CreateBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | **float32** | Disk size in GB for Bucket | 
**Label** | Pointer to **string** | Label of the Bucket. | [optional] 

## Methods

### NewCreateBucket

`func NewCreateBucket(sizeGB float32, ) *CreateBucket`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



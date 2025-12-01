# GenericBucketDiscoverInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Path of the Bucket | 
**MountingInformation** | [**[]GenericBucketMountingInformation**](GenericBucketMountingInformation.md) | Mounting information of the Bucket | 

## Methods

### NewGenericBucketDiscoverInformation

`func NewGenericBucketDiscoverInformation(path string, mountingInformation []GenericBucketMountingInformation, ) *GenericBucketDiscoverInformation`

NewGenericBucketDiscoverInformation instantiates a new GenericBucketDiscoverInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericBucketDiscoverInformationWithDefaults

`func NewGenericBucketDiscoverInformationWithDefaults() *GenericBucketDiscoverInformation`

NewGenericBucketDiscoverInformationWithDefaults instantiates a new GenericBucketDiscoverInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GenericBucketDiscoverInformation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GenericBucketDiscoverInformation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GenericBucketDiscoverInformation) SetPath(v string)`

SetPath sets Path field to given value.


### GetMountingInformation

`func (o *GenericBucketDiscoverInformation) GetMountingInformation() []GenericBucketMountingInformation`

GetMountingInformation returns the MountingInformation field if non-nil, zero value otherwise.

### GetMountingInformationOk

`func (o *GenericBucketDiscoverInformation) GetMountingInformationOk() (*[]GenericBucketMountingInformation, bool)`

GetMountingInformationOk returns a tuple with the MountingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingInformation

`func (o *GenericBucketDiscoverInformation) SetMountingInformation(v []GenericBucketMountingInformation)`

SetMountingInformation sets MountingInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



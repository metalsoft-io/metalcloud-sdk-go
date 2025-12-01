# GenericFileShareDiscoverInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Path of the File Share | 
**MountingInformation** | [**[]GenericFileShareMountingInformation**](GenericFileShareMountingInformation.md) | Mounting information of the File Share | 

## Methods

### NewGenericFileShareDiscoverInformation

`func NewGenericFileShareDiscoverInformation(path string, mountingInformation []GenericFileShareMountingInformation, ) *GenericFileShareDiscoverInformation`

NewGenericFileShareDiscoverInformation instantiates a new GenericFileShareDiscoverInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericFileShareDiscoverInformationWithDefaults

`func NewGenericFileShareDiscoverInformationWithDefaults() *GenericFileShareDiscoverInformation`

NewGenericFileShareDiscoverInformationWithDefaults instantiates a new GenericFileShareDiscoverInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GenericFileShareDiscoverInformation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GenericFileShareDiscoverInformation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GenericFileShareDiscoverInformation) SetPath(v string)`

SetPath sets Path field to given value.


### GetMountingInformation

`func (o *GenericFileShareDiscoverInformation) GetMountingInformation() []GenericFileShareMountingInformation`

GetMountingInformation returns the MountingInformation field if non-nil, zero value otherwise.

### GetMountingInformationOk

`func (o *GenericFileShareDiscoverInformation) GetMountingInformationOk() (*[]GenericFileShareMountingInformation, bool)`

GetMountingInformationOk returns a tuple with the MountingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingInformation

`func (o *GenericFileShareDiscoverInformation) SetMountingInformation(v []GenericFileShareMountingInformation)`

SetMountingInformation sets MountingInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



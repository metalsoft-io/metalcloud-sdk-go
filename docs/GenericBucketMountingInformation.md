# GenericBucketMountingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullPath** | **string** | Full path of the Bucket | 
**Host** | **string** | Host where the Bucket can be accessed | 
**Port** | Pointer to **float32** | Port where the Bucket can be accessed | [optional] 

## Methods

### NewGenericBucketMountingInformation

`func NewGenericBucketMountingInformation(fullPath string, host string, ) *GenericBucketMountingInformation`

NewGenericBucketMountingInformation instantiates a new GenericBucketMountingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericBucketMountingInformationWithDefaults

`func NewGenericBucketMountingInformationWithDefaults() *GenericBucketMountingInformation`

NewGenericBucketMountingInformationWithDefaults instantiates a new GenericBucketMountingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullPath

`func (o *GenericBucketMountingInformation) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *GenericBucketMountingInformation) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *GenericBucketMountingInformation) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.


### GetHost

`func (o *GenericBucketMountingInformation) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GenericBucketMountingInformation) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GenericBucketMountingInformation) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *GenericBucketMountingInformation) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GenericBucketMountingInformation) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GenericBucketMountingInformation) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GenericBucketMountingInformation) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# QuotaUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureServerGroupMaxCount** | Pointer to **float32** | Current number of server groups across all infrastructures owned by users in this scope | [optional] 
**InfrastructureDriveMaxCount** | Pointer to **float32** | Current number of drives across all infrastructures owned by users in this scope | [optional] 
**InfrastructureFileShareMaxCount** | Pointer to **float32** | Current number of file shares across all infrastructures owned by users in this scope | [optional] 
**InfrastructureBucketMaxCount** | Pointer to **float32** | Current number of buckets across all infrastructures owned by users in this scope | [optional] 
**InfrastructureVmInstanceGroupMaxCount** | Pointer to **float32** | Current number of VM instance groups across all infrastructures owned by users in this scope | [optional] 
**UserResourceServerTypeNameToMaxCount** | Pointer to **map[string]interface{}** | Map of server type name to the current total instance count in this scope | [optional] 

## Methods

### NewQuotaUsage

`func NewQuotaUsage() *QuotaUsage`

NewQuotaUsage instantiates a new QuotaUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaUsageWithDefaults

`func NewQuotaUsageWithDefaults() *QuotaUsage`

NewQuotaUsageWithDefaults instantiates a new QuotaUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureServerGroupMaxCount

`func (o *QuotaUsage) GetInfrastructureServerGroupMaxCount() float32`

GetInfrastructureServerGroupMaxCount returns the InfrastructureServerGroupMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureServerGroupMaxCountOk

`func (o *QuotaUsage) GetInfrastructureServerGroupMaxCountOk() (*float32, bool)`

GetInfrastructureServerGroupMaxCountOk returns a tuple with the InfrastructureServerGroupMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureServerGroupMaxCount

`func (o *QuotaUsage) SetInfrastructureServerGroupMaxCount(v float32)`

SetInfrastructureServerGroupMaxCount sets InfrastructureServerGroupMaxCount field to given value.

### HasInfrastructureServerGroupMaxCount

`func (o *QuotaUsage) HasInfrastructureServerGroupMaxCount() bool`

HasInfrastructureServerGroupMaxCount returns a boolean if a field has been set.

### GetInfrastructureDriveMaxCount

`func (o *QuotaUsage) GetInfrastructureDriveMaxCount() float32`

GetInfrastructureDriveMaxCount returns the InfrastructureDriveMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureDriveMaxCountOk

`func (o *QuotaUsage) GetInfrastructureDriveMaxCountOk() (*float32, bool)`

GetInfrastructureDriveMaxCountOk returns a tuple with the InfrastructureDriveMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDriveMaxCount

`func (o *QuotaUsage) SetInfrastructureDriveMaxCount(v float32)`

SetInfrastructureDriveMaxCount sets InfrastructureDriveMaxCount field to given value.

### HasInfrastructureDriveMaxCount

`func (o *QuotaUsage) HasInfrastructureDriveMaxCount() bool`

HasInfrastructureDriveMaxCount returns a boolean if a field has been set.

### GetInfrastructureFileShareMaxCount

`func (o *QuotaUsage) GetInfrastructureFileShareMaxCount() float32`

GetInfrastructureFileShareMaxCount returns the InfrastructureFileShareMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureFileShareMaxCountOk

`func (o *QuotaUsage) GetInfrastructureFileShareMaxCountOk() (*float32, bool)`

GetInfrastructureFileShareMaxCountOk returns a tuple with the InfrastructureFileShareMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureFileShareMaxCount

`func (o *QuotaUsage) SetInfrastructureFileShareMaxCount(v float32)`

SetInfrastructureFileShareMaxCount sets InfrastructureFileShareMaxCount field to given value.

### HasInfrastructureFileShareMaxCount

`func (o *QuotaUsage) HasInfrastructureFileShareMaxCount() bool`

HasInfrastructureFileShareMaxCount returns a boolean if a field has been set.

### GetInfrastructureBucketMaxCount

`func (o *QuotaUsage) GetInfrastructureBucketMaxCount() float32`

GetInfrastructureBucketMaxCount returns the InfrastructureBucketMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureBucketMaxCountOk

`func (o *QuotaUsage) GetInfrastructureBucketMaxCountOk() (*float32, bool)`

GetInfrastructureBucketMaxCountOk returns a tuple with the InfrastructureBucketMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureBucketMaxCount

`func (o *QuotaUsage) SetInfrastructureBucketMaxCount(v float32)`

SetInfrastructureBucketMaxCount sets InfrastructureBucketMaxCount field to given value.

### HasInfrastructureBucketMaxCount

`func (o *QuotaUsage) HasInfrastructureBucketMaxCount() bool`

HasInfrastructureBucketMaxCount returns a boolean if a field has been set.

### GetInfrastructureVmInstanceGroupMaxCount

`func (o *QuotaUsage) GetInfrastructureVmInstanceGroupMaxCount() float32`

GetInfrastructureVmInstanceGroupMaxCount returns the InfrastructureVmInstanceGroupMaxCount field if non-nil, zero value otherwise.

### GetInfrastructureVmInstanceGroupMaxCountOk

`func (o *QuotaUsage) GetInfrastructureVmInstanceGroupMaxCountOk() (*float32, bool)`

GetInfrastructureVmInstanceGroupMaxCountOk returns a tuple with the InfrastructureVmInstanceGroupMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureVmInstanceGroupMaxCount

`func (o *QuotaUsage) SetInfrastructureVmInstanceGroupMaxCount(v float32)`

SetInfrastructureVmInstanceGroupMaxCount sets InfrastructureVmInstanceGroupMaxCount field to given value.

### HasInfrastructureVmInstanceGroupMaxCount

`func (o *QuotaUsage) HasInfrastructureVmInstanceGroupMaxCount() bool`

HasInfrastructureVmInstanceGroupMaxCount returns a boolean if a field has been set.

### GetUserResourceServerTypeNameToMaxCount

`func (o *QuotaUsage) GetUserResourceServerTypeNameToMaxCount() map[string]interface{}`

GetUserResourceServerTypeNameToMaxCount returns the UserResourceServerTypeNameToMaxCount field if non-nil, zero value otherwise.

### GetUserResourceServerTypeNameToMaxCountOk

`func (o *QuotaUsage) GetUserResourceServerTypeNameToMaxCountOk() (*map[string]interface{}, bool)`

GetUserResourceServerTypeNameToMaxCountOk returns a tuple with the UserResourceServerTypeNameToMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResourceServerTypeNameToMaxCount

`func (o *QuotaUsage) SetUserResourceServerTypeNameToMaxCount(v map[string]interface{})`

SetUserResourceServerTypeNameToMaxCount sets UserResourceServerTypeNameToMaxCount field to given value.

### HasUserResourceServerTypeNameToMaxCount

`func (o *QuotaUsage) HasUserResourceServerTypeNameToMaxCount() bool`

HasUserResourceServerTypeNameToMaxCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



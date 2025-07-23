# CreateVMPoolOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Name of the cluster for the VM Pool | [optional] 

## Methods

### NewCreateVMPoolOptions

`func NewCreateVMPoolOptions() *CreateVMPoolOptions`

NewCreateVMPoolOptions instantiates a new CreateVMPoolOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMPoolOptionsWithDefaults

`func NewCreateVMPoolOptionsWithDefaults() *CreateVMPoolOptions`

NewCreateVMPoolOptionsWithDefaults instantiates a new CreateVMPoolOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *CreateVMPoolOptions) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateVMPoolOptions) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateVMPoolOptions) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CreateVMPoolOptions) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



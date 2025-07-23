# VMPoolOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Name of the cluster for the VM Pool | [optional] 
**DatastoreName** | Pointer to **string** | Name of the datastore to use for the VM Pool | [optional] 
**DatastoreNames** | Pointer to **[]string** | List of datastore names available on the VM Pool | [optional] 

## Methods

### NewVMPoolOptions

`func NewVMPoolOptions() *VMPoolOptions`

NewVMPoolOptions instantiates a new VMPoolOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolOptionsWithDefaults

`func NewVMPoolOptionsWithDefaults() *VMPoolOptions`

NewVMPoolOptionsWithDefaults instantiates a new VMPoolOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *VMPoolOptions) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VMPoolOptions) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VMPoolOptions) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *VMPoolOptions) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDatastoreName

`func (o *VMPoolOptions) GetDatastoreName() string`

GetDatastoreName returns the DatastoreName field if non-nil, zero value otherwise.

### GetDatastoreNameOk

`func (o *VMPoolOptions) GetDatastoreNameOk() (*string, bool)`

GetDatastoreNameOk returns a tuple with the DatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreName

`func (o *VMPoolOptions) SetDatastoreName(v string)`

SetDatastoreName sets DatastoreName field to given value.

### HasDatastoreName

`func (o *VMPoolOptions) HasDatastoreName() bool`

HasDatastoreName returns a boolean if a field has been set.

### GetDatastoreNames

`func (o *VMPoolOptions) GetDatastoreNames() []string`

GetDatastoreNames returns the DatastoreNames field if non-nil, zero value otherwise.

### GetDatastoreNamesOk

`func (o *VMPoolOptions) GetDatastoreNamesOk() (*[]string, bool)`

GetDatastoreNamesOk returns a tuple with the DatastoreNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreNames

`func (o *VMPoolOptions) SetDatastoreNames(v []string)`

SetDatastoreNames sets DatastoreNames field to given value.

### HasDatastoreNames

`func (o *VMPoolOptions) HasDatastoreNames() bool`

HasDatastoreNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



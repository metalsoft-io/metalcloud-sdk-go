# UpdateVMPoolOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreName** | Pointer to **string** | Name of the datastore to use for the VM Pool | [optional] 

## Methods

### NewUpdateVMPoolOptions

`func NewUpdateVMPoolOptions() *UpdateVMPoolOptions`

NewUpdateVMPoolOptions instantiates a new UpdateVMPoolOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMPoolOptionsWithDefaults

`func NewUpdateVMPoolOptionsWithDefaults() *UpdateVMPoolOptions`

NewUpdateVMPoolOptionsWithDefaults instantiates a new UpdateVMPoolOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreName

`func (o *UpdateVMPoolOptions) GetDatastoreName() string`

GetDatastoreName returns the DatastoreName field if non-nil, zero value otherwise.

### GetDatastoreNameOk

`func (o *UpdateVMPoolOptions) GetDatastoreNameOk() (*string, bool)`

GetDatastoreNameOk returns a tuple with the DatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreName

`func (o *UpdateVMPoolOptions) SetDatastoreName(v string)`

SetDatastoreName sets DatastoreName field to given value.

### HasDatastoreName

`func (o *UpdateVMPoolOptions) HasDatastoreName() bool`

HasDatastoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



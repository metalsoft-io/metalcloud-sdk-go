# FileShareHostBulkOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceArrayId** | **float32** | Id of the Instance Array Host that will be modified | 
**OperationType** | **string** | Operation type for the Instance Array Host | 

## Methods

### NewFileShareHostBulkOperation

`func NewFileShareHostBulkOperation(instanceArrayId float32, operationType string, ) *FileShareHostBulkOperation`

NewFileShareHostBulkOperation instantiates a new FileShareHostBulkOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareHostBulkOperationWithDefaults

`func NewFileShareHostBulkOperationWithDefaults() *FileShareHostBulkOperation`

NewFileShareHostBulkOperationWithDefaults instantiates a new FileShareHostBulkOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceArrayId

`func (o *FileShareHostBulkOperation) GetInstanceArrayId() float32`

GetInstanceArrayId returns the InstanceArrayId field if non-nil, zero value otherwise.

### GetInstanceArrayIdOk

`func (o *FileShareHostBulkOperation) GetInstanceArrayIdOk() (*float32, bool)`

GetInstanceArrayIdOk returns a tuple with the InstanceArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrayId

`func (o *FileShareHostBulkOperation) SetInstanceArrayId(v float32)`

SetInstanceArrayId sets InstanceArrayId field to given value.


### GetOperationType

`func (o *FileShareHostBulkOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *FileShareHostBulkOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *FileShareHostBulkOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



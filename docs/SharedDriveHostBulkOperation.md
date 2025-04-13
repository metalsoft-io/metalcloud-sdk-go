# SharedDriveHostBulkOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerInstanceGroupId** | **float32** | Id of the Server Instance Group Host that will be modified | 
**OperationType** | **string** | Operation type for the Server Instance Group Host | 

## Methods

### NewSharedDriveHostBulkOperation

`func NewSharedDriveHostBulkOperation(serverInstanceGroupId float32, operationType string, ) *SharedDriveHostBulkOperation`

NewSharedDriveHostBulkOperation instantiates a new SharedDriveHostBulkOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveHostBulkOperationWithDefaults

`func NewSharedDriveHostBulkOperationWithDefaults() *SharedDriveHostBulkOperation`

NewSharedDriveHostBulkOperationWithDefaults instantiates a new SharedDriveHostBulkOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerInstanceGroupId

`func (o *SharedDriveHostBulkOperation) GetServerInstanceGroupId() float32`

GetServerInstanceGroupId returns the ServerInstanceGroupId field if non-nil, zero value otherwise.

### GetServerInstanceGroupIdOk

`func (o *SharedDriveHostBulkOperation) GetServerInstanceGroupIdOk() (*float32, bool)`

GetServerInstanceGroupIdOk returns a tuple with the ServerInstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroupId

`func (o *SharedDriveHostBulkOperation) SetServerInstanceGroupId(v float32)`

SetServerInstanceGroupId sets ServerInstanceGroupId field to given value.


### GetOperationType

`func (o *SharedDriveHostBulkOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SharedDriveHostBulkOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SharedDriveHostBulkOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



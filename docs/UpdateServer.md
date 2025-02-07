# UpdateServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTypeId** | Pointer to **float32** | The id of the server type. | [optional] 
**ManagementAddress** | Pointer to **string** | The Management Address of the server. | [optional] 
**Username** | Pointer to **string** | The username to use. | [optional] 
**ServerCleanupPolicyId** | Pointer to **float32** | The cleanup policy id of the server. | [optional] 
**ServerComments** | Pointer to **string** | The comments of the server. | [optional] 
**ChassisRackId** | Pointer to **float32** | The chassis rack id of the server. | [optional] 
**RackName** | Pointer to **string** | The chassis rack name of the server. | [optional] 
**RackPositionUpperUnit** | Pointer to **string** | The chassis rack upper unit position of the server. | [optional] 
**RackPositionLowerUnit** | Pointer to **string** | The chassis rack lower unit position of the server. | [optional] 
**InventoryId** | Pointer to **string** | The inventory id of the server. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Server. | [optional] 
**ResourcePoolId** | Pointer to **float32** | Resource Pool ID | [optional] 
**ServerStatus** | Pointer to **string** | The status of the server. | [optional] 
**ServerClass** | Pointer to **string** | The server class. | [optional] 

## Methods

### NewUpdateServer

`func NewUpdateServer() *UpdateServer`

NewUpdateServer instantiates a new UpdateServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerWithDefaults

`func NewUpdateServerWithDefaults() *UpdateServer`

NewUpdateServerWithDefaults instantiates a new UpdateServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTypeId

`func (o *UpdateServer) GetServerTypeId() float32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *UpdateServer) GetServerTypeIdOk() (*float32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *UpdateServer) SetServerTypeId(v float32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *UpdateServer) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetManagementAddress

`func (o *UpdateServer) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *UpdateServer) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *UpdateServer) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *UpdateServer) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetServerCleanupPolicyId

`func (o *UpdateServer) GetServerCleanupPolicyId() float32`

GetServerCleanupPolicyId returns the ServerCleanupPolicyId field if non-nil, zero value otherwise.

### GetServerCleanupPolicyIdOk

`func (o *UpdateServer) GetServerCleanupPolicyIdOk() (*float32, bool)`

GetServerCleanupPolicyIdOk returns a tuple with the ServerCleanupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCleanupPolicyId

`func (o *UpdateServer) SetServerCleanupPolicyId(v float32)`

SetServerCleanupPolicyId sets ServerCleanupPolicyId field to given value.

### HasServerCleanupPolicyId

`func (o *UpdateServer) HasServerCleanupPolicyId() bool`

HasServerCleanupPolicyId returns a boolean if a field has been set.

### GetServerComments

`func (o *UpdateServer) GetServerComments() string`

GetServerComments returns the ServerComments field if non-nil, zero value otherwise.

### GetServerCommentsOk

`func (o *UpdateServer) GetServerCommentsOk() (*string, bool)`

GetServerCommentsOk returns a tuple with the ServerComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComments

`func (o *UpdateServer) SetServerComments(v string)`

SetServerComments sets ServerComments field to given value.

### HasServerComments

`func (o *UpdateServer) HasServerComments() bool`

HasServerComments returns a boolean if a field has been set.

### GetChassisRackId

`func (o *UpdateServer) GetChassisRackId() float32`

GetChassisRackId returns the ChassisRackId field if non-nil, zero value otherwise.

### GetChassisRackIdOk

`func (o *UpdateServer) GetChassisRackIdOk() (*float32, bool)`

GetChassisRackIdOk returns a tuple with the ChassisRackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisRackId

`func (o *UpdateServer) SetChassisRackId(v float32)`

SetChassisRackId sets ChassisRackId field to given value.

### HasChassisRackId

`func (o *UpdateServer) HasChassisRackId() bool`

HasChassisRackId returns a boolean if a field has been set.

### GetRackName

`func (o *UpdateServer) GetRackName() string`

GetRackName returns the RackName field if non-nil, zero value otherwise.

### GetRackNameOk

`func (o *UpdateServer) GetRackNameOk() (*string, bool)`

GetRackNameOk returns a tuple with the RackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackName

`func (o *UpdateServer) SetRackName(v string)`

SetRackName sets RackName field to given value.

### HasRackName

`func (o *UpdateServer) HasRackName() bool`

HasRackName returns a boolean if a field has been set.

### GetRackPositionUpperUnit

`func (o *UpdateServer) GetRackPositionUpperUnit() string`

GetRackPositionUpperUnit returns the RackPositionUpperUnit field if non-nil, zero value otherwise.

### GetRackPositionUpperUnitOk

`func (o *UpdateServer) GetRackPositionUpperUnitOk() (*string, bool)`

GetRackPositionUpperUnitOk returns a tuple with the RackPositionUpperUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionUpperUnit

`func (o *UpdateServer) SetRackPositionUpperUnit(v string)`

SetRackPositionUpperUnit sets RackPositionUpperUnit field to given value.

### HasRackPositionUpperUnit

`func (o *UpdateServer) HasRackPositionUpperUnit() bool`

HasRackPositionUpperUnit returns a boolean if a field has been set.

### GetRackPositionLowerUnit

`func (o *UpdateServer) GetRackPositionLowerUnit() string`

GetRackPositionLowerUnit returns the RackPositionLowerUnit field if non-nil, zero value otherwise.

### GetRackPositionLowerUnitOk

`func (o *UpdateServer) GetRackPositionLowerUnitOk() (*string, bool)`

GetRackPositionLowerUnitOk returns a tuple with the RackPositionLowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPositionLowerUnit

`func (o *UpdateServer) SetRackPositionLowerUnit(v string)`

SetRackPositionLowerUnit sets RackPositionLowerUnit field to given value.

### HasRackPositionLowerUnit

`func (o *UpdateServer) HasRackPositionLowerUnit() bool`

HasRackPositionLowerUnit returns a boolean if a field has been set.

### GetInventoryId

`func (o *UpdateServer) GetInventoryId() string`

GetInventoryId returns the InventoryId field if non-nil, zero value otherwise.

### GetInventoryIdOk

`func (o *UpdateServer) GetInventoryIdOk() (*string, bool)`

GetInventoryIdOk returns a tuple with the InventoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryId

`func (o *UpdateServer) SetInventoryId(v string)`

SetInventoryId sets InventoryId field to given value.

### HasInventoryId

`func (o *UpdateServer) HasInventoryId() bool`

HasInventoryId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateServer) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateServer) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateServer) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *UpdateServer) GetResourcePoolId() float32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *UpdateServer) GetResourcePoolIdOk() (*float32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *UpdateServer) SetResourcePoolId(v float32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *UpdateServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetServerStatus

`func (o *UpdateServer) GetServerStatus() string`

GetServerStatus returns the ServerStatus field if non-nil, zero value otherwise.

### GetServerStatusOk

`func (o *UpdateServer) GetServerStatusOk() (*string, bool)`

GetServerStatusOk returns a tuple with the ServerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatus

`func (o *UpdateServer) SetServerStatus(v string)`

SetServerStatus sets ServerStatus field to given value.

### HasServerStatus

`func (o *UpdateServer) HasServerStatus() bool`

HasServerStatus returns a boolean if a field has been set.

### GetServerClass

`func (o *UpdateServer) GetServerClass() string`

GetServerClass returns the ServerClass field if non-nil, zero value otherwise.

### GetServerClassOk

`func (o *UpdateServer) GetServerClassOk() (*string, bool)`

GetServerClassOk returns a tuple with the ServerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClass

`func (o *UpdateServer) SetServerClass(v string)`

SetServerClass sets ServerClass field to given value.

### HasServerClass

`func (o *UpdateServer) HasServerClass() bool`

HasServerClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



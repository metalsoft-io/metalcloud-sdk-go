# VMInstanceGroupInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Interface ID | 
**NetworkId** | **float32** | Network ID | 
**InterfaceIndex** | **float32** | Interface index | 
**ServiceStatus** | **string** | Service status of the VM Instance Group Interface. | 
**GroupId** | **float32** | VM Instance Group ID | 
**InfrastructureId** | **float32** | Infrastructure ID | 
**ChangeId** | **float32** | Operation ID | 
**Label** | **string** | Interface label | 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance Group Interface creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group Interface update. | 
**Operation** | **map[string]interface{}** | Operation object for the VM Instance Group. | 

## Methods

### NewVMInstanceGroupInterface

`func NewVMInstanceGroupInterface(id float32, networkId float32, interfaceIndex float32, serviceStatus string, groupId float32, infrastructureId float32, changeId float32, label string, createdTimestamp string, updatedTimestamp string, operation map[string]interface{}, ) *VMInstanceGroupInterface`

NewVMInstanceGroupInterface instantiates a new VMInstanceGroupInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupInterfaceWithDefaults

`func NewVMInstanceGroupInterfaceWithDefaults() *VMInstanceGroupInterface`

NewVMInstanceGroupInterfaceWithDefaults instantiates a new VMInstanceGroupInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMInstanceGroupInterface) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInstanceGroupInterface) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInstanceGroupInterface) SetId(v float32)`

SetId sets Id field to given value.


### GetNetworkId

`func (o *VMInstanceGroupInterface) GetNetworkId() float32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VMInstanceGroupInterface) GetNetworkIdOk() (*float32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VMInstanceGroupInterface) SetNetworkId(v float32)`

SetNetworkId sets NetworkId field to given value.


### GetInterfaceIndex

`func (o *VMInstanceGroupInterface) GetInterfaceIndex() float32`

GetInterfaceIndex returns the InterfaceIndex field if non-nil, zero value otherwise.

### GetInterfaceIndexOk

`func (o *VMInstanceGroupInterface) GetInterfaceIndexOk() (*float32, bool)`

GetInterfaceIndexOk returns a tuple with the InterfaceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIndex

`func (o *VMInstanceGroupInterface) SetInterfaceIndex(v float32)`

SetInterfaceIndex sets InterfaceIndex field to given value.


### GetServiceStatus

`func (o *VMInstanceGroupInterface) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *VMInstanceGroupInterface) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *VMInstanceGroupInterface) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetGroupId

`func (o *VMInstanceGroupInterface) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VMInstanceGroupInterface) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VMInstanceGroupInterface) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetInfrastructureId

`func (o *VMInstanceGroupInterface) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VMInstanceGroupInterface) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VMInstanceGroupInterface) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetChangeId

`func (o *VMInstanceGroupInterface) GetChangeId() float32`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *VMInstanceGroupInterface) GetChangeIdOk() (*float32, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *VMInstanceGroupInterface) SetChangeId(v float32)`

SetChangeId sets ChangeId field to given value.


### GetLabel

`func (o *VMInstanceGroupInterface) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstanceGroupInterface) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstanceGroupInterface) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *VMInstanceGroupInterface) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VMInstanceGroupInterface) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VMInstanceGroupInterface) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *VMInstanceGroupInterface) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstanceGroupInterface) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstanceGroupInterface) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetOperation

`func (o *VMInstanceGroupInterface) GetOperation() map[string]interface{}`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VMInstanceGroupInterface) GetOperationOk() (*map[string]interface{}, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VMInstanceGroupInterface) SetOperation(v map[string]interface{})`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



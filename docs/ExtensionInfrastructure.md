# ExtensionInfrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceArrays** | Pointer to [**[]ExtensionInstanceArray**](ExtensionInstanceArray.md) | Instance array details for the infrastructure requirement. | [optional] 
**SharedDrives** | Pointer to [**[]ExtensionSharedDrive**](ExtensionSharedDrive.md) | Shared drive details for the infrastructure requirement. | [optional] 
**LogicalNetworks** | Pointer to [**[]ExtensionInfrastructureLogicalNetwork**](ExtensionInfrastructureLogicalNetwork.md) | Logical networks for the infrastructure requirement. | [optional] 

## Methods

### NewExtensionInfrastructure

`func NewExtensionInfrastructure() *ExtensionInfrastructure`

NewExtensionInfrastructure instantiates a new ExtensionInfrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInfrastructureWithDefaults

`func NewExtensionInfrastructureWithDefaults() *ExtensionInfrastructure`

NewExtensionInfrastructureWithDefaults instantiates a new ExtensionInfrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceArrays

`func (o *ExtensionInfrastructure) GetInstanceArrays() []ExtensionInstanceArray`

GetInstanceArrays returns the InstanceArrays field if non-nil, zero value otherwise.

### GetInstanceArraysOk

`func (o *ExtensionInfrastructure) GetInstanceArraysOk() (*[]ExtensionInstanceArray, bool)`

GetInstanceArraysOk returns a tuple with the InstanceArrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrays

`func (o *ExtensionInfrastructure) SetInstanceArrays(v []ExtensionInstanceArray)`

SetInstanceArrays sets InstanceArrays field to given value.

### HasInstanceArrays

`func (o *ExtensionInfrastructure) HasInstanceArrays() bool`

HasInstanceArrays returns a boolean if a field has been set.

### GetSharedDrives

`func (o *ExtensionInfrastructure) GetSharedDrives() []ExtensionSharedDrive`

GetSharedDrives returns the SharedDrives field if non-nil, zero value otherwise.

### GetSharedDrivesOk

`func (o *ExtensionInfrastructure) GetSharedDrivesOk() (*[]ExtensionSharedDrive, bool)`

GetSharedDrivesOk returns a tuple with the SharedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrives

`func (o *ExtensionInfrastructure) SetSharedDrives(v []ExtensionSharedDrive)`

SetSharedDrives sets SharedDrives field to given value.

### HasSharedDrives

`func (o *ExtensionInfrastructure) HasSharedDrives() bool`

HasSharedDrives returns a boolean if a field has been set.

### GetLogicalNetworks

`func (o *ExtensionInfrastructure) GetLogicalNetworks() []ExtensionInfrastructureLogicalNetwork`

GetLogicalNetworks returns the LogicalNetworks field if non-nil, zero value otherwise.

### GetLogicalNetworksOk

`func (o *ExtensionInfrastructure) GetLogicalNetworksOk() (*[]ExtensionInfrastructureLogicalNetwork, bool)`

GetLogicalNetworksOk returns a tuple with the LogicalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworks

`func (o *ExtensionInfrastructure) SetLogicalNetworks(v []ExtensionInfrastructureLogicalNetwork)`

SetLogicalNetworks sets LogicalNetworks field to given value.

### HasLogicalNetworks

`func (o *ExtensionInfrastructure) HasLogicalNetworks() bool`

HasLogicalNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



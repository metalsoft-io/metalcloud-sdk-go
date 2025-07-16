# ExtensionInstanceArrayLogicalNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the logical network. | 
**NetworkConnection** | [**CreateServerInstanceGroupNetworkConnection**](CreateServerInstanceGroupNetworkConnection.md) | Network connection details for the logical network. | 

## Methods

### NewExtensionInstanceArrayLogicalNetworkDto

`func NewExtensionInstanceArrayLogicalNetworkDto(label string, networkConnection CreateServerInstanceGroupNetworkConnection, ) *ExtensionInstanceArrayLogicalNetworkDto`

NewExtensionInstanceArrayLogicalNetworkDto instantiates a new ExtensionInstanceArrayLogicalNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstanceArrayLogicalNetworkDtoWithDefaults

`func NewExtensionInstanceArrayLogicalNetworkDtoWithDefaults() *ExtensionInstanceArrayLogicalNetworkDto`

NewExtensionInstanceArrayLogicalNetworkDtoWithDefaults instantiates a new ExtensionInstanceArrayLogicalNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionInstanceArrayLogicalNetworkDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInstanceArrayLogicalNetworkDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInstanceArrayLogicalNetworkDto) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNetworkConnection

`func (o *ExtensionInstanceArrayLogicalNetworkDto) GetNetworkConnection() CreateServerInstanceGroupNetworkConnection`

GetNetworkConnection returns the NetworkConnection field if non-nil, zero value otherwise.

### GetNetworkConnectionOk

`func (o *ExtensionInstanceArrayLogicalNetworkDto) GetNetworkConnectionOk() (*CreateServerInstanceGroupNetworkConnection, bool)`

GetNetworkConnectionOk returns a tuple with the NetworkConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnection

`func (o *ExtensionInstanceArrayLogicalNetworkDto) SetNetworkConnection(v CreateServerInstanceGroupNetworkConnection)`

SetNetworkConnection sets NetworkConnection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



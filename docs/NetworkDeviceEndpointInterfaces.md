# NetworkDeviceEndpointInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NetworkDeviceEndpointInterface**](NetworkDeviceEndpointInterface.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceEndpointInterfaces

`func NewNetworkDeviceEndpointInterfaces(data []NetworkDeviceEndpointInterface, ) *NetworkDeviceEndpointInterfaces`

NewNetworkDeviceEndpointInterfaces instantiates a new NetworkDeviceEndpointInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceEndpointInterfacesWithDefaults

`func NewNetworkDeviceEndpointInterfacesWithDefaults() *NetworkDeviceEndpointInterfaces`

NewNetworkDeviceEndpointInterfacesWithDefaults instantiates a new NetworkDeviceEndpointInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NetworkDeviceEndpointInterfaces) GetData() []NetworkDeviceEndpointInterface`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkDeviceEndpointInterfaces) GetDataOk() (*[]NetworkDeviceEndpointInterface, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkDeviceEndpointInterfaces) SetData(v []NetworkDeviceEndpointInterface)`

SetData sets Data field to given value.


### GetLinks

`func (o *NetworkDeviceEndpointInterfaces) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceEndpointInterfaces) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceEndpointInterfaces) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceEndpointInterfaces) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



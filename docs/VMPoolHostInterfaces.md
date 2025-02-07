# VMPoolHostInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM Pool Host Interface ID | 
**HostId** | **float32** | VM Pool Host ID | 
**Name** | **string** | Name of the VM Pool Host Interface | 
**MacAddress** | **string** | MAC Address of the VM Pool Host Interface | 
**Fabric** | **string** | Fabric of the VM Pool Host Interface | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMPoolHostInterfaces

`func NewVMPoolHostInterfaces(id float32, hostId float32, name string, macAddress string, fabric string, links map[string]interface{}, ) *VMPoolHostInterfaces`

NewVMPoolHostInterfaces instantiates a new VMPoolHostInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolHostInterfacesWithDefaults

`func NewVMPoolHostInterfacesWithDefaults() *VMPoolHostInterfaces`

NewVMPoolHostInterfacesWithDefaults instantiates a new VMPoolHostInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPoolHostInterfaces) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPoolHostInterfaces) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPoolHostInterfaces) SetId(v float32)`

SetId sets Id field to given value.


### GetHostId

`func (o *VMPoolHostInterfaces) GetHostId() float32`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *VMPoolHostInterfaces) GetHostIdOk() (*float32, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *VMPoolHostInterfaces) SetHostId(v float32)`

SetHostId sets HostId field to given value.


### GetName

`func (o *VMPoolHostInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMPoolHostInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMPoolHostInterfaces) SetName(v string)`

SetName sets Name field to given value.


### GetMacAddress

`func (o *VMPoolHostInterfaces) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VMPoolHostInterfaces) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VMPoolHostInterfaces) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetFabric

`func (o *VMPoolHostInterfaces) GetFabric() string`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *VMPoolHostInterfaces) GetFabricOk() (*string, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *VMPoolHostInterfaces) SetFabric(v string)`

SetFabric sets Fabric field to given value.


### GetLinks

`func (o *VMPoolHostInterfaces) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMPoolHostInterfaces) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMPoolHostInterfaces) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



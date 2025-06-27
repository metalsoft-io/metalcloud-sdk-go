# EndpointInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Unique identifier for the endpoint interface | 
**Revision** | **string** | Revision number of the endpoint interface | 
**NetworkDeviceId** | **float32** | Network device id | 
**NetworkDeviceInterfaceId** | **float32** | Network device interface id | 
**NetworkDeviceInterfaceName** | **string** | Network device interface name | 
**MacAddress** | Pointer to **string** | Device interface mac address | [optional] 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewEndpointInterface

`func NewEndpointInterface(id float32, revision string, networkDeviceId float32, networkDeviceInterfaceId float32, networkDeviceInterfaceName string, createdTimestamp time.Time, updatedTimestamp time.Time, ) *EndpointInterface`

NewEndpointInterface instantiates a new EndpointInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInterfaceWithDefaults

`func NewEndpointInterfaceWithDefaults() *EndpointInterface`

NewEndpointInterfaceWithDefaults instantiates a new EndpointInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointInterface) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointInterface) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointInterface) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *EndpointInterface) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EndpointInterface) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EndpointInterface) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetNetworkDeviceId

`func (o *EndpointInterface) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *EndpointInterface) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *EndpointInterface) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkDeviceInterfaceId

`func (o *EndpointInterface) GetNetworkDeviceInterfaceId() float32`

GetNetworkDeviceInterfaceId returns the NetworkDeviceInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceIdOk

`func (o *EndpointInterface) GetNetworkDeviceInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceInterfaceIdOk returns a tuple with the NetworkDeviceInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceId

`func (o *EndpointInterface) SetNetworkDeviceInterfaceId(v float32)`

SetNetworkDeviceInterfaceId sets NetworkDeviceInterfaceId field to given value.


### GetNetworkDeviceInterfaceName

`func (o *EndpointInterface) GetNetworkDeviceInterfaceName() string`

GetNetworkDeviceInterfaceName returns the NetworkDeviceInterfaceName field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceNameOk

`func (o *EndpointInterface) GetNetworkDeviceInterfaceNameOk() (*string, bool)`

GetNetworkDeviceInterfaceNameOk returns a tuple with the NetworkDeviceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceName

`func (o *EndpointInterface) SetNetworkDeviceInterfaceName(v string)`

SetNetworkDeviceInterfaceName sets NetworkDeviceInterfaceName field to given value.


### GetMacAddress

`func (o *EndpointInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EndpointInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EndpointInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EndpointInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *EndpointInterface) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *EndpointInterface) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *EndpointInterface) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *EndpointInterface) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *EndpointInterface) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *EndpointInterface) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *EndpointInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EndpointInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



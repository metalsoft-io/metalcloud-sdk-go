# ExternalConnectionInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Unique identifier for the external connection interface | 
**NetworkDeviceInterfaceId** | **float32** | Network device interface id | 
**NetworkDeviceInterfaceName** | **string** | Network device interface name | 
**NetworkDeviceId** | **float32** | Network device id | 
**Revision** | **string** | Revision number of the external connection interface | 
**CreatedAt** | **time.Time** | The date and time the entity was created | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the entity was last updated | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExternalConnectionInterface

`func NewExternalConnectionInterface(id float32, networkDeviceInterfaceId float32, networkDeviceInterfaceName string, networkDeviceId float32, revision string, createdAt time.Time, updatedAt time.Time, ) *ExternalConnectionInterface`

NewExternalConnectionInterface instantiates a new ExternalConnectionInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalConnectionInterfaceWithDefaults

`func NewExternalConnectionInterfaceWithDefaults() *ExternalConnectionInterface`

NewExternalConnectionInterfaceWithDefaults instantiates a new ExternalConnectionInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalConnectionInterface) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalConnectionInterface) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalConnectionInterface) SetId(v float32)`

SetId sets Id field to given value.


### GetNetworkDeviceInterfaceId

`func (o *ExternalConnectionInterface) GetNetworkDeviceInterfaceId() float32`

GetNetworkDeviceInterfaceId returns the NetworkDeviceInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceIdOk

`func (o *ExternalConnectionInterface) GetNetworkDeviceInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceInterfaceIdOk returns a tuple with the NetworkDeviceInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceId

`func (o *ExternalConnectionInterface) SetNetworkDeviceInterfaceId(v float32)`

SetNetworkDeviceInterfaceId sets NetworkDeviceInterfaceId field to given value.


### GetNetworkDeviceInterfaceName

`func (o *ExternalConnectionInterface) GetNetworkDeviceInterfaceName() string`

GetNetworkDeviceInterfaceName returns the NetworkDeviceInterfaceName field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceNameOk

`func (o *ExternalConnectionInterface) GetNetworkDeviceInterfaceNameOk() (*string, bool)`

GetNetworkDeviceInterfaceNameOk returns a tuple with the NetworkDeviceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceName

`func (o *ExternalConnectionInterface) SetNetworkDeviceInterfaceName(v string)`

SetNetworkDeviceInterfaceName sets NetworkDeviceInterfaceName field to given value.


### GetNetworkDeviceId

`func (o *ExternalConnectionInterface) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *ExternalConnectionInterface) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *ExternalConnectionInterface) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetRevision

`func (o *ExternalConnectionInterface) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ExternalConnectionInterface) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ExternalConnectionInterface) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetCreatedAt

`func (o *ExternalConnectionInterface) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalConnectionInterface) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalConnectionInterface) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExternalConnectionInterface) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalConnectionInterface) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalConnectionInterface) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLinks

`func (o *ExternalConnectionInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalConnectionInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalConnectionInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalConnectionInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



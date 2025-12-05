# NetworkFabricInterconnectLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** | Revision number of the entity | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **int32** | The ID of the network fabric interconnect link | 
**InterconnectId** | **int32** | The ID of the network fabric interconnect | 
**Status** | [**NetworkFabricInterconnectLinkStatus**](NetworkFabricInterconnectLinkStatus.md) | The status of the network fabric interconnect link | 
**FabricAId** | **int32** | The ID of the network fabric to add to this interconnect | 
**FabricANetworkEquipmentId** | **int32** | The ID of the network equipment A | 
**FabricBId** | **int32** | The ID of the network fabric to add to this interconnect | 
**FabricBNetworkEquipmentId** | **int32** | The ID of the network equipment B | 

## Methods

### NewNetworkFabricInterconnectLink

`func NewNetworkFabricInterconnectLink(revision string, createdTimestamp time.Time, updatedTimestamp time.Time, id int32, interconnectId int32, status NetworkFabricInterconnectLinkStatus, fabricAId int32, fabricANetworkEquipmentId int32, fabricBId int32, fabricBNetworkEquipmentId int32, ) *NetworkFabricInterconnectLink`

NewNetworkFabricInterconnectLink instantiates a new NetworkFabricInterconnectLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectLinkWithDefaults

`func NewNetworkFabricInterconnectLinkWithDefaults() *NetworkFabricInterconnectLink`

NewNetworkFabricInterconnectLinkWithDefaults instantiates a new NetworkFabricInterconnectLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *NetworkFabricInterconnectLink) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabricInterconnectLink) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabricInterconnectLink) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetCreatedTimestamp

`func (o *NetworkFabricInterconnectLink) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabricInterconnectLink) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabricInterconnectLink) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabricInterconnectLink) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabricInterconnectLink) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabricInterconnectLink) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *NetworkFabricInterconnectLink) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkFabricInterconnectLink) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkFabricInterconnectLink) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkFabricInterconnectLink) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *NetworkFabricInterconnectLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabricInterconnectLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabricInterconnectLink) SetId(v int32)`

SetId sets Id field to given value.


### GetInterconnectId

`func (o *NetworkFabricInterconnectLink) GetInterconnectId() int32`

GetInterconnectId returns the InterconnectId field if non-nil, zero value otherwise.

### GetInterconnectIdOk

`func (o *NetworkFabricInterconnectLink) GetInterconnectIdOk() (*int32, bool)`

GetInterconnectIdOk returns a tuple with the InterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectId

`func (o *NetworkFabricInterconnectLink) SetInterconnectId(v int32)`

SetInterconnectId sets InterconnectId field to given value.


### GetStatus

`func (o *NetworkFabricInterconnectLink) GetStatus() NetworkFabricInterconnectLinkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricInterconnectLink) GetStatusOk() (*NetworkFabricInterconnectLinkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricInterconnectLink) SetStatus(v NetworkFabricInterconnectLinkStatus)`

SetStatus sets Status field to given value.


### GetFabricAId

`func (o *NetworkFabricInterconnectLink) GetFabricAId() int32`

GetFabricAId returns the FabricAId field if non-nil, zero value otherwise.

### GetFabricAIdOk

`func (o *NetworkFabricInterconnectLink) GetFabricAIdOk() (*int32, bool)`

GetFabricAIdOk returns a tuple with the FabricAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricAId

`func (o *NetworkFabricInterconnectLink) SetFabricAId(v int32)`

SetFabricAId sets FabricAId field to given value.


### GetFabricANetworkEquipmentId

`func (o *NetworkFabricInterconnectLink) GetFabricANetworkEquipmentId() int32`

GetFabricANetworkEquipmentId returns the FabricANetworkEquipmentId field if non-nil, zero value otherwise.

### GetFabricANetworkEquipmentIdOk

`func (o *NetworkFabricInterconnectLink) GetFabricANetworkEquipmentIdOk() (*int32, bool)`

GetFabricANetworkEquipmentIdOk returns a tuple with the FabricANetworkEquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricANetworkEquipmentId

`func (o *NetworkFabricInterconnectLink) SetFabricANetworkEquipmentId(v int32)`

SetFabricANetworkEquipmentId sets FabricANetworkEquipmentId field to given value.


### GetFabricBId

`func (o *NetworkFabricInterconnectLink) GetFabricBId() int32`

GetFabricBId returns the FabricBId field if non-nil, zero value otherwise.

### GetFabricBIdOk

`func (o *NetworkFabricInterconnectLink) GetFabricBIdOk() (*int32, bool)`

GetFabricBIdOk returns a tuple with the FabricBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricBId

`func (o *NetworkFabricInterconnectLink) SetFabricBId(v int32)`

SetFabricBId sets FabricBId field to given value.


### GetFabricBNetworkEquipmentId

`func (o *NetworkFabricInterconnectLink) GetFabricBNetworkEquipmentId() int32`

GetFabricBNetworkEquipmentId returns the FabricBNetworkEquipmentId field if non-nil, zero value otherwise.

### GetFabricBNetworkEquipmentIdOk

`func (o *NetworkFabricInterconnectLink) GetFabricBNetworkEquipmentIdOk() (*int32, bool)`

GetFabricBNetworkEquipmentIdOk returns a tuple with the FabricBNetworkEquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricBNetworkEquipmentId

`func (o *NetworkFabricInterconnectLink) SetFabricBNetworkEquipmentId(v int32)`

SetFabricBNetworkEquipmentId sets FabricBNetworkEquipmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



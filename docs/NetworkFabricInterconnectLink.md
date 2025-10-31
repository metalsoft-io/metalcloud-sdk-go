# NetworkFabricInterconnectLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** | Revision number of the entity | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **int32** | The ID of the network fabric interconnect link | 
**Status** | [**NetworkFabricInterconnectLinkStatus**](NetworkFabricInterconnectLinkStatus.md) | The status of the network fabric interconnect link | 
**FabricId** | **int32** | The ID of the network fabric to add to this interconnect | 
**InterconnectId** | **int32** | The ID of the network fabric interconnect | 

## Methods

### NewNetworkFabricInterconnectLink

`func NewNetworkFabricInterconnectLink(revision string, createdTimestamp time.Time, updatedTimestamp time.Time, id int32, status NetworkFabricInterconnectLinkStatus, fabricId int32, interconnectId int32, ) *NetworkFabricInterconnectLink`

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


### GetFabricId

`func (o *NetworkFabricInterconnectLink) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NetworkFabricInterconnectLink) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NetworkFabricInterconnectLink) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



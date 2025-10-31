# NetworkFabricInterconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The network fabric interconnect ID | 
**InterconnectType** | [**NetworkFabricInterconnectType**](NetworkFabricInterconnectType.md) | The network fabric interconnect type | 
**Label** | **string** | Unique label for the network fabric interconnect | 
**Name** | Pointer to **string** | Name of the network fabric interconnect | [optional] 
**Description** | Pointer to **string** | Short description of the network fabric interconnect | [optional] 
**Revision** | **string** | Revision number of the entity | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkFabricInterconnect

`func NewNetworkFabricInterconnect(id string, interconnectType NetworkFabricInterconnectType, label string, revision string, createdTimestamp time.Time, updatedTimestamp time.Time, ) *NetworkFabricInterconnect`

NewNetworkFabricInterconnect instantiates a new NetworkFabricInterconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectWithDefaults

`func NewNetworkFabricInterconnectWithDefaults() *NetworkFabricInterconnect`

NewNetworkFabricInterconnectWithDefaults instantiates a new NetworkFabricInterconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkFabricInterconnect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkFabricInterconnect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkFabricInterconnect) SetId(v string)`

SetId sets Id field to given value.


### GetInterconnectType

`func (o *NetworkFabricInterconnect) GetInterconnectType() NetworkFabricInterconnectType`

GetInterconnectType returns the InterconnectType field if non-nil, zero value otherwise.

### GetInterconnectTypeOk

`func (o *NetworkFabricInterconnect) GetInterconnectTypeOk() (*NetworkFabricInterconnectType, bool)`

GetInterconnectTypeOk returns a tuple with the InterconnectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectType

`func (o *NetworkFabricInterconnect) SetInterconnectType(v NetworkFabricInterconnectType)`

SetInterconnectType sets InterconnectType field to given value.


### GetLabel

`func (o *NetworkFabricInterconnect) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NetworkFabricInterconnect) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NetworkFabricInterconnect) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *NetworkFabricInterconnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFabricInterconnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFabricInterconnect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkFabricInterconnect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkFabricInterconnect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFabricInterconnect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFabricInterconnect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFabricInterconnect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRevision

`func (o *NetworkFabricInterconnect) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *NetworkFabricInterconnect) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *NetworkFabricInterconnect) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetCreatedTimestamp

`func (o *NetworkFabricInterconnect) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkFabricInterconnect) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkFabricInterconnect) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *NetworkFabricInterconnect) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *NetworkFabricInterconnect) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *NetworkFabricInterconnect) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *NetworkFabricInterconnect) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkFabricInterconnect) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkFabricInterconnect) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkFabricInterconnect) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



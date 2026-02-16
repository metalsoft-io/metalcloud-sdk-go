# LogicalNetworkInterconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** | Revision number of the entity | 
**Id** | **string** | Unique identifier for the logical network interconnect | 
**Label** | **string** | Unique label for the logical network interconnect | 
**Name** | **string** | Name of the logical network interconnect | 
**Annotations** | Pointer to **map[string]string** | JSON object containing additional metadata or annotations | [optional] 
**Kind** | [**LogicalNetworkInterconnectKind**](LogicalNetworkInterconnectKind.md) | Kind of the logical network interconnect | 
**FabricInterconnectId** | **int32** | Fabric Interconnect identifier | 
**TransportId** | Pointer to **int32** | Transport ID allocated from range 999999999-900000000 (descending) | [optional] 
**Status** | [**LogicalNetworkInterconnectStatus**](LogicalNetworkInterconnectStatus.md) | Status of the logical network interconnect | 
**CreatedAt** | **time.Time** | The date and time the entity was created | [readonly] 
**UpdatedAt** | **time.Time** | The date and time the entity was last updated | [readonly] 

## Methods

### NewLogicalNetworkInterconnect

`func NewLogicalNetworkInterconnect(revision string, id string, label string, name string, kind LogicalNetworkInterconnectKind, fabricInterconnectId int32, status LogicalNetworkInterconnectStatus, createdAt time.Time, updatedAt time.Time, ) *LogicalNetworkInterconnect`

NewLogicalNetworkInterconnect instantiates a new LogicalNetworkInterconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkInterconnectWithDefaults

`func NewLogicalNetworkInterconnectWithDefaults() *LogicalNetworkInterconnect`

NewLogicalNetworkInterconnectWithDefaults instantiates a new LogicalNetworkInterconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *LogicalNetworkInterconnect) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkInterconnect) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkInterconnect) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetId

`func (o *LogicalNetworkInterconnect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkInterconnect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkInterconnect) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *LogicalNetworkInterconnect) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LogicalNetworkInterconnect) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LogicalNetworkInterconnect) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *LogicalNetworkInterconnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalNetworkInterconnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalNetworkInterconnect) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *LogicalNetworkInterconnect) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *LogicalNetworkInterconnect) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *LogicalNetworkInterconnect) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *LogicalNetworkInterconnect) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *LogicalNetworkInterconnect) GetKind() LogicalNetworkInterconnectKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LogicalNetworkInterconnect) GetKindOk() (*LogicalNetworkInterconnectKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LogicalNetworkInterconnect) SetKind(v LogicalNetworkInterconnectKind)`

SetKind sets Kind field to given value.


### GetFabricInterconnectId

`func (o *LogicalNetworkInterconnect) GetFabricInterconnectId() int32`

GetFabricInterconnectId returns the FabricInterconnectId field if non-nil, zero value otherwise.

### GetFabricInterconnectIdOk

`func (o *LogicalNetworkInterconnect) GetFabricInterconnectIdOk() (*int32, bool)`

GetFabricInterconnectIdOk returns a tuple with the FabricInterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricInterconnectId

`func (o *LogicalNetworkInterconnect) SetFabricInterconnectId(v int32)`

SetFabricInterconnectId sets FabricInterconnectId field to given value.


### GetTransportId

`func (o *LogicalNetworkInterconnect) GetTransportId() int32`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *LogicalNetworkInterconnect) GetTransportIdOk() (*int32, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *LogicalNetworkInterconnect) SetTransportId(v int32)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *LogicalNetworkInterconnect) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.

### GetStatus

`func (o *LogicalNetworkInterconnect) GetStatus() LogicalNetworkInterconnectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogicalNetworkInterconnect) GetStatusOk() (*LogicalNetworkInterconnectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogicalNetworkInterconnect) SetStatus(v LogicalNetworkInterconnectStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *LogicalNetworkInterconnect) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogicalNetworkInterconnect) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogicalNetworkInterconnect) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogicalNetworkInterconnect) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogicalNetworkInterconnect) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogicalNetworkInterconnect) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



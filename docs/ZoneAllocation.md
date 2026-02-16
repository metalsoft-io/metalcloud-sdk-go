# ZoneAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**Status** | [**ResourceAllocationStatus**](ResourceAllocationStatus.md) |  | 
**ZoneName** | **string** |  | 

## Methods

### NewZoneAllocation

`func NewZoneAllocation(id int32, scope ResourceScope, status ResourceAllocationStatus, zoneName string, ) *ZoneAllocation`

NewZoneAllocation instantiates a new ZoneAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAllocationWithDefaults

`func NewZoneAllocationWithDefaults() *ZoneAllocation`

NewZoneAllocationWithDefaults instantiates a new ZoneAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneAllocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneAllocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneAllocation) SetId(v int32)`

SetId sets Id field to given value.


### GetScope

`func (o *ZoneAllocation) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ZoneAllocation) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ZoneAllocation) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *ZoneAllocation) GetStatus() ResourceAllocationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneAllocation) GetStatusOk() (*ResourceAllocationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneAllocation) SetStatus(v ResourceAllocationStatus)`

SetStatus sets Status field to given value.


### GetZoneName

`func (o *ZoneAllocation) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ZoneAllocation) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ZoneAllocation) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



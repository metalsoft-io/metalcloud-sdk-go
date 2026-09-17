# DhcpReservationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | MAC address to match (required for a manual allocation), as six colon-separated hex octets. Stored uppercase. | [optional] 
**CircuitId** | Pointer to **string** | Option 82 circuit id to match (circuit-strict) | [optional] 
**DeviceType** | Pointer to **string** | Device type guard/filter | [optional] 
**Vendor** | Pointer to **string** | Vendor guard/filter | [optional] 
**Giaddr** | Pointer to **string** | Relay gateway address (giaddr) to match | [optional] 
**Priority** | Pointer to **int32** | Tiebreaker among equally-specific matches (lower &#x3D; higher precedence); null sorts last | [optional] 
**PlannedDeviceId** | Pointer to **NullableInt64** | Optional link to a planned device record. Not used yet; null when unset (0 is treated as unset). | [optional] 
**Allocation** | Pointer to [**DhcpReservationAllocation**](DhcpReservationAllocation.md) |  | [optional] 

## Methods

### NewDhcpReservationUpdate

`func NewDhcpReservationUpdate() *DhcpReservationUpdate`

NewDhcpReservationUpdate instantiates a new DhcpReservationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationUpdateWithDefaults

`func NewDhcpReservationUpdateWithDefaults() *DhcpReservationUpdate`

NewDhcpReservationUpdateWithDefaults instantiates a new DhcpReservationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *DhcpReservationUpdate) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DhcpReservationUpdate) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DhcpReservationUpdate) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DhcpReservationUpdate) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetCircuitId

`func (o *DhcpReservationUpdate) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *DhcpReservationUpdate) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *DhcpReservationUpdate) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *DhcpReservationUpdate) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetDeviceType

`func (o *DhcpReservationUpdate) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DhcpReservationUpdate) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DhcpReservationUpdate) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DhcpReservationUpdate) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetVendor

`func (o *DhcpReservationUpdate) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DhcpReservationUpdate) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DhcpReservationUpdate) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DhcpReservationUpdate) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetGiaddr

`func (o *DhcpReservationUpdate) GetGiaddr() string`

GetGiaddr returns the Giaddr field if non-nil, zero value otherwise.

### GetGiaddrOk

`func (o *DhcpReservationUpdate) GetGiaddrOk() (*string, bool)`

GetGiaddrOk returns a tuple with the Giaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiaddr

`func (o *DhcpReservationUpdate) SetGiaddr(v string)`

SetGiaddr sets Giaddr field to given value.

### HasGiaddr

`func (o *DhcpReservationUpdate) HasGiaddr() bool`

HasGiaddr returns a boolean if a field has been set.

### GetPriority

`func (o *DhcpReservationUpdate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DhcpReservationUpdate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DhcpReservationUpdate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DhcpReservationUpdate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPlannedDeviceId

`func (o *DhcpReservationUpdate) GetPlannedDeviceId() int64`

GetPlannedDeviceId returns the PlannedDeviceId field if non-nil, zero value otherwise.

### GetPlannedDeviceIdOk

`func (o *DhcpReservationUpdate) GetPlannedDeviceIdOk() (*int64, bool)`

GetPlannedDeviceIdOk returns a tuple with the PlannedDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDeviceId

`func (o *DhcpReservationUpdate) SetPlannedDeviceId(v int64)`

SetPlannedDeviceId sets PlannedDeviceId field to given value.

### HasPlannedDeviceId

`func (o *DhcpReservationUpdate) HasPlannedDeviceId() bool`

HasPlannedDeviceId returns a boolean if a field has been set.

### SetPlannedDeviceIdNil

`func (o *DhcpReservationUpdate) SetPlannedDeviceIdNil(b bool)`

 SetPlannedDeviceIdNil sets the value for PlannedDeviceId to be an explicit nil

### UnsetPlannedDeviceId
`func (o *DhcpReservationUpdate) UnsetPlannedDeviceId()`

UnsetPlannedDeviceId ensures that no value is present for PlannedDeviceId, not even an explicit nil
### GetAllocation

`func (o *DhcpReservationUpdate) GetAllocation() DhcpReservationAllocation`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *DhcpReservationUpdate) GetAllocationOk() (*DhcpReservationAllocation, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *DhcpReservationUpdate) SetAllocation(v DhcpReservationAllocation)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *DhcpReservationUpdate) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



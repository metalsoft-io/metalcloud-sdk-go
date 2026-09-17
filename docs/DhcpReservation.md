# DhcpReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The reservation id | 
**Revision** | **int64** | Optimistic-lock revision, returned as the ETag. Echo it back in the If-Match header on update/delete. | 
**SiteId** | **int64** | The site (datacenter) id | 
**IpVersion** | **string** | IP version this reservation applies to | 
**MacAddress** | Pointer to **string** | MAC address to match (required for a manual allocation), as six colon-separated hex octets. Stored uppercase. | [optional] 
**CircuitId** | Pointer to **string** | Option 82 circuit id to match (circuit-strict) | [optional] 
**DeviceType** | Pointer to **string** | Device type guard/filter | [optional] 
**Vendor** | Pointer to **string** | Vendor guard/filter | [optional] 
**Giaddr** | Pointer to **string** | Relay gateway address (giaddr) to match | [optional] 
**Priority** | Pointer to **int32** | Tiebreaker among equally-specific matches (lower &#x3D; higher precedence); null sorts last | [optional] 
**PlannedDeviceId** | Pointer to **NullableInt64** | Optional link to a planned device record. Not used yet; null when unset (0 is treated as unset). | [optional] 
**Allocation** | [**DhcpReservationAllocation**](DhcpReservationAllocation.md) |  | 

## Methods

### NewDhcpReservation

`func NewDhcpReservation(id int64, revision int64, siteId int64, ipVersion string, allocation DhcpReservationAllocation, ) *DhcpReservation`

NewDhcpReservation instantiates a new DhcpReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationWithDefaults

`func NewDhcpReservationWithDefaults() *DhcpReservation`

NewDhcpReservationWithDefaults instantiates a new DhcpReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DhcpReservation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DhcpReservation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DhcpReservation) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *DhcpReservation) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DhcpReservation) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DhcpReservation) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetSiteId

`func (o *DhcpReservation) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DhcpReservation) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DhcpReservation) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.


### GetIpVersion

`func (o *DhcpReservation) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *DhcpReservation) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *DhcpReservation) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetMacAddress

`func (o *DhcpReservation) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DhcpReservation) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DhcpReservation) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DhcpReservation) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetCircuitId

`func (o *DhcpReservation) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *DhcpReservation) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *DhcpReservation) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *DhcpReservation) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetDeviceType

`func (o *DhcpReservation) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DhcpReservation) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DhcpReservation) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DhcpReservation) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetVendor

`func (o *DhcpReservation) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DhcpReservation) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DhcpReservation) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DhcpReservation) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetGiaddr

`func (o *DhcpReservation) GetGiaddr() string`

GetGiaddr returns the Giaddr field if non-nil, zero value otherwise.

### GetGiaddrOk

`func (o *DhcpReservation) GetGiaddrOk() (*string, bool)`

GetGiaddrOk returns a tuple with the Giaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiaddr

`func (o *DhcpReservation) SetGiaddr(v string)`

SetGiaddr sets Giaddr field to given value.

### HasGiaddr

`func (o *DhcpReservation) HasGiaddr() bool`

HasGiaddr returns a boolean if a field has been set.

### GetPriority

`func (o *DhcpReservation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DhcpReservation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DhcpReservation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DhcpReservation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPlannedDeviceId

`func (o *DhcpReservation) GetPlannedDeviceId() int64`

GetPlannedDeviceId returns the PlannedDeviceId field if non-nil, zero value otherwise.

### GetPlannedDeviceIdOk

`func (o *DhcpReservation) GetPlannedDeviceIdOk() (*int64, bool)`

GetPlannedDeviceIdOk returns a tuple with the PlannedDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDeviceId

`func (o *DhcpReservation) SetPlannedDeviceId(v int64)`

SetPlannedDeviceId sets PlannedDeviceId field to given value.

### HasPlannedDeviceId

`func (o *DhcpReservation) HasPlannedDeviceId() bool`

HasPlannedDeviceId returns a boolean if a field has been set.

### SetPlannedDeviceIdNil

`func (o *DhcpReservation) SetPlannedDeviceIdNil(b bool)`

 SetPlannedDeviceIdNil sets the value for PlannedDeviceId to be an explicit nil

### UnsetPlannedDeviceId
`func (o *DhcpReservation) UnsetPlannedDeviceId()`

UnsetPlannedDeviceId ensures that no value is present for PlannedDeviceId, not even an explicit nil
### GetAllocation

`func (o *DhcpReservation) GetAllocation() DhcpReservationAllocation`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *DhcpReservation) GetAllocationOk() (*DhcpReservationAllocation, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *DhcpReservation) SetAllocation(v DhcpReservationAllocation)`

SetAllocation sets Allocation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



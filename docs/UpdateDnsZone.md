# UpdateDnsZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The DNS zone description | [optional] 
**ZoneType** | Pointer to **string** | The type of zone, &#39;master&#39; is controlled by the application,                   &#39;slave&#39; is controlled by an external DNS server. | [optional] [default to "master"]
**SoaEmail** | Pointer to **string** | The email address of the DNS zone administrator | [optional] [default to "admin.<zone_name>"]
**Ttl** | Pointer to **int32** | TTL (Time to Live) for the DNS zone. | [optional] 
**NameServers** | Pointer to **[]string** | The name servers for this DNS zone. This is a list of DNS servers that are authoritative for the zone. | [optional] 
**Status** | Pointer to **string** | The current status of the DNS zone.       Status: READY           - The initial status when the DNS zone is registered in the system.           - The DNS zone can be updated or deleted.       Status: ACTIVE           - The DNS zone is configured on the DNS server(s).           - The DNS zone can be updated or deleted.           - DNS records can be associated with the DNS zone.       Status: USED           - The DNS zone has DNS records associated with it.           - The DNS zone is part of at least one infrastructure.       Status: ARCHIVED           - The DNS zone is no longer managed by the system.           - The DNS zone is kept in the system for historical purposes.           - The DNS zone cannot be updated or deleted (validation will prevent this).           - No DNS records can be associated, updated, or deleted. | [optional] [default to "ready"]
**Tags** | Pointer to **[]string** | The tags associated with the DNS zone | [optional] 

## Methods

### NewUpdateDnsZone

`func NewUpdateDnsZone() *UpdateDnsZone`

NewUpdateDnsZone instantiates a new UpdateDnsZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDnsZoneWithDefaults

`func NewUpdateDnsZoneWithDefaults() *UpdateDnsZone`

NewUpdateDnsZoneWithDefaults instantiates a new UpdateDnsZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateDnsZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDnsZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDnsZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDnsZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneType

`func (o *UpdateDnsZone) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *UpdateDnsZone) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *UpdateDnsZone) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *UpdateDnsZone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetSoaEmail

`func (o *UpdateDnsZone) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *UpdateDnsZone) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *UpdateDnsZone) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *UpdateDnsZone) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetTtl

`func (o *UpdateDnsZone) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *UpdateDnsZone) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *UpdateDnsZone) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *UpdateDnsZone) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetNameServers

`func (o *UpdateDnsZone) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *UpdateDnsZone) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *UpdateDnsZone) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *UpdateDnsZone) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateDnsZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDnsZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDnsZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDnsZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *UpdateDnsZone) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDnsZone) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDnsZone) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDnsZone) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



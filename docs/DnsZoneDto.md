# DnsZoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The DNS zone ID | [readonly] 
**Label** | **string** | The DNS zone label. It must be unique | [readonly] 
**Description** | Pointer to **string** | The DNS zone description | [optional] 
**ZoneName** | **string** | The name of the DNS zone (without a terminating dot).                   The DNS zone name, must be unique | 
**ZoneType** | **string** | The type of zone, &#39;master&#39; is controlled by the application,                   &#39;slave&#39; is controlled by an external DNS server. | [default to "master"]
**SoaEmail** | **string** | The email address of the DNS zone administrator | [default to "admin.<zone_name>"]
**SoaSerial** | **int32** | The serial number of the DNS zone | [readonly] 
**Ttl** | **int32** | TTL (Time to Live) for the DNS zone. | 
**NameServers** | **[]string** | The name servers for this DNS zone. This is a list of DNS servers that are authoritative for the zone. | 
**IsDefault** | **bool** | The default DNS zone to be used in deployments. | 
**Status** | **string** | The current status of the DNS zone.       Status: READY           - The initial status when the DNS zone is registered in the system.           - The DNS zone can be updated or deleted.       Status: ACTIVE           - The DNS zone is configured on the DNS server(s).           - The DNS zone can be updated or deleted.           - DNS records can be associated with the DNS zone.       Status: USED           - The DNS zone has DNS records associated with it.           - The DNS zone is part of at least one infrastructure.       Status: ARCHIVED           - The DNS zone is no longer managed by the system.           - The DNS zone is kept in the system for historical purposes.           - The DNS zone cannot be updated or deleted (validation will prevent this).           - No DNS records can be associated, updated, or deleted. | [default to "ready"]
**Tags** | Pointer to **[]string** | The tags associated with the DNS zone | [optional] 
**Revision** | **int32** | The revision number of the DNS zone | [readonly] 
**CreatedBy** | **int32** | The user ID of the user who created the DNS zone | 
**UpdatedBy** | Pointer to **int32** | The user ID of the user who last modified the DNS zone | [optional] 
**CreatedAt** | **time.Time** | The date and time the DNS zone was created | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the DNS zone was last updated | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewDnsZoneDto

`func NewDnsZoneDto(id int32, label string, zoneName string, zoneType string, soaEmail string, soaSerial int32, ttl int32, nameServers []string, isDefault bool, status string, revision int32, createdBy int32, createdAt time.Time, ) *DnsZoneDto`

NewDnsZoneDto instantiates a new DnsZoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneDtoWithDefaults

`func NewDnsZoneDtoWithDefaults() *DnsZoneDto`

NewDnsZoneDtoWithDefaults instantiates a new DnsZoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsZoneDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsZoneDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsZoneDto) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *DnsZoneDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DnsZoneDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DnsZoneDto) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *DnsZoneDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DnsZoneDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DnsZoneDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DnsZoneDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneName

`func (o *DnsZoneDto) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsZoneDto) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsZoneDto) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetZoneType

`func (o *DnsZoneDto) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *DnsZoneDto) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *DnsZoneDto) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.


### GetSoaEmail

`func (o *DnsZoneDto) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *DnsZoneDto) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *DnsZoneDto) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.


### GetSoaSerial

`func (o *DnsZoneDto) GetSoaSerial() int32`

GetSoaSerial returns the SoaSerial field if non-nil, zero value otherwise.

### GetSoaSerialOk

`func (o *DnsZoneDto) GetSoaSerialOk() (*int32, bool)`

GetSoaSerialOk returns a tuple with the SoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaSerial

`func (o *DnsZoneDto) SetSoaSerial(v int32)`

SetSoaSerial sets SoaSerial field to given value.


### GetTtl

`func (o *DnsZoneDto) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsZoneDto) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsZoneDto) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetNameServers

`func (o *DnsZoneDto) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *DnsZoneDto) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *DnsZoneDto) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.


### GetIsDefault

`func (o *DnsZoneDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DnsZoneDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DnsZoneDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetStatus

`func (o *DnsZoneDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DnsZoneDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DnsZoneDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *DnsZoneDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DnsZoneDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DnsZoneDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DnsZoneDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRevision

`func (o *DnsZoneDto) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DnsZoneDto) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DnsZoneDto) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatedBy

`func (o *DnsZoneDto) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DnsZoneDto) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DnsZoneDto) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *DnsZoneDto) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DnsZoneDto) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DnsZoneDto) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DnsZoneDto) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DnsZoneDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DnsZoneDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DnsZoneDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DnsZoneDto) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DnsZoneDto) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DnsZoneDto) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DnsZoneDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *DnsZoneDto) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DnsZoneDto) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DnsZoneDto) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DnsZoneDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



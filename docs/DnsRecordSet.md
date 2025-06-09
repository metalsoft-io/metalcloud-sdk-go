# DnsRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The DNS Resource Record Set (RRSet) ID | [readonly] 
**SiteId** | **int32** | The site ID | 
**InfrastructureId** | **int32** | The infrastructure ID | 
**ZoneId** | **int32** | The ID of the DNS zone | 
**ZoneName** | **string** | The name of the DNS zone (without a terminating dot) | 
**Name** | **string** | DNS Name for the RecordSet | 
**Type** | **string** | The type of DNS record (e.g., A, AAAA, CNAME, NS, PTR, TXT, SOA) | 
**Ttl** | Pointer to **int32** | TTL (Time to Live) for the DNS Record Set. | [optional] [default to 3600]
**Records** | **[]string** | The record data for this DNS record set | 
**Tags** | Pointer to **[]string** | The tags associated with the DNS Record Set | [optional] 
**Revision** | **int32** | The revision number of the DNS Record Set | [readonly] 
**CreatedBy** | **int32** | The user ID of the user who created the DNS Record Set | 
**UpdatedBy** | Pointer to **int32** | The user ID of the user who last modified the DNS Record Set | [optional] 
**CreatedAt** | **time.Time** | The date and time the DNS Record Set was created | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the DNS Record Set was last updated | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewDnsRecordSet

`func NewDnsRecordSet(id int32, siteId int32, infrastructureId int32, zoneId int32, zoneName string, name string, type_ string, records []string, revision int32, createdBy int32, createdAt time.Time, ) *DnsRecordSet`

NewDnsRecordSet instantiates a new DnsRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordSetWithDefaults

`func NewDnsRecordSetWithDefaults() *DnsRecordSet`

NewDnsRecordSetWithDefaults instantiates a new DnsRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsRecordSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecordSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecordSet) SetId(v int32)`

SetId sets Id field to given value.


### GetSiteId

`func (o *DnsRecordSet) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DnsRecordSet) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DnsRecordSet) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.


### GetInfrastructureId

`func (o *DnsRecordSet) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *DnsRecordSet) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *DnsRecordSet) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetZoneId

`func (o *DnsRecordSet) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsRecordSet) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsRecordSet) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.


### GetZoneName

`func (o *DnsRecordSet) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsRecordSet) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsRecordSet) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetName

`func (o *DnsRecordSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordSet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DnsRecordSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsRecordSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsRecordSet) SetType(v string)`

SetType sets Type field to given value.


### GetTtl

`func (o *DnsRecordSet) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordSet) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordSet) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordSet) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetRecords

`func (o *DnsRecordSet) GetRecords() []string`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *DnsRecordSet) GetRecordsOk() (*[]string, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *DnsRecordSet) SetRecords(v []string)`

SetRecords sets Records field to given value.


### GetTags

`func (o *DnsRecordSet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DnsRecordSet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DnsRecordSet) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DnsRecordSet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRevision

`func (o *DnsRecordSet) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DnsRecordSet) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DnsRecordSet) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatedBy

`func (o *DnsRecordSet) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DnsRecordSet) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DnsRecordSet) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *DnsRecordSet) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DnsRecordSet) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DnsRecordSet) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DnsRecordSet) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DnsRecordSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DnsRecordSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DnsRecordSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DnsRecordSet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DnsRecordSet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DnsRecordSet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DnsRecordSet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *DnsRecordSet) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DnsRecordSet) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DnsRecordSet) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DnsRecordSet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



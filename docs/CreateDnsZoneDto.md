# CreateDnsZoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The DNS zone description | [optional] 
**ZoneName** | **string** | The name of the DNS zone (without a terminating dot).                   The DNS zone name must be unique and follow the DNS domain naming conventions. | 
**ZoneType** | Pointer to **string** | The type of zone, &#39;master&#39; is controlled by the application,                   &#39;slave&#39; is controlled by an external DNS server. | [optional] [default to "master"]
**SoaEmail** | Pointer to **string** | The email address of the DNS zone administrator | [optional] [default to "admin.<zone_name>"]
**Ttl** | Pointer to **int32** | TTL (Time to Live) for the DNS zone. | [optional] 
**IsDefault** | **bool** | The default DNS zone to be used in deployments. | 
**NameServers** | **[]string** | The name servers for this DNS zone. This is a list of DNS servers that are authoritative for the zone. | 
**Tags** | Pointer to **[]string** | The tags associated with the DNS zone | [optional] 

## Methods

### NewCreateDnsZoneDto

`func NewCreateDnsZoneDto(zoneName string, isDefault bool, nameServers []string, ) *CreateDnsZoneDto`

NewCreateDnsZoneDto instantiates a new CreateDnsZoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsZoneDtoWithDefaults

`func NewCreateDnsZoneDtoWithDefaults() *CreateDnsZoneDto`

NewCreateDnsZoneDtoWithDefaults instantiates a new CreateDnsZoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateDnsZoneDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDnsZoneDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDnsZoneDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDnsZoneDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneName

`func (o *CreateDnsZoneDto) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *CreateDnsZoneDto) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *CreateDnsZoneDto) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetZoneType

`func (o *CreateDnsZoneDto) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *CreateDnsZoneDto) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *CreateDnsZoneDto) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *CreateDnsZoneDto) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetSoaEmail

`func (o *CreateDnsZoneDto) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *CreateDnsZoneDto) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *CreateDnsZoneDto) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *CreateDnsZoneDto) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetTtl

`func (o *CreateDnsZoneDto) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateDnsZoneDto) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateDnsZoneDto) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateDnsZoneDto) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateDnsZoneDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateDnsZoneDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateDnsZoneDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetNameServers

`func (o *CreateDnsZoneDto) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *CreateDnsZoneDto) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *CreateDnsZoneDto) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.


### GetTags

`func (o *CreateDnsZoneDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDnsZoneDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDnsZoneDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDnsZoneDto) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



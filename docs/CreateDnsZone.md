# CreateDnsZone

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

### NewCreateDnsZone

`func NewCreateDnsZone(zoneName string, isDefault bool, nameServers []string, ) *CreateDnsZone`

NewCreateDnsZone instantiates a new CreateDnsZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsZoneWithDefaults

`func NewCreateDnsZoneWithDefaults() *CreateDnsZone`

NewCreateDnsZoneWithDefaults instantiates a new CreateDnsZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateDnsZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDnsZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDnsZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDnsZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneName

`func (o *CreateDnsZone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *CreateDnsZone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *CreateDnsZone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetZoneType

`func (o *CreateDnsZone) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *CreateDnsZone) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *CreateDnsZone) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *CreateDnsZone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetSoaEmail

`func (o *CreateDnsZone) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *CreateDnsZone) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *CreateDnsZone) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *CreateDnsZone) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetTtl

`func (o *CreateDnsZone) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateDnsZone) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateDnsZone) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateDnsZone) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateDnsZone) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateDnsZone) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateDnsZone) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetNameServers

`func (o *CreateDnsZone) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *CreateDnsZone) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *CreateDnsZone) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.


### GetTags

`func (o *CreateDnsZone) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDnsZone) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDnsZone) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDnsZone) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



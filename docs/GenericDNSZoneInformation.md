# GenericDNSZoneInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | **string** | The name of the DNS zone. | 
**SoaEmail** | Pointer to **string** | The email address of the DNS zone administrator | [optional] 
**NameServers** | Pointer to **[]string** | The name servers for this DNS zone. This is a list of DNS servers that are authoritative for the zone. | [optional] 
**Ttl** | Pointer to **int32** | TTL (Time to Live) for the DNS zone. | [optional] 
**IsDefault** | Pointer to **bool** | Indicates if this DNS zone is the default zone. | [optional] 

## Methods

### NewGenericDNSZoneInformation

`func NewGenericDNSZoneInformation(zoneName string, ) *GenericDNSZoneInformation`

NewGenericDNSZoneInformation instantiates a new GenericDNSZoneInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericDNSZoneInformationWithDefaults

`func NewGenericDNSZoneInformationWithDefaults() *GenericDNSZoneInformation`

NewGenericDNSZoneInformationWithDefaults instantiates a new GenericDNSZoneInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *GenericDNSZoneInformation) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *GenericDNSZoneInformation) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *GenericDNSZoneInformation) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetSoaEmail

`func (o *GenericDNSZoneInformation) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *GenericDNSZoneInformation) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *GenericDNSZoneInformation) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *GenericDNSZoneInformation) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetNameServers

`func (o *GenericDNSZoneInformation) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *GenericDNSZoneInformation) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *GenericDNSZoneInformation) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *GenericDNSZoneInformation) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetTtl

`func (o *GenericDNSZoneInformation) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GenericDNSZoneInformation) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GenericDNSZoneInformation) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GenericDNSZoneInformation) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetIsDefault

`func (o *GenericDNSZoneInformation) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GenericDNSZoneInformation) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GenericDNSZoneInformation) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GenericDNSZoneInformation) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



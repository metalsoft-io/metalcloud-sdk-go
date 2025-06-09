# SiteConfigVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](Location.md) | Location details | [optional] 
**Repo** | Pointer to [**Repo**](Repo.md) | Repository details | [optional] 
**DnsZoneId** | Pointer to **int32** | ID of the DNS zone associated with the site | [optional] 
**DNSServers** | Pointer to **[]string** | List of DNS Servers | [optional] 
**NTPServers** | Pointer to **[]string** | List of NTP Servers | [optional] 
**NetworkDevicePolicy** | Pointer to [**NetworkDevicePolicy**](NetworkDevicePolicy.md) | Network device policies | [optional] 
**ServerPolicy** | Pointer to [**ServerPolicy**](ServerPolicy.md) | Server policies | [optional] 
**ControllerPolicy** | Pointer to [**ControllerPolicy**](ControllerPolicy.md) | Controller policies | [optional] 
**InfrastructurePolicy** | Pointer to [**InfrastructurePolicy**](InfrastructurePolicy.md) | Infrastructure policies | [optional] 

## Methods

### NewSiteConfigVariables

`func NewSiteConfigVariables() *SiteConfigVariables`

NewSiteConfigVariables instantiates a new SiteConfigVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteConfigVariablesWithDefaults

`func NewSiteConfigVariablesWithDefaults() *SiteConfigVariables`

NewSiteConfigVariablesWithDefaults instantiates a new SiteConfigVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SiteConfigVariables) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SiteConfigVariables) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SiteConfigVariables) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SiteConfigVariables) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRepo

`func (o *SiteConfigVariables) GetRepo() Repo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *SiteConfigVariables) GetRepoOk() (*Repo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *SiteConfigVariables) SetRepo(v Repo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *SiteConfigVariables) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetDnsZoneId

`func (o *SiteConfigVariables) GetDnsZoneId() int32`

GetDnsZoneId returns the DnsZoneId field if non-nil, zero value otherwise.

### GetDnsZoneIdOk

`func (o *SiteConfigVariables) GetDnsZoneIdOk() (*int32, bool)`

GetDnsZoneIdOk returns a tuple with the DnsZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneId

`func (o *SiteConfigVariables) SetDnsZoneId(v int32)`

SetDnsZoneId sets DnsZoneId field to given value.

### HasDnsZoneId

`func (o *SiteConfigVariables) HasDnsZoneId() bool`

HasDnsZoneId returns a boolean if a field has been set.

### GetDNSServers

`func (o *SiteConfigVariables) GetDNSServers() []string`

GetDNSServers returns the DNSServers field if non-nil, zero value otherwise.

### GetDNSServersOk

`func (o *SiteConfigVariables) GetDNSServersOk() (*[]string, bool)`

GetDNSServersOk returns a tuple with the DNSServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSServers

`func (o *SiteConfigVariables) SetDNSServers(v []string)`

SetDNSServers sets DNSServers field to given value.

### HasDNSServers

`func (o *SiteConfigVariables) HasDNSServers() bool`

HasDNSServers returns a boolean if a field has been set.

### GetNTPServers

`func (o *SiteConfigVariables) GetNTPServers() []string`

GetNTPServers returns the NTPServers field if non-nil, zero value otherwise.

### GetNTPServersOk

`func (o *SiteConfigVariables) GetNTPServersOk() (*[]string, bool)`

GetNTPServersOk returns a tuple with the NTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPServers

`func (o *SiteConfigVariables) SetNTPServers(v []string)`

SetNTPServers sets NTPServers field to given value.

### HasNTPServers

`func (o *SiteConfigVariables) HasNTPServers() bool`

HasNTPServers returns a boolean if a field has been set.

### GetNetworkDevicePolicy

`func (o *SiteConfigVariables) GetNetworkDevicePolicy() NetworkDevicePolicy`

GetNetworkDevicePolicy returns the NetworkDevicePolicy field if non-nil, zero value otherwise.

### GetNetworkDevicePolicyOk

`func (o *SiteConfigVariables) GetNetworkDevicePolicyOk() (*NetworkDevicePolicy, bool)`

GetNetworkDevicePolicyOk returns a tuple with the NetworkDevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePolicy

`func (o *SiteConfigVariables) SetNetworkDevicePolicy(v NetworkDevicePolicy)`

SetNetworkDevicePolicy sets NetworkDevicePolicy field to given value.

### HasNetworkDevicePolicy

`func (o *SiteConfigVariables) HasNetworkDevicePolicy() bool`

HasNetworkDevicePolicy returns a boolean if a field has been set.

### GetServerPolicy

`func (o *SiteConfigVariables) GetServerPolicy() ServerPolicy`

GetServerPolicy returns the ServerPolicy field if non-nil, zero value otherwise.

### GetServerPolicyOk

`func (o *SiteConfigVariables) GetServerPolicyOk() (*ServerPolicy, bool)`

GetServerPolicyOk returns a tuple with the ServerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPolicy

`func (o *SiteConfigVariables) SetServerPolicy(v ServerPolicy)`

SetServerPolicy sets ServerPolicy field to given value.

### HasServerPolicy

`func (o *SiteConfigVariables) HasServerPolicy() bool`

HasServerPolicy returns a boolean if a field has been set.

### GetControllerPolicy

`func (o *SiteConfigVariables) GetControllerPolicy() ControllerPolicy`

GetControllerPolicy returns the ControllerPolicy field if non-nil, zero value otherwise.

### GetControllerPolicyOk

`func (o *SiteConfigVariables) GetControllerPolicyOk() (*ControllerPolicy, bool)`

GetControllerPolicyOk returns a tuple with the ControllerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPolicy

`func (o *SiteConfigVariables) SetControllerPolicy(v ControllerPolicy)`

SetControllerPolicy sets ControllerPolicy field to given value.

### HasControllerPolicy

`func (o *SiteConfigVariables) HasControllerPolicy() bool`

HasControllerPolicy returns a boolean if a field has been set.

### GetInfrastructurePolicy

`func (o *SiteConfigVariables) GetInfrastructurePolicy() InfrastructurePolicy`

GetInfrastructurePolicy returns the InfrastructurePolicy field if non-nil, zero value otherwise.

### GetInfrastructurePolicyOk

`func (o *SiteConfigVariables) GetInfrastructurePolicyOk() (*InfrastructurePolicy, bool)`

GetInfrastructurePolicyOk returns a tuple with the InfrastructurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructurePolicy

`func (o *SiteConfigVariables) SetInfrastructurePolicy(v InfrastructurePolicy)`

SetInfrastructurePolicy sets InfrastructurePolicy field to given value.

### HasInfrastructurePolicy

`func (o *SiteConfigVariables) HasInfrastructurePolicy() bool`

HasInfrastructurePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



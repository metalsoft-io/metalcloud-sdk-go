# SiteConfigOSInstallationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoURL** | **string** | Repository details | 
**DnsZoneId** | Pointer to **int32** | ID of the DNS zone associated with the site | [optional] 
**DnsZoneName** | Pointer to **string** | The name of the DNS zone (without a terminating dot). | [optional] 
**DNSServers** | **[]string** | List of DNS Servers | 
**NTPServers** | **[]string** | List of NTP Servers | 

## Methods

### NewSiteConfigOSInstallationData

`func NewSiteConfigOSInstallationData(repoURL string, dNSServers []string, nTPServers []string, ) *SiteConfigOSInstallationData`

NewSiteConfigOSInstallationData instantiates a new SiteConfigOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteConfigOSInstallationDataWithDefaults

`func NewSiteConfigOSInstallationDataWithDefaults() *SiteConfigOSInstallationData`

NewSiteConfigOSInstallationDataWithDefaults instantiates a new SiteConfigOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoURL

`func (o *SiteConfigOSInstallationData) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *SiteConfigOSInstallationData) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *SiteConfigOSInstallationData) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.


### GetDnsZoneId

`func (o *SiteConfigOSInstallationData) GetDnsZoneId() int32`

GetDnsZoneId returns the DnsZoneId field if non-nil, zero value otherwise.

### GetDnsZoneIdOk

`func (o *SiteConfigOSInstallationData) GetDnsZoneIdOk() (*int32, bool)`

GetDnsZoneIdOk returns a tuple with the DnsZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneId

`func (o *SiteConfigOSInstallationData) SetDnsZoneId(v int32)`

SetDnsZoneId sets DnsZoneId field to given value.

### HasDnsZoneId

`func (o *SiteConfigOSInstallationData) HasDnsZoneId() bool`

HasDnsZoneId returns a boolean if a field has been set.

### GetDnsZoneName

`func (o *SiteConfigOSInstallationData) GetDnsZoneName() string`

GetDnsZoneName returns the DnsZoneName field if non-nil, zero value otherwise.

### GetDnsZoneNameOk

`func (o *SiteConfigOSInstallationData) GetDnsZoneNameOk() (*string, bool)`

GetDnsZoneNameOk returns a tuple with the DnsZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsZoneName

`func (o *SiteConfigOSInstallationData) SetDnsZoneName(v string)`

SetDnsZoneName sets DnsZoneName field to given value.

### HasDnsZoneName

`func (o *SiteConfigOSInstallationData) HasDnsZoneName() bool`

HasDnsZoneName returns a boolean if a field has been set.

### GetDNSServers

`func (o *SiteConfigOSInstallationData) GetDNSServers() []string`

GetDNSServers returns the DNSServers field if non-nil, zero value otherwise.

### GetDNSServersOk

`func (o *SiteConfigOSInstallationData) GetDNSServersOk() (*[]string, bool)`

GetDNSServersOk returns a tuple with the DNSServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSServers

`func (o *SiteConfigOSInstallationData) SetDNSServers(v []string)`

SetDNSServers sets DNSServers field to given value.


### GetNTPServers

`func (o *SiteConfigOSInstallationData) GetNTPServers() []string`

GetNTPServers returns the NTPServers field if non-nil, zero value otherwise.

### GetNTPServersOk

`func (o *SiteConfigOSInstallationData) GetNTPServersOk() (*[]string, bool)`

GetNTPServersOk returns a tuple with the NTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTPServers

`func (o *SiteConfigOSInstallationData) SetNTPServers(v []string)`

SetNTPServers sets NTPServers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the SiteConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteConfig{}

// SiteConfig struct for SiteConfig
type SiteConfig struct {
	// Location details
	Location *Location `json:"location,omitempty"`
	// Repository details
	Repo *Repo `json:"repo,omitempty"`
	// ID of the DNS zone associated with the site
	DnsZoneId *int32 `json:"dnsZoneId,omitempty"`
	// List of DNS Servers
	DNSServers []string `json:"DNSServers,omitempty"`
	// List of NTP Servers
	NTPServers []string `json:"NTPServers,omitempty"`
	// Network device policies
	NetworkDevicePolicy *NetworkDevicePolicy `json:"networkDevicePolicy,omitempty"`
	// Server policies
	ServerPolicy *ServerPolicy `json:"serverPolicy,omitempty"`
	// Controller policies
	ControllerPolicy *ControllerPolicy `json:"controllerPolicy,omitempty"`
	// Infrastructure policies
	InfrastructurePolicy *InfrastructurePolicy `json:"infrastructurePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteConfig SiteConfig

// NewSiteConfig instantiates a new SiteConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteConfig() *SiteConfig {
	this := SiteConfig{}
	return &this
}

// NewSiteConfigWithDefaults instantiates a new SiteConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteConfigWithDefaults() *SiteConfig {
	this := SiteConfig{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SiteConfig) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SiteConfig) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *SiteConfig) SetLocation(v Location) {
	o.Location = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise.
func (o *SiteConfig) GetRepo() Repo {
	if o == nil || IsNil(o.Repo) {
		var ret Repo
		return ret
	}
	return *o.Repo
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetRepoOk() (*Repo, bool) {
	if o == nil || IsNil(o.Repo) {
		return nil, false
	}
	return o.Repo, true
}

// HasRepo returns a boolean if a field has been set.
func (o *SiteConfig) HasRepo() bool {
	if o != nil && !IsNil(o.Repo) {
		return true
	}

	return false
}

// SetRepo gets a reference to the given Repo and assigns it to the Repo field.
func (o *SiteConfig) SetRepo(v Repo) {
	o.Repo = &v
}

// GetDnsZoneId returns the DnsZoneId field value if set, zero value otherwise.
func (o *SiteConfig) GetDnsZoneId() int32 {
	if o == nil || IsNil(o.DnsZoneId) {
		var ret int32
		return ret
	}
	return *o.DnsZoneId
}

// GetDnsZoneIdOk returns a tuple with the DnsZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetDnsZoneIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsZoneId) {
		return nil, false
	}
	return o.DnsZoneId, true
}

// HasDnsZoneId returns a boolean if a field has been set.
func (o *SiteConfig) HasDnsZoneId() bool {
	if o != nil && !IsNil(o.DnsZoneId) {
		return true
	}

	return false
}

// SetDnsZoneId gets a reference to the given int32 and assigns it to the DnsZoneId field.
func (o *SiteConfig) SetDnsZoneId(v int32) {
	o.DnsZoneId = &v
}

// GetDNSServers returns the DNSServers field value if set, zero value otherwise.
func (o *SiteConfig) GetDNSServers() []string {
	if o == nil || IsNil(o.DNSServers) {
		var ret []string
		return ret
	}
	return o.DNSServers
}

// GetDNSServersOk returns a tuple with the DNSServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetDNSServersOk() ([]string, bool) {
	if o == nil || IsNil(o.DNSServers) {
		return nil, false
	}
	return o.DNSServers, true
}

// HasDNSServers returns a boolean if a field has been set.
func (o *SiteConfig) HasDNSServers() bool {
	if o != nil && !IsNil(o.DNSServers) {
		return true
	}

	return false
}

// SetDNSServers gets a reference to the given []string and assigns it to the DNSServers field.
func (o *SiteConfig) SetDNSServers(v []string) {
	o.DNSServers = v
}

// GetNTPServers returns the NTPServers field value if set, zero value otherwise.
func (o *SiteConfig) GetNTPServers() []string {
	if o == nil || IsNil(o.NTPServers) {
		var ret []string
		return ret
	}
	return o.NTPServers
}

// GetNTPServersOk returns a tuple with the NTPServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetNTPServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NTPServers) {
		return nil, false
	}
	return o.NTPServers, true
}

// HasNTPServers returns a boolean if a field has been set.
func (o *SiteConfig) HasNTPServers() bool {
	if o != nil && !IsNil(o.NTPServers) {
		return true
	}

	return false
}

// SetNTPServers gets a reference to the given []string and assigns it to the NTPServers field.
func (o *SiteConfig) SetNTPServers(v []string) {
	o.NTPServers = v
}

// GetNetworkDevicePolicy returns the NetworkDevicePolicy field value if set, zero value otherwise.
func (o *SiteConfig) GetNetworkDevicePolicy() NetworkDevicePolicy {
	if o == nil || IsNil(o.NetworkDevicePolicy) {
		var ret NetworkDevicePolicy
		return ret
	}
	return *o.NetworkDevicePolicy
}

// GetNetworkDevicePolicyOk returns a tuple with the NetworkDevicePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetNetworkDevicePolicyOk() (*NetworkDevicePolicy, bool) {
	if o == nil || IsNil(o.NetworkDevicePolicy) {
		return nil, false
	}
	return o.NetworkDevicePolicy, true
}

// HasNetworkDevicePolicy returns a boolean if a field has been set.
func (o *SiteConfig) HasNetworkDevicePolicy() bool {
	if o != nil && !IsNil(o.NetworkDevicePolicy) {
		return true
	}

	return false
}

// SetNetworkDevicePolicy gets a reference to the given NetworkDevicePolicy and assigns it to the NetworkDevicePolicy field.
func (o *SiteConfig) SetNetworkDevicePolicy(v NetworkDevicePolicy) {
	o.NetworkDevicePolicy = &v
}

// GetServerPolicy returns the ServerPolicy field value if set, zero value otherwise.
func (o *SiteConfig) GetServerPolicy() ServerPolicy {
	if o == nil || IsNil(o.ServerPolicy) {
		var ret ServerPolicy
		return ret
	}
	return *o.ServerPolicy
}

// GetServerPolicyOk returns a tuple with the ServerPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetServerPolicyOk() (*ServerPolicy, bool) {
	if o == nil || IsNil(o.ServerPolicy) {
		return nil, false
	}
	return o.ServerPolicy, true
}

// HasServerPolicy returns a boolean if a field has been set.
func (o *SiteConfig) HasServerPolicy() bool {
	if o != nil && !IsNil(o.ServerPolicy) {
		return true
	}

	return false
}

// SetServerPolicy gets a reference to the given ServerPolicy and assigns it to the ServerPolicy field.
func (o *SiteConfig) SetServerPolicy(v ServerPolicy) {
	o.ServerPolicy = &v
}

// GetControllerPolicy returns the ControllerPolicy field value if set, zero value otherwise.
func (o *SiteConfig) GetControllerPolicy() ControllerPolicy {
	if o == nil || IsNil(o.ControllerPolicy) {
		var ret ControllerPolicy
		return ret
	}
	return *o.ControllerPolicy
}

// GetControllerPolicyOk returns a tuple with the ControllerPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetControllerPolicyOk() (*ControllerPolicy, bool) {
	if o == nil || IsNil(o.ControllerPolicy) {
		return nil, false
	}
	return o.ControllerPolicy, true
}

// HasControllerPolicy returns a boolean if a field has been set.
func (o *SiteConfig) HasControllerPolicy() bool {
	if o != nil && !IsNil(o.ControllerPolicy) {
		return true
	}

	return false
}

// SetControllerPolicy gets a reference to the given ControllerPolicy and assigns it to the ControllerPolicy field.
func (o *SiteConfig) SetControllerPolicy(v ControllerPolicy) {
	o.ControllerPolicy = &v
}

// GetInfrastructurePolicy returns the InfrastructurePolicy field value if set, zero value otherwise.
func (o *SiteConfig) GetInfrastructurePolicy() InfrastructurePolicy {
	if o == nil || IsNil(o.InfrastructurePolicy) {
		var ret InfrastructurePolicy
		return ret
	}
	return *o.InfrastructurePolicy
}

// GetInfrastructurePolicyOk returns a tuple with the InfrastructurePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetInfrastructurePolicyOk() (*InfrastructurePolicy, bool) {
	if o == nil || IsNil(o.InfrastructurePolicy) {
		return nil, false
	}
	return o.InfrastructurePolicy, true
}

// HasInfrastructurePolicy returns a boolean if a field has been set.
func (o *SiteConfig) HasInfrastructurePolicy() bool {
	if o != nil && !IsNil(o.InfrastructurePolicy) {
		return true
	}

	return false
}

// SetInfrastructurePolicy gets a reference to the given InfrastructurePolicy and assigns it to the InfrastructurePolicy field.
func (o *SiteConfig) SetInfrastructurePolicy(v InfrastructurePolicy) {
	o.InfrastructurePolicy = &v
}

func (o SiteConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Repo) {
		toSerialize["repo"] = o.Repo
	}
	if !IsNil(o.DnsZoneId) {
		toSerialize["dnsZoneId"] = o.DnsZoneId
	}
	if !IsNil(o.DNSServers) {
		toSerialize["DNSServers"] = o.DNSServers
	}
	if !IsNil(o.NTPServers) {
		toSerialize["NTPServers"] = o.NTPServers
	}
	if !IsNil(o.NetworkDevicePolicy) {
		toSerialize["networkDevicePolicy"] = o.NetworkDevicePolicy
	}
	if !IsNil(o.ServerPolicy) {
		toSerialize["serverPolicy"] = o.ServerPolicy
	}
	if !IsNil(o.ControllerPolicy) {
		toSerialize["controllerPolicy"] = o.ControllerPolicy
	}
	if !IsNil(o.InfrastructurePolicy) {
		toSerialize["infrastructurePolicy"] = o.InfrastructurePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteConfig) UnmarshalJSON(data []byte) (err error) {
	varSiteConfig := _SiteConfig{}

	err = json.Unmarshal(data, &varSiteConfig)

	if err != nil {
		return err
	}

	*o = SiteConfig(varSiteConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "location")
		delete(additionalProperties, "repo")
		delete(additionalProperties, "dnsZoneId")
		delete(additionalProperties, "DNSServers")
		delete(additionalProperties, "NTPServers")
		delete(additionalProperties, "networkDevicePolicy")
		delete(additionalProperties, "serverPolicy")
		delete(additionalProperties, "controllerPolicy")
		delete(additionalProperties, "infrastructurePolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteConfig struct {
	value *SiteConfig
	isSet bool
}

func (v NullableSiteConfig) Get() *SiteConfig {
	return v.value
}

func (v *NullableSiteConfig) Set(val *SiteConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteConfig(val *SiteConfig) *NullableSiteConfig {
	return &NullableSiteConfig{value: val, isSet: true}
}

func (v NullableSiteConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



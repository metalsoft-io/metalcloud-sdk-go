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
	"fmt"
)

// checks if the SiteConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteConfig{}

// SiteConfig struct for SiteConfig
type SiteConfig struct {
	// Location details
	Location Location `json:"location"`
	// Repository details
	Repo Repo `json:"repo"`
	// List of DNS Servers
	DNSServers []string `json:"DNSServers"`
	// List of NTP Servers
	NTPServers []string `json:"NTPServers"`
	// Network device policies
	NetworkDevicePolicy NetworkDevicePolicy `json:"networkDevicePolicy"`
	// Server policies
	ServerPolicy ServerPolicy `json:"serverPolicy"`
	// Controller policies
	ControllerPolicy ControllerPolicy `json:"controllerPolicy"`
	AdditionalProperties map[string]interface{}
}

type _SiteConfig SiteConfig

// NewSiteConfig instantiates a new SiteConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteConfig(location Location, repo Repo, dNSServers []string, nTPServers []string, networkDevicePolicy NetworkDevicePolicy, serverPolicy ServerPolicy, controllerPolicy ControllerPolicy) *SiteConfig {
	this := SiteConfig{}
	this.Location = location
	this.Repo = repo
	this.DNSServers = dNSServers
	this.NTPServers = nTPServers
	this.NetworkDevicePolicy = networkDevicePolicy
	this.ServerPolicy = serverPolicy
	this.ControllerPolicy = controllerPolicy
	return &this
}

// NewSiteConfigWithDefaults instantiates a new SiteConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteConfigWithDefaults() *SiteConfig {
	this := SiteConfig{}
	return &this
}

// GetLocation returns the Location field value
func (o *SiteConfig) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetLocationOk() (*Location, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *SiteConfig) SetLocation(v Location) {
	o.Location = v
}

// GetRepo returns the Repo field value
func (o *SiteConfig) GetRepo() Repo {
	if o == nil {
		var ret Repo
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetRepoOk() (*Repo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *SiteConfig) SetRepo(v Repo) {
	o.Repo = v
}

// GetDNSServers returns the DNSServers field value
func (o *SiteConfig) GetDNSServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DNSServers
}

// GetDNSServersOk returns a tuple with the DNSServers field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetDNSServersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DNSServers, true
}

// SetDNSServers sets field value
func (o *SiteConfig) SetDNSServers(v []string) {
	o.DNSServers = v
}

// GetNTPServers returns the NTPServers field value
func (o *SiteConfig) GetNTPServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NTPServers
}

// GetNTPServersOk returns a tuple with the NTPServers field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetNTPServersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NTPServers, true
}

// SetNTPServers sets field value
func (o *SiteConfig) SetNTPServers(v []string) {
	o.NTPServers = v
}

// GetNetworkDevicePolicy returns the NetworkDevicePolicy field value
func (o *SiteConfig) GetNetworkDevicePolicy() NetworkDevicePolicy {
	if o == nil {
		var ret NetworkDevicePolicy
		return ret
	}

	return o.NetworkDevicePolicy
}

// GetNetworkDevicePolicyOk returns a tuple with the NetworkDevicePolicy field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetNetworkDevicePolicyOk() (*NetworkDevicePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkDevicePolicy, true
}

// SetNetworkDevicePolicy sets field value
func (o *SiteConfig) SetNetworkDevicePolicy(v NetworkDevicePolicy) {
	o.NetworkDevicePolicy = v
}

// GetServerPolicy returns the ServerPolicy field value
func (o *SiteConfig) GetServerPolicy() ServerPolicy {
	if o == nil {
		var ret ServerPolicy
		return ret
	}

	return o.ServerPolicy
}

// GetServerPolicyOk returns a tuple with the ServerPolicy field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetServerPolicyOk() (*ServerPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerPolicy, true
}

// SetServerPolicy sets field value
func (o *SiteConfig) SetServerPolicy(v ServerPolicy) {
	o.ServerPolicy = v
}

// GetControllerPolicy returns the ControllerPolicy field value
func (o *SiteConfig) GetControllerPolicy() ControllerPolicy {
	if o == nil {
		var ret ControllerPolicy
		return ret
	}

	return o.ControllerPolicy
}

// GetControllerPolicyOk returns a tuple with the ControllerPolicy field value
// and a boolean to check if the value has been set.
func (o *SiteConfig) GetControllerPolicyOk() (*ControllerPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControllerPolicy, true
}

// SetControllerPolicy sets field value
func (o *SiteConfig) SetControllerPolicy(v ControllerPolicy) {
	o.ControllerPolicy = v
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
	toSerialize["location"] = o.Location
	toSerialize["repo"] = o.Repo
	toSerialize["DNSServers"] = o.DNSServers
	toSerialize["NTPServers"] = o.NTPServers
	toSerialize["networkDevicePolicy"] = o.NetworkDevicePolicy
	toSerialize["serverPolicy"] = o.ServerPolicy
	toSerialize["controllerPolicy"] = o.ControllerPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
		"repo",
		"DNSServers",
		"NTPServers",
		"networkDevicePolicy",
		"serverPolicy",
		"controllerPolicy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

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
		delete(additionalProperties, "DNSServers")
		delete(additionalProperties, "NTPServers")
		delete(additionalProperties, "networkDevicePolicy")
		delete(additionalProperties, "serverPolicy")
		delete(additionalProperties, "controllerPolicy")
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



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

// checks if the VMInstanceGroupInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMInstanceGroupInterface{}

// VMInstanceGroupInterface struct for VMInstanceGroupInterface
type VMInstanceGroupInterface struct {
	// Name of the VM Instance Group Interface.
	Label string `json:"label"`
	// Interface index
	Index float32 `json:"index"`
	// Network ID
	NetworkId *float32 `json:"networkId,omitempty"`
	// Subdomain of the VM Instance Group Interface.
	Subdomain *string `json:"subdomain,omitempty"`
	// Timestamp of the VM Instance Group Interface update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Interface ID
	Id float32 `json:"id"`
	// Revision of the VM Instance Group Interface Configuration
	Revision float32 `json:"revision"`
	// Service status of the VM Instance Group Interface.
	ServiceStatus string `json:"serviceStatus"`
	// VM Instance Group ID
	GroupId float32 `json:"groupId"`
	// Infrastructure ID
	InfrastructureId float32 `json:"infrastructureId"`
	// Subdomain permanent of the VM Instance Group Interface.
	SubdomainPermanent *string `json:"subdomainPermanent,omitempty"`
	// Id of the DNS subdomain for the VM Instance Group Interface.
	DnsSubdomainId *float32 `json:"dnsSubdomainId,omitempty"`
	// Id of the permanent DNS subdomain for the VM Instance Group Interface.
	DnsSubdomainPermanentId *float32 `json:"dnsSubdomainPermanentId,omitempty"`
	// The current changes to be deployed for the VM Instance Group Interface.
	Config VMInstanceGroupInterfaceConfiguration `json:"config"`
	// Meta information of the VM Instance Group Interface.
	Meta VMInstanceGroupInterfaceMeta `json:"meta"`
	// Timestamp of the VM Instance Group Interface creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// Links to other resources
	Links map[string]interface{} `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _VMInstanceGroupInterface VMInstanceGroupInterface

// NewVMInstanceGroupInterface instantiates a new VMInstanceGroupInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMInstanceGroupInterface(label string, index float32, updatedTimestamp string, id float32, revision float32, serviceStatus string, groupId float32, infrastructureId float32, config VMInstanceGroupInterfaceConfiguration, meta VMInstanceGroupInterfaceMeta, createdTimestamp string, links map[string]interface{}) *VMInstanceGroupInterface {
	this := VMInstanceGroupInterface{}
	this.Label = label
	this.Index = index
	this.UpdatedTimestamp = updatedTimestamp
	this.Id = id
	this.Revision = revision
	this.ServiceStatus = serviceStatus
	this.GroupId = groupId
	this.InfrastructureId = infrastructureId
	this.Config = config
	this.Meta = meta
	this.CreatedTimestamp = createdTimestamp
	this.Links = links
	return &this
}

// NewVMInstanceGroupInterfaceWithDefaults instantiates a new VMInstanceGroupInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMInstanceGroupInterfaceWithDefaults() *VMInstanceGroupInterface {
	this := VMInstanceGroupInterface{}
	return &this
}

// GetLabel returns the Label field value
func (o *VMInstanceGroupInterface) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VMInstanceGroupInterface) SetLabel(v string) {
	o.Label = v
}

// GetIndex returns the Index field value
func (o *VMInstanceGroupInterface) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *VMInstanceGroupInterface) SetIndex(v float32) {
	o.Index = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *VMInstanceGroupInterface) GetNetworkId() float32 {
	if o == nil || IsNil(o.NetworkId) {
		var ret float32
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetNetworkIdOk() (*float32, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *VMInstanceGroupInterface) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given float32 and assigns it to the NetworkId field.
func (o *VMInstanceGroupInterface) SetNetworkId(v float32) {
	o.NetworkId = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *VMInstanceGroupInterface) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *VMInstanceGroupInterface) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *VMInstanceGroupInterface) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *VMInstanceGroupInterface) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *VMInstanceGroupInterface) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetId returns the Id field value
func (o *VMInstanceGroupInterface) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMInstanceGroupInterface) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *VMInstanceGroupInterface) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *VMInstanceGroupInterface) SetRevision(v float32) {
	o.Revision = v
}

// GetServiceStatus returns the ServiceStatus field value
func (o *VMInstanceGroupInterface) GetServiceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceStatus
}

// GetServiceStatusOk returns a tuple with the ServiceStatus field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetServiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceStatus, true
}

// SetServiceStatus sets field value
func (o *VMInstanceGroupInterface) SetServiceStatus(v string) {
	o.ServiceStatus = v
}

// GetGroupId returns the GroupId field value
func (o *VMInstanceGroupInterface) GetGroupId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetGroupIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *VMInstanceGroupInterface) SetGroupId(v float32) {
	o.GroupId = v
}

// GetInfrastructureId returns the InfrastructureId field value
func (o *VMInstanceGroupInterface) GetInfrastructureId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructureId, true
}

// SetInfrastructureId sets field value
func (o *VMInstanceGroupInterface) SetInfrastructureId(v float32) {
	o.InfrastructureId = v
}

// GetSubdomainPermanent returns the SubdomainPermanent field value if set, zero value otherwise.
func (o *VMInstanceGroupInterface) GetSubdomainPermanent() string {
	if o == nil || IsNil(o.SubdomainPermanent) {
		var ret string
		return ret
	}
	return *o.SubdomainPermanent
}

// GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetSubdomainPermanentOk() (*string, bool) {
	if o == nil || IsNil(o.SubdomainPermanent) {
		return nil, false
	}
	return o.SubdomainPermanent, true
}

// HasSubdomainPermanent returns a boolean if a field has been set.
func (o *VMInstanceGroupInterface) HasSubdomainPermanent() bool {
	if o != nil && !IsNil(o.SubdomainPermanent) {
		return true
	}

	return false
}

// SetSubdomainPermanent gets a reference to the given string and assigns it to the SubdomainPermanent field.
func (o *VMInstanceGroupInterface) SetSubdomainPermanent(v string) {
	o.SubdomainPermanent = &v
}

// GetDnsSubdomainId returns the DnsSubdomainId field value if set, zero value otherwise.
func (o *VMInstanceGroupInterface) GetDnsSubdomainId() float32 {
	if o == nil || IsNil(o.DnsSubdomainId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainId
}

// GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetDnsSubdomainIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainId) {
		return nil, false
	}
	return o.DnsSubdomainId, true
}

// HasDnsSubdomainId returns a boolean if a field has been set.
func (o *VMInstanceGroupInterface) HasDnsSubdomainId() bool {
	if o != nil && !IsNil(o.DnsSubdomainId) {
		return true
	}

	return false
}

// SetDnsSubdomainId gets a reference to the given float32 and assigns it to the DnsSubdomainId field.
func (o *VMInstanceGroupInterface) SetDnsSubdomainId(v float32) {
	o.DnsSubdomainId = &v
}

// GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field value if set, zero value otherwise.
func (o *VMInstanceGroupInterface) GetDnsSubdomainPermanentId() float32 {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainPermanentId
}

// GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetDnsSubdomainPermanentIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		return nil, false
	}
	return o.DnsSubdomainPermanentId, true
}

// HasDnsSubdomainPermanentId returns a boolean if a field has been set.
func (o *VMInstanceGroupInterface) HasDnsSubdomainPermanentId() bool {
	if o != nil && !IsNil(o.DnsSubdomainPermanentId) {
		return true
	}

	return false
}

// SetDnsSubdomainPermanentId gets a reference to the given float32 and assigns it to the DnsSubdomainPermanentId field.
func (o *VMInstanceGroupInterface) SetDnsSubdomainPermanentId(v float32) {
	o.DnsSubdomainPermanentId = &v
}

// GetConfig returns the Config field value
func (o *VMInstanceGroupInterface) GetConfig() VMInstanceGroupInterfaceConfiguration {
	if o == nil {
		var ret VMInstanceGroupInterfaceConfiguration
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetConfigOk() (*VMInstanceGroupInterfaceConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *VMInstanceGroupInterface) SetConfig(v VMInstanceGroupInterfaceConfiguration) {
	o.Config = v
}

// GetMeta returns the Meta field value
func (o *VMInstanceGroupInterface) GetMeta() VMInstanceGroupInterfaceMeta {
	if o == nil {
		var ret VMInstanceGroupInterfaceMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetMetaOk() (*VMInstanceGroupInterfaceMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *VMInstanceGroupInterface) SetMeta(v VMInstanceGroupInterfaceMeta) {
	o.Meta = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *VMInstanceGroupInterface) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *VMInstanceGroupInterface) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetLinks returns the Links field value
func (o *VMInstanceGroupInterface) GetLinks() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupInterface) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *VMInstanceGroupInterface) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o VMInstanceGroupInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMInstanceGroupInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["index"] = o.Index
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	toSerialize["serviceStatus"] = o.ServiceStatus
	toSerialize["groupId"] = o.GroupId
	toSerialize["infrastructureId"] = o.InfrastructureId
	if !IsNil(o.SubdomainPermanent) {
		toSerialize["subdomainPermanent"] = o.SubdomainPermanent
	}
	if !IsNil(o.DnsSubdomainId) {
		toSerialize["dnsSubdomainId"] = o.DnsSubdomainId
	}
	if !IsNil(o.DnsSubdomainPermanentId) {
		toSerialize["dnsSubdomainPermanentId"] = o.DnsSubdomainPermanentId
	}
	toSerialize["config"] = o.Config
	toSerialize["meta"] = o.Meta
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VMInstanceGroupInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"index",
		"updatedTimestamp",
		"id",
		"revision",
		"serviceStatus",
		"groupId",
		"infrastructureId",
		"config",
		"meta",
		"createdTimestamp",
		"links",
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

	varVMInstanceGroupInterface := _VMInstanceGroupInterface{}

	err = json.Unmarshal(data, &varVMInstanceGroupInterface)

	if err != nil {
		return err
	}

	*o = VMInstanceGroupInterface(varVMInstanceGroupInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "index")
		delete(additionalProperties, "networkId")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "serviceStatus")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "subdomainPermanent")
		delete(additionalProperties, "dnsSubdomainId")
		delete(additionalProperties, "dnsSubdomainPermanentId")
		delete(additionalProperties, "config")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMInstanceGroupInterface struct {
	value *VMInstanceGroupInterface
	isSet bool
}

func (v NullableVMInstanceGroupInterface) Get() *VMInstanceGroupInterface {
	return v.value
}

func (v *NullableVMInstanceGroupInterface) Set(val *VMInstanceGroupInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableVMInstanceGroupInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableVMInstanceGroupInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMInstanceGroupInterface(val *VMInstanceGroupInterface) *NullableVMInstanceGroupInterface {
	return &NullableVMInstanceGroupInterface{value: val, isSet: true}
}

func (v NullableVMInstanceGroupInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMInstanceGroupInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



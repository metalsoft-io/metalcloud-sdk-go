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

// checks if the ServerInstanceGroupInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceGroupInterface{}

// ServerInstanceGroupInterface struct for ServerInstanceGroupInterface
type ServerInstanceGroupInterface struct {
	// The server instance group interface ID.
	Id *int32 `json:"id,omitempty"`
	// Revision number
	Revision float32 `json:"revision"`
	// The server instance group interface label.
	Label *string `json:"label,omitempty"`
	// Subdomain of the Server Group.
	Subdomain *string `json:"subdomain,omitempty"`
	// Subdomain permanent of the Server Group.
	SubdomainPermanent *string `json:"subdomainPermanent,omitempty"`
	// Id of the DNS subdomain for the Server Group.
	DnsSubdomainId *int32 `json:"dnsSubdomainId,omitempty"`
	// Id of the permanent DNS subdomain for the Server Group.
	DnsSubdomainPermanentId *int32 `json:"dnsSubdomainPermanentId,omitempty"`
	// Timestamp of the server instance group interface creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// Timestamp of the latest update for the server instance group interface.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// GUI settings in JSON format.
	Meta *GenericGUISettings `json:"meta,omitempty"`
	InfrastructureId int32 `json:"infrastructureId"`
	GroupId int32 `json:"groupId"`
	// The index of the interface (0-based) on this server.
	Index int32 `json:"index"`
	// The ID of the network to which this interface is to be attached to.
	NetworkId *int32 `json:"networkId,omitempty"`
	// Current status of the server instance group interface.
	ServiceStatus string `json:"serviceStatus"`
	Config *ServerInstanceGroupInterfaceConfiguration `json:"config,omitempty"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceGroupInterface ServerInstanceGroupInterface

// NewServerInstanceGroupInterface instantiates a new ServerInstanceGroupInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceGroupInterface(revision float32, createdTimestamp string, updatedTimestamp string, infrastructureId int32, groupId int32, index int32, serviceStatus string) *ServerInstanceGroupInterface {
	this := ServerInstanceGroupInterface{}
	this.Revision = revision
	this.CreatedTimestamp = createdTimestamp
	this.UpdatedTimestamp = updatedTimestamp
	this.InfrastructureId = infrastructureId
	this.GroupId = groupId
	this.Index = index
	this.ServiceStatus = serviceStatus
	return &this
}

// NewServerInstanceGroupInterfaceWithDefaults instantiates a new ServerInstanceGroupInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceGroupInterfaceWithDefaults() *ServerInstanceGroupInterface {
	this := ServerInstanceGroupInterface{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ServerInstanceGroupInterface) SetId(v int32) {
	o.Id = &v
}

// GetRevision returns the Revision field value
func (o *ServerInstanceGroupInterface) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *ServerInstanceGroupInterface) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServerInstanceGroupInterface) SetLabel(v string) {
	o.Label = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *ServerInstanceGroupInterface) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetSubdomainPermanent returns the SubdomainPermanent field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetSubdomainPermanent() string {
	if o == nil || IsNil(o.SubdomainPermanent) {
		var ret string
		return ret
	}
	return *o.SubdomainPermanent
}

// GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetSubdomainPermanentOk() (*string, bool) {
	if o == nil || IsNil(o.SubdomainPermanent) {
		return nil, false
	}
	return o.SubdomainPermanent, true
}

// HasSubdomainPermanent returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasSubdomainPermanent() bool {
	if o != nil && !IsNil(o.SubdomainPermanent) {
		return true
	}

	return false
}

// SetSubdomainPermanent gets a reference to the given string and assigns it to the SubdomainPermanent field.
func (o *ServerInstanceGroupInterface) SetSubdomainPermanent(v string) {
	o.SubdomainPermanent = &v
}

// GetDnsSubdomainId returns the DnsSubdomainId field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetDnsSubdomainId() int32 {
	if o == nil || IsNil(o.DnsSubdomainId) {
		var ret int32
		return ret
	}
	return *o.DnsSubdomainId
}

// GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetDnsSubdomainIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsSubdomainId) {
		return nil, false
	}
	return o.DnsSubdomainId, true
}

// HasDnsSubdomainId returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasDnsSubdomainId() bool {
	if o != nil && !IsNil(o.DnsSubdomainId) {
		return true
	}

	return false
}

// SetDnsSubdomainId gets a reference to the given int32 and assigns it to the DnsSubdomainId field.
func (o *ServerInstanceGroupInterface) SetDnsSubdomainId(v int32) {
	o.DnsSubdomainId = &v
}

// GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetDnsSubdomainPermanentId() int32 {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		var ret int32
		return ret
	}
	return *o.DnsSubdomainPermanentId
}

// GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetDnsSubdomainPermanentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		return nil, false
	}
	return o.DnsSubdomainPermanentId, true
}

// HasDnsSubdomainPermanentId returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasDnsSubdomainPermanentId() bool {
	if o != nil && !IsNil(o.DnsSubdomainPermanentId) {
		return true
	}

	return false
}

// SetDnsSubdomainPermanentId gets a reference to the given int32 and assigns it to the DnsSubdomainPermanentId field.
func (o *ServerInstanceGroupInterface) SetDnsSubdomainPermanentId(v int32) {
	o.DnsSubdomainPermanentId = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *ServerInstanceGroupInterface) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *ServerInstanceGroupInterface) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *ServerInstanceGroupInterface) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *ServerInstanceGroupInterface) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetMeta() GenericGUISettings {
	if o == nil || IsNil(o.Meta) {
		var ret GenericGUISettings
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetMetaOk() (*GenericGUISettings, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given GenericGUISettings and assigns it to the Meta field.
func (o *ServerInstanceGroupInterface) SetMeta(v GenericGUISettings) {
	o.Meta = &v
}

// GetInfrastructureId returns the InfrastructureId field value
func (o *ServerInstanceGroupInterface) GetInfrastructureId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetInfrastructureIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructureId, true
}

// SetInfrastructureId sets field value
func (o *ServerInstanceGroupInterface) SetInfrastructureId(v int32) {
	o.InfrastructureId = v
}

// GetGroupId returns the GroupId field value
func (o *ServerInstanceGroupInterface) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ServerInstanceGroupInterface) SetGroupId(v int32) {
	o.GroupId = v
}

// GetIndex returns the Index field value
func (o *ServerInstanceGroupInterface) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ServerInstanceGroupInterface) SetIndex(v int32) {
	o.Index = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetNetworkId() int32 {
	if o == nil || IsNil(o.NetworkId) {
		var ret int32
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetNetworkIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given int32 and assigns it to the NetworkId field.
func (o *ServerInstanceGroupInterface) SetNetworkId(v int32) {
	o.NetworkId = &v
}

// GetServiceStatus returns the ServiceStatus field value
func (o *ServerInstanceGroupInterface) GetServiceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceStatus
}

// GetServiceStatusOk returns a tuple with the ServiceStatus field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetServiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceStatus, true
}

// SetServiceStatus sets field value
func (o *ServerInstanceGroupInterface) SetServiceStatus(v string) {
	o.ServiceStatus = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetConfig() ServerInstanceGroupInterfaceConfiguration {
	if o == nil || IsNil(o.Config) {
		var ret ServerInstanceGroupInterfaceConfiguration
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetConfigOk() (*ServerInstanceGroupInterfaceConfiguration, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ServerInstanceGroupInterfaceConfiguration and assigns it to the Config field.
func (o *ServerInstanceGroupInterface) SetConfig(v ServerInstanceGroupInterfaceConfiguration) {
	o.Config = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServerInstanceGroupInterface) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupInterface) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServerInstanceGroupInterface) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ServerInstanceGroupInterface) SetLinks(v []Link) {
	o.Links = v
}

func (o ServerInstanceGroupInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceGroupInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["revision"] = o.Revision
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.SubdomainPermanent) {
		toSerialize["subdomainPermanent"] = o.SubdomainPermanent
	}
	if !IsNil(o.DnsSubdomainId) {
		toSerialize["dnsSubdomainId"] = o.DnsSubdomainId
	}
	if !IsNil(o.DnsSubdomainPermanentId) {
		toSerialize["dnsSubdomainPermanentId"] = o.DnsSubdomainPermanentId
	}
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["infrastructureId"] = o.InfrastructureId
	toSerialize["groupId"] = o.GroupId
	toSerialize["index"] = o.Index
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	toSerialize["serviceStatus"] = o.ServiceStatus
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceGroupInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"createdTimestamp",
		"updatedTimestamp",
		"infrastructureId",
		"groupId",
		"index",
		"serviceStatus",
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

	varServerInstanceGroupInterface := _ServerInstanceGroupInterface{}

	err = json.Unmarshal(data, &varServerInstanceGroupInterface)

	if err != nil {
		return err
	}

	*o = ServerInstanceGroupInterface(varServerInstanceGroupInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "subdomainPermanent")
		delete(additionalProperties, "dnsSubdomainId")
		delete(additionalProperties, "dnsSubdomainPermanentId")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "index")
		delete(additionalProperties, "networkId")
		delete(additionalProperties, "serviceStatus")
		delete(additionalProperties, "config")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceGroupInterface struct {
	value *ServerInstanceGroupInterface
	isSet bool
}

func (v NullableServerInstanceGroupInterface) Get() *ServerInstanceGroupInterface {
	return v.value
}

func (v *NullableServerInstanceGroupInterface) Set(val *ServerInstanceGroupInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceGroupInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceGroupInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceGroupInterface(val *ServerInstanceGroupInterface) *NullableServerInstanceGroupInterface {
	return &NullableServerInstanceGroupInterface{value: val, isSet: true}
}

func (v NullableServerInstanceGroupInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceGroupInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the FileShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileShare{}

// FileShare struct for FileShare
type FileShare struct {
	// Disk size in GB for File Share
	SizeGB float32 `json:"sizeGB"`
	// Timestamp of the File Share last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Id of the storage pool the File Share is assigned to
	StoragePoolId *float32 `json:"storagePoolId,omitempty"`
	// Label of the File Share.
	Label string `json:"label"`
	// Subdomain of the File Share.
	Subdomain *string `json:"subdomain,omitempty"`
	// Id of the Logical Network for the File Share.
	LogicalNetworkId *float32 `json:"logicalNetworkId,omitempty"`
	// Id of the File Share
	Id float32 `json:"id"`
	// Revision of the File Share
	Revision float32 `json:"revision"`
	// Infrastructure id of the File Share
	InfrastructureId float32 `json:"infrastructureId"`
	// Timestamp of the File Share creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// Service status of the File Share
	ServiceStatus string `json:"serviceStatus"`
	// Subdomain permanent of the File Share.
	SubdomainPermanent *string `json:"subdomainPermanent,omitempty"`
	// Id of the DNS subdomain for the File Share.
	DnsSubdomainId *float32 `json:"dnsSubdomainId,omitempty"`
	// Id of the VLAN for the File Share.
	NetworkVlanId *float32 `json:"networkVlanId,omitempty"`
	// Endpoint of the File Share.
	Endpoint *string `json:"endpoint,omitempty"`
	// The current changes to be deployed for the File Share.
	Config FileShareConfiguration `json:"config"`
	// Meta information of the File Share.
	Meta FileShareMeta `json:"meta"`
	AdditionalProperties map[string]interface{}
}

type _FileShare FileShare

// NewFileShare instantiates a new FileShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileShare(sizeGB float32, updatedTimestamp string, label string, id float32, revision float32, infrastructureId float32, createdTimestamp string, serviceStatus string, config FileShareConfiguration, meta FileShareMeta) *FileShare {
	this := FileShare{}
	this.SizeGB = sizeGB
	this.UpdatedTimestamp = updatedTimestamp
	this.Label = label
	this.Id = id
	this.Revision = revision
	this.InfrastructureId = infrastructureId
	this.CreatedTimestamp = createdTimestamp
	this.ServiceStatus = serviceStatus
	this.Config = config
	this.Meta = meta
	return &this
}

// NewFileShareWithDefaults instantiates a new FileShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileShareWithDefaults() *FileShare {
	this := FileShare{}
	return &this
}

// GetSizeGB returns the SizeGB field value
func (o *FileShare) GetSizeGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SizeGB
}

// GetSizeGBOk returns a tuple with the SizeGB field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetSizeGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeGB, true
}

// SetSizeGB sets field value
func (o *FileShare) SetSizeGB(v float32) {
	o.SizeGB = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *FileShare) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *FileShare) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetStoragePoolId returns the StoragePoolId field value if set, zero value otherwise.
func (o *FileShare) GetStoragePoolId() float32 {
	if o == nil || IsNil(o.StoragePoolId) {
		var ret float32
		return ret
	}
	return *o.StoragePoolId
}

// GetStoragePoolIdOk returns a tuple with the StoragePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetStoragePoolIdOk() (*float32, bool) {
	if o == nil || IsNil(o.StoragePoolId) {
		return nil, false
	}
	return o.StoragePoolId, true
}

// HasStoragePoolId returns a boolean if a field has been set.
func (o *FileShare) HasStoragePoolId() bool {
	if o != nil && !IsNil(o.StoragePoolId) {
		return true
	}

	return false
}

// SetStoragePoolId gets a reference to the given float32 and assigns it to the StoragePoolId field.
func (o *FileShare) SetStoragePoolId(v float32) {
	o.StoragePoolId = &v
}

// GetLabel returns the Label field value
func (o *FileShare) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FileShare) SetLabel(v string) {
	o.Label = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *FileShare) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *FileShare) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *FileShare) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetLogicalNetworkId returns the LogicalNetworkId field value if set, zero value otherwise.
func (o *FileShare) GetLogicalNetworkId() float32 {
	if o == nil || IsNil(o.LogicalNetworkId) {
		var ret float32
		return ret
	}
	return *o.LogicalNetworkId
}

// GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetLogicalNetworkIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LogicalNetworkId) {
		return nil, false
	}
	return o.LogicalNetworkId, true
}

// HasLogicalNetworkId returns a boolean if a field has been set.
func (o *FileShare) HasLogicalNetworkId() bool {
	if o != nil && !IsNil(o.LogicalNetworkId) {
		return true
	}

	return false
}

// SetLogicalNetworkId gets a reference to the given float32 and assigns it to the LogicalNetworkId field.
func (o *FileShare) SetLogicalNetworkId(v float32) {
	o.LogicalNetworkId = &v
}

// GetId returns the Id field value
func (o *FileShare) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FileShare) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *FileShare) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *FileShare) SetRevision(v float32) {
	o.Revision = v
}

// GetInfrastructureId returns the InfrastructureId field value
func (o *FileShare) GetInfrastructureId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructureId, true
}

// SetInfrastructureId sets field value
func (o *FileShare) SetInfrastructureId(v float32) {
	o.InfrastructureId = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *FileShare) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *FileShare) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetServiceStatus returns the ServiceStatus field value
func (o *FileShare) GetServiceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceStatus
}

// GetServiceStatusOk returns a tuple with the ServiceStatus field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetServiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceStatus, true
}

// SetServiceStatus sets field value
func (o *FileShare) SetServiceStatus(v string) {
	o.ServiceStatus = v
}

// GetSubdomainPermanent returns the SubdomainPermanent field value if set, zero value otherwise.
func (o *FileShare) GetSubdomainPermanent() string {
	if o == nil || IsNil(o.SubdomainPermanent) {
		var ret string
		return ret
	}
	return *o.SubdomainPermanent
}

// GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetSubdomainPermanentOk() (*string, bool) {
	if o == nil || IsNil(o.SubdomainPermanent) {
		return nil, false
	}
	return o.SubdomainPermanent, true
}

// HasSubdomainPermanent returns a boolean if a field has been set.
func (o *FileShare) HasSubdomainPermanent() bool {
	if o != nil && !IsNil(o.SubdomainPermanent) {
		return true
	}

	return false
}

// SetSubdomainPermanent gets a reference to the given string and assigns it to the SubdomainPermanent field.
func (o *FileShare) SetSubdomainPermanent(v string) {
	o.SubdomainPermanent = &v
}

// GetDnsSubdomainId returns the DnsSubdomainId field value if set, zero value otherwise.
func (o *FileShare) GetDnsSubdomainId() float32 {
	if o == nil || IsNil(o.DnsSubdomainId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainId
}

// GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetDnsSubdomainIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainId) {
		return nil, false
	}
	return o.DnsSubdomainId, true
}

// HasDnsSubdomainId returns a boolean if a field has been set.
func (o *FileShare) HasDnsSubdomainId() bool {
	if o != nil && !IsNil(o.DnsSubdomainId) {
		return true
	}

	return false
}

// SetDnsSubdomainId gets a reference to the given float32 and assigns it to the DnsSubdomainId field.
func (o *FileShare) SetDnsSubdomainId(v float32) {
	o.DnsSubdomainId = &v
}

// GetNetworkVlanId returns the NetworkVlanId field value if set, zero value otherwise.
func (o *FileShare) GetNetworkVlanId() float32 {
	if o == nil || IsNil(o.NetworkVlanId) {
		var ret float32
		return ret
	}
	return *o.NetworkVlanId
}

// GetNetworkVlanIdOk returns a tuple with the NetworkVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetNetworkVlanIdOk() (*float32, bool) {
	if o == nil || IsNil(o.NetworkVlanId) {
		return nil, false
	}
	return o.NetworkVlanId, true
}

// HasNetworkVlanId returns a boolean if a field has been set.
func (o *FileShare) HasNetworkVlanId() bool {
	if o != nil && !IsNil(o.NetworkVlanId) {
		return true
	}

	return false
}

// SetNetworkVlanId gets a reference to the given float32 and assigns it to the NetworkVlanId field.
func (o *FileShare) SetNetworkVlanId(v float32) {
	o.NetworkVlanId = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *FileShare) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShare) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *FileShare) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *FileShare) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetConfig returns the Config field value
func (o *FileShare) GetConfig() FileShareConfiguration {
	if o == nil {
		var ret FileShareConfiguration
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetConfigOk() (*FileShareConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *FileShare) SetConfig(v FileShareConfiguration) {
	o.Config = v
}

// GetMeta returns the Meta field value
func (o *FileShare) GetMeta() FileShareMeta {
	if o == nil {
		var ret FileShareMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FileShare) GetMetaOk() (*FileShareMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FileShare) SetMeta(v FileShareMeta) {
	o.Meta = v
}

func (o FileShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sizeGB"] = o.SizeGB
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	if !IsNil(o.StoragePoolId) {
		toSerialize["storagePoolId"] = o.StoragePoolId
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.LogicalNetworkId) {
		toSerialize["logicalNetworkId"] = o.LogicalNetworkId
	}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	toSerialize["infrastructureId"] = o.InfrastructureId
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	toSerialize["serviceStatus"] = o.ServiceStatus
	if !IsNil(o.SubdomainPermanent) {
		toSerialize["subdomainPermanent"] = o.SubdomainPermanent
	}
	if !IsNil(o.DnsSubdomainId) {
		toSerialize["dnsSubdomainId"] = o.DnsSubdomainId
	}
	if !IsNil(o.NetworkVlanId) {
		toSerialize["networkVlanId"] = o.NetworkVlanId
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	toSerialize["config"] = o.Config
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileShare) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sizeGB",
		"updatedTimestamp",
		"label",
		"id",
		"revision",
		"infrastructureId",
		"createdTimestamp",
		"serviceStatus",
		"config",
		"meta",
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

	varFileShare := _FileShare{}

	err = json.Unmarshal(data, &varFileShare)

	if err != nil {
		return err
	}

	*o = FileShare(varFileShare)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sizeGB")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "storagePoolId")
		delete(additionalProperties, "label")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "logicalNetworkId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "serviceStatus")
		delete(additionalProperties, "subdomainPermanent")
		delete(additionalProperties, "dnsSubdomainId")
		delete(additionalProperties, "networkVlanId")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "config")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileShare struct {
	value *FileShare
	isSet bool
}

func (v NullableFileShare) Get() *FileShare {
	return v.value
}

func (v *NullableFileShare) Set(val *FileShare) {
	v.value = val
	v.isSet = true
}

func (v NullableFileShare) IsSet() bool {
	return v.isSet
}

func (v *NullableFileShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileShare(val *FileShare) *NullableFileShare {
	return &NullableFileShare{value: val, isSet: true}
}

func (v NullableFileShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



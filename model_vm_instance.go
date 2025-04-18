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

// checks if the VMInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMInstance{}

// VMInstance struct for VMInstance
type VMInstance struct {
	// Name of the VM Instance.
	Label string `json:"label"`
	// Id of the VM.
	VmId *float32 `json:"vmId,omitempty"`
	// Id of the VM Type.
	TypeId float32 `json:"typeId"`
	// Subdomain of the VM Instance.
	Subdomain *string `json:"subdomain,omitempty"`
	// Id of the VM Pool.
	VmPoolId *float32 `json:"vmPoolId,omitempty"`
	// Disk size in GB of the VM Instance.
	DiskSizeGB float32 `json:"diskSizeGB"`
	// RAM size in GB of the VM Instance.
	RamGB float32 `json:"ramGB"`
	// Number of CPU cores for the VM Instance.
	CpuCores float32 `json:"cpuCores"`
	// Id of the template used by the VM Instance.
	OsTemplateId *float32 `json:"osTemplateId,omitempty"`
	// Custom variables for the VM Instance.
	CustomVariables map[string]interface{} `json:"customVariables,omitempty"`
	// Timestamp of the VM Instance last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// VM Instance ID
	Id float32 `json:"id"`
	// Revision of the VM Instance
	Revision float32 `json:"revision"`
	// Id of the VM Instance Group.
	GroupId float32 `json:"groupId"`
	// Id of the Infrastructure.
	InfrastructureId float32 `json:"infrastructureId"`
	// Infrastructure information
	Infrastructure ParentInfrastructureDto `json:"infrastructure"`
	// Service status of the VM Instance.
	ServiceStatus string `json:"serviceStatus"`
	// Subdomain permanent of the VM Instance.
	SubdomainPermanent *string `json:"subdomainPermanent,omitempty"`
	// Id of the DNS subdomain for the VM Instance.
	DnsSubdomainId *float32 `json:"dnsSubdomainId,omitempty"`
	// Id of the permanent DNS subdomain for the VM Instance.
	DnsSubdomainPermanentId *float32 `json:"dnsSubdomainPermanentId,omitempty"`
	// Unique index of the VM Instance.
	UniqueIndex *float32 `json:"uniqueIndex,omitempty"`
	// Timestamp of the VM Instance creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// The current changes to be deployed for the VM Instance.
	Config VMInstanceConfiguration `json:"config"`
	// Meta information of the VM Instance.
	Meta VMInstanceMeta `json:"meta"`
	// Links to other resources
	Links map[string]interface{} `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _VMInstance VMInstance

// NewVMInstance instantiates a new VMInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMInstance(label string, typeId float32, diskSizeGB float32, ramGB float32, cpuCores float32, updatedTimestamp string, id float32, revision float32, groupId float32, infrastructureId float32, infrastructure ParentInfrastructureDto, serviceStatus string, createdTimestamp string, config VMInstanceConfiguration, meta VMInstanceMeta, links map[string]interface{}) *VMInstance {
	this := VMInstance{}
	this.Label = label
	this.TypeId = typeId
	this.DiskSizeGB = diskSizeGB
	this.RamGB = ramGB
	this.CpuCores = cpuCores
	this.UpdatedTimestamp = updatedTimestamp
	this.Id = id
	this.Revision = revision
	this.GroupId = groupId
	this.InfrastructureId = infrastructureId
	this.Infrastructure = infrastructure
	this.ServiceStatus = serviceStatus
	this.CreatedTimestamp = createdTimestamp
	this.Config = config
	this.Meta = meta
	this.Links = links
	return &this
}

// NewVMInstanceWithDefaults instantiates a new VMInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMInstanceWithDefaults() *VMInstance {
	this := VMInstance{}
	return &this
}

// GetLabel returns the Label field value
func (o *VMInstance) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VMInstance) SetLabel(v string) {
	o.Label = v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *VMInstance) GetVmId() float32 {
	if o == nil || IsNil(o.VmId) {
		var ret float32
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetVmIdOk() (*float32, bool) {
	if o == nil || IsNil(o.VmId) {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *VMInstance) HasVmId() bool {
	if o != nil && !IsNil(o.VmId) {
		return true
	}

	return false
}

// SetVmId gets a reference to the given float32 and assigns it to the VmId field.
func (o *VMInstance) SetVmId(v float32) {
	o.VmId = &v
}

// GetTypeId returns the TypeId field value
func (o *VMInstance) GetTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *VMInstance) SetTypeId(v float32) {
	o.TypeId = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *VMInstance) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *VMInstance) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *VMInstance) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetVmPoolId returns the VmPoolId field value if set, zero value otherwise.
func (o *VMInstance) GetVmPoolId() float32 {
	if o == nil || IsNil(o.VmPoolId) {
		var ret float32
		return ret
	}
	return *o.VmPoolId
}

// GetVmPoolIdOk returns a tuple with the VmPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetVmPoolIdOk() (*float32, bool) {
	if o == nil || IsNil(o.VmPoolId) {
		return nil, false
	}
	return o.VmPoolId, true
}

// HasVmPoolId returns a boolean if a field has been set.
func (o *VMInstance) HasVmPoolId() bool {
	if o != nil && !IsNil(o.VmPoolId) {
		return true
	}

	return false
}

// SetVmPoolId gets a reference to the given float32 and assigns it to the VmPoolId field.
func (o *VMInstance) SetVmPoolId(v float32) {
	o.VmPoolId = &v
}

// GetDiskSizeGB returns the DiskSizeGB field value
func (o *VMInstance) GetDiskSizeGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetDiskSizeGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskSizeGB, true
}

// SetDiskSizeGB sets field value
func (o *VMInstance) SetDiskSizeGB(v float32) {
	o.DiskSizeGB = v
}

// GetRamGB returns the RamGB field value
func (o *VMInstance) GetRamGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamGB
}

// GetRamGBOk returns a tuple with the RamGB field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetRamGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamGB, true
}

// SetRamGB sets field value
func (o *VMInstance) SetRamGB(v float32) {
	o.RamGB = v
}

// GetCpuCores returns the CpuCores field value
func (o *VMInstance) GetCpuCores() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetCpuCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuCores, true
}

// SetCpuCores sets field value
func (o *VMInstance) SetCpuCores(v float32) {
	o.CpuCores = v
}

// GetOsTemplateId returns the OsTemplateId field value if set, zero value otherwise.
func (o *VMInstance) GetOsTemplateId() float32 {
	if o == nil || IsNil(o.OsTemplateId) {
		var ret float32
		return ret
	}
	return *o.OsTemplateId
}

// GetOsTemplateIdOk returns a tuple with the OsTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetOsTemplateIdOk() (*float32, bool) {
	if o == nil || IsNil(o.OsTemplateId) {
		return nil, false
	}
	return o.OsTemplateId, true
}

// HasOsTemplateId returns a boolean if a field has been set.
func (o *VMInstance) HasOsTemplateId() bool {
	if o != nil && !IsNil(o.OsTemplateId) {
		return true
	}

	return false
}

// SetOsTemplateId gets a reference to the given float32 and assigns it to the OsTemplateId field.
func (o *VMInstance) SetOsTemplateId(v float32) {
	o.OsTemplateId = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *VMInstance) GetCustomVariables() map[string]interface{} {
	if o == nil || IsNil(o.CustomVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetCustomVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return map[string]interface{}{}, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *VMInstance) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given map[string]interface{} and assigns it to the CustomVariables field.
func (o *VMInstance) SetCustomVariables(v map[string]interface{}) {
	o.CustomVariables = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *VMInstance) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *VMInstance) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetId returns the Id field value
func (o *VMInstance) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMInstance) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *VMInstance) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *VMInstance) SetRevision(v float32) {
	o.Revision = v
}

// GetGroupId returns the GroupId field value
func (o *VMInstance) GetGroupId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetGroupIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *VMInstance) SetGroupId(v float32) {
	o.GroupId = v
}

// GetInfrastructureId returns the InfrastructureId field value
func (o *VMInstance) GetInfrastructureId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructureId, true
}

// SetInfrastructureId sets field value
func (o *VMInstance) SetInfrastructureId(v float32) {
	o.InfrastructureId = v
}

// GetInfrastructure returns the Infrastructure field value
func (o *VMInstance) GetInfrastructure() ParentInfrastructureDto {
	if o == nil {
		var ret ParentInfrastructureDto
		return ret
	}

	return o.Infrastructure
}

// GetInfrastructureOk returns a tuple with the Infrastructure field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetInfrastructureOk() (*ParentInfrastructureDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Infrastructure, true
}

// SetInfrastructure sets field value
func (o *VMInstance) SetInfrastructure(v ParentInfrastructureDto) {
	o.Infrastructure = v
}

// GetServiceStatus returns the ServiceStatus field value
func (o *VMInstance) GetServiceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceStatus
}

// GetServiceStatusOk returns a tuple with the ServiceStatus field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetServiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceStatus, true
}

// SetServiceStatus sets field value
func (o *VMInstance) SetServiceStatus(v string) {
	o.ServiceStatus = v
}

// GetSubdomainPermanent returns the SubdomainPermanent field value if set, zero value otherwise.
func (o *VMInstance) GetSubdomainPermanent() string {
	if o == nil || IsNil(o.SubdomainPermanent) {
		var ret string
		return ret
	}
	return *o.SubdomainPermanent
}

// GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetSubdomainPermanentOk() (*string, bool) {
	if o == nil || IsNil(o.SubdomainPermanent) {
		return nil, false
	}
	return o.SubdomainPermanent, true
}

// HasSubdomainPermanent returns a boolean if a field has been set.
func (o *VMInstance) HasSubdomainPermanent() bool {
	if o != nil && !IsNil(o.SubdomainPermanent) {
		return true
	}

	return false
}

// SetSubdomainPermanent gets a reference to the given string and assigns it to the SubdomainPermanent field.
func (o *VMInstance) SetSubdomainPermanent(v string) {
	o.SubdomainPermanent = &v
}

// GetDnsSubdomainId returns the DnsSubdomainId field value if set, zero value otherwise.
func (o *VMInstance) GetDnsSubdomainId() float32 {
	if o == nil || IsNil(o.DnsSubdomainId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainId
}

// GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetDnsSubdomainIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainId) {
		return nil, false
	}
	return o.DnsSubdomainId, true
}

// HasDnsSubdomainId returns a boolean if a field has been set.
func (o *VMInstance) HasDnsSubdomainId() bool {
	if o != nil && !IsNil(o.DnsSubdomainId) {
		return true
	}

	return false
}

// SetDnsSubdomainId gets a reference to the given float32 and assigns it to the DnsSubdomainId field.
func (o *VMInstance) SetDnsSubdomainId(v float32) {
	o.DnsSubdomainId = &v
}

// GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field value if set, zero value otherwise.
func (o *VMInstance) GetDnsSubdomainPermanentId() float32 {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainPermanentId
}

// GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetDnsSubdomainPermanentIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainPermanentId) {
		return nil, false
	}
	return o.DnsSubdomainPermanentId, true
}

// HasDnsSubdomainPermanentId returns a boolean if a field has been set.
func (o *VMInstance) HasDnsSubdomainPermanentId() bool {
	if o != nil && !IsNil(o.DnsSubdomainPermanentId) {
		return true
	}

	return false
}

// SetDnsSubdomainPermanentId gets a reference to the given float32 and assigns it to the DnsSubdomainPermanentId field.
func (o *VMInstance) SetDnsSubdomainPermanentId(v float32) {
	o.DnsSubdomainPermanentId = &v
}

// GetUniqueIndex returns the UniqueIndex field value if set, zero value otherwise.
func (o *VMInstance) GetUniqueIndex() float32 {
	if o == nil || IsNil(o.UniqueIndex) {
		var ret float32
		return ret
	}
	return *o.UniqueIndex
}

// GetUniqueIndexOk returns a tuple with the UniqueIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstance) GetUniqueIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.UniqueIndex) {
		return nil, false
	}
	return o.UniqueIndex, true
}

// HasUniqueIndex returns a boolean if a field has been set.
func (o *VMInstance) HasUniqueIndex() bool {
	if o != nil && !IsNil(o.UniqueIndex) {
		return true
	}

	return false
}

// SetUniqueIndex gets a reference to the given float32 and assigns it to the UniqueIndex field.
func (o *VMInstance) SetUniqueIndex(v float32) {
	o.UniqueIndex = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *VMInstance) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *VMInstance) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetConfig returns the Config field value
func (o *VMInstance) GetConfig() VMInstanceConfiguration {
	if o == nil {
		var ret VMInstanceConfiguration
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetConfigOk() (*VMInstanceConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *VMInstance) SetConfig(v VMInstanceConfiguration) {
	o.Config = v
}

// GetMeta returns the Meta field value
func (o *VMInstance) GetMeta() VMInstanceMeta {
	if o == nil {
		var ret VMInstanceMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetMetaOk() (*VMInstanceMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *VMInstance) SetMeta(v VMInstanceMeta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *VMInstance) GetLinks() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *VMInstance) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *VMInstance) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o VMInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !IsNil(o.VmId) {
		toSerialize["vmId"] = o.VmId
	}
	toSerialize["typeId"] = o.TypeId
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.VmPoolId) {
		toSerialize["vmPoolId"] = o.VmPoolId
	}
	toSerialize["diskSizeGB"] = o.DiskSizeGB
	toSerialize["ramGB"] = o.RamGB
	toSerialize["cpuCores"] = o.CpuCores
	if !IsNil(o.OsTemplateId) {
		toSerialize["osTemplateId"] = o.OsTemplateId
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	toSerialize["groupId"] = o.GroupId
	toSerialize["infrastructureId"] = o.InfrastructureId
	toSerialize["infrastructure"] = o.Infrastructure
	toSerialize["serviceStatus"] = o.ServiceStatus
	if !IsNil(o.SubdomainPermanent) {
		toSerialize["subdomainPermanent"] = o.SubdomainPermanent
	}
	if !IsNil(o.DnsSubdomainId) {
		toSerialize["dnsSubdomainId"] = o.DnsSubdomainId
	}
	if !IsNil(o.DnsSubdomainPermanentId) {
		toSerialize["dnsSubdomainPermanentId"] = o.DnsSubdomainPermanentId
	}
	if !IsNil(o.UniqueIndex) {
		toSerialize["uniqueIndex"] = o.UniqueIndex
	}
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	toSerialize["config"] = o.Config
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VMInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"typeId",
		"diskSizeGB",
		"ramGB",
		"cpuCores",
		"updatedTimestamp",
		"id",
		"revision",
		"groupId",
		"infrastructureId",
		"infrastructure",
		"serviceStatus",
		"createdTimestamp",
		"config",
		"meta",
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

	varVMInstance := _VMInstance{}

	err = json.Unmarshal(data, &varVMInstance)

	if err != nil {
		return err
	}

	*o = VMInstance(varVMInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "vmId")
		delete(additionalProperties, "typeId")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "vmPoolId")
		delete(additionalProperties, "diskSizeGB")
		delete(additionalProperties, "ramGB")
		delete(additionalProperties, "cpuCores")
		delete(additionalProperties, "osTemplateId")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "infrastructure")
		delete(additionalProperties, "serviceStatus")
		delete(additionalProperties, "subdomainPermanent")
		delete(additionalProperties, "dnsSubdomainId")
		delete(additionalProperties, "dnsSubdomainPermanentId")
		delete(additionalProperties, "uniqueIndex")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "config")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMInstance struct {
	value *VMInstance
	isSet bool
}

func (v NullableVMInstance) Get() *VMInstance {
	return v.value
}

func (v *NullableVMInstance) Set(val *VMInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableVMInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableVMInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMInstance(val *VMInstance) *NullableVMInstance {
	return &NullableVMInstance{value: val, isSet: true}
}

func (v NullableVMInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



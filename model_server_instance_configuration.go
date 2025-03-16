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

// checks if the ServerInstanceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceConfiguration{}

// ServerInstanceConfiguration struct for ServerInstanceConfiguration
type ServerInstanceConfiguration struct {
	// Revision number
	Revision float32 `json:"revision"`
	// The server instance label.
	Label string `json:"label"`
	// Timestamp of the latest update for the Server Instance.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	GroupId int32 `json:"groupId"`
	DriveIdBootable *int32 `json:"driveIdBootable,omitempty"`
	// Subdomain of the Server Group.
	Subdomain *string `json:"subdomain,omitempty"`
	// The server type ID.
	ServerTypeId *int32 `json:"serverTypeId,omitempty"`
	// The ID of the server assigned to the instance.
	ServerId *int32 `json:"serverId,omitempty"`
	// The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal. 
	TemplateId *int32 `json:"templateId,omitempty"`
	InstanceWanMlagId *int32 `json:"instanceWanMlagId,omitempty"`
	CustomVariables map[string]interface{} `json:"customVariables,omitempty"`
	// RAID profile for the Instance Interface.
	RaidProfile *ServerInstanceStorageProfile `json:"raidProfile,omitempty"`
	// iSCSI Initiator IQN for the Instance Interface.
	IscsiInitiatorIqn *string `json:"iscsiInitiatorIqn,omitempty"`
	// iSCSI Initiator Username for the Instance Interface.
	IscsiInitiatorUsername *string `json:"iscsiInitiatorUsername,omitempty"`
	// iSCSI Initiator Password for the Instance Interface.
	IscsiInitiatorPasswordEncrypted *string `json:"iscsiInitiatorPasswordEncrypted,omitempty"`
	// Control panel url for the Instance Interface.
	ControlPanelUrl *string `json:"controlPanelUrl,omitempty"`
	// Number of empty edits
	EmptyEdit *int32 `json:"emptyEdit,omitempty"`
	// Server Instance deploy type
	DeployType string `json:"deployType"`
	// Server Instance deploy status
	DeployStatus string `json:"deployStatus"`
	// Id of the DNS subdomain for the Server Instance.
	DnsSubdomainChangeId *int32 `json:"dnsSubdomainChangeId,omitempty"`
	// Id of the deployment for the Server Instance.
	InfrastructureDeployId *int32 `json:"infrastructureDeployId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceConfiguration ServerInstanceConfiguration

// NewServerInstanceConfiguration instantiates a new ServerInstanceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceConfiguration(revision float32, label string, updatedTimestamp string, groupId int32, deployType string, deployStatus string) *ServerInstanceConfiguration {
	this := ServerInstanceConfiguration{}
	this.Revision = revision
	this.Label = label
	this.UpdatedTimestamp = updatedTimestamp
	this.GroupId = groupId
	this.DeployType = deployType
	this.DeployStatus = deployStatus
	return &this
}

// NewServerInstanceConfigurationWithDefaults instantiates a new ServerInstanceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceConfigurationWithDefaults() *ServerInstanceConfiguration {
	this := ServerInstanceConfiguration{}
	return &this
}

// GetRevision returns the Revision field value
func (o *ServerInstanceConfiguration) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *ServerInstanceConfiguration) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value
func (o *ServerInstanceConfiguration) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ServerInstanceConfiguration) SetLabel(v string) {
	o.Label = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *ServerInstanceConfiguration) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *ServerInstanceConfiguration) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetGroupId returns the GroupId field value
func (o *ServerInstanceConfiguration) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ServerInstanceConfiguration) SetGroupId(v int32) {
	o.GroupId = v
}

// GetDriveIdBootable returns the DriveIdBootable field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetDriveIdBootable() int32 {
	if o == nil || IsNil(o.DriveIdBootable) {
		var ret int32
		return ret
	}
	return *o.DriveIdBootable
}

// GetDriveIdBootableOk returns a tuple with the DriveIdBootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetDriveIdBootableOk() (*int32, bool) {
	if o == nil || IsNil(o.DriveIdBootable) {
		return nil, false
	}
	return o.DriveIdBootable, true
}

// HasDriveIdBootable returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasDriveIdBootable() bool {
	if o != nil && !IsNil(o.DriveIdBootable) {
		return true
	}

	return false
}

// SetDriveIdBootable gets a reference to the given int32 and assigns it to the DriveIdBootable field.
func (o *ServerInstanceConfiguration) SetDriveIdBootable(v int32) {
	o.DriveIdBootable = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *ServerInstanceConfiguration) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetServerTypeId returns the ServerTypeId field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetServerTypeId() int32 {
	if o == nil || IsNil(o.ServerTypeId) {
		var ret int32
		return ret
	}
	return *o.ServerTypeId
}

// GetServerTypeIdOk returns a tuple with the ServerTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetServerTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerTypeId) {
		return nil, false
	}
	return o.ServerTypeId, true
}

// HasServerTypeId returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasServerTypeId() bool {
	if o != nil && !IsNil(o.ServerTypeId) {
		return true
	}

	return false
}

// SetServerTypeId gets a reference to the given int32 and assigns it to the ServerTypeId field.
func (o *ServerInstanceConfiguration) SetServerTypeId(v int32) {
	o.ServerTypeId = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetServerId() int32 {
	if o == nil || IsNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetServerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *ServerInstanceConfiguration) SetServerId(v int32) {
	o.ServerId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ServerInstanceConfiguration) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetInstanceWanMlagId returns the InstanceWanMlagId field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetInstanceWanMlagId() int32 {
	if o == nil || IsNil(o.InstanceWanMlagId) {
		var ret int32
		return ret
	}
	return *o.InstanceWanMlagId
}

// GetInstanceWanMlagIdOk returns a tuple with the InstanceWanMlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetInstanceWanMlagIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceWanMlagId) {
		return nil, false
	}
	return o.InstanceWanMlagId, true
}

// HasInstanceWanMlagId returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasInstanceWanMlagId() bool {
	if o != nil && !IsNil(o.InstanceWanMlagId) {
		return true
	}

	return false
}

// SetInstanceWanMlagId gets a reference to the given int32 and assigns it to the InstanceWanMlagId field.
func (o *ServerInstanceConfiguration) SetInstanceWanMlagId(v int32) {
	o.InstanceWanMlagId = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetCustomVariables() map[string]interface{} {
	if o == nil || IsNil(o.CustomVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetCustomVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return map[string]interface{}{}, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given map[string]interface{} and assigns it to the CustomVariables field.
func (o *ServerInstanceConfiguration) SetCustomVariables(v map[string]interface{}) {
	o.CustomVariables = v
}

// GetRaidProfile returns the RaidProfile field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetRaidProfile() ServerInstanceStorageProfile {
	if o == nil || IsNil(o.RaidProfile) {
		var ret ServerInstanceStorageProfile
		return ret
	}
	return *o.RaidProfile
}

// GetRaidProfileOk returns a tuple with the RaidProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetRaidProfileOk() (*ServerInstanceStorageProfile, bool) {
	if o == nil || IsNil(o.RaidProfile) {
		return nil, false
	}
	return o.RaidProfile, true
}

// HasRaidProfile returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasRaidProfile() bool {
	if o != nil && !IsNil(o.RaidProfile) {
		return true
	}

	return false
}

// SetRaidProfile gets a reference to the given ServerInstanceStorageProfile and assigns it to the RaidProfile field.
func (o *ServerInstanceConfiguration) SetRaidProfile(v ServerInstanceStorageProfile) {
	o.RaidProfile = &v
}

// GetIscsiInitiatorIqn returns the IscsiInitiatorIqn field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetIscsiInitiatorIqn() string {
	if o == nil || IsNil(o.IscsiInitiatorIqn) {
		var ret string
		return ret
	}
	return *o.IscsiInitiatorIqn
}

// GetIscsiInitiatorIqnOk returns a tuple with the IscsiInitiatorIqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetIscsiInitiatorIqnOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiInitiatorIqn) {
		return nil, false
	}
	return o.IscsiInitiatorIqn, true
}

// HasIscsiInitiatorIqn returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasIscsiInitiatorIqn() bool {
	if o != nil && !IsNil(o.IscsiInitiatorIqn) {
		return true
	}

	return false
}

// SetIscsiInitiatorIqn gets a reference to the given string and assigns it to the IscsiInitiatorIqn field.
func (o *ServerInstanceConfiguration) SetIscsiInitiatorIqn(v string) {
	o.IscsiInitiatorIqn = &v
}

// GetIscsiInitiatorUsername returns the IscsiInitiatorUsername field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetIscsiInitiatorUsername() string {
	if o == nil || IsNil(o.IscsiInitiatorUsername) {
		var ret string
		return ret
	}
	return *o.IscsiInitiatorUsername
}

// GetIscsiInitiatorUsernameOk returns a tuple with the IscsiInitiatorUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetIscsiInitiatorUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiInitiatorUsername) {
		return nil, false
	}
	return o.IscsiInitiatorUsername, true
}

// HasIscsiInitiatorUsername returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasIscsiInitiatorUsername() bool {
	if o != nil && !IsNil(o.IscsiInitiatorUsername) {
		return true
	}

	return false
}

// SetIscsiInitiatorUsername gets a reference to the given string and assigns it to the IscsiInitiatorUsername field.
func (o *ServerInstanceConfiguration) SetIscsiInitiatorUsername(v string) {
	o.IscsiInitiatorUsername = &v
}

// GetIscsiInitiatorPasswordEncrypted returns the IscsiInitiatorPasswordEncrypted field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetIscsiInitiatorPasswordEncrypted() string {
	if o == nil || IsNil(o.IscsiInitiatorPasswordEncrypted) {
		var ret string
		return ret
	}
	return *o.IscsiInitiatorPasswordEncrypted
}

// GetIscsiInitiatorPasswordEncryptedOk returns a tuple with the IscsiInitiatorPasswordEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetIscsiInitiatorPasswordEncryptedOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiInitiatorPasswordEncrypted) {
		return nil, false
	}
	return o.IscsiInitiatorPasswordEncrypted, true
}

// HasIscsiInitiatorPasswordEncrypted returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasIscsiInitiatorPasswordEncrypted() bool {
	if o != nil && !IsNil(o.IscsiInitiatorPasswordEncrypted) {
		return true
	}

	return false
}

// SetIscsiInitiatorPasswordEncrypted gets a reference to the given string and assigns it to the IscsiInitiatorPasswordEncrypted field.
func (o *ServerInstanceConfiguration) SetIscsiInitiatorPasswordEncrypted(v string) {
	o.IscsiInitiatorPasswordEncrypted = &v
}

// GetControlPanelUrl returns the ControlPanelUrl field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetControlPanelUrl() string {
	if o == nil || IsNil(o.ControlPanelUrl) {
		var ret string
		return ret
	}
	return *o.ControlPanelUrl
}

// GetControlPanelUrlOk returns a tuple with the ControlPanelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetControlPanelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ControlPanelUrl) {
		return nil, false
	}
	return o.ControlPanelUrl, true
}

// HasControlPanelUrl returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasControlPanelUrl() bool {
	if o != nil && !IsNil(o.ControlPanelUrl) {
		return true
	}

	return false
}

// SetControlPanelUrl gets a reference to the given string and assigns it to the ControlPanelUrl field.
func (o *ServerInstanceConfiguration) SetControlPanelUrl(v string) {
	o.ControlPanelUrl = &v
}

// GetEmptyEdit returns the EmptyEdit field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetEmptyEdit() int32 {
	if o == nil || IsNil(o.EmptyEdit) {
		var ret int32
		return ret
	}
	return *o.EmptyEdit
}

// GetEmptyEditOk returns a tuple with the EmptyEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetEmptyEditOk() (*int32, bool) {
	if o == nil || IsNil(o.EmptyEdit) {
		return nil, false
	}
	return o.EmptyEdit, true
}

// HasEmptyEdit returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasEmptyEdit() bool {
	if o != nil && !IsNil(o.EmptyEdit) {
		return true
	}

	return false
}

// SetEmptyEdit gets a reference to the given int32 and assigns it to the EmptyEdit field.
func (o *ServerInstanceConfiguration) SetEmptyEdit(v int32) {
	o.EmptyEdit = &v
}

// GetDeployType returns the DeployType field value
func (o *ServerInstanceConfiguration) GetDeployType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetDeployTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployType, true
}

// SetDeployType sets field value
func (o *ServerInstanceConfiguration) SetDeployType(v string) {
	o.DeployType = v
}

// GetDeployStatus returns the DeployStatus field value
func (o *ServerInstanceConfiguration) GetDeployStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetDeployStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployStatus, true
}

// SetDeployStatus sets field value
func (o *ServerInstanceConfiguration) SetDeployStatus(v string) {
	o.DeployStatus = v
}

// GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetDnsSubdomainChangeId() int32 {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		var ret int32
		return ret
	}
	return *o.DnsSubdomainChangeId
}

// GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetDnsSubdomainChangeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		return nil, false
	}
	return o.DnsSubdomainChangeId, true
}

// HasDnsSubdomainChangeId returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasDnsSubdomainChangeId() bool {
	if o != nil && !IsNil(o.DnsSubdomainChangeId) {
		return true
	}

	return false
}

// SetDnsSubdomainChangeId gets a reference to the given int32 and assigns it to the DnsSubdomainChangeId field.
func (o *ServerInstanceConfiguration) SetDnsSubdomainChangeId(v int32) {
	o.DnsSubdomainChangeId = &v
}

// GetInfrastructureDeployId returns the InfrastructureDeployId field value if set, zero value otherwise.
func (o *ServerInstanceConfiguration) GetInfrastructureDeployId() int32 {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		var ret int32
		return ret
	}
	return *o.InfrastructureDeployId
}

// GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceConfiguration) GetInfrastructureDeployIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		return nil, false
	}
	return o.InfrastructureDeployId, true
}

// HasInfrastructureDeployId returns a boolean if a field has been set.
func (o *ServerInstanceConfiguration) HasInfrastructureDeployId() bool {
	if o != nil && !IsNil(o.InfrastructureDeployId) {
		return true
	}

	return false
}

// SetInfrastructureDeployId gets a reference to the given int32 and assigns it to the InfrastructureDeployId field.
func (o *ServerInstanceConfiguration) SetInfrastructureDeployId(v int32) {
	o.InfrastructureDeployId = &v
}

func (o ServerInstanceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["label"] = o.Label
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	toSerialize["groupId"] = o.GroupId
	if !IsNil(o.DriveIdBootable) {
		toSerialize["driveIdBootable"] = o.DriveIdBootable
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.ServerTypeId) {
		toSerialize["serverTypeId"] = o.ServerTypeId
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.InstanceWanMlagId) {
		toSerialize["instanceWanMlagId"] = o.InstanceWanMlagId
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	if !IsNil(o.RaidProfile) {
		toSerialize["raidProfile"] = o.RaidProfile
	}
	if !IsNil(o.IscsiInitiatorIqn) {
		toSerialize["iscsiInitiatorIqn"] = o.IscsiInitiatorIqn
	}
	if !IsNil(o.IscsiInitiatorUsername) {
		toSerialize["iscsiInitiatorUsername"] = o.IscsiInitiatorUsername
	}
	if !IsNil(o.IscsiInitiatorPasswordEncrypted) {
		toSerialize["iscsiInitiatorPasswordEncrypted"] = o.IscsiInitiatorPasswordEncrypted
	}
	if !IsNil(o.ControlPanelUrl) {
		toSerialize["controlPanelUrl"] = o.ControlPanelUrl
	}
	if !IsNil(o.EmptyEdit) {
		toSerialize["emptyEdit"] = o.EmptyEdit
	}
	toSerialize["deployType"] = o.DeployType
	toSerialize["deployStatus"] = o.DeployStatus
	if !IsNil(o.DnsSubdomainChangeId) {
		toSerialize["dnsSubdomainChangeId"] = o.DnsSubdomainChangeId
	}
	if !IsNil(o.InfrastructureDeployId) {
		toSerialize["infrastructureDeployId"] = o.InfrastructureDeployId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"label",
		"updatedTimestamp",
		"groupId",
		"deployType",
		"deployStatus",
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

	varServerInstanceConfiguration := _ServerInstanceConfiguration{}

	err = json.Unmarshal(data, &varServerInstanceConfiguration)

	if err != nil {
		return err
	}

	*o = ServerInstanceConfiguration(varServerInstanceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "driveIdBootable")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "serverTypeId")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "instanceWanMlagId")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "raidProfile")
		delete(additionalProperties, "iscsiInitiatorIqn")
		delete(additionalProperties, "iscsiInitiatorUsername")
		delete(additionalProperties, "iscsiInitiatorPasswordEncrypted")
		delete(additionalProperties, "controlPanelUrl")
		delete(additionalProperties, "emptyEdit")
		delete(additionalProperties, "deployType")
		delete(additionalProperties, "deployStatus")
		delete(additionalProperties, "dnsSubdomainChangeId")
		delete(additionalProperties, "infrastructureDeployId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceConfiguration struct {
	value *ServerInstanceConfiguration
	isSet bool
}

func (v NullableServerInstanceConfiguration) Get() *ServerInstanceConfiguration {
	return v.value
}

func (v *NullableServerInstanceConfiguration) Set(val *ServerInstanceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceConfiguration(val *ServerInstanceConfiguration) *NullableServerInstanceConfiguration {
	return &NullableServerInstanceConfiguration{value: val, isSet: true}
}

func (v NullableServerInstanceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



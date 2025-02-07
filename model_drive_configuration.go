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

// checks if the DriveConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriveConfiguration{}

// DriveConfiguration struct for DriveConfiguration
type DriveConfiguration struct {
	// Revision of the Drive Configuration
	Revision float32 `json:"revision"`
	// Label of the Drive.
	Label string `json:"label"`
	// Drive Array Id
	GroupId float32 `json:"groupId"`
	ContainerId *float32 `json:"containerId,omitempty"`
	InstanceId *float32 `json:"instanceId,omitempty"`
	// Id of the storage pool the Drive is assigned to
	StoragePoolId *float32 `json:"storagePoolId,omitempty"`
	// Disk size in MB for Drive
	SizeMb float32 `json:"sizeMb"`
	// The name of the storage image used by the Drive.
	StorageImageName *string `json:"storageImageName,omitempty"`
	// The iSCSI Index in hex format of the Drive.
	IscsiIndexHex *string `json:"iscsiIndexHex,omitempty"`
	// Template Id
	TemplateIdOrigin *float32 `json:"templateIdOrigin,omitempty"`
	// The OS Admin Username the Drive will use.
	OsAdminUsername *string `json:"osAdminUsername,omitempty"`
	// The OS Admin Password the Drive will use.
	OsAdminPasswordEncrypted *string `json:"osAdminPasswordEncrypted,omitempty"`
	// Service status of the Drive
	StorageType string `json:"storageType"`
	// Subdomain of the Drive.
	Subdomain *string `json:"subdomain,omitempty"`
	// Timestamp of the Drive last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// SSH port used by the Drive.
	SshPort *float32 `json:"sshPort,omitempty"`
	// Operating system information of the Drive.
	OperatingSystemInfo map[string]interface{} `json:"operatingSystemInfo,omitempty"`
	// Filesystem information of the Drive.
	FilesystemInfo map[string]interface{} `json:"filesystemInfo,omitempty"`
	// Id of the DNS subdomain for the Drive.
	DnsSubdomainChangeId *float32 `json:"dnsSubdomainChangeId,omitempty"`
	// Deploy type of the Drive
	DeployType string `json:"deployType"`
	// Deploy status of the Drive
	DeployStatus string `json:"deployStatus"`
	// Id of the deployment for the Drive.
	InfrastructureDeployId *float32 `json:"infrastructureDeployId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DriveConfiguration DriveConfiguration

// NewDriveConfiguration instantiates a new DriveConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriveConfiguration(revision float32, label string, groupId float32, sizeMb float32, storageType string, updatedTimestamp string, deployType string, deployStatus string) *DriveConfiguration {
	this := DriveConfiguration{}
	this.Revision = revision
	this.Label = label
	this.GroupId = groupId
	this.SizeMb = sizeMb
	this.StorageType = storageType
	this.UpdatedTimestamp = updatedTimestamp
	this.DeployType = deployType
	this.DeployStatus = deployStatus
	return &this
}

// NewDriveConfigurationWithDefaults instantiates a new DriveConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriveConfigurationWithDefaults() *DriveConfiguration {
	this := DriveConfiguration{}
	var storageType string = "iscsi_ssd"
	this.StorageType = storageType
	var deployType string = "create"
	this.DeployType = deployType
	var deployStatus string = "not_started"
	this.DeployStatus = deployStatus
	return &this
}

// GetRevision returns the Revision field value
func (o *DriveConfiguration) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *DriveConfiguration) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value
func (o *DriveConfiguration) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *DriveConfiguration) SetLabel(v string) {
	o.Label = v
}

// GetGroupId returns the GroupId field value
func (o *DriveConfiguration) GetGroupId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetGroupIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *DriveConfiguration) SetGroupId(v float32) {
	o.GroupId = v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *DriveConfiguration) GetContainerId() float32 {
	if o == nil || IsNil(o.ContainerId) {
		var ret float32
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetContainerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *DriveConfiguration) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given float32 and assigns it to the ContainerId field.
func (o *DriveConfiguration) SetContainerId(v float32) {
	o.ContainerId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *DriveConfiguration) GetInstanceId() float32 {
	if o == nil || IsNil(o.InstanceId) {
		var ret float32
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetInstanceIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *DriveConfiguration) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given float32 and assigns it to the InstanceId field.
func (o *DriveConfiguration) SetInstanceId(v float32) {
	o.InstanceId = &v
}

// GetStoragePoolId returns the StoragePoolId field value if set, zero value otherwise.
func (o *DriveConfiguration) GetStoragePoolId() float32 {
	if o == nil || IsNil(o.StoragePoolId) {
		var ret float32
		return ret
	}
	return *o.StoragePoolId
}

// GetStoragePoolIdOk returns a tuple with the StoragePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetStoragePoolIdOk() (*float32, bool) {
	if o == nil || IsNil(o.StoragePoolId) {
		return nil, false
	}
	return o.StoragePoolId, true
}

// HasStoragePoolId returns a boolean if a field has been set.
func (o *DriveConfiguration) HasStoragePoolId() bool {
	if o != nil && !IsNil(o.StoragePoolId) {
		return true
	}

	return false
}

// SetStoragePoolId gets a reference to the given float32 and assigns it to the StoragePoolId field.
func (o *DriveConfiguration) SetStoragePoolId(v float32) {
	o.StoragePoolId = &v
}

// GetSizeMb returns the SizeMb field value
func (o *DriveConfiguration) GetSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetSizeMbOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeMb, true
}

// SetSizeMb sets field value
func (o *DriveConfiguration) SetSizeMb(v float32) {
	o.SizeMb = v
}

// GetStorageImageName returns the StorageImageName field value if set, zero value otherwise.
func (o *DriveConfiguration) GetStorageImageName() string {
	if o == nil || IsNil(o.StorageImageName) {
		var ret string
		return ret
	}
	return *o.StorageImageName
}

// GetStorageImageNameOk returns a tuple with the StorageImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetStorageImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.StorageImageName) {
		return nil, false
	}
	return o.StorageImageName, true
}

// HasStorageImageName returns a boolean if a field has been set.
func (o *DriveConfiguration) HasStorageImageName() bool {
	if o != nil && !IsNil(o.StorageImageName) {
		return true
	}

	return false
}

// SetStorageImageName gets a reference to the given string and assigns it to the StorageImageName field.
func (o *DriveConfiguration) SetStorageImageName(v string) {
	o.StorageImageName = &v
}

// GetIscsiIndexHex returns the IscsiIndexHex field value if set, zero value otherwise.
func (o *DriveConfiguration) GetIscsiIndexHex() string {
	if o == nil || IsNil(o.IscsiIndexHex) {
		var ret string
		return ret
	}
	return *o.IscsiIndexHex
}

// GetIscsiIndexHexOk returns a tuple with the IscsiIndexHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetIscsiIndexHexOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiIndexHex) {
		return nil, false
	}
	return o.IscsiIndexHex, true
}

// HasIscsiIndexHex returns a boolean if a field has been set.
func (o *DriveConfiguration) HasIscsiIndexHex() bool {
	if o != nil && !IsNil(o.IscsiIndexHex) {
		return true
	}

	return false
}

// SetIscsiIndexHex gets a reference to the given string and assigns it to the IscsiIndexHex field.
func (o *DriveConfiguration) SetIscsiIndexHex(v string) {
	o.IscsiIndexHex = &v
}

// GetTemplateIdOrigin returns the TemplateIdOrigin field value if set, zero value otherwise.
func (o *DriveConfiguration) GetTemplateIdOrigin() float32 {
	if o == nil || IsNil(o.TemplateIdOrigin) {
		var ret float32
		return ret
	}
	return *o.TemplateIdOrigin
}

// GetTemplateIdOriginOk returns a tuple with the TemplateIdOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetTemplateIdOriginOk() (*float32, bool) {
	if o == nil || IsNil(o.TemplateIdOrigin) {
		return nil, false
	}
	return o.TemplateIdOrigin, true
}

// HasTemplateIdOrigin returns a boolean if a field has been set.
func (o *DriveConfiguration) HasTemplateIdOrigin() bool {
	if o != nil && !IsNil(o.TemplateIdOrigin) {
		return true
	}

	return false
}

// SetTemplateIdOrigin gets a reference to the given float32 and assigns it to the TemplateIdOrigin field.
func (o *DriveConfiguration) SetTemplateIdOrigin(v float32) {
	o.TemplateIdOrigin = &v
}

// GetOsAdminUsername returns the OsAdminUsername field value if set, zero value otherwise.
func (o *DriveConfiguration) GetOsAdminUsername() string {
	if o == nil || IsNil(o.OsAdminUsername) {
		var ret string
		return ret
	}
	return *o.OsAdminUsername
}

// GetOsAdminUsernameOk returns a tuple with the OsAdminUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetOsAdminUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.OsAdminUsername) {
		return nil, false
	}
	return o.OsAdminUsername, true
}

// HasOsAdminUsername returns a boolean if a field has been set.
func (o *DriveConfiguration) HasOsAdminUsername() bool {
	if o != nil && !IsNil(o.OsAdminUsername) {
		return true
	}

	return false
}

// SetOsAdminUsername gets a reference to the given string and assigns it to the OsAdminUsername field.
func (o *DriveConfiguration) SetOsAdminUsername(v string) {
	o.OsAdminUsername = &v
}

// GetOsAdminPasswordEncrypted returns the OsAdminPasswordEncrypted field value if set, zero value otherwise.
func (o *DriveConfiguration) GetOsAdminPasswordEncrypted() string {
	if o == nil || IsNil(o.OsAdminPasswordEncrypted) {
		var ret string
		return ret
	}
	return *o.OsAdminPasswordEncrypted
}

// GetOsAdminPasswordEncryptedOk returns a tuple with the OsAdminPasswordEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetOsAdminPasswordEncryptedOk() (*string, bool) {
	if o == nil || IsNil(o.OsAdminPasswordEncrypted) {
		return nil, false
	}
	return o.OsAdminPasswordEncrypted, true
}

// HasOsAdminPasswordEncrypted returns a boolean if a field has been set.
func (o *DriveConfiguration) HasOsAdminPasswordEncrypted() bool {
	if o != nil && !IsNil(o.OsAdminPasswordEncrypted) {
		return true
	}

	return false
}

// SetOsAdminPasswordEncrypted gets a reference to the given string and assigns it to the OsAdminPasswordEncrypted field.
func (o *DriveConfiguration) SetOsAdminPasswordEncrypted(v string) {
	o.OsAdminPasswordEncrypted = &v
}

// GetStorageType returns the StorageType field value
func (o *DriveConfiguration) GetStorageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *DriveConfiguration) SetStorageType(v string) {
	o.StorageType = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DriveConfiguration) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DriveConfiguration) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DriveConfiguration) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *DriveConfiguration) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *DriveConfiguration) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *DriveConfiguration) GetSshPort() float32 {
	if o == nil || IsNil(o.SshPort) {
		var ret float32
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetSshPortOk() (*float32, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *DriveConfiguration) HasSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given float32 and assigns it to the SshPort field.
func (o *DriveConfiguration) SetSshPort(v float32) {
	o.SshPort = &v
}

// GetOperatingSystemInfo returns the OperatingSystemInfo field value if set, zero value otherwise.
func (o *DriveConfiguration) GetOperatingSystemInfo() map[string]interface{} {
	if o == nil || IsNil(o.OperatingSystemInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.OperatingSystemInfo
}

// GetOperatingSystemInfoOk returns a tuple with the OperatingSystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetOperatingSystemInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.OperatingSystemInfo) {
		return map[string]interface{}{}, false
	}
	return o.OperatingSystemInfo, true
}

// HasOperatingSystemInfo returns a boolean if a field has been set.
func (o *DriveConfiguration) HasOperatingSystemInfo() bool {
	if o != nil && !IsNil(o.OperatingSystemInfo) {
		return true
	}

	return false
}

// SetOperatingSystemInfo gets a reference to the given map[string]interface{} and assigns it to the OperatingSystemInfo field.
func (o *DriveConfiguration) SetOperatingSystemInfo(v map[string]interface{}) {
	o.OperatingSystemInfo = v
}

// GetFilesystemInfo returns the FilesystemInfo field value if set, zero value otherwise.
func (o *DriveConfiguration) GetFilesystemInfo() map[string]interface{} {
	if o == nil || IsNil(o.FilesystemInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.FilesystemInfo
}

// GetFilesystemInfoOk returns a tuple with the FilesystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetFilesystemInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FilesystemInfo) {
		return map[string]interface{}{}, false
	}
	return o.FilesystemInfo, true
}

// HasFilesystemInfo returns a boolean if a field has been set.
func (o *DriveConfiguration) HasFilesystemInfo() bool {
	if o != nil && !IsNil(o.FilesystemInfo) {
		return true
	}

	return false
}

// SetFilesystemInfo gets a reference to the given map[string]interface{} and assigns it to the FilesystemInfo field.
func (o *DriveConfiguration) SetFilesystemInfo(v map[string]interface{}) {
	o.FilesystemInfo = v
}

// GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field value if set, zero value otherwise.
func (o *DriveConfiguration) GetDnsSubdomainChangeId() float32 {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainChangeId
}

// GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		return nil, false
	}
	return o.DnsSubdomainChangeId, true
}

// HasDnsSubdomainChangeId returns a boolean if a field has been set.
func (o *DriveConfiguration) HasDnsSubdomainChangeId() bool {
	if o != nil && !IsNil(o.DnsSubdomainChangeId) {
		return true
	}

	return false
}

// SetDnsSubdomainChangeId gets a reference to the given float32 and assigns it to the DnsSubdomainChangeId field.
func (o *DriveConfiguration) SetDnsSubdomainChangeId(v float32) {
	o.DnsSubdomainChangeId = &v
}

// GetDeployType returns the DeployType field value
func (o *DriveConfiguration) GetDeployType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetDeployTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployType, true
}

// SetDeployType sets field value
func (o *DriveConfiguration) SetDeployType(v string) {
	o.DeployType = v
}

// GetDeployStatus returns the DeployStatus field value
func (o *DriveConfiguration) GetDeployStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetDeployStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployStatus, true
}

// SetDeployStatus sets field value
func (o *DriveConfiguration) SetDeployStatus(v string) {
	o.DeployStatus = v
}

// GetInfrastructureDeployId returns the InfrastructureDeployId field value if set, zero value otherwise.
func (o *DriveConfiguration) GetInfrastructureDeployId() float32 {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		var ret float32
		return ret
	}
	return *o.InfrastructureDeployId
}

// GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveConfiguration) GetInfrastructureDeployIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		return nil, false
	}
	return o.InfrastructureDeployId, true
}

// HasInfrastructureDeployId returns a boolean if a field has been set.
func (o *DriveConfiguration) HasInfrastructureDeployId() bool {
	if o != nil && !IsNil(o.InfrastructureDeployId) {
		return true
	}

	return false
}

// SetInfrastructureDeployId gets a reference to the given float32 and assigns it to the InfrastructureDeployId field.
func (o *DriveConfiguration) SetInfrastructureDeployId(v float32) {
	o.InfrastructureDeployId = &v
}

func (o DriveConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriveConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["label"] = o.Label
	toSerialize["groupId"] = o.GroupId
	if !IsNil(o.ContainerId) {
		toSerialize["containerId"] = o.ContainerId
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.StoragePoolId) {
		toSerialize["storagePoolId"] = o.StoragePoolId
	}
	toSerialize["sizeMb"] = o.SizeMb
	if !IsNil(o.StorageImageName) {
		toSerialize["storageImageName"] = o.StorageImageName
	}
	if !IsNil(o.IscsiIndexHex) {
		toSerialize["iscsiIndexHex"] = o.IscsiIndexHex
	}
	if !IsNil(o.TemplateIdOrigin) {
		toSerialize["templateIdOrigin"] = o.TemplateIdOrigin
	}
	if !IsNil(o.OsAdminUsername) {
		toSerialize["osAdminUsername"] = o.OsAdminUsername
	}
	if !IsNil(o.OsAdminPasswordEncrypted) {
		toSerialize["osAdminPasswordEncrypted"] = o.OsAdminPasswordEncrypted
	}
	toSerialize["storageType"] = o.StorageType
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	if !IsNil(o.SshPort) {
		toSerialize["sshPort"] = o.SshPort
	}
	if !IsNil(o.OperatingSystemInfo) {
		toSerialize["operatingSystemInfo"] = o.OperatingSystemInfo
	}
	if !IsNil(o.FilesystemInfo) {
		toSerialize["filesystemInfo"] = o.FilesystemInfo
	}
	if !IsNil(o.DnsSubdomainChangeId) {
		toSerialize["dnsSubdomainChangeId"] = o.DnsSubdomainChangeId
	}
	toSerialize["deployType"] = o.DeployType
	toSerialize["deployStatus"] = o.DeployStatus
	if !IsNil(o.InfrastructureDeployId) {
		toSerialize["infrastructureDeployId"] = o.InfrastructureDeployId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DriveConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"label",
		"groupId",
		"sizeMb",
		"storageType",
		"updatedTimestamp",
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

	varDriveConfiguration := _DriveConfiguration{}

	err = json.Unmarshal(data, &varDriveConfiguration)

	if err != nil {
		return err
	}

	*o = DriveConfiguration(varDriveConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "containerId")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "storagePoolId")
		delete(additionalProperties, "sizeMb")
		delete(additionalProperties, "storageImageName")
		delete(additionalProperties, "iscsiIndexHex")
		delete(additionalProperties, "templateIdOrigin")
		delete(additionalProperties, "osAdminUsername")
		delete(additionalProperties, "osAdminPasswordEncrypted")
		delete(additionalProperties, "storageType")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "sshPort")
		delete(additionalProperties, "operatingSystemInfo")
		delete(additionalProperties, "filesystemInfo")
		delete(additionalProperties, "dnsSubdomainChangeId")
		delete(additionalProperties, "deployType")
		delete(additionalProperties, "deployStatus")
		delete(additionalProperties, "infrastructureDeployId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDriveConfiguration struct {
	value *DriveConfiguration
	isSet bool
}

func (v NullableDriveConfiguration) Get() *DriveConfiguration {
	return v.value
}

func (v *NullableDriveConfiguration) Set(val *DriveConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDriveConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDriveConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriveConfiguration(val *DriveConfiguration) *NullableDriveConfiguration {
	return &NullableDriveConfiguration{value: val, isSet: true}
}

func (v NullableDriveConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriveConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the Storage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Storage{}

// Storage struct for Storage
type Storage struct {
	// Id of the Storage
	Id float32 `json:"id"`
	// Revision of the Storage
	Revision float32 `json:"revision"`
	// Id of the owner
	UserId *float32 `json:"userId,omitempty"`
	// Id of the site
	SiteId float32 `json:"siteId"`
	// The name of the datacenter where the storage is located.
	DatacenterName string `json:"datacenterName"`
	// Storage driver
	Driver string `json:"driver"`
	// Storage technology
	Technologies []string `json:"technologies"`
	// Storage type
	Type string `json:"type"`
	// Storage status
	Status string `json:"status"`
	// Name of the storage
	Name string `json:"name"`
	// ISCSI host
	IscsiHost *string `json:"iscsiHost,omitempty"`
	// ISCSI port
	IscsiPort *float32 `json:"iscsiPort,omitempty"`
	// Management host
	ManagementHost string `json:"managementHost"`
	// Username
	Username string `json:"username"`
	// Password encrypted
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty"`
	// Options for the storage
	Options *StorageOptions `json:"options,omitempty"`
	// Specifies if the storage is in maintenance
	InMaintenance *float32 `json:"inMaintenance,omitempty"`
	// Target IQN
	TargetIQN *string `json:"targetIQN,omitempty"`
	// Specifies if the storage is experimental
	IsExperimental *float32 `json:"isExperimental,omitempty"`
	// Specifies the drive priority
	DrivePriority *float32 `json:"drivePriority,omitempty"`
	// Specifies the shared drive priority
	SharedDrivePriority *float32 `json:"sharedDrivePriority,omitempty"`
	// Alternate SAN IPs
	AlternateSanIPs []string `json:"alternateSanIPs,omitempty"`
	// Tags
	Tags []string `json:"tags,omitempty"`
	// Port group allocation order
	PortGroupAllocationOrder map[string]interface{} `json:"portGroupAllocationOrder,omitempty"`
	// Port group physical ports
	PortGroupPhysicalPorts map[string]interface{} `json:"portGroupPhysicalPorts,omitempty"`
	// Subnet type
	SubnetType string `json:"subnetType"`
	JobStatistics *JobGroupStatistics `json:"jobStatistics,omitempty"`
	// The extension execution info of the storage.
	ExtensionInfo *ExtensionExecutionInfo `json:"extensionInfo,omitempty"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Storage Storage

// NewStorage instantiates a new Storage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorage(id float32, revision float32, siteId float32, datacenterName string, driver string, technologies []string, type_ string, status string, name string, managementHost string, username string, subnetType string) *Storage {
	this := Storage{}
	this.Id = id
	this.Revision = revision
	this.SiteId = siteId
	this.DatacenterName = datacenterName
	this.Driver = driver
	this.Technologies = technologies
	this.Type = type_
	this.Status = status
	this.Name = name
	this.ManagementHost = managementHost
	this.Username = username
	this.SubnetType = subnetType
	return &this
}

// NewStorageWithDefaults instantiates a new Storage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageWithDefaults() *Storage {
	this := Storage{}
	return &this
}

// GetId returns the Id field value
func (o *Storage) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Storage) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Storage) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *Storage) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *Storage) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *Storage) SetRevision(v float32) {
	o.Revision = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Storage) GetUserId() float32 {
	if o == nil || IsNil(o.UserId) {
		var ret float32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetUserIdOk() (*float32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Storage) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given float32 and assigns it to the UserId field.
func (o *Storage) SetUserId(v float32) {
	o.UserId = &v
}

// GetSiteId returns the SiteId field value
func (o *Storage) GetSiteId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *Storage) GetSiteIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *Storage) SetSiteId(v float32) {
	o.SiteId = v
}

// GetDatacenterName returns the DatacenterName field value
func (o *Storage) GetDatacenterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value
// and a boolean to check if the value has been set.
func (o *Storage) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatacenterName, true
}

// SetDatacenterName sets field value
func (o *Storage) SetDatacenterName(v string) {
	o.DatacenterName = v
}

// GetDriver returns the Driver field value
func (o *Storage) GetDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Driver
}

// GetDriverOk returns a tuple with the Driver field value
// and a boolean to check if the value has been set.
func (o *Storage) GetDriverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Driver, true
}

// SetDriver sets field value
func (o *Storage) SetDriver(v string) {
	o.Driver = v
}

// GetTechnologies returns the Technologies field value
func (o *Storage) GetTechnologies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Technologies
}

// GetTechnologiesOk returns a tuple with the Technologies field value
// and a boolean to check if the value has been set.
func (o *Storage) GetTechnologiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Technologies, true
}

// SetTechnologies sets field value
func (o *Storage) SetTechnologies(v []string) {
	o.Technologies = v
}

// GetType returns the Type field value
func (o *Storage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Storage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Storage) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *Storage) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Storage) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Storage) SetStatus(v string) {
	o.Status = v
}

// GetName returns the Name field value
func (o *Storage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Storage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Storage) SetName(v string) {
	o.Name = v
}

// GetIscsiHost returns the IscsiHost field value if set, zero value otherwise.
func (o *Storage) GetIscsiHost() string {
	if o == nil || IsNil(o.IscsiHost) {
		var ret string
		return ret
	}
	return *o.IscsiHost
}

// GetIscsiHostOk returns a tuple with the IscsiHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetIscsiHostOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiHost) {
		return nil, false
	}
	return o.IscsiHost, true
}

// HasIscsiHost returns a boolean if a field has been set.
func (o *Storage) HasIscsiHost() bool {
	if o != nil && !IsNil(o.IscsiHost) {
		return true
	}

	return false
}

// SetIscsiHost gets a reference to the given string and assigns it to the IscsiHost field.
func (o *Storage) SetIscsiHost(v string) {
	o.IscsiHost = &v
}

// GetIscsiPort returns the IscsiPort field value if set, zero value otherwise.
func (o *Storage) GetIscsiPort() float32 {
	if o == nil || IsNil(o.IscsiPort) {
		var ret float32
		return ret
	}
	return *o.IscsiPort
}

// GetIscsiPortOk returns a tuple with the IscsiPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetIscsiPortOk() (*float32, bool) {
	if o == nil || IsNil(o.IscsiPort) {
		return nil, false
	}
	return o.IscsiPort, true
}

// HasIscsiPort returns a boolean if a field has been set.
func (o *Storage) HasIscsiPort() bool {
	if o != nil && !IsNil(o.IscsiPort) {
		return true
	}

	return false
}

// SetIscsiPort gets a reference to the given float32 and assigns it to the IscsiPort field.
func (o *Storage) SetIscsiPort(v float32) {
	o.IscsiPort = &v
}

// GetManagementHost returns the ManagementHost field value
func (o *Storage) GetManagementHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagementHost
}

// GetManagementHostOk returns a tuple with the ManagementHost field value
// and a boolean to check if the value has been set.
func (o *Storage) GetManagementHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementHost, true
}

// SetManagementHost sets field value
func (o *Storage) SetManagementHost(v string) {
	o.ManagementHost = v
}

// GetUsername returns the Username field value
func (o *Storage) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Storage) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Storage) SetUsername(v string) {
	o.Username = v
}

// GetPasswordEncrypted returns the PasswordEncrypted field value if set, zero value otherwise.
func (o *Storage) GetPasswordEncrypted() string {
	if o == nil || IsNil(o.PasswordEncrypted) {
		var ret string
		return ret
	}
	return *o.PasswordEncrypted
}

// GetPasswordEncryptedOk returns a tuple with the PasswordEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetPasswordEncryptedOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordEncrypted) {
		return nil, false
	}
	return o.PasswordEncrypted, true
}

// HasPasswordEncrypted returns a boolean if a field has been set.
func (o *Storage) HasPasswordEncrypted() bool {
	if o != nil && !IsNil(o.PasswordEncrypted) {
		return true
	}

	return false
}

// SetPasswordEncrypted gets a reference to the given string and assigns it to the PasswordEncrypted field.
func (o *Storage) SetPasswordEncrypted(v string) {
	o.PasswordEncrypted = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Storage) GetOptions() StorageOptions {
	if o == nil || IsNil(o.Options) {
		var ret StorageOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetOptionsOk() (*StorageOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Storage) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given StorageOptions and assigns it to the Options field.
func (o *Storage) SetOptions(v StorageOptions) {
	o.Options = &v
}

// GetInMaintenance returns the InMaintenance field value if set, zero value otherwise.
func (o *Storage) GetInMaintenance() float32 {
	if o == nil || IsNil(o.InMaintenance) {
		var ret float32
		return ret
	}
	return *o.InMaintenance
}

// GetInMaintenanceOk returns a tuple with the InMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetInMaintenanceOk() (*float32, bool) {
	if o == nil || IsNil(o.InMaintenance) {
		return nil, false
	}
	return o.InMaintenance, true
}

// HasInMaintenance returns a boolean if a field has been set.
func (o *Storage) HasInMaintenance() bool {
	if o != nil && !IsNil(o.InMaintenance) {
		return true
	}

	return false
}

// SetInMaintenance gets a reference to the given float32 and assigns it to the InMaintenance field.
func (o *Storage) SetInMaintenance(v float32) {
	o.InMaintenance = &v
}

// GetTargetIQN returns the TargetIQN field value if set, zero value otherwise.
func (o *Storage) GetTargetIQN() string {
	if o == nil || IsNil(o.TargetIQN) {
		var ret string
		return ret
	}
	return *o.TargetIQN
}

// GetTargetIQNOk returns a tuple with the TargetIQN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetTargetIQNOk() (*string, bool) {
	if o == nil || IsNil(o.TargetIQN) {
		return nil, false
	}
	return o.TargetIQN, true
}

// HasTargetIQN returns a boolean if a field has been set.
func (o *Storage) HasTargetIQN() bool {
	if o != nil && !IsNil(o.TargetIQN) {
		return true
	}

	return false
}

// SetTargetIQN gets a reference to the given string and assigns it to the TargetIQN field.
func (o *Storage) SetTargetIQN(v string) {
	o.TargetIQN = &v
}

// GetIsExperimental returns the IsExperimental field value if set, zero value otherwise.
func (o *Storage) GetIsExperimental() float32 {
	if o == nil || IsNil(o.IsExperimental) {
		var ret float32
		return ret
	}
	return *o.IsExperimental
}

// GetIsExperimentalOk returns a tuple with the IsExperimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetIsExperimentalOk() (*float32, bool) {
	if o == nil || IsNil(o.IsExperimental) {
		return nil, false
	}
	return o.IsExperimental, true
}

// HasIsExperimental returns a boolean if a field has been set.
func (o *Storage) HasIsExperimental() bool {
	if o != nil && !IsNil(o.IsExperimental) {
		return true
	}

	return false
}

// SetIsExperimental gets a reference to the given float32 and assigns it to the IsExperimental field.
func (o *Storage) SetIsExperimental(v float32) {
	o.IsExperimental = &v
}

// GetDrivePriority returns the DrivePriority field value if set, zero value otherwise.
func (o *Storage) GetDrivePriority() float32 {
	if o == nil || IsNil(o.DrivePriority) {
		var ret float32
		return ret
	}
	return *o.DrivePriority
}

// GetDrivePriorityOk returns a tuple with the DrivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetDrivePriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.DrivePriority) {
		return nil, false
	}
	return o.DrivePriority, true
}

// HasDrivePriority returns a boolean if a field has been set.
func (o *Storage) HasDrivePriority() bool {
	if o != nil && !IsNil(o.DrivePriority) {
		return true
	}

	return false
}

// SetDrivePriority gets a reference to the given float32 and assigns it to the DrivePriority field.
func (o *Storage) SetDrivePriority(v float32) {
	o.DrivePriority = &v
}

// GetSharedDrivePriority returns the SharedDrivePriority field value if set, zero value otherwise.
func (o *Storage) GetSharedDrivePriority() float32 {
	if o == nil || IsNil(o.SharedDrivePriority) {
		var ret float32
		return ret
	}
	return *o.SharedDrivePriority
}

// GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetSharedDrivePriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.SharedDrivePriority) {
		return nil, false
	}
	return o.SharedDrivePriority, true
}

// HasSharedDrivePriority returns a boolean if a field has been set.
func (o *Storage) HasSharedDrivePriority() bool {
	if o != nil && !IsNil(o.SharedDrivePriority) {
		return true
	}

	return false
}

// SetSharedDrivePriority gets a reference to the given float32 and assigns it to the SharedDrivePriority field.
func (o *Storage) SetSharedDrivePriority(v float32) {
	o.SharedDrivePriority = &v
}

// GetAlternateSanIPs returns the AlternateSanIPs field value if set, zero value otherwise.
func (o *Storage) GetAlternateSanIPs() []string {
	if o == nil || IsNil(o.AlternateSanIPs) {
		var ret []string
		return ret
	}
	return o.AlternateSanIPs
}

// GetAlternateSanIPsOk returns a tuple with the AlternateSanIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetAlternateSanIPsOk() ([]string, bool) {
	if o == nil || IsNil(o.AlternateSanIPs) {
		return nil, false
	}
	return o.AlternateSanIPs, true
}

// HasAlternateSanIPs returns a boolean if a field has been set.
func (o *Storage) HasAlternateSanIPs() bool {
	if o != nil && !IsNil(o.AlternateSanIPs) {
		return true
	}

	return false
}

// SetAlternateSanIPs gets a reference to the given []string and assigns it to the AlternateSanIPs field.
func (o *Storage) SetAlternateSanIPs(v []string) {
	o.AlternateSanIPs = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Storage) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Storage) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Storage) SetTags(v []string) {
	o.Tags = v
}

// GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field value if set, zero value otherwise.
func (o *Storage) GetPortGroupAllocationOrder() map[string]interface{} {
	if o == nil || IsNil(o.PortGroupAllocationOrder) {
		var ret map[string]interface{}
		return ret
	}
	return o.PortGroupAllocationOrder
}

// GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetPortGroupAllocationOrderOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PortGroupAllocationOrder) {
		return map[string]interface{}{}, false
	}
	return o.PortGroupAllocationOrder, true
}

// HasPortGroupAllocationOrder returns a boolean if a field has been set.
func (o *Storage) HasPortGroupAllocationOrder() bool {
	if o != nil && !IsNil(o.PortGroupAllocationOrder) {
		return true
	}

	return false
}

// SetPortGroupAllocationOrder gets a reference to the given map[string]interface{} and assigns it to the PortGroupAllocationOrder field.
func (o *Storage) SetPortGroupAllocationOrder(v map[string]interface{}) {
	o.PortGroupAllocationOrder = v
}

// GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field value if set, zero value otherwise.
func (o *Storage) GetPortGroupPhysicalPorts() map[string]interface{} {
	if o == nil || IsNil(o.PortGroupPhysicalPorts) {
		var ret map[string]interface{}
		return ret
	}
	return o.PortGroupPhysicalPorts
}

// GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetPortGroupPhysicalPortsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PortGroupPhysicalPorts) {
		return map[string]interface{}{}, false
	}
	return o.PortGroupPhysicalPorts, true
}

// HasPortGroupPhysicalPorts returns a boolean if a field has been set.
func (o *Storage) HasPortGroupPhysicalPorts() bool {
	if o != nil && !IsNil(o.PortGroupPhysicalPorts) {
		return true
	}

	return false
}

// SetPortGroupPhysicalPorts gets a reference to the given map[string]interface{} and assigns it to the PortGroupPhysicalPorts field.
func (o *Storage) SetPortGroupPhysicalPorts(v map[string]interface{}) {
	o.PortGroupPhysicalPorts = v
}

// GetSubnetType returns the SubnetType field value
func (o *Storage) GetSubnetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetType
}

// GetSubnetTypeOk returns a tuple with the SubnetType field value
// and a boolean to check if the value has been set.
func (o *Storage) GetSubnetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetType, true
}

// SetSubnetType sets field value
func (o *Storage) SetSubnetType(v string) {
	o.SubnetType = v
}

// GetJobStatistics returns the JobStatistics field value if set, zero value otherwise.
func (o *Storage) GetJobStatistics() JobGroupStatistics {
	if o == nil || IsNil(o.JobStatistics) {
		var ret JobGroupStatistics
		return ret
	}
	return *o.JobStatistics
}

// GetJobStatisticsOk returns a tuple with the JobStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetJobStatisticsOk() (*JobGroupStatistics, bool) {
	if o == nil || IsNil(o.JobStatistics) {
		return nil, false
	}
	return o.JobStatistics, true
}

// HasJobStatistics returns a boolean if a field has been set.
func (o *Storage) HasJobStatistics() bool {
	if o != nil && !IsNil(o.JobStatistics) {
		return true
	}

	return false
}

// SetJobStatistics gets a reference to the given JobGroupStatistics and assigns it to the JobStatistics field.
func (o *Storage) SetJobStatistics(v JobGroupStatistics) {
	o.JobStatistics = &v
}

// GetExtensionInfo returns the ExtensionInfo field value if set, zero value otherwise.
func (o *Storage) GetExtensionInfo() ExtensionExecutionInfo {
	if o == nil || IsNil(o.ExtensionInfo) {
		var ret ExtensionExecutionInfo
		return ret
	}
	return *o.ExtensionInfo
}

// GetExtensionInfoOk returns a tuple with the ExtensionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetExtensionInfoOk() (*ExtensionExecutionInfo, bool) {
	if o == nil || IsNil(o.ExtensionInfo) {
		return nil, false
	}
	return o.ExtensionInfo, true
}

// HasExtensionInfo returns a boolean if a field has been set.
func (o *Storage) HasExtensionInfo() bool {
	if o != nil && !IsNil(o.ExtensionInfo) {
		return true
	}

	return false
}

// SetExtensionInfo gets a reference to the given ExtensionExecutionInfo and assigns it to the ExtensionInfo field.
func (o *Storage) SetExtensionInfo(v ExtensionExecutionInfo) {
	o.ExtensionInfo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Storage) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storage) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Storage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Storage) SetLinks(v []Link) {
	o.Links = v
}

func (o Storage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Storage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	toSerialize["siteId"] = o.SiteId
	toSerialize["datacenterName"] = o.DatacenterName
	toSerialize["driver"] = o.Driver
	toSerialize["technologies"] = o.Technologies
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["name"] = o.Name
	if !IsNil(o.IscsiHost) {
		toSerialize["iscsiHost"] = o.IscsiHost
	}
	if !IsNil(o.IscsiPort) {
		toSerialize["iscsiPort"] = o.IscsiPort
	}
	toSerialize["managementHost"] = o.ManagementHost
	toSerialize["username"] = o.Username
	if !IsNil(o.PasswordEncrypted) {
		toSerialize["passwordEncrypted"] = o.PasswordEncrypted
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.InMaintenance) {
		toSerialize["inMaintenance"] = o.InMaintenance
	}
	if !IsNil(o.TargetIQN) {
		toSerialize["targetIQN"] = o.TargetIQN
	}
	if !IsNil(o.IsExperimental) {
		toSerialize["isExperimental"] = o.IsExperimental
	}
	if !IsNil(o.DrivePriority) {
		toSerialize["drivePriority"] = o.DrivePriority
	}
	if !IsNil(o.SharedDrivePriority) {
		toSerialize["sharedDrivePriority"] = o.SharedDrivePriority
	}
	if !IsNil(o.AlternateSanIPs) {
		toSerialize["alternateSanIPs"] = o.AlternateSanIPs
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.PortGroupAllocationOrder) {
		toSerialize["portGroupAllocationOrder"] = o.PortGroupAllocationOrder
	}
	if !IsNil(o.PortGroupPhysicalPorts) {
		toSerialize["portGroupPhysicalPorts"] = o.PortGroupPhysicalPorts
	}
	toSerialize["subnetType"] = o.SubnetType
	if !IsNil(o.JobStatistics) {
		toSerialize["jobStatistics"] = o.JobStatistics
	}
	if !IsNil(o.ExtensionInfo) {
		toSerialize["extensionInfo"] = o.ExtensionInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Storage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"revision",
		"siteId",
		"datacenterName",
		"driver",
		"technologies",
		"type",
		"status",
		"name",
		"managementHost",
		"username",
		"subnetType",
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

	varStorage := _Storage{}

	err = json.Unmarshal(data, &varStorage)

	if err != nil {
		return err
	}

	*o = Storage(varStorage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "datacenterName")
		delete(additionalProperties, "driver")
		delete(additionalProperties, "technologies")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "name")
		delete(additionalProperties, "iscsiHost")
		delete(additionalProperties, "iscsiPort")
		delete(additionalProperties, "managementHost")
		delete(additionalProperties, "username")
		delete(additionalProperties, "passwordEncrypted")
		delete(additionalProperties, "options")
		delete(additionalProperties, "inMaintenance")
		delete(additionalProperties, "targetIQN")
		delete(additionalProperties, "isExperimental")
		delete(additionalProperties, "drivePriority")
		delete(additionalProperties, "sharedDrivePriority")
		delete(additionalProperties, "alternateSanIPs")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "portGroupAllocationOrder")
		delete(additionalProperties, "portGroupPhysicalPorts")
		delete(additionalProperties, "subnetType")
		delete(additionalProperties, "jobStatistics")
		delete(additionalProperties, "extensionInfo")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorage struct {
	value *Storage
	isSet bool
}

func (v NullableStorage) Get() *Storage {
	return v.value
}

func (v *NullableStorage) Set(val *Storage) {
	v.value = val
	v.isSet = true
}

func (v NullableStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorage(val *Storage) *NullableStorage {
	return &NullableStorage{value: val, isSet: true}
}

func (v NullableStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



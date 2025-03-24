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

// checks if the CreateStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStorage{}

// CreateStorage struct for CreateStorage
type CreateStorage struct {
	// Id of the owner
	UserId *float32 `json:"userId,omitempty"`
	// Id of the site
	SiteId float32 `json:"siteId"`
	// Storage driver
	StorageDriver string `json:"storageDriver"`
	// Storage technology
	StorageTechnology string `json:"storageTechnology"`
	// Storage type
	StorageType string `json:"storageType"`
	// Storage status
	Status string `json:"status"`
	// Total capacity in MB
	TotalCapacity *float32 `json:"totalCapacity,omitempty"`
	// Usable capacity in MB
	UsableCapacity *float32 `json:"usableCapacity,omitempty"`
	// Free capacity in MB
	FreeCapacity *float32 `json:"freeCapacity,omitempty"`
	// Virtual used capacity in MB
	VirtualUsedCapacity *float32 `json:"virtualUsedCapacity,omitempty"`
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
	PasswordEncrypted string `json:"passwordEncrypted"`
	// Specifies if the storage is in maintenance
	InMaintenance float32 `json:"inMaintenance"`
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
	// Default IO limit policy
	DefaultIoLimitPolicy *string `json:"defaultIoLimitPolicy,omitempty"`
	// Subnet type
	SubnetType string `json:"subnetType"`
	// Array id
	ArrayId *string `json:"arrayId,omitempty"`
	// Director id
	DirectorId *string `json:"directorId,omitempty"`
	// S3 hostname
	S3Hostname *string `json:"s3Hostname,omitempty"`
	// S3 port
	S3Port *string `json:"s3Port,omitempty"`
	JobInfo *JobInfo `json:"jobInfo,omitempty"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	// Options for the storage
	Options *UpdateStorageOptions `json:"options,omitempty"`
	// The password to use.
	Password string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _CreateStorage CreateStorage

// NewCreateStorage instantiates a new CreateStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStorage(siteId float32, storageDriver string, storageTechnology string, storageType string, status string, name string, managementHost string, username string, passwordEncrypted string, inMaintenance float32, subnetType string, password string) *CreateStorage {
	this := CreateStorage{}
	this.SiteId = siteId
	this.StorageDriver = storageDriver
	this.StorageTechnology = storageTechnology
	this.StorageType = storageType
	this.Status = status
	this.Name = name
	this.ManagementHost = managementHost
	this.Username = username
	this.PasswordEncrypted = passwordEncrypted
	this.InMaintenance = inMaintenance
	this.SubnetType = subnetType
	this.Password = password
	return &this
}

// NewCreateStorageWithDefaults instantiates a new CreateStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStorageWithDefaults() *CreateStorage {
	this := CreateStorage{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CreateStorage) GetUserId() float32 {
	if o == nil || IsNil(o.UserId) {
		var ret float32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetUserIdOk() (*float32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateStorage) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given float32 and assigns it to the UserId field.
func (o *CreateStorage) SetUserId(v float32) {
	o.UserId = &v
}

// GetSiteId returns the SiteId field value
func (o *CreateStorage) GetSiteId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetSiteIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *CreateStorage) SetSiteId(v float32) {
	o.SiteId = v
}

// GetStorageDriver returns the StorageDriver field value
func (o *CreateStorage) GetStorageDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageDriver
}

// GetStorageDriverOk returns a tuple with the StorageDriver field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetStorageDriverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageDriver, true
}

// SetStorageDriver sets field value
func (o *CreateStorage) SetStorageDriver(v string) {
	o.StorageDriver = v
}

// GetStorageTechnology returns the StorageTechnology field value
func (o *CreateStorage) GetStorageTechnology() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageTechnology
}

// GetStorageTechnologyOk returns a tuple with the StorageTechnology field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetStorageTechnologyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageTechnology, true
}

// SetStorageTechnology sets field value
func (o *CreateStorage) SetStorageTechnology(v string) {
	o.StorageTechnology = v
}

// GetStorageType returns the StorageType field value
func (o *CreateStorage) GetStorageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *CreateStorage) SetStorageType(v string) {
	o.StorageType = v
}

// GetStatus returns the Status field value
func (o *CreateStorage) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateStorage) SetStatus(v string) {
	o.Status = v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *CreateStorage) GetTotalCapacity() float32 {
	if o == nil || IsNil(o.TotalCapacity) {
		var ret float32
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetTotalCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCapacity) {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *CreateStorage) HasTotalCapacity() bool {
	if o != nil && !IsNil(o.TotalCapacity) {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given float32 and assigns it to the TotalCapacity field.
func (o *CreateStorage) SetTotalCapacity(v float32) {
	o.TotalCapacity = &v
}

// GetUsableCapacity returns the UsableCapacity field value if set, zero value otherwise.
func (o *CreateStorage) GetUsableCapacity() float32 {
	if o == nil || IsNil(o.UsableCapacity) {
		var ret float32
		return ret
	}
	return *o.UsableCapacity
}

// GetUsableCapacityOk returns a tuple with the UsableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetUsableCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.UsableCapacity) {
		return nil, false
	}
	return o.UsableCapacity, true
}

// HasUsableCapacity returns a boolean if a field has been set.
func (o *CreateStorage) HasUsableCapacity() bool {
	if o != nil && !IsNil(o.UsableCapacity) {
		return true
	}

	return false
}

// SetUsableCapacity gets a reference to the given float32 and assigns it to the UsableCapacity field.
func (o *CreateStorage) SetUsableCapacity(v float32) {
	o.UsableCapacity = &v
}

// GetFreeCapacity returns the FreeCapacity field value if set, zero value otherwise.
func (o *CreateStorage) GetFreeCapacity() float32 {
	if o == nil || IsNil(o.FreeCapacity) {
		var ret float32
		return ret
	}
	return *o.FreeCapacity
}

// GetFreeCapacityOk returns a tuple with the FreeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetFreeCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.FreeCapacity) {
		return nil, false
	}
	return o.FreeCapacity, true
}

// HasFreeCapacity returns a boolean if a field has been set.
func (o *CreateStorage) HasFreeCapacity() bool {
	if o != nil && !IsNil(o.FreeCapacity) {
		return true
	}

	return false
}

// SetFreeCapacity gets a reference to the given float32 and assigns it to the FreeCapacity field.
func (o *CreateStorage) SetFreeCapacity(v float32) {
	o.FreeCapacity = &v
}

// GetVirtualUsedCapacity returns the VirtualUsedCapacity field value if set, zero value otherwise.
func (o *CreateStorage) GetVirtualUsedCapacity() float32 {
	if o == nil || IsNil(o.VirtualUsedCapacity) {
		var ret float32
		return ret
	}
	return *o.VirtualUsedCapacity
}

// GetVirtualUsedCapacityOk returns a tuple with the VirtualUsedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetVirtualUsedCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.VirtualUsedCapacity) {
		return nil, false
	}
	return o.VirtualUsedCapacity, true
}

// HasVirtualUsedCapacity returns a boolean if a field has been set.
func (o *CreateStorage) HasVirtualUsedCapacity() bool {
	if o != nil && !IsNil(o.VirtualUsedCapacity) {
		return true
	}

	return false
}

// SetVirtualUsedCapacity gets a reference to the given float32 and assigns it to the VirtualUsedCapacity field.
func (o *CreateStorage) SetVirtualUsedCapacity(v float32) {
	o.VirtualUsedCapacity = &v
}

// GetName returns the Name field value
func (o *CreateStorage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateStorage) SetName(v string) {
	o.Name = v
}

// GetIscsiHost returns the IscsiHost field value if set, zero value otherwise.
func (o *CreateStorage) GetIscsiHost() string {
	if o == nil || IsNil(o.IscsiHost) {
		var ret string
		return ret
	}
	return *o.IscsiHost
}

// GetIscsiHostOk returns a tuple with the IscsiHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetIscsiHostOk() (*string, bool) {
	if o == nil || IsNil(o.IscsiHost) {
		return nil, false
	}
	return o.IscsiHost, true
}

// HasIscsiHost returns a boolean if a field has been set.
func (o *CreateStorage) HasIscsiHost() bool {
	if o != nil && !IsNil(o.IscsiHost) {
		return true
	}

	return false
}

// SetIscsiHost gets a reference to the given string and assigns it to the IscsiHost field.
func (o *CreateStorage) SetIscsiHost(v string) {
	o.IscsiHost = &v
}

// GetIscsiPort returns the IscsiPort field value if set, zero value otherwise.
func (o *CreateStorage) GetIscsiPort() float32 {
	if o == nil || IsNil(o.IscsiPort) {
		var ret float32
		return ret
	}
	return *o.IscsiPort
}

// GetIscsiPortOk returns a tuple with the IscsiPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetIscsiPortOk() (*float32, bool) {
	if o == nil || IsNil(o.IscsiPort) {
		return nil, false
	}
	return o.IscsiPort, true
}

// HasIscsiPort returns a boolean if a field has been set.
func (o *CreateStorage) HasIscsiPort() bool {
	if o != nil && !IsNil(o.IscsiPort) {
		return true
	}

	return false
}

// SetIscsiPort gets a reference to the given float32 and assigns it to the IscsiPort field.
func (o *CreateStorage) SetIscsiPort(v float32) {
	o.IscsiPort = &v
}

// GetManagementHost returns the ManagementHost field value
func (o *CreateStorage) GetManagementHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagementHost
}

// GetManagementHostOk returns a tuple with the ManagementHost field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetManagementHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementHost, true
}

// SetManagementHost sets field value
func (o *CreateStorage) SetManagementHost(v string) {
	o.ManagementHost = v
}

// GetUsername returns the Username field value
func (o *CreateStorage) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateStorage) SetUsername(v string) {
	o.Username = v
}

// GetPasswordEncrypted returns the PasswordEncrypted field value
func (o *CreateStorage) GetPasswordEncrypted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordEncrypted
}

// GetPasswordEncryptedOk returns a tuple with the PasswordEncrypted field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetPasswordEncryptedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordEncrypted, true
}

// SetPasswordEncrypted sets field value
func (o *CreateStorage) SetPasswordEncrypted(v string) {
	o.PasswordEncrypted = v
}

// GetInMaintenance returns the InMaintenance field value
func (o *CreateStorage) GetInMaintenance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InMaintenance
}

// GetInMaintenanceOk returns a tuple with the InMaintenance field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetInMaintenanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InMaintenance, true
}

// SetInMaintenance sets field value
func (o *CreateStorage) SetInMaintenance(v float32) {
	o.InMaintenance = v
}

// GetTargetIQN returns the TargetIQN field value if set, zero value otherwise.
func (o *CreateStorage) GetTargetIQN() string {
	if o == nil || IsNil(o.TargetIQN) {
		var ret string
		return ret
	}
	return *o.TargetIQN
}

// GetTargetIQNOk returns a tuple with the TargetIQN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetTargetIQNOk() (*string, bool) {
	if o == nil || IsNil(o.TargetIQN) {
		return nil, false
	}
	return o.TargetIQN, true
}

// HasTargetIQN returns a boolean if a field has been set.
func (o *CreateStorage) HasTargetIQN() bool {
	if o != nil && !IsNil(o.TargetIQN) {
		return true
	}

	return false
}

// SetTargetIQN gets a reference to the given string and assigns it to the TargetIQN field.
func (o *CreateStorage) SetTargetIQN(v string) {
	o.TargetIQN = &v
}

// GetIsExperimental returns the IsExperimental field value if set, zero value otherwise.
func (o *CreateStorage) GetIsExperimental() float32 {
	if o == nil || IsNil(o.IsExperimental) {
		var ret float32
		return ret
	}
	return *o.IsExperimental
}

// GetIsExperimentalOk returns a tuple with the IsExperimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetIsExperimentalOk() (*float32, bool) {
	if o == nil || IsNil(o.IsExperimental) {
		return nil, false
	}
	return o.IsExperimental, true
}

// HasIsExperimental returns a boolean if a field has been set.
func (o *CreateStorage) HasIsExperimental() bool {
	if o != nil && !IsNil(o.IsExperimental) {
		return true
	}

	return false
}

// SetIsExperimental gets a reference to the given float32 and assigns it to the IsExperimental field.
func (o *CreateStorage) SetIsExperimental(v float32) {
	o.IsExperimental = &v
}

// GetDrivePriority returns the DrivePriority field value if set, zero value otherwise.
func (o *CreateStorage) GetDrivePriority() float32 {
	if o == nil || IsNil(o.DrivePriority) {
		var ret float32
		return ret
	}
	return *o.DrivePriority
}

// GetDrivePriorityOk returns a tuple with the DrivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetDrivePriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.DrivePriority) {
		return nil, false
	}
	return o.DrivePriority, true
}

// HasDrivePriority returns a boolean if a field has been set.
func (o *CreateStorage) HasDrivePriority() bool {
	if o != nil && !IsNil(o.DrivePriority) {
		return true
	}

	return false
}

// SetDrivePriority gets a reference to the given float32 and assigns it to the DrivePriority field.
func (o *CreateStorage) SetDrivePriority(v float32) {
	o.DrivePriority = &v
}

// GetSharedDrivePriority returns the SharedDrivePriority field value if set, zero value otherwise.
func (o *CreateStorage) GetSharedDrivePriority() float32 {
	if o == nil || IsNil(o.SharedDrivePriority) {
		var ret float32
		return ret
	}
	return *o.SharedDrivePriority
}

// GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetSharedDrivePriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.SharedDrivePriority) {
		return nil, false
	}
	return o.SharedDrivePriority, true
}

// HasSharedDrivePriority returns a boolean if a field has been set.
func (o *CreateStorage) HasSharedDrivePriority() bool {
	if o != nil && !IsNil(o.SharedDrivePriority) {
		return true
	}

	return false
}

// SetSharedDrivePriority gets a reference to the given float32 and assigns it to the SharedDrivePriority field.
func (o *CreateStorage) SetSharedDrivePriority(v float32) {
	o.SharedDrivePriority = &v
}

// GetAlternateSanIPs returns the AlternateSanIPs field value if set, zero value otherwise.
func (o *CreateStorage) GetAlternateSanIPs() []string {
	if o == nil || IsNil(o.AlternateSanIPs) {
		var ret []string
		return ret
	}
	return o.AlternateSanIPs
}

// GetAlternateSanIPsOk returns a tuple with the AlternateSanIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetAlternateSanIPsOk() ([]string, bool) {
	if o == nil || IsNil(o.AlternateSanIPs) {
		return nil, false
	}
	return o.AlternateSanIPs, true
}

// HasAlternateSanIPs returns a boolean if a field has been set.
func (o *CreateStorage) HasAlternateSanIPs() bool {
	if o != nil && !IsNil(o.AlternateSanIPs) {
		return true
	}

	return false
}

// SetAlternateSanIPs gets a reference to the given []string and assigns it to the AlternateSanIPs field.
func (o *CreateStorage) SetAlternateSanIPs(v []string) {
	o.AlternateSanIPs = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateStorage) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateStorage) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateStorage) SetTags(v []string) {
	o.Tags = v
}

// GetPortGroupAllocationOrder returns the PortGroupAllocationOrder field value if set, zero value otherwise.
func (o *CreateStorage) GetPortGroupAllocationOrder() map[string]interface{} {
	if o == nil || IsNil(o.PortGroupAllocationOrder) {
		var ret map[string]interface{}
		return ret
	}
	return o.PortGroupAllocationOrder
}

// GetPortGroupAllocationOrderOk returns a tuple with the PortGroupAllocationOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetPortGroupAllocationOrderOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PortGroupAllocationOrder) {
		return map[string]interface{}{}, false
	}
	return o.PortGroupAllocationOrder, true
}

// HasPortGroupAllocationOrder returns a boolean if a field has been set.
func (o *CreateStorage) HasPortGroupAllocationOrder() bool {
	if o != nil && !IsNil(o.PortGroupAllocationOrder) {
		return true
	}

	return false
}

// SetPortGroupAllocationOrder gets a reference to the given map[string]interface{} and assigns it to the PortGroupAllocationOrder field.
func (o *CreateStorage) SetPortGroupAllocationOrder(v map[string]interface{}) {
	o.PortGroupAllocationOrder = v
}

// GetPortGroupPhysicalPorts returns the PortGroupPhysicalPorts field value if set, zero value otherwise.
func (o *CreateStorage) GetPortGroupPhysicalPorts() map[string]interface{} {
	if o == nil || IsNil(o.PortGroupPhysicalPorts) {
		var ret map[string]interface{}
		return ret
	}
	return o.PortGroupPhysicalPorts
}

// GetPortGroupPhysicalPortsOk returns a tuple with the PortGroupPhysicalPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetPortGroupPhysicalPortsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PortGroupPhysicalPorts) {
		return map[string]interface{}{}, false
	}
	return o.PortGroupPhysicalPorts, true
}

// HasPortGroupPhysicalPorts returns a boolean if a field has been set.
func (o *CreateStorage) HasPortGroupPhysicalPorts() bool {
	if o != nil && !IsNil(o.PortGroupPhysicalPorts) {
		return true
	}

	return false
}

// SetPortGroupPhysicalPorts gets a reference to the given map[string]interface{} and assigns it to the PortGroupPhysicalPorts field.
func (o *CreateStorage) SetPortGroupPhysicalPorts(v map[string]interface{}) {
	o.PortGroupPhysicalPorts = v
}

// GetDefaultIoLimitPolicy returns the DefaultIoLimitPolicy field value if set, zero value otherwise.
func (o *CreateStorage) GetDefaultIoLimitPolicy() string {
	if o == nil || IsNil(o.DefaultIoLimitPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultIoLimitPolicy
}

// GetDefaultIoLimitPolicyOk returns a tuple with the DefaultIoLimitPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetDefaultIoLimitPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultIoLimitPolicy) {
		return nil, false
	}
	return o.DefaultIoLimitPolicy, true
}

// HasDefaultIoLimitPolicy returns a boolean if a field has been set.
func (o *CreateStorage) HasDefaultIoLimitPolicy() bool {
	if o != nil && !IsNil(o.DefaultIoLimitPolicy) {
		return true
	}

	return false
}

// SetDefaultIoLimitPolicy gets a reference to the given string and assigns it to the DefaultIoLimitPolicy field.
func (o *CreateStorage) SetDefaultIoLimitPolicy(v string) {
	o.DefaultIoLimitPolicy = &v
}

// GetSubnetType returns the SubnetType field value
func (o *CreateStorage) GetSubnetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetType
}

// GetSubnetTypeOk returns a tuple with the SubnetType field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetSubnetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetType, true
}

// SetSubnetType sets field value
func (o *CreateStorage) SetSubnetType(v string) {
	o.SubnetType = v
}

// GetArrayId returns the ArrayId field value if set, zero value otherwise.
func (o *CreateStorage) GetArrayId() string {
	if o == nil || IsNil(o.ArrayId) {
		var ret string
		return ret
	}
	return *o.ArrayId
}

// GetArrayIdOk returns a tuple with the ArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetArrayIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArrayId) {
		return nil, false
	}
	return o.ArrayId, true
}

// HasArrayId returns a boolean if a field has been set.
func (o *CreateStorage) HasArrayId() bool {
	if o != nil && !IsNil(o.ArrayId) {
		return true
	}

	return false
}

// SetArrayId gets a reference to the given string and assigns it to the ArrayId field.
func (o *CreateStorage) SetArrayId(v string) {
	o.ArrayId = &v
}

// GetDirectorId returns the DirectorId field value if set, zero value otherwise.
func (o *CreateStorage) GetDirectorId() string {
	if o == nil || IsNil(o.DirectorId) {
		var ret string
		return ret
	}
	return *o.DirectorId
}

// GetDirectorIdOk returns a tuple with the DirectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetDirectorIdOk() (*string, bool) {
	if o == nil || IsNil(o.DirectorId) {
		return nil, false
	}
	return o.DirectorId, true
}

// HasDirectorId returns a boolean if a field has been set.
func (o *CreateStorage) HasDirectorId() bool {
	if o != nil && !IsNil(o.DirectorId) {
		return true
	}

	return false
}

// SetDirectorId gets a reference to the given string and assigns it to the DirectorId field.
func (o *CreateStorage) SetDirectorId(v string) {
	o.DirectorId = &v
}

// GetS3Hostname returns the S3Hostname field value if set, zero value otherwise.
func (o *CreateStorage) GetS3Hostname() string {
	if o == nil || IsNil(o.S3Hostname) {
		var ret string
		return ret
	}
	return *o.S3Hostname
}

// GetS3HostnameOk returns a tuple with the S3Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetS3HostnameOk() (*string, bool) {
	if o == nil || IsNil(o.S3Hostname) {
		return nil, false
	}
	return o.S3Hostname, true
}

// HasS3Hostname returns a boolean if a field has been set.
func (o *CreateStorage) HasS3Hostname() bool {
	if o != nil && !IsNil(o.S3Hostname) {
		return true
	}

	return false
}

// SetS3Hostname gets a reference to the given string and assigns it to the S3Hostname field.
func (o *CreateStorage) SetS3Hostname(v string) {
	o.S3Hostname = &v
}

// GetS3Port returns the S3Port field value if set, zero value otherwise.
func (o *CreateStorage) GetS3Port() string {
	if o == nil || IsNil(o.S3Port) {
		var ret string
		return ret
	}
	return *o.S3Port
}

// GetS3PortOk returns a tuple with the S3Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetS3PortOk() (*string, bool) {
	if o == nil || IsNil(o.S3Port) {
		return nil, false
	}
	return o.S3Port, true
}

// HasS3Port returns a boolean if a field has been set.
func (o *CreateStorage) HasS3Port() bool {
	if o != nil && !IsNil(o.S3Port) {
		return true
	}

	return false
}

// SetS3Port gets a reference to the given string and assigns it to the S3Port field.
func (o *CreateStorage) SetS3Port(v string) {
	o.S3Port = &v
}

// GetJobInfo returns the JobInfo field value if set, zero value otherwise.
func (o *CreateStorage) GetJobInfo() JobInfo {
	if o == nil || IsNil(o.JobInfo) {
		var ret JobInfo
		return ret
	}
	return *o.JobInfo
}

// GetJobInfoOk returns a tuple with the JobInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetJobInfoOk() (*JobInfo, bool) {
	if o == nil || IsNil(o.JobInfo) {
		return nil, false
	}
	return o.JobInfo, true
}

// HasJobInfo returns a boolean if a field has been set.
func (o *CreateStorage) HasJobInfo() bool {
	if o != nil && !IsNil(o.JobInfo) {
		return true
	}

	return false
}

// SetJobInfo gets a reference to the given JobInfo and assigns it to the JobInfo field.
func (o *CreateStorage) SetJobInfo(v JobInfo) {
	o.JobInfo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateStorage) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateStorage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *CreateStorage) SetLinks(v []Link) {
	o.Links = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateStorage) GetOptions() UpdateStorageOptions {
	if o == nil || IsNil(o.Options) {
		var ret UpdateStorageOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetOptionsOk() (*UpdateStorageOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateStorage) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given UpdateStorageOptions and assigns it to the Options field.
func (o *CreateStorage) SetOptions(v UpdateStorageOptions) {
	o.Options = &v
}

// GetPassword returns the Password field value
func (o *CreateStorage) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateStorage) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateStorage) SetPassword(v string) {
	o.Password = v
}

func (o CreateStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	toSerialize["siteId"] = o.SiteId
	toSerialize["storageDriver"] = o.StorageDriver
	toSerialize["storageTechnology"] = o.StorageTechnology
	toSerialize["storageType"] = o.StorageType
	toSerialize["status"] = o.Status
	if !IsNil(o.TotalCapacity) {
		toSerialize["totalCapacity"] = o.TotalCapacity
	}
	if !IsNil(o.UsableCapacity) {
		toSerialize["usableCapacity"] = o.UsableCapacity
	}
	if !IsNil(o.FreeCapacity) {
		toSerialize["freeCapacity"] = o.FreeCapacity
	}
	if !IsNil(o.VirtualUsedCapacity) {
		toSerialize["virtualUsedCapacity"] = o.VirtualUsedCapacity
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.IscsiHost) {
		toSerialize["iscsiHost"] = o.IscsiHost
	}
	if !IsNil(o.IscsiPort) {
		toSerialize["iscsiPort"] = o.IscsiPort
	}
	toSerialize["managementHost"] = o.ManagementHost
	toSerialize["username"] = o.Username
	toSerialize["passwordEncrypted"] = o.PasswordEncrypted
	toSerialize["inMaintenance"] = o.InMaintenance
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
	if !IsNil(o.DefaultIoLimitPolicy) {
		toSerialize["defaultIoLimitPolicy"] = o.DefaultIoLimitPolicy
	}
	toSerialize["subnetType"] = o.SubnetType
	if !IsNil(o.ArrayId) {
		toSerialize["arrayId"] = o.ArrayId
	}
	if !IsNil(o.DirectorId) {
		toSerialize["directorId"] = o.DirectorId
	}
	if !IsNil(o.S3Hostname) {
		toSerialize["s3Hostname"] = o.S3Hostname
	}
	if !IsNil(o.S3Port) {
		toSerialize["s3Port"] = o.S3Port
	}
	if !IsNil(o.JobInfo) {
		toSerialize["jobInfo"] = o.JobInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateStorage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"siteId",
		"storageDriver",
		"storageTechnology",
		"storageType",
		"status",
		"name",
		"managementHost",
		"username",
		"passwordEncrypted",
		"inMaintenance",
		"subnetType",
		"password",
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

	varCreateStorage := _CreateStorage{}

	err = json.Unmarshal(data, &varCreateStorage)

	if err != nil {
		return err
	}

	*o = CreateStorage(varCreateStorage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userId")
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "storageDriver")
		delete(additionalProperties, "storageTechnology")
		delete(additionalProperties, "storageType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "totalCapacity")
		delete(additionalProperties, "usableCapacity")
		delete(additionalProperties, "freeCapacity")
		delete(additionalProperties, "virtualUsedCapacity")
		delete(additionalProperties, "name")
		delete(additionalProperties, "iscsiHost")
		delete(additionalProperties, "iscsiPort")
		delete(additionalProperties, "managementHost")
		delete(additionalProperties, "username")
		delete(additionalProperties, "passwordEncrypted")
		delete(additionalProperties, "inMaintenance")
		delete(additionalProperties, "targetIQN")
		delete(additionalProperties, "isExperimental")
		delete(additionalProperties, "drivePriority")
		delete(additionalProperties, "sharedDrivePriority")
		delete(additionalProperties, "alternateSanIPs")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "portGroupAllocationOrder")
		delete(additionalProperties, "portGroupPhysicalPorts")
		delete(additionalProperties, "defaultIoLimitPolicy")
		delete(additionalProperties, "subnetType")
		delete(additionalProperties, "arrayId")
		delete(additionalProperties, "directorId")
		delete(additionalProperties, "s3Hostname")
		delete(additionalProperties, "s3Port")
		delete(additionalProperties, "jobInfo")
		delete(additionalProperties, "links")
		delete(additionalProperties, "options")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateStorage struct {
	value *CreateStorage
	isSet bool
}

func (v NullableCreateStorage) Get() *CreateStorage {
	return v.value
}

func (v *NullableCreateStorage) Set(val *CreateStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStorage(val *CreateStorage) *NullableCreateStorage {
	return &NullableCreateStorage{value: val, isSet: true}
}

func (v NullableCreateStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



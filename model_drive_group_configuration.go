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

// checks if the DriveGroupConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DriveGroupConfiguration{}

// DriveGroupConfiguration struct for DriveGroupConfiguration
type DriveGroupConfiguration struct {
	// Revision of the Drive Group Configuration
	Revision float32 `json:"revision"`
	// Label of the Drive Group.
	Label string `json:"label"`
	// Infrastructure id of the Drive Group
	InfrastructureId float32 `json:"infrastructureId"`
	// Template Id
	TemplateId *float32 `json:"templateId,omitempty"`
	// Default disk size in MB for new Drives in the Drive Group
	DriveSizeMbDefault float32 `json:"driveSizeMbDefault"`
	ServerInstanceGroupId *float32 `json:"serverInstanceGroupId,omitempty"`
	ContainerArrayId *float32 `json:"containerArrayId,omitempty"`
	// Flag to determine whether the Drive Group should be expanded with a Server Instance Group by adding one drive for each instance
	ExpandWithServerInstanceGroup float32 `json:"expandWithServerInstanceGroup"`
	// The IO limit policy of the Drive Group.
	IoLimitPolicy *string `json:"ioLimitPolicy,omitempty"`
	// Service status of the Drive Group
	StorageType string `json:"storageType"`
	// Filesystem information of the Drive Group.
	FilesystemInfo map[string]interface{} `json:"filesystemInfo,omitempty"`
	// Subdomain of the Drive Group.
	Subdomain *string `json:"subdomain,omitempty"`
	// Timestamp of the Drive Group last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Id of the DNS subdomain for the Drive Group.
	DnsSubdomainChangeId *float32 `json:"dnsSubdomainChangeId,omitempty"`
	// Deploy type of the Drive Group
	DeployType string `json:"deployType"`
	// Deploy status of the Drive Group
	DeployStatus string `json:"deployStatus"`
	// Id of the deployment for the Drive Group.
	InfrastructureDeployId *float32 `json:"infrastructureDeployId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DriveGroupConfiguration DriveGroupConfiguration

// NewDriveGroupConfiguration instantiates a new DriveGroupConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriveGroupConfiguration(revision float32, label string, infrastructureId float32, driveSizeMbDefault float32, expandWithServerInstanceGroup float32, storageType string, updatedTimestamp string, deployType string, deployStatus string) *DriveGroupConfiguration {
	this := DriveGroupConfiguration{}
	this.Revision = revision
	this.Label = label
	this.InfrastructureId = infrastructureId
	this.DriveSizeMbDefault = driveSizeMbDefault
	this.ExpandWithServerInstanceGroup = expandWithServerInstanceGroup
	this.StorageType = storageType
	this.UpdatedTimestamp = updatedTimestamp
	this.DeployType = deployType
	this.DeployStatus = deployStatus
	return &this
}

// NewDriveGroupConfigurationWithDefaults instantiates a new DriveGroupConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriveGroupConfigurationWithDefaults() *DriveGroupConfiguration {
	this := DriveGroupConfiguration{}
	var storageType string = "iscsi_ssd"
	this.StorageType = storageType
	var deployType string = "create"
	this.DeployType = deployType
	var deployStatus string = "not_started"
	this.DeployStatus = deployStatus
	return &this
}

// GetRevision returns the Revision field value
func (o *DriveGroupConfiguration) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *DriveGroupConfiguration) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value
func (o *DriveGroupConfiguration) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *DriveGroupConfiguration) SetLabel(v string) {
	o.Label = v
}

// GetInfrastructureId returns the InfrastructureId field value
func (o *DriveGroupConfiguration) GetInfrastructureId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructureId, true
}

// SetInfrastructureId sets field value
func (o *DriveGroupConfiguration) SetInfrastructureId(v float32) {
	o.InfrastructureId = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetTemplateId() float32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret float32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetTemplateIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given float32 and assigns it to the TemplateId field.
func (o *DriveGroupConfiguration) SetTemplateId(v float32) {
	o.TemplateId = &v
}

// GetDriveSizeMbDefault returns the DriveSizeMbDefault field value
func (o *DriveGroupConfiguration) GetDriveSizeMbDefault() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DriveSizeMbDefault
}

// GetDriveSizeMbDefaultOk returns a tuple with the DriveSizeMbDefault field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetDriveSizeMbDefaultOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DriveSizeMbDefault, true
}

// SetDriveSizeMbDefault sets field value
func (o *DriveGroupConfiguration) SetDriveSizeMbDefault(v float32) {
	o.DriveSizeMbDefault = v
}

// GetServerInstanceGroupId returns the ServerInstanceGroupId field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetServerInstanceGroupId() float32 {
	if o == nil || IsNil(o.ServerInstanceGroupId) {
		var ret float32
		return ret
	}
	return *o.ServerInstanceGroupId
}

// GetServerInstanceGroupIdOk returns a tuple with the ServerInstanceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetServerInstanceGroupIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerInstanceGroupId) {
		return nil, false
	}
	return o.ServerInstanceGroupId, true
}

// HasServerInstanceGroupId returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasServerInstanceGroupId() bool {
	if o != nil && !IsNil(o.ServerInstanceGroupId) {
		return true
	}

	return false
}

// SetServerInstanceGroupId gets a reference to the given float32 and assigns it to the ServerInstanceGroupId field.
func (o *DriveGroupConfiguration) SetServerInstanceGroupId(v float32) {
	o.ServerInstanceGroupId = &v
}

// GetContainerArrayId returns the ContainerArrayId field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetContainerArrayId() float32 {
	if o == nil || IsNil(o.ContainerArrayId) {
		var ret float32
		return ret
	}
	return *o.ContainerArrayId
}

// GetContainerArrayIdOk returns a tuple with the ContainerArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetContainerArrayIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ContainerArrayId) {
		return nil, false
	}
	return o.ContainerArrayId, true
}

// HasContainerArrayId returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasContainerArrayId() bool {
	if o != nil && !IsNil(o.ContainerArrayId) {
		return true
	}

	return false
}

// SetContainerArrayId gets a reference to the given float32 and assigns it to the ContainerArrayId field.
func (o *DriveGroupConfiguration) SetContainerArrayId(v float32) {
	o.ContainerArrayId = &v
}

// GetExpandWithServerInstanceGroup returns the ExpandWithServerInstanceGroup field value
func (o *DriveGroupConfiguration) GetExpandWithServerInstanceGroup() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpandWithServerInstanceGroup
}

// GetExpandWithServerInstanceGroupOk returns a tuple with the ExpandWithServerInstanceGroup field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetExpandWithServerInstanceGroupOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpandWithServerInstanceGroup, true
}

// SetExpandWithServerInstanceGroup sets field value
func (o *DriveGroupConfiguration) SetExpandWithServerInstanceGroup(v float32) {
	o.ExpandWithServerInstanceGroup = v
}

// GetIoLimitPolicy returns the IoLimitPolicy field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetIoLimitPolicy() string {
	if o == nil || IsNil(o.IoLimitPolicy) {
		var ret string
		return ret
	}
	return *o.IoLimitPolicy
}

// GetIoLimitPolicyOk returns a tuple with the IoLimitPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetIoLimitPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.IoLimitPolicy) {
		return nil, false
	}
	return o.IoLimitPolicy, true
}

// HasIoLimitPolicy returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasIoLimitPolicy() bool {
	if o != nil && !IsNil(o.IoLimitPolicy) {
		return true
	}

	return false
}

// SetIoLimitPolicy gets a reference to the given string and assigns it to the IoLimitPolicy field.
func (o *DriveGroupConfiguration) SetIoLimitPolicy(v string) {
	o.IoLimitPolicy = &v
}

// GetStorageType returns the StorageType field value
func (o *DriveGroupConfiguration) GetStorageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *DriveGroupConfiguration) SetStorageType(v string) {
	o.StorageType = v
}

// GetFilesystemInfo returns the FilesystemInfo field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetFilesystemInfo() map[string]interface{} {
	if o == nil || IsNil(o.FilesystemInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.FilesystemInfo
}

// GetFilesystemInfoOk returns a tuple with the FilesystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetFilesystemInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FilesystemInfo) {
		return map[string]interface{}{}, false
	}
	return o.FilesystemInfo, true
}

// HasFilesystemInfo returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasFilesystemInfo() bool {
	if o != nil && !IsNil(o.FilesystemInfo) {
		return true
	}

	return false
}

// SetFilesystemInfo gets a reference to the given map[string]interface{} and assigns it to the FilesystemInfo field.
func (o *DriveGroupConfiguration) SetFilesystemInfo(v map[string]interface{}) {
	o.FilesystemInfo = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DriveGroupConfiguration) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *DriveGroupConfiguration) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *DriveGroupConfiguration) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetDnsSubdomainChangeId() float32 {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainChangeId
}

// GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		return nil, false
	}
	return o.DnsSubdomainChangeId, true
}

// HasDnsSubdomainChangeId returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasDnsSubdomainChangeId() bool {
	if o != nil && !IsNil(o.DnsSubdomainChangeId) {
		return true
	}

	return false
}

// SetDnsSubdomainChangeId gets a reference to the given float32 and assigns it to the DnsSubdomainChangeId field.
func (o *DriveGroupConfiguration) SetDnsSubdomainChangeId(v float32) {
	o.DnsSubdomainChangeId = &v
}

// GetDeployType returns the DeployType field value
func (o *DriveGroupConfiguration) GetDeployType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetDeployTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployType, true
}

// SetDeployType sets field value
func (o *DriveGroupConfiguration) SetDeployType(v string) {
	o.DeployType = v
}

// GetDeployStatus returns the DeployStatus field value
func (o *DriveGroupConfiguration) GetDeployStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetDeployStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployStatus, true
}

// SetDeployStatus sets field value
func (o *DriveGroupConfiguration) SetDeployStatus(v string) {
	o.DeployStatus = v
}

// GetInfrastructureDeployId returns the InfrastructureDeployId field value if set, zero value otherwise.
func (o *DriveGroupConfiguration) GetInfrastructureDeployId() float32 {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		var ret float32
		return ret
	}
	return *o.InfrastructureDeployId
}

// GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveGroupConfiguration) GetInfrastructureDeployIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		return nil, false
	}
	return o.InfrastructureDeployId, true
}

// HasInfrastructureDeployId returns a boolean if a field has been set.
func (o *DriveGroupConfiguration) HasInfrastructureDeployId() bool {
	if o != nil && !IsNil(o.InfrastructureDeployId) {
		return true
	}

	return false
}

// SetInfrastructureDeployId gets a reference to the given float32 and assigns it to the InfrastructureDeployId field.
func (o *DriveGroupConfiguration) SetInfrastructureDeployId(v float32) {
	o.InfrastructureDeployId = &v
}

func (o DriveGroupConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DriveGroupConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["label"] = o.Label
	toSerialize["infrastructureId"] = o.InfrastructureId
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	toSerialize["driveSizeMbDefault"] = o.DriveSizeMbDefault
	if !IsNil(o.ServerInstanceGroupId) {
		toSerialize["serverInstanceGroupId"] = o.ServerInstanceGroupId
	}
	if !IsNil(o.ContainerArrayId) {
		toSerialize["containerArrayId"] = o.ContainerArrayId
	}
	toSerialize["expandWithServerInstanceGroup"] = o.ExpandWithServerInstanceGroup
	if !IsNil(o.IoLimitPolicy) {
		toSerialize["ioLimitPolicy"] = o.IoLimitPolicy
	}
	toSerialize["storageType"] = o.StorageType
	if !IsNil(o.FilesystemInfo) {
		toSerialize["filesystemInfo"] = o.FilesystemInfo
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
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

func (o *DriveGroupConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"label",
		"infrastructureId",
		"driveSizeMbDefault",
		"expandWithServerInstanceGroup",
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

	varDriveGroupConfiguration := _DriveGroupConfiguration{}

	err = json.Unmarshal(data, &varDriveGroupConfiguration)

	if err != nil {
		return err
	}

	*o = DriveGroupConfiguration(varDriveGroupConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "driveSizeMbDefault")
		delete(additionalProperties, "serverInstanceGroupId")
		delete(additionalProperties, "containerArrayId")
		delete(additionalProperties, "expandWithServerInstanceGroup")
		delete(additionalProperties, "ioLimitPolicy")
		delete(additionalProperties, "storageType")
		delete(additionalProperties, "filesystemInfo")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "dnsSubdomainChangeId")
		delete(additionalProperties, "deployType")
		delete(additionalProperties, "deployStatus")
		delete(additionalProperties, "infrastructureDeployId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDriveGroupConfiguration struct {
	value *DriveGroupConfiguration
	isSet bool
}

func (v NullableDriveGroupConfiguration) Get() *DriveGroupConfiguration {
	return v.value
}

func (v *NullableDriveGroupConfiguration) Set(val *DriveGroupConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDriveGroupConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDriveGroupConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriveGroupConfiguration(val *DriveGroupConfiguration) *NullableDriveGroupConfiguration {
	return &NullableDriveGroupConfiguration{value: val, isSet: true}
}

func (v NullableDriveGroupConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriveGroupConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



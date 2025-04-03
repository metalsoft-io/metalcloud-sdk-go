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

// checks if the VM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VM{}

// VM struct for VM
type VM struct {
	// VM ID
	Id float32 `json:"id"`
	// Name of the VM
	Name string `json:"name"`
	// Id of the site for the VM
	SiteId float32 `json:"siteId"`
	// Datacenter of the VM
	DatacenterName string `json:"datacenterName"`
	// ID of the infrastructure where this VM is deployed
	InfrastructureId float32 `json:"infrastructureId"`
	// ID of the user that owns this VM
	UserId float32 `json:"userId"`
	// Email of the user that owns this VM
	UserEmail string `json:"userEmail"`
	// ID of the instance where this VM is deployed
	InstanceId float32 `json:"instanceId"`
	// The id of the VM Instance. This is a number.
	VmInstanceId float32 `json:"vmInstanceId"`
	// Name of the host
	Host string `json:"host"`
	// List of hosts
	Hosts []string `json:"hosts"`
	// Number of CPU cores for the VM
	CpuCores float32 `json:"cpuCores"`
	// RAM in GB for the VM
	RamGB float32 `json:"ramGB"`
	// Disk size in GB for the VM
	DiskSizeGB float32 `json:"diskSizeGB"`
	// The id of the VM Type. This is a number.
	TypeId float32 `json:"typeId"`
	// The id of the VM Pool. This is a number.
	PoolId float32 `json:"poolId"`
	// The administration state of the VM.
	AdministrationState string `json:"administrationState"`
	// VM comments.
	Comments *float32 `json:"comments,omitempty"`
	// The power state of the VM.
	PowerState string `json:"powerState"`
	// Timestamp when the VM power state was last updated.
	PowerStateLastUpdatedTimestamp string `json:"powerStateLastUpdatedTimestamp"`
	// Timestamp when the VM was created
	CreatedTimestamp string `json:"createdTimestamp"`
	// Timestamp when the VM was allocated
	AllocationTimestamp string `json:"allocationTimestamp"`
	// Tags for the VM. This is a JSON object.
	Tags []string `json:"tags,omitempty"`
	// The disks of the VM.
	Disks []VMDisk `json:"disks"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VM VM

// NewVM instantiates a new VM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVM(id float32, name string, siteId float32, datacenterName string, infrastructureId float32, userId float32, userEmail string, instanceId float32, vmInstanceId float32, host string, hosts []string, cpuCores float32, ramGB float32, diskSizeGB float32, typeId float32, poolId float32, administrationState string, powerState string, powerStateLastUpdatedTimestamp string, createdTimestamp string, allocationTimestamp string, disks []VMDisk) *VM {
	this := VM{}
	this.Id = id
	this.Name = name
	this.SiteId = siteId
	this.DatacenterName = datacenterName
	this.InfrastructureId = infrastructureId
	this.UserId = userId
	this.UserEmail = userEmail
	this.InstanceId = instanceId
	this.VmInstanceId = vmInstanceId
	this.Host = host
	this.Hosts = hosts
	this.CpuCores = cpuCores
	this.RamGB = ramGB
	this.DiskSizeGB = diskSizeGB
	this.TypeId = typeId
	this.PoolId = poolId
	this.AdministrationState = administrationState
	this.PowerState = powerState
	this.PowerStateLastUpdatedTimestamp = powerStateLastUpdatedTimestamp
	this.CreatedTimestamp = createdTimestamp
	this.AllocationTimestamp = allocationTimestamp
	this.Disks = disks
	return &this
}

// NewVMWithDefaults instantiates a new VM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMWithDefaults() *VM {
	this := VM{}
	return &this
}

// GetId returns the Id field value
func (o *VM) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VM) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VM) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VM) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VM) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VM) SetName(v string) {
	o.Name = v
}

// GetSiteId returns the SiteId field value
func (o *VM) GetSiteId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *VM) GetSiteIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *VM) SetSiteId(v float32) {
	o.SiteId = v
}

// GetDatacenterName returns the DatacenterName field value
func (o *VM) GetDatacenterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatacenterName
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value
// and a boolean to check if the value has been set.
func (o *VM) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatacenterName, true
}

// SetDatacenterName sets field value
func (o *VM) SetDatacenterName(v string) {
	o.DatacenterName = v
}

// GetInfrastructureId returns the InfrastructureId field value
func (o *VM) GetInfrastructureId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value
// and a boolean to check if the value has been set.
func (o *VM) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfrastructureId, true
}

// SetInfrastructureId sets field value
func (o *VM) SetInfrastructureId(v float32) {
	o.InfrastructureId = v
}

// GetUserId returns the UserId field value
func (o *VM) GetUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *VM) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *VM) SetUserId(v float32) {
	o.UserId = v
}

// GetUserEmail returns the UserEmail field value
func (o *VM) GetUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value
// and a boolean to check if the value has been set.
func (o *VM) GetUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserEmail, true
}

// SetUserEmail sets field value
func (o *VM) SetUserEmail(v string) {
	o.UserEmail = v
}

// GetInstanceId returns the InstanceId field value
func (o *VM) GetInstanceId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *VM) GetInstanceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *VM) SetInstanceId(v float32) {
	o.InstanceId = v
}

// GetVmInstanceId returns the VmInstanceId field value
func (o *VM) GetVmInstanceId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VmInstanceId
}

// GetVmInstanceIdOk returns a tuple with the VmInstanceId field value
// and a boolean to check if the value has been set.
func (o *VM) GetVmInstanceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmInstanceId, true
}

// SetVmInstanceId sets field value
func (o *VM) SetVmInstanceId(v float32) {
	o.VmInstanceId = v
}

// GetHost returns the Host field value
func (o *VM) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *VM) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *VM) SetHost(v string) {
	o.Host = v
}

// GetHosts returns the Hosts field value
func (o *VM) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *VM) GetHostsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *VM) SetHosts(v []string) {
	o.Hosts = v
}

// GetCpuCores returns the CpuCores field value
func (o *VM) GetCpuCores() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value
// and a boolean to check if the value has been set.
func (o *VM) GetCpuCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuCores, true
}

// SetCpuCores sets field value
func (o *VM) SetCpuCores(v float32) {
	o.CpuCores = v
}

// GetRamGB returns the RamGB field value
func (o *VM) GetRamGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamGB
}

// GetRamGBOk returns a tuple with the RamGB field value
// and a boolean to check if the value has been set.
func (o *VM) GetRamGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamGB, true
}

// SetRamGB sets field value
func (o *VM) SetRamGB(v float32) {
	o.RamGB = v
}

// GetDiskSizeGB returns the DiskSizeGB field value
func (o *VM) GetDiskSizeGB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value
// and a boolean to check if the value has been set.
func (o *VM) GetDiskSizeGBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskSizeGB, true
}

// SetDiskSizeGB sets field value
func (o *VM) SetDiskSizeGB(v float32) {
	o.DiskSizeGB = v
}

// GetTypeId returns the TypeId field value
func (o *VM) GetTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *VM) GetTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *VM) SetTypeId(v float32) {
	o.TypeId = v
}

// GetPoolId returns the PoolId field value
func (o *VM) GetPoolId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *VM) GetPoolIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *VM) SetPoolId(v float32) {
	o.PoolId = v
}

// GetAdministrationState returns the AdministrationState field value
func (o *VM) GetAdministrationState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdministrationState
}

// GetAdministrationStateOk returns a tuple with the AdministrationState field value
// and a boolean to check if the value has been set.
func (o *VM) GetAdministrationStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdministrationState, true
}

// SetAdministrationState sets field value
func (o *VM) SetAdministrationState(v string) {
	o.AdministrationState = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VM) GetComments() float32 {
	if o == nil || IsNil(o.Comments) {
		var ret float32
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VM) GetCommentsOk() (*float32, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VM) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given float32 and assigns it to the Comments field.
func (o *VM) SetComments(v float32) {
	o.Comments = &v
}

// GetPowerState returns the PowerState field value
func (o *VM) GetPowerState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value
// and a boolean to check if the value has been set.
func (o *VM) GetPowerStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerState, true
}

// SetPowerState sets field value
func (o *VM) SetPowerState(v string) {
	o.PowerState = v
}

// GetPowerStateLastUpdatedTimestamp returns the PowerStateLastUpdatedTimestamp field value
func (o *VM) GetPowerStateLastUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerStateLastUpdatedTimestamp
}

// GetPowerStateLastUpdatedTimestampOk returns a tuple with the PowerStateLastUpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VM) GetPowerStateLastUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerStateLastUpdatedTimestamp, true
}

// SetPowerStateLastUpdatedTimestamp sets field value
func (o *VM) SetPowerStateLastUpdatedTimestamp(v string) {
	o.PowerStateLastUpdatedTimestamp = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *VM) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VM) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *VM) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetAllocationTimestamp returns the AllocationTimestamp field value
func (o *VM) GetAllocationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AllocationTimestamp
}

// GetAllocationTimestampOk returns a tuple with the AllocationTimestamp field value
// and a boolean to check if the value has been set.
func (o *VM) GetAllocationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocationTimestamp, true
}

// SetAllocationTimestamp sets field value
func (o *VM) SetAllocationTimestamp(v string) {
	o.AllocationTimestamp = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VM) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VM) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VM) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VM) SetTags(v []string) {
	o.Tags = v
}

// GetDisks returns the Disks field value
func (o *VM) GetDisks() []VMDisk {
	if o == nil {
		var ret []VMDisk
		return ret
	}

	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value
// and a boolean to check if the value has been set.
func (o *VM) GetDisksOk() ([]VMDisk, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disks, true
}

// SetDisks sets field value
func (o *VM) SetDisks(v []VMDisk) {
	o.Disks = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VM) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VM) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VM) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *VM) SetLinks(v []Link) {
	o.Links = v
}

func (o VM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["siteId"] = o.SiteId
	toSerialize["datacenterName"] = o.DatacenterName
	toSerialize["infrastructureId"] = o.InfrastructureId
	toSerialize["userId"] = o.UserId
	toSerialize["userEmail"] = o.UserEmail
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["vmInstanceId"] = o.VmInstanceId
	toSerialize["host"] = o.Host
	toSerialize["hosts"] = o.Hosts
	toSerialize["cpuCores"] = o.CpuCores
	toSerialize["ramGB"] = o.RamGB
	toSerialize["diskSizeGB"] = o.DiskSizeGB
	toSerialize["typeId"] = o.TypeId
	toSerialize["poolId"] = o.PoolId
	toSerialize["administrationState"] = o.AdministrationState
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["powerState"] = o.PowerState
	toSerialize["powerStateLastUpdatedTimestamp"] = o.PowerStateLastUpdatedTimestamp
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	toSerialize["allocationTimestamp"] = o.AllocationTimestamp
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["disks"] = o.Disks
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"siteId",
		"datacenterName",
		"infrastructureId",
		"userId",
		"userEmail",
		"instanceId",
		"vmInstanceId",
		"host",
		"hosts",
		"cpuCores",
		"ramGB",
		"diskSizeGB",
		"typeId",
		"poolId",
		"administrationState",
		"powerState",
		"powerStateLastUpdatedTimestamp",
		"createdTimestamp",
		"allocationTimestamp",
		"disks",
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

	varVM := _VM{}

	err = json.Unmarshal(data, &varVM)

	if err != nil {
		return err
	}

	*o = VM(varVM)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "datacenterName")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "userEmail")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "vmInstanceId")
		delete(additionalProperties, "host")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "cpuCores")
		delete(additionalProperties, "ramGB")
		delete(additionalProperties, "diskSizeGB")
		delete(additionalProperties, "typeId")
		delete(additionalProperties, "poolId")
		delete(additionalProperties, "administrationState")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "powerState")
		delete(additionalProperties, "powerStateLastUpdatedTimestamp")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "allocationTimestamp")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "disks")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVM struct {
	value *VM
	isSet bool
}

func (v NullableVM) Get() *VM {
	return v.value
}

func (v *NullableVM) Set(val *VM) {
	v.value = val
	v.isSet = true
}

func (v NullableVM) IsSet() bool {
	return v.isSet
}

func (v *NullableVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVM(val *VM) *NullableVM {
	return &NullableVM{value: val, isSet: true}
}

func (v NullableVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



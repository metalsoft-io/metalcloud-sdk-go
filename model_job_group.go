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

// checks if the JobGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobGroup{}

// JobGroup struct for JobGroup
type JobGroup struct {
	// Group Id
	Id int32 `json:"id"`
	// Group type
	Type string `json:"type"`
	// Group context parameters
	Context map[string]interface{} `json:"context,omitempty"`
	// Group description
	Description string `json:"description"`
	// Group created timestamp
	CreatedTimestamp string `json:"createdTimestamp"`
	// Group finished timestamp
	FinishedTimestamp *string `json:"finishedTimestamp,omitempty"`
	// Group parameters
	Params map[string]interface{} `json:"params,omitempty"`
	// Group archived status
	Archived *int32 `json:"archived,omitempty"`
	// Infrastructure Id
	InfrastructureId *int32 `json:"infrastructureId,omitempty"`
	// Drive Id
	DriveId *int32 `json:"driveId,omitempty"`
	// Server Id
	ServerId *int32 `json:"serverId,omitempty"`
	// Network device Id
	NetworkDeviceId *int32 `json:"networkDeviceId,omitempty"`
	// VM Pool Id
	VmPoolId *int32 `json:"vmPoolId,omitempty"`
	// Storage Pool Id
	StorageId *int32 `json:"storageId,omitempty"`
	// Links to other resources
	Links map[string]interface{} `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _JobGroup JobGroup

// NewJobGroup instantiates a new JobGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobGroup(id int32, type_ string, description string, createdTimestamp string, links map[string]interface{}) *JobGroup {
	this := JobGroup{}
	this.Id = id
	this.Type = type_
	this.Description = description
	this.CreatedTimestamp = createdTimestamp
	this.Links = links
	return &this
}

// NewJobGroupWithDefaults instantiates a new JobGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobGroupWithDefaults() *JobGroup {
	this := JobGroup{}
	return &this
}

// GetId returns the Id field value
func (o *JobGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobGroup) SetId(v int32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *JobGroup) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobGroup) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobGroup) SetType(v string) {
	o.Type = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *JobGroup) GetContext() map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *JobGroup) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *JobGroup) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetDescription returns the Description field value
func (o *JobGroup) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *JobGroup) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *JobGroup) SetDescription(v string) {
	o.Description = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *JobGroup) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *JobGroup) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *JobGroup) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetFinishedTimestamp returns the FinishedTimestamp field value if set, zero value otherwise.
func (o *JobGroup) GetFinishedTimestamp() string {
	if o == nil || IsNil(o.FinishedTimestamp) {
		var ret string
		return ret
	}
	return *o.FinishedTimestamp
}

// GetFinishedTimestampOk returns a tuple with the FinishedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetFinishedTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.FinishedTimestamp) {
		return nil, false
	}
	return o.FinishedTimestamp, true
}

// HasFinishedTimestamp returns a boolean if a field has been set.
func (o *JobGroup) HasFinishedTimestamp() bool {
	if o != nil && !IsNil(o.FinishedTimestamp) {
		return true
	}

	return false
}

// SetFinishedTimestamp gets a reference to the given string and assigns it to the FinishedTimestamp field.
func (o *JobGroup) SetFinishedTimestamp(v string) {
	o.FinishedTimestamp = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *JobGroup) GetParams() map[string]interface{} {
	if o == nil || IsNil(o.Params) {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *JobGroup) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *JobGroup) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *JobGroup) GetArchived() int32 {
	if o == nil || IsNil(o.Archived) {
		var ret int32
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetArchivedOk() (*int32, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *JobGroup) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given int32 and assigns it to the Archived field.
func (o *JobGroup) SetArchived(v int32) {
	o.Archived = &v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *JobGroup) GetInfrastructureId() int32 {
	if o == nil || IsNil(o.InfrastructureId) {
		var ret int32
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetInfrastructureIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InfrastructureId) {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *JobGroup) HasInfrastructureId() bool {
	if o != nil && !IsNil(o.InfrastructureId) {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given int32 and assigns it to the InfrastructureId field.
func (o *JobGroup) SetInfrastructureId(v int32) {
	o.InfrastructureId = &v
}

// GetDriveId returns the DriveId field value if set, zero value otherwise.
func (o *JobGroup) GetDriveId() int32 {
	if o == nil || IsNil(o.DriveId) {
		var ret int32
		return ret
	}
	return *o.DriveId
}

// GetDriveIdOk returns a tuple with the DriveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetDriveIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DriveId) {
		return nil, false
	}
	return o.DriveId, true
}

// HasDriveId returns a boolean if a field has been set.
func (o *JobGroup) HasDriveId() bool {
	if o != nil && !IsNil(o.DriveId) {
		return true
	}

	return false
}

// SetDriveId gets a reference to the given int32 and assigns it to the DriveId field.
func (o *JobGroup) SetDriveId(v int32) {
	o.DriveId = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *JobGroup) GetServerId() int32 {
	if o == nil || IsNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetServerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *JobGroup) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *JobGroup) SetServerId(v int32) {
	o.ServerId = &v
}

// GetNetworkDeviceId returns the NetworkDeviceId field value if set, zero value otherwise.
func (o *JobGroup) GetNetworkDeviceId() int32 {
	if o == nil || IsNil(o.NetworkDeviceId) {
		var ret int32
		return ret
	}
	return *o.NetworkDeviceId
}

// GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetNetworkDeviceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkDeviceId) {
		return nil, false
	}
	return o.NetworkDeviceId, true
}

// HasNetworkDeviceId returns a boolean if a field has been set.
func (o *JobGroup) HasNetworkDeviceId() bool {
	if o != nil && !IsNil(o.NetworkDeviceId) {
		return true
	}

	return false
}

// SetNetworkDeviceId gets a reference to the given int32 and assigns it to the NetworkDeviceId field.
func (o *JobGroup) SetNetworkDeviceId(v int32) {
	o.NetworkDeviceId = &v
}

// GetVmPoolId returns the VmPoolId field value if set, zero value otherwise.
func (o *JobGroup) GetVmPoolId() int32 {
	if o == nil || IsNil(o.VmPoolId) {
		var ret int32
		return ret
	}
	return *o.VmPoolId
}

// GetVmPoolIdOk returns a tuple with the VmPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetVmPoolIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VmPoolId) {
		return nil, false
	}
	return o.VmPoolId, true
}

// HasVmPoolId returns a boolean if a field has been set.
func (o *JobGroup) HasVmPoolId() bool {
	if o != nil && !IsNil(o.VmPoolId) {
		return true
	}

	return false
}

// SetVmPoolId gets a reference to the given int32 and assigns it to the VmPoolId field.
func (o *JobGroup) SetVmPoolId(v int32) {
	o.VmPoolId = &v
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *JobGroup) GetStorageId() int32 {
	if o == nil || IsNil(o.StorageId) {
		var ret int32
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobGroup) GetStorageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StorageId) {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *JobGroup) HasStorageId() bool {
	if o != nil && !IsNil(o.StorageId) {
		return true
	}

	return false
}

// SetStorageId gets a reference to the given int32 and assigns it to the StorageId field.
func (o *JobGroup) SetStorageId(v int32) {
	o.StorageId = &v
}

// GetLinks returns the Links field value
func (o *JobGroup) GetLinks() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *JobGroup) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *JobGroup) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o JobGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["description"] = o.Description
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	if !IsNil(o.FinishedTimestamp) {
		toSerialize["finishedTimestamp"] = o.FinishedTimestamp
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.InfrastructureId) {
		toSerialize["infrastructureId"] = o.InfrastructureId
	}
	if !IsNil(o.DriveId) {
		toSerialize["driveId"] = o.DriveId
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.NetworkDeviceId) {
		toSerialize["networkDeviceId"] = o.NetworkDeviceId
	}
	if !IsNil(o.VmPoolId) {
		toSerialize["vmPoolId"] = o.VmPoolId
	}
	if !IsNil(o.StorageId) {
		toSerialize["storageId"] = o.StorageId
	}
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"description",
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

	varJobGroup := _JobGroup{}

	err = json.Unmarshal(data, &varJobGroup)

	if err != nil {
		return err
	}

	*o = JobGroup(varJobGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "context")
		delete(additionalProperties, "description")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "finishedTimestamp")
		delete(additionalProperties, "params")
		delete(additionalProperties, "archived")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "driveId")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "networkDeviceId")
		delete(additionalProperties, "vmPoolId")
		delete(additionalProperties, "storageId")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobGroup struct {
	value *JobGroup
	isSet bool
}

func (v NullableJobGroup) Get() *JobGroup {
	return v.value
}

func (v *NullableJobGroup) Set(val *JobGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableJobGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableJobGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobGroup(val *JobGroup) *NullableJobGroup {
	return &NullableJobGroup{value: val, isSet: true}
}

func (v NullableJobGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



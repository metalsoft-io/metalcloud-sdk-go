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

// checks if the ServerDisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerDisk{}

// ServerDisk struct for ServerDisk
type ServerDisk struct {
	// The id of the disk.
	Id float32 `json:"id"`
	// The id of the server.
	ServerId float32 `json:"serverId"`
	// The model of the disk
	Model string `json:"model"`
	// The size of the disk in GB
	DiskSizeGb float32 `json:"diskSizeGb"`
	// The serial of the disk
	Serial string `json:"serial"`
	// The vendor of the disk
	Vendor string `json:"vendor"`
	// The status of the disk
	Status string `json:"status"`
	// The type of the disk
	Type string `json:"type"`
	// The id of the storage controller
	ServerStorageControllerId *float32 `json:"serverStorageControllerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerDisk ServerDisk

// NewServerDisk instantiates a new ServerDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerDisk(id float32, serverId float32, model string, diskSizeGb float32, serial string, vendor string, status string, type_ string) *ServerDisk {
	this := ServerDisk{}
	this.Id = id
	this.ServerId = serverId
	this.Model = model
	this.DiskSizeGb = diskSizeGb
	this.Serial = serial
	this.Vendor = vendor
	this.Status = status
	this.Type = type_
	return &this
}

// NewServerDiskWithDefaults instantiates a new ServerDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerDiskWithDefaults() *ServerDisk {
	this := ServerDisk{}
	return &this
}

// GetId returns the Id field value
func (o *ServerDisk) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerDisk) SetId(v float32) {
	o.Id = v
}

// GetServerId returns the ServerId field value
func (o *ServerDisk) GetServerId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetServerIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *ServerDisk) SetServerId(v float32) {
	o.ServerId = v
}

// GetModel returns the Model field value
func (o *ServerDisk) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ServerDisk) SetModel(v string) {
	o.Model = v
}

// GetDiskSizeGb returns the DiskSizeGb field value
func (o *ServerDisk) GetDiskSizeGb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskSizeGb
}

// GetDiskSizeGbOk returns a tuple with the DiskSizeGb field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetDiskSizeGbOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskSizeGb, true
}

// SetDiskSizeGb sets field value
func (o *ServerDisk) SetDiskSizeGb(v float32) {
	o.DiskSizeGb = v
}

// GetSerial returns the Serial field value
func (o *ServerDisk) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *ServerDisk) SetSerial(v string) {
	o.Serial = v
}

// GetVendor returns the Vendor field value
func (o *ServerDisk) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *ServerDisk) SetVendor(v string) {
	o.Vendor = v
}

// GetStatus returns the Status field value
func (o *ServerDisk) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ServerDisk) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ServerDisk) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServerDisk) SetType(v string) {
	o.Type = v
}

// GetServerStorageControllerId returns the ServerStorageControllerId field value if set, zero value otherwise.
func (o *ServerDisk) GetServerStorageControllerId() float32 {
	if o == nil || IsNil(o.ServerStorageControllerId) {
		var ret float32
		return ret
	}
	return *o.ServerStorageControllerId
}

// GetServerStorageControllerIdOk returns a tuple with the ServerStorageControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDisk) GetServerStorageControllerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerStorageControllerId) {
		return nil, false
	}
	return o.ServerStorageControllerId, true
}

// HasServerStorageControllerId returns a boolean if a field has been set.
func (o *ServerDisk) HasServerStorageControllerId() bool {
	if o != nil && !IsNil(o.ServerStorageControllerId) {
		return true
	}

	return false
}

// SetServerStorageControllerId gets a reference to the given float32 and assigns it to the ServerStorageControllerId field.
func (o *ServerDisk) SetServerStorageControllerId(v float32) {
	o.ServerStorageControllerId = &v
}

func (o ServerDisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerDisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serverId"] = o.ServerId
	toSerialize["model"] = o.Model
	toSerialize["diskSizeGb"] = o.DiskSizeGb
	toSerialize["serial"] = o.Serial
	toSerialize["vendor"] = o.Vendor
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	if !IsNil(o.ServerStorageControllerId) {
		toSerialize["serverStorageControllerId"] = o.ServerStorageControllerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerDisk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serverId",
		"model",
		"diskSizeGb",
		"serial",
		"vendor",
		"status",
		"type",
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

	varServerDisk := _ServerDisk{}

	err = json.Unmarshal(data, &varServerDisk)

	if err != nil {
		return err
	}

	*o = ServerDisk(varServerDisk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "model")
		delete(additionalProperties, "diskSizeGb")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "serverStorageControllerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerDisk struct {
	value *ServerDisk
	isSet bool
}

func (v NullableServerDisk) Get() *ServerDisk {
	return v.value
}

func (v *NullableServerDisk) Set(val *ServerDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableServerDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableServerDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerDisk(val *ServerDisk) *NullableServerDisk {
	return &NullableServerDisk{value: val, isSet: true}
}

func (v NullableServerDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ServerUnmanagedImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerUnmanagedImport{}

// ServerUnmanagedImport struct for ServerUnmanagedImport
type ServerUnmanagedImport struct {
	// The management address of the server.
	ManagementAddress *string `json:"managementAddress,omitempty"`
	// The site id where the server is located.
	SiteId float32 `json:"siteId"`
	// The server type id.
	ServerTypeId float32 `json:"serverTypeId"`
	// Flag to indicate if the server supports SOL.
	ServerSupportsOobProvisioning *float32 `json:"serverSupportsOobProvisioning,omitempty"`
	// The interfaces of the server.
	ServerInterfaces []ServerUnmanagedImportInternalInterface `json:"serverInterfaces"`
	// The server Serial Number.
	ServerSerialNumber *string `json:"serverSerialNumber,omitempty"`
	// The server UUID.
	ServerUUID *string `json:"serverUUID,omitempty"`
	// The server hostname.
	Hostname *string `json:"hostname,omitempty"`
	// The server IPMI username.
	ServerIpmiUsername *string `json:"serverIpmiUsername,omitempty"`
	// The server IPMI password.
	ServerIpmiPassword *string `json:"serverIpmiPassword,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerUnmanagedImport ServerUnmanagedImport

// NewServerUnmanagedImport instantiates a new ServerUnmanagedImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerUnmanagedImport(siteId float32, serverTypeId float32, serverInterfaces []ServerUnmanagedImportInternalInterface) *ServerUnmanagedImport {
	this := ServerUnmanagedImport{}
	this.SiteId = siteId
	this.ServerTypeId = serverTypeId
	this.ServerInterfaces = serverInterfaces
	return &this
}

// NewServerUnmanagedImportWithDefaults instantiates a new ServerUnmanagedImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerUnmanagedImportWithDefaults() *ServerUnmanagedImport {
	this := ServerUnmanagedImport{}
	return &this
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetManagementAddress() string {
	if o == nil || IsNil(o.ManagementAddress) {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetManagementAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementAddress) {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasManagementAddress() bool {
	if o != nil && !IsNil(o.ManagementAddress) {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *ServerUnmanagedImport) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetSiteId returns the SiteId field value
func (o *ServerUnmanagedImport) GetSiteId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetSiteIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *ServerUnmanagedImport) SetSiteId(v float32) {
	o.SiteId = v
}

// GetServerTypeId returns the ServerTypeId field value
func (o *ServerUnmanagedImport) GetServerTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerTypeId
}

// GetServerTypeIdOk returns a tuple with the ServerTypeId field value
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerTypeId, true
}

// SetServerTypeId sets field value
func (o *ServerUnmanagedImport) SetServerTypeId(v float32) {
	o.ServerTypeId = v
}

// GetServerSupportsOobProvisioning returns the ServerSupportsOobProvisioning field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetServerSupportsOobProvisioning() float32 {
	if o == nil || IsNil(o.ServerSupportsOobProvisioning) {
		var ret float32
		return ret
	}
	return *o.ServerSupportsOobProvisioning
}

// GetServerSupportsOobProvisioningOk returns a tuple with the ServerSupportsOobProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerSupportsOobProvisioningOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerSupportsOobProvisioning) {
		return nil, false
	}
	return o.ServerSupportsOobProvisioning, true
}

// HasServerSupportsOobProvisioning returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasServerSupportsOobProvisioning() bool {
	if o != nil && !IsNil(o.ServerSupportsOobProvisioning) {
		return true
	}

	return false
}

// SetServerSupportsOobProvisioning gets a reference to the given float32 and assigns it to the ServerSupportsOobProvisioning field.
func (o *ServerUnmanagedImport) SetServerSupportsOobProvisioning(v float32) {
	o.ServerSupportsOobProvisioning = &v
}

// GetServerInterfaces returns the ServerInterfaces field value
func (o *ServerUnmanagedImport) GetServerInterfaces() []ServerUnmanagedImportInternalInterface {
	if o == nil {
		var ret []ServerUnmanagedImportInternalInterface
		return ret
	}

	return o.ServerInterfaces
}

// GetServerInterfacesOk returns a tuple with the ServerInterfaces field value
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerInterfacesOk() ([]ServerUnmanagedImportInternalInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerInterfaces, true
}

// SetServerInterfaces sets field value
func (o *ServerUnmanagedImport) SetServerInterfaces(v []ServerUnmanagedImportInternalInterface) {
	o.ServerInterfaces = v
}

// GetServerSerialNumber returns the ServerSerialNumber field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetServerSerialNumber() string {
	if o == nil || IsNil(o.ServerSerialNumber) {
		var ret string
		return ret
	}
	return *o.ServerSerialNumber
}

// GetServerSerialNumberOk returns a tuple with the ServerSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ServerSerialNumber) {
		return nil, false
	}
	return o.ServerSerialNumber, true
}

// HasServerSerialNumber returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasServerSerialNumber() bool {
	if o != nil && !IsNil(o.ServerSerialNumber) {
		return true
	}

	return false
}

// SetServerSerialNumber gets a reference to the given string and assigns it to the ServerSerialNumber field.
func (o *ServerUnmanagedImport) SetServerSerialNumber(v string) {
	o.ServerSerialNumber = &v
}

// GetServerUUID returns the ServerUUID field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetServerUUID() string {
	if o == nil || IsNil(o.ServerUUID) {
		var ret string
		return ret
	}
	return *o.ServerUUID
}

// GetServerUUIDOk returns a tuple with the ServerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServerUUID) {
		return nil, false
	}
	return o.ServerUUID, true
}

// HasServerUUID returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasServerUUID() bool {
	if o != nil && !IsNil(o.ServerUUID) {
		return true
	}

	return false
}

// SetServerUUID gets a reference to the given string and assigns it to the ServerUUID field.
func (o *ServerUnmanagedImport) SetServerUUID(v string) {
	o.ServerUUID = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ServerUnmanagedImport) SetHostname(v string) {
	o.Hostname = &v
}

// GetServerIpmiUsername returns the ServerIpmiUsername field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetServerIpmiUsername() string {
	if o == nil || IsNil(o.ServerIpmiUsername) {
		var ret string
		return ret
	}
	return *o.ServerIpmiUsername
}

// GetServerIpmiUsernameOk returns a tuple with the ServerIpmiUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerIpmiUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerIpmiUsername) {
		return nil, false
	}
	return o.ServerIpmiUsername, true
}

// HasServerIpmiUsername returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasServerIpmiUsername() bool {
	if o != nil && !IsNil(o.ServerIpmiUsername) {
		return true
	}

	return false
}

// SetServerIpmiUsername gets a reference to the given string and assigns it to the ServerIpmiUsername field.
func (o *ServerUnmanagedImport) SetServerIpmiUsername(v string) {
	o.ServerIpmiUsername = &v
}

// GetServerIpmiPassword returns the ServerIpmiPassword field value if set, zero value otherwise.
func (o *ServerUnmanagedImport) GetServerIpmiPassword() string {
	if o == nil || IsNil(o.ServerIpmiPassword) {
		var ret string
		return ret
	}
	return *o.ServerIpmiPassword
}

// GetServerIpmiPasswordOk returns a tuple with the ServerIpmiPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUnmanagedImport) GetServerIpmiPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ServerIpmiPassword) {
		return nil, false
	}
	return o.ServerIpmiPassword, true
}

// HasServerIpmiPassword returns a boolean if a field has been set.
func (o *ServerUnmanagedImport) HasServerIpmiPassword() bool {
	if o != nil && !IsNil(o.ServerIpmiPassword) {
		return true
	}

	return false
}

// SetServerIpmiPassword gets a reference to the given string and assigns it to the ServerIpmiPassword field.
func (o *ServerUnmanagedImport) SetServerIpmiPassword(v string) {
	o.ServerIpmiPassword = &v
}

func (o ServerUnmanagedImport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerUnmanagedImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagementAddress) {
		toSerialize["managementAddress"] = o.ManagementAddress
	}
	toSerialize["siteId"] = o.SiteId
	toSerialize["serverTypeId"] = o.ServerTypeId
	if !IsNil(o.ServerSupportsOobProvisioning) {
		toSerialize["serverSupportsOobProvisioning"] = o.ServerSupportsOobProvisioning
	}
	toSerialize["serverInterfaces"] = o.ServerInterfaces
	if !IsNil(o.ServerSerialNumber) {
		toSerialize["serverSerialNumber"] = o.ServerSerialNumber
	}
	if !IsNil(o.ServerUUID) {
		toSerialize["serverUUID"] = o.ServerUUID
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.ServerIpmiUsername) {
		toSerialize["serverIpmiUsername"] = o.ServerIpmiUsername
	}
	if !IsNil(o.ServerIpmiPassword) {
		toSerialize["serverIpmiPassword"] = o.ServerIpmiPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerUnmanagedImport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"siteId",
		"serverTypeId",
		"serverInterfaces",
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

	varServerUnmanagedImport := _ServerUnmanagedImport{}

	err = json.Unmarshal(data, &varServerUnmanagedImport)

	if err != nil {
		return err
	}

	*o = ServerUnmanagedImport(varServerUnmanagedImport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "managementAddress")
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "serverTypeId")
		delete(additionalProperties, "serverSupportsOobProvisioning")
		delete(additionalProperties, "serverInterfaces")
		delete(additionalProperties, "serverSerialNumber")
		delete(additionalProperties, "serverUUID")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "serverIpmiUsername")
		delete(additionalProperties, "serverIpmiPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerUnmanagedImport struct {
	value *ServerUnmanagedImport
	isSet bool
}

func (v NullableServerUnmanagedImport) Get() *ServerUnmanagedImport {
	return v.value
}

func (v *NullableServerUnmanagedImport) Set(val *ServerUnmanagedImport) {
	v.value = val
	v.isSet = true
}

func (v NullableServerUnmanagedImport) IsSet() bool {
	return v.isSet
}

func (v *NullableServerUnmanagedImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerUnmanagedImport(val *ServerUnmanagedImport) *NullableServerUnmanagedImport {
	return &NullableServerUnmanagedImport{value: val, isSet: true}
}

func (v NullableServerUnmanagedImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerUnmanagedImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



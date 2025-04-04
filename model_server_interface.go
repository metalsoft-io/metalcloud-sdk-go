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

// checks if the ServerInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInterface{}

// ServerInterface struct for ServerInterface
type ServerInterface struct {
	// The id of the server interface.
	ServerInterfaceId float32 `json:"serverInterfaceId"`
	// The id of the server.
	ServerId float32 `json:"serverId"`
	// The MAC address of the server interface.
	MacAddress string `json:"macAddress"`
	// The index of the server interface.
	Index float32 `json:"index"`
	// The switch port id of the server interface.
	SwitchPortId *string `json:"switchPortId,omitempty"`
	// The switch hostname of the server interface.
	SwitchHostname *string `json:"switchHostname,omitempty"`
	// The capacity in Mbps of the server interface.
	CapacityMbps *float32 `json:"capacityMbps,omitempty"`
	// The add-on MAC address of the server interface.
	AddOnMac *string `json:"addOnMac,omitempty"`
	// The add-on type of the server interface.
	AddOnType *string `json:"addOnType,omitempty"`
	// The add-on info of the server interface.
	AddOnInfo map[string]interface{} `json:"addOnInfo,omitempty"`
	// Flag to indicate if PXE is enabled.
	PxeEnabled *float32 `json:"pxeEnabled,omitempty"`
	// Flag to indicate if the server interface supports iSCSI boot
	SupportsIscsiBoot *float32 `json:"supportsIscsiBoot,omitempty"`
	// The WWPN of the fibre channel.
	FibreChannelWwpn *string `json:"fibreChannelWwpn,omitempty"`
	// The description of the server interface.
	Description *string `json:"description,omitempty"`
	// The alias index of the server interface.
	AliasIndex *float32 `json:"aliasIndex,omitempty"`
	// The network device linked to the server.
	NetworkDevice map[string]interface{} `json:"networkDevice,omitempty"`
	// The network device interface linked to the server.
	NetworkDeviceInterface map[string]interface{} `json:"networkDeviceInterface,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInterface ServerInterface

// NewServerInterface instantiates a new ServerInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInterface(serverInterfaceId float32, serverId float32, macAddress string, index float32) *ServerInterface {
	this := ServerInterface{}
	this.ServerInterfaceId = serverInterfaceId
	this.ServerId = serverId
	this.MacAddress = macAddress
	this.Index = index
	return &this
}

// NewServerInterfaceWithDefaults instantiates a new ServerInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInterfaceWithDefaults() *ServerInterface {
	this := ServerInterface{}
	return &this
}

// GetServerInterfaceId returns the ServerInterfaceId field value
func (o *ServerInterface) GetServerInterfaceId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerInterfaceId
}

// GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field value
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetServerInterfaceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerInterfaceId, true
}

// SetServerInterfaceId sets field value
func (o *ServerInterface) SetServerInterfaceId(v float32) {
	o.ServerInterfaceId = v
}

// GetServerId returns the ServerId field value
func (o *ServerInterface) GetServerId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetServerIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *ServerInterface) SetServerId(v float32) {
	o.ServerId = v
}

// GetMacAddress returns the MacAddress field value
func (o *ServerInterface) GetMacAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacAddress, true
}

// SetMacAddress sets field value
func (o *ServerInterface) SetMacAddress(v string) {
	o.MacAddress = v
}

// GetIndex returns the Index field value
func (o *ServerInterface) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ServerInterface) SetIndex(v float32) {
	o.Index = v
}

// GetSwitchPortId returns the SwitchPortId field value if set, zero value otherwise.
func (o *ServerInterface) GetSwitchPortId() string {
	if o == nil || IsNil(o.SwitchPortId) {
		var ret string
		return ret
	}
	return *o.SwitchPortId
}

// GetSwitchPortIdOk returns a tuple with the SwitchPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetSwitchPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchPortId) {
		return nil, false
	}
	return o.SwitchPortId, true
}

// HasSwitchPortId returns a boolean if a field has been set.
func (o *ServerInterface) HasSwitchPortId() bool {
	if o != nil && !IsNil(o.SwitchPortId) {
		return true
	}

	return false
}

// SetSwitchPortId gets a reference to the given string and assigns it to the SwitchPortId field.
func (o *ServerInterface) SetSwitchPortId(v string) {
	o.SwitchPortId = &v
}

// GetSwitchHostname returns the SwitchHostname field value if set, zero value otherwise.
func (o *ServerInterface) GetSwitchHostname() string {
	if o == nil || IsNil(o.SwitchHostname) {
		var ret string
		return ret
	}
	return *o.SwitchHostname
}

// GetSwitchHostnameOk returns a tuple with the SwitchHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetSwitchHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchHostname) {
		return nil, false
	}
	return o.SwitchHostname, true
}

// HasSwitchHostname returns a boolean if a field has been set.
func (o *ServerInterface) HasSwitchHostname() bool {
	if o != nil && !IsNil(o.SwitchHostname) {
		return true
	}

	return false
}

// SetSwitchHostname gets a reference to the given string and assigns it to the SwitchHostname field.
func (o *ServerInterface) SetSwitchHostname(v string) {
	o.SwitchHostname = &v
}

// GetCapacityMbps returns the CapacityMbps field value if set, zero value otherwise.
func (o *ServerInterface) GetCapacityMbps() float32 {
	if o == nil || IsNil(o.CapacityMbps) {
		var ret float32
		return ret
	}
	return *o.CapacityMbps
}

// GetCapacityMbpsOk returns a tuple with the CapacityMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetCapacityMbpsOk() (*float32, bool) {
	if o == nil || IsNil(o.CapacityMbps) {
		return nil, false
	}
	return o.CapacityMbps, true
}

// HasCapacityMbps returns a boolean if a field has been set.
func (o *ServerInterface) HasCapacityMbps() bool {
	if o != nil && !IsNil(o.CapacityMbps) {
		return true
	}

	return false
}

// SetCapacityMbps gets a reference to the given float32 and assigns it to the CapacityMbps field.
func (o *ServerInterface) SetCapacityMbps(v float32) {
	o.CapacityMbps = &v
}

// GetAddOnMac returns the AddOnMac field value if set, zero value otherwise.
func (o *ServerInterface) GetAddOnMac() string {
	if o == nil || IsNil(o.AddOnMac) {
		var ret string
		return ret
	}
	return *o.AddOnMac
}

// GetAddOnMacOk returns a tuple with the AddOnMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetAddOnMacOk() (*string, bool) {
	if o == nil || IsNil(o.AddOnMac) {
		return nil, false
	}
	return o.AddOnMac, true
}

// HasAddOnMac returns a boolean if a field has been set.
func (o *ServerInterface) HasAddOnMac() bool {
	if o != nil && !IsNil(o.AddOnMac) {
		return true
	}

	return false
}

// SetAddOnMac gets a reference to the given string and assigns it to the AddOnMac field.
func (o *ServerInterface) SetAddOnMac(v string) {
	o.AddOnMac = &v
}

// GetAddOnType returns the AddOnType field value if set, zero value otherwise.
func (o *ServerInterface) GetAddOnType() string {
	if o == nil || IsNil(o.AddOnType) {
		var ret string
		return ret
	}
	return *o.AddOnType
}

// GetAddOnTypeOk returns a tuple with the AddOnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetAddOnTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AddOnType) {
		return nil, false
	}
	return o.AddOnType, true
}

// HasAddOnType returns a boolean if a field has been set.
func (o *ServerInterface) HasAddOnType() bool {
	if o != nil && !IsNil(o.AddOnType) {
		return true
	}

	return false
}

// SetAddOnType gets a reference to the given string and assigns it to the AddOnType field.
func (o *ServerInterface) SetAddOnType(v string) {
	o.AddOnType = &v
}

// GetAddOnInfo returns the AddOnInfo field value if set, zero value otherwise.
func (o *ServerInterface) GetAddOnInfo() map[string]interface{} {
	if o == nil || IsNil(o.AddOnInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AddOnInfo
}

// GetAddOnInfoOk returns a tuple with the AddOnInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetAddOnInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AddOnInfo) {
		return map[string]interface{}{}, false
	}
	return o.AddOnInfo, true
}

// HasAddOnInfo returns a boolean if a field has been set.
func (o *ServerInterface) HasAddOnInfo() bool {
	if o != nil && !IsNil(o.AddOnInfo) {
		return true
	}

	return false
}

// SetAddOnInfo gets a reference to the given map[string]interface{} and assigns it to the AddOnInfo field.
func (o *ServerInterface) SetAddOnInfo(v map[string]interface{}) {
	o.AddOnInfo = v
}

// GetPxeEnabled returns the PxeEnabled field value if set, zero value otherwise.
func (o *ServerInterface) GetPxeEnabled() float32 {
	if o == nil || IsNil(o.PxeEnabled) {
		var ret float32
		return ret
	}
	return *o.PxeEnabled
}

// GetPxeEnabledOk returns a tuple with the PxeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetPxeEnabledOk() (*float32, bool) {
	if o == nil || IsNil(o.PxeEnabled) {
		return nil, false
	}
	return o.PxeEnabled, true
}

// HasPxeEnabled returns a boolean if a field has been set.
func (o *ServerInterface) HasPxeEnabled() bool {
	if o != nil && !IsNil(o.PxeEnabled) {
		return true
	}

	return false
}

// SetPxeEnabled gets a reference to the given float32 and assigns it to the PxeEnabled field.
func (o *ServerInterface) SetPxeEnabled(v float32) {
	o.PxeEnabled = &v
}

// GetSupportsIscsiBoot returns the SupportsIscsiBoot field value if set, zero value otherwise.
func (o *ServerInterface) GetSupportsIscsiBoot() float32 {
	if o == nil || IsNil(o.SupportsIscsiBoot) {
		var ret float32
		return ret
	}
	return *o.SupportsIscsiBoot
}

// GetSupportsIscsiBootOk returns a tuple with the SupportsIscsiBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetSupportsIscsiBootOk() (*float32, bool) {
	if o == nil || IsNil(o.SupportsIscsiBoot) {
		return nil, false
	}
	return o.SupportsIscsiBoot, true
}

// HasSupportsIscsiBoot returns a boolean if a field has been set.
func (o *ServerInterface) HasSupportsIscsiBoot() bool {
	if o != nil && !IsNil(o.SupportsIscsiBoot) {
		return true
	}

	return false
}

// SetSupportsIscsiBoot gets a reference to the given float32 and assigns it to the SupportsIscsiBoot field.
func (o *ServerInterface) SetSupportsIscsiBoot(v float32) {
	o.SupportsIscsiBoot = &v
}

// GetFibreChannelWwpn returns the FibreChannelWwpn field value if set, zero value otherwise.
func (o *ServerInterface) GetFibreChannelWwpn() string {
	if o == nil || IsNil(o.FibreChannelWwpn) {
		var ret string
		return ret
	}
	return *o.FibreChannelWwpn
}

// GetFibreChannelWwpnOk returns a tuple with the FibreChannelWwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetFibreChannelWwpnOk() (*string, bool) {
	if o == nil || IsNil(o.FibreChannelWwpn) {
		return nil, false
	}
	return o.FibreChannelWwpn, true
}

// HasFibreChannelWwpn returns a boolean if a field has been set.
func (o *ServerInterface) HasFibreChannelWwpn() bool {
	if o != nil && !IsNil(o.FibreChannelWwpn) {
		return true
	}

	return false
}

// SetFibreChannelWwpn gets a reference to the given string and assigns it to the FibreChannelWwpn field.
func (o *ServerInterface) SetFibreChannelWwpn(v string) {
	o.FibreChannelWwpn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerInterface) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerInterface) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerInterface) SetDescription(v string) {
	o.Description = &v
}

// GetAliasIndex returns the AliasIndex field value if set, zero value otherwise.
func (o *ServerInterface) GetAliasIndex() float32 {
	if o == nil || IsNil(o.AliasIndex) {
		var ret float32
		return ret
	}
	return *o.AliasIndex
}

// GetAliasIndexOk returns a tuple with the AliasIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetAliasIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.AliasIndex) {
		return nil, false
	}
	return o.AliasIndex, true
}

// HasAliasIndex returns a boolean if a field has been set.
func (o *ServerInterface) HasAliasIndex() bool {
	if o != nil && !IsNil(o.AliasIndex) {
		return true
	}

	return false
}

// SetAliasIndex gets a reference to the given float32 and assigns it to the AliasIndex field.
func (o *ServerInterface) SetAliasIndex(v float32) {
	o.AliasIndex = &v
}

// GetNetworkDevice returns the NetworkDevice field value if set, zero value otherwise.
func (o *ServerInterface) GetNetworkDevice() map[string]interface{} {
	if o == nil || IsNil(o.NetworkDevice) {
		var ret map[string]interface{}
		return ret
	}
	return o.NetworkDevice
}

// GetNetworkDeviceOk returns a tuple with the NetworkDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetNetworkDeviceOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NetworkDevice) {
		return map[string]interface{}{}, false
	}
	return o.NetworkDevice, true
}

// HasNetworkDevice returns a boolean if a field has been set.
func (o *ServerInterface) HasNetworkDevice() bool {
	if o != nil && !IsNil(o.NetworkDevice) {
		return true
	}

	return false
}

// SetNetworkDevice gets a reference to the given map[string]interface{} and assigns it to the NetworkDevice field.
func (o *ServerInterface) SetNetworkDevice(v map[string]interface{}) {
	o.NetworkDevice = v
}

// GetNetworkDeviceInterface returns the NetworkDeviceInterface field value if set, zero value otherwise.
func (o *ServerInterface) GetNetworkDeviceInterface() map[string]interface{} {
	if o == nil || IsNil(o.NetworkDeviceInterface) {
		var ret map[string]interface{}
		return ret
	}
	return o.NetworkDeviceInterface
}

// GetNetworkDeviceInterfaceOk returns a tuple with the NetworkDeviceInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInterface) GetNetworkDeviceInterfaceOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NetworkDeviceInterface) {
		return map[string]interface{}{}, false
	}
	return o.NetworkDeviceInterface, true
}

// HasNetworkDeviceInterface returns a boolean if a field has been set.
func (o *ServerInterface) HasNetworkDeviceInterface() bool {
	if o != nil && !IsNil(o.NetworkDeviceInterface) {
		return true
	}

	return false
}

// SetNetworkDeviceInterface gets a reference to the given map[string]interface{} and assigns it to the NetworkDeviceInterface field.
func (o *ServerInterface) SetNetworkDeviceInterface(v map[string]interface{}) {
	o.NetworkDeviceInterface = v
}

func (o ServerInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverInterfaceId"] = o.ServerInterfaceId
	toSerialize["serverId"] = o.ServerId
	toSerialize["macAddress"] = o.MacAddress
	toSerialize["index"] = o.Index
	if !IsNil(o.SwitchPortId) {
		toSerialize["switchPortId"] = o.SwitchPortId
	}
	if !IsNil(o.SwitchHostname) {
		toSerialize["switchHostname"] = o.SwitchHostname
	}
	if !IsNil(o.CapacityMbps) {
		toSerialize["capacityMbps"] = o.CapacityMbps
	}
	if !IsNil(o.AddOnMac) {
		toSerialize["addOnMac"] = o.AddOnMac
	}
	if !IsNil(o.AddOnType) {
		toSerialize["addOnType"] = o.AddOnType
	}
	if !IsNil(o.AddOnInfo) {
		toSerialize["addOnInfo"] = o.AddOnInfo
	}
	if !IsNil(o.PxeEnabled) {
		toSerialize["pxeEnabled"] = o.PxeEnabled
	}
	if !IsNil(o.SupportsIscsiBoot) {
		toSerialize["supportsIscsiBoot"] = o.SupportsIscsiBoot
	}
	if !IsNil(o.FibreChannelWwpn) {
		toSerialize["fibreChannelWwpn"] = o.FibreChannelWwpn
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AliasIndex) {
		toSerialize["aliasIndex"] = o.AliasIndex
	}
	if !IsNil(o.NetworkDevice) {
		toSerialize["networkDevice"] = o.NetworkDevice
	}
	if !IsNil(o.NetworkDeviceInterface) {
		toSerialize["networkDeviceInterface"] = o.NetworkDeviceInterface
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverInterfaceId",
		"serverId",
		"macAddress",
		"index",
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

	varServerInterface := _ServerInterface{}

	err = json.Unmarshal(data, &varServerInterface)

	if err != nil {
		return err
	}

	*o = ServerInterface(varServerInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverInterfaceId")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "macAddress")
		delete(additionalProperties, "index")
		delete(additionalProperties, "switchPortId")
		delete(additionalProperties, "switchHostname")
		delete(additionalProperties, "capacityMbps")
		delete(additionalProperties, "addOnMac")
		delete(additionalProperties, "addOnType")
		delete(additionalProperties, "addOnInfo")
		delete(additionalProperties, "pxeEnabled")
		delete(additionalProperties, "supportsIscsiBoot")
		delete(additionalProperties, "fibreChannelWwpn")
		delete(additionalProperties, "description")
		delete(additionalProperties, "aliasIndex")
		delete(additionalProperties, "networkDevice")
		delete(additionalProperties, "networkDeviceInterface")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInterface struct {
	value *ServerInterface
	isSet bool
}

func (v NullableServerInterface) Get() *ServerInterface {
	return v.value
}

func (v *NullableServerInterface) Set(val *ServerInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInterface(val *ServerInterface) *NullableServerInterface {
	return &NullableServerInterface{value: val, isSet: true}
}

func (v NullableServerInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



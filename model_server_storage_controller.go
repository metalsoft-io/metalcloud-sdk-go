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

// checks if the ServerStorageController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStorageController{}

// ServerStorageController struct for ServerStorageController
type ServerStorageController struct {
	// The id of the storage controller.
	Id float32 `json:"id"`
	// The id of the server.
	ServerId float32 `json:"serverId"`
	// The name of the storage controller.
	Name string `json:"name"`
	// The label of the storage controller.
	Label string `json:"label"`
	// The description of the storage controller.
	Description string `json:"description"`
	// The options of the storage controller.
	Options ServerTypeStorageControllerOptions `json:"options"`
	// The mode of the storage controller.
	Mode string `json:"mode"`
	AdditionalProperties map[string]interface{}
}

type _ServerStorageController ServerStorageController

// NewServerStorageController instantiates a new ServerStorageController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStorageController(id float32, serverId float32, name string, label string, description string, options ServerTypeStorageControllerOptions, mode string) *ServerStorageController {
	this := ServerStorageController{}
	this.Id = id
	this.ServerId = serverId
	this.Name = name
	this.Label = label
	this.Description = description
	this.Options = options
	this.Mode = mode
	return &this
}

// NewServerStorageControllerWithDefaults instantiates a new ServerStorageController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStorageControllerWithDefaults() *ServerStorageController {
	this := ServerStorageController{}
	return &this
}

// GetId returns the Id field value
func (o *ServerStorageController) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerStorageController) SetId(v float32) {
	o.Id = v
}

// GetServerId returns the ServerId field value
func (o *ServerStorageController) GetServerId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetServerIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *ServerStorageController) SetServerId(v float32) {
	o.ServerId = v
}

// GetName returns the Name field value
func (o *ServerStorageController) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerStorageController) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *ServerStorageController) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ServerStorageController) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value
func (o *ServerStorageController) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServerStorageController) SetDescription(v string) {
	o.Description = v
}

// GetOptions returns the Options field value
func (o *ServerStorageController) GetOptions() ServerTypeStorageControllerOptions {
	if o == nil {
		var ret ServerTypeStorageControllerOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetOptionsOk() (*ServerTypeStorageControllerOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ServerStorageController) SetOptions(v ServerTypeStorageControllerOptions) {
	o.Options = v
}

// GetMode returns the Mode field value
func (o *ServerStorageController) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ServerStorageController) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *ServerStorageController) SetMode(v string) {
	o.Mode = v
}

func (o ServerStorageController) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStorageController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["serverId"] = o.ServerId
	toSerialize["name"] = o.Name
	toSerialize["label"] = o.Label
	toSerialize["description"] = o.Description
	toSerialize["options"] = o.Options
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerStorageController) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"serverId",
		"name",
		"label",
		"description",
		"options",
		"mode",
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

	varServerStorageController := _ServerStorageController{}

	err = json.Unmarshal(data, &varServerStorageController)

	if err != nil {
		return err
	}

	*o = ServerStorageController(varServerStorageController)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "options")
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerStorageController struct {
	value *ServerStorageController
	isSet bool
}

func (v NullableServerStorageController) Get() *ServerStorageController {
	return v.value
}

func (v *NullableServerStorageController) Set(val *ServerStorageController) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStorageController) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStorageController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStorageController(val *ServerStorageController) *NullableServerStorageController {
	return &NullableServerStorageController{value: val, isSet: true}
}

func (v NullableServerStorageController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStorageController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



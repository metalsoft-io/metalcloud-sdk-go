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

// checks if the ServerInstanceStorageProfileController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceStorageProfileController{}

// ServerInstanceStorageProfileController struct for ServerInstanceStorageProfileController
type ServerInstanceStorageProfileController struct {
	// The id of the Storage Controller.
	Id string `json:"id"`
	// The mode of the Storage Controller.
	Mode string `json:"mode"`
	// The information for each Volume of the Storage Controller.
	Volumes []ServerInstanceStorageProfileControllerVolume `json:"volumes"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceStorageProfileController ServerInstanceStorageProfileController

// NewServerInstanceStorageProfileController instantiates a new ServerInstanceStorageProfileController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceStorageProfileController(id string, mode string, volumes []ServerInstanceStorageProfileControllerVolume) *ServerInstanceStorageProfileController {
	this := ServerInstanceStorageProfileController{}
	this.Id = id
	this.Mode = mode
	this.Volumes = volumes
	return &this
}

// NewServerInstanceStorageProfileControllerWithDefaults instantiates a new ServerInstanceStorageProfileController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceStorageProfileControllerWithDefaults() *ServerInstanceStorageProfileController {
	this := ServerInstanceStorageProfileController{}
	return &this
}

// GetId returns the Id field value
func (o *ServerInstanceStorageProfileController) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceStorageProfileController) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerInstanceStorageProfileController) SetId(v string) {
	o.Id = v
}

// GetMode returns the Mode field value
func (o *ServerInstanceStorageProfileController) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceStorageProfileController) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *ServerInstanceStorageProfileController) SetMode(v string) {
	o.Mode = v
}

// GetVolumes returns the Volumes field value
func (o *ServerInstanceStorageProfileController) GetVolumes() []ServerInstanceStorageProfileControllerVolume {
	if o == nil {
		var ret []ServerInstanceStorageProfileControllerVolume
		return ret
	}

	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceStorageProfileController) GetVolumesOk() ([]ServerInstanceStorageProfileControllerVolume, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volumes, true
}

// SetVolumes sets field value
func (o *ServerInstanceStorageProfileController) SetVolumes(v []ServerInstanceStorageProfileControllerVolume) {
	o.Volumes = v
}

func (o ServerInstanceStorageProfileController) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceStorageProfileController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["mode"] = o.Mode
	toSerialize["volumes"] = o.Volumes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceStorageProfileController) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"mode",
		"volumes",
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

	varServerInstanceStorageProfileController := _ServerInstanceStorageProfileController{}

	err = json.Unmarshal(data, &varServerInstanceStorageProfileController)

	if err != nil {
		return err
	}

	*o = ServerInstanceStorageProfileController(varServerInstanceStorageProfileController)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "volumes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceStorageProfileController struct {
	value *ServerInstanceStorageProfileController
	isSet bool
}

func (v NullableServerInstanceStorageProfileController) Get() *ServerInstanceStorageProfileController {
	return v.value
}

func (v *NullableServerInstanceStorageProfileController) Set(val *ServerInstanceStorageProfileController) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceStorageProfileController) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceStorageProfileController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceStorageProfileController(val *ServerInstanceStorageProfileController) *NullableServerInstanceStorageProfileController {
	return &NullableServerInstanceStorageProfileController{value: val, isSet: true}
}

func (v NullableServerInstanceStorageProfileController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceStorageProfileController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


